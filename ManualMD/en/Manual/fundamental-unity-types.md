# Fundamental Unity types

Unity has some fundamental built-in classes that are particularly important for scripting. These are classes which your own custom types can inherit from to integrate with Editor and Engine functionality. It’s helpful to understand these types, their behavior, and why you should inherit from or use them.

For a complete reference of all the built-in classes and every member available, refer to the [Script Reference](../ScriptReference/index.md).

| **Topic** | **Description** |
| --- | --- |
| **[Object](class-Object.md)** | `UnityEngine.Object` is the base class for all objects the Editor can reference from fields in the **Inspector** window. |
| **[MonoBehaviour](class-MonoBehaviour.md)** | Inherit from `MonoBehaviour` to make your script a component and control the behaviour of GameObjects and make them responsive to events. |
| **[ScriptableObject](class-ScriptableObject.md)** | Inherit from `ScriptableObject` to store data that’s independent of GameObjects. |
| **[Unity attributes](unity-attributes.md)** | Use Unity-specific C# attributes to define special behavior for your code. |

## Additional resources

* [Unity Scripting reference](../ScriptReference/index.md)