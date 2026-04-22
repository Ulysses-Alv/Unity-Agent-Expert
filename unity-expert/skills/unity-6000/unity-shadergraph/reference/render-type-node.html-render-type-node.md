# Render Type node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/render-type-node.html)

# Render Type node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

The Render Type node outputs the current render type the shader is processing. You can use this node to create custom logic based on the render type. For example, you can use the output of the Render Type node to control logic that changes behavior depending on the render type.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Solid | Output | Boolean | True when the shader is processing the Solid render type. |
| Texture | Output | Boolean | True when the shader is processing the Texture render type. |
| SDF Text | Output | Boolean | True when the shader is processing the SDF Text render type. |
| Bitmap Text | Output | Boolean | True when the shader is processing the Bitmap Text render type. |
| Gradient | Output | Boolean | True when the shader is processing the Gradient render type. |

## Additional resources

* [Render Type Branch node](render-type-branch-node.html)