# ZTest command in ShaderLab reference

Sets the conditions under which geometry passes or fails depth testing.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ZTest** | Yes | Yes | Yes | Yes |

## Syntax

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

| **Signature** | **Example syntax** | **Function** |
| --- | --- | --- |
| ZTest [operation] | ZTest Less | Set the conditions under which geometry passes or fails depth testing. |

## Parameters

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| operation |
|  | Disabled | Disable the depth test. |
|  | Never | Draw no geometry, regardless of distance. |
|  | Less | Draw geometry that is in front of existing geometry. Do not draw geometry that is at the same distance as or behind existing geometry. |
|  | Equal | Draw geometry that is at the same distance as existing geometry. Do not draw geometry that is in front of or behind existing geometry. |
|  | LEqual | Draw geometry that is in front of or at the same distance as existing geometry. Do not draw geometry that is behind existing geometry.  This is the default value. |
|  | Greater | Draw geometry that is behind existing geometry. Do not draw geometry that is at the same distance as or in front of existing geometry. |
|  | NotEqual | Draw geometry that is not at the same distance as existing geometry. Do not draw geometry that is at the same distance as existing geometry. |
|  | GEqual | Draw geometry that is behind or at the same distance as existing geometry. Do not draw geometry that is in front of existing geometry. |
|  | Always | No depth testing occurs. Draw all geometry, regardless of distance. |

## Additional resources

* [Set the depth bias in a shader](writing-shader-set-depth-bias.md)
* [Set the depth clip mode in a shader](writing-shader-set-zclip.md)
* [Set the depth testing mode in a shader](writing-shader-set-ztest.md)