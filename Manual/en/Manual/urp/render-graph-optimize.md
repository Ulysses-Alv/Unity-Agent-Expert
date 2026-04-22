# Optimize a render graph

To optimize a render graph, merge or reduce the number of render passes. The more render passes you have, the more data the CPU and GPU need to store and retrieve from memory. This slows down rendering, especially on devices that use tile-based deferred randering (TBDR).

## Use existing copies of the color or depth buffers

If you need a copy of the color or **depth buffers**, avoid copying them yourself if you can. Use the copies URP creates by default instead, to avoid creating unnecessary render passes.

Use the [ConfigureInput](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest/index.html?subfolder=/api/UnityEngine.Rendering.Universal.ScriptableRenderPass.html#UnityEngine_Rendering_Universal_ScriptableRenderPass_ConfigureInput_UnityEngine_Rendering_Universal_ScriptableRenderPassInput_) API to make sure URP generates the texture you need in the frame data.

To check if URP creates copies during the frame that you can use, check for the following passes in the [Render Graph Viewer](render-graph-view.md):

* **Copy Color** - copies color from `_CameraTargetAttachment` to `cameraOpaqueTexture` in the [frame data](render-graph-frame-data-reference.md).
* **Copy Depth** - copies depth from `_CameraDepthAttachment` to `cameraDepthTexture` in the [frame data](render-graph-frame-data-reference.md).

## Merge render passes

Use the [Render Graph Viewer](render-graph-view.md) to check the reason why URP can’t merge render passes, and fix the issue if you can. On devices that use tile-based deferred randering (TBDR), merging passes helps the device use less energy and run for longer.

You can do the following to make sure URP merges render passes:

* Use `AddRasterRenderPass` instead of other types of render pass as much as possible.
* If a render pass needs only to read the current **pixel** from the framebuffer instead of neighboring pixels, use the `SetInputAttachment` API and the `LOAD_FRAMEBUFFER_X_INPUT` macro. For more information, refer to [Get the current framebuffer from GPU memory](render-graph-framebuffer-fetch.md).

## Reduce the number of render passes

Don’t create unnecessary render passes to organize your code into smaller, more manageable chunks. Each render pass you create requires more processing time on the CPU.

To write combined render passes, you can use the `AddUnsafePass` API, but rendering might be slower because URP can’t optimize the render pass. For more information, refer to [Use the CommandBuffer interface in a render graph](render-graph-unsafe-pass.md).

## Avoid blitting to and from the color buffer

To avoid creating two render passes that **blit** from and to the **camera** color texture, use the `ContextContainer` object to read and write to the color buffer directly.

For example:

```
public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameData)
{

    // Fetch the frame data textures
    var resourceData = frameData.Get<UniversalResourceData>();

    // Set the source as the color texture the camera currently targets
    var source = resourceData.activeColorTexture;
    
    // Create a destination texture, with the same dimensions as the source
    var destinationDescriptor = renderGraph.GetTextureDesc(source);
    destinationDescriptor.name = "DestinationTexture";
    destinationDescriptor.clearBuffer = false;
    TextureHandle destination = renderGraph.CreateTexture(destinationDescriptor);

    // Use the AddBlitPass API to create a simple blit from the source to the destination
    RenderGraphUtils.BlitMaterialParameters parameters = new(source, destination, BlitMaterial, 0);
    renderGraph.AddBlitPass(parameters, passName: "MyRenderPass");

    // Set the main color texture for the camera as the destination texture
    resourceData.cameraColor = destination;
}
```

## Additional resources

* [Analyze a render graph](render-graph-view.md)