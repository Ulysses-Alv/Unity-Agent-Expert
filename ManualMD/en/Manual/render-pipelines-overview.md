# Introduction to render pipelines

A **render pipeline** takes the objects in a **scene** and displays them on-screen.

## How a render pipeline works

![Stages of a render pipeline workflow.](../uploads/Main/BestPracticeLightingPipeline3.svg)

Stages of a render pipeline workflow.

A render pipeline follows these steps:

1. Culling, where the pipeline decides which objects from the scene to display. This usually means it removes objects that are outside the **camera** view ([frustum culling](UnderstandingFrustum.md)) or hidden behind other objects ([occlusion culling](OcclusionCulling.md)).
2. Rendering, where the pipeline draws the objects with their correct lighting into **pixel** buffers.
3. **Post-processing**, where the pipeline modifies the pixel buffers to generate the final output frame for the display. Example of modifications include color grading, bloom, and **depth of field**.

A render pipeline repeats these steps each time Unity generates a new frame.

## Render pipelines in Unity

In Unity, you can choose between different render pipelines. Unity provides three prebuilt render pipelines with different capabilities and performance characteristics, or you can create your own.

The [Universal Render Pipeline (URP)](universal-render-pipeline.md) is a Scriptable Render Pipeline that you can customize. It lets you create scalable graphics across a wide range of platforms.

![URP 3D sample](../uploads/Main/render-pipelines-overview-urp.jpg)

URP 3D sample

The [High Definition Render Pipeline (HDRP)](high-definition-render-pipeline.md) is a Scriptable Render Pipeline that lets you create cutting-edge, high-fidelity graphics on high-end platforms.

![HDRP 3D sample: The Enemies Unity demo](../uploads/Main/render-pipelines-overview-hdrp.jpg)

HDRP 3D sample: The Enemies Unity demo

The [Built-In Render Pipeline](built-in-render-pipeline.md) is a general purpose render pipeline with limited options for customization.

![Built-In Render Pipeline 3D sample](../uploads/Main/render-pipelines-overview-builtin.jpg)

Built-In Render Pipeline 3D sample

The Scriptable Render Pipelines let you inspect and change how culling, rendering, and post-processing work directly in C#. This level of customization is also possible in the Built-In Render Pipeline when you [purchase access to the Unity engine’s source code](https://unity.com/products/source-code) in C++.

If you’re an experienced graphics developer with advanced customization needs, you can also [create your own custom render pipeline](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/manual/srp-custom.html) using Unity’s Scriptable Render Pipeline API.

Refer to [Choose a render pipeline](choose-a-render-pipeline.md) for more information about choosing the right pipeline for your project.

## Additional resources

* [Introduction to Lighting and Rendering tutorial](https://learn.unity.com/tutorial/introduction-to-lighting-and-rendering-2019-3)