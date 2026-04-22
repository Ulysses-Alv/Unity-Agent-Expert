# Move static colliders to prevent performance issues

Properly manage static **colliders** when they move to avoid performance issues.

A [Static collider](collider-types-introduction.md#static-colliders) is a **GameObject** with a Collider component but no **Rigidbody** or ArticulationBody component attached. You can use static colliders for objects that don’t move during gameplay, such as **terrain**, buildings, or other environmental features.

When you move a static collider by changing its transform values, the physics system detects the change and updates its internal spatial structures during the next physics step or when [`Physics.SyncTransforms`](../ScriptReference/Physics.SyncTransforms.md) is called. If you want to make frequent changes to the transform values of a static collider between physics simulations steps when you execute gameplay code, use a Kinematic Rigidbody component instead.

If you want to move a static collider, the recommended best practice is that you don’t add a Rigidbody component to a static object solely to move that GameObject. If it doesn’t need a physics simulation, you’re adding an unnecessary performance burden.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [Introduction to collider types](collider-types-introduction.md)