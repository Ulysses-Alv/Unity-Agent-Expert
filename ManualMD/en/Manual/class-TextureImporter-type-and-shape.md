# Texture type and shape settings reference

![The Texture Import Settings window with all the settings up to the Advanced section highlighted.](../uploads/Main/texture-import-settings-top.png)

The Texture Import Settings window with all the settings up to the **Advanced** section highlighted.

**Note:** Some of the less commonly used properties are hidden by default. Expand the [Advanced](#advanced) section in the **Inspector** window to view these properties.

## Texture Type

Use the **Texture Type** property to select the type of Texture you want to create from the source image file. The other properties in the Texture Import settings window change depending on the value you set.

| **Property** | **Function** |
| --- | --- |
| **Default** | This is the most common setting used for all Textures. It provides access to most of the properties for Texture importing. For more information, see the [Default](texture-type-default.md) Texture type. |
| **Normal map** | The **Normal map** texture type formats the texture asset so it’s suitable for real-time normal mapping. For more information, see the [Normal map](texture-type-normal-map.md) texture type documentation.   For more information on normal mapping in general, see [Importing Textures](texture-type-normal-map.md). |
| **Editor GUI and Legacy GUI** | The **Editor GUI and Legacy GUI** texture type formats the texture asset so it’s suitable for HUD and GUI controls. For more information, see the [Editor GUI and Legacy GUI](texture-type-editor-gui-and-legacy-gui.md) texture type documentation. |
| **Sprite (2D and UI)** | The **Sprite (2D and UI)** texture type formats the texture asset so it’s suitable to use in 2D applications as a [Sprite](sprite/sprite-landing.md). For more information, see the [Sprite (2D and UI)](texture-type-sprite.md) texture type documentation. |
| **Cursor** | The **Cursor** texture type formats the texture asset so it’s suitable to use as a custom mouse cursor. For more information, see the [Cursor](texture-type-cursor.md) texture type documentation. |
| **Cookie** | The **Cookie** texture type formats the texture asset so it’s suitable to use as a [light cookie](Cookies.md) in the Built-in **Render Pipeline**. For more information, see the [Cookie](texture-type-cookie.md) texture type documentation. |
| **Lightmap** | The **Lightmap** texture type formats the texture asset so it’s suitable to use as a [Lightmap](class-LightmapParameters.md). This option enables encoding into a specific format (RGBM or dLDR depending on the platform) and a **post-processing** step on texture data (a push-pull dilation pass). For more information, see the [Lightmap](texture-type-lightmap.md) texture type documentation. |
| **Directional Lightmap** | The **Directional Lightmap** texture type formats the texture asset so it’s suitable to use as a directional [Lightmap](class-LightmapParameters.md). For more information, see the [Directional Lightmap](texture-type-directional-lightmap.md) texture type documentation. |
| **Shadowmask** | The **Shadowmask** texture type formats the texture asset so it’s suitable to use as a [shadowmask](lighting-mode.md#shadowmask). For more information, see the [Shadowmask](texture-type-shadowmask.md) texture type documentation. |
| **Single Channel** | The **Single Channel** texture type formats the texture asset so it only has one channel. For information on the properties available only for the this type, see the [Single Channel](texture-type-singlechannel.md) texture type documentation. |

## Texture Shape

Use the **Texture Shape** property to select and define the shape and structure of the Texture. There are four shape types:

* **2D** is the most common setting for all Textures; it defines the image file as a 2D Texture. These are used to map Textures to 3D Meshes and GUI elements, among other Project elements.
* **Cube** defines the Texture as a [cubemap](class-Cubemap-landing.md). You could use this for Skyboxes or **Reflection Probes**, for example. This type is only available with the [Default](texture-type-default.md), [Normal Map](texture-type-normal-map.md), and [Single Channel](texture-type-singlechannel.md) Texture types.
* **2D Array** defines the Texture as a [2D array texture](class-Texture2DArray.md). This is commonly used as an optimization for some rendering techniques, where many textures of the same size & format are used.
* **3D** defines the Texture as a [3D texture](class-Texture3D.md). 3D textures are used by some rendering techniques to represent volumetric data.

### 2D Array and 3D columns and rows

When you set the **Texture Shape** property to **2D Array** or **3D**, Unity displays the **Columns** and **Rows** properties. Use these to tell Unity how to divide the flipbook texture into cells.

| **Property:** | **Function:** |
| --- | --- |
| **Columns** | The number of columns that the source flipbook texture is divided into. |
| **Rows** | The number of rows that the source flipbook texture is divided into. |

## Other settings

Depending on which **Texture Type** you select, different properties can appear in the Texture Import Settings window. Some of these properties are specific to the Texture Type itself, such as **Sprite Mode** available with the [Sprite (2D and UI)](texture-type-sprite.md) type.

Use Advanced settings to make finer adjustments to the way Unity handles the Texture. The order and availability of these settings can vary depending on the **Texture Type** you choose.

For information on the properties for each texture type, refer to the documentation for that texture type:

* [Default](texture-type-default.md)
* [Normal map](texture-type-normal-map.md)
* [Editor GUI and Legacy GUI](texture-type-editor-gui-and-legacy-gui.md)
* [Sprite (2D and UI)](texture-type-sprite.md)
* [Cursor](texture-type-cursor.md)
* [Cookie](texture-type-cookie.md)
* [Lightmap](texture-type-lightmap.md)
* [Directional Lightmap](texture-type-directional-lightmap.md)
* [Shadowmask](texture-type-shadowmask.md)
* [Single Channel](texture-type-singlechannel.md)

## Additional resources

* [Platform-specific texture overrides panel reference](class-TextureImporter-type-specific.md)