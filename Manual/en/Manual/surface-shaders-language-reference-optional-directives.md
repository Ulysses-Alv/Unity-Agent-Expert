# Surface Shader optional directives reference for the Built-In Render Pipeline

**Transparency and alpha testing** is controlled by `alpha` and `alphatest` directives. Transparency can typically be of two kinds: traditional alpha blending (used for fading objects out) or more physically plausible “premultiplied blending” (which allows semitransparent surfaces to retain proper specular reflections). Enabling semitransparency makes the generated surface **shader** code contain [blending](SL-Blend.md) commands; whereas enabling alpha cutout will do a fragment discard in the generated **pixel** shader, based on the given variable.

* `alpha` or `alpha:auto` - Will pick fade-transparency (same as `alpha:fade`) for simple lighting functions, and premultiplied transparency (same as `alpha:premul`) for physically based lighting functions.
* `alpha:blend` - Enable alpha blending.
* `alpha:fade` - Enable traditional fade-transparency.
* `alpha:premul` - Enable premultiplied alpha transparency.
* `alphatest:VariableName` - Enable alpha cutout transparency. Cutoff value is in a float variable with VariableName. You’ll likely also want to use `addshadow` directive to generate proper shadow caster pass.
* `keepalpha` - By default opaque **surface shaders** write 1.0 (white) into alpha channel, no matter what’s output in the Alpha of output struct or what’s returned by the lighting function. Using this option allows keeping lighting function’s alpha value even for opaque surface shaders.
* `decal:add` - Additive decal shader (e.g. terrain AddPass). This is meant for objects that lie atop of other surfaces, and use additive blending. See [Surface Shader Examples](SL-SurfaceShaderExamples.md)
* `decal:blend` - Semitransparent decal shader. This is meant for objects that lie atop of other surfaces, and use alpha blending. See [Surface Shader Examples](SL-SurfaceShaderExamples.md)

**Custom modifier functions** can be used to alter or compute incoming vertex data, or to alter final computed fragment color.

* `vertex:VertexFunction` - Custom vertex modification function. This function is invoked at start of generated **vertex shader**, and can modify or compute per-vertex data. See [Surface Shader Examples](SL-SurfaceShaderExamples.md).
* `finalcolor:ColorFunction` - Custom final color modification function. See [Surface Shader Examples](SL-SurfaceShaderExamples.md).
* `finalgbuffer:ColorFunction` - Custom deferred path for altering G-buffer content.

**Shadows and Tessellation** - additional directives can be given to control how shadows and tessellation is handled.

* `addshadow` - Generate a shadow caster pass. Commonly used with custom vertex modification, so that shadow casting also gets any procedural vertex animation. Often shaders don’t need any special shadows handling, as they can just use shadow caster pass from their fallback.
* `fullforwardshadows` - Support all light shadow types in [Forward](RenderTech-ForwardRendering.md) **rendering path**. By default shaders only support shadows from one directional light in **forward rendering** (to save on internal shader variant count). If you need point or Spot Light shadows in forward rendering, use this directive.
* `tessellate:TessFunction` - use DX11 GPU tessellation; the function computes tessellation factors. See [Surface Shader Tessellation](SL-SurfaceShaderTessellation.md) for details.

**Code generation options** - by default generated surface shader code tries to handle all possible lighting/shadowing/**lightmap** scenarios. However in some cases you know you won’t need some of them, and it is possible to adjust generated code to skip them. This can result in smaller shaders that are faster to load.

* `exclude_path:deferred`, `exclude_path:forward` - Do not generate passes for given rendering path ([Deferred Shading](RenderTech-DeferredShading.md), [Forward](RenderTech-ForwardRendering.md) respectively).
* `noshadow` - Disables all shadow receiving support in this shader.
* `noambient` - Do not apply any ambient lighting or **light probes**.
* `novertexlights` - Do not apply any light probes or per-vertex lights in Forward rendering.
* `nolightmap` - Disables all lightmapping support in this shader.
* `nodynlightmap` - Disables runtime dynamic **global illumination** support in this shader.
* `nodirlightmap` - Disables directional lightmaps support in this shader.
* `nofog` - Disables all built-in Fog support.
* `nometa` - Does not generate a “meta” pass (that’s used by lightmapping & dynamic global illumination to extract surface information).
* `noforwardadd` - Disables [Forward](RenderTech-ForwardRendering.md) rendering additive pass. This makes the shader support one full directional light, with all other lights computed per-vertex/SH. Makes shaders smaller as well.
* `nolppv` - Disables **Light Probe Proxy Volume** support in this shader.
* `noshadowmask` - Disables Shadowmask support for this shader (both [Shadowmask](lighting-mode.md#shadowmask) and Distance Shadowmask).

**Miscellaneous options**

* `softvegetation` - Makes the surface shader only be rendered when Soft Vegetation is on.
* `interpolateview` - Compute view direction in the vertex shader and interpolate it; instead of computing it in the pixel shader. This can make the pixel shader faster, but uses up one more texture interpolator.
* `halfasview` - Pass half-direction vector into the lighting function instead of view-direction. Half-direction will be computed and normalized per vertex. This is faster, but not entirely correct.
* `approxview` - Removed in Unity 5.0. Use `interpolateview` instead.
* `dualforward` - Use dual lightmaps in [forward](RenderTech-ForwardRendering.md) rendering path.
* `dithercrossfade` - Makes the surface shader support dithering effects. You can then apply this shader to **GameObjects** that use an [LOD Group](class-LODGroup.md) component configured for cross-fade transition mode.

To see what exactly is different from using different options above, it can be helpful to use “Show Generated Code” button in the [Shader Inspector](class-Shader.md).