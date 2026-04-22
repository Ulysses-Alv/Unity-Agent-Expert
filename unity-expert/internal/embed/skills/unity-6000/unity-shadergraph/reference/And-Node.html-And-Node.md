# And Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/And-Node.html)

# And Node

## Description

Returns true if both the inputs **A** and **B** are true. This is useful for [Branching](Branch-Node.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| A | Input | Boolean | None | First input value |
| B | Input | Boolean | None | Second input value |
| Out | Output | Boolean | None | Output value |

## Generated Code Example

```
void Unity_And(float A, float B, out float Out)
{
    Out = A && B;
}
```