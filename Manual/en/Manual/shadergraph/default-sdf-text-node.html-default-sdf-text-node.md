# Default SDF Text node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/default-sdf-text-node.html)

# Default SDF Text node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Outputs the text color for SDF text rendering and includes a tint input you can use to modify the color of the text. For example, if you connect a **Color** node to the tint input and set it to red, and connect the output to SDF text render type, the text color of your SDF text becomes red.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Tint | Input | Color | The tint color to apply to the text. |
| SDF Text | Output | Texture | The rendered SDF text as a texture. |

## Additional resources

* [Default Solid node](default-solid-node.html)
* [Default Gradient node](default-gradient-node.html)
* [Default Texture node](default-texture-node.html)
* [Default Bitmap Text node](default-bitmap-text-node.html)
* [Render Type Branch node](render-type-branch-node.html)