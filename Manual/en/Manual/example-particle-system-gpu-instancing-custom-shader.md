# Example of Particle System GPU instancing in the Built-In Render Pipeline

Here is a complete working example of a Custom **Shader** using **particle** system GPU instancing. This custom shader adds a feature which the standard particle shader does not have - a fade between the individual frames of a [texture sheet animation](PartSysTexSheetAnimModule.md).

```
Shader "Instanced/ParticleMeshesCustom"
{
    Properties
    {
        _MainTex("Albedo", 2D) = "white" {}
        [Toggle(_TSANIM_BLENDING)] _TSAnimBlending("Texture Sheet Animation Blending", Int) = 0
    }
    SubShader
    {
        Tags{ "RenderType" = "Opaque" }
        LOD 100
        Pass
        {
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #pragma multi_compile __ _TSANIM_BLENDING
            #pragma multi_compile_instancing
            #pragma instancing_options procedural:vertInstancingSetup
            #include "UnityCG.cginc"
            #include "UnityStandardParticleInstancing.cginc"
            struct appdata
            {
                float4 vertex : POSITION;
                fixed4 color : COLOR;
                float2 texcoord : TEXCOORD0;
                UNITY_VERTEX_INPUT_INSTANCE_ID
            };
            struct v2f
            {
                float4 vertex : SV_POSITION;
                fixed4 color : COLOR;
                float2 texcoord : TEXCOORD0;
#ifdef _TSANIM_BLENDING
                float3 texcoord2AndBlend : TEXCOORD1;   
#endif
            };
            sampler2D _MainTex;
            float4 _MainTex_ST;
            fixed4 readTexture(sampler2D tex, v2f IN)
            {
                fixed4 color = tex2D(tex, IN.texcoord);
#ifdef _TSANIM_BLENDING
                fixed4 color2 = tex2D(tex, IN.texcoord2AndBlend.xy);
                color = lerp(color, color2, IN.texcoord2AndBlend.z);
#endif
                return color;
            }
            v2f vert(appdata v)
            {
                v2f o;
                UNITY_SETUP_INSTANCE_ID(v);
                o.color = v.color;
                o.texcoord = v.texcoord;
                vertInstancingColor(o.color);
#ifdef _TSANIM_BLENDING
                vertInstancingUVs(v.texcoord, o.texcoord, o.texcoord2AndBlend);
#else
                vertInstancingUVs(v.texcoord, o.texcoord);
#endif
                o.vertex = UnityObjectToClipPos(v.vertex);
                return o;
            }
            fixed4 frag(v2f i) : SV_Target
            {
                half4 albedo = readTexture(_MainTex, i);
                return i.color * albedo;
            }
            ENDHLSL
        }
    }
}
```

This example contains the same set-up code as the **Surface Shader** for loading the positional data:

```
        #pragma instancing_options procedural:vertInstancingSetup
        #include "UnityStandardParticleInstancing.cginc"
```

The modification to the vertex function is very similar to the Surface Shader too:

```
                vertInstancingColor(o.color);
#ifdef _TSANIM_BLENDING
                vertInstancingUVs(v.texcoord, o.texcoord, o.texcoord2AndBlend);
#else
                vertInstancingUVs(v.texcoord, o.texcoord);
#endif
```

The only difference here, compared with the first example above, is the texture sheet animation blending. This means that the shader requires an extra set of texture coordinates to read two frames of the texture sheet animation instead of just one, and blends them together.

Finally, the fragment shader reads the textures and calculates the final color.