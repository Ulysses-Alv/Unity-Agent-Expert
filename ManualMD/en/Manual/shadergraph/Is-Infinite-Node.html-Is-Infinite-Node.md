# Is Infinite Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Is-Infinite-Node.html)

# Is Infinite Node

## Description

Returns true if the input **In** is an infinite value. This is useful for [Branching](Branch-Node.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| In | Input | Float | None | Input value |
| Out | Output | Boolean | None | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_IsInfinite_float(float In, out float Out)
{
    Out = isinf(In);
}
```