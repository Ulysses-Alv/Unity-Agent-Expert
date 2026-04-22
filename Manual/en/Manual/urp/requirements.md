# Requirements and compatibility for URP

This page contains information on system requirements and compatibility of the Universal Render Pipeline (URP) package.

## Unity Editor compatibility

URP is a [core Unity package](../pack-core). For each alpha, beta or patch release of Unity, the main Unity installer contains the up-to-date version of the package.

The Package Manager window displays only the major and minor revision of the package. For example, version 17.2.0 for all Unity 6.2.x releases.

You can install a different version of a graphics package from disk using the Package Manager, or by modifying the `manifest.json` file.

## Render pipeline compatibility

Projects made using URP are not compatible with the High Definition Render Pipeline (HDRP) or the Built-in Render Pipeline. Before you start development, you must decide which render pipeline to use in your Project. For information on choosing a render pipeline, refer to the [Render Pipelines](../render-pipelines.md) page.

## Graphics API compatibility

URP supports the following graphics APIs:

* DirectX 11 (feature level 11\_0 and later)
* DirectX 12
* Vulkan
* Metal
* OpenGL ES 3.0 and later
* OpenGL Core
* WebGL2
* WebGPU (experimental)

## Unity Player system requirements

This package does not add any extra platform-specific requirements. General system requirements for the Unity Player apply. For more information on Unity system requirements, check the [System requirements for Unity](../system-requirements.md).

## Additional resources

* [Render pipeline feature comparison reference](../render-pipelines-feature-comparison.md)
* [Hardware requirements for the Built-In Render Pipeline](../RenderTech-HardwareRequirements.md)