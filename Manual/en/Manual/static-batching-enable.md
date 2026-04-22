# Batch meshes at runtime

For more information about batching, refer to [Introduction to batching meshes](DrawCallBatching.md).

**Note:** Unity uses additional CPU memory to store the combined meshes, even if the meshes are the same. For example, marking trees as static in a dense forest environment can have a large impact on memory. If **static batching** uses too much memory, [combine meshes manually](combining-meshes.md) instead.

## Batch static GameObjects at runtime

To batch static **GameObjects** at runtime, for example GameObjects you create procedurally, follow these steps:

1. Select each **mesh** you want to batch, then in the ****Inspector**** window enable **Read/Write enabled**.
2. Use the [StaticBatchingUtility](../ScriptReference/StaticBatchingUtility.md) API to batch the meshes. The [StaticBatchingUtility.Combine](../ScriptReference/StaticBatchingUtility.Combine.md) method combines GameObjects and prepares them for static batching.

**Note:** To avoid the read/write mesh using double the amount of memory, you can now delete one copy of the mesh by calling the [`Mesh.UploadMeshData`](../ScriptReference/Mesh.UploadMeshData.md) API with `markNoLongerReadable` set to `true`.

You don’t need to enable the static batching setting in the URP Asset, HDRP Asset, or ****Player Settings****.

If a mesh is culled because it’s outside the **camera** view, Unity might split the batch into two batches, to avoid a gap in the combined vertex buffers.

You can only change the position, rotation, and scale of the whole batch at runtime, not the individual meshes.

## Batch moving GameObjects at runtime

To batch moving GameObjects, enable **dynamic batching**.

**Warning:** For most uses, dynamic batching is no longer recommended, because the CPU overhead can be greater than the overhead of a draw call. If you use URP or HDRP, use the [Scriptable Render Pipeline (SRP) Batcher](SRPBatcher.md) instead. For more information, refer to [Choose a draw call optimization method](optimizing-draw-calls-choose-method.md).

Follow these steps in the Built-In **Render Pipeline**:

1. Go to **Edit** > **Project Settings** > **Player**.
2. In the **Other Settings** section, enable **Dynamic Batching**.

Dynamic batching transforms all geometry into world space. To disable dynamic batching in a custom **shader** so you can access object space, enable **Disable batching** in the [Shader Import Settings](class-ShaderImporter.md) window of the shader.

You can also disable batching by setting the `DisableBatching` tag to `True` in the subshader block of a shader. For example:

```
    SubShader {
        Tags { "DisableBatching" = "True" }
        ...
    }
```

## Additional resources

* [`SubShader` tags in ShaderLab reference](SL-SubShaderTags.md)