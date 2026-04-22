# Metal requirements and compatibility

This page lists the requirements for using Metal as well as the features that Metal is compatible with.

## Platform compatibility

Unity supports Metal for the Unity Player on iOS, tvOS, and macOS. Unity also supports Metal for the Unity Editor on macOS.

## Hardware compatibility

Unity supports Metal for all Apple devices that Unity supports.

## Render pipeline compatibility

| **Feature** | **Built-in **Render Pipeline**** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom Scriptable Render Pipeline (SRP)** |
| --- | --- | --- | --- | --- |
| **Metal** | Yes | Yes | Yes (macOS only) | Yes |

## Shader compatibility

* Metal supports **shaders** that have a minimum [shader compilation target](SL-ShaderCompileTargets.md) of 3.5.
* Metal doesn’t support geometry shaders.

## Additional resources

* [Use 16-bit precision in shaders](SL-Use16BitPrecisionInShaders.md)
* [Writing shaders for different graphics APIs](SL-PlatformDifferences.md)