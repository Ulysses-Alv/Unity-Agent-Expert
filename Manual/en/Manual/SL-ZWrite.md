# ZWrite command in ShaderLab reference

Sets whether the **depth buffer** contents are updated during rendering. Normally, ZWrite is enabled for opaque objects and disabled for semi-transparent ones.

Disabling ZWrite can lead to incorrect depth ordering. In this case, you need to sort geometry on the CPU.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ZWrite** | Yes | Yes | Yes | Yes |

## Syntax

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

| **Signature** | **Example syntax** | **Function** |
| --- | --- | --- |
| ZWrite [state] | ZWrite Off | Enables or disables writing to the depth buffer. |

## Parameters

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| state | On | Enables writing to the depth buffer. |
|  | Off | Disables writing to the depth buffer. |

## Additional resources

* [Set the depth bias in a shader](writing-shader-set-depth-bias.md)
* [Set the depth clip mode in a shader](writing-shader-set-zclip.md)
* [Set the depth testing mode in a shader](writing-shader-set-ztest.md)