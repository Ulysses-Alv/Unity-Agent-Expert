# Write HLSL for different graphics APIs

There are differences in how graphics rendering behaves between different graphics APIs. Most of the time the Unity Editor hides the differences, but there are some situations where Unity can’t do this for you.

## Check floating point exception handling

Desktop platforms and mobile platforms can handle exceptions differently during floating point operations such as division by zero. Exceptions can result in `NaN`, zero, or another value. Test your **shaders** on the different platforms you target.

## Use consistent semantics

Use the following semantics to ensure shaders work on all platforms:

* `SV_POSITION` instead of `POSITION` for the position the **vertex shader** outputs.
* `SV_Target` instead of `COLOR` or `COLOR0` for the output of the fragment shader.
* `PSIZE` if you render meshes as points. For example, set `PSIZE` to 1. This avoids some platforms like Metal reading point size as undefined.

For more information about semantics, refer to [High-Level Shader Language (HLSL) semantics reference](SL-HLSLSemantics.md).

## Use const for constants

Don’t use `const` for values that come from outside the shader or that you calculate at runtime, as this only works in HLSL, not OpenGL Shading Language (GLSL). Use `const` only for compile-time constants.

## Avoid mismatched buffer layouts

If you use the standard HLSL `cbuffer` type, or Unity’s `CBUFFERSTART` and `CBUFFEREND` macros, a `float3` might become a `float4`, or a `float` might become a `float2`.

To ensure all graphics APIs compile a buffer with the same data layout, do the following:

* Use `float4` and `float4x4` instead of `float3` and `float3x3`, because `float4` variables are the same size on all graphics APIs, while `float3` variables can become a different size on some graphics APIs.
* Declare variables in decreasing size order, for example `float4` then `float2` then `float`, so all graphics APIs structure the data in the same way.

For example:

```
cbuffer myConstantBuffer {
    float4x4 matWorld;
    float4 vObjectPosition; // Uses a float4 instead of a float3
    float arrayIndex;
}
```

**Note:** You can’t add structs to constant buffers.

## Handle reversed coordinate spaces

DirectX, Metal, and Vulkan use different coordinate systems to other graphics APIs. To make sure shaders work correctly on all platforms, adjust your code to handle reversed textures.

To check for reversed textures, use pre-processor conditionals like `#if` and `#ifdef` with Unity’s built-in macros or methods.

### Check for upside-down render textures

If you sample a **render texture** at `(0,0)`, DirectX, Metal, and Vulkan return the top-left of the texture. Other graphics APIs return the bottom-left of the texture. Unity handles this internally for most render textures by flipping render textures upside-down internally where needed, so coordinates you use return the same texels on all platforms.

To check if the current platform uses reversed coordinates for render textures, use either of the following methods:

* Check if Unity defines `UNITY_UV_STARTS_AT_TOP`, which means the graphics API uses `(0,0)` for the top-left of the texture.

  For example, to check for reversed render textures then flip vertical uv coordinates:

  ```
  #if UNITY_UV_STARTS_AT_TOP
      if (_MainTex_TexelSize.y < 0)
          uv.y = 1-uv.y;
  #endif
  ```
* Check the built-in `ProjectionParams.x` variable. If the value is `-1`, Unity flipped the render texture, so you must also flip vertical UV coordinates.

**Note**: If you use [GrabPass](SL-GrabPass.md) in the Built-In **Render Pipeline**, Unity might not reverse the texture. To calculate the coordinates to use to sample the texture, use the `ComputeGrabScreenPos` method. For more information, refer to [Use built-in shader functions in the Built-In Render Pipeline](SL-BuiltinFunctions.md).

### Check for reversed depth direction

The direction of depth in clip space, the **depth buffer**, and depth textures is different depending on the graphics API you target.

| **Coordinate space** | **Z coordinate in DirectX, Metal, and Vulkan** | **Z coordinate in OpenGL ES and **WebGL**** |
| --- | --- | --- |
| Depth buffer and depth textures | `1.0` near to `0.0` far | `0.0` near to `1.0` far |
| Clip space | `1.0` near to `0.0` far | `-1.0` near to `1.0` far |

Use the following in custom HLSL code to handle these differences:

* Use [SystemInfo.usesReversedZBuffer](../ScriptReference/SystemInfo-usesReversedZBuffer.md) to check if the graphics API uses 1.0 to 0.0 for depth.
* Check if Unity defines `UNITY_REVERSED_Z`, which means the graphics API uses 1.0 to 0.0 for depth, then flip the z value.

  For example, in the Built-In Render Pipeline:

  ```
  float z = tex2D(_CameraDepthTexture, uv);
  #if defined(UNITY_REVERSED_Z)
      z = 1.0f - z;
  #endif
  ```

You might also need to flip depth bias values if you use a code rendering **plug-in**.

For more macros that help with depth, refer to [Shader methods in the Built-In Render Pipeline](use-built-in-shader-methods-birp.md) or [Shader methods in the Universal Render Pipeline (URP)](urp/use-built-in-shader-methods.md).

## Compiling on DirectX platforms

If you target the DirectX graphics API, the compiler is stricter than other compilers. Check the following:

* Set all the components of a vector. For example, if you declare a `float4`, ensure you set all four components.
* Don’t use the `tex2D` method in the vertex shader. Instead, use the `tex2Dlod` method to sample a specific mip level, and use `#pragma target 3.0` to ensure compatibility.
* If you use a vertex shader with an `out` parameter in a **surface shader**, initialize the output using the `UNITY_INITIALIZE_OUTPUT` macro.

## Additional resources

* [Use 16-bit precision in shaders](SL-Use16BitPrecisionInShaders.md).
* [Platform development](PlatformSpecific.md)