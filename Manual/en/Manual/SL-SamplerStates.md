# Sample a texture in HLSL

To sample a texture in High-Level **Shader** Language (HLSL), declare a texture and its sampler. The sampler contains the properties the GPU uses to sample the texture, for example the wrap mode and filter mode. The properties usually come from the [texture import settings](class-TextureImporter.md).

**Note:** If your project uses the Universal **Render Pipeline** (URP), refer instead to [Draw a texture in a shader in URP](urp/writing-shaders-urp-unlit-texture.md).

Follow these steps:

1. In the `HLSLPROGRAM` block, use the `UNITY_DECLARE_TEX2D` macro to declare the texture and its sampler. For example:

   ```
   UNITY_DECLARE_TEX2D(_MainTex);
   ```

   Unity creates a standard HLSL `Texture2D` object for the texture, and a standard HLSL `SamplerState` object for its sampler.
2. To sample the texture, use the `UNITY_SAMPLE_TEX2D` macro. For example:

   ```
   half4 color = UNITY_SAMPLE_TEX2D(_MainTex, uv);
   ```

If the graphics API doesn’t support separate textures and samplers, Unity does the following:

* The `UNITY_DECLARE_TEX2D` macro falls back to creating a standard HLSL `sampler2D` object instead, which combines the texture and its sampler.
* The `UNITY_SAMPLE_TEX2D_SAMPLER` macro samples the texture using the combined texture and sampler.

## Sample other textures with the same sampler

To sample a second texture with the sampler from a texture, follow these steps:

1. In the `HLSLPROGRAM` block, declare the second texture using the `UNITY_DECLARE_TEX2D_NOSAMPLER` macro. For example:

   ```
   UNITY_DECLARE_TEX2D_NOSAMPLER(_SecondTex);
   ```
2. To sample the texture, use the `UNITY_SAMPLE_TEX2D_SAMPLER` macro, passing in the first texture as the sampler source. For example:

   ```
   half4 color = UNITY_SAMPLE_TEX2D_SAMPLER(_SecondTex, _FirstTex, uv);
   ```

## Create a custom sampler

To create and use a custom sampler directly in a shader, also known as an inline sampler, follow these steps:

1. Declare a `SamplerState` object and use its name to set its properties. For example, the following code creates a custom sampler that uses point filtering and the clamp wrapping mode:

   ```
   SamplerState my_point_clamp_sampler;
   ```
2. Sample the texture using the `Sample` method of a `Texture2D` object, and pass in the custom sampler. For example:

   ```
   _MainTex.Sample(my_point_clamp_sampler, uv);
   ```

**Note**: Inline samplers don’t work if you target the OpenGL ES or OpenGL Core graphics APIs.

For more information, refer to [Built-in shader variables reference](SL-UnityShaderVariables.md).

## Example

The following example reuses the same sampler for three textures:

```
Shader "Custom/SeparateSamplerExample"
{
    Properties
    {
        _MainTex ("Main Texture", 2D) = "white" {}
        _SecondTex ("Second Texture", 2D) = "white" {}
        _ThirdTex ("Third Texture", 2D) = "white" {}
    }

    SubShader
    {
        Pass
        {
            HLSLPROGRAM

            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc"

            // Declare the first texture and its sampler.
            UNITY_DECLARE_TEX2D(_MainTex);

            // Declare the second and third textures without samplers.
            UNITY_DECLARE_TEX2D_NOSAMPLER(_SecondTex);
            UNITY_DECLARE_TEX2D_NOSAMPLER(_ThirdTex);

            struct v2f {
                float4 pos : SV_POSITION;
                float2 uv : TEXCOORD0;
            };

            v2f vert (appdata_base v)
            {
                v2f o;
                o.pos = UnityObjectToClipPos(v.vertex);
                o.uv = v.texcoord;
                return o;
            }

            half4 frag (v2f i) : SV_Target
            {
                // Sample the first texture using its own sampler.
                half4 color1 = UNITY_SAMPLE_TEX2D(_MainTex, i.uv);

                // Sample the second and third textures using the sampler from the first texture.
                half4 color2 = UNITY_SAMPLE_TEX2D_SAMPLER(_SecondTex, _MainTex, i.uv);
                half4 color3 = UNITY_SAMPLE_TEX2D_SAMPLER(_ThirdTex, _MainTex, i.uv);

                // Return the combined color from all three textures.
                return (color1 + color2 + color3) / 3.0;
            }
            ENDHLSL
        }
    }

}
```

## Additional resources

* [Use built-in shader functions](SL-BuiltinFunctions.md)
* [HLSL documentation](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-reference) on the Microsoft website