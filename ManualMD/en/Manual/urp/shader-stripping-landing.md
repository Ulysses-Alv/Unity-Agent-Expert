# Reducing shader variants in URP

The **shaders** in the Universal **Render Pipeline** (URP) use [shader keywords](../shader-keywords.md) to support many different features, which might mean Unity compiles a lot of [shader variants](../shader-variants.md).

The following resources are about speeding up builds by reducing the number of shader variants URP compiles.

For more information, refer to [Reduce shader variants](../shader-variant-stripping.md), which applies to all render pipelines.

| **Page** | **Description** |
| --- | --- |
| [Check how many shader variants your build has](shader-stripping-check.md) | Log how many shader variants Unity compiles and strips. |
| [Strip shader variants](shader-stripping-features.md) | Remove shader variants for features you don’t use. |
| [Enable dynamic branching in shaders](shader-stripping-fog.md) | Make Unity use dynamic branching in prebuilt or custom shaders, instead of keywords and shader variants. |
| [Settings and keywords reference for shader stripping](shader-stripping-features-and-keywords.md) | Explore the settings and shader keywords you can use to strip shader variants. |

## Additional resources

* [Graphics performance and profiling in URP](../graphics-performance-and-profiling-in-urp.md)
* [Troubleshooting shaders](../shader-troubleshooting.md)