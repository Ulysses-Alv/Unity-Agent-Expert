# Adjust Rigidbody component solver iterations

Adjust solver iteration counts for individual ****Rigidbody**** components to improve simulation results.

If a specific **Rigidbody** component in your **scene** requires higher simulation accuracy or stability, such as preventing jittering in stacks or ensure **joint** stability, adjust its solver iteration count individually. Increasing solver iterations uses more CPU time for that specific **Rigidbody** component. This targeted approach is more efficient than increasing the global solver iteration count, which would affect all **Rigidbody** components.

To adjust the solver iterations for a specific **Rigidbody** component, use [`Rigidbody.solverIterations`](../ScriptReference/Rigidbody-solverIterations.md) on that specific component. This setting overrides the global project setting for that **Rigidbody** component only.

To configure the global setting for all **Rigidbody** components in your project:

1. Select **Edit > Project Settings** to open the Project Settings window.
2. Select the **Physics > Settings** tab.
3. Select the ****GameObject**** tab.
4. Adjust the **Default Solver Iterations** value.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [`Rigidbody.solverIterations`](../ScriptReference/Rigidbody-solverIterations.md)