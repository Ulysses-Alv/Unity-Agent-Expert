# Example of Particle System GPU instancing in a Surface Shader

Here is a complete working example of Surface **Shader** using **Particle** System GPU Instancing:

```
Shader "Instanced/ParticleMeshesSurface" {
    Properties {
        _Color ("Color", Color) = (1,1,1,1)
        _MainTex ("Albedo (RGB)", 2D) = "white" {}
        _Glossiness ("Smoothness", Range(0,1)) = 0.5
        _Metallic ("Metallic", Range(0,1)) = 0.0
    }
    SubShader {
        Tags { "RenderType"="Opaque" }
        LOD 200

        CGPROGRAM
        // Physically based Standard lighting model, and enable shadows on all light types
        // And generate the shadow pass with instancing support
        #pragma surface surf Standard nolightmap nometa noforwardadd keepalpha fullforwardshadows addshadow vertex:vert
        // Enable instancing for this shader
        #pragma multi_compile_instancing
        #pragma instancing_options procedural:vertInstancingSetup
        #pragma exclude_renderers gles
        #include "UnityStandardParticleInstancing.cginc"
        sampler2D _MainTex;
        struct Input {
            float2 uv_MainTex;
            fixed4 vertexColor;
        };
        fixed4 _Color;
        half _Glossiness;
        half _Metallic;
        void vert (inout appdata_full v, out Input o)
        {
            UNITY_INITIALIZE_OUTPUT(Input, o);
            vertInstancingColor(o.vertexColor);
            vertInstancingUVs(v.texcoord, o.uv_MainTex);
        }

        void surf (Input IN, inout SurfaceOutputStandard o) {
            // Albedo comes from a texture tinted by color
            fixed4 c = tex2D (_MainTex, IN.uv_MainTex) * IN.vertexColor * _Color;
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

There are a number of small differences to a regular [Surface Shader](SL-SurfaceShaders.md) in the above example, which make it work with particle instancing.

Firstly, you must add the following two lines to enable Procedural Instancing, and specify the built-in vertex setup function. This function lives in UnityStandardParticleInstancing.cginc, and loads the per-instance (per-particle) positional data:

```
        #pragma instancing_options procedural:vertInstancingSetup
        #include "UnityStandardParticleInstancing.cginc"
```

The other modification in the example is to the Vertex function, which has two extra lines that apply per-instance attributes, specifically, particle colors and [Texture Sheet Animation](PartSysTexSheetAnimModule.md) texture coordinates:

```
            vertInstancingColor(o.vertexColor);
            vertInstancingUVs(v.texcoord, o.uv_MainTex);
```