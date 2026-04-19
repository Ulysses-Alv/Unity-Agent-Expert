# Set the color channels the GPU renders to

By default, the GPU writes to all channels (RGBA). For some effects you might want to leave certain channels unmodified; for example, you can disable color rendering to render uncolored shadows. Another common use case is to disable color writes completely so that you can populate one buffer with data without writing to others; for example, you might want to populate the **stencil buffer** without writing to the render target.

## Examples

```
Shader "Examples/CommandExample"
{
    SubShader
    {
         // The rest of the code that defines the SubShader goes here.

        Pass
        {    
              // Enable writing only to the RGB channels for this Pass, which disables writing to the alpha channel
              ColorMask RGB

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
         // Enable writing only to the RGB channels for this SubShader, which disables writing to the alpha channel
         ColorMask RGB

         // The rest of the code that defines the SubShader goes here.        

        Pass
        {    
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [ColorMask command in ShaderLab reference](SL-ColorMask.md)