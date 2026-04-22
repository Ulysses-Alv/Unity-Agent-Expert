# Offset command in ShaderLab reference

Sets the depth bias on the GPU.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Offset** | Yes | Yes | Yes | Yes |

## Syntax

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

| **Signature** | **Example syntax** | **Function** |
| --- | --- | --- |
| `Offset <factor>, <units>` | `Offset 1, 1` | Draw geometry closer to or further away from the camera, based on the given values. |

## Parameters

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| **factor** | Float, range –1 through 1. | Scales the maximum Z slope, also called the depth slope, to produce a variable depth offset for each polygon.  Polygons that are not parallel to the near and far clip planes have Z slope. Adjust this value to avoid visual artifacts on such polygons. |
| **units** | Float, range –1 through 1. | Scales the minimum resolvable depth buffer value, to produce a constant depth offset. The minimum depth resolvable depth buffer value (one *unit*) varies by device.  A negative value means that the GPU draws the polygon closer to the camera. A positive value means that the GPU draws the polygon further away from the camera. |

## Additional resources

* [Set the depth bias in a shader](writing-shader-set-depth-bias.md)