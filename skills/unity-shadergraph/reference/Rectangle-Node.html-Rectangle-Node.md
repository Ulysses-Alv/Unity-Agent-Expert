# Rectangle Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Rectangle-Node.html)

# Rectangle Node

## Description

Generates a rectangle shape based on input **UV** at the size specified by inputs **Width** and **Height**. The generated shape can be offset or tiled by connecting a [Tiling And Offset Node](Tiling-And-Offset-Node.html). Note that in order to preserve the ability to offset the shape within the UV space the shape will not automatically repeat if tiled. To achieve a repeating rectangle effect first connect your input through a [Fraction Node](Fraction-Node.html).

NOTE: This [Node](Node.html) can only be used in the **Fragment** [Shader Stage](Shader-Stage.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| UV | Input | Vector 2 | UV | Input UV value |
| Width | Input | Float | None | Rectangle width |
| Height | Input | Float | None | Rectangle height |
| Out | Output | Float | None | Output value |

## Controls

| Control | Description |
| --- | --- |
| (Dropdown) | Select the robustness of computation. The options are:  * **Fastest** * **Nicest** |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
void Unity_Rectangle_float(float2 UV, float Width, float Height, out float Out)
{
    float2 d = abs(UV * 2 - 1) - float2(Width, Height);
    d = 1 - d / fwidth(d);
    Out = saturate(min(d.x, d.y));
}
```