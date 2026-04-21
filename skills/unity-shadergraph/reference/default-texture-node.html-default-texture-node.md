# Default Texture node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/default-texture-node.html)

# Default Texture node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Provides the texture assigned to the UI element.

You can use this node to access the texture assigned to a UI element, such as a Texture 2D background image. The node includes UV and tint inputs that allow you to modify how the texture is applied. For example, you can connect a **Tiling and Offset** node to the UV input to create a repeating effect for the background image, or connect a **Color** node to the tint input to adjust the tint color of the background image.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| UV | Input | Vector2 | The UV coordinates for the texture. |
| Tint | Input | Color | The tint color to apply to the texture. |
| Texture | Output | Texture | The resulting texture. |

## Additional resources

* [Default Solid node](default-solid-node.html)
* [Default Gradient node](default-gradient-node.html)
* [Default SDF Text node](default-sdf-text-node.html)
* [Default Bitmap Text node](default-bitmap-text-node.html)