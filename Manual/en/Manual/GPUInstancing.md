# Introduction to GPU instancing

GPU instancing is a [draw call optimization](optimizing-draw-calls.md) method that uses a single draw call to render multiple **GameObjects** that use the same **mesh** and material. This speeds up rendering when you draw things that appear multiple times in a **scene**, for example, trees or bushes.

GPU instancing is a built-in functionality of GPUs. Each copy of the mesh is called an instance. Each instance can have different properties, such as color or scale.

The performance benefits of GPU instancing depend on the platform and the GPU. For each draw call, Unity has to collect, combine, and upload properties from various memory locations, so the performance overhead might outweigh the benefits. The performance benefits are better on mobile platforms than on desktop platforms.

## Render pipeline compatibility

GPU instancing is compatible with all Unity **render pipelines**, with the following limitations:

* If you use the Universal Render Pipeline (URP) or High Definition Render Pipeline (HDRP), GPU instancing works with custom **shaders** only if you disable the [Scriptable Render Pipeline (SRP) Batcher](SRPBatcher.md) or [make a shader incompatible with the SRP Batcher](SRPBatcher-Incompatible.md).
* If you use the Built-in Render Pipeline (BiRP), GPU Instancing doesn’t work with Shader Graph shaders.

For information on draw call optimization methods you can use instead of GPU instancing, refer to [Choose a method for optimizing draw calls](optimizing-draw-calls-choose-method.md).

## Indirect lighting compatibility

GPU instancing supports the following types of GameObject:

* Dynamic GameObjects that get lighting from **Light Probes**.
* Static GameObjects that get lighting from **lightmaps**, if they have **Contribute GI** enabled in their [Static Editor Flags](StaticObjects.md), and they bake to the same lightmap texture.
* GameObjects that use **Light Probe Proxy Volumes** (LPPV). You must bake the LPPV for the entire space that contains all the instances.

## Shader and mesh compatibility

The following meshes are compatible if you use prebuilt materials:

* [Mesh Renderer components](class-MeshRenderer.md) in your scene. Skinned **Mesh Renderer** components are not supported.
* Meshes you render in a script using APIs that support GPU instancing in prebuilt materials, such as [Graphics.RenderMeshInstanced](../ScriptReference/Graphics.RenderMeshInstanced.md).

The following shaders support GPU instancing:

* Most [prebuilt materials](built-in-materials-and-shaders.md). Compatible shaders have an **Enable GPU Instancing** property.
* [Shader Graph](shader-graph.md) materials, if you use URP or HDRP.

To create a custom shader that supports GPU instancing, refer to the following:

* [Creating custom shaders that support GPU instancing in the Built-In Render Pipeline](gpu-instancing-shader.md)
* [Indirect & Procedural Rendering in Shader Graph](https://discussions.unity.com/t/indirect-procedural-rendering/1664601) on the Unity Discussions site if you use URP or HDRP.

## Additional resources

* [Choose a method for optimizing draw calls](optimizing-draw-calls-choose-method.md)