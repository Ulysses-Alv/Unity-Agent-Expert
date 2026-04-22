# Description | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Tangent-Vector-Node.html)

## Description

Provides access to the mesh vertex or fragment's **Tangent Vector**. The coordinate space of the output value can be selected with the **Space** dropdown parameter.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Vector 3 | None | Mesh's **Tangent Vector**. |

## Parameters

| Name | Type | Options | Description |
| --- | --- | --- | --- |
| Space | Dropdown | Object, View, World, Tangent | Selects coordinate space of **Tangent Vector** to output. |