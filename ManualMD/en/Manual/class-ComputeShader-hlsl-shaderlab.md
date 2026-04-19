# Use HLSL and ShaderLab in a compute shader

## Use texture samplers

Textures and samplers aren’t separate objects in Unity, so to use them in compute **shaders** you must follow one of the following Unity-specific rules:

* Use the same name as the Texture name, with `sampler` at the beginning (for example, `Texture2D MyTex`; `SamplerState samplerMyTex`). In this case, the sampler is initialized to that Texture’s filter/wrap/aniso settings.
* Use a predefined sampler. For this, the name has to have `Linear` or `Point` (for filter mode) and `Clamp` or `Repeat` (for wrap mode). For example, `SamplerState MyLinearClampSampler` creates a sampler that has Linear filter mode and Clamp wrap mode.

For more information, see documentation on [Sampler States](SL-SamplerStates.md).

## Use keywords

You can use keywords to produce multiple variants of compute shaders, the same as you can for graphics shaders.

For general information on variants, see [Shader variants](shader-variants.md). For information on how to implement these features in compute shaders, see [Declaring and using shader keywords in HLSL](SL-MultipleProgramVariants.md) and the [ComputeShader API documentation](../ScriptReference/ComputeShader.md).

## HLSL-only or GLSL-only compute shaders

Usually, compute shader files are written in [HLSL](https://en.wikipedia.org/wiki/High-Level_Shading_Language), and compiled or translated into all necessary platforms automatically. However, it is possible to either prevent translation to other languages (that is, only keep HLSL platforms), or to write [GLSL](https://en.wikipedia.org/wiki/OpenGL_Shading_Language) compute code manually.

The following information only applies to HLSL-only or GLSL-only compute shaders, not cross-platform builds. This is because this information can result in compute shader source being excluded from some platforms.

* Compute shader source surrounded by `CGPROGRAM` and `ENDCG` keywords is not processed for non-HLSL platforms.
* Compute shader source surrounded by `GLSLPROGRAM` and `ENDGLSL` keywords is treated as GLSL source, and emitted verbatim. This only works when targeting OpenGL or GLSL platforms. You should also note that while automatically translated shaders follow HLSL data layout on buffers, manually written GLSL shaders follow GLSL layout rules.