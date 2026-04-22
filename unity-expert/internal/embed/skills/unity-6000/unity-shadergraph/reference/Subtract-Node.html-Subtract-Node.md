# Subtract Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Subtract-Node.html)

# Subtract Node

## Description

Returns the result of input **A** minus input **B**.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| A | Input | Dynamic Vector | First input value |
| B | Input | Dynamic Vector | Second input value |
| Out | Output | Dynamic Vector | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Subtract_float4(float4 A, float4 B, out float4 Out)
{
    Out = A - B;
}
```