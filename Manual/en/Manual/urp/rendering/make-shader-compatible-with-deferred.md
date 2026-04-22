# Make a shader compatible with the Deferred or Deferred+ rendering paths in URP

## Use a shader in the Deferred or Deferred+ rendering paths

To use a **shader** in the Deferred **rendering path**, add the `UniversalGBuffer` tag to the Pass in your **ShaderLab** code. Unity executes the shader during the G-buffer render pass.

For example:

```
Shader "Examples/ExamplePassFlag"
{
    SubShader
    {
        Pass
        {    
              Tags
              { 
                "RenderPipeline" = "UniversalPipeline"
                "LightMode" = "UniversalGBuffer"
              }
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Use a shader in the forward pass of the Deferred or Deferred+ rendering paths

To use a shader in the Deferred rendering path, add the `UniversalForwardOnly` and `DepthNormalsOnly` tags to the Pass in your ShaderLab code. Unity executes the shader during the G-buffer render pass.

For example:

```
Shader "Examples/ExamplePassFlag"
{
    SubShader
    {
        Pass
        {    
              Tags { 
                "RenderPipeline" = "UniversalPipeline"
                "LightMode" = "UniversalForwardOnly"
                "LightMode" = "DepthNormalsOnly"
              }
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Specify the shader lighting model

To specify the shader lighting model as Lit or Simple Lit, use the `UniversalMaterialType` tag. For example:

```
Shader "Examples/ExamplePassFlag"
{
    SubShader
    {
        Pass
        {    
              Tags
              { 
                "RenderPipeline" = "UniversalPipeline"
                "LightMode" = "UniversalGBuffer"
                "UniversalMaterialType" = "Lit" 
              }
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [ShaderLab Pass tags in URP](../urp-shaders/urp-shaderlab-pass-tags.md)
* [Pass tags in ShaderLab reference](../../SL-PassTags.md)