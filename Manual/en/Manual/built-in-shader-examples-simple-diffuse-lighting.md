# Simple diffuse lighting shader example in the Built-In Render Pipeline

![A cat-like character lit using simple diffuse lighting, so surfaces facing away from the light are much darker.](../uploads/SL/ExampleDiffuseLighting.png)

A cat-like character lit using simple diffuse lighting, so surfaces facing away from the light are much darker.

The first thing we need to do is to indicate that our **shader** does in fact need lighting information passed to it. Unity’s [rendering pipeline](SL-RenderPipeline.md) supports various ways of rendering; here we’ll be using the default [forward rendering](RenderTech-ForwardRendering.md) one.

We’ll start by only supporting one directional light. Forward rendering in Unity works by rendering the main directional light, ambient, **lightmaps** and reflections in a single pass called **ForwardBase**. In the shader, this is indicated by adding a [pass tag](SL-PassTags.md): **Tags {“LightMode”=“ForwardBase”}**. This will make directional light data be passed into shader via some [built-in variables](SL-UnityShaderVariables.md).

Here’s the shader that computes simple diffuse lighting per vertex, and uses a single main texture:

```
Shader "Lit/Simple Diffuse"
{
    Properties
    {
        [NoScaleOffset] _MainTex ("Texture", 2D) = "white" {}
    }
    SubShader
    {
        Pass
        {
            // indicate that our pass is the "base" pass in forward
            // rendering pipeline. It gets ambient and main directional
            // light data set up; light direction in _WorldSpaceLightPos0
            // and color in _LightColor0
            Tags {"LightMode"="ForwardBase"}
        
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc" // for UnityObjectToWorldNormal
            #include "UnityLightingCommon.cginc" // for _LightColor0

            struct v2f
            {
                float2 uv : TEXCOORD0;
                fixed4 diff : COLOR0; // diffuse lighting color
                float4 vertex : SV_POSITION;
            };

            v2f vert (appdata_base v)
            {
                v2f o;
                o.vertex = UnityObjectToClipPos(v.vertex);
                o.uv = v.texcoord;
                // get vertex normal in world space
                half3 worldNormal = UnityObjectToWorldNormal(v.normal);
                // dot product between normal and light direction for
                // standard diffuse (Lambert) lighting
                half nl = max(0, dot(worldNormal, _WorldSpaceLightPos0.xyz));
                // factor in the light color
                o.diff = nl * _LightColor0;
                return o;
            }
            
            sampler2D _MainTex;

            fixed4 frag (v2f i) : SV_Target
            {
                // sample texture
                fixed4 col = tex2D(_MainTex, i.uv);
                // multiply by lighting
                col *= i.diff;
                return col;
            }
            ENDHLSL
        }
    }
}
```

This makes the object react to light direction - parts of it facing the light are illuminated, and parts facing away are not illuminated at all.