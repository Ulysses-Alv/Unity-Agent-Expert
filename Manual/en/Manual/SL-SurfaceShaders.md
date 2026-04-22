# Introduction to surface shaders in the Built-In Render Pipeline

In the Built-in **Render Pipeline**, Surface **Shaders** are a streamlined way of writing shaders that interact with lighting.

## Render pipeline compatibility

| **Feature name** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| ****Surface Shaders**** | No | No | No | Yes |

To create custom lighting in URP and HDRP, do one the following instead:

* For a streamlined way of creating **Shader objects**, refer to [Shader Graph](shader-graph.md).
* Copy and modify the shader code URP or HDRP provides in the package files. Modifying might be difficult because the shader code is large and complex. For more information about finding package files, refer to [Inspecting packages](upm-inspect.md).

## Overview

Writing shaders that interact with lighting is complex. There are different light types, different shadow options, different **rendering paths** (forward and deferred rendering), and the shader should somehow handle all that complexity.

Surface Shaders is a code generation approach that makes it much easier to write lit shaders than using low level [vertex/pixel shader programs](writing-shader-writing-shader-programs-hlsl.md).

For some examples, take a look at [Surface Shader Examples](SL-SurfaceShaderExamples.md).

## How it works

You define a “surface function” that takes any UVs or data you need as input, and fills in output structure `SurfaceOutput`. SurfaceOutput basically describes *properties of the surface* (its albedo color, normal, emission, specularity etc.). You write this code in HLSL.

Surface Shader compiler then figures out what inputs are needed, what outputs are filled and so on, and generates actual [vertex&pixel shaders](writing-shader-writing-shader-programs-hlsl.md), as well as rendering passes to handle forward and deferred rendering.

## Surface Shaders and DirectX 11 HLSL syntax

Currently some parts of surface shader compilation pipeline do not understand [DirectX 11](UsingDX11GL3Features.md)-specific HLSL syntax, so if you’re using HLSL features like StructuredBuffers, RWTextures and other non-DX9 syntax, you have to wrap it into a DX11-only preprocessor macro.

See [Platform Specific Differences](SL-PlatformDifferences.md) and [Using HLSL](writing-shader-writing-shader-programs-hlsl.md) pages for details.

## Additional resources

* [Writing custom shaders in URP](urp/writing-custom-shaders-urp.md)