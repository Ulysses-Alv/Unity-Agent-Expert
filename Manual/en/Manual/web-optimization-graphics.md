# Recommended Graphics settings to optimize your Web build

Use the following recommended **Graphics** settings to optimize your builds for the Unity Web platform.

Find these settings under **Edit** > **Project Settings** > **Graphics**. For more information on each setting, refer to the details in [Graphics](class-GraphicsSettings.md).

| **Setting** | **Recommended setting** | **Description** |
| --- | --- | --- |
| [Lightmap modes](#lightmap-modes) | Automatic (default) | Automatically strip variants that aren’t used. |
| [Fog modes](#fog-modes) | Automatic (default) | Automatically strip variants that aren’t used. |
| [Instancing Variants](#instancing-variants) | Strip Unused (default) | Only include a **shader** variant if at least one material uses that variant. |
| [BatchRendererGroup Variants](#brg-variants) | Strip all | Removes all BatchRendererGroup shader variants. |
| [Always Included Shaders](#always-included-shaders) |  | Remove any shaders from this list that aren’t used in your project |

## Lightmap Modes

Use the ****Lightmap** Modes** setting to change the shader variant stripping behavior for lightmap-related shaders. For examples of lightmap-related shader variants, refer to [Graphics](class-GraphicsSettings.md).

The recommended setting for each of these settings is **Automatic (default)**, which removes any shader variants that aren’t used in your build. This setting is useful because if you don’t strip unused shaders, it can increase build times, file size and memory usage.

## Fog Modes

Use the **Fog Modes** setting to change the shader variant stripping behavior for shaders that relate to built-in Unity fog effect. For examples of fog shader variants and more information, refer to [Graphics](class-GraphicsSettings.md).

The recommended setting is **Automatic (Default)**, which removes unused fog shaders from your build. It’s best to remove unused shaders because they can increase build times, file size and memory usage.

## Instancing Variants

Use the **Instancing Variants** setting to change how much Unity should strip shader variants for GPU instancing. For more information, refer to [Graphics](class-GraphicsSettings.md).

The recommended setting is **Strip Unused**, which removes any instancing variant shaders that your project doesn’t use. It’s best to remove unused shaders because they can increase build times, file size and memory usage.

To keep some unused shaders for future use or for other shaders to reference them, choose **Keep All** instead.

## Batch renderer group variants

Use the **Batch renderer group variants** setting to change the shader variant stripping behavior for shaders related to batch renderer groups (BRGs). For more information about BRGs, refer to [BatchRendererGroup](batch-renderer-group.md).

If your project doesn’t use BRGs, set **Batch renderer group variants** to **Strip all**, which removes all BRG shader variants. Unused shaders can increase build times, file size and memory usage. If your project uses BRGs, ignore this recommendation.

## Always included shaders

**Always included shaders** is a list of shaders for which Unity includes all possible variants in every build. If your project doesn’t use any of the shaders in the list, it’s best to remove them from the list because unused shaders can increase build times, file size and memory usage.

For more information, refer to [Graphics](class-GraphicsSettings.md).

### Edit the Always included shaders list via C

To change the **Always included shaders** list via script instead, create a list of shaders you want to have in the list and assign it like this:

```
GraphicsSettings.alwaysIncludedShaders = newShadersList.ToArray();
```

## Additional resources

* [Graphics](class-GraphicsSettings.md)
* [Recommended Player settings to optimize your Web build](web-optimization-player.md)
* [Recommended Quality settings to optimize your Web build](web-optimization-quality.md)
* [Remove unused resources from your Web build](web-optimization-remove-resources.md)
* [Optimize your Web build](web-optimization.md)
* [Optimize Web platform for mobile](web-optimization-mobile.md)