# Write a shader for a Custom Render Texture

To update a Custom **Render Texture** manually, you can write a specialized Custom Render Texture **shader**.

To help you write your Custom Render Texture shaders, here are two example frameworks that contain utility functions and built-in helper variables.

The following shader example fills the texture with a color multiplied by a color. When you write a shader for a Custom Render Texture, you must do the following:

* `#include "UnityCustomRenderTexture.cginc"`.
* Use the provided **Vertex Shader** `CustomRenderTextureVertexShader`.
* Use the provided input structure `v2f_customrendertexture` for the **pixel** shader.

```
Shader "CustomRenderTexture/Simple"
{
    Properties
    {
        _Color ("Color", Color) = (1,1,1,1)
        _Tex("InputTex", 2D) = "white" {}
     }

     SubShader
     {
        Lighting Off
        Blend One Zero

        Pass
        {
            HLSLPROGRAM
            #include "UnityCustomRenderTexture.cginc"
            #pragma vertex CustomRenderTextureVertexShader
            #pragma fragment frag
            #pragma target 3.0

            float4      _Color;
            sampler2D   _Tex;

            float4 frag(v2f_customrendertexture IN) : COLOR
            {
                return _Color * tex2D(_Tex, IN.localTexcoord.xy);
            }
            ENDHLSL
            }
    }
}
```

The following example is a shader for a material you can use to initialize a Custom Render Texture. When you write a shader for an initialization Material, the following steps are mandatory:

* `#include "UnityCustomRenderTexture.cginc"`
* Use the provided Vertex Shader `CustomRenderTextureVertexShader`
* Use the provided input structure `v2f_customrendertexture` for the pixel shader.

```
Shader "CustomRenderTexture/CustomTextureInit"
{
    Properties
    {
        _Color ("Color", Color) = (1,1,1,1)
        _Tex("InputTex", 2D) = "white" {}
    }

    SubShader
    {
        Lighting Off
        Blend One Zero

        Pass
        {
            HLSLPROGRAM
            #include "UnityCustomRenderTexture.cginc"

            #pragma vertex InitCustomRenderTextureVertexShader
            #pragma fragment frag
            #pragma target 3.0

            float4      _Color;
            sampler2D   _Tex;

            float4 frag(v2f_init_customrendertexture IN) : COLOR
            {
                return _Color * tex2D(_Tex, IN.texcoord.xy);
            }
            ENDHLSL
        }
    }
}
```