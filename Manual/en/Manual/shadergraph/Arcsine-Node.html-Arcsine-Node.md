# Arcsine Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Arcsine-Node.html)

# Arcsine Node

## Description

Returns the arcsine of each component of the input **In** as a vector of the same dimension and equal length. Each component should be within the range of -Pi/2 to Pi/2.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| In | Input | Dynamic Vector | Input value |
| Out | Output | Dynamic Vector | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Arcsine_float4(float4 In, out float4 Out)
{
    Out = asin(In);
}
```