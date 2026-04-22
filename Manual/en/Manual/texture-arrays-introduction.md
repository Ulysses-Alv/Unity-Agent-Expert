# Introduction to 2D texture arrays

A texture array is a collection of 2D textures with the same size, format, and flags. The individual texture in the array are called slices or layers. In a **shader**, you can use an index to sample each slice.

Texture arrays are useful for implementing custom **terrain** rendering systems or other special effects where you need an efficient way to access many textures of the same size and format.

## Platform support

Texture arrays are supported if the platform supports one of the following graphics APIs:

* DirectX11 and DirectX12 (Windows)
* Metal (iOS, macOS)
* OpenGL Core (macOS, Linux)
* OpenGL ES 3.0 (Android, WebGL 2.0)
* Vulkan (Windows, Linux)

To check if a platform supports texture arrays, use the [SystemInfo.supports2DArrayTextures](../ScriptReference/SystemInfo-supports2DArrayTextures.md) API at runtime.