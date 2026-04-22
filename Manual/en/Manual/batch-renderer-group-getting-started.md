# Set up your project for the BatchRendererGroup API in URP

Before you use BRG, your project must support it. BRG requires your project to:

* Use the SRP Batcher. To enable the SRP Batcher, see [Using the SRP Batcher](SRPBatcher-Enable.md).
* Keep BRG [shader variants](shader-variants.md). To do this, select **Edit** > **Project Settings** > **Graphics**, and set **BatchRendererGroup variants** to **Keep all**.
* If your project uses URP, it’s best practice to disable the **Strip Unused Variants** [Global Setting](urp/urp-global-settings.md). This helps to avoid problems with Unity stripping necessary DOTS Instancing variants. For more information, see [DOTS Instancing shaders](dots-instancing-shaders.md). To find this setting, select **Edit** > **Project Settings** > **URP Global Settings**.
* Allow [unsafe code](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/unsafe-code). To do this, enable the **Allow ‘unsafe’ Code** [Player Setting](class-PlayerSettings.md).

**Note:** The BatchRendererGroup uses [DOTS Instancing shaders](dots-instancing-shaders.md), but it doesn’t require any DOTS packages. The name reflects the new data-oriented way to load instance data, and also helps with backward compatibility with existing Hybrid Renderer compatible **shaders**.

For information on how to use BRG to create a basic renderer, see [Creating a renderer with BatchRendererGroup](batch-renderer-group-creating-a-renderer.md).