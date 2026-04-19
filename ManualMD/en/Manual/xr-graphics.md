# XR graphics

Graphics and rendering in an **XR** app follow the same principles as they do in any Unity application, with some differences arising from the need for stereo rendering and increased efficiency.

XR apps generally require very high and consistent frame rates for user comfort. At the same time, rendering in stereo means that every visible object must be drawn twice from different perspectives. Techniques like single-pass rendering can improve rendering efficiency by reducing the duplication of effort, but also require changes to **shader** code.

| Topic | Description |
| --- | --- |
| [Universal Render Pipeline compatibility in XR](xr-render-pipeline-compatibility.md) | Understand which Universal **Render Pipeline** features are compatible with XR platforms. |
| [Stereo rendering](SinglePassStereoRendering.md) | Understand stereo rendering for XR platforms. |
| [Foveated rendering](xr-foveated-rendering.md) | Learn about foveated rendering for XR. |
| [Multiview Render Regions](xr-multiview-render-regions.md) | Understand the Multiview Render Regions feature for XR devices. |
| [VR frame timing](VRFrameTiming.md) | Learn about **VR** frame timing. |
| [Optimize for untethered XR devices in URP](xr-untethered-device-optimization.md) | Understand the ways you can optimize your URP project for XR devices. |
| [Resolution control in XR projects](xr-graphics-resolution-scaling.md) | Understand how to control the resolution in your untethered XR project. |
| [Tile-based rendering in XR](xr-graphics-on-tile-rendering.md) | Understand on-tile rendering in XR. |
| [On-tile post-processing (URP)](xr-graphics-on-tile-post-processing.md) | Learn about on-tile **post-processing** for untethered XR devices. |