# DirectX

Unity supports the DirectX graphics API including both DirectX 11 and DirectX 12. However, not all features are available in DirectX 11. For more information, refer to [Feature comparison of DirectX 11 and DirectX 12 in Unity](#comparison-of-directx11-and-directx12-in-unity).

## Set your default graphics API to DirectX

You can choose to set DirectX 11 (DX11) or DirectX 12 (DX12) as your default Graphics API in the Editor or Standalone Player:

1. Open the [Player settings](class-PlayerSettings.md) (menu: **Edit** > **Project Settings** > **Player**).
2. In the **Other Settings** > **Rendering** section, disable the **Auto Graphics API for a platform (Windows/Mac/Linux)** option.
3. Select the **Add** (**+**) button, then select **Direct3D11** or **Direct3D12** from the list of the supported Graphics APIs.

## Feature comparison of DirectX 11 and DirectX 12 in Unity

The following lists include the features introduced with the DirectX 12 graphics API, which are unavailable in DirectX 11.

| **APIs** | **DirectX 11** | **DirectX 12** |
| --- | --- | --- |
| [Dynamic resolution](DynamicResolution-introduction.md) | Unsupported | Supported |
| [Asynchronous compute](../ScriptReference/Graphics.ExecuteCommandBufferAsync.md) | Unsupported | Supported |
| [Native render pass](../ScriptReference/Rendering.ScriptableRenderContext.BeginRenderPass.md) | Unsupported | Supported |
| [Ray tracing acceleration](../ScriptReference/Rendering.RayTracingAccelerationStructure.md) | Unsupported | Supported |
| [Graphics state collection](../ScriptReference/Experimental.Rendering.GraphicsStateCollection.md) | Unsupported | Supported |
| [XR foveated rendering](xr-foveated-rendering.md) | Unsupported | Supported |

| **Render Threading Mode** | **DirectX 11** | **DirectX 12** |
| --- | --- | --- |
| [Direct](../ScriptReference/Rendering.RenderingThreadingMode.Direct.md) | Supported | Supported |
| [Single threaded](../ScriptReference/Rendering.RenderingThreadingMode.SingleThreaded.md) | Supported | Supported |
| [Main + render thread](../ScriptReference/Rendering.RenderingThreadingMode.MultiThreaded.md) | Supported | Supported |
| [Legacy jobified](../ScriptReference/Rendering.RenderingThreadingMode.LegacyJobified.md) | Supported | Supported |
| [Native graphics jobs](../ScriptReference/Rendering.RenderingThreadingMode.NativeGraphicsJobs.md) | Unsupported | Supported |
| [Split graphics jobs](../ScriptReference/Rendering.RenderingThreadingMode.NativeGraphicsJobsSplitThreading.md) | Unsupported | Supported |

| ****Shader** features** | **DirectX 11** | **DirectX 12** |
| --- | --- | --- |
| [Ray tracing shader](../ScriptReference/Rendering.RayTracingShader.md) | Unsupported | Supported |
| [Inline ray tracing](../ScriptReference/ComputeShader.SetRayTracingAccelerationStructure.md) (`#pragma multi_compile _ UNITY_DEVICE_SUPPORTS_INLINE_RAY_TRACING`) | Unsupported | Supported |
| [Native 16-bit](SL-ShaderCompileTargets.md) (`#pragma multi_compile _ UNITY_DEVICE_SUPPORTS_NATIVE_16BIT`) | Unsupported | Supported |
| [Wave function](SL-ShaderCompileTargets.md) (`#pragma multi_compile _ UNITY_DEVICE_SUPPORTS_WAVE_ANY`) | Unsupported | Supported |

| **Universal **Render Pipeline**** | **DirectX 11** | **DirectX 12** |
| --- | --- | --- |
| [Raster pass merging](urp/render-graph-introduction.md) | Unsupported | Supported (Windows on ARM) |
| [Native RenderPass](urp/urp-universal-renderer.md#native-renderpass) | Unsupported | Supported (Windows on ARM) |
| [Framebuffer fetch](urp/render-graph-framebuffer-fetch.md) | Unsupported | Supported (Windows on ARM) |

| **High Definition Render Pipeline** | **DirectX 11** | **DirectX 12** |
| --- | --- | --- |
| [Hardware dynamic resolution](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Dynamic-Resolution.html) | Unsupported | Supported |
| [Asynchronous Compute Shaders](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/frame-settings-reference.html#asynchronous-compute-shaders) | Unsupported | Supported |
| [Ray traced ambient occlusion](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Ray-Traced-Ambient-Occlusion.html) | Unsupported | Supported |
| [Ray traced contact shadows](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Ray-Traced-Contact-Shadows.html) | Unsupported | Supported |
| [Ray traced global illumination](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Ray-Traced-Global-Illumination.html) | Unsupported | Supported |
| [Ray traced reflections](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Ray-Traced-Reflections.html) | Unsupported | Supported |
| [Ray traced shadows](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Ray-Traced-Shadows.html) | Unsupported | Supported |
| [Ray traced recursion](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Ray-Tracing-Recursive-Rendering.html) | Unsupported | Supported |
| [Ray traced subsurface scattering](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Ray-Traced-Subsurface-Scattering.html) | Unsupported | Supported |

## Additional resources

* [Configure graphics APIs](configure-graphicsAPIs.md)
* [Surface Shaders with DX11 / OpenGL Core Tessellation](SL-SurfaceShaderTessellation.md)
* [Surface shaders and DirectX 11 HLSL syntax](SL-SurfaceShaders.md#surface-shaders-directX)