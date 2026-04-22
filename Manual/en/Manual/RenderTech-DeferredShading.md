# Deferred rendering path in the Built-In Render Pipeline

This page details the **Deferred Shading** [rendering path](RenderingPaths.md) in Unity’s Built-in **Render Pipeline**. See [Wikipedia: deferred shading](http://en.wikipedia.org/wiki/Deferred_shading) for an introductory technical overview of deferred shading.

## Overview

When using deferred shading, there is no limit on the number of lights that can affect a **GameObject**. All lights are evaluated per-pixel, which means that they all interact correctly with **normal maps**, etc. Additionally, all lights can have cookies and shadows.

Deferred shading has the advantage that the processing overhead of lighting is proportional to the number of **pixels** the light shines on. This is determined by the size of the light volume in the **Scene** regardless of how many GameObjects it illuminates. Therefore, performance can be improved by keeping lights small. Deferred shading also has highly consistent and predictable behaviour. The effect of each light is computed per-pixel, so there are no lighting computations that break down on large triangles.

On the downside, deferred shading has no real support for anti-aliasing and can’t handle semi-transparent GameObjects (these are rendered using [forward](RenderTech-ForwardRendering.md) rendering). There is also no support for the **Mesh** Renderer’s Receive Shadows flag and **culling masks** are only supported in a limited way. You can only use up to four culling masks. That is, your culling **layer mask** must at least contain all layers minus four arbitrary layers, so 28 of the 32 layers must be set. Otherwise you get graphical artifacts.

## Requirements

It requires a graphics card with Multiple Render Targets (MRT), **Shader** Model 3.0 (or later) and support for Depth **render textures**. Most PC graphics cards made after 2006 support deferred shading, starting with GeForce 8xxx, Radeon X2400, Intel G45.

On mobile, deferred shading is supported on all devices running at least OpenGL ES 3.0.

**Note**: Deferred rendering isn’t supported when using Orthographic projection. If the **Camera**’s projection mode is set to Orthographic, the Camera falls back to Forward rendering.

**Note**: Deferred Rendering doesn’t support Single Pass Stereo Instancing in the Built-in Render Pipeline..

## Performance considerations

The rendering overhead of real-time lights in deferred shading is proportional to the number of pixels illuminated by the light and *not* dependent on Scene complexity. So small Point Lights or Spot Lights are very cheap to render and if they are fully or partially occluded by Scene GameObjects then they are even cheaper.

Of course, lights with shadows are much more expensive than lights without shadows. In deferred shading, shadow-casting GameObjects still need to be rendered once or more for each shadow-casting light. Furthermore, the lighting shader that applies shadows has a higher rendering overhead than the one used when shadows are disabled.

## Implementation details

Objects with Shaders that do not support deferred shading are rendered after deferred shading is complete, using the [forward rendering](RenderTech-ForwardRendering.md) path.

The default layout of the render targets (RT0 - RT4) in the G-buffer is listed below. Data types are placed in the various channels of each render target. The channels used are shown in parentheses.

* RT0, ARGB32 format: Diffuse color (RGB), occlusion (A).
* RT1, ARGB32 format: **Specular color** (RGB), smoothness (A).
* RT2, ARGB2101010 format: World space normal (RGB), unused (A).
* RT3, ARGB2101010 (non-HDR) or ARGBHalf (HDR) format: Emission + lighting + **lightmaps** + **reflection probes** buffer.
* Depth+**Stencil buffer**.

So the default G-buffer layout is 160 bits/pixel (non-HDR) or 192 bits/pixel (HDR).

If using the [Shadowmask](lighting-mode.md#shadowmask) or Distance Shadowmask modes for Mixed lighting, a fifth target is used:

* RT4, ARGB32 format: Light occlusion values (RGBA).

And thus the G-buffer layout is 192 bits/pixel (non-HDR) or 224 bits/pixel (HDR).

If the hardware does not support five concurrent rendertargets then objects using shadowmasks will fallback to the forward rendering path.
Emission+lighting buffer (RT3) is logarithmically encoded to provide greater dynamic range than is usually possible with an ARGB32 texture, when the Camera is not using **HDR**.

Note that when the Camera is using HDR rendering, there’s no separate rendertarget being created for Emission+lighting buffer (RT3); instead the rendertarget that the Camera renders into (that is, the one that is passed to the image effects) is used as RT3.

### G-Buffer pass

The G-buffer pass renders each GameObject once, into textures that store diffuse and specular colors, surface smoothness, world space normals, and emission, ambient, reflections, and lightmaps. Unity sets up G-buffer textures as global shader properties for later access by shaders. The property names are `_CameraGBufferTexture0` to `_CameraGBufferTexture3`.

### Lighting pass

The lighting pass computes lighting based on G-buffer and depth. Lighting is computed in screen space, so the time it takes to process is independent of Scene complexity. Lighting is added to the emission buffer.

Point Lights and Spot Lights that do not cross the Camera’s near plane are rendered as 3D shapes, with the Z buffer’s test against the Scene enabled. This makes partially or fully occluded Point Lights and Spot Lights very cheap to render. Directional Lights and Point or Spot Lights that cross the near plane are rendered as fullscreen **quads**.

If a light has shadows enabled then they are also rendered and applied in this pass. Note that shadows do not come for “free”; shadow casters need to be rendered and a more complex light shader must be applied.

The only lighting model available is Standard. If you need a different lighting model, you can create a customized version of the default deferred shader file. Follow these steps:

1. [Download the source code](StandardShaderMakeYourOwn.md) for the Built-In Render Pipeline shaders.
2. Find the `Internal-DeferredShading.shader` file in the `DefaultResourcesExtra` folder.
3. Edit the file to change the lighting model.
4. In the **Project** window, add the file to your **Assets** folder, in a folder called **Resources**.
5. From the main menu, select **Edit** > **Project Settings** > **Graphics**.
6. In the **Built-in Shader Settings** section, set **Deferred** to **Custom shader**, then select your shader file.

---

* 2017–06–08 Page published
* Light Modes (**Shadowmask** and **Distance Shadowmask**) added in 5.6