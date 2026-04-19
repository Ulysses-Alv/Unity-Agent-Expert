# Mipmaps

A mipmap level is a version of a texture with a specific resolution. Mipmap levels exist in sets called mipmaps. Mipmaps contain progressively smaller and lower resolution versions of a single texture.

For example, a mipmap might contain four versions of a texture, from the original texture (mipmap level 0), to mipmap level 1, mipmap level 2, and mipmap level 3:

![Mipmap levels](../uploads/Main/mipmaps.png)

Mipmap levels

Mipmaps are commonly used for rendering objects in 3D **scenes**, where textured objects can vary in distance from the **camera**. A higher mipmap level is used for objects closer to the camera, and lower mipmap levels are used for more distant objects.

Mipmaps can speed up rendering operations and reduce rendering artifacts in situations where the GPU renders a texture at less than its full resolution. A mipmap is effectively a cached, downsampled version of the original texture. Instead of performing many sampling operations on the original, full resolution texture, the GPU can perform a smaller number of operations on the already downsampled version.

Sometimes, mipmaps aren’t beneficial. Mipmaps increase the size of a texture by 33%, both on disk and in memory. They also provide no benefit when a texture is only rendered at its full resolution, such as a UI texture that isn’t scaled.

You can create a mipmap for a texture manually, or you can instruct Unity to generate a mipmap for you. To automatically generate a mipmap, ensure that your original texture’s resolution is a power of two value, as shown in the example mipmap image.

You can enable or disable mipmaps for a texture asset in the [Texture Import Settings Inspector](class-TextureImporter.md).

## How the GPU samples mipmap levels

When the GPU samples a texture, it determines which mipmap level to use based on the texture coordinates (UVs) for the current **pixel**, and two internal values that the GPU calculates: DDX and DDY. DDX and DDY provide information about the UVs of the pixels beside and above the current pixel, including distances and angles.

The GPU uses these values to determine how much of a texture’s detail is visible to the camera. A greater distance and a more extreme angle between the current pixel and its neighbors means that the GPU picks a lower resolution mipmap level; a shorter distance and less extreme angle means that the GPU picks a mipmap level with a higher resolution.

The GPU can also blend the texture information from two mipmap levels together with trilinear filtering. Blending mipmap levels while sampling can make the transition from one mipmap level to another less noticeable. To blend mipmap levels, the GPU takes a specific percentage of texture information from one mipmap level and the rest from another mipmap level.

## Mipmap bias

A setting called mipmap bias can do two things while sampling, based on the sampler settings:

* The mipmap bias can change the threshold for the GPU selecting a lower or higher mipmap level for a sample. The GPU selects a specific mipmap level when you use point and linear filtering in a sampler. For example, the GPU’s might decide that the texture at a set of UVs uses a sample from mipmap level 3. With a mipmap bias of –2, the GPU would use the higher resolution mipmap level 1 for the sample, instead.
* The mipmap bias can tell the GPU to prefer one mipmap level over another by an exact percentage when blending samples from different mipmap levels. The GPU blends mipmap levels when you use trilinear filtering in a sampler. For example, the GPU’s calculations might return a value of 0.5. The 0.5 value tells the GPU to take 50% of the texture information it needs from one mipmap level, and the remaining 50% from the next mipmap level. With an added mipmap bias of 0.2, the 0.5 value would change to 0.7, and the GPU would take 70% of the texture information from the first mipmap level and only 30% from the second.

The GPU has a global mipmap bias that it applies to its mipmap level selection by default. Textures can have their own mipmap bias, which Unity adds or subtracts from the global mipmap bias. You can also specify your own mipmap bias for an individual texture sampling operation in a **shader**.

To set the mipmap bias for an individual texture, see [Texture.mipMapBias](../ScriptReference/Texture-mipMapBias.md). To set a mipmap bias for a texture sampling operation in a hand-coded shader, use HLSL functions such as [tex2dbias](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-tex2dbias). To set a mipmap bias for a texture sampling operation in Shader Graph, see [Sample texture 2D Array node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Sample-Texture-2D-Array-Node.html) or [Sample texture 2D node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Sample-Texture-2D-Node.html).

## How Unity loads mipmaps

You can control the way that Unity loads mipmaps at runtime with [Mipmap Streaming](TextureStreaming.md).

## Additional resources

* [Optimizing GPU texture memory with mipmap streaming](TextureStreaming.md)