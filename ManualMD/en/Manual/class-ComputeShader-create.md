# Create a compute shader

Similar to [shader assets](class-Shader.md), compute **shader** assets are files in your project. with a *.compute* file extension. They are written in DirectX 11 style [HLSL](http://msdn.microsoft.com/en-us/library/windows/desktop/bb509561.aspx) language, with a minimal number of #pragma compilation directives to indicate which functions to compile as compute shader kernels.

## Check if a platform supports compute shaders

Compute shader support can be queried at runtime using [SystemInfo.supportsComputeShaders](../ScriptReference/SystemInfo-supportsComputeShaders.md).

## Create a compute shader asset

Here’s a basic example of a compute shader file, which fills the output texture with red:

```
// test.compute

#pragma kernel FillWithRed

RWTexture2D<float4> res;

[numthreads(1,1,1)]
void FillWithRed (uint3 dtid : SV_DispatchThreadID)
{
    res[dtid.xy] = float4(1,0,0,1);
}
```

The language is standard DX11 HLSL, with an additional `#pragma kernel FillWithRed` directive. One compute shader Asset file must contain at least one`compute kernel` that can be invoked, and that function is indicated by the `#pragma directive`. There can be more kernels in the file; just add multiple `#pragma kernel` lines.

When using multiple `#pragma kernel` lines, note that comments of the style `// text` are not permitted on the same line as the `#pragma kernel` directives, and cause compilation errors if used.

The `#pragma kernel` line can optionally be followed by a number of preprocessor macros to define while compiling that kernel, for example:

```
#pragma kernel KernelOne SOME_DEFINE DEFINE_WITH_VALUE=1337
#pragma kernel KernelTwo OTHER_DEFINE
// ...
```