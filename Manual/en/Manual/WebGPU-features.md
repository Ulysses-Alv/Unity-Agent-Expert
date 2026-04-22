# Introduction to the WebGPU graphics API

Enhance your Web project with WebGPU’s advanced graphics capabilities.

WebGPU is a web graphics API that provides low-level, high-performance access to modern GPU features. This information covers the features WebGPU uniquely supports and some features it doesn’t support yet.

## The features of WebGPU

In the browser, WebGPU uses modern GPU frameworks such as DirectX 12, Vulkan, or Metal (depending on your device) to improve visuals and optimize performance. Access to these frameworks let WebGPU introduce advanced graphics capabilities that aren’t possible with [WebGL2](WebGL2.md), including:

* [Compute shaders](class-ComputeShader.md)
* [Indirect rendering](../ScriptReference/Graphics.RenderMeshIndirect.md)
* [GPU skinning](../ScriptReference/PlayerSettings-gpuSkinning.md)
* [VFX Graph](https://unity.com/visual-effect-graph)

## Features with no support in WebGPU

Although WebGPU unlocks useful graphics features, there are a few it doesn’t support yet:

* Async compute
* [Dynamic resolution](DynamicResolution-landing.md)
* [Cubemap arrays](class-CubemapArray-create.md)

WebGPU also has some limitations you need to be aware of. For more information, refer to [Limitations of WebGPU](WebGPU-limitations.md).

## Additional resources

* [WebGPU (Experimental)](WebGPU.md)
* [Limitations of the WebGPU graphics API](WebGPU-limitations.md)
* [Recommendations for both WebGL2 and WebGPU APIs](web-graphics-apis-intro.md)
* [Enable the WebGPU graphics API](WebGPU-enable.md)