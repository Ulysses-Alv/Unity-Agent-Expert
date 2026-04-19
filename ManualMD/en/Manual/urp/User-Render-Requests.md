# Render to a render texture outside the URP rendering loop

To trigger a **camera** to render to a **render texture** outside of the Universal **Render Pipeline** (URP) rendering loop, use the `SingleCameraRequest` and `SubmitRenderRequest` APIs in a C# script.

Follow these steps:

1. Create a render request of the [`UniversalRenderPipeline.SingleCameraRequest`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.UniversalRenderPipeline.SingleCameraRequest.html) type. For example:

   ```
   UniversalRenderPipeline.SingleCameraRequest request = new UniversalRenderPipeline.SingleCameraRequest();
   ```
2. Check whether the camera supports the render request type, using the [`RenderPipeline.SupportsRenderRequest`](../../ScriptReference/Rendering.RenderPipeline.SupportsRenderRequest.md) API. For example, to check the main camera:

   ```
   Camera mainCamera = Camera.main;

   if (RenderPipeline.SupportsRenderRequest(mainCamera, request))
   {
       ...
   }
   ```
3. Set the target of the camera to a `RenderTexture` object, using the `destination` parameter of the render request. For example:

   ```
   request.destination = myRenderTexture;
   ```
4. Render to the render texture using the [SubmitRenderRequest](../../ScriptReference/Rendering.RenderPipeline.SubmitRenderRequest.md) API. For example:

   ```
   RenderPipeline.SubmitRenderRequest(mainCamera, request);
   ```

To make sure all cameras finish rendering before you render to the render texture, use either of the following approaches:

* A coroutine that waits for the end of the frame. For more information, refer to the [WaitForEndOfFrame](../../ScriptReference/WaitForEndOfFrame.md) API.
* A callback. For more information, refer to the [RenderPipelineManager.endContextRendering](../../ScriptReference/Rendering.RenderPipelineManager-endContextRendering.md) API.

## Example

The following example renders multiple cameras to multiple render textures. To use the example, follow these steps:

1. In your Unity project, add the code to a new C# script called `SingleCameraRenderRequest.cs`.
2. Add the script to a **GameObject** in your **scene**.
3. In the **Inspector** window of the GameObject, assign the cameras and render textures. Make sure the number of cameras is the same as the number of render textures.
4. Enter Play mode.

```
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Rendering;
using UnityEngine.Rendering.Universal;

public class SingleCameraRenderRequest : MonoBehaviour
{
    public Camera[] cameras;
    public RenderTexture[] renderTextures;

    void Start()
    {
        // Make sure all data is valid before you start the component
        if (cameras == null || cameras.Length == 0 || renderTextures == null || cameras.Length != renderTextures.Length)
        {
            Debug.LogError("Invalid setup");
            return;
        }

        // Start the asynchronous coroutine
        StartCoroutine(RenderSingleRequestNextFrame());
        
        // Call a method called OnEndContextRendering when a camera finishes rendering
        RenderPipelineManager.endContextRendering += OnEndContextRendering;
    }

    void OnEndContextRendering(ScriptableRenderContext context, List<Camera> cameras)
    {
        // Create a log to show cameras have finished rendering
        Debug.Log("All cameras have finished rendering.");
    }

    void OnDestroy()
    {
        // End the subscription to the callback
        RenderPipelineManager.endContextRendering -= OnEndContextRendering;
    }

    IEnumerator RenderSingleRequestNextFrame()
    {
        // Wait for the main camera to finish rendering
        yield return new WaitForEndOfFrame();

        // Enqueue one render request for each camera
        SendSingleRenderRequests();

        // Wait for the end of the frame
        yield return new WaitForEndOfFrame();

        // Restart the coroutine
        StartCoroutine(RenderSingleRequestNextFrame());
    }

    void SendSingleRenderRequests()
    {
        //Iterates over the cameras array.        
        for (int i = 0; i < cameras.Length; i++)
        {
            UniversalRenderPipeline.SingleCameraRequest request =
                new UniversalRenderPipeline.SingleCameraRequest();

            // Check if the active render pipeline supports the render request
            if (RenderPipeline.SupportsRenderRequest(cameras[i], request))
            {
                // Set the destination of the camera output to the matching RenderTexture
                request.destination = renderTextures[i];
                
                // Render the camera output to the RenderTexture synchronously
                RenderPipeline.SubmitRenderRequest(cameras[i], request);

                // At this point, the RenderTexture in renderTextures[i] contains the scene rendered from the point
                // of view of the Camera in cameras[i]
            }
        }
    }
}
```

## Additional resources

* [Cameras in URP](urp-cameras-landing.md)
* [Rendering to a texture](../render-texture-landing.md)