# Bitangent Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Bitangent-Vector-Node.html)

# Bitangent Node

## Description

Provides access to the mesh vertex or fragment's **Bitangent Vector**, depending on the effective [Shader Stage](Shader-Stage.html) of the graph section the [Node](Node.html) is part of.

The bitangent vector is derived from the normal and tangent vectors and is orthogonal to both. The three vectors provide a reference frame to perform complex light calculations, for example.

You can select the coordinate space of the output with the **Space** dropdown parameter.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Vector 3 | None | **Bitangent Vector** for the Mesh Vertex/Fragment. |

## Controls

| Name | Type | Options | Description |
| --- | --- | --- | --- |
| Space | Dropdown | Object, View, World, Tangent | Selects coordinate space of **Bitangent Vector** to output. |