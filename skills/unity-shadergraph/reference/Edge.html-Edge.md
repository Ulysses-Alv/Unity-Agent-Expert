# Edge | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Edge.html)

# Edge

## Description

An **Edge** defines a connection between two [Ports](Port.html). **Edges** define how data flows through the [Shader Graph](index.html) node network. They can only be connected from an input [Port](Port.html) to an output [Port](Port.html).

Each **Edge** has a [Data Type](Data-Types.html) which defines what [Ports](Port.html) it can be connected to. Each [Data Type](Data-Types.html) has an associated color for identifying its type.

You can create a new **Edge** by clicking and dragging from a [Port](Port.html) with the left mouse button. Edges can be deleted with Delete (Windows), Command + Backspace (OSX) or from the context menu by right clicking on the edge.

You can open a contextual [Create Node Menu](Create-Node-Menu.html) by dragging an **Edge** from a [Port](Port.html) with the left mouse button and releasing it in an empty area of the workspace.