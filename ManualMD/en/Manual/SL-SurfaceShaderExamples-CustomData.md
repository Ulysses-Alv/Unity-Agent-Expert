# Custom data Surface Shader example in the Built-In Render Pipeline

Using a vertex modifier function, it is also possible to compute custom data in a vertex **Shader**, which then will be passed to the **Surface Shader** function per-pixel. The same compilation directive `vertex:functionName` is used, but the function should take two parameters: `inout appdata_full` and `out Input`. You can fill in any Input member that is not a built-in value there.

**Note:** Custom Input members used in this way must not have names beginning with ‘uv’ or they won’t work properly.

The example below defines a custom `float3 customColor` member, which is computed in a vertex function:

```
  Shader "Example/Custom Vertex Data" {
    Properties {
      _MainTex ("Texture", 2D) = "white" {}
    }
    SubShader {
      Tags { "RenderType" = "Opaque" }
      CGPROGRAM
      #pragma surface surf Lambert vertex:vert
      struct Input {
          float2 uv_MainTex;
          float3 customColor;
      };
      void vert (inout appdata_full v, out Input o) {
          UNITY_INITIALIZE_OUTPUT(Input,o);
          o.customColor = abs(v.normal);
      }
      sampler2D _MainTex;
      void surf (Input IN, inout SurfaceOutput o) {
          o.Albedo = tex2D (_MainTex, IN.uv_MainTex).rgb;
          o.Albedo *= IN.customColor;
      }
      ENDCG
    } 
    Fallback "Diffuse"
  }
```

In this example `customColor` is set to the absolute value of the normal:

![Shader example that calculates values in the vertex shader and uses them in the fragment shader.](../uploads/Main/SurfaceShaderCustomVertexData.jpg)

Shader example that calculates values in the vertex shader and uses them in the fragment shader.

More practical uses could be computing any per-vertex data that is not provided by built-in Input variables; or optimizing Shader computations. For example, it’s possible to compute Rim lighting at the **GameObject**’s vertices, instead of doing that in the Surface Shader per-pixel.