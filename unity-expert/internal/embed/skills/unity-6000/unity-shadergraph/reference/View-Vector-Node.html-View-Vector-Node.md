# View Vector Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/View-Vector-Node.html)

# View Vector Node

## Description

This node provides access to an unnormalized version of the mesh vertex or fragment's **View Direction** vector. It does not normalize any of the values it stores. For a normalized option, see [View Direction Node](View-Direction-Node.html).

Select a **Space** to modify the output value's coordinate space.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Vector 3 | None | View Vector for the Mesh Vertex/Fragment. |

## Controls

| Name | Type | Options | Description |
| --- | --- | --- | --- |
| Space | Dropdown | Object, View, World, Tangent | Selects coordinate space of **View Direction** to output. |