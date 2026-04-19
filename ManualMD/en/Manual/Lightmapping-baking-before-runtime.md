# Baking lightmaps before runtime

Resources for precalculating the lit colors of **GameObjects** and storing them in textures called **lightmaps**.

| **Page** | **Description** |
| --- | --- |
| [Lightmaps and baking](Lightmappers.md) | Learn about precalculating the color and brightness of surfaces in a **scene**, then storing the result in a texture called a lightmap for later use. |
| [Set up your scene and lights for baking](Lightmapping.md) | Set up meshes and GameObjects so that they store precalculated color in lightmaps. |
| [Preview baked lighting](Lightmapping-preview.md) | Edit a scene and preview changes to lightmaps and lighting without affecting the baked lightmaps. |
| [Bake lighting](Lightmapping-bake.md) | Generate lightmaps for your scene. |
| [Configuring baking lightmaps](Lightmapping-configure.md) | Resources for selecting CPU or GPU lightmapping, and changing how GameObjects store lighting in lightmaps. |
| [Troubleshooting baked lightmaps](Lightmapping-troubleshooting.md) | Solve common issues with baked lightmaps, such as hard edges or light bleeding. |
| [Troubleshooting materials missing specular highlights](ts-missing-spec-response.md) | Troubleshoot issues causing materials missing specular highlights. |
| [Optimize baked lightmaps](GPUProgressiveLightmapper.md) | Speed up baking time by configuring baking settings or selecting different GPUs for rendering and baking. |
| [Lightmapping settings in the Lighting Settings Asset reference](Lightmaps-reference.md) | Explore the properties and settings in the Lighting Settings Asset window to customize baking lightmaps. |

## Additional resources

* [Configure lightmapping with a Lighting Settings Asset](global-illumination-configure.md)
* [Lighting Settings Asset Inspector window reference](class-LightingSettings.md)