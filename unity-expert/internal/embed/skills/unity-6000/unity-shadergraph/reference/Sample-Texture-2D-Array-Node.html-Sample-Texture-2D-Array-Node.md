# Sample Texture 2D Array node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Sample-Texture-2D-Array-Node.html)

# Sample Texture 2D Array node

The Sample Texture 2D Array node samples a **Texture 2D Array** asset and returns a **Vector 4** color value. You can specify the **UV** coordinates for a texture sample and use a [Sampler State node](Sampler-State-Node.html) to define a specific Sampler State. The node's **Index** input port specifies which index of a Texture 2D Array to sample.

For more information about Texture 2D Arrays, see [Texture Arrays](https://docs.unity3d.com/Manual/class-Texture2DArray.html) in the Unity User manual.

##### Note

If you experience texture sampling errors for this node in a graph with Custom Function Nodes or Sub Graphs, upgrade to Shader Graph version 10.3 or later.

![An image that displays the Graph window with a Sample Texture 2D Array node.](images/sg-sample-texture-2d-array-node.png)

## Create Node menu category

The Sample Texture 2D Array node is under the **Input** > **Texture** category in the Create Node menu.

## Limitations

With its default settings, this node can only connect to a Block node in the **Fragment** Context of a Shader Graph. To sample a Texture for the **Vertex** Context of a Shader Graph, set the [Mip Sampling Mode](#additional-node-settings) to **LOD**.

## Inputs

The Sample Texture 3D node has the following input ports:

| **Name** | **Type** | **Binding** | **Description** |
| --- | --- | --- | --- |
| **Texture Array** | Texture 2D Array | None | The Texture 2D Array asset to sample. |
| **Index** | Float | None | The index of the specific Texture in the Texture array to sample. The index value is the Texture's location in the Texture array. The index values in an array always start at 0. An array with four textures would have locations 0, 1, 2, and 3. |
| **UV** | Vector 2 | None | UV coordinates to use to sample the Texture. |
| **Sampler** | Sampler State | Default Sampler State | The Sampler State and settings to use to sample the texture. |
| **LOD** | Float | LOD | **NOTE**: The **LOD** Input port only displays if **Mip Sampling Mode** is **LOD**. For more information, refer to [Additional node settings](#additional-node-settings). The specific mip to use when sampling the Texture. |
| **UV** | Vector 2 | UV | The UV coordinates to use to sample the texture. |
| **Sampler** | Sampler State | Default Sampler State | The Sampler State and settings to use to sample the texture. |
| **LOD** | Float | LOD | The specific mip to use when sampling the Texture. **NOTE** The **LOD** Input port only displays if **Mip Sampling Mode** is **LOD**. For more information, refer to [Additional node settings](#additional-node-settings). |
| **Bias** | Float | Bias | **NOTE**: The **Bias** Input port only displays if **Mip Sampling Mode** is **Bias**. For more information, refer to [Additional node settings](#additional-node-settings). If **Use Global Mip Bias** is enabled, Unity adds this Bias amount to the Global Mip Bias for a texture's mip calculation. If **Global Mip Bias** is disabled, Unity uses this Bias amount instead of the Global Mip Bias. |
| **DDX** | Float | DDY | **NOTE**: The **DDX** Input port only displays if **Mip Sampling Mode** is **Gradient**. For more information, refer to [Additional node settings](#additional-node-settings). The specific DDX value to use to calculate the texture's mip when sampling. For more information on DDX values for mipmaps, refer to [Mipmaps introduction](https://docs.unity3d.com/Documentation/Manual/texture-mipmaps-introduction.html) in the Unity User Manual. |
| **DDY** | Float | DDY | **NOTE** The **DDY** Input port only displays if **Mip Sampling Mode** is **Gradient**. For more information, refer to [Additional node settings](#additional-node-settings). The specific DDY value to use to calculate the texture's mip when sampling. For more information on DDY values for mipmaps, refer to [Mipmaps introduction](https://docs.unity3d.com/Documentation/Manual/texture-mipmaps-introduction.html)> in the Unity User Manual. |

## Additional node settings

The Sample Texture 3D node has some additional settings that you can access from the Graph Inspector:

| **Option** | **Description** |
| --- | --- |
| **Use Global Mip Bias** | Enable Use Global Mip Bias to use the render pipeline's Global Mip Bias. This bias adjusts the percentage of texture information taken from a specific mip when sampling. For more information on mip bias, refer to Mipmaps introduction in the Unity User Manual. The options are:  * **Enabled**: Shader Graph uses the render pipeline's Global Mip Bias to adjust the texture information taken when sampling. * **Disabled**: Shader Graph doesn't use the render pipeline's Global Mip Bias to adjust texture information when sampling. |
| **Mip Sampling Mode** | Choose the sampling mode to use to calculate the mip level of the texture. The options are:  * **Standard**: The render pipeline calculates and automatically selects the mip for the texture. * **LOD**: The render pipeline lets you set an explicit mip for the texture on the node. The texture will always use this mip, regardless of the DDX or DDY calculations between pixels. Set the Mip Sampling Mode to LOD to connect the node to a Block node in the Vertex Context. For more information on Block nodes and Contexts, refer to Master Stack. * **Gradient**: The render pipeline lets you set the DDX and DDY values to use for its mip calculation, instead of using the values calculated from the texture's UV coordinates. For more information on DDX and DDY values, refer to Mipmaps introduction in the User Manual. * **Bias**: The render pipeline lets you set a bias to adjust the calculated mip for a texture up or down. Negative values bias the mip to a higher resolution. Positive values bias the mip to a lower resolution. The render pipeline can add this value to the value of the Global Mip Bias, or use this value instead of its Global Mip Bias. For more information on mip bias, refer to Mipmaps introduction in the User Manual. |

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

In the following example, the Sample Texture 2D Array node samples a Texture array that has 4 different cloth normal maps. Change the number given to the **Index** port as an input, and the Sample Texture 2D Array node can sample a specific normal map from the array. The **Index** value changes the output the node sends to the Normal Unpack node, and the Normal (Tangent Space) Block node in the Master Stack.

![An image of the Graph window, that displays a Sample Texture 2D Array node. The node has a Sampler State node attached as an input and sends its RGBA output to the Normal Unpack node. The Normal Unpack node's Out output port connects to the Normal (Tangent Space) Block node in the Master Stack. The Index is set to 2, which makes the sphere in the Main Preview window render with a leather-like Texture.](images/sg-sample-texture-2d-array-node-example.png)

![An image of the Graph window, that displays a Sample Texture 2D Array node. The node has a Sampler State node attached as an input and sends its RGBA output to the Normal Unpack node. The Normal Unpack node's Out output port connects to the Normal (Tangent Space) Block node in the Master Stack. The Index is set to 0, which makes the sphere in the Main Preview window render with a ridged fabric Texture.](images/sg-sample-texture-2d-array-node-example-2.png)

## Generated code example

The following code represents this node in Unity's shader code:

```
float4 _SampleTexture2DArray_RGBA = SAMPLE_TEXTURE2D_ARRAY(Texture, Sampler, UV, Index);
float _SampleTexture2DArray_R = _SampleTexture2DArray_RGBA.r;
float _SampleTexture2DArray_G = _SampleTexture2DArray_RGBA.g;
float _SampleTexture2DArray_B = _SampleTexture2DArray_RGBA.b;
float _SampleTexture2DArray_A = _SampleTexture2DArray_RGBA.a;
```

## Related nodes

The following nodes are related or similar to the Sample Texture 3D node:

* [Sample Texture 2D node](Sample-Texture-2D-Node.html)
* [Sample Texture 3D node](Sample-Texture-3D-Node.html)
* [Sampler State node](Sampler-State-Node.html)