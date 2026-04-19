# HLSL semantics reference

Explore the semantics you can use to input **mesh** data into custom shaders in High-Level **Shader** Language (HLSL) in Unity.

| **Semantic** | **Description** | **Type** |
| --- | --- | --- |
| `POSITION` | The vertex position. | `float3` or `float4` |
| `NORMAL` | The vertex normal. | `float3` |
| `TEXCOORD0` | The first UV coordinates. | `float2`, `float3`, or `float4` |
| `TEXCOORD1` | The second UV coordinates. | `float2`, `float3`, or `float4` |
| `TEXCOORD2` | The third UV coordinates. | `float2`, `float3`, or `float4` |
| `TEXCOORD3` | The fourth UV coordinates. | `float2`, `float3`, or `float4` |
| `TANGENT` | The tangent vector used for normal mapping. | `float4` |
| `COLOR` | The per-vertex color. | `float4` |
| `VPOS` | The **pixel** position in screen space. You can use `VPOS` only in the fragment shader. | Use `UNITY_VPOS_TYPE` for cross-platform compatibility, which is `float4` on most platforms. |
| `VFACE` | Indicates whether the pixel surface is facing the **camera**. | `float` |
| `SV_VertexID` | The index of the vertex. You can’t use `SV_VertexID` in **Surface Shaders**. | `uint` |

## Additional resources

* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [Semantics](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-semantics) in the Microsoft HLSL documentation