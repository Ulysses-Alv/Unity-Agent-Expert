# Normal Reconstruct Z Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Normal-Reconstruct-Z-Node.html)

# Normal Reconstruct Z Node

## Description

Derives the correct Z value for generated normal maps using a given **X** and **Y** value from input **In**.

## Ports

| Name | Direction | Type | Description |
| --- | --- | --- | --- |
| In | Input | Vector 2 | Normal X and Y value |
| Out | Output | Vector 3 | Output value |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_NormalReconstructZ_float(float2 In, out float3 Out)
{
    float reconstructZ = sqrt(1.0 - saturate(dot(In.xy, In.xy)));
    float3 normalVector = float3(In.x, In.y, reconstructZ);
    Out = normalize(normalVector);
}
```