# Frame data textures reference for URP

You can fetch the following textures from the frame data.

## Color data

| **Property** | **Texture** | **URP **shader** pass that writes to the texture** |
| --- | --- | --- |
| `activeColorTexture` | The color texture the **camera** currently targets. | Any pass, depending on your settings |
| `afterPostProcessColor` | The main color texture after URP’s post processing passes. | `UberPost` |
| `backBufferColor` | The color texture of the screen back buffer. If you use [post-processing](integration-with-post-processing.md), URP writes to this texture at the end of rendering, unless you enable [HDR Debug Views](post-processing/hdr-output-debug-views-urp.md). Refer to `debugScreenTexture` for more information. | Any pass, depending on your settings |
| `cameraColor` | The main color texture for the camera. You can store multiple samples in this texture if you enable [Multisample Anti-aliasing (MSAA)](anti-aliasing.md#msaa). To use this texture, set `ScriptableRenderPass.requiresIntermediateTexture` to `true`. | Any pass, depending on your settings |
| `cameraOpaqueTexture` | A texture with the opaque objects in the **scene**, if you enable **Opaque Texture** in the [URP Asset](universalrp-asset.md). | `CopyColor` |
| `debugScreenTexture` | If you enable [HDR Debug Views](post-processing/hdr-output-debug-views-urp.md), URP writes the output of [post-processing](integration-with-post-processing.md) to this texture instead of `backBufferColor`. | `uberPost` and `finalPost` |

## Depth data

| **Property** | **Texture** | **URP shader pass that writes to the texture** |
| --- | --- | --- |
| `activeDepthTexture` | The **depth buffer** the GPU currently renders to. This is either `backBufferDepth` or `cameraDepth`. | Any pass, depending on your settings |
| `backBufferDepth` | The depth buffer of the screen back buffer. If you target `backBufferDepth`, any changes you make are overwritten when URP **blits** `cameraDepth` to the back buffer near the end of a frame. | Any pass, depending on your settings |
| `cameraDepth` | The depth buffer from the **render texture** the camera currently renders to. Avoid targeting `cameraDepth`, because URP uses this buffer for most of its own rendering. To use this texture, set `ScriptableRenderPass.requiresIntermediateTexture` to `true`. | Any pass, depending on your settings |
| `cameraDepthTexture` | A depth texture copy of the depth buffer, if you enable **Depth Priming Mode** in the [renderer](urp-universal-renderer.md) or **Depth Texture** in the active [URP Asset](universalrp-asset.md). If you use the [Deferred render path](rendering/deferred-rendering-path-landing.md), `cameraDepthTexture` is a color format instead of a depth format. | `CopyDepth` or `DepthPrepass` |
| `cameraNormalsTexture` | The scene normals texture. Contains the scene depth for objects with shaders that have a `DepthNormals` pass. | `DepthNormals` prepass |

## Shadow data

| **Property** | **Texture** | **URP shader pass that writes to the texture** |
| --- | --- | --- |
| `additionalShadowsTexture` | The additional shadow map. | `ShadowCaster` |
| `mainShadowsTexture` | The main shadow map. | `ShadowCaster` |

## Decal data

| **Property** | **Texture** | **URP shader pass that writes to the texture** |
| --- | --- | --- |
| `dBuffer` | The Decals texture. For more information about the decals texture, refer to [DBuffer](renderer-feature-decal-reference.md#dbuffer). | `Decals` |
| `dBufferDepth` | The Decals depth texture. Refer to [DBuffer](renderer-feature-decal-reference.md#dbuffer). | `Decals` |

## Motion vector data

| **Property** | **Texture** | **URP shader pass that writes to the texture** |
| --- | --- | --- |
| `motionVectorColor` | The motion vectors color texture. Refer to [motion vectors](features/motion-vectors.md). | `Camera Motion Vectors` and `MotionVectors` |
| `motionVectorDepth` | The motion vectors depth texture. Refer to [motion vectors](features/motion-vectors.md). | `Camera Motion Vectors` and `MotionVectors` |

## Other data

| **Property** | **Texture** | **URP shader pass that writes to the texture** |
| --- | --- | --- |
| `gBuffer` | The G-buffer textures. Refer to [G-buffer](rendering/g-buffer-layout.md). | `GBuffer` |
| `internalColorLut` | The internal look-up textures (LUT) texture. | `InternalLut` |
| `overlayUITexture` | The overlay UI texture. | `DrawScreenSpaceUI` |
| `renderingLayersTexture` | The Rendering Layers texture. Refer to [Rendering layers](features/rendering-layers.md) | `DrawOpaques` or the `DepthNormals` prepass, depending on your settings. |
| `ssaoTexture` | The Screen Space Ambient Occlusion (SSAO) texture. Refer to [Ambient occlusion](post-processing-ssao.md). | `SSAO` |

## Additional resources

* [Using frame data in the render graph system](render-graph-frame-data.md)