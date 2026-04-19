# Camera output

Techniques for outputting depth, depth-normal and motion vector information from a **camera**.

| **Page** | **Description** |
| --- | --- |
| [Introduction to camera output](CameraOutput-introduction.md) | Learn about using the `DepthTextureMode` API to output a depth or motion texture from a camera. |
| [Output a depth texture from a camera](SL-CameraDepthTexture.md) | Output a texture with data about the object depth or surface normals in the camera view. |
| [Output a motion vector texture from a camera](SL-CameraDepthTexture-motionvectors.md) | Output a texture with data about object motion in the camera view. |
| [Sample output textures in a shader](CameraOutput-shader.md) | Read the data from a depth texture or a motion vector texture in custom **shader** code. |
| [Troubleshooting camera output](CameraOutput-troubleshoot.md) | Solve common issues with camera output, such as visual artefacts. |

## Additional resources

* [Motion vectors in URP](urp/features/motion-vectors.md)
* [Camera render order in URP](urp/cameras-advanced.md)
* [Access camera data with the Universal Additional Camera Data component in URP](urp/universal-additional-camera-data.md)