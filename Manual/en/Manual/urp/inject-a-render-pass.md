# Adding a Scriptable Render Pass to the frame rendering loop in URP

Add a custom render pass to the Universal **Render Pipeline** (URP) frame rendering loop by creating a Scriptable Renderer Feature, or using the `RenderPipelineManager` API.

| Page | Description |
| --- | --- |
| [Inject a render pass using a Scriptable Renderer Feature](renderer-features/scriptable-renderer-features/inject-a-pass-using-a-scriptable-renderer-feature.md) | Write a class that inherits `ScriptableRendererFeature`, and use it to creates an instance of the custom render pass you created, and insert the custom render pass into the rendering pipeline. |
| [Inject a render pass using the RenderPipelineManager API](customize/inject-render-pass-via-script.md) | Use the `RenderPipelineManager` API to insert a custom render pass into the rendering pipeline. |
| [Restrict a render pass to a scene area](customize/restrict-render-pass-scene-area.md) | Enable a custom rendering effect only if the **camera** is inside a volume in a **scene**. |
| [Injection points reference](customize/custom-pass-injection-points.md) | URP contains multiple injection points that let you inject render passes at different points in the frame rendering loop. |

## Additional resources

* [Custom post-processing](post-processing/custom-post-processing.md)