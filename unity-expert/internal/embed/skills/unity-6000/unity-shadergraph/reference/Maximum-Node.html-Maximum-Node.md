# Maximum Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Maximum-Node.html)

# Maximum Node

## Description

Returns the largest of the two inputs values **A** and **B**.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| A | Input | Dynamic Vector | First input value |
| B | Input | Dynamic Vector | Second input value |
| Out | Output | Dynamic Vector | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Maximum_float4(float4 A, float4 B, out float4 Out)
{
    Out = max(A, B);
}
```