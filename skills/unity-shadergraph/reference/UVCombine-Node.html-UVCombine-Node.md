# UVCombine node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/UVCombine-Node.html)

# UVCombine node

The UVCombine node lets you select which UV channel you want to use for mapping your shader to geometry in your application. You can also choose to apply tiling and offset to your UV coordinates.

![An image showing the UVCombine node in the Shader Graph window](images/sg-uvcombine-node.png)

##### Note

This node is a Subgraph node: it represents a Subgraph instead of directly representing Unity's shader code. Double-click the node in any Shader Graph to view its Subgraph.

## Create Node menu category

The UVCombine node is under the **Utility** > **High Definition Render Pipeline** category in the Create Node menu.

## Compatibility

node is supported on the following render pipelines:

| **Built-in Render Pipeline** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** |
| --- | --- | --- |
| No | No | Yes |

For more information on the HDRP, see [Unity's HDRP package documentation](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest).

node can also be connected to any Block node in either Context. For more information on Block nodes and Contexts, refer to [Master Stack](Master-Stack.html).

## Inputs

node has the following input ports:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| **UV Channel Mask** | Vector 4 | Select which UV channel you want to use for your UV coordinates by entering a `1` in the corresponding default input on the port:  * **X**: UV channel 0 * **Y**: UV channel 1 * **Z**: UV channel 2 * **W**: UV channel 3  Set all other default inputs to `0`. You can also connect a node that outputs a Vector 4. |
| **UV Tile and Offset** | Vector 4 | Use the port's default input to specify the amount of offset or tiling that you want to apply to your shader's UV coordinates:  * Use **X** and **Y** to specify the tiling. * Use **W** and **Z** to specify the offset.  You can also connect a node that outputs a Vector 4. |

## Outputs

node has one output port:

| **Name** | **Type** | **Binding** | **Description** |
| --- | --- | --- | --- |
| **UV** | Vector 2 | UV | The final UV output, after selecting a UV channel and, if specified, any tiling or offset. |

## Example graph usage

For an example use of the UVCombine node, see either of the HDRP's Fabric shaders.

To view these Shader Graphs:

1. Create a new material and assign it the **HDRP** > **Fabric** > **Silk** or **HDRP** > **Fabric** > **CottonWool** shader, as described in the Unity User Manual section [Creating a material asset, and assigning a shader to it](https://docs.unity3d.com/Documentation/Manual/materials-introduction.html).
2. Next to the **Shader** dropdown, select **Edit**.

Your chosen Fabric's Shader Graph opens. You can view the UVCombine node, its Subgraph, and the other nodes that create HDRP's Fabric shaders.