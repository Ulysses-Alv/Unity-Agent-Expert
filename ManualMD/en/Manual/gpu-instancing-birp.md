# GPU Instancing in the Built-In Render Pipeline

This section contains information on how to add GPU instancing support to a custom Unity **shader**. It first explains the shader keywords, variables, and functions custom Unity shaders require to support GPU instancing. Then it includes examples of how to add per-instance data to both **surface shaders** and vertex/fragment shaders.

## Render pipeline compatibility

| **Feature** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom Scriptable Render Pipeline (SRP)** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Custom GPU instanced shaders** | No | No | No | Yes |

## Additional resources

* [GPU instancing](GPUInstancing-landing.md)
* [Graphics performance and profiling](graphics-performance-profiling.md)