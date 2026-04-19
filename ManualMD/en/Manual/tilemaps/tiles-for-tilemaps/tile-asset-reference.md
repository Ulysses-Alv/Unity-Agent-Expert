# Tile asset reference

**Tiles** are **Assets** that are arranged on ****Tilemaps**** to construct a 2D environment. Each Tile references a selected ****Sprite****, which is then rendered at the Tile’s location on the [Tilemap Grid](../grid-reference.md). Refer to [Creating Tiles](./create-tile-assets.md) for more information about preparing and importing sprites for your Tiles, and the different methods for creating the Assets in the Editor.

## Asset properties

| Property | Function |
| --- | --- |
| **Preview** | Displays a visual preview of the selected Tile. |
| **Sprite** | Select the Sprite to be rendered on this Tile. Click the circle icon to the right to open the object picker window to select from the available Sprite Assets, or drag a Sprite directly onto this box. |
| **Color** | Tints the Sprite placed on this Tile with the selected Color. When set to white, the Sprite is rendered without a tint. |
| ****Collider** Type** | Defines the shape of the Collider generated for the Tile. |
| **None** | No Collider is generated. |
| **Sprite** | The Collider shape is generated based on the selected **Sprite’s** outline. |
| **Grid** | The Collider shape is based on a cell of the **Tilemap**. The shape of a cell depends on the [Cell Layout](../grid-reference.md) of the Tilemap. |

## Collider shape

The **Collider Type** set in the **Tile Asset** properties affects the generation of Collider shapes for each Tile in the Tilemap. The component’s shape generation behavior corresponds to the **Collider Types** in the following ways:

| Collider Type | Function |
| --- | --- |
| **None** | The Tilemap Collider 2D component does not generate any Collider shapes for this Tile. |
| **Sprite** | The Tilemap Collider 2D component generates a Collider shape based on the Sprite assigned to the Tile. The Collider shape is based on the Custom Physics Shapeset for the Sprite. For more information, refer to [Create collision shapes for a sprite](../../sprite/create-collision-geometry.md). |
| **Grid** | The Tilemap Collider 2D component generates a Collider shape based on the shape of the Grid cell, which is determined by the selected Cell Layout of the Grid component. |

---

* Page content and screenshots updated for [2020.1](https://docs.unity3d.com/2020.1/Documentation/Manual/30_search.html?q=newin20201) NewIn20201
* Tilemaps added in [2017.2](https://docs.unity3d.com/2017.2/Documentation/Manual/30_search.html?q=newin20172) NewIn20172