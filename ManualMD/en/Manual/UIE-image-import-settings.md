# Image import settings

After you have imported an image to your project, for the most intuitive results, it’s recommended that you apply certain import settings for Textures, **Sprites**, and Vector images before you use them as a background for a **visual element** in the UI Builder.

## Texture

The recommended import settings for a **Texture** image that you use as a background for a visual element are as follows.

Use the following settings for Default Texture types, and Editor GUI and Legacy GUI Texture types:

| **Property** | **Value** |
| --- | --- |
| **Non-Power of 2** | None |
| ****Compression**** | None |
| **Alpha Is Transparency** | true |

Use the following settings for Sprite (2D and UI) Texture types:

| **Property** | **Value** |
| --- | --- |
| **Compression** | None |
| **Alpha Is Transparency** | true |
| **Sprite Mode** | Single |
| ****Mesh** Type** | Tight |

## Sprites

The recommended import settings for a **Sprite** image that you use as a background for a visual element:

| **Property:** | **Value:** |
| --- | --- |
| **Compression** | None |
| **Alpha Is Transparency** | true |
| **Sprite Mode** | Multiple if file contains multiple sprites, Single otherwise. |
| **Mesh Type** | Tight |

## Vector images

When you import a vector image, Unity generates a **VectorImage** asset. You can then use this asset in your UI. The VectorImage asset is a lightweight representation of the SVG file that Unity can render efficiently.

### Recommended import settings

The following are the recommended import settings for a SVG **Vector** image that you use as a background for a visual element:

| **Property:** | **Value:** |
| --- | --- |
| **Generated Asset Type** | UI Toolkit Vector Image |
| **Gradient Resolution** | Texture size used to store gradients |
| **Tessellation Mode** | Basic Triangulation |
| **Tessellation Settings** | Basic |

**Note**: When you use **Basic** tessellation settings, you only need to provide a **Target Resolution** (use the lowest value that produces satisfactory results) and a **Zoom Factor**. The importer then automatically configures the advanced settings to ensure your SVG document renders at a high enough tessellation for that resolution at that zoom factor.

### Tessellation Mode

The **Tessellation Mode** determines how the SVG file is converted into a format that Unity can render. The available modes are:

* **Basic Triangulation**: This mode provides a good balance between performance and visual fidelity. It uses a simple triangulation algorithm to convert the SVG paths into triangles.
* **Antialiased Arc Encoding**: This mode is designed to handle curves more smoothly by using a higher density of vertices along the arcs, resulting in better visual quality for curved paths.

### Advanced tessellation

If you want full control over the tessellation of the SVG document, you can specify the following advanced settings:

| **Property** | **Function** |
| --- | --- |
| **Step Distance** | Distance at which the importer generates vertices along the paths. Lower values result in a more dense tessellation. |
| **Sampling Steps** | Number of samples the importer evaluates on paths. More samples can result in more precise curves, especially when the curves have sharp corners. |
| **Max Cord Enabled** | Whether to enable the maximum chord length constraint. |
| **Max Cord Deviation** | Distance on the cord to a straight line between two points after which the importer generates more tessellation. |
| **Max Tangent Enabled** | Whether to enable the maximum tangent length constraint. |
| **Max Tangent Angle** | Maximum tangent angle in degrees after which the importer generates tessellation. |

The following image shows the constraints for the SVG importer:

![Max Cord](../uploads/Main/uss/constraints.png)

Max Cord

The importer subdivides curves for as long as every enabled constraint isn’t satisfied.

**Tip**:

* You can [apply default presets to Assets by folder](DefaultPresetsByFolder.md) to automatically set your desired import settings.
* All image types are subject to [dynamic atlasing](../ScriptReference/UIElements.DynamicAtlasSettings.md) if they’re not already in an atlas. (An image is in an atlas if imported as a Sprite with **Sprite Mode** set to **Multiple**, or if you have manually added it to a [Sprite Atlas](sprite/atlas/atlas-landing.md) asset.) You can configure dynamic atlasing in a [Panel Settings](UIE-Runtime-Panel-Settings.md) asset.

## Additional resources

* [UIE-uxml-element-Image](UIE-uxml-element-Image.md)