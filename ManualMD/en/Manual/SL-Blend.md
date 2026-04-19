# Blend command in ShaderLab reference

Determines how the GPU combines the output of the fragment **shader** with the render target.

The functionality of this command depends on the blending operation, which you can set using the BlendOp command. Note that while blending itself is supported on all graphics APIs and hardware, some blending operations have more limited support.

Enabling blending disables some optimizations on the GPU (mostly hidden surface removal/Early-Z), which can lead to increased GPU frame times.

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Blend** | Yes | Yes | Yes | Yes |

## Syntax

| **Signature** | **Example syntax** | **Function** |
| --- | --- | --- |
| `Blend <state>` | `Blend Off` | Disables blending for the default render target. This is the default value. |
| `Blend <render target> <state>` | `Blend 1 Off` | As above, but for a given render target. (1) |
| `Blend <source factor> <destination factor>` | `Blend One Zero` | Enables blending for the default render target. Sets blend factors for RGBA values. |
| `Blend <render target> <source factor> <destination factor>` | `Blend 1 One Zero` | As above, but for a given render target. (1) |
| `Blend <source factor RGB> <destination factor RGB>, <source factor alpha> <destination factor alpha>` | `Blend One Zero, Zero One` | Enables blending the default render target. Sets separate blend factors for RGB and alpha values. (2) |
| `Blend <render target> <source factor RGB> <destination factor RGB>, <source factor alpha> <destination factor alpha>` | `Blend 1 One Zero, Zero One` | As above, but for a given render target. (1) (2) |

If blending is enabled, the following occurs:

* If the BlendOp command is used, the blending operation is set to that value. Otherwise, the blending operation defaults to `Add`.
* If the blending operation is `Add`, `Sub`, `RevSub`, `Min`, or `Max`, the GPU multiplies the value of the output of the fragment shader by the source factor.
* If the blending operation is `Add`, `Sub`, `RevSub`, `Min`, or `Max`, the GPU multiplies the value that is already in the render target by the destination factor.
* The GPU performs the blending operation on the resulting values.

**Notes:**

1. Any signature that specifies a render target requires OpenGL 4.0+, `GL_ARB_draw_buffers_blend`, or OpenGL ES 3.2.
2. Separate RGB and alpha blending is not compatible with [advanced OpenGL blending operations](SL-BlendOp.md).

## Parameters

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| **render target** | Integer, range 0 through 7 | The render target index. |
| **state** | `Off` | Disables blending. |
| **factor** | `One` | The value of this input is one. Use this to use the value of the source or the destination color. |
|  | `Zero` | The value of this input is zero. Use this to remove either the source or the destination values. |
|  | `SrcColor` | The GPU multiplies the value of this input by the source color value. |
|  | `SrcAlpha` | The GPU multiplies the value of this input by the source alpha value. |
|  | `SrcAlphaSaturate` | The GPU multiplies the value of this input by the minimum value of `source alpha` and `(1 - destination alpha)` |
|  | `DstColor` | The GPU multiplies the value of this input by the frame buffer source color value. |
|  | `DstAlpha` | The GPU multiplies the value of this input by the frame buffer source alpha value. |
|  | `OneMinusSrcColor` | The GPU multiplies the value of this input by (1 - source color). |
|  | `OneMinusSrcAlpha` | The GPU multiplies the value of this input by (1 - source alpha). |
|  | `OneMinusDstColor` | The GPU multiplies the value of this input by (1 - destination color). |
|  | `OneMinusDstAlpha` | The GPU multiplies the value of this input by (1 - destination alpha). |

The blending equation is:

```
finalValue = sourceFactor * sourceValue operation destinationFactor * destinationValue
```

In this equation:

* `finalValue` is the value that the GPU writes to the destination buffer.
* `sourceFactor` is defined in the Blend command.
* `sourceValue` is the value output by the fragment shader.
* `operation` is the blending operation.
* `destinationFactor` is defined in the Blend command.
* `destinationValue` is the value already in the destination buffer.

## Additional resources

* [Set the blending mode in a shader](writing-shader-blending-modes.md)