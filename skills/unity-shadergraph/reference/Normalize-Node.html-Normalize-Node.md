# Normalize Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Normalize-Node.html)

# Normalize Node

## Description

Returns the normalized value of input **In**. The output vector will have the same direction as input **In** but a length of 1.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| In | Input | Dynamic Vector | Input value |
| Out | Output | Dynamic Vector | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Normalize_float4(float4 In, out float4 Out)
{
    Out = normalize(In);
}
```