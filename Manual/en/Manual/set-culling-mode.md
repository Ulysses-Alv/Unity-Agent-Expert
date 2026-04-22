# Set the culling mode in a shader

Culling is the process of determining what not to draw. Culling improves rendering efficiency, by not wasting GPU time drawing things that would not be visible in the final image.

By default, the GPU performs back-face culling; this means that it does not draw polygons that face away from the viewer. In general, the more you can reduce the rendering workload, the better; you should therefore change this setting only when necessary.

## Example

```
Shader "Examples/CommandExample"
{
    SubShader
    {
         // The rest of the code that defines the SubShader goes here.

        Pass
        {    
              // Disable culling for this Pass.
              // You would typically do this for special effects, such as transparent objects or double-sided walls.
              Cull Off
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [Cull command in ShaderLab reference](SL-Cull.md)