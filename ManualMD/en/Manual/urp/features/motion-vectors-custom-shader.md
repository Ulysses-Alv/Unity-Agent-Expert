# Output a motion vector texture in a custom shader in URP

For URP to render the **MotionVectors** pass for a **ShaderLab** **shader**, make sure that its active SubShader contains a pass with the following [LightMode tag](../urp-shaders/urp-shaderlab-pass-tags.md#lightmode):

```
Tags { "LightMode" = "MotionVectors" }
```

For example:

```
Shader “Example/MyCustomShaderWithPerObjectMotionVectors"
{
    SubShader
    {
        // ...other passes, SubShader tags and commands
    
        Pass
        {
            Tags { "LightMode" = "MotionVectors" }
            ColorMask RG
            
            HLSLPROGRAM
            
            // Your shader code goes here.
            
            ENDHLSL
        }
    }
}
```

For an example of adding motion vector pass support for features such as alpha clipping, **LOD** cross-fade, or alembic animation, refer to the implementation of the `MotionVectors` pass in URP pre-built ShaderLab shaders, for example, the `Unlit.shader` file. The rendering of your `MotionVectors` pass should match your non-motion-vector passes and should reflect any custom deformation and/or vertex animation a pass is performing.

If a custom shader is only intended for objects with transform motion or skinned animation, without using alpha clipping, LOD cross-fade, alembic animation, custom deformation, or vertex animation, the motion vector fallback shader provided by URP might be enough. To add the pre-built fallback shader, add the following ShaderLab command to your SubShader blocks:

```
Shader “Example/MyCustomShaderWithPerObjectMotionVectorFallback"
{
    SubShader
    {
        // ...other passes, SubShader tags, and commands
    
        UsePass "Hidden/Universal Render Pipeline/ObjectMotionVectorFallback/MOTIONVECTORS"
    }
}
```

**Note:** In Unity versions earlier than 2023.2, URP would automatically use the fallback pass for all SubShader blocks which don’t have a pass tagged with the `MotionVectors` [LightMode tag](../urp-shaders/urp-shaderlab-pass-tags.md). Starting from Unity 2023.2, this fallback logic is disabled for the following reasons:

* The fallback logic was only an initial implementation detail meant for URP’s own Materials.
* URP Materials now use material-type-specific motion vector passes to support features like alpha clip, LOD cross-fade, or alembic animation, making the fallback obsolete.
* The fallback logic was causing unintended visual artifacts for content which it was not applicable to (for example, decals which should not draw any object motion vectors).