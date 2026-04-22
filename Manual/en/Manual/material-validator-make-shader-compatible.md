# Make a shader compatible with the Material Validator in the Built-In Render Pipeline

The Material Validator works with any Materials that use Unity’s [Standard shader](shader-StandardShader.md) or [surface shaders](SL-SurfaceShaders.md). However, custom **shaders** require a pass named `“META”`. Most custom shaders that support lightmapping already have this pass defined. See documentation on [Lightmapping and shaders](MetaPass.md) for more details.

Carry out the following steps to make your custom shader compatible with the Material Validator:

1. Add the following pragma to the meta pass: `#pragma shader_feature EDITOR_VISUALIZATION`
2. In the `UnityMetaInput` structure, assign the **specular color** of the Material to the field called `SpecularColor`, as shown in the code example below.

Here is an example of a custom meta pass:

```
Pass
{
    Name "META" 
    Tags { "LightMode"="Meta" }

    Cull Off

    CGPROGRAM
    #pragma vertex vert_meta
    #pragma fragment frag_meta

    #pragma shader_feature _EMISSION
    #pragma shader_feature _METALLICGLOSSMAP
    #pragma shader_feature _ _SMOOTHNESS_TEXTURE_ALBEDO_CHANNEL_A
    #pragma shader_feature ___ _DETAIL_MULX2
    #pragma shader_feature EDITOR_VISUALIZATION

    float4 frag_meta(v2f_meta i) : SV_TARGET
    {
        UnityMetaInput input;
        UNITY_INITIALIZE_OUTPUT(UnityMetaInput, input);
        float4 materialSpecularColor = float4(1.0f, 0.0f, 0.0f, 1.0f);
        float4 materialAlbedo = float4(0.0f, 1.0f, 0.0f, 1.0f);
        input.SpecularColor = materialSpecularColor;
        input.Albedo = materialAlbedo;

        return UnityMetaFragment(input);
    }  
}
```