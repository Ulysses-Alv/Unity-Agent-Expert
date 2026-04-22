# Sample Texture 3D node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Sample-Texture-3D-Node.html)

# Sample Texture 3D node

The Sample Texture 3D node samples a **Texture 3D** asset and returns a **Vector 4** color value. You can specify the **UV** coordinates for a texture sample and use a [Sampler State node](Sampler-State-Node.html) to define a specific Sampler State.

For more information about Texture 3D assets, refer to [3D textures](https://docs.unity3d.com/Manual/class-Texture3D.html) in the Unity User manual.

##### Note

If you experience texture sampling errors for this node in a graph with Custom Function Nodes or Sub Graphs, upgrade to Shader Graph version 10.3 or later.

![An image that displays the Graph window with a Sample Texture 3D node.](images/sg-sample-texture-3d-node.png)

## Create Node menu category

The Sample Texture 2D node is under the **Input** > **Texture** category in the Create Node menu.

## Limitations

With its default settings, this node can only connect to a Block node in the **Fragment** Context of a Shader Graph. To sample a Texture for the **Vertex** Context of a Shader Graph, set the [Mip Sampling Mode](#additional-node-settings) to **LOD**.

## Inputs

The Sample Texture 3D node has the following input ports:

| **Name** | **Type** | **Binding** | **Description** |
| --- | --- | --- | --- |
| **Texture** | Texture 3D | None | The 3D Texture asset to sample. |
| **UV** | Vector 3 | None | The 3D UV coordinates to use to sample the Texture. |
| **Sampler** | Sampler State | Default Sampler State | The Sampler State and settings to use to sample the texture. |
| **LOD** | Float | LOD | The specific mip to use when sampling the Texture. **NOTE** The **LOD** Input port only displays if **Mip Sampling Mode** is **LOD**. For more information, refer to [Additional node settings](#additional-node-settings). |

## Additional node settings

The Sample Texture 3D node has some additional settings that you can access from the Graph Inspector:

| Name | Type | Option | Description |
| --- | --- | --- | --- |
| Mip Sampling Mode | Dropdown | N/A | Choose the sampling mode that the Sample Texture 3D node uses to calculate the mip level of the texture. |
| N/A | N/A | Standard | The mip is calculated and selected automatically for the texture. |
| N/A | N/A | LOD | Set an explicit mip for the texture. The texture will always use this mip, regardless of the DDX or DDY calculations between pixels. If the Mip Sampling Mode is set to LOD, you can connect the node to a Block node in the Vertex Context. For more information on Block nodes and Contexts, refer to [Master Stack](Master-Stack.html). |

## Outputs

The Sample Texture 3D node has the following output ports:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| **RGBA** | Vector 4 | The full RGBA Vector 4 color value of the texture sample. |
| **R** | Float | The Red channel or X component of the texture sample. |
| **G** | Float | The Green channel or Y component of the texture sample. |
| **B** | Float | The Blue channel or Z component of the texture sample. |
| **A** | Float | The Alpha channel or W component of the texture sample. |

## Example graph usage

In the following example, the Sample Texture 3D node samples a 3D fractal noise Texture asset. It takes its input UV coordinates from a Position node, set to **Object** Space.

The Sample Texture 3D node needs a Vector 3 for its UV coordinate input, rather than a Vector 2, because the Texture asset exists as a volume in imaginary 3D space. The node uses the default Sampler State because there is no Sampler State node connected.

This specific Texture 3D asset stores its Texture data in the Alpha channel, so the Sample Texture 3D node uses its **A** output port as an input for the Base Color Block node in the **Fragment** Context of the Master Stack:

![An image of the Graph window, that displays a Position node connected to the UV input port on a Sample Texture 3D node. The Sample Texture 3D node's A output port connects to the Base Color Block node in the Fragment Context of the Master Stack.](images/sg-sample-texture-3d-node-example.png)

## Generated code example

The following code represents this node in Unity's shader code:

```
float4 _SampleTexture3D_Out = SAMPLE_TEXTURE3D(Texture, Sampler, UV);
```

## Related nodes

The following nodes are related or similar to the Sample Texture 3D node:

* [Sample Texture 2D Array node](Sample-Texture-2D-Array-Node.html)
* [Sample Texture 2D node](Sample-Texture-2D-Node.html)
* [Sampler State node](Sampler-State-Node.html)