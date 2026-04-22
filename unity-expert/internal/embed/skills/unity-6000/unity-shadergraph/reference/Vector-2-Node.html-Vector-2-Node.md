# Vector 2 Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Vector-2-Node.html)

# Vector 2 Node

## Description

Defines a **Vector 2** value in the shader. If [Ports](Port.html) **X** and **Y** are not connected with [Edges](Edge.html) this [Node](Node.html) defines a constant **Vector 2**, otherwise this [Node](Node.html) can be used to combine various **Float** values.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| X | Input | Float | None | Input x component value |
| Y | Input | Float | None | Input y component value |
| Out | Output | Vector 2 | None | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
float2 _Vector2_Out = float2(X, Y);
```