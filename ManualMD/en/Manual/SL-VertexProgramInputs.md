# Input data into HLSL

To pass data into a custom **shader** in High-Level Shader Language (HLSL), create a `struct` that declares shader variables and connects them to the **mesh** vertex data.

Use one of the following methods:

* Include a prebuilt vertex structure from Unity’s shader library.
* Declare a custom vertex structure.

## Include a prebuilt structure

Unity has a library of High-Level Shader Language (HLSL) shader files that contain prebuilt HLSL structures with common vertex inputs. To use a prebuilt structure, follow these steps:

1. Add `#include "UnityCG.cginc"` at the top of your `HLSLPROGRAM`. This imports the `UnityCG.cginc` file, which contains the prebuilt structures.
2. When you declare your **vertex shader** function, use one of the prebuilt structures as the input parameter type. For example, to use `appdata_base`, use `v2f vert (appdata_base vertex_data)`.

**Note:** If your project uses the Universal **Render Pipeline** (URP), don’t use `UnityCG.cginc` because it’s from the Built-in Render Pipeline. For information about URP shader library files, refer to [Shader methods in URP](urp/use-built-in-shader-methods.md) and [Upgrade custom shaders for URP compatibility](urp/urp-shaders/birp-urp-custom-shader-upgrade-guide.md)

For a list of prebuilt structures, refer to [Built-in shader variables reference](SL-UnityShaderVariables.md).

### Example

The following shader uses `appdata_base`, and colors the mesh based on its normals.

```
Shader "VertexInputSimple" {
    SubShader {
        Pass {
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc"
         
            // Declare the structure for data passed from the vertex shader to the fragment shader.
            struct v2f {
                float4 position : SV_POSITION;
                half4 color : COLOR;
            };
            
            // Use the built-in appdata_base structure for the vertex shader.
            v2f vert (appdata_base v)
            {
                v2f output_data;

                // Use the position from appdata_base.
                output_data.position = UnityObjectToClipPos(v.vertex);

                // Use the normal from appdata_base. 
                output_data.color.xyz = v.normal * 0.5 + 0.5;

                output_data.color.w = 1.0;
                return output_data;
            }

            half4 frag (v2f i) : SV_Target { 
                return i.color; 
            }

            ENDHLSL
        }
    } 
}
```

## Declare a custom structure

To declare a custom vertex structure, follow these steps:

1. At the top of your `HLSLPROGRAM`, add a `struct` that declares a variable for each input you want to access from the mesh.
2. In each variable declaration, add a semantic that connects the variable to the mesh data. For more information about semantics, refer to [Semantics](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-semantics) in the Microsoft HLSL documentation.

For a list of semantics, refer to [HLSL semantics reference](SL-HLSLSemantics.md).

For example:

```
Shader "VertexInputCustom" {
    SubShader {
        Pass {

            HLSLPROGRAM

            ...

            struct custom_vertex_input {

                // Declare a variable for the vertex position, and connect it to the mesh data using the POSITION semantic.
                float4 vertex : POSITION;

                // Declare a variable for the vertex normal, and connect it to the mesh data using the NORMAL semantic.
                float3 normal : NORMAL;
            };

            ...

            v2f vert (custom_vertex_input v) {
                ...

            }

            ENDHLSL

        }
    }
}
```

If a variable contains more components than the mesh data, the extra components are 0 for x, y, and z, and 1 for w. For example, if you declare a `float4` with `TEXCOORD0`, the value is (u, v, 0, 1).

## Passing data from the vertex shader to the fragment shader

The data you pass from the vertex shader to the fragment shader, known as interpolators, has limits based on the target platform and GPU.

To minimize the number of interpolators you use, pack data into single variables. For example, combine two sets of `float2` UV coordinates into one `float4` variable, where x and y contain the first UV coordinate and z and w contain the second UV coordinate.

## Add a custom buffer for uniform variables

DirectX groups all uniform variables into constant buffers. Rendering might be faster if you group uniform variables into separate constant buffers depending on how often they change.

To create a constant buffer, use the `CBUFFER_START` and `CBUFFER_END` Unity macros. For example:

```
CBUFFER_START(RarelyUpdatedVariables)
    float4 lightPosition;
CBUFFER_END

CBUFFER_START(FrequentlyUpdatedVariables)
    float4 colorTint;
CBUFFER_END
```

## Additional resources

* [Visualizing vertex data shader examples in the Built-In Render Pipeline](built-in-shader-examples-vertex-data.md)
* [Built-in shader variables reference](SL-UnityShaderVariables.md)
* [Customize how shaders contribute lightmap data in the Built-In Render Pipeline](MetaPass.md)
* [HLSL shader examples in the Built-In Render Pipeline](built-in-shader-examples.md)
* [Semantics](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-semantics) in the Microsoft HLSL documentation