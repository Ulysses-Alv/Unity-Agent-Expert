# Add a texture to the frame data in URP

To pass a texture from one render pass to another within the same render graph, you can add a texture to the [frame data](accessing-frame-data.md).

Follow these steps:

1. Create a class that inherits [`ContextItem`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/api/UnityEngine.Rendering.ContextItem.html) and contains a texture handle field.

   For example:

   ```
   public class MyCustomData : ContextItem {
       public TextureHandle textureToTransfer;
   }
   ```
2. You must implement the `Reset()` method in your class, to reset the texture when the frame resets.

   For example:

   ```
   public class MyCustomData : ContextItem {
       public TextureHandle textureToTransfer;

       public override void Reset()
       {
           textureToTransfer = TextureHandle.nullHandle;
       }
   }
   ```
3. In your `RecordRenderGraph` method, add an instance of your class to the frame data.

   For example:

   ```
   public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameContext)
   {
       using (var builder = renderGraph.AddRasterRenderPass<PassData>("Get frame data", out var passData))
       {
           var customData = frameContext.Create<MyCustomData>();
       }
   }
   ```
4. Set the texture handle to your texture. For more information, refer to [Read or write to a texture in a render pass in URP](render-graph-read-write-texture.md).
5. In a later render pass, in your `RecordRenderGraph` method, you can get your custom data and fetch your texture:

For example:

```
// Get the custom data
MyCustomData customData = frameData.Get<MyCustomData>();

// Get the texture
TextureHandle customTexture = customData.textureToTransfer;
```

For more information about frame data, refer to [Use frame data](accessing-frame-data.md).

## Example

The following example adds a `CustomData` class that contains a texture. The first render pass clears the texture to yellow, and the second render pass fetches the yellow texture and draws a triangle onto it. To see the render passes, open the [Frame Debugger](../FrameDebugger.md).

```
using UnityEngine;
using UnityEngine.Rendering.Universal;
using UnityEngine.Rendering.RenderGraphModule;
using UnityEngine.Rendering;

public class AddOwnTextureToFrameData : ScriptableRendererFeature
{
    AddOwnTexturePass customPass1;
    DrawTrianglePass customPass2;

    public override void Create()
    {
        customPass1 = new AddOwnTexturePass();
        customPass2 = new DrawTrianglePass();

        customPass1.renderPassEvent = RenderPassEvent.AfterRenderingOpaques;
        customPass2.renderPassEvent = RenderPassEvent.AfterRenderingOpaques;
    }

    public override void AddRenderPasses(ScriptableRenderer renderer, ref RenderingData renderingData)
    {
        renderer.EnqueuePass(customPass1);
        renderer.EnqueuePass(customPass2);
    }
    
    // Create the first render pass, which creates a texture and adds it to the frame data
    class AddOwnTexturePass : ScriptableRenderPass
    {

        class PassData
        {
            internal TextureHandle copySourceTexture;
        }

        // Create the custom data class that contains the new texture
        public class CustomData : ContextItem {
            public TextureHandle newTextureForFrameData;

            public override void Reset()
            {
                newTextureForFrameData = TextureHandle.nullHandle;
            }
        }

        public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameContext)
        {
            using (var builder = renderGraph.AddRasterRenderPass<PassData>("Create new texture", out var passData))
            {
                // Create a texture and set it as the render target
                UniversalResourceData frameData = frameContext.Get<UniversalResourceData>();
                TextureDesc textureDesc = frameData.activeColorTexture.GetDescriptor(renderGraph);
                textureDesc.msaaSamples = MSAASamples.None;
                TextureHandle texture = renderGraph.CreateTexture(textureDesc);
                CustomData customData = frameContext.Create<CustomData>();
                customData.newTextureForFrameData = texture;
                builder.SetRenderAttachment(texture, 0, AccessFlags.Write);
    
                // Make sure the render graph system keeps the render pass, even if it's not used in the final frame.
                // Don't use this in production code, because it prevents the render graph system from removing the render pass if it's not needed.
                builder.AllowPassCulling(false);

                builder.SetRenderFunc(static (PassData data, RasterGraphContext context) => ExecutePass(data, context));
            }
        }

        static void ExecutePass(PassData data, RasterGraphContext context)
        {          
            // Clear the render target (the texture) to yellow
            context.cmd.ClearRenderTarget(true, true, Color.yellow);
        }
 
    }

    // Create the second render pass, which fetches the texture and writes to it
    class DrawTrianglePass : ScriptableRenderPass
    {

        class PassData
        {
            // No local pass data needed
        }      

        public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameContext)
        {
            using (var builder = renderGraph.AddRasterRenderPass<PassData>("Fetch texture and draw triangle", out var passData))
            {                                
                // Fetch the yellow texture from the frame data and set it as the render target
                var customData = frameContext.Get<AddOwnTexturePass.CustomData>();
                var customTexture = customData.newTextureForFrameData;
                builder.SetRenderAttachment(customTexture, 0, AccessFlags.Write);

                // Make sure the render graph system keeps the render pass, even if it's not used in the final frame.
                // Don't use this in production code, because it prevents the render graph system from removing the render pass if it's not needed.
                builder.AllowPassCulling(false);

                builder.SetRenderFunc(static (PassData data, RasterGraphContext context) => ExecutePass(data, context));
            }
        }

        static void ExecutePass(PassData data, RasterGraphContext context)
        {          
            // Generate a triangle mesh
            Mesh mesh = new Mesh();
            mesh.vertices = new Vector3[] { new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(0, 1, 0) };
            mesh.triangles = new int[] { 0, 1, 2 };
            
            // Draw a triangle to the render target (the yellow texture)
            context.cmd.DrawMesh(mesh, Matrix4x4.identity, new Material(Shader.Find("Universal Render Pipeline/Unlit")));
        }
    }
}
```