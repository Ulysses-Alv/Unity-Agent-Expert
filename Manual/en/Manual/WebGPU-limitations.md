# Limitations of the WebGPU graphics API

The [WebGPU](WebGPU.md) graphics API provides useful graphics features. But it also has some limitations you need to be aware of:

* [Limited browser support](#limited-browser-support)
* [Compute shader limitations](#compute-shader-limitations)
* [GPU readback](#gpu-readback)
* [Texture format limitations](#texture-format-limitations)
* [Render texture channels](#render-texture-channels)
* [Terrain texture limits](#terrain-texture-limits)

For further best practise recommendations, refer to [Recommendations for both WebGL2 and WebGPU APIs](web-graphics-apis-intro.md).

## Limited browser support

WebGPU is currently experimental and not supported by all browsers and devices. For more information, refer to the MDN Web docs on [Accessing a device](https://developer.mozilla.org/en-US/docs/Web/API/WebGPU_API#accessing_a_device) and [Browser compatibility](https://developer.mozilla.org/en-US/docs/Web/API/WebGPU_API#browser_compatibility).

## Compute shader limitations

There are a few limitations to compute **shader** in Web.

### No RWBuffer support

You can use structured buffers, `RWStructuredBuffer`, but not `RWBuffer` in compute shaders.
Limited read-write storage texture formats
WebGPU only supports a limited set of **texture formats** for read-write storage textures in compute shaders. HLSL doesn’t explicitly state the access type of a storage texture, so it’s based on usage.

```
RWTexture2D<float> tex; // storage texture
tex[uv] = 1.0f; // write only
tex[uv] += 1.0f; // read-write
```

Write-only storage textures are generally much better supported. Refer to [Texture format limitations](#texture-format-limitations).

### Strict uniformity analysis

With WebGPU, you must only call barrier functions from non-uniform blocks. Otherwise, WebGPU will fail to compile the shader. Non-uniform blocks are shader code blocks that aren’t uniformly consistent across multiple threads of execution.

The following are HLSL barrier functions:

* `GroupMemoryBarrier`
* `GroupMemoryBarrirWithGroupSync`
* `DeviceMemoryBarrier`
* `DeviceMemoryBarrierWithGroupSync`
* `AllMemoryBarrier`
* `AllMemoryBarrierWithGroupSync`

### Async compute

WebGPU doesn’t support async compute.

### No HLSL extended features

WebGPU doesn’t provide support for the following features:

* Wave Intrinsics (WaveBasic, WaveVote, WaveBallot, WaveMath, WaveMultiPrefix)
* QuadShuffle
* Int64
* Native16Bit

## GPU readback

WebGPU doesn’t support synchronous readback of buffer or texture data from the GPU to the CPU.
The following functions will not work:

* [Texture2D.GetPixels](https://docs.unity3d.com/ScriptReference/Texture2D.GetPixels.html)
* [ComputeBuffer.GetData](https://docs.unity3d.com/ScriptReference/ComputeBuffer.GetData.html)
* [ScreenCapture.CaptureScreenshot](https://docs.unity3d.com/ScriptReference/ScreenCapture.CaptureScreenshot.html)
* [ScreenCapture.CaptureScreenshotAsTexture](https://docs.unity3d.com/ScriptReference/ScreenCapture.CaptureScreenshotAsTexture.html)

Instead, use the `AsyncGPUReadback` API to read buffer or texture data from the GPU to the CPU. For screen captures, use `ScreenCapture.CaptureScreenshotIntoRenderTexture` along with `AsyncGPUReadback`.

## Texture format limitations

For information about WebGPU texture format capabilities, refer to the W3 documentation on [WebGPU](https://www.w3.org/TR/webgpu). Some formats are more restrictive than on other APIs, especially for storage texture usage.

For example, WebGPU doesn’t support `RGBA8` or `RHalf` as a read-write storage texture format. If you do create a RenderTexture with that format, you must set `enableRandomWrite` to `false`, which is the default value.

However, WebGPU supports `RFloat` as a storage texture format, so if you create a RenderTexture with `RFloat` format, `enableRandomWrite` doesn’t need to be `false`.

You can use [SystemInfo.GetCompatibleFormat](https://docs.unity3d.com/ScriptReference/SystemInfo.GetCompatibleFormat.html) to find a format that will work for a given usage.

## Render texture channels

The output type of a fragment shader must have a greater or equal number of channels as the target **render texture** format.

If the RenderTexture is RGBA, it has 4 channels. The fragment shader must output a `float4` or `half4`. If it returns a float or half, WebGPU will encounter an error because there are channels of the render texture that will be uninitialized.

If a RenderTexture is R8, it has 1 channel. The fragment shader can then output a float or half. The fragment shader can also return a float4 or half4, and the extra output channels will be ignored.

## Terrain texture limits

The number of textures you can access from a shader is limited to 16 on many platforms. This limit can hit by the **terrain** renderer, especially with layer textures. This can make terrain rendering problematic for WebGPU. It is recommended that you limit the number of terrain texture layers you use.

## Additional resources

* [WebGPU (Experimental)](WebGPU.md)
* [Introduction to the WebGPU graphics API](WebGPU-features.md)
* [Recommendations for both WebGL2 and WebGPU APIs](web-graphics-apis-intro.md)
* [Enable the WebGPU graphics API](WebGPU-enable.md)
* [WebGL2](WebGL2.md)