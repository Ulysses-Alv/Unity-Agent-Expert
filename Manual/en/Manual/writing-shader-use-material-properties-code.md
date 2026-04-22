# Access material properties in a script

Material properties are represented in C# code by the [MaterialProperty](../ScriptReference/MaterialProperty.md) class.

To access variables defined in your HLSL code, you can call [Material.GetFloat](../ScriptReference/Material.GetFloat.md), [Material.SetFloat](../ScriptReference/Material.SetFloat.md). There are other, similar methods; see the [Material API documentation](../ScriptReference/Material.md) for a full list. When you access HLSL variables using these APIs, it doesn’t matter whether the variable is a material property or not.