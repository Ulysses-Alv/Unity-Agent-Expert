# Optimize transform value syncing

Optimize the synchronization of Transform values with the physics system to improve performance and query accuracy.

By default, Unity defers physics transform syncing. If you change a **Transform** value and then immediately query the physics world, the physics system might not become aware of the change unless it’s explicitly communicated. You can control this behavior by enabling or disabling **Auto Sync Transforms**.

The recommended best practice is to disable **Auto Sync Transforms**, which is disabled by default. If you modify a **Rigidbody** or a **Collider** component’s transform values directly and then immediately need to perform a physics query that depends on that object’s new position in the same frame, manually call [`Physics.SyncTransforms`](../ScriptReference/Physics.SyncTransforms.md) before the query. This ensures the physics world is up-to-date with the transform values changes for accurate query results. **Note**: `Physics.SyncTransforms` is crucial for accurate queries when **Auto Sync Transforms** is disabled, especially if `Physics.simulationMode` is set to `Script`, and you’re querying as detailed in [Optimizing physics for query-only or non-simulating games](physics-optimization-cpu-query-only.md).

To enable or disable **Auto Sync Transforms** in the Editor:

1. Select **Edit > Project Settings** to open the Project Settings window.
2. Select the **Physics > Settings** tab.
3. Select the ****GameObject**** tab.
4. Enable or disable **Auto Sync Transforms**. By default, **Auto Sync Transforms** is not enabled. If you enable it, you add implicit synchronization points before every physics query, which can negatively affect performance.

To enable or disable **Auto Sync Transforms** in script, set [`Physics.autoSyncTransforms`](../ScriptReference/Physics-autoSyncTransforms.md) to `true` or `false`. Setting `Physics.autoSyncTransforms = true;` has the same effect as enabling **Auto Sync Transforms** in the **Project Settings**.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [Optimize physics for query-only or non-simulating games](physics-optimization-cpu-query-only.md)