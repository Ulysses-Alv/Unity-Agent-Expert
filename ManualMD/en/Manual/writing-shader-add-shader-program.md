# Add an HLSL shader program

To add a High-Level **Shader** Language (HLSL) program to a [shader file](class-Shader.md), add the code to the `Pass` block. Wrap the code in `HLSLPROGRAM` and `ENDHLSL` macros. For example:

```
Shader "Examples/ExampleShader"
{
    SubShader
    {
        Pass
        {                
              Name "ExamplePassName"

              HLSLPROGRAM

                // Add HLSL shader program code here
                
              ENDHLSL
        }
    }
}
```

## Additional resources

* [ShaderLab: adding shader programs](shader-shaderlab-code-blocks.md)