# Reduce aliasing with AlphaToMask mode

Alpha-to-coverage mode can reduce the excessive aliasing that occurs when you use multisample anti-aliasing (MSAA) with shaders that use alpha testing, such as vegetation shaders. To do this, it modifies the multisample coverage mask proportionally to the alpha value in the output of the fragment **shader** result.

This command is intended for use with MSAA. If you enable alpha-to-coverage mode when you are not using MSAA, the results can be unpredictable; different graphics APIs and GPUs handle this differently.

## Examples

```
Shader "Examples/CommandExample"
{
    SubShader
    {
         // The rest of the code that defines the SubShader goes here.

        Pass
        {    
              // Enable alpha-to-coverage mode for this Pass
              AlphaToMask On
            
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
         // Enable alpha-to-coverage mode for this SubShader
         AlphaToMask On

         // The rest of the code that defines the SubShader goes here.        

        Pass
        {    
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [AlphaToMask command in ShaderLab reference](SL-AlphaToMask.md)