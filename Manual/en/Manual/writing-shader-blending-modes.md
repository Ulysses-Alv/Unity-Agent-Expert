# Set the blending mode in a shader

The `Blend` command determines how the GPU combines the output of the fragment **shader** with the render target.

The functionality of this command depends on the blending operation, which you can set using the BlendOp command. Note that while blending itself is supported on all graphics APIs and hardware, some blending operations have more limited support.

Enabling blending disables some optimizations on the GPU (mostly hidden surface removal/Early-Z), which can lead to increased GPU frame times.

## Common blend types

Here is the syntax for the most common blend types:

```
Blend SrcAlpha OneMinusSrcAlpha // Traditional transparency
Blend One OneMinusSrcAlpha // Premultiplied transparency
Blend One One // Additive
Blend OneMinusDstColor One // Soft additive
Blend DstColor Zero // Multiplicative
Blend DstColor SrcColor // 2x multiplicative
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
              // Enable regular alpha blending for this Pass
      Blend SrcAlpha OneMinusSrcAlpha
            
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
         // Enable regular alpha blending for this SubShader
         Blend SrcAlpha OneMinusSrcAlpha

         // The rest of the code that defines the SubShader goes here.        

        Pass
        {    
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [Blend command in ShaderLab reference](SL-Blend.md)
* [BlendOp command in ShaderLab reference](SL-BlendOp.md)