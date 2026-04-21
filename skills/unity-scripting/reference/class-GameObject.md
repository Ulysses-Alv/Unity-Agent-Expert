# The GameObject class

[Switch to Scripting](../ScriptReference/GameObject.md "Go to GameObject page in the Scripting Reference")

Unity’s **GameObject** class represents anything that can exist in a ****Scene****. GameObjects are the building blocks for scenes in Unity and act as a container for functional components which determine how the GameObject looks and what it does.

The [GameObject class](../ScriptReference/GameObject.md) class provides a collection of methods which allow you to work with GameObjects in your code. These include methods for finding, making connections, and sending messages between GameObjects, adding or removing components attached to the GameObject, and setting values relating to their status within the scene.

For a complete reference of every member of the GameObject class, refer to the [GameObject script reference](../ScriptReference/GameObject.md).

For information on using GameObjects in the Scene and Hierarchy in the Unity Editor, refer to [Introduction to GameObjects](GameObjects.md).

## Scene status properties

All GameObjects share a set of controls at the top of the **Inspector** relating to the GameObject’s status within the scene, and these can be controlled via the GameObject’s scripting API.

![A typical GameObject viewed in the Inspector. In this case, a directional light. The Scene status properties are outlined in red.](../uploads/Main/GameObjectSceneStatusProperties.png)

A typical GameObject viewed in the Inspector. In this case, a directional light. The Scene status properties are outlined in red.

## Active status

![The Active status of a GameObject](../uploads/Main/GOInspectorActiveSetting.png)

The Active status of a GameObject

GameObjects are active by default, but can be deactivated, which turns off all components attached to the GameObject. This generally means it will become invisible, and not receive any of the normal callbacks or events such as [`Update`](../ScriptReference/MonoBehaviour.Update.md) or [`FixedUpdate`](../ScriptReference/MonoBehaviour.FixedUpdate.md).

The GameObject’s **active** status is represented by the checkbox to the left of the GameObject’s name. You can control this using [`GameObject.SetActive`](../ScriptReference/GameObject.SetActive.md).

You can also use [`GameObject.activeSelf`](../ScriptReference/GameObject-activeSelf.md) to read the current active state of a GameObject. Use [`GameObject.activeInHierarchy`](../ScriptReference/GameObject-activeInHierarchy.md) to read whether the GameObject is actually active in the scene. [`GameObject.activeInHierarchy`](../ScriptReference/GameObject-activeInHierarchy.md) is necessary because whether a GameObject is actually active is determined by its own active state and the active state of all its parents. If any of its parents aren’t active, then it’s not active despite its own active setting.

## Static status

![The Static status of a GameObject](../uploads/Main/GOInspectorStaticSetting.png)

The Static status of a GameObject

Some of Unity’s systems, such as **Global Illumination**, Occlusion, Batching, Navigation, and **Reflection Probes**, rely on the static status of a GameObject. You can control which of Unity’s systems consider the GameObject to be **static** by using [`GameObjectUtility.SetStaticEditorFlags`](../ScriptReference/GameObjectUtility.SetStaticEditorFlags.md). For more information, refer to [Static GameObjects](https://docs.unity3d.com/Manual/StaticObjects.html).

## Tags and Layers

![The Tag and Layer fields of a GameObject](../uploads/Main/GOInspectorTagsAndLayers.png)

The Tag and Layer fields of a GameObject

**Tags** provide a way of marking and identifying types of GameObject in your scene and **Layers** provide a similar but distinct way of including or excluding groups of GameObjects from certain built-in actions, such as [rendering](class-MeshRenderer.md) or [physics collisions](CollidersOverview.md).

For more information about how to use Tags and Layers in the editor, refer to the main user manual pages for [Tags](Tags.md) and [Layers](Layers.md).

You can modify tag and layer values via script using the [`GameObject.tag`](../ScriptReference/GameObject-tag.md) and [`GameObject.layer`](../ScriptReference/GameObject-layer.md) properties. You can also check a GameObject’s tag efficiently by using the [`CompareTag`](../ScriptReference/GameObject.CompareTag.md) method, which includes validation of whether the tag exists, and doesn’t cause any memory allocation.

## Add and remove components

You can add or remove components at runtime, which can be useful for procedurally creating GameObjects, or modifying how a GameObject behaves. Note, you can also [`enable` or `disable`](../ScriptReference/Behaviour-enabled.md) script components, and some types of built-in component, via script without destroying them.

The best way to add a component at runtime is to use [`AddComponent<Type>`](../ScriptReference/GameObject.AddComponent.md), specifying the type of component within angle brackets as shown. To remove a component, you must use [`Object.Destroy`](../ScriptReference/Object.Destroy.md) method on the component itself.

## Access components

The simplest case is where a script on a GameObject needs to access another Component attached to the same GameObject (remember, other scripts attached to a GameObject are also Components themselves). To do this, the first step is to get a reference to the Component instance you want to work with. This is done with the [GetComponent](../ScriptReference/Component.GetComponent.md) method. Typically, you want to assign the Component object to a variable, which is done in using the following code. In this example the script is getting a reference to a **Rigidbody** component on the same GameObject:

```
void Start ()
{
    Rigidbody rb = GetComponent<Rigidbody>();
}
```

Once you have a reference to a Component instance, you can set the values of its properties much as you would in the Inspector:

```
void Start ()
{
    Rigidbody rb = GetComponent<Rigidbody>();

    // Change the mass of the object's Rigidbody.
    rb.mass = 10f;
}
```

You can also call methods on the Component reference, for example:

```
void Start ()
{
    Rigidbody rb = GetComponent<Rigidbody>();

    // Add a force to the Rigidbody.
    rb.AddForce(Vector3.up * 10f);
}
```

Note: you can have multiple custom **scripts** attached to the same GameObject. If you need to access one script from another, you can use GetComponent as usual and just use the name of the script class (or the file name) to specify the Component type you want.

If you attempt to retrieve a Component type that hasn’t actually been added to the GameObject then GetComponent will return null. You will get a null reference error at runtime if you try to change any values on a null object.

## Access components on other GameObjects

Although they sometimes operate in isolation, it’s common for scripts to keep track of other GameObjects, or more commonly, components on other GameObjects. For example, in a cooking game, a chef might need to know the position of the stove. Unity provides different ways to retrieve other objects, each appropriate to certain situations.

### Link to GameObjects with variables in the Inspector

The most straightforward way to find a related GameObject is to add a public GameObject variable to the script:

```
public class Chef : MonoBehaviour
{
    public GameObject stove;

    // Other variables and functions...
}
```

This variable will be visible in the Inspector, as a **GameObject field**.

You can now drag an object from the scene or Hierarchy panel onto this variable to assign it.

The GetComponent function and Component access variables are available for this object as with any other, so you can use code like the following:

```
public class Chef : MonoBehaviour {

    public GameObject stove;

    void Start() {
        // Start the chef 2 units in front of the stove.
        transform.position = stove.transform.position + Vector3.forward * 2f;
    }
}
```

Additionally, if you declare a public variable of a Component type in your script, you can drag any GameObject that has that Component attached onto it. This accesses the Component directly rather than the GameObject itself.

```
public Transform playerTransform;
```

Linking objects together with variables is most useful when you are dealing with individual objects that have permanent connections. You can use an array variable to link several objects of the same type, but the connections must still be made in the Unity editor rather than at runtime. It’s often convenient to locate objects at runtime and Unity provides two basic ways to do this, as described below.

### Find child GameObjects

Sometimes, a game Scene makes use of multiple GameObjects of the same type, such as collectibles, waypoints, and obstacles. These might need to be tracked by a particular script that supervises or reacts to them (for example, all waypoints might need to be available to a pathfinding script). Using variables to link these GameObjects is a possibility but it makes the design process tedious if each new waypoint has to be dragged to a variable on a script. Likewise, if a waypoint is deleted, then it’s a nuisance to have to remove the variable reference to the missing GameObject. In cases like this, it is often better to manage a set of GameObjects by making them all children of one parent GameObject. The child GameObjects can be retrieved using the parent’s **Transform component** (because all GameObjects implicitly have a Transform):

```
using UnityEngine;

public class WaypointManager : MonoBehaviour {
    public Transform[] waypoints;

    void Start()
    {
        waypoints = new Transform[transform.childCount];
        int i = 0;

        foreach (Transform t in transform)
        {
            waypoints[i++] = t;
        }
    }
}
```

You can also locate a specific child object by name using the [Transform.Find](../ScriptReference/Transform.Find.md) method:

```
transform.Find("Frying Pan");
```

This can be useful when a GameObject has a child GameObject that can be added and removed during gameplay. A tool or utensil that can be picked up and put down during gameplay is a good example of this.

## Send and broadcast messages

While editing your project you can set up references between GameObjects in the Inspector. However, sometimes it’s impossible to set these up in advance (for example, finding the nearest item to a character in your game, or making references to GameObjects that were instantiated after the Scene loaded). In these cases, you can find references and send messages between GameObjects at runtime.

[`BroadcastMessage`](../ScriptReference/GameObject.BroadcastMessage.md) allows you to send out a call to a named method, without being specific about where that method should be implemented. You can use it to call a named method on every MonoBehaviour on a particular GameObject or any of its children. You can optionally choose to enforce that there must be at least one receiver (or an error is generated).

[`SendMessage`](../ScriptReference/GameObject.SendMessage.md) is a little more specific, and only sends the call to a named method on the GameObject itself, and not its children.

[`SendMessageUpwards`](../ScriptReference/GameObject.SendMessageUpwards.md) is similar, but sends out the call to a named method on the GameObject and all its *parents*.

## Find GameObjects by Name or Tag

It’s always possible to locate GameObjects anywhere in the Scene hierarchy as long as you have some information to identify them. Individual objects can be retrieved by name using the [GameObject.Find](../ScriptReference/GameObject.Find.md) function:

```
GameObject player;

void Start()
{
    player = GameObject.Find("MainHeroCharacter");
}
```

An object or a collection of objects can also be located by their tag using the [GameObject.FindWithTag](../ScriptReference/GameObject.FindWithTag.md) and [GameObject.FindGameObjectsWithTag](../ScriptReference/GameObject.FindGameObjectsWithTag.md) methods.

For example, in a cooking game with one chef character, and multiple stoves in the kitchen (each tagged “Stove”):

```
GameObject chef;
GameObject[] stoves;

void Start()
{
    chef = GameObject.FindWithTag("Chef");
    stoves = GameObject.FindGameObjectsWithTag("Stove");
}
```

## Create and destroy GameObjects

You can create and destroy GameObjects while your project is running. In Unity, a GameObject can be created using the [Instantiate](../ScriptReference/Object.Instantiate.md) method which makes a new copy of an existing object.

For a full description and examples of how to instantiate GameObjects, see [Instantiating Prefabs at Runtime](instantiating-prefabs.md).

The [Destroy](../ScriptReference/Object.Destroy.md) method destroys an object after the frame update has finished or optionally after a short time delay:

```
void OnCollisionEnter(Collision otherObj) {
    if (otherObj.gameObject.tag == "Garbage can") {
        Destroy(gameObject, 0.5f);
    }
}
```

Note that the Destroy function can destroy individual components and not affect the GameObject itself. A common mistake is to write `this` and assume it destroys the GameObject the script is attached to:

```
Destroy(this);
```

`this` represents the script, and not the GameObject. It will actually just destroy the script component that calls it and leave the GameObject intact but with the script component removed.

## Primitives

The GameObject class offers script-based alternatives to the options available in Unity’s GameObject menu that allows you to create [primitive objects](PrimitiveObjects.md).

To create instances of Unity’s built-in primitives, use [GameObject.CreatePrimitive](../ScriptReference/GameObject.CreatePrimitive.md), which instantiates a primitive of the type that you specify. The available primitive types are [Sphere](../ScriptReference/PrimitiveType.Sphere.md), [Capsule](../ScriptReference/PrimitiveType.Capsule.md), [Cylinder](../ScriptReference/PrimitiveType.Cylinder.md), [Cube](../ScriptReference/PrimitiveType.Cube.md), [Plane](../ScriptReference/PrimitiveType.Plane.md) and [Quad](../ScriptReference/PrimitiveType.Quad.md).

![The Primitive shapes available in Unitys GameObject menu](../uploads/Main/AllPrimitives.png)

The Primitive shapes available in Unity’s GameObject menu

## Additional resources

* [Introduction to GameObjects](GameObjects.md)
* [GameObject API reference](../ScriptReference/GameObject.md)