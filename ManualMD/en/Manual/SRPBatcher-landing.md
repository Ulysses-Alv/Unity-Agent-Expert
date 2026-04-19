# Scriptable Render Pipeline (SRP) Batcher in URP

Resources for using the Scriptable **Render Pipeline** (SRP) Batcher to reduce the number of render state changes between draw calls.

| **Page** | **Description** |
| --- | --- |
| [Introduction to the SRP Batcher](SRPBatcher.md) | Understand how the SRP Batcher reduces render state changes between draw calls. |
| [Check whether a GameObject is compatible with the SRP Batcher](SRPBatcher-Materials.md) | Find out if Unity can include a **GameObject** and a **shader** in the SRP Batcher. |
| [Enable the SRP Batcher](SRPBatcher-Enable.md) | Enable the SRP Batcher in the URP Asset. |
| [Troubleshoot the SRP Batcher](SRPBatcher-Profile.md) | Use the Frame Debugger to solve common issues with the SRP Batcher, such as a low number of draw calls in batches. |
| [Remove SRP Batcher compatibility for GameObjects](SRPBatcher-Incompatible.md) | Make a shader or renderer incompatible with the SRP Batcher, for example if you want to use [GPU instancing](GPUInstancing.md). |

## Additional resources

* [Choose a method for optimizing draw calls](optimizing-draw-calls-choose-method.md)
* [GPU instancing](GPUInstancing.md)