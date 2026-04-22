# Check how many shader variants you have

You can use logging and profiling tools to check how many **shader** variants Unity compiles, and identify ways you can remove ([strip](shader-variant-stripping.md)) variants to improve build times and reduce memory usage. You can do the following:

* [Get a list of shader variants the Editor uses](#get-shader-variants-editor-uses)
* [Get a list of shader variants Unity creates at build time](#get-shader-variants-unity-creates)
* [Get a list of shader variants Unity compiles at runtime](#get-shader-variants-unity-compiles)
* [Check how much memory shaders use at runtime](#check-memory-shaders-use-runtime)
* [Highlight missing shaders at runtime](#highlight-missing-shaders)

## Get a list of shader variants Unity creates at build time

After you build your project, open the `Editor.log` log file and search for `Compiling shader` to see which variants Unity compiles and strips. For the location of `Editor.log`, refer to [log files](log-files.md).

For example:

```
Compiling shader "Sprites/Default" pass "" (vp)
    Full variant space:         8
    After settings filtering:   8
    After built-in stripping:   4
    After scriptable stripping: 4
    Processed in 0.00 seconds
    starting compilation...
    finished in 0.03 seconds. Local cache hits 0 (0.00s CPU time), remote cache hits 0 (0.00s CPU time), compiled 4 variants (0.09s CPU time), skipped 0 variants
    Prepared data for serialisation in 0.00s
```

If you use the Universal **Render Pipeline** (URP) or the High Definition Render Pipeline (HDRP), refer to the following:

* [Reducing shader variants in URP](urp/shader-stripping-landing.md)
* [Reduce shader variants](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/reduce-shader-variants.html) in HDRP

## Get a list of shader variants Unity compiles at runtime

To get a list of shader variants Unity compiles at runtime, record which pipeline state objects (PSOs) your application uses with the experimental `GraphicStateCollection` API. For more information, refer to [Trace and manage PSO data collections](shader-pso-trace.md).

You can also log compiled shader variants in the **Console** window and in a log file. Follow these steps:

1. Go to **Edit > Project Settings > Graphics**.
2. Under **Shader Loading**, enable **Log Shader Compilation**.
3. When you build your project, enable ****Development Build**** in the [build settings](BuildSettings.md).
4. In the [Console Window](Console.md), select **Editor** and enable **Full Log [Developer Mode Only]**.
5. Start the application you built.

Unity prints a `Compiled shader` message in the Console Window when it compiles a shader for the GPU.

## Check how much memory shaders use at runtime

Use the [Memory Profiler module](ProfilerMemory.md) or the [Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@1.0/manual/index.html) to check how much memory shaders are using at runtime. If a shader uses a lot of memory, you can experiment with stripping its variants.

## Highlight missing shaders at runtime

In Unity 2022.3 and above, you can force Unity to show a pink error shader during runtime, when a Material tries to use a missing shader variant.

1. Go to **Edit > Project Settings > Player**.
2. Under **Other Settings**, in the **Shader Settings** section, select **Strict shader variant matching**.

You can also enable this in C# **scripts** using [`strictShaderVariantMatching`](../ScriptReference/PlayerSettings-strictShaderVariantMatching.md).

When you do this, Unity shows a warning in the console with the missing variant and its keywords. You can use this during [stripping](shader-variant-stripping.md) to check you don’t remove shader variants your project needs.