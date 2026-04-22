# Branch On Input Connection node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Branch-On-Input-Connection-Node.html)

# Branch On Input Connection node

The Branch On Input Connection node allows you to change the behavior of a [Sub Graph](Sub-graphs.html) based on the connected state of an input property in the parent shader graph. Use the Branch On Input Connection node to create a default input for a port.

For more information, refer to [Set default inputs for a Sub Graph](Sub-Graph-Default-Property-Values.html).

The node generates branching HLSL source code, but during compilation Unity optimizes the branch out of your shader.

You can only use a Branch on Input Connection node in a [Sub Graph](Sub-graphs.html).

## Inputs

The Branch On Input Connection node has the following input ports:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| **Input** | Property | The property that determines the branching logic in the node, based on its connection state in the parent Shader Graph. |
| **Connected** | Dynamic Vector | The value to send to the **Out** port when **Input** is connected in the parent Shader Graph. |
| **NotConnected** | Dynamic Vector | The value to send to the **Out** port when **Input** isn't connected in the parent Shader Graph. |

## Outputs

The Branch On Input Connection node has one output port:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| **Out** | Dynamic Vector | Outputs the value of either **Connected** or **NotConnected**, based on the **Input** property's connection state in the parent Shader Graph. |

## Related nodes

The following nodes are related or similar to the Branch On Input Connection node:

* [Branch node](Branch-Node.html)
* [Subgraph node](Sub-graph-Node.html)