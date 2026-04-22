# Optimize physics for query-only or non-simulating games

Prevent the default physics update loop from running to reduce unnecessary performance overhead by changing the **Simulation Mode** to **Script**.

Preventing the default physics update loop from running can be beneficial if your game doesn’t require physics simulation for the default physics **scene**. For example, your game might only require physics queries such as raycasts, spherecasts, or overlap checks, but not **Rigidbody** components with forces, gravity, or physics-driven movement.

This approach uses physics queries for **collision** detection or environmental interaction without the performance cost of a full physics simulation, and it maintains control over when transform data is synchronized.

Note the following if you set the **Simulation Mode** of the default scene to **Script**:

* Physics queries still function by checking against the physics world’s representation of **colliders**.
* If you have objects in your scene that are moved by directly manipulating their `Transform` components, their positions within the **physics engine**’s internal spatial data structures are not automatically updated to match their `Transform` positions frame-to-frame (since no simulation step is occurring by default in this mode).
* To ensure raycasts and other queries against these colliders are accurate, assuming `Physics.autoSyncTransforms` is `false` (its default and recommended setting), you must manually call `Physics.SyncTransforms` once per frame before performing your queries. This function explicitly updates the physics system with the current state of all the transform values of **GameObjects** that have colliders.
* If `Physics.autoSyncTransforms` is set to `true`, this synchronization happens implicitly before each query, but this is generally not the recommended best practice for the potential overhead of syncing before every query. Note that `Physics.autoSyncTransforms` is deprecated. To learn how to sync transforms, refer to [Optimize transform value syncing](physics-optimization-cpu-transform-sync.md).
* Without an explicit `Physics.SyncTransforms` call (when `autoSyncTransforms` is `false`), queries might use outdated object positions and lead to incorrect results.

To prevent the default physics update from running in the Editor:

1. Select **Edit > Project Settings** to open the Project Settings window.
2. Select the **Physics > Settings** tab.
3. Select the **GameObject** tab.
4. Set **Simulation Mode** to **Script**.

To disable automatic simulation in script, set `Physics.simulationMode = SimulationMode.Script;` at runtime.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [Introduction to collider types](collider-types-introduction.md)
* [Optimize transform value syncing](physics-optimization-cpu-transform-sync.md)