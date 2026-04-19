# Introduction to primitive collider shapes

Primitive colliders are the most computationally efficient type of **collider** in Unity. They are called “primitive” because they are defined by simple geometric shapes such as boxes, spheres, and capsules. They match the [Primitive Objects](PrimitiveObjects.md), which are built-in **GameObject** shapes.

There are three primitive collider shapes in Unity:

* [Box collider](class-BoxCollider.md): A rectangular box-shaped collider that is suitable for most rectangular objects.
* [Sphere collider](class-SphereCollider.md): A spherical collider that is suitable for most circular objects.
* [Capsule collider](class-CapsuleCollider.md): A cylindrical collider that is suitable for objects that have a cylindrical shape, or for rounding out the **collision** edges on sharp corners. Capsule colliders are also useful for player and non-player characters.

Primitive colliders are efficient, but they have limitations. For example, you cannot change or deform their shape, only their scale. Unlike **Mesh** colliders, they are not made up of triangles; their shape is fixed (note that the [Physics Debug visualization](PhysicsDebugVisualization.md) does indicate a triangle-based mesh on primitive colliders, but these are for visualization purposes only and do not reflect the collider’s construction).

Primitive colliders are usually not the best option for complex shapes, wheels, or **Terrain** collision. Unity has specific colliders available for more complex collider shapes (see [Mesh colliders](mesh-colliders.md), [Terrain colliders](terrain-colliders.md), and [Wheel colliders](wheel-colliders.md)). However, if you need a collider shape that fits a complex shape, but does not need to be too accurate, you can use primitive colliders to create a [Compound collider](compound-colliders.md). A Compound collider is a collection of primitive colliders in an arrangement that fits the collider shape you need.

In summary, primitive colliders are an efficient but sometimes inaccurate way to add **collision detection** to objects in Unity. They are suitable for primitive-shaped GameObjects with regular shapes. However, for more complex objects with irregular shapes, or for more accurate collision detection, you should use a more complex collider shape.