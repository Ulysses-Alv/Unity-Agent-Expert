# Inject a render pass with a Scriptable Renderer Feature in URP

Use the `ScriptableRenderFeature` API to insert a [Scriptable Render Pass](../../renderer-features/intro-to-scriptable-render-passes.md) into the Universal **Render Pipeline** (URP) frame rendering loop.

Follow these steps:

1. Create a new C# script.
2. Replace the code with a class that inherits from the `ScriptableRendererFeature` class.

   ```
   using UnityEngine;
   using UnityEngine.Rendering.Universal;

   public class MyRendererFeature : ScriptableRendererFeature
   {
   }
   ```
3. In the class, override the `Create` method. For example:

   ```
   public override void Create()
   {
   }
   ```

   URP calls the `Create` methods on the following events:

   * When the Scriptable Renderer Feature loads the first time.
   * When you enable or disable the Scriptable Renderer Feature.
   * When you change a property in the ****Inspector**** window of the Renderer Feature.
4. In the `Create` method, create an instance of your Scriptable Render Pass, and inject it into the renderer.

   For example, if you have a Scriptable Render Pass called `RedTintRenderPass`:

   ```
   // Define an instance of the Scriptable Render Pass
   private RedTintRenderPass redTintRenderPass;

   public override void Create()
   {
       // Create an instance of the Scriptable Render Pass
       redTintRenderPass = new RedTintRenderPass();

       // Inject the render pass after rendering the skybox
       redTintRenderPass.renderPassEvent = RenderPassEvent.AfterRenderingSkybox;
   }
   ```

   You can also use the `Create` method to initialize any other resources the Scriptable Renderer Feature needs, such as materials.
5. Override the `AddRenderPasses` method.

   ```
   public override void AddRenderPasses(ScriptableRenderer renderer, ref RenderingData renderingData)
   {
   }
   ```

   URP calls the `AddRenderPasses` method every frame, once for each **camera**. Don’t create or instantiate any resources within this method.
6. Use the `EnqueuePass` API to inject the Scriptable Render Pass into the frame rendering loop.

   ```
   public override void AddRenderPasses(ScriptableRenderer renderer, ref RenderingData renderingData)
   {
       renderer.EnqueuePass(redTintRenderPass);
   }
   ```

   To add a render pass to a specific camera only, use an `if` statement to check the camera type. For example:

   ```
   if (renderingData.cameraData.cameraType == CameraType.Game)
   {
       renderer.EnqueuePass(redTintRenderPass);
   }
   ```

You can now add the Scriptable Renderer Feature to the active URP asset. Refer to [How to add a Renderer Feature to a Renderer](../../urp-renderer-feature.md) for more information.

To dispose of any resources, override the [`Dispose`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest/index.html?subfolder=/api/UnityEngine.Rendering.Universal.ScriptableRendererFeature.html#UnityEngine_Rendering_Universal_ScriptableRendererFeature_Dispose) method.

## Example

For a complete example of a Scriptable Renderer Feature with a custom render pass that uses a material to **blit**, follow these steps:

1. Import the URP render graph samples. For more information, refer to [Importing a package sample in URP](../../package-sample-urp-package-samples.md).
2. In the **Project** window, go to **Samples** > **Universal Render Pipeline** > <your version> > **URP RenderGraph Samples** > **BlitWithMaterial**.

## Additional resources

* [`ScriptableRendererFeature`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest/index.html?subfolder=/api/UnityEngine.Rendering.Universal.ScriptableRendererFeature.html)
* [Add a Renderer Feature to a URP Renderer](../../urp-renderer-feature.md)
* [Restrict a render pass to a scene area](../../customize/restrict-render-pass-scene-area.md)
* [Create a custom post-processing effect with Volume support in URP](../../post-processing/custom-post-processing-with-volume.md)
* [Custom render pass workflow in URP](../custom-rendering-pass-workflow-in-urp.md)
* [Render graph system](../../render-graph.md)