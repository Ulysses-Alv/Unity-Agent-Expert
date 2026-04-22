# Enable 2D lighting with the Tilemap Renderer in URP

To enable 2D lighting in URP, set up the required settings to use the Tilemap Renderer component.

You can use the Tilemap Renderer with URP 2D to enable [2D lighting](../Lights-2D-intro.md) on [tiles](../../tilemaps/tiles-for-tilemaps/tiles-landing.md) and [tilemaps](../../tilemaps/work-with-tilemaps/work-with-tilemaps-landing.md), especially [isometric tilemaps](../../tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md) which simulate pseudo-depth. Doing so requires you to setup your project and adjust the Tilemap Renderer’s settings in the following ways.

When you select the [2D Renderer Data asset](../2DRendererData-overview.md) for your project’s [Scriptable Render Pipeline](../../scriptable-render-pipeline-introduction.md), the Renderer Data asset assumes control of the Tilemap Renderer’s **Transparency Sort Mode** property settings and requires you to adjust the settings under the 2D Renderer Data asset’s property settings instead of in the **Project Settings**.

To optimize the rendering of the Tilemap Renderer component with the 2D lighting system, Unity can batch the rendering of the [Tilemap Renderer component](../../tilemaps/work-with-tilemaps/tilemap-renderer-reference.md) with the [Scriptable Render Pipeline Batcher](../../SRPBatcher.md) to improve the rendering performance of the Tilemap Renderer with other **sprite** renderers with the same rendering characteristics.

## Adjusting the Transparency Sort Mode settings in the 2D Renderer Data asset

When you [create an isometric tilemap](../../tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md), one of the steps requires you to set the **Transparency Sort Mode** (**Edit** > **Project Settings…** > **Graphics** > **Camera Settings**) to **Custom Axis** and to set it to the [required values](https://docs.unity3d.com/Manual/Tilemap-Isometric-CreateIso.html#customaxis) to have Unity render the Tiles with the pseudo-depth of an isometric perspective.

![The default location of the Transparency Sort Mode settings when no specific URP Pipeline is selected.](../../../uploads/urp/2D/non-renderer2d-project-settings.png)

The default location of the **Transparency Sort Mode** settings when no specific URP Pipeline is selected.

If you want to use the 2D Renderer with the Tilemap Renderer in your Project, first create the URP Asset and its associated 2D Renderer Data asset by right-clicking the Asset window and go to **Create > Rendering > URP Asset (with 2D Renderer)**. Then go to your Project’s **Scriptable Render Pipeline Settings** (menu: **Edit > Project Settings… > Graphics**) and select the Universal Render Pipeline (URP) 2D asset.

When you do so, the ****Camera** Settings** (including the **Transparency Sort Mode** property) become hidden.

![Camera Settings hidden after selecting a Universal Render Pipeline Asset.](../../../uploads/urp/2D/renderer-2d-selected.png)

Camera Settings hidden after selecting a Universal Render Pipeline Asset.

The 2D Renderer Data asset now controls the **Transparency Sort Mode** property settings, and the values set in the active Renderer 2D Data asset supersedes the values set in the **Project Settings**. Select the 2D Renderer Data asset go to its **Inspector** window. In the General section, set **Transparency Sort Mode** to **Custom Axis**.

![Select Custom Axis in the 2D Renderer Data asset.](../../../uploads/urp/2D/renderer-2d-custom-axis.png)

Select **Custom Axis** in the 2D Renderer Data asset.

The **Transparency Sort Axis** property settings then become available. Use the same values for the **Transparency Sort Axis** as those used to for the [rendering of tiles on an isometric tilemap](https://docs.unity3d.com/Manual/Tilemap-Isometric-CreateIso.html#customaxis).

![The Transparency Sort Axis setting for isometric tilemaps in the Renderer 2D asset properties.](../../../uploads/urp/2D/renderer-data-asset-properties.png)

The **Transparency Sort Axis** setting for isometric tilemaps in the Renderer 2D asset properties.

## Enabling the Scriptable Render Pipeline Batcher

To prepare the Tilemap Renderer for SRP batching:

1. Fulfill the requirements and steps for enabling the [SRP Batcher](https://docs.unity3d.com/Manual/SRPBatcher.html#using-the-srp-batcher) in URP.
2. Select the Tilemap Renderer component and go to its **Mode** property setting.
3. Select either **Individual** or **SRP Batch** (only supported in the Universal Render Pipeline version 15 onwards). **Note**: **Chunk** mode is incompatible with the SRP batcher.

## Additional resources

* [2D Sorting](../../2d-renderer-sorting.md)
* [Scriptable Render Pipeline fundamentals](../../scriptable-render-pipeline-introduction.md)
* [Isometric Tilemaps](../../tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md)