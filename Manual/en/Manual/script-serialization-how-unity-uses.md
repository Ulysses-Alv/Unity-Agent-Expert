# How Unity uses serialization

## Saving and loading

Unity uses serialization to load and save [scenes](CreatingScenes.md), [Assets](AssetWorkflow.md), and [AssetBundles](AssetBundlesIntro.md) to and from your device’s memory. This includes data saved in your own scripting API objects such as [MonoBehaviour](../ScriptReference/MonoBehaviour.md) components and [ScriptableObjects](class-ScriptableObject.md).

Many of the features in the Unity Editor are built on top of the core serialization system. Two things to be particularly aware of with serialization are the [Inspector window](UsingTheInspector.md), and hot reloading.

### The Inspector window

The **Inspector** window shows the value of the serialized fields of the inspected objects. When you change a value in the Inspector, the Inspector updates the serialized data and triggers a deserialization that updates the inspected object.

The same applies for both built-in Unity objects, and scripting objects such as MonoBehaviour-derived classes.

Unity doesn’t call any C# property getters and setters when you view or change values in the Inspector window. Instead, Unity accesses the serialized backing field directly.

### Hot reloading

Hot reloading of script code is performed as part of an asset database refresh. It refers to the process of reloading and applying code changes directly while the Editor is running, without having to restart it. For more information, refer to [Refreshing the Asset Database](AssetDatabaseRefreshing.md) and [Hot reloading](AssetDatabaseRefreshing.md#hotreloading).

**Note**: Hot reloading is a special serialization case. Unlike in other serialization cases, Unity serializes private fields by default when reloading, even if they don’t have the [`[SerializeField]`](../ScriptReference/SerializeField.md) attribute.

When Unity reloads **scripts**:

1. Unity serializes and stores all variables in all loaded scripts.
2. Unity restores them to their original, pre-serialization values:
   * Unity restores all variables - **including private variables** - that fulfill the requirements for serialization, even if a variable has no `[SerializeField]` attribute.
   * Unity never restores static variables, so don’t use static variables for states that you need to keep after Unity reloads a script because the reloading process will discard them.

### Preventing serialization and restoration

Sometimes, you need to prevent Unity from restoring private variables, for example, if you want a reference to be null after reloading from scripts. In this case, use the [`[NonSerialized]`](https://learn.microsoft.com/en-us/dotnet/api/system.nonserializedattribute?view=net-9.0) attribute.

Unity’s restore behavior also applies to the private backing fields that store the value of [automatically implemented properties](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/auto-implemented-properties) in C#. For example, consider the following automatically implemented property:

```
public string MyProperty { get; set; }
```

In this case, C# adds an invisible private field to the class to store the actual value of `MyProperty`. Unity would serialize and restore this value during a hot reload. To prevent this, use `[field: NonSerialized]` on the property.

The following example demonstrates using both the `[NonSerialized]` attribute and `[field: NonSerialized]` to prevent serialization and restoration of regular fields and an auto-property backing field respectively:

```
class Test
{
    // p will not be shown in the Inspector or serialized
    [System.NonSerialized]
    public int p = 5;

    // neverSerializeMe will never be serialized, even during an hot reload.
    [System.NonSerialized]
    private int neverSerializeMe;

    // The backing field for NeverSerializedProperty property will never be serialized,
    // even during a hot reload
    [field: System.NonSerialized]
    public int NeverSerializedProperty { get; set; }
}
```

## Prefabs

A [prefab](Prefabs.md) is the serialized data of one or more [GameObjects](GameObjects.md) or [components](Components.md). A prefab instance contains a reference to both the prefab source and a list of modifications to it. The modifications are what Unity needs to do to the prefab source to create that particular prefab instance.

The prefab instance only exists while you edit your project in the Unity Editor. The Unity Editor instantiates a GameObject from its two sets of serialization data: the prefab source and the prefab instance’s modifications.

## Instantiation

When you call [`Instantiate`](../ScriptReference/Object.Instantiate.md) on anything that exists in a scene, such as a prefab or a GameObject:

1. Unity serializes it. This happens both at runtime and in the Editor. Unity can serialize everything that derives from `UnityEngine.Object`.
2. Unity creates a new GameObject and deserializes the data onto the new GameObject.
3. Unity runs the same serialization code in a different variant to report which other `UnityEngine.Objects` it references. It checks all referenced `UnityEngine.Objects` to determine if they’re part of the data Unity instantiates. If the reference points to something external, such as a Texture, Unity keeps that reference as it is. If the reference points to something internal, such as a child GameObject, Unity patches the reference to the corresponding copy.

## Unloading unused assets

`EditorUtility.UnloadUnusedAssetsImmediate` is the native Unity garbage collector and has a different purpose to the standard C# garbage collector. It runs after you load a scene and checks for objects (like Textures) that it no longer references and unloads them safely. The native Unity garbage collector runs the serializer in a variation in which objects report all references to external `UnityEngine.Objects`. This is how Textures that one scene uses, the garbage collector unloads in the next.

## Differences between Editor and runtime serialization

Most serialization happens in the Editor, whereas deserialization is the focus at runtime. Unity serializes some features only in the Editor, while it can serialize other features in both the Editor and at runtime:

| **Feature** | **Editor** | **Runtime** |
| --- | --- | --- |
| **Assets in Binary Format** | Read/write supported | Read supported |
| **Assets in YAML format** | Read/write supported | Not supported |
| **Saving scenes, prefabs and other assets** | Supported, unless in Play mode | Not supported |
| **Serialization of individual objects with [JsonUtility](json-serialization.md)** | Read/write support with JsonUtility.  Support for additional types of objects with EditorJsonUtility | Read/write support with JsonUtility |
| **[SerializeReference](../ScriptReference/SerializeReference.md)** | Supported | Supported |
| **[ISerializationCallbackReceiver](../ScriptReference/ISerializationCallbackReceiver.md)** | Supported | Supported |
| **[FormerlySerializedAs](../ScriptReference/Serialization.FormerlySerializedAsAttribute.md)** | Supported | Not supported |

Objects can have additional fields that only the Editor serializes, such as when you declare fields within the UNITY\_EDITOR [scripting symbol](platform-dependent-compilation.md):

```
public class SerializeRules : MonoBehaviour
{
#if UNITY_EDITOR
public int m_intEditorOnly;
#endif
}
```

In the previous example, the `m_intEditorOnly` field is only serialized in the editor and isn’t included in the build. This allows you to save memory by omitting data that’s only required in the Editor from your build. Any code that uses that field would also need to be conditionally compiled, for example within `#if UNITY_EDITOR` blocks, so that the class can compile at build time.

The Editor doesn’t support objects with fields that Unity only serializes at runtime, (for example, when you declare fields within the UNITY\_STANDALONE directive).

## Additional resources

* [Serialization rules](script-serialization-rules.md)
* [JSONSerialization](json-serialization.md)
* [Serialization best practices](script-serialization-best-practices.md)