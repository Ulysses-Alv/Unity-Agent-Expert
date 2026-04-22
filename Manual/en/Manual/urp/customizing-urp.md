# Custom rendering and post-processing in URP

Customize and extend the rendering process in the Universal **Render Pipeline** (URP). Create a custom render pass in a C# script and inject it into the URP frame rendering loop.

| Page | Description |
| --- | --- |
| [Introduction to Scriptable Render Passes](renderer-features/intro-to-scriptable-render-passes.md) | Learn about using Scriptable Render Passes to alter how Unity renders a **scene** or the objects within a scene. |
| [Adding pre-built effects with Renderer Features in URP](urp-renderer-feature-landing.md) | Resources for adding pre-built render passes to URP, and configuring their behaviour. |
| [Custom render pass workflow in URP](renderer-features/custom-rendering-pass-workflow-in-urp.md) | Add and inject a custom render pass. |
| [Blit](customize/blit-overview.md) | Understand the different ways to perform a blit operation in URP and best practices to follow when writing custom render passes. |
| [Render graph system](render-graph.md) | Resources and approaches for using the `RenderGraph` APIs to create a Scriptable Render Pass. |
| [Adding a Scriptable Render Pass to the frame rendering loop](inject-a-render-pass.md) | Resources and techniques for injecting a custom render pass via a Scriptable Renderer Feature, or the `RenderPipelineManager` API. |
| [Modify URP source code](customize/modify-urp-source-code.md) | Modify URP source code to implement advanced customizations. |

## Additional resources

* [Rendering](rendering-in-universalrp.md)
* [Render pipeline concepts](urp-concepts.md)
* [Pre-built effects (Renderer Features)](urp-renderer-feature.md)
* [How to create a custom post-processing effect](post-processing/post-processing-custom-effect-low-code.md)
* [Execute rendering commands in a custom render pipeline](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/manual/srp-using-scriptable-render-context.html)
* [Universal Render Pipeline scripting API](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.2/api/index.html)