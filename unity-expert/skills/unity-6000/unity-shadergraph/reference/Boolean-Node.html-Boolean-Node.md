# Boolean Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Boolean-Node.html)

# Boolean Node

## Description

Defines a constant **Boolean** value in the [Shader Graph](index.html), although internally to the shader this is treated as a constant **float** value that is ether 0 or 1, similar to Shaderlab's [Toggle](https://docs.unity3d.com/ScriptReference/MaterialPropertyDrawer.html) property. Can be converted to a **Boolean** type [Property](Property-Types.html) via the [Node's](Node.html) context menu.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Boolean | None | Output value |

## Controls

| Control | Description |
| --- | --- |
| (Checkbox) | Defines the output value. |

## Generated Code Example

The following basic test code represents one possible outcome of this node with the Boolean value set to 0:

```
float _Boolean = 0;
```