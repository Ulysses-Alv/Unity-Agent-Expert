# Final color modifier Surface Shader example in the Built-In Render Pipeline

It is possible to use a “final color modifier” function that will modify the final color computed by the **Shader**.The **Surface Shader** compilation directive `finalcolor:functionName` is used for this, with a function that takes `Input IN, SurfaceOutput o, inout fixed4 color` parameters.

Here’s a simple Shader that applies tint to the final color. This is different from just applying tint to the surface Albedo color: this tint will also affect any color that comes from **Lightmaps**, **Light Probes** and similar extra sources.

```
  Shader "Example/Tint Final Color" {
    Properties {
      _MainTex ("Texture", 2D) = "white" {}
      _ColorTint ("Tint", Color) = (1.0, 0.6, 0.6, 1.0)
    }
    SubShader {
      Tags { "RenderType" = "Opaque" }
      CGPROGRAM
      #pragma surface surf Lambert finalcolor:mycolor
      struct Input {
          float2 uv_MainTex;
      };
      fixed4 _ColorTint;
      void mycolor (Input IN, SurfaceOutput o, inout fixed4 color)
      {
          color *= _ColorTint;
      }
      sampler2D _MainTex;
      void surf (Input IN, inout SurfaceOutput o) {
           o.Albedo = tex2D (_MainTex, IN.uv_MainTex).rgb;
      }
      ENDCG
    } 
    Fallback "Diffuse"
  }
```

![Shader example that uses the final color modifier method.](../uploads/Main/SurfaceShaderFinalColorSimple.jpg)

Shader example that uses the final color modifier method.