# Use the layer collision matrix to reduce overlaps

Reduce **collision** calculation overhead by configuring interaction rules between **GameObjects** with collision layers.

Use the **Layer Collision Matrix** to control how objects on different layers interact with each other in the physics simulation. You can assign objects to a layer in their **Inspector**, and then use the **Layer Collision Matrix** to define the interaction rules between them. Defining how objects interact on a layer basis can reduce the number of overlap checks and collision calculations performed each frame.

The benefits of layer-based filtering are:

* Reduced broad phase workload: The physics system skips entire groups of objects when calculating potential collisions.
* Improved narrow phase efficiency: By ignoring irrelevant interactions, the simulation can focus computational resources on meaningful collisions.
* Better scalability: Helps maintain performance as the number of objects and complexity in your **scene** increases.

To begin optimizing with collision layers, first define and assign layers to your GameObjects. Then, configure how these layers interact in the **Layer Collision Matrix**. This matrix directly tells the physics system which layers must ignore each other during collision checks, which effectively culls potential interactions early in the simulation.

The recommended best practices for layer-based **collision detection** are:

* Enable interactions only between layers that are absolutely necessary for gameplay.
* Revisit the collision matrix periodically as your project grows. Disable unused layer pairs to increase performance.
* Use clear and consistent naming conventions for your layers to simplify debugging and future maintenance.
* Use the [Profiler](Profiler.md) to monitor the effect of your changes. In the [Physics Profiler module](ProfilerPhysics.md), observe counters related to broad phase pairs, narrow phase checks, and active contacts. Reducing unnecessary layer interactions leads to a decrease in these values, especially the broad phase work.

To learn more, refer to [Layer-based collision detection](LayerBasedCollision.md).

To define and assign layers to your GameObjects and configure the Layer Collision Matrix:

1. Select **Edit > Project Settings** to open the Project Settings window.
2. Select the **Tags and Layers** tab.
3. In the **Layers** section, define custom layers under the **User Layers** section. For example, you can create a “Player” “Environment”, or “Projectiles” layer.
4. In the Editor, assign a GameObject to a layer in the **Inspector** window. To learn more, refer to [Layers](Layers.md).
5. Select \*\*Edit > Project Settings.
6. Select the **Physics > Settings** tab.
7. Select the **Shared** tab.
8. In the **Layer Collision Matrix**, uncheck any layer pairs that you do not want to interact. For example, you might not want projectiles to collide with other projectiles.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [Layer-based collision detection](LayerBasedCollision.md)
* [Layers](Layers.md)