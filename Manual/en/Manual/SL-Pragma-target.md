# HLSL pragma target command reference

Here is the list of **shader** models that Unity uses, and the combination of `#pragma require` values that each corresponds to.

**Note:** Unity’s shader models are similar to DirectX shader models and OpenGL version requirements, but they do not correspond exactly. Read the descriptions carefully to ensure that you understand the differences.

| **Value** | **Description** | **Support** | **Equivalent `#pragma require` values** |
| --- | --- | --- | --- |
| `2.0` | Equivalent to [DirectX shader model 2.0](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-sm2).  Limited amount of arithmetic and texture instructions; 8 interpolators; no vertex texture sampling; no derivatives in fragment shaders; no explicit LOD texture sampling. | Works on all platforms supported by Unity. | N/A |
| `2.5` | Almost the same as 3.0, but with only 8 interpolators, and no explicit LOD texture sampling. | DirectX 11 feature level 9+ OpenGL 3.2+ Vulkan Metal | `derivatives` |
| `3.0` | Equivalent to [DirectX shader model 3.0](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-sm3).  . | DirectX 11 feature level 10 + OpenGL 3.2+ OpenGL ES 3.0+ Vulkan Metal | Everything in `2.5`, plus: `interpolators10 samplelod fragcoord` |
| `3.5` | Equivalent to [OpenGL ES 3.0](https://en.wikipedia.org/wiki/OpenGL_ES#OpenGL_ES_3.0). | DirectX 11 feature level 10+ OpenGL 3.2+ OpenGL ES 3+ Vulkan Metal | Everything in `3.0`, plus: `interpolators15 mrt4 integers 2darray instancing` |
| `4.0` | Equivalent to [DirectX shader model 4.0](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-sm4), but without the requirement to support 8 MRTs. | DirectX 11 feature level 10+ Desktop OpenGL 3.x OpenGL ES 3.1 OpenGL ES 3.1 + Android Extension Pack (AEP) OpenGL 3.2+ Vulkan Metal (if no geometry stage is defined) | Everything in `3.5`, plus: `geometry` |
| `gl4.1` | Equivalent to [OpenGL 4.1](https://en.wikipedia.org/wiki/OpenGL#OpenGL_4.1) | Desktop OpenGL 4.1 Shader model 4.0 + tessellation to match macOS 10.9 capabilities. | Everything in `4.0`, plus: `cubearray tesshw tessellation msaatex` |
| `4.5` | Equivalent to [OpenGL ES 3.1](https://en.wikipedia.org/wiki/OpenGL_ES#OpenGL_ES_3.1). | DirectX 11 feature level 11+ OpenGL 4.3+ OpenGL ES 3.1 Vulkan Metal | Everything in `3.5`, plus: `compute randomwrite msaatex` |
| `4.6` | Equivalent to [OpenGL 4.1](https://en.wikipedia.org/wiki/OpenGL#OpenGL_4.1).  This is the highest OpenGL level supported on a Mac. | DirectX 11 feature level 11+ OpenGL 4.1+ OpenGL ES 3.1 + AEP Vulkan Metal (if no geometry stage is defined, and no hull or domain stage is defined) | Everything in `4.0`, plus: `cubearray tesshw tessellation msaatex` |
| `5.0` | Equivalent to [DirectX shader model 5.0](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/d3d11-graphics-reference-sm5), but without the requirement to support 32 interpolators or cubemap arrays. | DirectX 11 feature level 11+ Desktop OpenGL >= 4.2 OpenGL 4.3+ OpenGL ES 3.1 + AEP Vulkan Metal (if no geometry stage is defined, and no hull or domain stage is defined) | Everything in `4.0`, plus: `compute randomwrite msaatex tesshw tessellation` |

For information on shader model support for console platforms, see the platform-specific documentation.

**Notes:**

* In the DirectX definitions, shader model 4.0 includes `mrt8`, and shader model 5.0 includes `interpolators32` and `cubearray`. Unity does not include these, for broader compatibility. To require these features, use an explicit `#pragma require` directive.
* If you use a target that requires `geometry` but your shader does not define a geometry stage, Unity removes `geometry` from the list of requirements at compile time.
* If you use a target that requires `tessellation` but your shader does not define a hull or domain stage, Unity removes `tessellation` from the list of requirements at compile time.