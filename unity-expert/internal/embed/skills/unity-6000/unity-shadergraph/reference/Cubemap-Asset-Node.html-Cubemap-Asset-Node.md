# Cubemap Asset Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Cubemap-Asset-Node.html)

# Cubemap Asset Node

## Description

Defines a constant **Cubemap Asset** for use in the shader. To sample the **Cubemap Asset** it should be used in conjunction with a [Sample Cubemap Node](Sample-Cubemap-Node.html). When using a separate **Cubemap Asset Node** you can sample a **Cubemap** twice, with different parameters, without defining the **Cubemap** itself twice.

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Cubemap | None | Output value |

## Controls

| Control | Description |
| --- | --- |
| (Cubemap) | Defines the cubemap asset from the project. |