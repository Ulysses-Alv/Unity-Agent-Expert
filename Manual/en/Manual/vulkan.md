# Vulkan

[Vulkan](https://www.vulkan.org) is a modern graphics API that works across multiple platforms including mobile and desktop. Unity supports Vulkan allowing you to target Vulkan-capable devices and providing direct control over the GPU.

## Platform compatibility

Unity supports the Vulkan API for developing applications on the following platforms:

* Android
* Embedded Linux
* Linux
* Windows

For developing [XR](XR.md) applications, Unity supports Vulkan on the Windows and Android platform.

## Vulkan on Android

When you create a new Android project, Vulkan is enabled as the default graphics API. For other platforms where Vulkan is not the default graphics API, you can enable it through Unity Editor’s **Player settings**. For more information, refer to [Override default graphics APIs](configure-graphicsAPIs.md#override-default-gfxapi).

When developing a Unity Android project with Vulkan, consider the following recommendations and features:

### Vulkan device filtering

You can manage Vulkan support on Android devices to fine-tune which devices use the Vulkan API with the [Vulkan Device Filtering Asset](introduction-vulkan-device-filtering-asset.md). This asset allows you to exclude specific Android devices on which Vulkan might lead to suboptimal performance.

### Swapchain pre-rotation

It is recommended to enable framebuffer orientation through the Android Player settings. For more information, refer to the documentation on [Framebuffer orientation](vulkan-swapchain-pre-rotation.md).

### Vulkan validation layers

You can enable validation layers for debugging native Vulkan API **plug-ins** through the `-force-vulkan-layers` command-line argument. For more information about the plug-in requirements when using validation layers, refer to [Plug-in using the Vulkan graphics API](low-level-native-plugin-rendering-extensions.md#vulkan-plugin).

## Advantages and limitations

Consider the following advantages and limitations of using Vulkan as the graphics API in your Unity project.

* Vulkan enables [High Definition Render Pipeline (HDRP)](high-definition-render-pipeline.md) support on Linux.
* On Linux and Android, Vulkan generally supports more capabilities, such as graphics jobs and direct control over the GPU, which aren’t available in [OpenGL Core](OpenGLCoreDetails.md) on Linux or OpenGL ES on Android.
* Unity doesn’t support Vulkan for Windows on Arm devices. If your application targets such devices, consider using alternative graphics APIs supported by Unity.
* On Windows, some features such as async compute and **ray tracing** are currently not implemented for Vulkan. For these features, [D3D12](UsingDX11GL3Features.md) is recommended in Unity.

## Key features and differences in Vulkan API

Vulkan supports several features and capabilities that differentiate it from other graphics APIs, such as OpenGL.

* [Graphics jobs](class-PlayerSettingsAndroid.md#GraphicsJobsMode): Vulkan supports graphics jobs, a feature that allows rendering tasks to run on multiple threads. This might help improve performance of your application. Graphics jobs feature isn’t supported with OpenGL.
* [Dynamic Resolution](DynamicResolution-introduction.md): Vulkan supports dynamic resolution scaling on devices allowing to adjust the resolution of the individual render targets in real time. This feature reduces workload on the GPU thereby improving performance of your application.
* [Memoryless Render Textures](../ScriptReference/RenderTexture-memorylessMode.md): On some Vulkan compatible platforms, memoryless **render textures** temporarily store data in the on-tile memory during rendering. This can help improve performance and reduce memory usage of your application.
* [Direct control over GPU](PlayerCommandLineArguments.md): Unity uses Vulkan’s direct GPU control capability to support the `-force-device-index` option, allowing you to specify which GPU Unity’s Standalone Player should use. This capability of Vulkan allows you to optimize your application for better performance. This feature isn’t supported with OpenGL.
* [XR features](xr-untethered-device-optimization.md#vulkan): Vulkan is recommended for Android XR and Meta Quest devices because some newer graphics features are only supported with Vulkan. For example, [Multiview Render Regions](xr-multiview-render-regions.md) is a Vulkan-specific feature. Vulkan API is more stable and provides better performance compared to OpenGL in URP projects targeting XR platforms.

## Additional resources

* [Vulkan design guidelines](https://developer.android.com/ndk/guides/graphics/design-notes?authuser=3)
* [Vulkan validation layers on Android](https://developer.android.com/ndk/guides/graphics/validation-layer)