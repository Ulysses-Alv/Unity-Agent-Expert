# Introduction to mipmap streaming

Use mipmap streaming to limit the size of textures in GPU memory.

## How mipmap streaming works

For each texture in the **camera** view, Unity automatically calculates and loads mipmap levels only up to a specific level, instead of loading all of them. This means that Unity only transfers some mipmap levels per texture from disk to the CPU and the GPU.

Unity loads mipmap levels at the highest resolution possible, but uses lower mipmap levels if higher resolution mipmap levels don’t fit in the memory limit you set. For more information about setting a memory limit, refer to [Configure mipmap streaming](TextureStreaming-configure.md).

Unity caches mipmap levels on the GPU, to avoid repeatedly loading mipmap levels from disk and the CPU.

## Limitations

* Unity doesn’t support mipmap streaming on [terrain](script-Terrain.md) textures, because Unity needs the highest resolution mipmap levels at all times.
* Unity doesn’t support mipmap streaming with texture arrays, **cubemap** arrays, or 3D textures.
* If you use an API such as [Graphics.DrawMeshNow](../ScriptReference/Graphics.DrawMeshNow.md) to render a texture, Unity doesn’t have the information it needs to calculate a mipmap level. Set the mipmap level for the texture manually with the [Texture2D.requestedMipmapLevel](../ScriptReference/Texture2D-requestedMipmapLevel.md) API, or disable mipmap streaming on the texture.
* Unity might not be able to calculate the correct mipmap level for textures that don’t use the special `_ST` property for texture tiling and offset. For more information about the `_ST` property, refer to [Accessing shader properties in Cg/HLSL](SL-PropertiesInPrograms.md).