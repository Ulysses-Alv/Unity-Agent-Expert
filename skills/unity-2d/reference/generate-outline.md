# Crop a sprite

Unity renders a **sprite** on a **mesh**. To crop the mesh to a different shape, use the **Custom Outline** tab of the **Sprite Editor** window.

You can use the Custom Outline tab to do the following:

* Remove transparent **pixels** from the mesh, to improve performance by reducing the number of unneeded pixels Unity renders.
* Crop a sprite to a custom shape.

## Open the Custom Outline tab of the Sprite Editor

Follow these steps:

1. In the **Hierarchy** window, select the sprite **GameObject**.
2. In the ****Inspector**** window, in the ****Sprite Renderer**** component, select **Open Sprite Editor**.
3. Select the **Custom Outline** tab in the top-left dropdown. Unity displays the **Outline Tools** overlay.

## Remove transparent pixels from the mesh

Follow these steps in the **Custom Outline** tab:

1. Select **Generate** to automatically generate a mesh that follows the opaque parts of the sprite. Unity displays the outline of the mesh in white. Each point is a vertex of the mesh.
2. Edit the outline if needed. For more information, refer to the [Edit an outline](#edit-an-outline) section.
3. To save the shape, select **Apply** in the **toolbar**. Unity now renders only the area within the outline.

To change how closely the outline follows the opaque parts of the sprite. adjust the **Outline Detail** and **Alpha Tolerance** properties, then regenerate the mesh.

For more information, refer to [Custom Outline tab reference for the Sprite Editor window](custom-outline-editor-reference.md).

![Left: An automatically generated mesh outline with a low Outline Detail value. Right: An automatically generated mesh outline with a higher Outline Detail value.](../../../uploads/Main/2DCustomOutline_7.png)

Left: An automatically generated mesh outline with a low **Outline Detail** value. Right: An automatically generated mesh outline with a higher **Outline Detail** value.

## Crop a sprite to a custom shape

Do the following in the **Custom Outline** tab:

1. To create an outline, click and drag a rectangle. You can also follow the steps in the previous section to generate an outline.

   You can add multiple outlines to the same sprite.
2. Edit the outline. For more information, refer to the [Edit an outline](#edit-an-outline) section.
3. To save the shape, select **Apply** in the toolbar. Unity now renders only the area within the outline.

## Edit an outline

To edit an outline, do the following:

* To move a point, select and drag it.
* To add a point, click on an edge.
* To delete a point, select it and press **Delete**.
* To move an edge, hold **Ctrl** and select and drag the edge.

To save the new outline, select **Apply** in the toolbar.

## Additional resources

* [Custom Outline tab reference for the Sprite Editor window](custom-outline-editor-reference.md)
* [Create sprites from a texture](../sprite-editor/use-editor.md)