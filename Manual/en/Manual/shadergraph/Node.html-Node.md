# Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Node.html)

# Node

## Description

A **Node** defines an input, output or operation on the Shader Graph, depending on its available [Ports](Port.html). A **Node** may have any number of input and/or output ports. You create a Shader Graph by connecting these ports with [Edges](Edge.html). A **Node** might also have any number of **Controls**, these are controls on the **Node** that do not have ports.

You can collapse a **Node** by clicking the **Collapse** button in the top-right corner of the **Node**. This will hide all unconnected ports.

For components of a **Node** see:

* [Port](Port.html)
* [Edge](Edge.html)

There are many available **Nodes** in Shader Graph. For a full list of all available **Nodes** see the [Node Library](Node-Library.html).

## Preview

Many nodes include a preview. This preview displays the main output value at that stage in the graph. Hide this preview with the Collapse control that displays when you hover over the node. You can also collapse and expand node previews via the Context Menu in the [Shader Graph Window](Shader-Graph-Window.html). To configure the appearance of node previews, see [Preview Mode Control](Preview-Mode-Control.html).

## Context Menu

Right clicking on a **Node** will open a context menu. This menu contains many operations that can be performed on the **Node**. Note that when multiple nodes are selected, these operations will be applied to the entire selection.

| Item | Description |
| --- | --- |
| Copy Shader | Copies the generated HLSL code at this stage in the graph to the clipboard |
| Disconnect All | Removes all edges from all ports on the **Node(s)** |
| Cut | Cuts selected **Node(s)** to the clipboard |
| Copy | Copies selected **Nodes(s)** to the clipboard |
| Paste | Pastes **Node(s)** in the clipboard |
| Delete | Deletes selected **Node(s)** |
| Duplicate | Duplicates selected **Node(s)** |
| Convert To Sub-graph | Creates a new [Sub-graph](Sub-graph.html) with the selected **Node(s)** included |
| Convert To Inline Node | Converts a [Property Node](Property-Types.html) into a regular node of the appropriate [Data Type](Data-Types.html) |
| Convert To Property | Converts a **Node** into a new **Property** on the [Blackboard](Blackboard.html) of the appropriate [Property Type](Property-Types.html) |
| Open Documentation | Opens a new web browser to the selected **Nodes** documentation page in the [Node Library](Node-Library.html) |

## Color Mode

**Nodes** interact with the Shader Graph Window's Color Modes. Colors are displayed on nodes underneath the text on the node title bar. See [Color Modes](Color-Modes.html) for more information on available colors for nodes.

Unity applies each component of T as a weight factor to each component to A and B. If T has fewer components than A and B, Unity casts T to the required number of components. Unity copies the values of the original components of T to the added components.