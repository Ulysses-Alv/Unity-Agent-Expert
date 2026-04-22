# ColorMask command in ShaderLab reference

Sets the color channel writing mask, which prevents the GPU from writing to channels in the render target.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ColorMask** | Yes | Yes | Yes | Yes |

## Syntax

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

| **Signature** | **Example syntax** | **Function** |
| --- | --- | --- |
| `ColorMask <channels>` | `ColorMask RGB` | Write to the given channels of the default render target. |
| `ColorMask <channels> <render target>` | `ColorMask RGB 2` | As above, but for a given render target. |

## Parameters

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| **render target** | Integer, 0 through 7. | The render target index. |
| **channels** | `0` | Disables color writes to the R, G, B, and A channels. |
|  | `R` | Enables color writes to the red channel. |
|  | `G` | Enables color writes to the green channel. |
|  | `B` | Enables color writes to the blue channel. |
|  | `A` | Enables color writes to the alpha channel. |
|  | Any combination of `R`, `G`, `B`, and `A` without spaces. For example: `RB` | Enables color writes to the given channels. |

## Additional resources

* [Set the color channels the GPU renders to](writing-shader-color-mask.md)