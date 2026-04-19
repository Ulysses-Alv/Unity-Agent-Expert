# Check or write to the stencil buffer in a shader

The **stencil buffer** stores an 8-bit integer value for each **pixel** in the frame buffer. Before executing the fragment **shader** for a given pixel, the GPU can compare the current value in the stencil buffer against a given reference value. This is called a stencil test. If the stencil test passes, the GPU performs the depth test. If the stencil test fails, the GPU skips the rest of the processing for that pixel. This means that you can use the stencil buffer as a mask to tell the GPU which pixels to draw, and which pixels to discard.

You would typically use the stencil buffer for special effects such as portals or mirrors. Additionally, the stencil buffer is sometimes used when rendering hard shadows, or [constructive solid geometry (CSG)](https://en.wikipedia.org/wiki/Constructive_solid_geometry?).

You can use the `Stencil` command to do two different things: to configure the stencil test, and to configure what the GPU writes to the stencil buffer. You can do both of these things in the same command, but the most common use case is to create one **Shader object** that masks out areas of the screen that other Shader objects cannot draw to. To do this, you need to configure the first Shader object to always pass the stencil test and write to the stencil buffer, and configure the others to perform a stencil test and not write to the stencil buffer.

Use the `Ref`, `ReadMask`, and `Comp` parameters to configure the stencil test. Use the `Ref`, `WriteMask`, `Pass`, `Fail`, and `ZFail` parameters to configure the stencil write operation.

This command makes a change to the render state. Use it in a `Pass` block to set the render state for that Pass, or use it in a `SubShader` block to set the render state for all Passes in that SubShader.

The stencil test equation is:

```
(ref & readMask) comparisonFunction (stencilBufferValue & readMask)
```

## Examples

```
Shader "Examples/CommandExample"
{
    SubShader
    {
         // The rest of the code that defines the SubShader goes here.

        Pass
        {    
             // All pixels in this Pass will pass the stencil test and write a value of 2 to the stencil buffer
             // You would typically do this if you wanted to prevent subsequent shaders from drawing to this area of the render target or restrict them to render to this area only
             Stencil
             {
                 Ref 2
                 Comp Always
                 Pass Replace
             }            

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
             // All pixels in this SubShader pass the stencil test only if the current value of the stencil buffer is less than 2
             // You would typically do this if you wanted to only draw to areas of the render target that were not "masked out"
             Stencil
             {
                 Ref 2
                 Comp Less
             }  

         // The rest of the code that defines the SubShader goes here.        

        Pass
        {    
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [Stencil command in ShaderLab reference](SL-Stencil.md)