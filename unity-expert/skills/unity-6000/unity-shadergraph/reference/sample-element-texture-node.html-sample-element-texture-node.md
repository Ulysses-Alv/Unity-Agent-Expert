# Sample Element Texture node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/sample-element-texture-node.html)

# Sample Element Texture node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

The Sample Element Texture node samples a texture at specific UV coordinates. You can use this node to get multiple samples of the texture assigned to the element, for example to create complex visual effects or manipulate the texture.

Sampling multiple times with this node is more efficient than using several separate nodes for individual samples. This is because each node introduces overhead by traversing internal branches to select the correct texture slot. By combining multiple samples into a single node, you reduce this overhead and improve performance.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| UV 0 | Input | Vector2 | The UV coordinates to sample. |
| UV 1 | Input | Vector2 | The second set of UV coordinates to sample. |
| UV 2 | Input | Vector2 | The third set of UV coordinates to sample. |
| UV 3 | Input | Vector2 | The fourth set of UV coordinates to sample. |
| Color 0 | Output | Color | The sampled color from UV 0. |
| Color 1 | Output | Color | The sampled color from UV 1. |
| Color 2 | Output | Color | The sampled color from UV 2. |
| Color 3 | Output | Color | The sampled color from UV 3. |

## Additional resources

* [Introduction to UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/introduction-to-ui-shader-graph.html)