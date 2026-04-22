# Prebuilt shaders render pipeline compatibility reference

| ****Shaders**** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| [**URP shaders**](urp/shaders-in-universalrp.md) | Yes | No | No | No |
| [**HDRP shaders**](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/materials-and-surfaces.html) | No | Yes | No | No |
| [**Standard Shader**](shader-StandardShader-landing.md) | No  Can convert Standard Shader to equivalent on import | No  Can convert Standard Shader to equivalent on import | No | Yes |
| [**Standard Particle Shaders**](shader-StandardParticleShadersLanding.md) | No | No | No | Yes |
| [**Legacy shaders**](Built-inShaderGuide.md) | Yes  Simple unlit Legacy shaders will likely render fine, but they might not be SRP Batcher compatible. | Yes  Simple unlit Legacy shaders will likely render fine, but they might not be SRP Batcher compatible. They also do not support any HDRP features. For an unlit shader that supports HDRP features, use the HDRP/Unlit shader. | Yes  Simple unlit Legacy shaders will likely render fine, but they might not be SRP Batcher compatible. | Yes |