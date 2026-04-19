# Pass information to shader compilers in HLSL

To pass information to the **shader** compilers in High-Level Shader Language (HLSL), use one of the following directives:

* The standard HLSL `#define` directive.
* The standard HLSL `#pragma` directive, to pass standard HLSL or Unity-specific directives.
* The Unity `#define_for_platform_compiler` directive.

## Add a #define directive

Unity supports the standard HLSL `#define` directive. For more information, see [#define directive](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-appendix-pre-define) in the Microsoft HLSL documentation. For example:

```
Shader "Custom/ExampleShader"
{
    SubShader
    {
        Pass
        {
            HLSLPROGRAM

            // Define a symbol called EXAMPLE_SYMBOL
            #define EXAMPLE_SYMBOL

            // Add your shader code here

            ENDHLSL
        }
    }
}
```

## Add a #pragma directive

Use a `#pragma` directive to do the following:

* Pass Unity-specific directives, for example to add shader keywords or target specific graphics APIs. For more information, refer to [HLSL pragma directives reference](SL-PragmaDirectives.md).
* Pass standard HLSL directives. For more information, refer to [#pragma directive](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-appendix-pre-pragma) in the Microsoft HLSL documentation.

Add a `#pragma` statement anywhere in your `HLSLPROGRAM`, but usually at the start. For example:

```
Shader "Custom/ExampleShader"
{
    SubShader
    {
        Pass
        {
            HLSLPROGRAM

            // Target shader model 3.0
            #pragma target 3.0

            // Compile this shader only for Direct3D 11 and OpenGL Core
            #pragma only_renderers d3d11 glcore

            // Add your shader code here

            ENDHLSL
        }
    }
}
```

### Limitations

Unity recognizes `#pragma` directives only in the following contexts:

* Unity-specific directives: In `.shader` files, or files you include with the `#include_with_pragmas` directive.
* Standard HLSL directives: In files you include with the `#include` directive.

Unity ignores `#pragma` directives outside these contexts.

**Note:** If a shader file uses `#include` to import a file that contains an `#include_with_pragmas` directive, Unity ignores the `#pragma` directives in the file the `#include_with_pragmas` directive references.

If you put a `#pragma` directive inside an `#if` statement, the condition must depend on one of the following:

* A custom symbol you define in a `#define` directive.
* The `SHADER_API_MOBILE` or `SHADER_API_DESKTOP` [platform macros](shader-branching-api.md) for branching based on the target platform.
* The `UNITY_NO_RGBM` or `UNITY_FRAMEBUFFER_FETCH_AVAILABLE` macros for [branching based on platform features](shader-branching-platform.md).
* The `UNITY_USE_NATIVE_HDR` macro.
* The `UNITY_NO_CUBEMAP_ARRAY` macro.
* The `UNITY_VERSION` macro for [branching based on Unity version](shader-branching-unity-version.md).

## Add a #define\_for\_platform\_compiler directive

To pass information to the platform-specific shader compiler, for example the FXC compiler if you target a platform that uses DirectX, add a `#define_for_platform_compiler` directive. For example:

```
Shader "Custom/ExampleShader"
{
    SubShader
    {
        Pass
        {
            HLSLPROGRAM

            // Define a symbol called EXAMPLE_SYMBOL for the platform-specific shader compiler
            #define_for_platform_compiler EXAMPLE_SYMBOL

            // Add your shader code here

            ENDHLSL
        }
    }
}
```

The directive sends a `#define` directive to the platform-specific shader compiler. For example, `#define_for_platform_compiler EXAMPLE_SYMBOL` sends a `#define EXAMPLE_SYMBOL` directive to the shader compiler.

The Unity preprocessor doesn’t use symbols you define with `#define_for_platform_compiler`. If you use the symbols inside conditional statements in your own shader code, the code won’t run.

## Additional resources

* [Shader compilation](shader-compilation.md)
* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [HLSL pragma target command reference](SL-Pragma-target.md)
* [HLSL pragma require command reference](SL-Pragma-require.md)
* [HLSL documentation](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-reference) on the Microsoft website