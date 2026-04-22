# Get the current framebuffer with the GrabPass command in the Built-In Render Pipeline

GrabPass is a command that creates a special type of Pass that grabs the contents of the frame buffer into a texture. This texture can be used in subsequent Passes to do advanced image based effects.

This command can significantly increase both CPU and GPU frame times. You should generally avoid using this command other than for quick prototyping, and attempt to achieve your effect in other ways. If you do use this command, try to reduce the number of screen grabbing operations as much as possible; either by reducing your usage of this command, or by using the signature that grabs the screen to a named texture if applicable.

GrabPass works only on the frame buffer. You cannot use this command to grab the contents of other render targets, the **depth buffer**, and so on.

## Examples

This example for the Built-in **Render Pipeline** demonstrates using `GrabPass` to invert the colors of the render target. Note that this is not an efficient way to achieve this effect, and is intended only to demonstrate the use of GrabPass; the same effect could be achieved more efficiently using an invert blend mode.

```
Shader "GrabPassInvert"
{
    SubShader
    {
        // Draw after all opaque geometry
        Tags { "Queue" = "Transparent" }

        // Grab the screen behind the object into _BackgroundTexture
        GrabPass
        {
            "_BackgroundTexture"
        }

        // Render the object with the texture generated above, and invert the colors
        Pass
        {
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc"

            struct v2f
            {
                float4 grabPos : TEXCOORD0;
                float4 pos : SV_POSITION;
            };

            v2f vert(appdata_base v) {
                v2f o;
                // use UnityObjectToClipPos from UnityCG.cginc to calculate 
                // the clip-space of the vertex
                o.pos = UnityObjectToClipPos(v.vertex);

                // use ComputeGrabScreenPos function from UnityCG.cginc
                // to get the correct texture coordinate
                o.grabPos = ComputeGrabScreenPos(o.pos);
                return o;
            }

            sampler2D _BackgroundTexture;

            half4 frag(v2f i) : SV_Target
            {
                half4 bgcolor = tex2Dproj(_BackgroundTexture, i.grabPos);
                return 1 - bgcolor;
            }
            ENDHLSL
        }

    }
}
```

##

## Provide fragment color as input

Some GPUs (most notably PowerVR-based ones on iOS) allow you to do a form of programmable blending by providing current fragment color as input to the Fragment **Shader** (refer to `EXT_shader_framebuffer_fetch` on [khronos.org](https://www.khronos.org/registry/gles/extensions/EXT/EXT_shader_framebuffer_fetch.txt)). This process is sometimes called framebuffer fetch.

To do this, use the `inout` color argument when you write a fragment shader. For example:

```
HLSLPROGRAM
// only compile Shader for platforms that can potentially
// do it (currently gles,gles3,metal)
#pragma only_renderers framebufferfetch

void frag (v2f i, inout half4 ocol : SV_Target)
{
    // ocol can be read (current framebuffer color)
    // and written into (will change color to that one)
    // ...
}
ENDHLSL
```

## Additional resources

* [GrabPass directive in ShaderLab reference](SL-GrabPass.md)
* [Writing shaders in code](shader-writing.md)