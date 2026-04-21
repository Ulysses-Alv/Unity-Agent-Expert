# Is Front Face Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Is-Front-Face-Node.html)

# Is Front Face Node

## Description

Returns true if currently rendering a front face and false if rendering a back face. This value is always true unless the [Master Node](Master-Node.md)'s **Two Sided** value is set to true in the **Material Options**. This is useful for [Branching](Branch-Node.html).

NOTE: This [Node](Node.html) can only be used in the **Fragment** [Shader Stage](Shader-Stage.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Boolean | None | Output value |