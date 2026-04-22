# Universal Render Pipeline compatibility in XR

Support for **XR** features in the Universal **Render Pipeline** (URP) varies by URP package version. This page details compatibility between XR features in Unity 6 and the latest compatible URP version.

To determine which version of URP is compatible with your current Unity version, refer to the [Requirements and compatibility](urp/requirements.md) page in the Universal Render Pipeline documentation.

Unity 6 supports the following **AR** and **VR** features in the Universal Render Pipeline:

| **Feature** | **Supported in XR** |
| --- | --- |
| **Post-processing** effects: Bloom | Yes |
| Post-processing effects: MotionBlur | Yes |
| Post-processing effects: Lens Distortion | No |
| Post-processing effects: **Depth of Field** | Yes |
| Post-processing effects: **ToneMapping** | Yes |
| Spatial Temporal Post-Processing | No |
| Other post-processing effects (color adjustment, etc.) | Yes |
| GI (Global Illumination) | Yes |
| **HDR** | Yes |
| MSAA | Yes |
| Physical **Camera** | No |
| CopyColor / ColorDepth | Yes |
| Multi Display | No |
| Camera Stacking | Yes |
| Cascaded Shadow | Yes |
| sRGB | Yes |
| **Skybox** | Yes |
| Fog | Yes |
| **Billboard** | Yes |
| **Shader** Graph | Yes (1) |
| **Particles** | Yes |
| **Terrain** | Yes |
| 2D UI (Canvas Renderer, Text Mesh Pro) | Yes |
| URP Debug (Scene View Mode, Frame Debug) | Yes (2) |

* (1) Although Shader Graph shaders can run in XR, Shader Graph doesn’t currently support the XR utility feature to create SPI-compatible shader input textures. Unity will expand support for Shader Graph functionality in future releases.
* (2) Unity supports frame debugging for mock HMDs. Currently, there is no support for Meta/Oculus.

In Unity 6.3 and newer, you can use on-tile post-processing to enable certain URP post-processing effects on XR devices. Refer to [On-tile post-processing (URP)](xr-graphics-on-tile-post-processing.md) to learn more about on-tile post-processing and how to enable it in your XR project.

To learn more about post-processing effects, refer to the [Effect list](post-processing-effect-availability-reference.md) page in the Universal Render Pipeline documentation.