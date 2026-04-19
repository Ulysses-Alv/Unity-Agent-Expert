# Optimize the physics system for CPU usage

You can optimize how the Unity physics system uses CPU resources in several ways. For example, you can adjust simulation frequency, carefully manage **collider** types, configure **Rigidbody** component behaviors, and more. Effective CPU optimization helps ensure your game maintains a high frame rate and responsive physics interactions.

Use the guidance in these pages to maintain your target frame rate and ensure smooth, responsive gameplay. The instructions in these pages address issues identified by Unity Editor diagnostic tools. Before you apply these optimizations described in the documentation in this section and throughout your development, you must be familiar with these diagnostic tools:

* [**The Unity Profiler**](Profiler.md)
* [**The Memory Profiler**](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [**The Physics Debug window**](PhysicsDebugVisualization.md)

| **Topic** | **Description** |
| --- | --- |
| **[Set fixed timestep to optimize physics simulation frequency](physics-optimization-cpu-frequency.md)** | Configure the fixed time step and manage potential performance spirals to control physics update frequency. |
| **[Manually set physics simulation](physics-optimization-cpu-manual-simulation.md)** | Control over when physics calculations occur to align them with game performance. |
| **[Optimize physics for query-only or non-simulating games](physics-optimization-cpu-query-only.md)** | Prevent the default physics update loop from running to reduce unnecessary performance overhead. |
| **[Optimize transform value syncing](physics-optimization-cpu-transform-sync.md)** | Optimize the synchronization of Transform values with the physics system to improve performance and query accuracy. |
| **[Move static colliders](physics-optimization-cpu-static-colliders.md)** | Understand best practices for moving static colliders and when to use Kinematic **Rigidbody** components instead. |
| **[Use the layer collision matrix to reduce overlaps](physics-optimization-cpu-collision-layers.md)** | Reduce **collision** calculation overhead by configuring interaction rules between **GameObjects** with collision layers. |
| **[Select a broad phase pruning algorithm](physics-optimization-cpu-broad-phase.md)** | Optimize physics performance in large **scenes** by selecting the most efficient broad phase pruning algorithm. |
| **[Collider types and performance](physics-optimization-cpu-collider-types.md)** | Select the most efficient collider types for different GameObjects. |
| **[Configure Mesh Collider component cooking options for optimization](physics-optimization-cpu-mesh-cooking-options.md)** | Optimize physics calculations by configuring cooking options for ****Mesh** Collider** components |
| **[Use Rigidbody sleeping to improve physics performance](physics-optimization-cpu-rigidbody-sleeping.md)** | Reduce CPU load and improve physics performance by enabling **Rigidbody** sleeping for stationary objects. |
| **[Adjust Rigidbody component solver iterations](physics-optimization-cpu-rigidbody-solver.md)** | Adjust solver iteration counts for a **Rigidbody** component to improve simulation accuracy. |
| **[Optimize Rigidbody component collision detection modes](physics-optimization-cpu-rigidbody-collision-modes.md)** | Balance collision accuracy and CPU performance by choosing appropriate detection modes for **Rigidbody** components. |

## Additional resources

* [Unity Profiler](Profiler.md)
* [Built-in 3D Physics](PhysicsOverview.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)