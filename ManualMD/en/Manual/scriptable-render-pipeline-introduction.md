# Scriptable Render Pipeline fundamentals

This page explains how Unity’s Scriptable **Render Pipeline** (SRP) works, and introduces some key concepts and terminology. The information on this page is applicable to the Universal Render Pipeline (URP), the High Definition Render Pipeline (HDRP), and custom render pipelines that are based on SRP.

The Scriptable Render Pipeline is a thin API layer that lets you schedule and configure rendering commands using C# **scripts**. Unity passes these commands to its low-level graphics architecture, which then sends instructions to the graphics API.

URP and HDRP are built on top of SRP. You can also create your own custom render pipeline on top of SRP.

## Render Pipeline Instance and Render Pipeline Asset

Every render pipeline based on SRP has two key customized elements:

* A **Render Pipeline Instance**. This is an instance of a class that defines the functionality of your render pipeline. Its script inherits from [RenderPipeline](../ScriptReference/Rendering.RenderPipeline.md), and overrides its `Render()` method.
* A **Render Pipeline Asset**. This is an asset in your Unity Project that stores data about which Render Pipeline Instance to use, and how to configure it. Its script inherits from [RenderPipelineAsset](../ScriptReference/Rendering.RenderPipelineAsset.md) and overrides its `CreatePipeline()` method.

For more information on these elements, and instructions on how to create them in a custom render pipeline, see [Creating a Render Pipeline Asset and a Render Pipeline Instance](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/manual/index.html).

## ScriptableRenderContext

`ScriptableRenderContext` is a class that acts as an interface between the custom C# code in the render pipeline and Unity’s low-level graphics code.

Use the [ScriptableRenderContext](../ScriptReference/Rendering.ScriptableRenderContext.md) API to schedule and execute rendering commands. For information, see [Scheduling and executing rendering commands in the Scriptable Render Pipeline](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/manual/srp-using-scriptable-render-context.html).

## Additional resources

* [Universal Render Pipeline](universal-render-pipeline.md)
* [High Definition Render Pipeline](high-definition-render-pipeline.md)
* [Creating a custom render pipeline](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/manual/srp-custom.html)