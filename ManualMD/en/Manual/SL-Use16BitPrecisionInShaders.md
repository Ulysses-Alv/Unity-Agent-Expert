# Use 16-bit precision in shaders

By default, GPUs use 32-bit precision. You can use 16-bit precision instead in GPU calculations, which has the following benefits on mobile platforms:

* **Shaders** use less memory, bandwidth, and power.
* Calculations are faster. Using fewer bits can improve how the GPU allocates registers.

## Create a 16-bit variable

To use 16-bit precision in a shader, use `half` when you declare a scalar, a vector, or a matrix. For example:

```
half _Glossiness;
half4 _Color;
half4x4 _Matrix;
```

**Note:** Unity doesn’t support HLSL floating point suffixes. For example if you use `2.0h` to create a half precision float, Unity treats it as a high precision float `2.0f` instead.

To use 16-bit precision in a texture sampler, add `half4` when you declare it. For example:

```
Texture2D<half4> _MainTex;
```

16-bit precision might not be enough for some shader calculations. This might cause visible errors such as color banding or stuttering geometry. To check for errors, run your project on a platform that supports `half`. If there are errors, use `float` instead.

A `half` variable is stored in buffers with a size and alignment of 32 bits.

## Use 16-bit values on more platforms

By default, `half` has no effect on higher performance platforms in Unity, for example platforms that use MacOS. A `half` variable becomes a `float`, and the GPU uses 32-bit values for calculations.

To use 16-bit precision on more platforms, go to **Edit** > **Project Settings** > **Player** and set **Shader Precision Model** to **Uniform**. Unity then treats `half` in your HLSL code as the following:

* `min16float` for scalars
* `float` for texture samplers

When Unity compiles shaders, `min16float` becomes a platform data type that allows the GPU to use 16-bit precision for calculations if it supports them. For example:

* If you use DirectX 11 or 12, the variable remains a `min16float`.
* If you use OpenGL, the variable becomes `mediump`.
* If you use Vulkan, the variable becomes a `RelaxedPrecision` float.
* If you use Metal, the variable becomes a `float`, but the GPU uses 16-bit values for calculations.

To override the **Shader Precision Model** setting for textures, add `_half` when you declare a texture sampler. For example `Texture2D<half4> _MainTex`.

**Note**: In versions before Unity 6, the size of `min16float` varied depending on the target platform if you targeted the Metal graphics API. To enable this behavior in Unity 6 and later, include `#pragma metal_fxc_allow_float16_in_cpu_visible_buffers` in your shader.

## Additional resources

* [Data Types (HLSL)](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-data-types) in the Microsoft HLSL documentation
* [Optimize shaders](SL-ShaderPerformance.md)
* [Writing shaders for different graphics APIs](SL-PlatformDifferences.md)