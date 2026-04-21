# Physics Low Level Settings 2D asset Inspector window reference

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics) instead.

Explore the properties and settings you can use to configure the defaults and global behaviour of the `LowLevelPhysics2D` API. For more information, refer to [Configure global 2D physics settings](2d-physics-api/2d-physics-api-global-settings.md).

## Layers

The properties in this section configure the layers Unity uses to detect **collision**. For more information, refer to [Configure collisions between LowLevelPhysics2D API objects](2d-physics-api/2d-physics-api-collisions-enable.md).

| **Property** | **Description** |
| --- | --- |
| **Physics Layer Names** | Sets the names of the 64 layers the `PhysicsMask` API uses if you enable **Use Full Layers**. By default, the first layer has the name **Default** and the other layers have no name. |
| **Use Full Layers** | Enables the `PhysicsMask` API using the 64 layers listed under **Physics Layer Names**. If you disable this property, the `PhysicsMask` API uses the 32 [GameObject layers](Layers.md) instead. |

## Default Definitions

The properties in this section set the default values when you create a new definition object. For the individual sections, refer to the following:

* [World definition reference](2d-physics-api/2d-physics-api-reference-world.md)
* [Body definition reference](2d-physics-api/2d-physics-api-reference-body.md)
* [Shape definition reference](2d-physics-api/2d-physics-api-reference-shape.md)
* [Chain definition reference](2d-physics-api/2d-physics-api-reference-chain.md)
* [Distance joint definition reference](2d-physics-api/2d-physics-api-reference-joint-distance.md)
* [Fixed joint definition reference](2d-physics-api/2d-physics-api-reference-joint-fixed.md)
* [Hinge joint definition reference](2d-physics-api/2d-physics-api-reference-joint-hinge.md)
* [Relative joint definition reference](2d-physics-api/2d-physics-api-reference-joint-relative.md)
* [Slider joint definition reference](2d-physics-api/2d-physics-api-reference-joint-slider.md)
* [Wheel joint definition reference](2d-physics-api/2d-physics-api-reference-joint-wheel.md)

## Globals

| **Property** | **Description** |
| --- | --- |
| **Concurrent Simulations** | Sets the number of physics simulations that can run at the same time. The default is 2. |
| **Length Units Per Meter** | Sets the number of **Unity units** that correspond to one meter in the physics simulation. The default is 1. |
| **Draw In Build** | Draws the debug visualization in a built application. For more information, refer to [Draw a debug visualization of objects](2d-physics-api/2d-physics-api-debug-drawing.md). |
| **Bypass Low Level** | Disables the `LowLevelPhysics2D` API physics system. |

## Additional resources

* [Configure global 2D physics settings](2d-physics-api/2d-physics-api-global-settings.md)
* [Configure 2D physics properties using a definition](2d-physics-api/2d-physics-api-definitions.md)