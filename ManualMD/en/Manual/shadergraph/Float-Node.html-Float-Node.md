# Float Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Float-Node.html)

# Float Node

## Description

Defines a **Float** value in the shader. If [Port](Port.html) **X** is not connected with an [Edge](Edge.html) this [Node](Node.html) defines a constant **Float**.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| X | Input | Float | None | Input x component value |
| Out | Output | Float | None | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
float _Vector1_Out = X;
```