# Integer Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Integer-Node.html)

# Integer Node

## Description

Defines a constant **Float** value in the shader using an **Integer** field. Can be converted to a **Float** type [Property](Property-Types.html) with a **Mode** setting of **Integer** via the [Node's](Node.html) context menu.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Float | None | Output value |

## Controls

| Control | Description |
| --- | --- |
| (Integer) | Defines the output value. |

## Generated Code Example

The following example code represents one possible outcome of this node.

```
float _Integer = 1;
```