# Introduction to asynchronous shader compilation

Shader objects can contain of hundreds or thousands of [shader variants](shader-variants.md). If the Editor compiled all variants when loading a **Shader** object, the import process would be very slow. Instead, the Editor compiles shader variants on-demand, the first time it encounters them.

Compiling these shader variants can cause the Editor to stall for milliseconds or even seconds, depending on the graphics API and the complexity of the **Shader object**. To avoid these stalls, you can use asynchronous shader compilation to compile the shader variants in the background, and use placeholder Shader objects in the meantime.

## How asynchronous shader compilation works

Asynchronous shader compilation works like this:

1. When the Editor first encounters an uncompiled shader variant, it adds the shader variant to a compilation queue on a job thread. The progress bar in the bottom right corner of the Editor shows the status of the compilation queue.
2. While the shader variant is loading, Editor renders the geometry with a placeholder shader, which appears as a plain cyan color.
3. When the Editor has finished compiling the shader variant, it renders the geometry using the shader variant.

![Unity renders shader variants that are still compiling with a cyan dummy Shader until compilation has finished. The bottom right progress bar indicates the compilation queue progress.](../uploads/Main/cyan_dummy_shaders.png)

Unity renders shader variants that are still compiling with a cyan dummy Shader until compilation has finished. The bottom right progress bar indicates the compilation queue progress.

## Exceptions

The following exceptions apply:

* If you draw geometry using [`DrawProcedural`](../ScriptReference/Graphics.DrawProcedural.md) or [`CommandBuffer.DrawProcedural`](../ScriptReference/Rendering.CommandBuffer.DrawProcedural.md), the Editor doesn’t use a placeholder shader. Instead, the Editor just skips rendering this geometry until it has compiled the shader variant.
* The Unity Editor does not use asynchronous shader compilation with **Blit** operations. This is to guarantee correct output in the most common use cases.