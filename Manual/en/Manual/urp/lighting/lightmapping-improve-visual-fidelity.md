# Improve the visual fidelity of lightmaps

Use bicubic sampling to improve the visual fidelity of **lightmaps**.

Bicubic sampling smooths sharp or jagged edges in your lightmap to improve its visual fidelity. Bicubic sampling is helpful for low resolution lightmaps and shadows. However, it performs more texture sampling so it introduces additional performance overhead on lower-end platforms.

![A black hexagonal shadow that has jagged edges, without bicubic sampling](../../../uploads/Main/bicubic-sampling-before.png)

A black hexagonal shadow that has jagged edges, without bicubic sampling

![A black hexagonal shadow that has smooth edges, with bicubic sampling](../../../uploads/Main/bicubic-sampling-after.png)

A black hexagonal shadow that has smooth edges, with bicubic sampling

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Bicubic Lightmap Sampling** | Yes | Yes | No | No |

## Use bicubic sampling

To use bicubic sampling in your project, enable the **Use Bicubic Lightmap Sampling** property in [URP graphics settings](urp/urp-global-settings.md) or [HDRP graphics settings](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest?subfolder=/manual/Default-Settings-Window.html).

## Fix texture bleeding

Bicubic sampling can cause texture bleeding, where Unity samples unintended texels from neighboring regions of a texture, leading to visual artifacts. To prevent texture bleeding, there needs to be enough margin between [lightmap UV charts](LightingGiUvs-landing.md). To fully remove texture bleeding artifacts, set the number of texels between each UV chart to at least 2 for bilinear filtering and at least 4 for bicubic filtering.

To change the margin between charts, refer to [Fix light bleeding in lightmaps](ProgressiveLightmapper-UVOverlap).

## Additional resources

* [Introduction to lightmaps and baking](Lightmappers.md)
* [Configuring lightmaps and baking](Lightmapping-configure.md)
* [Optimize baking](GPUProgressiveLightmapper.md)
* [Introduction to lightmap UVs](LightingGiUvs.md)