# Decals Surface Shader example in the Built-In Render Pipeline

Decals are commonly used to add details to Materials at run time (for example, bullet impacts). They are especially useful in deferred rendering, because they alter the GBuffer before it is lit, therefore saving on performance.

In a typical scenario, Decals should be rendered after the opaque objects and should not be shadow casters, as seen in the **ShaderLab** “Tags” in the example below.

```
Shader "Example/Decal" {
  Properties {
    _MainTex ("Base (RGB)", 2D) = "white" {}
  }
  SubShader {
    Tags { "RenderType"="Opaque" "Queue"="Geometry+1" "ForceNoShadowCasting"="True" }
    LOD 200
    Offset -1, -1
    
    CGPROGRAM
    #pragma surface surf Lambert decal:blend
    
    sampler2D _MainTex;
    
    struct Input {
      float2 uv_MainTex;
    };
    
    void surf (Input IN, inout SurfaceOutput o) {
        half4 c = tex2D (_MainTex, IN.uv_MainTex);
        o.Albedo = c.rgb;
        o.Alpha = c.a;
      }
    ENDCG
    }
}
```