# Set the depth testing mode in a shader

Depth testing allows GPUs that have “Early-Z” functionality to reject geometry early in the pipeline, and also ensures correct ordering of the geometry. You can change the conditions of depth testing to achieve visual effects such as object occlusion.

## Examples

This example code demonstrates the syntax for using this command in a Pass block.

```
Shader "Examples/CommandExample"
{
    SubShader
    {
         // The rest of the code that defines the SubShader goes here.

        Pass
        {    
              // Sets the depth test operation to Equal for all pixels in this Pass
              // You would typically do this if you want to render the geometry exactly where already rendered geometry is
              ZTest Equal
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

This example code demonstrates the syntax for using this command in a SubShader block.

```
Shader "Examples/CommandExample"
{
    SubShader
    {
        // Sets the depth test operation to Equal for all pixels in this Pass
        // You would typically do this if you want to render the geometry exactly where already rendered geometry is
        ZTest Equal

         // The rest of the code that defines the SubShader goes here.        

        Pass
        {    
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [ZTest command in ShaderLab reference](SL-ZTest.md)