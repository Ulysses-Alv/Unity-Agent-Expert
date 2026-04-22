# Create a temporary texture for a single frame in URP

To create a temporary texture in the render graph system, use the `CreateTexture` method of the render graph instance.

When you create a texture inside a render pass, the render graph system handles the creation and disposal of the texture. This process means the texture might not exist in the next frame, and other **cameras** might not be able to use it. To make sure a texture is available across frames and cameras, refer to [Use a texture in multiple render passes](render-graph-pass-textures-between-passes.md) instead.

When the render graph system optimizes the **render pipeline**, it might not create a texture if the final frame doesn’t use the texture, to reduce the memory and bandwidth the render passes use.

## Create a temporary texture

In the `RecordRenderGraph` method of your `ScriptableRenderPass` class, follow these steps:

1. Create a [`TextureDesc`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest/index.html?subfolder=/api/UnityEngine.Rendering.RenderGraphModule.TextureDesc.html) object with the texture properties you need.
2. Use the [`CreateTexture`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest/index.html?subfolder=/api/UnityEngine.Rendering.RenderGraphModule.RenderGraph.html#UnityEngine_Rendering_RenderGraphModule_RenderGraph_CreateTexture_UnityEngine_Rendering_RenderGraphModule_TextureDesc__) method of the render graph instance to create a texture and return a texture handle.

For example, the following creates a texture the same size as the screen.

```
UniversalResourceData resourceData = frameData.Get<UniversalResourceData>();
TextureDesc textureDesc = resourceData.activeColorTexture.GetDescriptor(renderGraph);
TextureHandle textureHandle = renderGraph.CreateTexture(textureDesc);
```

**Note**: Avoid using `UniversalRenderer.CreateRenderGraphTexture`, because it can introduce subtle bugs.

The render graph system manages the lifetime of textures you create with `CreateTexture`, so you don’t need to manually dispose of the memory they use when you’re finished with them.

### Inspect how URP handles load actions on a texture

You can use the [Render Graph Viewer window](render-graph-viewer-reference.md) to inspect how URP handles load actions on textures. This is useful to debug your results and confirm whether Unity clears, preserves, or discards a texture.

**Note**: The `clear` parameter and the `AccessFlags` value affect how URP sets the load action for a texture. This can impact rendering results:

* If `clear` is `true` and you use [`AccessFlags.WriteAll`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest/index.html?subfolder=/api/UnityEngine.Rendering.RenderGraphModule.AccessFlags.html), the load action is `DontCare`, not `Clear`. This is because Unity assumes you’re overwriting the entire texture, so it skips clearing the texture to optimize performance.
* If `clear` is `false` and you use [`AccessFlags.Write`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest/index.html?subfolder=/api/UnityEngine.Rendering.RenderGraphModule.AccessFlags.html), the load action is `Clear`. This is because Unity clears the texture automatically to avoid rendering artifacts if you only write part of the texture. To skip Unity clearing the texture, use `AccessFlags.WriteAll` to explicitly tell Unity you’re writing the full texture.

### Example

The following Scriptable Renderer Feature contains an example render pass that creates a texture and clears it to yellow. For more information about adding the render pass to the render pipeline, refer to [Inject a pass using a Scriptable Renderer Feature](renderer-features/scriptable-renderer-features/inject-a-pass-using-a-scriptable-renderer-feature.md).

Use the [Frame Debugger](https://docs.unity3d.com/2023.3/Documentation/Manual/frame-debugger-window.html) to check the texture the render pass adds.

```
using UnityEngine;
using UnityEngine.Rendering.Universal;
using UnityEngine.Rendering.RenderGraphModule;
using UnityEngine.Rendering;

public class CreateYellowTextureFeature : ScriptableRendererFeature
{
    CreateYellowTexture customPass;

    public override void Create()
    {
        customPass = new CreateYellowTexture();
        customPass.renderPassEvent = RenderPassEvent.AfterRenderingPostProcessing;
    }

    public override void AddRenderPasses(ScriptableRenderer renderer, ref RenderingData renderingData)
    {
        renderer.EnqueuePass(customPass);
    }

    class CreateYellowTexture : ScriptableRenderPass
    {
        class PassData
        {
            public TextureHandle cameraColorTexture;
        }

        public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameData)
        {
            using (var builder = renderGraph.AddRasterRenderPass<PassData>("Create yellow texture", out var passData))
            {
                // Get the frame data
                UniversalResourceData resourceData = frameData.Get<UniversalResourceData>();

                // Create texture properties that match the screen size
                TextureDesc textureDesc = resourceData.activeColorTexture.GetDescriptor(renderGraph);
                textureDesc.msaaSamples = MSAASamples.None;

                // Create a temporary texture
                TextureHandle texture = renderGraph.CreateTexture(textureDesc);

                // Set the texture as the render target
                builder.SetRenderAttachment(texture, 0, AccessFlags.Write);
    
                // Make sure the render graph system keeps the render pass, even if it's not used in the final frame.
                // Don't use this in production code, because it prevents the render graph system from removing the render pass if it's not needed.
                builder.AllowPassCulling(false);

                builder.SetRenderFunc(static (PassData data, RasterGraphContext context) => ExecutePass(data, context));
            }
        }

        static void ExecutePass(PassData data, RasterGraphContext context)
        {          
            // Clear the render target to yellow
            context.cmd.ClearRenderTarget(true, true, Color.yellow);            
        }
    }

}
```

## Additional resources

* [Use a texture in multiple render passes](render-graph-pass-textures-between-passes.md)
* [Use frame data](accessing-frame-data.md)
* [Textures](https://docs.unity3d.com/Manual/Textures.html)