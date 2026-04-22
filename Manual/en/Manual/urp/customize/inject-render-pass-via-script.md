# Inject a render pass using the RenderPipelineManager API in URP

To inject a custom render pass into the rendering loop in the Universal Rendering Pipeline (URP), subscribe to one of the events in the [RenderPipelineManager](https://docs.unity3d.com/ScriptReference/Rendering.RenderPipelineManager.html) API and use the [`EnqueuePass`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.2/api/UnityEngine.Rendering.Universal.ScriptableRenderer.html#UnityEngine_Rendering_Universal_ScriptableRenderer_EnqueuePass_UnityEngine_Rendering_Universal_ScriptableRenderPass) API.

For example, you can render extra cameras to **render textures**, and use those textures for effects like planar reflections or surveillance **camera** views.

**Note**: `RenderPipelineManager` events trigger only if a camera is active.

If you want to apply a render pass across multiple cameras, **scenes**, or across your entire project, the recommended best practice is to use a Scriptable Renderer Feature instead. For more information, refer to [Inject a custom render pass using a Scriptable Renderer Feature](../renderer-features/scriptable-renderer-features/inject-a-pass-using-a-scriptable-renderer-feature.md).

Follow these steps to use the `RenderPipelineManager` API:

1. Attach a new `MonoBehaviour` script to a **GameObject**, then create an instance of a custom render pass. For example:

   ```
   public class InjectPass : MonoBehaviour
   {
       // Declare a custom render pass
       private ExampleRenderPass exampleRenderPass;

       private void OnEnable()
       {
           // Create the custom render pass
           exampleRenderPass = new ExampleRenderPass(settings);
       }
   }
   ```

   For more information about writing a render pass, refer to [Write a render pass using the render graph system](../render-graph-write-render-pass.md).
2. Subscribe a method to one of the events in the `RenderPipelineManager` API. For example:

   ```
       private void OnEnable()
       {
           ...
           // Call the InjectRenderPass method when the beginCameraRendering event is raised.
           RenderPipelineManager.beginCameraRendering += InjectRenderPass;            
       }
   ```
3. In the subscribed method, use the [`EnqueuePass`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.2/api/UnityEngine.Rendering.Universal.ScriptableRenderer.html#UnityEngine_Rendering_Universal_ScriptableRenderer_EnqueuePass_UnityEngine_Rendering_Universal_ScriptableRenderPass) API method to inject a custom render pass into the URP rendering loop.

   ```
   public class InjectPass : MonoBehaviour
   {
       ...

       // Add the method that Unity calls after the beginCameraRendering event is raised.
       // The method must accept the ScriptableRenderContext and Camera parameters the event delegate defines.
       private void InjectRenderPass(ScriptableRenderContext context, Camera cam)
       {
           ...
           // Use the EnqueuePass API to inject the custom render pass.
           cam.GetUniversalAdditionalCameraData().scriptableRenderer.EnqueuePass(exampleRenderPass);
       }
   }
   ```
4. Unsubscribe from the event in the `OnDisable` method. For example:

   ```
   public class InjectPass : MonoBehaviour
   {
       ...

       private void OnDisable()
       {
           RenderPipelineManager.beginCameraRendering -= InjectRenderPass;
       }
   }
   ```

## Additional resources

* [Write a render pass using the render graph system](../render-graph-write-render-pass.md)
* [Custom render pass workflow](../renderer-features/custom-rendering-pass-workflow-in-urp.md)