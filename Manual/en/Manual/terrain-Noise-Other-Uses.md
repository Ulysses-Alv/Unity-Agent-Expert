# Other ways to use noise

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

Alongside the Noise height sculpting tool, you can use the noise mask with other brush-based tools, as a source for 2D textures and 3D assets, and as a shareable settings asset.

## Use a noise mask on brush-based tools

You can add noise as a mask to any height or texture brush to add variety without creating multiple brushes. The mask uses the same **Noise Editor** window that the Noise height sculpting tool uses, so it has the same options. For more information, refer to [Noise reference](terrain-Noise-Ref.md) and to [Brushes](class-Brush.md).

## Export noise output as 2D or 3D textures

You can also export a noise output as a texture to use as:

* A brush for any other brush-based tool from a 2D .png. For information about brush-based workflows and masks, refer to [Brushes](class-Brush.md).
* A **heightmap** for new **Terrain** tiles from a 2D .png.
* A 3D texture that you can use elsewhere in the Editor, for example to create a **shader**.

To export the noise output, use the **Noise Editor** window from either:

* **Window** > **Terrain Tools** > **Noise Editor**.
* **Brush Mask** > **Noise** mask > **Edit**.
* A Terrain tile’s **Inspector** window > **Paint Terrain** tool > tool dropdown > **Sculpt** > **Noise**. can export from the **Noise Editor** window for a Noise mask, or from the **Noise** height sculpting tool.

### Export options

| **Option** | **Description** |
| --- | --- |
| **Noise Settings** | The settings file the Noise Editor is using as a template for exporting noise settings. You can’t change this file. |
| **Texture Type** | Choose between a 2D texture (exported as .png), or a 3D asset. |
| **Dimensions** | Specify the texture’s size, where X is width and Y is height. If you set the **Texture Type** to **Texture 3D**, you can also specify the z-axis. |
| ****Texture Format**** | Choose the texture format. For a list of supported formats, refer to [GraphicsFormat](../ScriptReference/Experimental.Rendering.GraphicsFormat.md). |

## Reuse and share noise settings

The Unity Editor always saves noise settings to a settings asset file. You can create multiple setting files to reuse or share.

You can:

* Save your noise settings to the current noise settings asset.
* Revert changes you’ve made to the noise settings.
* Save your noise settings to a new noise settings asset.

To create an asset of noise definitions, use the **Noise Editor** window from either:

* **Window** > **Terrain Tools** > **Noise Editor**.
* **Brush Mask** > **Noise** mask > **Edit**.
* A Terrain tile’s **Inspector** window > **Paint Terrain** tool > tool dropdown > **Sculpt** > **Noise**.

To go back to the default definition asset:

1. In the Noise Settings Assets field, select the **prefab** picker (circle).
2. In the Select Noise Settings window, select None.

## Additional resources

* [Noise reference](terrain-Noise-Ref.md)
* [Sculpt height with the Noise tool](terrain-Sculpt-Noise-Landing.md)
* [Noise and fractal types](terrain-Noise-Types.md)