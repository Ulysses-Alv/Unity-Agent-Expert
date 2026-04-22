# Customize how shaders contribute lightmap data in the Built-In Render Pipeline

This page contains information about how to make **shaders** compatible with Unity’s **lightmappers**.

## The Meta Pass

A Meta Pass is a Shader pass that provides albedo and emission values to the **Global Illumination** system. These values are separate from those used in real-time rendering, meaning that you can use the Meta Pass to control how a **GameObject** looks from the point of view of the lighting baking system without affecting its appearance at runtime.

An example of when this would be useful is if you wanted the green moss on a cliff to generate exaggerated green indirect light in your **lightmaps**, but you didn’t want to recolor the **terrain** in the real-time pass of shader.

All of Unity’s built-in Materials have a Meta Pass, and the Standard Shader contains a Meta pass. If you are using these, you do not need to do anything to enable the Meta Pass. If you are using a custom Shader, you can add your own Meta Pass.

### Example Shader with a Meta pass

The Shader below allows you to specify an albedo color and albedo Texture that are used only by the lighting baking system, without affecting the appearance of the Material at runtime. In this example, emission is taken from the UVs; but any value could be used to control it.

```
Shader "Custom/metaPassShader"{

    Properties {
        _Color ("Color", Color)=(1,1,1,1)
        _MainTex ("Albedo (RGB)",2D)="white"{}
        _Glossiness ("Smoothness", Range(0,1))=0.5
        _Metallic ("Metallic", Range(0,1))=0.0

        _GIAlbedoColor ("Color Albedo (GI)", Color)=(1,1,1,1)
        _GIAlbedoTex ("Albedo (GI)",2D)="white"{}
    }

    SubShader {
    // ------------------------------------------------------------------
    // Extracts information for lightmapping, GI (emission, albedo, ...)
    // This pass is not used during regular rendering.
        Pass
        {
            Name "META"
            Tags {"LightMode"="Meta"}
            Cull Off
            CGPROGRAM

            #include"UnityStandardMeta.cginc"

            sampler2D _GIAlbedoTex;
            fixed4 _GIAlbedoColor;
            float4 frag_meta2 (v2f_meta i): SV_Target
            {
                // We're interested in diffuse & specular colors
                // and surface roughness to produce final albedo.

                FragmentCommonData data = UNITY_SETUP_BRDF_INPUT (i.uv);
                UnityMetaInput o;
                UNITY_INITIALIZE_OUTPUT(UnityMetaInput, o);
                fixed4 c = tex2D (_GIAlbedoTex, i.uv);
                o.Albedo = fixed3(c.rgb * _GIAlbedoColor.rgb);
                o.Emission = Emission(i.uv.xy);
                return UnityMetaFragment(o);
            }

            #pragma vertex vert_meta
            #pragma fragment frag_meta2
            #pragma shader_feature _EMISSION
            #pragma shader_feature _METALLICGLOSSMAP
            #pragma shader_feature ___ _DETAIL_MULX2
            ENDCG
        }

        Tags {"RenderType"="Opaque"}
        LOD 200

        CGPROGRAM
        // Physically-based Standard lighting model, and enable shadows on all light types
        #pragma surface surf Standard fullforwardshadows nometa
        // Use Shader model 3.0 target, to get nicer looking lighting
        #pragma target 3.0

        sampler2D _MainTex;

        struct Input {
            float2 uv_MainTex;
        };

        half _Glossiness;
        half _Metallic;
        fixed4 _Color;

        void surf (Input IN,inout SurfaceOutputStandard o){
            // Albedo comes from a texture tinted by color
            fixed4 c = tex2D (_MainTex, IN.uv_MainTex)* _Color;
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

### Meta Pass technical information

Enlighten Realtime Global Illumination and lightmapping use Unity’s Meta Pass to extract albedo values from surfaces and handle diffuse transport themselves by using surface albedo on each bounce.

Metallic surfaces with a black (or almost black) albedo bounce very little light diffusely. Because the lightmappers handle only diffuse light transport, this means that you may see very little bounced light from these types of surfaces. Unity’s built-in Meta Passes account for this by providing a boosted version of the metal hue color instead of a physically correct albedo. This means that you get some bounce even from metallic materials. If you want a different behaviour, you can create a custom meta pass.

The built-in Meta Passes do not handle spectral specular reflectance.

Note: If you are using **Enlighten** Realtime Global Illumination, the Meta pass in the Player is not as fast as [DynamicGI.SetEmissive](http://docs.unity3d.com/ScriptReference/DynamicGI.SetEmissive.html), but it is more flexible because you are not limited to a single color.

## Custom RGB transparency

By default, shaders in Unity use monochrome transparency. This means that Unity uses the alpha channel of the material color or albedo texture to evaluate light transmission through the material.

During lightmapping, you can use custom RGB transparency instead. This means that Unity uses the values of a given texture to evaluate light transmission through the material. This is useful when you want color-based transparency that is independent of the material color or albedo texture; for example, if you want to bake lighting that simulates the behavior of a light shining through a stained glass window.

To use custom RGB transparency during lightmapping, add the following line to your [ShaderLab](SL-Reference.md) code:

`_TransparencyLM ("Transmissive Texture", 2D) = "white" {}`

This creates a [material property](SL-Properties.md) that appears in the Material **Inspector** with the name “Transmissive Texture”. Assign the desired texture to this field.

## Additional resources

* [Direct and indirect lighting](direct-and-indirect-lighting.md)
* [Configuring lightmaps and baking](Lightmapping-configure.md)