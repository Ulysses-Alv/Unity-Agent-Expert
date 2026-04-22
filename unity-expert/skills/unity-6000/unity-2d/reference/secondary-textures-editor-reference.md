# Secondary Textures tab reference for the Sprite Editor window

Explore the properties you use to add secondary textures to a **sprite** texture, for example **normal maps** or lighting mask maps. For more information, refer to [Add a normal map or a mask map to a sprite in URP](../../urp/SecondaryTextures.md).

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

To display a secondary texture in the **Sprite Editor** window, select it in the **Secondary Textures** overlay. To display the sprite texture again, select anywhere outside of the overlay.

## Secondary Textures overlay

| **Property** | **Description** |
| --- | --- |
| **Add** (**+**) | Adds a secondary texture. The maximum number of secondary textures is 8. |
| **Remove** (**-**) | Removes a secondary texture. Select a secondary texture to highlight it, then select **Remove**. |
| **Name** | Sets the name of the secondary texture. To add a normal map or a lighting mask map, select **\_NormalMap** or **\_MaskTex** from the dropdown. **Note:** The dropdown might include names used by Unity packages you installed, even if you have since uninstalled the packages.  For more information about adding a normal map of a lighting mask map, refer to [Add a normal map or a mask map to a sprite in URP](../../urp/SecondaryTextures.md). |
| **Texture** | Sets the texture to use as the secondary texture. Drag a texture from the **Project** window, or select the picker (**⊙**). To save the texture, select **Apply** in the **toolbar**. |

## Additional resources

* [Add normal map and mask textures to a sprite in URP](../../urp/SecondaryTextures.md)
* [Sprite Editor window reference](sprite-editor-window-reference-landing.md)