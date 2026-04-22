# Sign Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Sign-Node.html)

# Sign Node

## Description

Per component, returns -1 if the value of input **In** is less than zero, 0 if equal to zero and 1 if greater than zero.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| In | Input | Dynamic Vector | Input value |
| Out | Output | Dynamic Vector | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Sign_float4(float4 In, out float4 Out)
{
    Out = sign(In);
}
```