# Blit using the render graph system in URP

To **blit** from one texture to another in the render graph system in the Universal **Render Pipeline** (URP), use the [`AddBlitPass`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.Util.RenderGraphUtils.html#UnityEngine_Rendering_RenderGraphModule_Util_RenderGraphUtils_AddBlitPass_UnityEngine_Rendering_RenderGraphModule_RenderGraph_UnityEngine_Rendering_RenderGraphModule_Util_RenderGraphUtils_BlitMaterialParameters_System_String_System_Boolean_) API. The API generates a render pass automatically, so you don’t need to use a method like `AddRasterRenderPass`.

Follow these steps:

1. To create a shader and material that works with a blit render pass, from the main menu select **Assets** > **Create** > **Shader** > **SRP Blit Shader**, then create a material from it.

   To use **shader** graph instead, refer to [Create a low-code custom post-processing effect](post-processing/post-processing-custom-effect-low-code.md).

   For more information about writing shaders in URP, refer to [Writing custom shaders in URP](writing-custom-shaders-urp.md).
2. Add `using UnityEngine.Rendering.RenderGraphModule.Util` to your render pass script.
3. In your render pass, create a field for the blit material. For example:

   ```
   public class MyBlitPass : ScriptableRenderPass
   {
       Material blitMaterial;
   }
   ```
4. Set up the texture to blit from and blit to. For example:

   ```
   TextureHandle sourceTexture = renderGraph.CreateTexture(sourceTextureProperties);
   TextureHandle destinationTexture = renderGraph.CreateTexture(destinationTextureProperties);
   ```

   For more information, refer to [Get data from the current frame in URP](accessing-frame-data.md) and [Create a temporary texture for a single render pass](render-graph-create-a-texture.md).
5. To set up the material, textures, and shader pass for the blit operation, create a [`RenderGraphUtils.BlitMaterialParameters`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.Util.RenderGraphUtils.BlitMaterialParameters.html) object. For example:

   ```
   // Create a BlitMaterialParameters object with the blit material, source texture, destination texture, and shader pass to use.
   var blitParams = new RenderGraphUtils.BlitMaterialParameters(sourceTexture, destinationTexture, blitMaterial, 0);
   ```
6. To add a blit pass, call the [`AddBlitPass`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.Util.RenderGraphUtils.html#UnityEngine_Rendering_RenderGraphModule_Util_RenderGraphUtils_AddBlitPass_UnityEngine_Rendering_RenderGraphModule_RenderGraph_UnityEngine_Rendering_RenderGraphModule_Util_RenderGraphUtils_BlitMaterialParameters_System_String_System_Boolean_) method with the blit parameters. For example:

   ```
   renderGraph.AddBlitPass(blitParams, "Pass created with AddBlitPass");
   ```

For a complete example, refer to the example called **BlitWithMaterial** in the render graph examples of the [Universal Render Pipeline (URP) package samples](package-sample-urp-package-samples.md#rendergraph-samples).

If you use `AddBlitPass` with a default material, Unity might use the `AddCopyPass` API instead, to optimize the render pass so it accesses the framebuffer from the on-chip memory of the GPU instead of video memory. This process is sometimes called framebuffer fetch. For more information, refer to [`AddCopyPass`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.Util.RenderGraphUtils.html#UnityEngine_Rendering_RenderGraphModule_Util_RenderGraphUtils_AddCopyPass_UnityEngine_Rendering_RenderGraphModule_RenderGraph_UnityEngine_Rendering_RenderGraphModule_TextureHandle_UnityEngine_Rendering_RenderGraphModule_TextureHandle_System_String_System_Boolean_) API.

## Avoid blitting back

After a blit, you usually blit the destination texture back to the active color texture. However in the render graph system, you can update the frame data to point to the destination texture instead, so you only blit once. For example:

```
// Set the camera color as the destination texture you blitted to.
frameData.cameraColor = destinationTexture;
```

For more information, refer to [Avoid blitting to and from the color buffer](render-graph-optimize.md#avoid-blitting-to-and-from-the-color-buffer).

## Customize the render pass

To customize the render pass that the methods generate, for example to change settings or add more resources, call the APIs with the `returnBuilder` parameter set to `true`. The APIs then return the `IBaseRenderGraphBuilder` object that you usually receive as `var builder` from a method like `AddRasterRenderPass`.

For example:

```
using (var builder = renderGraph.AddBlitPass(blitParams, "Pass created with AddBlitPass", returnBuilder: true))
{
    // Use the builder variable to customize the render pass here.
}
```

## Additional resources

* [Get the current framebuffer from GPU memory](render-graph-framebuffer-fetch.md)
* [Render graph system](render-graph.md)
* [Universal Render Pipeline (URP) package samples](package-sample-urp-package-samples.md#rendergraph-samples): refer to the **Blit** example which uses **AddCopyPass**.