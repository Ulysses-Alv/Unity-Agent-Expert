# Reflection Surface Shader examples in the Built-In Render Pipeline

Here’s a **Shader** that does cubemapped reflection using built-in `worldRefl` input. It’s very similar to built-in Reflective/Diffuse Shader:

```
  Shader "Example/WorldRefl" {
    Properties {
      _MainTex ("Texture", 2D) = "white" {}
      _Cube ("Cubemap", CUBE) = "" {}
    }
    SubShader {
      Tags { "RenderType" = "Opaque" }
      CGPROGRAM
      #pragma surface surf Lambert
      struct Input {
          float2 uv_MainTex;
          float3 worldRefl;
      };
      sampler2D _MainTex;
      samplerCUBE _Cube;
      void surf (Input IN, inout SurfaceOutput o) {
          o.Albedo = tex2D (_MainTex, IN.uv_MainTex).rgb * 0.5;
          o.Emission = texCUBE (_Cube, IN.worldRefl).rgb;
      }
      ENDCG
    } 
    Fallback "Diffuse"
  }
```

Because it assigns the reflection color as **Emission**, we get a very shiny soldier:

![Surface shader example that uses the world reflection texture as emission texture.](../uploads/Main/SurfaceShaderWorldRefl.jpg)

Surface shader example that uses the world reflection texture as emission texture.

If you want to do reflections that are affected by **normal maps**, it needs to be slightly more involved: `INTERNAL_DATA` needs to be added to the `Input` structure, and `WorldReflectionVector` function used to compute per-pixel reflection vector after you’ve written the Normal output.

```
  Shader "Example/WorldRefl Normalmap" {
    Properties {
      _MainTex ("Texture", 2D) = "white" {}
      _BumpMap ("Bumpmap", 2D) = "bump" {}
      _Cube ("Cubemap", CUBE) = "" {}
    }
    SubShader {
      Tags { "RenderType" = "Opaque" }
      CGPROGRAM
      #pragma surface surf Lambert
      struct Input {
          float2 uv_MainTex;
          float2 uv_BumpMap;
          float3 worldRefl;
          INTERNAL_DATA
      };
      sampler2D _MainTex;
      sampler2D _BumpMap;
      samplerCUBE _Cube;
      void surf (Input IN, inout SurfaceOutput o) {
          o.Albedo = tex2D (_MainTex, IN.uv_MainTex).rgb * 0.5;
          o.Normal = UnpackNormal (tex2D (_BumpMap, IN.uv_BumpMap));
          o.Emission = texCUBE (_Cube, WorldReflectionVector (IN, o.Normal)).rgb;
      }
      ENDCG
    } 
    Fallback "Diffuse"
  }
```

Here’s a normal-mapped shiny soldier:

![Surface shader example that uses the world reflection texture as emission texture and applies a normal texture to it.](../uploads/Main/SurfaceShaderWorldReflNormalmap.jpg)

Surface shader example that uses the world reflection texture as emission texture and applies a normal texture to it.