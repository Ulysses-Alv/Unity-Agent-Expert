# Custom Render Texture Self Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Custom-Render-Texture-Self-Node.html)

# Custom Render Texture Self Node

## Description

Provides the texture that contains the result of the previous update of the **Custom Render Texture**. Use the output that corresponds to the type of **Custom Render Texture** used.

For more information on Custom Render Textures, refer to the [Unity Manual](https://docs.unity3d.com/Manual/class-CustomRenderTexture.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Self Texture 2D | Output | Texture2D | None | 2D Texture object that contains the update result of the previous **Custom Render Texture**. |
| Self Texture Cube | Output | Cubemap | None | Cubemap object that contains the update result of the previous **Custom Render Texture**. |
| Self Texture 3D | Output | Texture3D | None | 3D Texture object that contains the update result of the previous **Custom Render Texture**. |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
UnityBuildTexture2DStructNoScale(_SelfTexture2D)
```