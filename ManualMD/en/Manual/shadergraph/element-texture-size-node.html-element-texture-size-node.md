# Element Texture Size node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/element-texture-size-node.html)

# Element Texture Size node

##### Note

Use this node only in shaders with **Material** set to **UI** for UI Toolkit. Using it elsewhere can lead to unexpected behavior or errors. For information on how to create shaders for UI Toolkit, refer to [UI Shader Graph](https://docs.unity3d.com/6000.6/Documentation/Manual/ui-systems/ui-shader-graph.html).

Outputs the texture size that's assigned to the UI element. The output is undefined if the render type is solid.

This node is similar to the [**Texture Size**](Texture-Size-Node.html) node, but it specifically retrieves the size of the texture assigned to the UI element. This is important because UI elements can have textures assigned through styles or images, and these textures might differ from the material's main texture. For example, a button might have a background image set through its style, which is different from the material's texture.

Follow these guidelines to decide which node to use:

* Use the **Texture Size** node if you want to apply a texture-based effect that's not element-specific (for example, a soft mask). In this case, set the texture explicitly on the material and use the Texture Size node to get its size.
* Use the **Element Texture Size** node if you need the size of the texture assigned to a specific VisualElement, such as a background image or an image element. This includes textures set via styles or textures used in custom meshes drawn with `MeshGenerationContext.DrawMesh(texture)`.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| Width | Output | Float | The width of the texture. |
| Height | Output | Float | The height of the texture. |
| Texel Width | Output | Float | The texel width of the texture. |
| Texel Height | Output | Float | The texel height of the texture. |

## Additional resources

* [Element Texture UV node](element-texture-uv-node.html)
* [Element Layout UV node](element-layout-uv-node.html)
* [Texture Size node](Texture-Size-Node.html)