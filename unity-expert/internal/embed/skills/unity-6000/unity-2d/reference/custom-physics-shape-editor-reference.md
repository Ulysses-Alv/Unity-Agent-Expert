# Custom Physics Shape tab reference for the Sprite Editor window

Explore the properties you use to configure the geometry that Unity uses to detect if a **sprite** collides with another sprite. For more information, refer to [Create collision shapes for a sprite](create-collision-geometry).

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

## Main area

The main area of the **Sprite Editor** window displays the texture. Each individual sprite in the texture displays as a white outline, also known as a `SpriteRect`.

Left-click a sprite to select it.

![A detail from the Sprite Editor with a rock environment texture, sliced into four sprites. One sprite is selected.](../../../uploads/Main/sprite-editor_sprite-selected.jpg)

A detail from the Sprite Editor with a rock environment texture, sliced into four sprites. One sprite is selected.

Unity displays the following:

* **A**: The sprite area as a blue outline with filled circular handles.
* **B**: The border as a green outline with filled square handles.
* **C**: The pivot point as a hollow blue circle.

Select a sprite to display the **collision** geometry in white. Each square point is a vertex of a collision shape.

To edit a collision shape, do the following:

* To move a point, select and drag it.
* To add a point, click on an edge.
* To delete a point, select it and press **Delete**.
* To move an edge, hold **Ctrl** and select and drag the edge.

For more information, refer to [Create collision shapes for a sprite](create-collision-geometry).

## Outline Tool

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Outline Detail** | N/A | Sets how detailed the outline of the collision geometry is. A higher value creates a more detailed outline that follows the shape of the sprite more closely, but also increases the number of vertices in the geometry, which can affect performance. |
| **Alpha Tolerance** | N/A | Sets the minimum alpha value for Unity to consider a **pixel** of the sprite as opaque. For example, if you set **Alpha Tolerance** to 128, Unity considers pixels with an alpha value of below 128 transparent. |
| **Snap** | N/A | Snaps the vertices to the nearest pixel. |
| **Generate** | N/A | Generates the collision geometry. To save the geometry, select **Apply** in the **toolbar**. |
| **Generate All** | N/A | Generates collision geometry for all the sprites in the texture that don’t already have collision geometry. This button is available only if you don’t select a sprite. |
| **Force Generate All** | N/A | Generates collision geometry for all the sprites in the texture, even if they already have collision geometry. This button is available only if you don’t select a sprite, and you select the checkbox next to **Generate All**. |
| **Copy** | N/A | Stores the current geometry. Unity removes the copy from the clipboard if you close the Sprite Editor window or select a different tab. |
| **Paste** | N/A | Retrieves the geometry you stored. To paste to a different sprite, select the destination sprite in the **Project** window first. If the position of a vertex is outside the sprite frame, Unity clamps the vertex position to inside the frame. This button is only available if you select **Copy** first. |
| **Paste All** | N/A | Retrieves the geometry you stored and applies it to all the sprites in the texture, regardless of which sprites are selected. If the position of a vertex is outside the area of a sprite, Unity clamps the vertex position to inside the area. This button is only available if you select **Copy** first. |
| **From Custom Outline** | N/A | Copies the shape from the **Custom Outline** tab. |
| N/A | **Paste** | Copies the shape from the **Custom Outline** tab to this sprite. |
| N/A | **Paste All** | Copies the shape from the **Custom Outline** tab to all the sprites in the texture, regardless of which sprites are selected. |

## Additional resources

* [Create collision shapes for a sprite](../create-collision-geometry.md)
* [2D Physics](../../2d-physics/2d-physics.md)