# Enable GPU instancing for a prebuilt material

To enable GPU instancing for a [prebuilt materials](built-in-materials-and-shaders.md), follow these steps:

1. Select the material in the **Project** window.
2. In the **Advanced Options** section of the ****Inspector**** window, enable **Enable GPU Instancing**.

**Note:** If you use the Universal **Render Pipeline** (URP) or High Definition Render Pipeline (HDRP), the recommended best practice is to use the [SRP Batcher](SRPBatcher.md) instead, which is enabled by default. GPU instancing works only if you [disable the SRP Batcher](SRPBatcher-Incompatible.md).

If multiple **GameObjects** use the same **mesh** and material, Unity now uses GPU instancing to render them in single draw calls. To check this, open the **Frame Debugger** and look for render passes called **Draw Mesh (Instanced)**.

If there’s no **Enable GPU Instancing** property, the prebuilt **shader** doesn’t support GPU instancing.

## Change the properties of instances

To change the properties of instances at runtime, for example to give each instance a different color or position, create a script to do the following:

1. Create a [MaterialPropertyBlock](../ScriptReference/MaterialPropertyBlock.md) with the property value for the instance.
2. Attach the script to the [Mesh Renderer](class-MeshRenderer.md) component of the GameObject.

For more information, refer to [MaterialPropertyBlock](../ScriptReference/MaterialPropertyBlock.md).

## Create custom instances

To render multiple instances of a mesh in a script using GPU instancing, use one of the following APIs:

* [Graphics.RenderMeshInstanced](../ScriptReference/Graphics.RenderMeshInstanced.md)
* [Graphics.RenderMeshIndirect](../ScriptReference/Graphics.RenderMeshIndirect.md)
* [Graphics.RenderMeshPrimitives](../ScriptReference/Graphics.RenderMeshPrimitives.md)

## Additional resources

* [Creating custom shaders that support GPU instancing in the Built-In Render Pipeline](gpu-instancing-shader.md)
* [Indirect & Procedural Rendering in Shader Graph](https://discussions.unity.com/t/indirect-procedural-rendering/1664601)