# Default Gradient node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/default-gradient-node.html)

# Default Gradient node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Outputs the gradient specified for your UI elements. For example, if you set the background image of a button to use a vector graphic with a linear gradient from top red to bottom green, the Default Gradient node outputs that gradient for the button.

You can use the Default Gradient node combined with other nodes to create custom effects for the gradient render type. For example, you can multiply the output of this node with a **Color** node to filter unwanted colors from the gradient.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Gradient | Output | Gradient | The gradient specified for the UI element. |

## Additional resources

* [Default Solid node](default-solid-node.html)
* [Default Texture node](default-texture-node.html)
* [Default SDF Text node](default-sdf-text-node.html)
* [Default Bitmap Text node](default-bitmap-text-node.html)
* [Render Type Branch node](render-type-branch-node.html)