# Combine meshes using batching

Resources and approaches for improving performance by combining static **GameObjects** or moving GameObjects into fewer draw calls.

| **Page** | **Description** |
| --- | --- |
| [Introduction to batching](DrawCallBatching.md) | Understand how Unity creates batches of static and dynamic GameObjects to reduce draw calls. |
| [Set up GameObjects for batching](DrawCallBatching-SetUp.md) | Get the best results from batching, and make sure Unity batches GameObjects. |
| [Batch static meshes at build time](DrawCallBatching-Enable.md) | Use **static batching** to combine static meshes at build time. |
| [Batch meshes at runtime](static-batching-enable.md) | Batch static objects with the `StaticBatchingUtility` API, or dynamic GameObjects with **dynamic batching**. |
| [Combine meshes manually](combining-meshes.md) | Merge multiple meshes into a single **mesh** that Unity can render in a single draw call. |
| [Access properties in combined meshes](DrawCallBatching-Properties.md) | Use MaterialPropertyBlocks to change properties of combined meshes without breaking batching. |

## Additional resources

* [Choose a method for optimizing draw calls](optimizing-draw-calls-choose-method.md)
* [Reduce rendering work on the CPU or GPU](OptimizingGraphicsPerformance.md)