# Cosine Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Cosine-Node.html)

# Cosine Node

## Description

Returns the cosine of the value of input **In**.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| In | Input | Dynamic Vector | Input value in radians |
| Out | Output | Dynamic Vector | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Cosine_float4(float4 In, out float4 Out)
{
    Out = cos(In);
}
```