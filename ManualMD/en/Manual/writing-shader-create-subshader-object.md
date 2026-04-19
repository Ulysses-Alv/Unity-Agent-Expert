# Add a subshader in a custom shader

A **Shader** object contains one or more SubShaders. SubShaders let you define different GPU settings and shader programs for different hardware, **render pipelines**, and runtime settings. Some **Shader objects** contain only a single SubShader; others contain multiple SubShaders to support a range of different configurations.

For information on how a Shader object works, and the relationship between Shader objects, SubShaders and Passes, see [Shader objects introduction](shader-objects.md).

In **ShaderLab**, you define a SubShader by placing a `SubShader` block inside a `Shader` block.

## Example

This example code demonstrates the syntax for creating a Shader object that contains a single SubShader, which in turn contains a single Pass.

```
Shader "Examples/SinglePass"
{
    SubShader
    {
        Tags { "ExampleSubShaderTagKey" = "ExampleSubShaderTagValue" }
        LOD 100

         // ShaderLab commands that apply to the whole SubShader go here. 

        Pass
        {                
              Name "ExamplePassName"
              Tags { "ExamplePassTagKey" = "ExamplePassTagValue" }

              // ShaderLab commands that apply to this Pass go here.

              // HLSL code goes here.
        }
    }
}
```

## Additional resources

* [SubShader block in ShaderLab reference](SL-SubShader.md)
* [SubShader tags in ShaderLab reference](SL-SubShaderTags.md)