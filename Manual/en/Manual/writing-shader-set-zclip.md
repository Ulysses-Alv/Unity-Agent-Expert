# Set the depth clip mode in a shader

Setting the GPU’s depth clip mode to clamp can be useful for stencil shadow rendering; it means that there is no need for special case handling when the geometry goes beyond the far plane, which results in fewer rendering operations. However, it can result in incorrect Z ordering.

## Example

This example code demonstrates the syntax for using this command in a Pass block.

```
Shader "Examples/CommandExample"
{
    SubShader
    {
         // The rest of the code that defines the SubShader goes here.

        Pass
        {    
              // Sets the GPU's depth clip mode to clamp for this Pass
              // You would typically do this if you are rendering stencil shadows
              ZClip False
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [ZClip command in ShaderLab reference](SL-ZClip.md)