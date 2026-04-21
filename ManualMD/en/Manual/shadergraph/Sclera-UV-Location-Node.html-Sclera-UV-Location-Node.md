# Sclera UV Location Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Sclera-UV-Location-Node.html)

# Sclera UV Location Node

This node converts the object position of the sclera to a UV Sampling coordinate.

## Render pipeline compatibility

| **Node** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** |
| --- | --- | --- |
| **Sclera UV Location Node** | No | Yes |

## Ports

| name | **Direction** | type | description |
| --- | --- | --- | --- |
| **PositionOS** | Input | Vector3 | **Position of the fragment to shade in object space.** |
| **ScleraUV** | Output | Vector2 | Normalized UV coordinates that can be used to sample either a texture or procedurally generate a Sclera Texture. |