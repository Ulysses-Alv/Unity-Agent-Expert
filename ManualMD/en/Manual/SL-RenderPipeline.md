# Surface Shaders and rendering paths in the Built-In Render Pipeline

In the Built-in **Render Pipeline**, when using a Surface **Shader**, how lighting is applied and which [Passes](SL-Pass.md) of the shader are used depends on which [rendering path](RenderingPaths.md) is used. Each pass in a shader communicates its lighting type via [Pass Tags](SL-PassTags.md).

## Render pipeline compatibility

| **Feature name** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Surface Shaders** | No  For a streamlined way of creating Shader objects in URP, see [Shader Graph](shader-graph.md). | No  For a streamlined way of creating Shader objects in HDRP, see [Shader Graph](shader-graph.md). | No | Yes |

## Rendering paths

* In [Forward Rendering](RenderTech-ForwardRendering.md), `ForwardBase` and `ForwardAdd` passes are used.
* In [Deferred Shading](RenderTech-DeferredShading.md), `Deferred` pass is used.
* In [legacy Vertex Lit](RenderTech-VertexLit.md), `Vertex`, `VertexLMRGBM` and `VertexLM` passes are used.
* In any of the above, to render [Shadows](Shadows.md) or a depth texture, `ShadowCaster` pass is used.

## Forward Rendering path

`ForwardBase` pass renders ambient, **lightmaps**, main directional light and not important (vertex/SH) lights at once. `ForwardAdd` pass is used for any additive per-pixel lights; one invocation per object illuminated by such light is done. See [Forward Rendering](RenderTech-ForwardRendering.md) for details.

If forward rendering is used, but a shader does not have forward-suitable passes (i.e. neither `ForwardBase` nor `ForwardAdd` pass types are present), then that object is rendered just like it would in Vertex Lit path, see below.

## Deferred Shading path

`Deferred` pass renders all information needed for lighting (in built-in shaders: diffuse color, specular color, smoothness,
world space normal, emission). It also adds lightmaps, reflection probes and ambient lighting into the emission channel. See [Deferred Shading](RenderTech-DeferredShading.md) for details.

## Legacy Vertex Lit Rendering path

Since vertex lighting is most often used on platforms that do not support programmable shaders, Unity can’t create multiple shader variants internally to handle lightmapped vs. non-lightmapped cases. So to handle lightmapped and non-lightmapped objects, multiple passes have to be written explicitly.

* `Vertex` pass is used for non-lightmapped objects. All lights are rendered at once, using a fixed function OpenGL/Direct3D lighting model ([Blinn-Phong](http://en.wikipedia.org/wiki/Blinn-Phong_shading))
* `VertexLMRGBM` pass is used for lightmapped objects, when lightmaps are RGBM encoded (PC and consoles). No real-time lighting is applied; pass is expected to combine textures with a lightmap.
* `VertexLM` pass is used for lightmapped objects, when lightmaps are double-LDR encoded (mobile platforms). No real-time lighting is applied; pass is expected to combine textures with a lightmap.