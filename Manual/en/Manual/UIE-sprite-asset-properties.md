# Sprite asset properties reference

This page describes the **sprite** asset properties.

| Property | Description |
| --- | --- |
| **Update Sprite Asset** | Sync the sprite asset with the latest changes in the Sprite Editor window. |

## Atlas & Material

| Property | Description |
| --- | --- |
| ****Sprite Atlas****Graphics:\*\* A utility that packs several sprite textures tightly together within a single texture known as an atlas. [More info](sprite/atlas/v2/v2-landing). **2D:** A texture that is composed of several smaller textures. Also referred to as a texture atlas, image sprite, sprite sheet or packed texture. [More info](sprite/atlas/atlas-landing.md). See in [Glossary](Glossary.md#SpriteAtlas)\*\* | A reference to the sprite asset’s source texture. |
| **Default Material** | A reference to the sprite asset’s material, which it uses to render sprites. |

## Fallback Sprite Assets

When a glyph doesn’t exist in this sprite asset, TextCore searches the fallback sprite assets list for the missing glyph.

| Property | Description |
| --- | --- |
| **Fallback Sprite Asset List** | This is the local fallback list. The local fallback list has precedence over the global fallback list set in the [UITK Text Settings](UIE-text-setting-asset.md) asset. You can add or remove a sprite asset in the fallback list. You can also drag the handles on the left side of any sprite asset to reorder the list. |

## Sprite Character Table

Manage the sprites in this asset.

* Click a sprite to make it active.
* Click **Up** or **Down** to move the sprite up or down in the list.
* To move the sprite to a specific position, enter the index ID of the position in the text field and then click **Goto** to move the sprite to that position in the list.  
  **Note:** Moving a sprite updates its index ID and the index IDs of all preceding sprites accordingly.
* Click **+** to add a copy of the sprite to the list.
* Click **-** to remove the sprite from the list.

| Property | Description |
| --- | --- |
| **Index** | Unique index ID for the sprite, based on its position in the list.  You can use it to reference the sprite in the [`<sprite>` rich text tag](UIE-supported-tags.md#sprite). |
| **Unicode** | Unicode for the sprite. For more information, refer to [Assign and use Unicode for a sprite](UIE-sprite.md#assign-and-use-unicode-for-a-sprite). |
| **Name** | Unique name for the sprite. You can use it to reference the sprite in the [`<sprite>` rich text tag](UIE-supported-tags.md#sprite). |

## Sprite Glyph Table

Manage the glyphs in this asset.

| Property | Description |
| --- | --- |
| **Glyph Rect** | Define the rectangular area that the character occupies in the sprite atlas.  * **X, Y**: The coordinates of the glyph. * **W, H**: The width and height of the glyph. |
| **Glyph Metrics** | Metrics for the sprite.  * **W**: The width of the sprite. * **H**: The height of the sprite. * **BX, BY**: Stand for “Bearing X” and “Bearing Y,” which define the horizontal and vertical offsets from the glyph bounding box’s origin to the start of the drawings. * **AD**: Horizontal advance refers to the distance from the hozitontal origin of the glyph to the start of the next glyph. |

## Global Offsets & Scale

Offsets and scale settings for all sprites in the asset.

| Property | Description |
| --- | --- |
| **OX, OY** | Control the placement of the sprite, defined at its top-left corner relative to its origin on the baseline. |
| **ADV** | How far to advance along the baseline to place the next sprite. |
| **SF** | Adjust the size of the sprite. |

## Additional resources

* [Include sprites in text](UIE-sprite.md)
* [Rich text tags](UIE-rich-text-tags.md)