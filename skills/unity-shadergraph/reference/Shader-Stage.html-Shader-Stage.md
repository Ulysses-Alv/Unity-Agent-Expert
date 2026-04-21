# Shader Stage | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Shader-Stage.html)

# Shader Stage

## Description

**Shader Stage** refers to the part of the shader pipeline a [Node](Node.html) or [Port](Port.html) is part of. For example, **Vertex** or **Fragment**.

In Shader Graph, **Shader Stage** is defined per port but often all ports on a node are locked to the same **Shader Stage**. Ports on some nodes are unavailable in certain **Shader Stages** due to limitations in the underlying shader language. See the [Node Library](Node-Library.html) documentation for nodes that have **Shader Stage** restrictions.

## Shader Stage List

| Name | Description |
| --- | --- |
| Vertex | Operations calculated per vertex |
| Fragment | Operations calculated per fragment |