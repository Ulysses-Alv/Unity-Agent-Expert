# Default Bitmap Text node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/default-bitmap-text-node.html)

# Default Bitmap Text node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Outputs the text color set for bitmap text rendering and includes a tint input you can use to modify the color of the text. For example, if you connect a **Color** node to the tint input and set it to red, and connect the output to bitmap text render type, the text color of your bitmap text becomes red.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Tint | Input | Color | The color tint to apply to the text. |
| Bitmap text | Output | Texture | The rendered bitmap of the text. |

## Additional resources

* [Default Solid node](default-solid-node.html)
* [Default Gradient node](default-gradient-node.html)
* [Default Texture node](default-texture-node.html)
* [Default SDF Text node](default-sdf-text-node.html)
* [Render Type Branch node](render-type-branch-node.html)