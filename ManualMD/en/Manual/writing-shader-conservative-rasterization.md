# Enable conservative rasterization in a shader

Rasterization is a rendering technique that converts vector data (triangle projections) to **pixel** data (render target) by determining which pixels are covered by triangles. Normally, a GPU determines whether or not to rasterize a pixel by sampling points within the pixel to determine whether they are covered by the triangle; if the coverage is sufficient, then the GPU determines that the pixel is covered. Conservative **rasterization** means that the GPU rasterizes a pixel that is partially covered by a triangle, regardless of coverage. This is useful when certainty is required, such as when performing **occlusion culling**, **collision** detection on the GPU, or visibility detection.

Conservative rasterization means that the GPU generates more fragments on triangle edges; this leads to more fragment **shader** invocations, which can lead to increased GPU frame times.

Check for hardware support using the [SystemInfo.supportsConservativeRaster](../ScriptReference/Device.SystemInfo-supportsConservativeRaster.md) API. On hardware that does not support this command, it is ignored.

## Example

```
Shader "Examples/CommandExample"
{
    SubShader
    {
         // The rest of the code that defines the SubShader goes here.

        Pass
        {    
              // Enable conservative rasterization for this Pass.
              Conservative True
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [Conservative command in ShaderLab reference](SL-Conservative.md)