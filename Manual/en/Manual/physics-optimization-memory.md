# Optimize the physics system for memory usage

You can optimize how the Unity physics system uses memory resources by managing **collision** callback allocations and optimizing various physics queries. Efficient memory usage helps reduce garbage collection overhead and ensures smoother gameplay.

Use the guidance in these pages to maintain your target frame rate and ensure smooth, responsive gameplay. The instructions in these pages address issues identified by Unity Editor diagnostic tools. Before you apply these optimizations described in the documentation in this section and throughout your development, you must be familiar with these diagnostic tools:

* [**The Unity Profiler**](Profiler.md)
* [**The Memory Profiler**](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [**The Physics Debug window**](PhysicsDebugVisualization.md)

| **Topic** | **Description** |
| --- | --- |
| **[Optimize collision callbacks](physics-optimization-collision-callbacks.md)** | Reduce memory allocations caused by frequent collision events. |
| **[Optimize raycasts and other physics queries](physics-optimization-raycasts-queries.md)** | Optimize physics query performance and reduce garbage collection overhead by using efficient query versions and batch processing. |

## Additional resources

* [Unity Profiler](Profiler.md)
* [Built-in 3D Physics](PhysicsOverview.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)