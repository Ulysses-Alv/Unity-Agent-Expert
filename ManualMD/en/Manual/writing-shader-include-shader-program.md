# Reuse HLSL code

To tell the compiler to include existing High-Level **Shader** Language (HLSL) code in another file or section, use the following methods:

* Include an HLSL file in another, using the standard HLSL `#include` directive.
* Duplicate HLSL in multiple places using the `HLSLINCLUDE` and `ENDHLSL` macros.
* Apply a `#pragma` directive to multiple HLSL files, using the Unity `#include_with_pragmas` directive.

## Include an HLSL file in another

Unity supports the standard HLSL `#include` directive. For more information, refer to [include directive](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-appendix-pre-include) in the Microsoft HLSL documentation.

## Duplicate HLSL in multiple places

To duplicate HLSL code in all the sections of your code, for example to add the same variables in multiple `Pass` blocks, wrap the code you want to reuse in `HLSLINCLUDE` and `ENDHLSL` macros. Unity adds the code at the top of every `HLSLPROGRAM` block.

For example:

```
Shader "Examples/ExampleShader"
{
    HLSLINCLUDE
        // Add HLSL code you want to share across the file.
    ENDHLSL

    SubShader
    {
        Pass
        {                
              ...

              HLSLPROGRAM
                // Unity pastes the contents of the HLSLINCLUDE block here.
              ENDHLSL
        }

        Pass
        {                
              ...

              HLSLPROGRAM
                // Unity pastes the contents of the HLSLINCLUDE block here.
              ENDHLSL
        }

    }
}
```

## Apply a pragma directive to multiple files

To apply a `#pragma` directive to multiple shader files, use the Unity `#include_with_pragmas` directive.

You can use this method to add an argument in multiple files, but toggle the arguments on and off by updating only one file. For example, you can create and toggle a debugging keyword that enables and disables conditional debugging code in multiple shaders.

1. Add a `#pragma` directive to a separate shader file. For example, create a one-line HLSL file with a pragma that enables shader debug symbols:

```
  #pragma enable_d3d11_debug_symbols
```

1. In each shader you want to include the directive in, add an `#include_with_pragmas` directive to include the file. For example:

```
  #include_with_pragmas "path-to-include-file"
```

**Note:** If a shader file uses `#include` to import a file that contains an `#include_with_pragmas` directive, Unity ignores the `#pragma` directives in the file the `#include_with_pragmas` directive references. Unity warns you about this in the console.

## Additional resources

* [ShaderLab: adding shader programs](shader-shaderlab-code-blocks.md)