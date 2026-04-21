# Default Solid node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/default-solid-node.html)

# Default Solid node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Outputs the solid color specified for your UI elements, such as the background color of a button. For example, if you set the background color of a button to yellow, the **Default Solid** node outputs yellow for that button.

You can use this node combined with other nodes to create custom effects for the Solid color render type. For example, you can multiply the output of this node with a **Color** node to filter unwanted colors.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Solid | Output | Color | The solid color specified for the UI element. |

## Additional resources

* [Default Gradient node](default-gradient-node.html)
* [Default Texture node](default-texture-node.html)
* [Default SDF Text node](default-sdf-text-node.html)
* [Default Bitmap Text node](default-bitmap-text-node.html)
* [Render Type Branch node](render-type-branch-node.html)