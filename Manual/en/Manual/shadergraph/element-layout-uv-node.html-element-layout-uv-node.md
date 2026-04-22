# Element Layout UV node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/element-layout-uv-node.html)

# Element Layout UV node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Outputs the geometric coordinates (UV) relative to the UI element, such as a button. It allows you to determine your position within the element itself, regardless of the texture applied. It represents the relative coordinates within the layout rect of the visual element, where `(0,0)` is the bottom-left corner and `(1,1)` is the top-right corner.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Layout UV | Output | Vector2 | The UV coordinates for the element. |

## Additional resources

* [Element Texture UV node](element-texture-uv-node.html)
* [Element Texture Size node](element-texture-size-node.html)