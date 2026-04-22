# Tile Set properties reference

| **Property** | **Description** |
| --- | --- |
| **Open Tile Palette Window** | Opens the [Tile Palette Window](tile-palette-editor-reference.md) to start painting with the Tile Set. |
| **Texture Sources** | The list of **sprite** textures and Tile Templates that create Tiles for the Tile Set. For more information, refer to [Texture Sources](#texture-sources). |
| **Palette Grid** | Settings for the Tile Palette the Tile Set creates. For more information, refer to [Palette Grid](#palette-grid). |
| **Create Atlas** | If enabled, Unity creates a Sprite Atlas, and atlases the textures in Tile Set. |
| **Sprite Atlas Settings** | Settings for the [Sprite Atlas](../../sprite/atlas/sprite-atlas-reference.md) Unity creates. For more information, refer to [Sprite Atlas reference](../../sprite/atlas/sprite-atlas-reference.md). This section is available only if you enable **Create Atlas**. |

## Texture Sources

Each texture you add has the following properties in the **Element** dropdown.

| **Property** | **Description** |
| --- | --- |
| **Texture** | Selects a texture with sprites you want to use to create Tiles for the Tile Set. If the texture changes and Unity reimports it, the Tile Set also reimports and updates. |
| **Tile Template** | Selects the Tile Template asset Unity uses to create Tiles for the Tile Set. If the Tile Template changes and Unity reimports it, the Tile Set also reimports and updates. If no Tile Template asset is selected, Unity creates a Tile from each sprite in the **Texture**. |

## Palette Grid

The **Palette Grid** section has the following properties.

| **Property** | **Description** |
| --- | --- |
| **Cell Layout** | Sets the grid layout for the Tile Palette that Unity creates. The options are:  * **Rectangle** * **Hexagon** * **Isometric** * **Isometric Z as Y** |
| **Hexagon Layout** | Sets the layout of the Tile Palette if you set **Cell Layout** to **Hexagon**. The options are:  * **Point Top** * **Flat Top**  For more information, refer to [Hexagonal Tilemaps](../../tilemaps/work-with-tilemaps/hexagonal-tilemaps.md). |
| **Cell Sizing** | Controls the size of the cell you paint a tile on. The options are:  * **Automatic**: Calculates the cell size of the Tile Palette from the Tiles placed in the Tile Palette. * **Manual**: Uses **Cell Size** for the cell property. |
| **Cell Size** | Sets the size of the cell you paint a tile on. This option is available only if you set **Cell Sizing** to **Manual**. |
| **Transparency Sort Mode** | Determines the transparency sort mode of renderers in the Tile Palette. The options are:  * **Default** * **Perspective** * **Orthographic** * **Custom Axis** |
| **Transparency Sort Axis** | Sets the x-axis, y-axis, and z-axis of the sorting axis. This option is available only if you set **Transparency Sort Mode** to **Custom Axis**. |