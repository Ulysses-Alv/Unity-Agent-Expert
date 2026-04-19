# Troubleshooting asynchronous shader compilation

Advanced rendering solutions rely on generating data once and reusing it in later frames. If the Editor uses a placeholder **shader** during this process, it might pollute the generated data. If this happens, you can see the cyan color or other rendering artifacts in your **scene**, even after the shader variants have finished compiling.

To avoid this, you can [Disable asynchronous shader compilation](AsynchronousShaderCompilation-enable-or-disable.md).

## Customize compile time rendering

You can make your custom tools draw something other than the placeholder shader for each material. This way, you can avoid rendering in plain cyan, and instead draw something else while the shader variant compiles.

To check if a specific shader variant has compiled, see [Detecting asynchronous shader compilation](AsynchronousShaderCompilation-detect.md).

To trigger compilation manually, you can use [`ShaderUtil.CompilePass`](../ScriptReference/ShaderUtil.CompilePass.md). This way, you can avoid rendering with the cyan placeholder, and draw something else while the shader variant compiles.