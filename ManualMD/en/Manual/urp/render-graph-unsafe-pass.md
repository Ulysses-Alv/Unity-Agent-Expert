# Use the CommandBuffer interface in a render graph

You can use the render graph `AddUnsafePass` API to use [CommandBuffer](../../ScriptReference/Rendering.CommandBuffer.md) interface APIs such as `SetRenderTarget` in render graph system render passes.

The `AddUnsafePass` API gives access to the [UnsafeCommandBuffer](../../ScriptReference/Rendering.UnsafeCommandBuffer.md), which gives access to more `CommandBuffer` functions than the [RasterCommandBuffer](../../ScriptReference/Rendering.RasterCommandBuffer.md) API.

**Note**: Both `AddUnsafePass` and `RasterCommandBuffer` restrict access to certain `commandBuffer` functions so that `RenderGraph` can better optimize the frame. The most important function `UnsafeCommandbuffer` gives you access to is `SetRenderTarget`.

If you use the `AddUnsafePass` API, the following applies:

* You can’t use the `SetRenderAttachment` method in the `RecordRenderGraph` method. Use `SetRenderTarget` in the `SetRenderFunc` method instead.
* Rendering might be slower because the render graph system can’t optimize the render pass. For example, if your render pass writes to the active color buffer, the render graph system can’t detect if a later render pass writes to the same buffer. As a result, the render graph system can’t merge the two render passes, and the GPU unnecessarily transfers the buffer in and out of memory.

## Create an unsafe render pass

To create an unsafe render pass, follow these steps:

1. In your `RecordRenderGraph` method, use the `AddUnsafePass` method instead of the `AddRasterRenderPass` method.

   For example:

   ```
   using (var builder = renderGraph.AddUnsafePass<PassData>("My unsafe render pass", out var passData))
   ```
2. When you call the `SetRenderFunc` method, use the `UnsafeGraphContext` type instead of `RasterGraphContext`.

   For example:

   ```
   builder.SetRenderFunc(
       static (PassData passData, UnsafeGraphContext context) => ExecutePass(passData, context)
   );
   ```
3. If your render pass writes to a texture, add the texture as a field in your pass data class.

   For example:

   ```
   class PassData
   {
       public TextureHandle textureToWriteTo;
   }
   ```
4. If your render pass writes to a texture, set the texture as writeable using the `UseTexture` method.

   For example:

   ```
   builder.UseTexture(passData.textureToWriteTo, AccessFlags.Write);
   ```

You can now use CommandBuffer interface APIs in your `SetRenderFunc` method.

## Example

The following example uses the CommandBuffer interface `SetRenderTarget` API to set the render target to the active color buffer during the render pass, then draw objects using their surface normals as colors.

```
using UnityEngine;
using UnityEngine.Rendering.RenderGraphModule;
using UnityEngine.Rendering;
using UnityEngine.Rendering.Universal;

public class DrawNormalsToActiveColorTexture : ScriptableRendererFeature
{

    DrawNormalsPass unsafePass;

    public override void Create()
    {
        unsafePass = new DrawNormalsPass();
        unsafePass.renderPassEvent = RenderPassEvent.AfterRenderingPostProcessing;
    }

    public override void AddRenderPasses(ScriptableRenderer renderer, ref RenderingData renderingData)
    {
        renderer.EnqueuePass(unsafePass);
    }

    class DrawNormalsPass : ScriptableRenderPass
    {
        class PassData
        {
            public TextureHandle activeColorBuffer;
            public TextureHandle cameraNormalsTexture;
        }

        public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameContext)
        {
            using (var builder = renderGraph.AddUnsafePass<PassData>("Draw normals", out var passData))
            {
                // Make sure URP generates the normals texture
                ConfigureInput(ScriptableRenderPassInput.Normal);

                // Get the frame data
                UniversalResourceData resourceData = frameContext.Get<UniversalResourceData>();

                // Add the active color buffer to our pass data, and set it as writeable 
                passData.activeColorBuffer = resourceData.activeColorTexture;
                builder.UseTexture(passData.activeColorBuffer, AccessFlags.Write);                

                // Add the camera normals texture to our pass data 
                passData.cameraNormalsTexture = resourceData.cameraNormalsTexture;
                builder.UseTexture(passData.cameraNormalsTexture);

                // Make sure the render graph system keeps the render pass, even if it's not used in the final frame.
                // Don't use this in production code, because it prevents the render graph system from removing the render pass if it's not needed.
                builder.AllowPassCulling(false);

                builder.SetRenderFunc(static (PassData data, UnsafeGraphContext context) => ExecutePass(data, context));
            }
        }

        static void ExecutePass(PassData passData, UnsafeGraphContext context)
        {
            // Create a command buffer for a list of rendering methods
            CommandBuffer unsafeCommandBuffer = CommandBufferHelpers.GetNativeCommandBuffer(context.cmd);

            // Add a command to set the render target to the active color buffer so URP draws to it
            context.cmd.SetRenderTarget(passData.activeColorBuffer);

            // Add a command to copy the camera normals texture to the render target
            Blitter.BlitTexture(unsafeCommandBuffer, passData.cameraNormalsTexture, new Vector4(1, 1, 0, 0), 0, false);
        }

    }

}
```

For another example, refer to the example called **UnsafePass** in the [Universal Render Pipeline (URP) package samples](package-samples.md).