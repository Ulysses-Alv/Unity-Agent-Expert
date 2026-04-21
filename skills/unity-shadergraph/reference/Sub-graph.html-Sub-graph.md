# Introduction to Sub Graphs | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Sub-graph.html)

# Introduction to Sub Graphs

A Sub Graph is a type of shader graph that you include in other shader graphs. Use a Sub Graph to perform the same operations multiple times within a single shader graph, or across multiple shader graphs. It's similar to a reusable function in programming.

Create and edit Sub Graphs in the same way as regular shader graphs. Sub Graphs differ from regular shader graphs in the following ways:

* A Sub Graph is saved in its own asset file.
* A Sub Graph doesn't have a Master Stack. Instead it has an Output Node.
* When you add a Sub Graph to a shader graph, Unity creates a [Sub Graph Node](Sub-graph-Node.html). The inputs are the blackboard properties of the Sub Graph, and the outputs are from the Output node.
* To change the behavior of a Sub Graph, add a Dropdown node to the Sub Graph. For more information, refer to [Change the behavior of a Sub Graph with a dropdown](Change-Behaviour-Sub-Graph-Dropdown.html).

For more information, refer to [Create a Sub Graph](Create-Sub-Graph.html).

## Examples

The following Sub Graph uses a Multiply node to brighten an input color.

![A Color property connected to the A input of a Multiply node. The B input of the Multiply node is set to (5, 5, 5, 1), which brightens the color. The output is connected to an Output node, with a Sub Graph output called Brighter_Color.](images/sub-graph-example.png)

The following shader graph uses the Sub Graph to brighten a custom color.

![A dark Color property connected to the input of the Brighten Color Sub Graph node. The Sub Graph node displays a preview of the multiplied brighter color. The output connects to the Base Color of the Fragment context.](images/sub-graph-example-parent-graph.png)

Refer to the following for other example Sub Graphs:

* [Shader Graph samples](ShaderGraph-Samples.html)
* Built-in nodes that are Sub Graphs, for example the [ThreadMapDetail node](ThreadMapDetail-Node.html) or the [SpeedTree](SpeedTree8-SubGraphAssets.html) nodes.
* The example shader graph on the [Sample Texture 2D node](Sample-Texture-2D-Node.html) page.

## Additional resources

* [Branch On Input Connection node](Branch-On-Input-Connection-Node.html)
* [Custom Function Node](Custom-Function-Node.html)