# Conservative command in ShaderLab reference

Enables or disables conservative **rasterization**.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Conservative** | Yes | Yes | Yes | Yes |

## Syntax

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

| **Signature** | **Example syntax** | **Function** |
| --- | --- | --- |
| `Conservative <enabled>` | `Conservative True` | Enables or disables conservative rasterization.  Requires DX 11.3+, or GL\_NV\_conservative\_raster. |

## Parameters

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| **enabled** | `True` | Enables conservative rasterization. |
|  | `False` | Disables conservative rasterization. |

## Additional resources

* [Enable conservative rasterization in a shader](writing-shader-conservative-rasterization.md)