# Single-pass instanced rendering and custom shaders

URP, HDRP, ShaderGraph, Surface **shaders**, and built-in shaders already support single-pass stereo instanced rendering. However, shaders from the **Asset Store**, from other third parties, or those that you have written yourself might need to be updated.

For more information about supporting instanced rendering in your shaders, see [GPU Instancing](GPUInstancing.md). The information in this section specifically talks about stereo rendering and might not include all changes you need to make to support instanced rendering in general.

## Update the vertex input attributes struct

Add the `UNITY_VERTEX_INPUT_INSTANCE_ID` macro to the `appdata` struct.

Example:

```
struct appdata
{
    float4 vertex : POSITION;
    float2 uv : TEXCOORD0;

    UNITY_VERTEX_INPUT_INSTANCE_ID //Insert
};
```

## Update the vertex output attributes struct

Add `UNITY_VERTEX_OUTPUT_STEREO` macro to the `v2f` output struct.

Example:

```
struct v2f
{
    float2 uv : TEXCOORD0;
    float4 vertex : SV_POSITION;

    UNITY_VERTEX_OUTPUT_STEREO //Insert
};
```

## Update the main vertex shader function

Add the following macros to the beginning of your main `vert` method (in order):

1. `UNITY_SETUP_INSTANCE_ID()`
2. `UNITY_INITIALIZE_OUTPUT(v2f, o)`
3. `UNITY_INITIALIZE_VERTEX_OUTPUT_STEREO()`

`UNITY_SETUP_INSTANCE_ID()` calculates and sets the built-in `unity_StereoEyeIndex` and `unity_InstanceID` shader variables to the correct values based on which eye the GPU is currently rendering.

`UNITY_INITIALIZE_VERTEX_OUTPUT_STEREO` tells the GPU which eye in the texture array it should render to, based on the value of `unity_StereoEyeIndex`. This macro also transfers the value of `unity_StereoEyeIndex` from the **vertex shader** so that it will be accessible in the fragment shader only if `UNITY_SETUP_STEREO_EYE_INDEX_POST_VERTEX` is called in the fragment shader `frag` method.

`UNITY_INITALIZE_OUTPUT(v2f,o)` initializes all `v2f` values to 0.

Example:

```
v2f vert (appdata v)
{
    v2f o;

    UNITY_SETUP_INSTANCE_ID(v); //Insert
    UNITY_INITIALIZE_OUTPUT(v2f, o); //Insert
    UNITY_INITIALIZE_VERTEX_OUTPUT_STEREO(o); //Insert

    o.vertex = UnityObjectToClipPos(v.vertex);

    o.uv = v.uv;

    return o;
}
```

## Post-Processing shaders

If you want your **Post-Processing** shaders to support single-pass stereo instancing, follow the steps for custom shaders as well as the steps below.

**Note:** You can download all Unity base shader **scripts** from the [Unity website](https://unity3d.com/get-unity/download/archive).

Do the following for each Post-Processing shader that you want to support single-pass instancing:

1. Add the `UNITY_DECLARE_SCREENSPACE_TEXTURE(tex)` macro outside the frag method (see the example below for placement) in your Shader script, so that when you use a particular stereo rendering method the GPU uses the appropriate texture sampler. For example, if you use multi-pass rendering, the GPU uses a texture 2D sampler. For single-pass instancing or multi-view rendering, the texture sampler is a texture array.
2. Add `UNITY_SETUP_STEREO_EYE_INDEX_POST_VERTEX(i)` at the beginning of the fragment shader frag method (See the example below for placement). You only need to add this macro if you want to use the `unity_StereoEyeIndex` built-in shader variable to find out which eye the GPU is rendering to. This is useful when testing post-processing effects.
3. Use the `UNITY_SAMPLE_SCREENSPACE_TEXTURE()` macro when sampling 2D textures (See the example below). Standard shaders use a 2D texture-based back buffer to sample textures. Single-pass stereo instancing does not use this type of back buffer, so if you do not specify a different method for 2D texture sampling, your shader does not render correctly. To prevent rendering issues, the `UNITY_SAMPLE_SCREENSPACE_TEXTURE()` macro detects which stereo **rendering path** you are using and then automatically samples the texture in the correct manner. See Unity documentation on [HLSLSupport.cginc](SL-BuiltinIncludes.md) to learn more about similar macros used for depth textures and screen-space shadow maps.

Example:

```
UNITY_DECLARE_SCREENSPACE_TEXTURE(_MainTex); //Insert

fixed4 frag (v2f i) : SV_Target
{
    UNITY_SETUP_STEREO_EYE_INDEX_POST_VERTEX(i); //Insert
    
    fixed4 col = UNITY_SAMPLE_SCREENSPACE_TEXTURE(_MainTex, i.uv); //Insert
    
    // just invert the colors
    
    col = 1 - col;
    
    return col;
}
```

## Full sample shader code

Below is a simple example of the template image effect shader with all of the previously mentioned changes applied to support single-pass stereo instancing. The lines added to the shader code are marked with the comment: `//Insert`.

```
struct appdata
{
    float4 vertex : POSITION;
    float2 uv : TEXCOORD0;
    
    UNITY_VERTEX_INPUT_INSTANCE_ID //Insert
};

//v2f output struct

struct v2f
{

    float2 uv : TEXCOORD0;
    float4 vertex : SV_POSITION;
    
    UNITY_VERTEX_OUTPUT_STEREO //Insert
};

v2f vert (appdata v)
{
    v2f o;
    
    UNITY_SETUP_INSTANCE_ID(v); //Insert
    UNITY_INITIALIZE_OUTPUT(v2f, o); //Insert
    UNITY_INITIALIZE_VERTEX_OUTPUT_STEREO(o); //Insert
    
    o.vertex = UnityObjectToClipPos(v.vertex);
    o.uv = v.uv;
    return o;
}

UNITY_DECLARE_SCREENSPACE_TEXTURE(_MainTex); //Insert

fixed4 frag (v2f i) : SV_Target
{
    UNITY_SETUP_STEREO_EYE_INDEX_POST_VERTEX(i); //Insert
    
    fixed4 col = UNITY_SAMPLE_SCREENSPACE_TEXTURE(_MainTex, i.uv); //Insert
    
    // invert the colors
    
    col = 1 - col;
    
    return col;
}
```

## Procedural geometry

If you use the [Graphics.DrawProceduralIndirect()](../ScriptReference/Graphics.DrawProcedural.md) and [CommandBuffer.DrawProceduralIndirect()](../ScriptReference/Graphics.DrawProceduralIndirect.md) methods to draw fully procedural geometry on the GPU, note that both methods receive their arguments from a compute buffer. This means that it is difficult to increase the instance count at run time. To increase the instance count, you must manually double the instance count contained in your compute buffers.

## Debugging your shader

The following shader code renders a **GameObject** as green for a user’s left eye and red for their right eye. This shader is useful for debugging your stereo rendering, because it allows you to verify that all stereo graphics work and are functioning correctly.

```
Shader "XR/StereoEyeIndexColor"
{
   Properties
   {
       _LeftEyeColor("Left Eye Color", COLOR) = (0,1,0,1)
       _RightEyeColor("Right Eye Color", COLOR) = (1,0,0,1)
   }

   SubShader
   {
      Tags { "RenderType" = "Opaque" }

      Pass
      {
         CGPROGRAM

         #pragma vertex vert
         #pragma fragment frag

         float4 _LeftEyeColor;
         float4 _RightEyeColor;

         #include "UnityCG.cginc"

         struct appdata
         {
            float4 vertex : POSITION;

            UNITY_VERTEX_INPUT_INSTANCE_ID
         };

         struct v2f
         {
            float4 vertex : SV_POSITION;

            UNITY_VERTEX_INPUT_INSTANCE_ID 
            UNITY_VERTEX_OUTPUT_STEREO
         };

         v2f vert (appdata v)
         {
            v2f o;

            UNITY_SETUP_INSTANCE_ID(v);
            UNITY_INITIALIZE_OUTPUT(v2f, o);
            UNITY_INITIALIZE_VERTEX_OUTPUT_STEREO(o);

            o.vertex = UnityObjectToClipPos(v.vertex);

            return o;
         }

         fixed4 frag (v2f i) : SV_Target
         {
            UNITY_SETUP_STEREO_EYE_INDEX_POST_VERTEX(i);

            return lerp(_LeftEyeColor, _RightEyeColor, unity_StereoEyeIndex);
         }
         ENDCG
      }
   }
}
```

### ShaderGraph debug shader

ShaderGraph automatically adds the macros required to support single-pass stereo rendering. To implement the debug shader in ShaderGraph you can use a [Custom Function Node](https://docs.unity3d.com/Packages/com.unity.shadergraph@10.10/manual/Custom-Function-Node.html) that sets the base color based on the eye index.

![A custom function node in ShaderGraph.](../uploads/Main/xr-debug-shadergraph.png)

A custom function node in ShaderGraph.

Use the `unity_StereoEyeIndex` shader attribute to determine the base color depending on which eye instance is being rendered. The Custom Function Node in the example above contains the following code:

```
Out = lerp(LeftColor, RightColor, unity_StereoEyeIndex);
```