# Matrix Determinant | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Matrix-Determinant-Node.html)

# Matrix Determinant

## Description

Returns the determinant of the matrix defined by input **In**. It can be viewed as the scaling factor of the transformation described by the matrix.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| In | Input | Dynamic Matrix | Input value |
| Out | Output | Float | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_MatrixDeterminant_float4x4(float4x4 In, out float Out)
{
    Out = determinant(In);
}
```