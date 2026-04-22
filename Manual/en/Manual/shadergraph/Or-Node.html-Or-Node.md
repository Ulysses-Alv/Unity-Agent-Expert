# Or Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Or-Node.html)

# Or Node

## Description

Returns true if either of the inputs **A** and **B** are true. This is useful for [Branching](Branch-Node.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| A | Input | Boolean | None | First input value |
| B | Input | Boolean | None | Second input value |
| Out | Output | Boolean | None | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Or_float(float In, out float Out)
{
    Out = A || B;
}
```