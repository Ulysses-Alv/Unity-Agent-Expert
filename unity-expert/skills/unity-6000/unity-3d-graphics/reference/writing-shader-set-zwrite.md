# Disable writing to the depth buffer in a shader

`ZWrite` sets whether the **depth buffer** contents are updated during rendering. Normally, ZWrite is enabled for opaque objects and disabled for semi-transparent ones.

Disabling ZWrite can lead to incorrect depth ordering. In this case, you need to sort geometry on the CPU.

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
              // Disables writing to the depth buffer for this Pass
              ZWrite Off
            
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
         // Disables writing to the depth buffer for this SubShader
         ZWrite Off

         // The rest of the code that defines the SubShader goes here.        

        Pass
        {    
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional Resources

* [ZWrite command in ShaderLab reference](SL-ZWrite.md)