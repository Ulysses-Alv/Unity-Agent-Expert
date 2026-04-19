# Built-in shader variables reference

Unity’s built-in include files contain global variables for your **shaders**: things like current object’s transformation matrices, light parameters, current time and so on. You use them in [shader programs](writing-shader-writing-shader-programs-hlsl.md) like any other variable, but if you include the relevant include file, you don’t have to declare them.

For more information on include files, see [Built-in include files](SL-BuiltinIncludes.md).

## Vertex input structures

| **Structure** | **Description** |
| --- | --- |
| `appdata_base` | Declares `vertex` for position, `normal`, and `texcoord`. |
| `appdata_tan` | Declares `vertex` for position, `tangent`, `normal`, and `texcoord`. |
| `appdata_full` | Declares `vertex` for position, `tangent`, `normal`, `color`, `texcoord`, `texcoord1`, `texcoord2`, and `texcoord3`. |
| `appdata_img` | Declares `vertex` for position, and `texcoord`. |

For more information, refer to [Input data into a custom shader](SL-VertexProgramInputs.md).

## Inline sampler suffixes

Add the following suffixes when you declare an inline `SamplerState` sampler in HLSL code, to manually set texture filtering and wrapping modes. Separate multiple words with an underscore (`_`). For example, to set a sampler to use linear filtering and repeat wrapping mode, name it `sampler_linear_repeat`.

**Note**: These sampler suffixes only work if you target the DX11, DX12, Metal or Vulkan graphics API.

You must include a single filtering mode and a single texture wrapping mode. The other suffixes are optional.

| **Suffix** | **Description** |
| --- | --- |
| `point` | Enables nearest texture filtering. |
| `linear` | Enables bilinear texture filtering. |
| `trilinear` | Enables trilinear texture filtering. |
| `clamp` | Sets the texture wrapping mode to clamp. |
| `repeat` | Sets the texture wrapping mode to repeat. |
| `mirror` | Sets the texture wrapping mode to mirror. |
| `mirroronce` | Sets the texture wrapping mode to mirror once. If you target a mobile platform that doesn’t support `mirroronce`, Unity falls back to `mirror`. |
| `compare` | Sets the sampler to compare depth values. For more information refer to [SampleCmp](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-to-samplecmp) and [SampleCmpLevelZero](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-to-samplecmplevelzero) in the Microsoft HLSL documentation. |
| `aniso2`, `aniso4`, `aniso8`, `aniso16` | Enables anisotropic filtering with 2x, 4x, 8x, or 16x sampling. Depending on the graphics API and platform you target, Unity might clamp the filtering to a lower value, or disable anisotropic filtering completely. |

Add `U`, `V` or `W` to a wrap mode to set the wrapping mode for that axis only. For example, `sampler_clampu_point` sets the U axis to clamp mode, and leaves the V and W axes at their default wrapping mode.

For more information, refer to:

* [`Texture.wrapMode`](../ScriptReference/Texture-wrapMode.md)
* [`Texture.anisoLevel`](../ScriptReference/Texture-anisoLevel.md)
* [`Texture.filterMode`](../ScriptReference/Texture-filterMode.md)

## Transformations

All these matrices are `float4x4` type, and are column major.

| **Name** | **Value** |
| --- | --- |
| UNITY\_MATRIX\_MVP | Current model \* view \* projection matrix. |
| UNITY\_MATRIX\_MV | Current model \* view matrix. |
| UNITY\_MATRIX\_V | Current view matrix. |
| UNITY\_MATRIX\_P | Current projection matrix. |
| UNITY\_MATRIX\_VP | Current view \* projection matrix. |
| UNITY\_MATRIX\_T\_MV | Transpose of model \* view matrix. |
| UNITY\_MATRIX\_IT\_MV | Inverse transpose of model \* view matrix. |
| unity\_ObjectToWorld | Current model matrix. |
| unity\_WorldToObject | Inverse of current world matrix. |

## Camera and screen

These variables will correspond to the [Camera](class-Camera.md) that is rendering. For example during shadowmap rendering, they will still refer to the Camera component values, and not the “virtual camera” that is used for the shadowmap projection.

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| \_WorldSpaceCameraPos | float3 | World space position of the camera. |
| \_ProjectionParams | float4 | `x` is 1.0 (or –1.0 if currently rendering with a [flipped projection matrix](SL-PlatformDifferences.md)), `y` is the camera’s near plane, `z` is the camera’s far plane and `w` is 1/FarPlane. |
| \_ScreenParams | float4 | `x` is the width of the camera’s target texture in **pixels**, `y` is the height of the camera’s target texture in pixels, `z` is 1.0 + 1.0/width and `w` is 1.0 + 1.0/height. |
| \_ZBufferParams | float4 | Used to linearize Z buffer values.  * `x`: `1 - far/near`, or `–1 + far/near` if `UNITY_REVERSED_Z` is set to 1.  For more information about `UNITY_REVERSED_Z`, refer to [Branch based on platform features](shader-branching-platform). * `y`: `far/near`, or `1` if `UNITY_REVERSED_Z` is set to 1.  For more information about `UNITY_REVERSED_Z`, refer to [Branch based on platform features](shader-branching-platform). * `z`: `x / far` * `w`: `y / far` |
| unity\_OrthoParams | float4 | The following parameters:  * `x`: half the width of the camera view in world space * `y`: half the height of the camera view in world space * `z`: unused * `w`: 1.0 if the camera is orthographic, or 0.0 if the camera is perspective |
| unity\_CameraProjection | float4x4 | Camera’s projection matrix. |
| unity\_CameraInvProjection | float4x4 | Inverse of camera’s projection matrix. |
| unity\_CameraWorldClipPlanes[6] | float4 | Camera frustum plane world space equations, in this order: left, right, bottom, top, near, far. |

## Time

Time is measured in seconds, and is scaled by the Time multiplier in your Project’s [Time settings](class-TimeManager.md). There is no built-in variable that provides access to unscaled time.

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| \_Time | float4 | Time since level load (t/20, t, t\*2, t\*3), use to animate things inside the shaders. |
| \_SinTime | float4 | Sine of time: (t/8, t/4, t/2, t). |
| \_CosTime | float4 | Cosine of time: (t/8, t/4, t/2, t). |
| unity\_DeltaTime | float4 | Delta time: (dt, 1/dt, smoothDt, 1/smoothDt). |

## Lighting

Light parameters are passed to shaders in different ways depending on which [Rendering Path](RenderingPaths.md) is used,
and which LightMode [Pass Tag](SL-PassTags.md) is used in the shader.

[Forward rendering](RenderTech-ForwardRendering.md) (`ForwardBase` and `ForwardAdd` pass types):

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| \_LightColor0 *(declared in UnityLightingCommon.cginc)* | fixed4 | Light color. |
| \_WorldSpaceLightPos0 | float4 | Directional lights: (world space direction, 0). Other lights: (world space position, 1). |
| unity\_WorldToLight *(declared in AutoLight.cginc)* | float4x4 | World-to-light matrix. Used to sample cookie & attenuation textures. |
| unity\_4LightPosX0, unity\_4LightPosY0, unity\_4LightPosZ0 | float4 | *(ForwardBase pass only)* world space positions of first four non-important point lights. |
| unity\_4LightAtten0 | float4 | *(ForwardBase pass only)* attenuation factors of first four non-important point lights. |
| unity\_LightColor | half4[4] | *(ForwardBase pass only)* colors of of first four non-important point lights. |
| unity\_WorldToShadow | float4x4[4] | World-to-shadow matrices. One matrix for Spot Lights, up to four for directional light cascades. |

Deferred shading, used in the lighting pass shader (all declared in UnityDeferredLibrary.cginc):

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| \_LightColor | float4 | Light color. |
| unity\_WorldToLight | float4x4 | World-to-light matrix. Used to sample cookie & attenuation textures. |
| unity\_WorldToShadow | float4x4[4] | World-to-shadow matrices. One matrix for Spot Lights, up to four for directional light cascades. |

Spherical harmonics coefficients (used by ambient and light probes) are set up for `ForwardBase` and `Deferred`
pass types. They contain 3rd order SH to be evaluated by world space normal (see `ShadeSH9` from [UnityCG.cginc](SL-BuiltinIncludes.md)).
The variables are all half4 type, `unity_SHAr` and similar names.

[Vertex-lit rendering](RenderTech-VertexLit.md) (`Vertex` pass type):

Up to 8 lights are set up for a `Vertex` pass type; always sorted starting from the brightest one. So if you want
to render objects affected by two lights at once, you can just take first two entries in the arrays. If there are fewer than eight
lights affecting the object, the rest will have their color set to black.

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| unity\_LightColor | half4[8] | Light colors. |
| unity\_LightPosition | float4[8] | View-space light positions. (-direction,0) for directional lights; (position,1) for Point or Spot Lights. |
| unity\_LightAtten | half4[8] | Light attenuation factors. *x* is cos(spotAngle/2) or –1 for non-Spot Lights; *y* is 1/cos(spotAngle/4) or 1 for non-Spot Lights; *z* is quadratic attenuation; *w* is squared light range. |
| unity\_SpotDirection | float4[8] | View-space Spot Lights positions; (0,0,1,0) for non-Spot Lights. |

## Lightmaps

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| unity\_Lightmap | Texture2D | Contains **lightmap** information. |
| unity\_LightmapST | float4[8] | Scales and translates the UV information to the correct range to sample the lightmap texture. |

## Fog and Ambient

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| unity\_AmbientSky | fixed4 | Sky ambient lighting color in gradient ambient lighting case. |
| unity\_AmbientEquator | fixed4 | Equator ambient lighting color in gradient ambient lighting case. |
| unity\_AmbientGround | fixed4 | Ground ambient lighting color in gradient ambient lighting case. |
| unity\_IndirectSpecColor | fixed4 | If you use a [skybox](skyboxes-using.md), this is the average color of the skybox, which Unity calculates using the `DC` component of the [spherical harmonics data](LightProbes-TechnicalInformation.md) in the [ambient probe](../ScriptReference/RenderSettings-ambientProbe.md). Otherwise this is black. |
| UNITY\_LIGHTMODEL\_AMBIENT | fixed4 | Ambient lighting color (sky color in gradient ambient case). Legacy variable. |
| unity\_FogColor | fixed4 | Fog color. |
| unity\_FogParams | float4 | Parameters for fog calculation: (density / sqrt(ln(2)), density / ln(2), –1/(end-start), end/(end-start)). *x* is useful for Exp2 fog mode, *y* for Exp mode, *z* and *w* for Linear mode. |

## Various

| **Name** | **Type** | **Value** |
| --- | --- | --- |
| unity\_LODFade | float4 | Level-of-detail fade when using [LODGroup](class-LODGroup.md). *x* is fade (0..1), *y* is fade quantized to 16 levels, *z* and *w* unused. |
| \_TextureSampleAdd | float4 | Set automatically by Unity **for UI only** based on whether the texture being used is in Alpha8 format (the value is set to (1,1,1,0)) or not (the value is set to (0,0,0,0)). |