# Batch static meshes at build time

To batch static **GameObjects** at build time, enable **static batching**. For more information about batching, refer to [Introduction to batching meshes](DrawCallBatching.md).

Follow these steps:

1. Go to **Edit** > **Project Settings** > **Player**.
2. In **Other Settings**, enable **Static Batching**.
3. In the **Scene** view or **Hierarchy** window, select the GameObject that you want to batch. You can select multiple GameObjects at the same time to enable static batching for all of them.
4. In the ****Inspector**** window, select the [Static Editor Flags](StaticObjects.md) dropdown and enable **Batching Static**.

If you use static batching at build time, Unity doesn’t use any CPU resources at runtime to generate the **mesh** data for the static batch.

**Note:** Unity uses additional GPU memory to store the combined meshes, even if the meshes are the same. For example, marking trees as static in a dense forest environment can have a large impact on memory. If static batching uses too much memory, [combine meshes manually](combining-meshes.md) instead.

If a mesh is culled because it’s outside the **camera** view, Unity might not be able to draw all the meshes in a single draw call. This is because the culled mesh leaves a gap in the combined vertex buffers.

## Additional resources

* [Batch meshes at runtime](static-batching-enable.md)