# Receiving shadows shader example in the Built-In Render Pipeline

![A cat-like character with accurate shadowing from the sun.](../uploads/SL/ExampleShadowReceiving.png)

A cat-like character with accurate shadowing from the sun.

Implementing support for receiving shadows will require compiling the base lighting pass into
several variants, to handle cases of “directional light without shadows” and “directional light with shadows” properly. **#pragma multi\_compile\_fwdbase** directive does this (see
[multiple shader variants](SL-MultipleProgramVariants.md) for details). In fact it does a lot more:
it also compiles variants for the different **lightmap** types, **Enlighten** Realtime **Global Illumination** being on or off etc. Currently we don’t need all that, so we’ll explicitly skip these variants.

Then to get actual shadowing computations, we’ll **#include “AutoLight.cginc”** **shader** [include file](SL-BuiltinIncludes.md) and use SHADOW\_COORDS, TRANSFER\_SHADOW, SHADOW\_ATTENUATION macros from it.

Here’s the shader:

```
Shader "Lit/Diffuse With Shadows"
{
    Properties
    {
        [NoScaleOffset] _MainTex ("Texture", 2D) = "white" {}
    }
    SubShader
    {
        Pass
        {
            Tags {"LightMode"="ForwardBase"}
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc"
            #include "Lighting.cginc"

            // compile shader into multiple variants, with and without shadows
            // (we don't care about any lightmaps yet, so skip these variants)
            #pragma multi_compile_fwdbase nolightmap nodirlightmap nodynlightmap novertexlight
            // shadow helper functions and macros
            #include "AutoLight.cginc"

            struct v2f
            {
                float2 uv : TEXCOORD0;
                SHADOW_COORDS(1) // put shadows data into TEXCOORD1
                fixed3 diff : COLOR0;
                fixed3 ambient : COLOR1;
                float4 pos : SV_POSITION;
            };
            v2f vert (appdata_base v)
            {
                v2f o;
                o.pos = UnityObjectToClipPos(v.vertex);
                o.uv = v.texcoord;
                half3 worldNormal = UnityObjectToWorldNormal(v.normal);
                half nl = max(0, dot(worldNormal, _WorldSpaceLightPos0.xyz));
                o.diff = nl * _LightColor0.rgb;
                o.ambient = ShadeSH9(half4(worldNormal,1));
                // compute shadows data
                TRANSFER_SHADOW(o)
                return o;
            }

            sampler2D _MainTex;

            fixed4 frag (v2f i) : SV_Target
            {
                fixed4 col = tex2D(_MainTex, i.uv);
                // compute shadow attenuation (1.0 = fully lit, 0.0 = fully shadowed)
                fixed shadow = SHADOW_ATTENUATION(i);
                // darken light's illumination with shadow, keep ambient intact
                fixed3 lighting = i.diff * shadow + i.ambient;
                col.rgb *= lighting;
                return col;
            }
            ENDHLSL
        }

        // shadow casting support
        UsePass "Legacy Shaders/VertexLit/SHADOWCASTER"
    }
}
```

Look, we have shadows now!