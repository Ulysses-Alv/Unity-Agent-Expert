# All Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/All-Node.html)

# All Node

## Description

Returns true if all components of the input **In** are non-zero. This is useful for [Branching](Branch-Node.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| In | Input | Dynamic Vector | None | Input value |
| Out | Output | Boolean | None | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_All_float4(float4 In, out float Out)
{
    Out = all(In);
}
```