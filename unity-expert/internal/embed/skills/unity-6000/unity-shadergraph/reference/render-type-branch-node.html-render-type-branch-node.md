# Render Type Branch node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/render-type-branch-node.html)

# Render Type Branch node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

The Render Type Branch node routes inputs based on the current render type and outputs the appropriate results for the Fragment node. You can connect the inputs to various nodes to define how each render type is processed.

UI Shader Graph provides default nodes, such as the [Default Solid node](default-solid-node.html) or [Default Texture node](default-texture-node.html), for each render type. You can use these nodes as starting points when customizing your shaders. For best performance, if you want an input to use its default value for a render type, leave it disconnected rather than connecting it to a default node. The Render Type Branch node automatically uses the default values for that input and handles branching more efficiently when you disconnect inputs. Only connect these default nodes if you want to customize the shader’s behavior.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Solid | Input | Color | The color to use for backgrounds and borders. |
| Texture | Input | Texture | The texture to use for texture graphics. |
| SDF Text | Input | Texture | The texture to use for SDF text. |
| Bitmap Text | Input | Texture | The texture to use for bitmap text. |
| Gradient | Input | Texture | The texture to use for vector graphic gradients. |
| Color | Output | Color | The output color. |
| Alpha | Output | Float | The output alpha value. |

## Additional resources

* [Render Type node](render-type-node.html)
* [Default Solid node](default-solid-node.html)
* [Default Texture node](default-texture-node.html)
* [Default SDF Text node](default-sdf-text-node.html)
* [Default Bitmap Text node](default-bitmap-text-node.html)
* [Default Gradient node](default-gradient-node.html)