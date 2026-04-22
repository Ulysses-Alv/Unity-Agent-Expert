# Convert quality settings from the Built-In Render Pipeline to URP

This page provides recommended URP graphics [Quality Level](../../class-QualitySettings.md) settings for **Low** and **High** quality level values. These settings approximately match the performance of the equivalent **Low** and **High** default presets in the Built-In **Render Pipeline**.

URP changes the implementation of many features and settings, and as a result they often have a different performance impact to the Built-In Render Pipeline equivalent. When you upgrade a project from the Built-In Render Pipeline to URP, your existing quality levels might provide a different level of performance, and you might need to update or create new quality levels for your project. You can use the values on this page as a starting point.

This page is split into the following sections:

* [Project Settings](#project-settings)
* [URP Asset](#urp-asset)

> **Note**: In URP, many quality level settings have moved from the Project Settings window to the URP Asset. For more information on where to find these settings in URP projects, refer to [Built-In Render Pipeline Quality Settings Reference](quality-settings-location.md).

## Project Settings

You can change the following settings in [Project Settings](../../class-QualitySettings.md) under **Project Settings** > **Quality**.

| **Setting** | **“Low” preset value** | **“High” preset value** |
| --- | --- | --- |
| **Rendering** |  |  |
| Real-time **Reflection Probes** | No | Yes |
| Resolution Scaling Fixed DPI Factor | 1 | 1 |
| **VSync** Count | Don’t Sync | Every V Blank |
| **Textures** |  |  |
| Global Mipmap Limit | Half Resolution | Full Resolution |
| Anisotropic Textures | Disabled | Disabled |
| Texture Streaming | No | No |
| ****Particles**** |  |  |
| Particle Raycast Budget | 16 | 256 |
| ****Terrain**** |  |  |
| **Billboards** Face **Camera** Position | No | Yes |
| **Shadows** |  |  |
| **Shadowmask** Mode | Shadowmask | **Distance Shadowmask** |
| **Async Asset Upload** |  |  |
| Time Slice | 2 | 2 |
| Buffer Size | 16 | 16 |
| Persistent Buffer | Yes | Yes |
| ****Level of Detail**** |  |  |
| LOD Bias | 0.4 | 1 |
| Maximum LOD level | 0 | 0 |
| **Meshes** |  |  |
| Skin Weights | 4 Bones | Unlimited |

## URP Asset

You can change the following settings inside any [URP Asset](./../universalrp-asset.md).

| **Setting** | **“Low” preset value** | **“High” preset value** |
| --- | --- | --- |
| **Rendering** |  |  |
| Depth Texture | No | No |
| Opaque Texture | No | No |
| Terrain Holes | Yes | Yes |
| **Quality** |  |  |
| **HDR** | Yes | Yes |
| Anti Aliasing (MSAA) | Disabled | 2x |
| Render Scale | 1 | 1 |
| **Lighting** |  |  |
| Main Light | Per **Pixel** | Per Pixel |
| Cast Shadows | No | Yes |
| Shadows Resolution | N/A | 2048 |
| Additional Lights | Disabled | Per Pixel |
| Per Object Limit | N/A | 4 |
| Cast Shadows | N/A | Yes |
| Shadow Atlas Resolution | N/A | 2048 |
| Shadow Resolution Tiers | N/A |  |
| Low | N/A | 512 |
| Medium | N/A | 1024 |
| High | N/A | 2048 |
| Cookie Atlas Resolution | N/A | 2048 |
| Cookie Atlas Format | N/A | Color High |
| Reflection Probes |  |  |
| Probe Blending | No | Yes |
| Box Projection | No | No |
| **Shadows** |  |  |
| Max Distance | N/A | 50 |
| Cascade Count | N/A | 3 |
| Split 1 | N/A | 12.5 |
| Split 2 | N/A | 33.8 |
| Last Border | N/A | 3.8 |
| Working Unit | N/A | Metric |
| Depth Bias | N/A | 1 |
| Normal Bias | N/A | 1 |
| Soft Shadows | N/A | Yes |
| ****Post-processing**** |  |  |
| Grading Mode | Low Dynamic Range | Low Dynamic Range |
| LUT Size | 16 | 32 |
| Fast sRGB/Linear Conversion | No | No |

## Additional resources

* [Find Quality Settings in URP](./quality-settings-location.md)
* [URP Asset](./../universalrp-asset.md)