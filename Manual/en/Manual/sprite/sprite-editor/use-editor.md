# Cut out sprites from a texture

To cut out sprites from a texture, use the default ****Sprite** Editor** tab of the **Sprite Editor** window.

For each sprite, Unity creates a sprite asset () as a child of the texture in the **Project** window.

## Open the Sprite Editor

Follow these steps:

1. In the **Project** window, select the texture you want to edit.

   **Note**: You can’t edit a sprite by selecting it in the ****Scene**** view.
2. In the ****Inspector**** window, make sure the **Texture Type** is set to **Sprite (2D and UI)**.
3. Select **Sprite Editor**.

## Adjust the area of a single sprite

If the **Sprite Mode** of the texture is set to **Single**, the entire texture is a single sprite. To adjust the sprite in the **Sprite Editor** window, follow these steps:

1. Select the texture. Unity displays a rectangular selection area with handles in the corners. This is the sprite area, also known as a `SpriteRect`.
2. Drag the blue handles or edges to set the area of the sprite.
3. Configure the properties of the sprite using the **Sprite** overlay. For more information, refer to [Sprite Editor tab reference](sprite-editor-window-reference.md#sprite-overlay).
4. To save the sprite, select **Apply** in the **toolbar**.

## Cut out multiple sprites from a single texture

If the **Sprite Mode** of the texture is set to **Multiple**, the texture can contain multiple sprites. Each sprite is defined by a `SpriteRect` that you create in the **Sprite Editor** window.

Do either of the following in the **Sprite Editor** window:

* Manually select areas to use as sprites.
* Slice into sprites automatically, for example from a sprite sheet that uses a regular layout, or transparent **pixels** to separate each image.

### Manually select areas to use as sprites

Follow these steps:

1. Left-click the texture and drag to create a `SpriteRect`.
2. Repeat step 1 to create more sprites.
3. To save the sprites, select **Apply** in the toolbar.

### Slice into sprites automatically

To slice a texture into multiple sprites automatically, select the **Slice** dropdown in the toolbar.

Choose from the following slicing methods:

* **Automatic** to slice based on the transparency of the texture. Use this option if each image in the texture is separated by transparent pixels.
* **Grid By Cell Size** to slice into sprites of equal size.
* **Grid By Cell Count** to slice into a specific number of rows and columns.
* **Isometric Grid** to slice into isometric diamond-shaped sprites instead of rectangular sprites.

When you select an option other than **Automatic** or adjust the settings in the **Slice** overlay, Unity displays a preview of the sprites as red outlines. Unity might not display the outlines if you already sliced the texture.

After you select **Slice** to apply the slicing method, to adjust a generated sprite, select it then drag its blue handles or edges.

For more information, refer to [Sprite Editor tab reference](sprite-editor-window-reference.md#slice-overlay).

## Warn before applying or reverting changes

To make Unity display a warning when you select **Revert** or **Apply** in the Sprite Editor, follow these steps:

1. From the main menu, select **Edit** > **Preferences**.
2. Select **2D > Sprite Editor Window**.
3. Enable **Show Apply Confirmation** or **Show Revert Confirmation**.

## Additional resources

* [Sprite Editor tab reference](sprite-editor-window-reference.md)
* [Sprite asset reference](../../class-Sprite.md)