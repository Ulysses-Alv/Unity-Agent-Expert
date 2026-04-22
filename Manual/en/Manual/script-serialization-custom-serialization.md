# Custom serialization

When you want to serialize something that Unity’s serializer doesn’t support (for example, a C# Dictionary) you can implement the [ISerializationCallbackReceiver](../ScriptReference/ISerializationCallbackReceiver.md) interface in your class. This allows you to implement callbacks that Unity invokes at key points during serialization and deserialization.

You can use serialization callbacks to give your hard-to-serialize data a different representation at runtime than at serialization time. You can transform your data into something Unity understands just before Unity serializes it. After Unity has written the data to your fields, you can transform the serialized data back into the form you want it to have at runtime.

1. When an object is about to be serialized, Unity invokes the `OnBeforeSerialize()` callback. In this callback you can transform your data into something Unity understands. For example, to serialize a C# Dictionary, copy the data from the Dictionary into an array of keys and an array of values.
2. After the `OnBeforeSerialize()` callback is complete, Unity serializes the arrays.
3. Later, when the object is deserialized, Unity invokes the `OnAfterDeserialize()` callback. In this callback you can transform the data back into a form that’s convenient for the object in memory. For example, use the key and value arrays to repopulate the C# Dictionary.

## Performance issues with Unity’s default serialization

If you want to have a tree data structure and you let Unity serialize the data structure directly, the missing support for null limitation causes your data stream to become large. This leads to performance degradations in many systems:

```
using UnityEngine;
using System.Collections.Generic;
using System;

// This example demonstrates why directly serializing a tree
// can lead to performance issues due to Unity's serializer limitations (e.g., no null support, deep copying).
// It's provided to illustrate a problematic pattern, not a recommended one.
public class ProblematicTreeSerialization : MonoBehaviour
{
    [Serializable]
    public class Node
    {
        public string interestingValue = "value";
        // This field creates a class cycle (Node -> List<Node> -> Node),
        // which can make serialization data very large and inefficient for Unity's default serializer.
        public List<Node> children = new List<Node>();
    }

    // This root node will be serialized by Unity.
    public Node root = new Node();

}
```

To resolve this issue, you can tell Unity not to serialize the tree directly and create a separate serializable representation of the tree that Unity understands. You can then use the `ISerializationCallbackReceiver` interface to convert between your runtime tree and this serializable format.

```
using UnityEngine;
using System.Collections.Generic;
using System;

// This example demonstrates how to use ISerializationCallbackReceiver to custom serialize
// a tree structure, avoiding the performance issues of direct serialization shown in ProblematicTreeSerialization.
public class CustomTreeSerialization : MonoBehaviour, ISerializationCallbackReceiver
{
    // Node class used at runtime. This is not directly serialized by Unity.
    public class Node
    {
        public string interestingValue = "value";
        public List<Node> children = new List<Node>();
    }

    // Serializable struct that represents a node for Unity's serializer.
    [Serializable]
    public struct SerializableNode
    {
        public string interestingValue;
        public int childCount;
        public int indexOfFirstChild;
    }

    // The root node used for runtime tree representation. Not directly serialized by Unity.
    private Node root = new Node();

    // This list is the only data Unity will serialize for the tree.
    // It's marked [SerializeField] so Unity's serializer can access it.
    [SerializeField] private List<SerializableNode> serializedNodes = new List<SerializableNode>();

    // Called just before Unity serializes the object.
    public void OnBeforeSerialize()
    {
        // Ensure root exists for serialization.
        if (root == null)
        {
            root = new Node();
        }

        serializedNodes.Clear(); 
        AddNodeToSerializedNodes(root); 
    }

    // Recursively adds nodes to the 'serializedNodes' list in depth-first order.
    private void AddNodeToSerializedNodes(Node n)
    {
        var serializedNode = new SerializableNode
        {
            interestingValue = n.interestingValue,
            childCount = n.children.Count,
            // The index of the first child will be the current size of 'serializedNodes' + 1
            // (since the current node is added next, and then its children).
            indexOfFirstChild = serializedNodes.Count + 1
        };

        serializedNodes.Add(serializedNode);

        foreach (var child in n.children)
        {
            AddNodeToSerializedNodes(child);
        }
    }

    // Called just after Unity deserializes the object.
    public void OnAfterDeserialize()
    {
        // If there's serialized data, reconstruct the runtime tree.
        if (serializedNodes != null && serializedNodes.Count > 0)
        {
            // Start reconstruction from the first serialized node (index 0).
            ReadNodeFromSerializedNodes(0, out root);
        }
        else
        {
            root = new Node();
        }
    }

    // Recursively reads nodes from 'serializedNodes' and reconstructs the runtime tree.
    // Returns the next index to read from.
    private int ReadNodeFromSerializedNodes(int index, out Node node)
    {
        var serializedNode = serializedNodes[index];

        // Create a new runtime node and transfer data.
        Node newNode = new Node()
        {
            interestingValue = serializedNode.interestingValue,
            children = new List<Node>()
        };

        // The children were serialized immediately after their parent (depth-first).
        // Increment the index and recursively read children.
        int currentIndex = index + 1; 
        for (int i = 0; i < serializedNode.childCount; i++)
        {
            Node childNode;
            currentIndex = ReadNodeFromSerializedNodes(currentIndex, out childNode);
            newNode.children.Add(childNode);
        }

        node = newNode;
        return currentIndex;
    }

}
```

## Directly serializing a Dictionary

Unity’s default serializer doesn’t support `Dictionary<TKey, TValue>`. By implementing [`ISerializationCallbackReceiver`](../ScriptReference/ISerializationCallbackReceiver.md), you can convert the dictionary into two `List` instances (one for keys, one for values) before serialization, and then reconstruct the dictionary from these lists after deserialization.

```
using UnityEngine;
using System;
using System.Collections.Generic;

public class SerializationCallbackScript : MonoBehaviour, ISerializationCallbackReceiver
{
    public List<int> keys = new List<int> { 3, 4, 5 };
    public List<string> values = new List<string> { "I", "Love", "Unity" };

    // Create a Dictionary. The Unity serializer doesn't support Dictionary types.
    public Dictionary<int, string>  myDictionary = new Dictionary<int, string>();

    public void OnBeforeSerialize()
    {
        keys.Clear();
        values.Clear();
        // For each key/value pair in the dictionary, add the key to the keys list and the value to the values list
        foreach (var kvp in myDictionary)
        {
            keys.Add(kvp.Key);
            values.Add(kvp.Value);
        }
    }

    public void OnAfterDeserialize()
    {
        myDictionary = new Dictionary<int, string>();
        // Loop through the list of keys and values and add each key/value pair to the dictionary
        for (int i = 0; i != Math.Min(keys.Count, values.Count); i++)
            myDictionary.Add(keys[i], values[i]);
    }

    void OnGUI()
    {
        // This callback displays the following output on three separate labels in the GameView:
        // "Key: 3 value: I"
        // "Key: 4 value: Love"
        // "Key: 5 value: Unity"
        foreach (var kvp in myDictionary)
            GUILayout.Label("Key: " + kvp.Key + " value: " + kvp.Value);
    }
}
```

## Polymorphic deserialization

Unity’s default serializer has limited support for polymorphism with plain C# classes. To serialize a list of objects that inherit from a common base class, you often need to manually store type information and the specific data for each derived type. This example shows how to achieve this using `ISerializationCallbackReceiver`.

```
using UnityEngine;
using System.Collections.Generic;
using System;

// Base class for polymorphic objects. Not directly serializable by Unity.
public abstract class Animal
{
    public string Name { get; set; }
    public abstract string GetSound();
}

// Derived class
public class Dog : Animal
{
    public int BarkVolume { get; set; }
    public Dog(string name, int barkVolume)
    {
        Name = name;
        BarkVolume = barkVolume;
    }
    public override string GetSound() => $"Woof! (Volume: {BarkVolume})";
}

// Derived class
public class Cat : Animal
{
    public bool IsCute { get; set; }
    public Cat(string name, bool isCute)
    {
        Name = name;
        IsCute = isCute;
    }
    public override string GetSound() => $"Meow! (Cute: {IsCute})";
}

// Serializable wrapper struct for Dog's data.
// Marked [Serializable] for Unity.
[Serializable]
public struct SerializableDog
{
    public string name;
    public int barkVolume;
}

// Serializable wrapper struct for Cat's data.
// Marked [Serializable] for Unity.
[Serializable]
public struct SerializableCat
{
    public string name;
    public bool isCute;
}

// Enum to identify the type of animal for deserialization.
public enum AnimalType
{
    Dog,
    Cat
}

// Example demonstrating polymorphic deserialization using ISerializationCallbackReceiver.
// This approach requires manual management of type information during serialization.
public class PolymorphicAnimalSerialization : MonoBehaviour, ISerializationCallbackReceiver
{
    // The runtime list of polymorphic Animal objects. Not directly serialized by Unity.
    public List<Animal> animals = new List<Animal>();

    // These lists store the serializable data for each derived type,
    // along with a list of types to reconstruct the original order.
    // Marked [SerializeField] to be saved by Unity.
    [SerializeField] private List<AnimalType> _animalTypes = new List<AnimalType>();
    [SerializeField] private List<SerializableDog> _dogs = new List<SerializableDog>();
    [SerializeField] private List<SerializableCat> _cats = new List<SerializableCat>();

    // Called just before Unity serializes the object.
    public void OnBeforeSerialize()
    {
        _animalTypes.Clear();
        _dogs.Clear();
        _cats.Clear();

        foreach (var animal in animals)
        {
            if (animal is Dog dog)
            {
                _animalTypes.Add(AnimalType.Dog);
                _dogs.Add(new SerializableDog { name = dog.Name, barkVolume = dog.BarkVolume });
            }
            else if (animal is Cat cat)
            {
                _animalTypes.Add(AnimalType.Cat);
                _cats.Add(new SerializableCat { name = cat.Name, isCute = cat.IsCute });
            }
        }
    }

    // Called just after Unity deserializes the object.
    public void OnAfterDeserialize()
    {
        animals.Clear();
        
        // Keep track of which serialized item to read next from each list.
        int dogIndex = 0;
        int catIndex = 0;

        // Iterate through the type list to reconstruct the animals in the correct order.
        foreach (var type in _animalTypes)
        {
            switch (type)
            {
                case AnimalType.Dog:
                    if (dogIndex < _dogs.Count)
                    {
                        var sDog = _dogs[dogIndex];
                        animals.Add(new Dog(sDog.name, sDog.barkVolume));
                        dogIndex++;
                    }
                    else
                    {
                        Debug.LogError("Mismatch in serialized Dog data during deserialization.");
                    }
                    break;
                case AnimalType.Cat:
                    if (catIndex < _cats.Count)
                    {
                        var sCat = _cats[catIndex];
                        animals.Add(new Cat(sCat.name, sCat.isCute));
                        catIndex++;
                    }
                    else
                    {
                        Debug.LogError("Mismatch in serialized Cat data during deserialization.");
                    }
                    break;
            }
        }
    }
}
```

## Additional resources

* [Serialization rules](script-serialization-rules.md)
* [How Unity uses serialization](script-serialization-how-unity-uses.md)
* [JSONSerialization](json-serialization.md)
* [Serialization best practices](script-serialization-best-practices.md)