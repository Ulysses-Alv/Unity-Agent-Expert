# Import a texture into a render pass in URP

To make a texture available across frames and **cameras**, create it outside the render graph system, then point a render graph to it using the `ImportTexture` API.

The render graph system doesn’t manage the lifetime of imported textures, so it can’t optimize using the texture, and there might be memory leaks. Where possible, [Create a temporary texture for a single frame](render-graph-create-a-texture.md) instead, so the render graph system manages its lifetime.

To fetch textures created by URP instead of your own code, for example the camera color and depth textures, refer to [Frame data in the render graph system](render-graph-frame-data.md).

## Create a render texture

To create a texture once but use it across multiple frames, create it before you call `AddRasterRenderPass`. For example:

1. Declare the texture using the [RTHandle](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/api/UnityEngine.Rendering.RTHandle.html) API.

   For example:

   ```
   public RTHandle renderTextureHandle;
   ```
2. Create a [`TextureDesc`](xref:UnityEngine.Rendering.RenderGraphModule.TextureDesc) object with the texture properties you need.
3. Allocate the **render texture** using the [ReAllocateHandleIfNeeded](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.4/api/UnityEngine.Rendering.Universal.RenderingUtils.html#UnityEngine_Rendering_Universal_RenderingUtils_ReAllocateHandleIfNeeded_UnityEngine_Rendering_RTHandle__UnityEngine_RenderTextureDescriptor__UnityEngine_FilterMode_UnityEngine_TextureWrapMode_System_Int32_System_Single_System_String_) API. This method checks if the render texture already exists, so it only creates the render texture once.

   For example:

   ```
   RenderingUtils.ReAllocateHandleIfNeeded(ref renderTextureHandle, textureDesc, name: "My render texture");
   ```

## Point the render graph to the texture

To use the texture after you call `AddRasterRenderPass`, use the `ImportTexture` API to get a handle to the render texture that render graph can use. For example:

```
TextureHandle texture = renderGraph.ImportTexture(renderTextureHandle);
```

You can then use the `TextureHandle` object to [read from or write to the render texture](render-graph-read-write-texture.md), and the texture won’t be disposed of at the end of the render pass.

## Dispose of the render texture

At the end of rendering, you must free the memory a render texture uses at the end of a render pass.

For example, you can create the following `Dispose` method in your `ScriptableRendererFeature` class.

```
public void Dispose()
{
    renderTexture.Release();
}
```

## Example

The following Scriptable Renderer Feature creates a texture outside the render graph system, then uses it in a render pass to draw a random triangle each frame. To see the texture, open the [Frame Debugger](https://docs.unity3d.com/Manual/FrameDebugger.html) and select the **Test render pass** item.

```
using UnityEngine;
using UnityEngine.Rendering;
using UnityEngine.Rendering.RenderGraphModule;
using UnityEngine.Rendering.RenderGraphModule.Util;
using UnityEngine.Rendering.Universal;
using UnityEngine.Experimental.Rendering;

public class RandomTriangles : ScriptableRendererFeature
{
    class DrawRandomTriangles : ScriptableRenderPass
    {

        RTHandle destinationTexture;

        class PassData
        {
            public TextureHandle destination;
            public Material material;
        }

        public DrawRandomTriangles(RTHandle externalTexture)
        {
            destinationTexture = externalTexture;
        }

        public override void RecordRenderGraph(RenderGraph renderGraph, ContextContainer frameContext)
        {

            using (var builder = renderGraph.AddRasterRenderPass<PassData>("Test render pass", out var passData))
            {
                // Make sure the render graph system keeps the render pass, even if it's not used in the final frame.
                // Don't use this in production code, because it prevents the render graph system from removing the render pass if it's not needed.
                builder.AllowPassCulling(false);

                // Create a material for drawing a mesh.
                passData.material = new Material(Shader.Find("Unlit/Texture"));

                // Point the render graph to the external texture.
                passData.destination = renderGraph.ImportTexture(destinationTexture);
                builder.SetRenderAttachment(passData.destination, 0, AccessFlags.Write);

                builder.SetRenderFunc(static (PassData data, RasterGraphContext context) =>
                    {
                        // Create a triangle mesh.
                        Mesh mesh = new Mesh();

                        Vector3[] vertices = new Vector3[3]
                        {
                            new Vector3(0, 0, 0),
                            new Vector3(1f, 0, 0),
                            new Vector3(0, 1f, 0),
                        };

                        int[] triangles = new int[3]
                        {
                            0, 2, 1,
                        };

                        mesh.vertices = vertices;
                        mesh.triangles = triangles;

                        // Draw the triangle mesh at a random position.
                        Vector3 randomPos = new Vector3(Random.Range(-10, 10), Random.Range(-10, 10), 0);
                        Matrix4x4 trs = Matrix4x4.TRS(randomPos, Quaternion.Euler(0, 0, 0), Vector3.one);
                        context.cmd.DrawMesh(mesh, trs, data.material, 0);
                    });
            }
        }
    }

    DrawRandomTriangles renderPass;
    public RTHandle myRenderTexture;
    public override void Create()
    {
        // Allocate a render texture outside the shader graph
        TextureDesc textureDesc = new TextureDesc(Vector2.one)
        {
            colorFormat = GraphicsFormat.R8G8B8A8_UNorm,
            width = 256,
            height = 256,
            clearBuffer = true,
            clearColor = Color.red,
            name = "My temporary texture"
        };
        RenderingUtils.ReAllocateHandleIfNeeded(ref myRenderTexture, textureDesc, name: "My render texture");

        // Create the render pass and pass in the render texture
        renderPass = new DrawRandomTriangles(myRenderTexture);
        renderPass.renderPassEvent = RenderPassEvent.AfterRenderingPostProcessing;
    }

    public override void AddRenderPasses(ScriptableRenderer renderer, ref RenderingData renderingData)
    {
        renderer.EnqueuePass(renderPass);
    }

    public new void Dispose()
    {
        myRenderTexture.Release();
    }    

}
```

## Additional resources

* [Using the RTHandle system](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.4/manual/rthandle-system-using.html)
* [Textures](https://docs.unity3d.com/Manual/Textures.html)
* [Render Texture assets](https://docs.unity3d.com/Manual/class-RenderTexture.html)
* [Custom Render Texture assets](https://docs.unity3d.com/Manual/class-CustomRenderTexture.html)