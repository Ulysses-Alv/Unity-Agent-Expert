# G-buffer layout in the Deferred and Deferred+ rendering paths in URP

The following illustration shows the data structure for each **pixel** of the render targets that Unity uses in the G-buffer in the Deferred and Deferred+ **rendering path**.

![Data structure of the render targets that Unity uses in the Deferred and Deferred+ rendering paths](../../../uploads/urp/rendering-deferred/data-structure-render-targets-g-buffer.png)

Data structure of the render targets that Unity uses in the Deferred and Deferred+ rendering paths

## Albedo (sRGB)

This field contains the albedo color in 24-bit sRGB format.

## MaterialFlags

This field is a bit field that contains material flags:

* Bit 0, **ReceiveShadowsOff**: if set, the pixel does not receive dynamic shadows.
* Bit 1, **SpecularHighlightsOff**: if set, the pixel does not receive specular highlights.
* Bit 2, **SubtractiveMixedLighting**: if set, the pixel uses the Subtractive Lighting Mode.
* Bit 3, **SpecularSetup**: if set, the material uses the specular workflow.

Bits 4 to 7 are reserved for future use.

## Specular

This field contains the following 24-bit values:

* Simple Lit material: RGB **specular color** stored in 24 bits.
* Lit material with metallic workflow: Reflectivity stored in 8 bits. The other 16 bits aren’t used.
* Lit material with specular workflow: RGB specular color stored in 24 bits.

## Occlusion

This field contains the baked occlusion value from baked lighting. For realtime lighting, Unity calculates the **ambient occlusion** value by combining the baked occlusion value with the screen space ambient occlusion (SSAO) value.

## Normal

This field contains the world space normals encoded in 24 bits. For more information, refer to [Enable accurate G-buffer normals in the Deferred rendering path](accurate-g-buffer-normals.md).

## Smoothness

This field stores the smoothness value for the Simple Lit and Lit materials.

## Emissive/GI/Lighting

This render target contains the material emissive output and baked lighting. Unity fills this field during the G-buffer render pass. During the **deferred shading** pass, Unity stores lighting results in this render target.

The default format is `B10G11R11_UFloatPack32`.

Unity uses `R16G6B16A16_SFloat` if either of the following is true:

* In the [URP Asset](../universalrp-asset.md#quality), **Quality Settings** > **HDR** is enabled but the build platform doesn’t support HDR.
* In the [Player settings window](../../class-PlayerSettings.md), **PreserveFramebufferAlpha** is enabled.

If Unity can’t use `B10G11R11_UFloatPack32` or `R16G6B16A16_SFloat`, it uses the [default HDR format](https://docs.unity3d.com/ScriptReference/Experimental.Rendering.DefaultFormat.HDR.html).

## ShadowMask

Unity adds this render target to the G-buffer layout when you set the [Lighting Mode](../../lighting-mode.md) to **Subtractive** or ****Shadowmask****.

## Rendering Layer Mask

Unity adds this render target to the G-buffer layout when you enable [Rendering Layers](../features/rendering-layers.md).

## Depth as Color

Unity adds this render target to the G-buffer layout when you enable Native Render Passes, and the platform supports them. Unity renders depth as a color into this render target. This render target has the following purpose:

* Improves performance on Vulkan devices.
* Lets Unity get the **depth buffer** on platforms that use the Metal graphics API, which doesn’t allow fetching the depth from the DepthStencil buffer within the same render pass.

The format of this render target is `GraphicsFormat.R32_SFloat`.

## DepthStencil

Unity reserves the four highest bits of this render target for the material type. For more information, refer to [URP ShaderLab Pass tags](../urp-shaders/urp-shaderlab-pass-tags.md#universalmaterialtype).

The format of this render target is `D32F_S8` or `D24S8`, depending on the platform.

## Additional resources

* [Render passes in the Deferred rendering path](render-passes-deferred.md)
* [Rendering Layers performance](../features/rendering-layers-introduction.md#performance)