# Element Texture UV node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/element-texture-uv-node.html)

# Element Texture UV node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Outputs the texture coordinates (UV) of the texture mapped onto a UI element.

This node might output different coordinates than the [**Element Layout UV**](element-layout-uv-node.html) node, which provides coordinates within the element's bounding rectangle. The coordinates are more likely to be different if you tile, offset, or transform the texture.

If the texture is part of an atlas, its UV coordinates only map to a specific region within the atlas. If you repeat UV coordinates or sample outside them, the data comes from other textures in the atlas. Use texture coordinates (UV) when you need precise control over how a texture appears on a UI element, and be mindful of atlas constraints.

The texture UV can also originate from a custom mesh when you call [`MeshGenerationContext.DrawMesh`](https://docs.unity3d.com/6000.6/Documentation/ScriptReference/UIElements.MeshGenerationContext.DrawMesh.html). In such cases, the UV values might vary depending on the mesh data.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Texture UV | Output | Vector2 | The UV coordinates for the texture. |

## Additional resources

* [Element Layout UV node](element-layout-uv-node.html)
* [Element Texture Size node](element-texture-size-node.html)