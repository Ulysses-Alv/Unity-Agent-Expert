# Sprite Editor tab reference for the Sprite Editor window

Explore the properties and settings you use to configure **sprites** and convert a texture into multiple sprites.

**Note:** This page is about the default **Sprite Editor** tab in the **Sprite Editor** window. For information about other tabs such as **Custom Outline** and ****Skinning** Editor**, refer to [Sprite Editor window reference](sprite-editor-window-reference-landing.md).

**Note:** If the **Sprite Mode** of the texture is set to **Polygon**, the tab is called the **Sprite Polygon Mode Editor** tab.

## Toolbar

The following properties appear in all the different tabs of the **Sprite Editor** window.

| **Property** | **Icon** | **Description** |
| --- | --- | --- |
| **Window name** | N/A | Sets the tab to display in the **Sprite Editor** window. The options are:  * **Sprite Editor**: Displays the [Sprite Editor](sprite-editor-window-reference.md) tab. * **Custom Outline**: Displays the [Custom Outline](custom-outline-editor-reference.md) tab. * **Custom Physics Shape**: Displays the [Custom Physics Shape](custom-physics-shape-editor-reference.md) tab. * **Secondary Textures**: Displays the [Secondary Textures](secondary-textures-editor-reference.md) tab. * **Skinning Editor**: Displays the **Skinning Editor** tab. For more information, refer to the [2D Animation package documentation](https://docs.unity3d.com/Packages/com.unity.2d.animation@latest/index?subfolder=/manual/index.html). |
| **Preview** | Preview icon | Previews your changes in the **Scene** view. |
| **Revert** | N/A | Discards your changes. |
| **Apply** | N/A | Saves your changes. |
| **Color** | Color icon | Toggles between displaying the color channels of the texture and its alpha channel. |
| **Zoom** | N/A | Sets the zoom level at which to display the texture. |
| **Mipmap level** | Mipmap level icons and slider | Sets the mipmap level of the texture to display. The slider ranges from the lowest resolution mipmap level on the left to the highest resolution mipmap level on the right. This property is available only if the texture has mipmap levels. |

### Sprite Editor toolbar

The following options appear only in the **Sprite Editor** tab of the Sprite Editor window.

| **Property** | **Icon** | **Description** |
| --- | --- | --- |
| **Change Shape** | N/A | Hides or reveals the **Sides** overlay in the main area. This button is available only if the **Sprite Mode** of the texture is set to **Polygon**. For more information, refer to the [Main area](#main-area) section. |
| **Slice** | N/A | Enables the **Slice** overlay that allows you to automatically slice a texture into multiple sprites. This option is available only if you set the **Sprite Mode** of the texture to **Multiple**. For more information, refer to the [Slice overlay](#slice-overlay) section. |
| **Trim** | N/A | Resizes the selected sprite so it fits tightly around the edge of the opaque part of the texture. This property is available only if the **Sprite Mode** of the texture is set to **Multiple**. |
| **Locks** | Locks icon | Locks editing the properties of all the sprites. The options are:  * **Name**: Locks or unlocks editing the names of sprites. * **Size**: Locks or unlocks editing the size of sprites. * **Position**: Locks or unlocks editing the positions of sprites. * **Border**: Locks or unlocks editing the borders of sprites. * **Create/Delete**: Locks or unlocks creating and deleting sprites.  This property is available only if you set the **Sprite Mode** of the texture to **Multiple**. |

## Slice overlay

Select the **Slice** overlay to display and set the properties for [creating multiple sprites from a single texture](use-editor.md#create-multiple-sprites).

| **Property** | **Description** |
| --- | --- |
| **Slice on Import** | Slices the texture again automatically when it changes externally and Unity reimports it. If you enable this property, the recommended best practice is to set **Method** to either **Smart** or **Safe**, to make sure Unity keeps existing sprites. |
| **Type** | Selects the slicing method. The options are:  * **Automatic**: Slices based on the transparency of the texture. Use this option if each image in the texture is separated by transparent pixels. * **Grid By Cell Size**: Slices into sprites of equal size. * **Grid By Cell Count**: Slices into a specific number of rows and columns. * **Isometric Grid**: Slices into isometric diamond-shaped sprites instead of rectangular sprites. |
| **Column & Row** | Sets the number of rows and columns in the sliced texture. This property is available only if you set **Type** to **Grid By Cell Count**. |
| ****Pixel** Size** | Sets the width and height of each sprite in pixels. This property is available only if you set **Type** to **Grid By Cell Size** or **Isometric Grid**. |
| **Offset** | Adjusts the sprite grid along the x-axis and y-axis. This property is available only if you set **Type** to an option other than **Automatic**. |
| **Padding** | Sets the amount of space to add between sprites in pixels. This property is available only if you set **Type** to **Grid By Cell Size** or **Grid By Cell Count**. |
| **Keep Empty Rects** | Keeps sprites that have no opaque pixels. Use this option to organize sprites based on their position in the texture. This property is available only if you set **Type** to an option other than **Automatic**. |
| **Is Alternate** | If you set **Type** to **Isometric Grid**, this property staggers the isometric diamonds across alternate rows. Unity assumes the first diamond in the top row starts half a pixel from the left side. |
| **Pivot** | Sets the position of the points Unity uses for transformations such as rotation. Select **Custom** to set a custom pivot point. |
| **Pivot Unit Mode** | Sets the units the **Custom Pivot** parameter uses. The options are:  * **Normalized**: The values range from 0, 0 at the bottom-left of the sprite to 1, 1 at the top-right of the sprite. * **Pixels**: The values are in pixels, and range from 0, 0 at the bottom-left of the sprite. |
| **Custom Pivot** | Sets the positions of the pivot point if you set **Pivot** to **Custom**. Right-click on the **Custom Pivot** label to copy or paste values. |
| **Method** | Keeps or deletes the existing sprites when you select **Slice**. The options are:  * **Delete Existing**: Deletes the existing sprites then adds the new sprites. * **Smart**: For each new sprite, checks if it overlaps any existing sprites. If it overlaps, Unity ignores the new sprite, but uses its position and size for the best existing overlapping sprite. * **Safe**: Keeps all the original sprites, and ignores any new sprites that overlap an existing sprite. |
| **Slice** | Slices the texture according to the settings. |

## Main area

The main area of the **Sprite Editor** window displays the texture. Each individual sprite in the texture displays as a white outline, also known as a `SpriteRect`.

Left-click a sprite to select it.

![A detail from the Sprite Editor with a rock environment texture, sliced into four sprites. One sprite is selected.](../../../uploads/Main/sprite-editor_sprite-selected.jpg)

A detail from the Sprite Editor with a rock environment texture, sliced into four sprites. One sprite is selected.

Unity displays the following:

* **A**: The sprite area as a blue outline with filled circular handles.
* **B**: The border as a green outline with filled square handles.
* **C**: The pivot point as a hollow blue circle.

## Polygon texture properties

Unity displays the following properties if the **Sprite Mode** of the texture is set to **Polygon**.

| **Property** | **Description** |
| --- | --- |
| **Sides** | The number of sides the polygon has. To change the number of sides, enter a new value and select **Change**. This property is available only if the **Sprite Mode** of the texture is set to **Polygon**. |
| **Change** | Updates the polygon to have the number of sides set in **Sides**. |

## Sprite overlay

Select a sprite to display and set its properties in the **Sprite** overlay.

| **Property** | **Description** |
| --- | --- |
| **Name** | Sets the name of the sprite. |
| **Position** | Sets the area of the texture to use as the sprite. The properties are:  * **X**: The x coordinate of the left of the sprite in pixels. * **Y**: The y coordinate of the bottom of the sprite in pixels. * **W**: The width of the sprite in pixels. * **H**: The height of the sprite in pixels. |
| **Border** | Sets the border to use for [9-slicing](../9-slice/9-slice-landing.md). Unity displays the border as a green outline in the Scene view. The properties are:  * **L**: The width of the left border in pixels. * **R**: The width of the right border in pixels. * **T**: The height of the top border in pixels. * **B**: The height of the bottom border in pixels. |
| **Pivot** | Sets the position of the point Unity uses for transformations such as rotation. Select **Custom** to set a custom pivot point. |
| **Pivot Unit Mode** | Sets the units the **Custom Pivot** parameter uses. The options are:  * **Normalized**: The values range from 0, 0 at the bottom-left of the sprite to 1, 1 at the top-right of the sprite. * **Pixels**: The values are in pixels, and range from 0, 0 at the bottom-left of the sprite. |
| **Custom Pivot** | Sets the positions of the pivot point if you set **Pivot** to **Custom**. The units depend on **Pivot Unit Mode**. Right-click on the **Custom Pivot** label to copy or paste values. |

## Additional resources

* [Create sprites from a texture](use-editor.md)