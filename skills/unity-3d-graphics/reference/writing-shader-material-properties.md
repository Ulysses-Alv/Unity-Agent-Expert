# Introduction to material properties

In your **ShaderLab** code, you can define **material properties**. A material property is a property that Unity stores as part of the material asset. This allows artists to create, edit, and share materials with different configurations.

If you use material properties:

* You can get or set the value of a variable in a **Shader** object by calling functions on the material (such as [Material.SetFloat](../ScriptReference/Material.SetFloat.md)).
* You can view and edit the values using the material **Inspector**.
* Unity saves the changes that you make as part of the material asset, so they persist between sessions.

If you do not use material properties:

* You can still get or set the value of a variable in a **Shader object** by calling a function on a material.
* There is no visual editor for these values.
* Changes do not persist between sessions.

The only times that you would normally *not* create a material property is if you want to set shader property values entirely using **scripts** (for example, if you are making procedural content), if the properties cannot be made into material properties, or if you don’t want them to be edited in the Inspector.

## How property values are provided to shaders

Shader property values are found and provided to shaders from these places:

* Per-Renderer values set in [MaterialPropertyBlock](../ScriptReference/MaterialPropertyBlock.md). This is typically “per-instance” data (e.g. customized tint color for a lot of objects that all share the same material).
* Values set in the [Material](class-Material.md) that’s used on the rendered object.
* Global shader properties, set either by Unity rendering code itself (see [built-in shader variables](SL-UnityShaderVariables.md)), or from your own scripts (e.g. [Shader.SetGlobalTexture](../ScriptReference/Shader.SetGlobalTexture.md)).

The order of precedence is like above: per-instance data overrides everything; then Material data is used; and finally if shader property does not exist in these two places then global property value is used. Finally, if there’s no shader property value defined anywhere, then “default” (zero for floats, black for colors, empty white texture for textures) value will be provided.

## Serialized and Runtime Material properties

[Materials](class-Material.md) can contain both serialized and runtime-set property values.

Serialized data is all the properties defined in shader’s [Properties](SL-Properties.md) block. Typically these are values that need to be stored in the material, and are tweakable by the user in Material Inspector.

A material can also have some properties that are used by the shader, but not declared in shader’s [Properties](SL-Properties.md) block. Typically this is for properties that are set from script code at runtime, e.g. via [Material.SetColor](../ScriptReference/Material.SetColor.md). Note that matrices and arrays can only exist as non-serialized runtime properties (since there’s no way to define them in Properties block).