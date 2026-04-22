# Set a shader to require GPU features

You can use [`#pragma` directives](SL-PragmaDirectives.md) to indicate that a **shader** pass requires certain GPU features. At runtime, Unity compiles the shader only if the platform supports the features.

Use one of the following in the `HLSLPROGRAM` block of a shader:

* `#pragma require` to specify individual GPU features that a shader requires, for example `#pragma require integers mrt8`.
* `#pragma target` to specify a shader model that a shader requires, for example `#pragma target 4.0`. A shader model represents a specific group of GPU features.

For the full list of values, refer to [HLSL pragma require command reference](SL-Pragma-require.md) and [HLSL pragma target command reference](SL-Pragma-target.md).

For example:

```
SubShader
{
    Pass
    {

        HLSLPROGRAM

        // Require multiple GPU features
        #pragma require integers mrt8
        ...

        ENDHLSL
    }
}
```

You can add multiple `#pragma` statements. You can also add multiple GPU features or shader models to a single `#pragma` statement, using a space between each.

By default, Unity uses:

* `#pragma require geometry` if your shader uses `#pragma geometry` to define a geometry stage.
* `#pragma require tessellation` if your shader uses `#pragma hull` or `#pragma domain` to define a tessellation stage.

Unity displays a warning message when it adds `geometry` or `tessellation`. To avoid this warning, add the requirements explicitly or use an appropriate `#pragma target` value in your code.

## Limit the requirement to specific shader variants

To compile shader variants only when the runtime platform supports a GPU feature or shader model, use the following syntax:

* In `#pragma target`, add the keywords of the shader variants, separated by spaces. For example `#pragma target 4.0 EXAMPLE_KEYWORD`.
* In `#pragma require`, add a colon (`:`), followed by the keywords of the shader variants, separated by spaces. For example `#pragma require mrt8 : EXAMPLE_KEYWORD`.

In these examples, at build time, Unity compiles shader variants that use the `EXAMPLE_KEYWORD` only if the target platform supports the GPU feature or shader model.

If you use two or more keywords in `#pragma target` or `#pragma require`, Unity compiles the shader variants that use any of the keywords.

## DirectX12 (DX12), Vulkan and Metal graphics APIs features

**Note:** If you use the following shader keywords, Unity compiles shaders using the DXC compiler. DXC support in Unity is experimental, not supported on all platforms, and not ready for production use.

If you use the DirectX12 (DX12), Vulkan or Metal graphics APIs, you can use a shader keyword to target the following GPU features:

* Support for 16-bit data types.
* Support for wave operations in [compute shaders](class-ComputeShader.md).

Use the following syntax:

```
#pragma multi_compile _ <keyword>
```

You don’t need to add a `pragma require` directive.

Unity then does the following:

* Uses a compatible shader compiler and turns on compiler flags related to the feature.
* Automatically compiles both a shader variant with the keyword on and a shader variant with the keyword off.
* At runtime, automatically checks if the hardware supports the feature and selects the correct variant.

You can use an `#if` statement to make parts of your shader code conditional on whether the GPU supports the feature.

### Keywords

| multi-compile keyword | GPU feature | Keyword for conditional shader code |
| --- | --- | --- |
| `UNITY_DEVICE_SUPPORTS_NATIVE_16BIT` | Supports 16-bit data types. If you use this keyword, the layout of shader buffers might change, because data types such as `half` and `min16float` convert to 16-bit. | `UNITY_DEVICE_SUPPORTS_NATIVE_16BIT` |
| `UNITY_DEVICE_SUPPORTS_WAVE_ANY` | Supports wave operations of any size. Use this keyword only if you use wave operations where the size of the waves doesn’t matter. | `UNITY_HW_SUPPORTS_WAVE` |
| `UNITY_DEVICE_SUPPORTS_WAVE_8` | Supports wave operations with a wave size of 8. | `UNITY_HW_SUPPORTS_WAVE` |
| `UNITY_DEVICE_SUPPORTS_WAVE_16` | Supports wave operations with a wave size of 16. | `UNITY_HW_SUPPORTS_WAVE` |
| `UNITY_DEVICE_SUPPORTS_WAVE_32` | Supports wave operations with a wave size of 32. | `UNITY_HW_SUPPORTS_WAVE` |
| `UNITY_DEVICE_SUPPORTS_WAVE_64` | Supports wave operations with a wave size of 64. | `UNITY_HW_SUPPORTS_WAVE` |
| `UNITY_DEVICE_SUPPORTS_WAVE_128` | Supports wave operations with a wave size of 128. | `UNITY_HW_SUPPORTS_WAVE` |

If you use a keyword that targets a specific wave size, Unity sets a `UNITY_HW_WAVE_SIZE` define to the same wave size so you can use it in shader code.

Refer to [Declaring and using shader keywords in HLSL](SL-MultipleProgramVariants.md) for more information.

## Additional resources

* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [HLSL pragma target command reference](SL-Pragma-target.md)
* [HLSL pragma require command reference](SL-Pragma-require.md)