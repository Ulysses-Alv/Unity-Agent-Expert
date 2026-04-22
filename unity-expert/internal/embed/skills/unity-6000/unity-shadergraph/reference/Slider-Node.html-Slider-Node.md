# Slider Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Slider-Node.html)

# Slider Node

## Description

Defines a constant **Float** value in the shader using a **Slider** field. Can be converted to a **Float** type [Property](Property-Types.html) with a **Mode** setting of **Slider** via the [Node's](Node.html) context menu.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Float | None | Output value |

## Controls

| Name | Type | Options | Description |
| --- | --- | --- | --- |
|  | Slider |  | Defines the output value. |
| Min | Float |  | Defines the slider parameter's minimum value. |
| Max | Float |  | Defines the slider parameter's maximum value. |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
float _Slider_Out = 1.0;
```