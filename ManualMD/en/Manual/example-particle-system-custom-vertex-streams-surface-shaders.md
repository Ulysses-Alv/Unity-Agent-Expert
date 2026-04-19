# Example of Particle System Surface Shader custom vertex streams in the Built-In Render Pipeline

It’s possible to use Surface **Shaders** with custom vertex streams, although there are some extra things to be aware of:

* The input structure to your surface function is not the same as the input structure to the **vertex Shader**. You need to provide your own vertex Shader input structure. See below for an example, where it is called `appdata_particles`.
* When **surface Shaders** are built, there is automatic handling of variables whose names begin with certain tokens. The most notable one is `uv`. To prevent the automatic handling from causing problems here, be sure to give your UV inputs different names (for example, “texcoord”).

Here is the same functionality as the Standard Shader, but in a Surface Shader:

```
Shader "Particles/Anim Alpha Blend Surface" {
    Properties {
        _Color ("Color", Color) = (1,1,1,1)
        _MainTex ("Albedo (RGB)", 2D) = "white" {}
        _Glossiness ("Smoothness", Range(0,1)) = 0.5
        _Metallic ("Metallic", Range(0,1)) = 0.0
    }
    SubShader {
        Tags {"Queue"="Transparent" "RenderType"="Transparent"}
        Blend SrcAlpha OneMinusSrcAlpha
        ZWrite off
        LOD 200
        
        CGPROGRAM
        // Physically based Standard lighting model, and enable shadows on all light types
        #pragma surface surf Standard alpha vertex:vert

        // Use shader model 3.0 target, to get nicer looking lighting
        #pragma target 3.0

        sampler2D _MainTex;

         struct appdata_particles {
            float4 vertex : POSITION;
            float3 normal : NORMAL;
            float4 color : COLOR;
            float4 texcoords : TEXCOORD0;
            float texcoordBlend : TEXCOORD1;
            };

        struct Input {
            float2 uv_MainTex;
            float2 texcoord1;
            float blend;
            float4 color;
        };

        void vert(inout appdata_particles v, out Input o) {
            UNITY_INITIALIZE_OUTPUT(Input,o);
            o.uv_MainTex = v.texcoords.xy;
            o.texcoord1 = v.texcoords.zw;
            o.blend = v.texcoordBlend;
            o.color = v.color;
          }

        half _Glossiness;
        half _Metallic;
        fixed4 _Color;

        void surf (Input IN, inout SurfaceOutputStandard o) {
            fixed4 colA = tex2D(_MainTex, IN.uv_MainTex);
            fixed4 colB = tex2D(_MainTex, IN.texcoord1);
            fixed4 c = 2.0f * IN.color * lerp(colA, colB, IN.blend) * _Color;
                 
            o.Albedo = c.rgb;
            // Metallic and smoothness come from slider variables
            o.Metallic = _Metallic;
            o.Smoothness = _Glossiness;
            o.Alpha = c.a;
        }
        ENDCG
    }
    FallBack "Diffuse"
}
```