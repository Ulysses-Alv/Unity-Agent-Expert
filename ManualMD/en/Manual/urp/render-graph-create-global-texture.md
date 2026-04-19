# Create a texture as a global texture in URP

If you need to use a texture as the input for the **shader** on a **GameObject**, you can set a texture as a global texture. A global texture is available to all shaders and render passes.

Setting a texture as a global texture can make rendering slower. Refer to [SetGlobalTexture](https://docs.unity3d.com/ScriptReference/Shader.SetGlobalTexture.html).

Don’t use an [unsafe render pass](render-graph-unsafe-pass.md) and `CommandBuffer.SetGlobal` to set a texture as a global texture, because it might cause errors.

To set a global texture, in the `RecordRenderGraph` method, use the [`SetGlobalTextureAfterPass`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest/index.html?subfolder=/api/UnityEngine.Rendering.RenderGraphModule.IBaseRenderGraphBuilder.html%23UnityEngine_Rendering_RenderGraphModule_IBaseRenderGraphBuilder_SetGlobalTextureAfterPass_UnityEngine_Rendering_RenderGraphModule_TextureHandle__System_Int32_) method.

For example:

```
// Allocate a global shader texture called _GlobalTexture
private int globalTextureID = Shader.PropertyToID("_GlobalTexture")

using (var builder = renderGraph.AddRasterRenderPass<PassData>("MyPass", out var passData)){

    // Set a texture to the global texture
    builder.SetGlobalTextureAfterPass(texture, globalTextureID);
}
```

If you don’t already call `SetRenderFunc`, you must also add an empty render function. For example:

```
    builder.SetRenderFunc(static (PassData data, RasterGraphContext context) => { });
```

You can now:

* Access the texture in a different render pass with the [`UseGlobalTexture()`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest/index.html?subfolder=/api/UnityEngine.Rendering.RenderGraphModule.IBaseRenderGraphBuilder.html%23UnityEngine_Rendering_RenderGraphModule_IBaseRenderGraphBuilder_UseGlobalTexture_System_Int32_UnityEngine_Rendering_RenderGraphModule_AccessFlags_) or [`UseAllGlobalTextures()`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest/index.html?subfolder=/api/UnityEngine.Rendering.RenderGraphModule.IBaseRenderGraphBuilder.html%23UnityEngine_Rendering_RenderGraphModule_IBaseRenderGraphBuilder_UseAllGlobalTextures_System_Boolean_) API.
* Use the texture on any material in your **scene** by referencing its [`nameID`](../../ScriptReference/Shader.PropertyToID.md).

## Access a global texture

Access a specific texture set as a global texture in a different render pass with the `nameID` of the texture (retrieve it with [`Shader.PropertyToId`](../../ScriptReference/Shader.PropertyToID.md)). Make sure to set the appropriate access flags depending on how the pass uses the global texture.

For example:

```
    class AccessGlobalTexturePass : ScriptableRenderPass
    {

        // The nameID of the globalTexture you want to use - which you have set in a previous pass
        private int globalTextureID = Shader.PropertyToID("_GlobalTexture")

        class PassData
        {
            // No local pass data needed
        }       

        public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameContext)
        {
            using (var builder = renderGraph.AddRasterRenderPass<PassData>("Fetch texture and draw triangle", out var passData))
            {
                
                // Set the inputs and outputs of your pass
                // builder.SetRenderAttachment(/*...*/);
                // builder.UseTexture(/*...*/);

                // Use the global texture in this pass
                builder.UseGlobalTexture(globalTextureID, AccessFlags.Read);

                builder.SetRenderFunc(static (PassData data, RasterGraphContext context) => ExecutePass(data, context));
            }
        }

        static void ExecutePass(PassData data, RasterGraphContext context)
        {
            // ...
        }
    }
```

**Note**: You can also use all the global textures in your pass with `builder.UseAllGlobalTextures(true);` instead of a single one. However, this uses more memory.