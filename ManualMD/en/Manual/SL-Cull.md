# Cull command in ShaderLab reference

Sets which polygons the GPU should cull, based on the direction that they are facing relative to the **camera**.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Cull** | Yes | Yes | Yes | Yes |

## Syntax

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

| **Signature** | **Example syntax** | **Function** |
| --- | --- | --- |
| `Cull <state>` | `Cull Back` | Sets which polygons the GPU should cull, based on the direction that they face relative to the camera. |

## Parameters

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| **state** | `Back` | Cull polygons that face away from the camera. This is called back-face culling.  This is the default value. |
|  | `Front` | Cull polygons that face towards the camera. This is called front-face culling.  Use this for turning geometry inside-out. |
|  | `Off` | Do not cull polygons based on the direction that they face.  Use this for special effects, such as transparent objects or double-sided walls. |

## Additional resources

* [Set the culling mode in a shader](set-culling-mode.md)