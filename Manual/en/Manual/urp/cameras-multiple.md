# Multiple cameras in URP

![An example of the effect camera stacking can produce in URP](../../uploads/urp/camera-stacking-example.png)

An example of the effect camera stacking can produce in URP

Resources and approaches for using multiple cameras to work with multiple **camera** outputs and targets in the Universal **Render Pipeline** (URP).

**Note:** If you use multiple cameras, it might make rendering slower. An active camera runs through the entire rendering loop even if it renders nothing.

| Page | Description |
| --- | --- |
| [Camera stacking](cameras/camera-stacking-concepts.md) | Learn about the fundamental concepts of camera stacking. |
| [Set up a camera stack](camera-stacking.md) | Stack cameras to layer the outputs of multiple cameras into a single combined output. |
| [Add and remove cameras in a camera stack](cameras/add-and-remove-cameras-in-a-stack.md) | Add, remove, and reorder cameras within a camera stack. |
| [Set up split-screen rendering](rendering-to-the-same-render-target.md) | Render multiple camera outputs to a single render target to create effects such as split screen rendering. |
| [Apply different post processing effects to separate cameras](cameras/apply-different-post-proc-to-cameras.md) | Apply different **post-processing** setups to individual cameas within a **scene**. |
| [Render a camera’s output to a Render Texture](rendering-to-a-render-texture.md) | Render to a **Render Texture** to create effects such as in-game CCTV monitors. |
| [Create a render request](User-Render-Requests.md) | Trigger a camera to render to a Render Texture outside of URP rendering loop. |

## Additional resources

* [Multiple cameras](../MultipleCameras-landing.md)
* [Camera render types in URP](camera-types-and-render-type.md)
* [Camera render order in URP](cameras-advanced.md)