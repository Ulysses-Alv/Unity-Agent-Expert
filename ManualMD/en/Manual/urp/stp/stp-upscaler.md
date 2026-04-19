# Introduction to Spatial-Temporal Post-processing in URP

Spatial-Temporal **Post-Processing** (STP) uses spatial and temporal upsampling techniques to produce a high quality, anti-aliased image.

STP is a software-based upscaler.

## Requirements

STP uses compute **shaders**, so target platforms must support [Shader Model 5.0](https://learn.microsoft.com/en-us/windows/win32/direct3dhlsl/d3d11-graphics-reference-sm5).

STP doesn’t support OpenGL ES, even if the platform supports compute shaders.

STP requires [temporal anti-aliasing (TAA)](../anti-aliasing.md) pre-processing, it will implicitly enable it if not selected already.

## STP performance

STP configures itself automatically to provide the best balance of performance and quality based on the platform your application runs on. You don’t need to configure it for each different platform.

On high-performance platforms, like PCs and consoles, STP uses higher quality image filtering logic and additional deringing logic to improve image quality when it upscales images. These techniques require additional processing power and Unity uses them on high-performance devices where the performance impact is not significant.

On mobile devices, STP uses more performant image filtering logic to provide a balance between performance and image quality. This minimizes the performance impact of STP on less powerful devices, while still delivering a high quality image.

## Additional resources

* [Enable Spatial Temporal Post-processing](stp-enable.md)
* [Spatial Temporal Post-processing debug views](stp-debug-views.md)