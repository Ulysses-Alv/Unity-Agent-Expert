# Optimize collision callbacks

Reduce garbage collector (GC) load and improve performance by optimizing physics **collision** callback allocations.

Unity provides physics collision callback messages like:

* [`OnCollisionEnter`](../ScriptReference/Collider.OnCollisionEnter.md)
* [`OnCollisionStay`](../ScriptReference/Collider.OnCollisionStay.md)
* [`OnCollisionExit`](../ScriptReference/Collider.OnCollisionExit.md)

By default, each time these callbacks are invoked, Unity allocates new `Collision` (or `Collider` for triggers) data objects in memory to hold information about the collision. These frequent small allocations can increase the load on the garbage collector.

To reduce the number of allocations from these callbacks, you can enable the **Reuse Collision Callbacks** setting, to make Unity reuse a smaller pool of `Collision` data instances. This reduces memory overhead and improves performance, especially in scenarios with many collision events per frame. Note that if you reuse collision callbacks, the `Collision` data passed to the callbacks is then only valid during that specific callback.

To learn more about garbage collection, refer to [Garbage collector overview](performance-garbage-collector.md).

By default, **Reuse Collision Callbacks** is enabled in the **Project Settings**. To make sure that you are reusing collision callbacks and reducing GC allocations from these callbacks in the Editor:

1. Go to **Edit > Project Settings** to open the Project Settings window.
2. Select the **Physics > Settings** tab.
3. Select the ****GameObject**** tab.
4. Make sure that **Reuse Collision Callbacks** is enabled.

To reduce GC allocations from these callbacks in script, set `Physics.reuseCollisionCallbacks = true`.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [Garbage collector overview](performance-garbage-collector.md)