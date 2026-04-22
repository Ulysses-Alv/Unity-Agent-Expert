# Introduction to Mesh colliders

[Mesh colliders](class-MeshCollider.md) match the shape of a [Mesh](class-Mesh.md) exactly, for extremely accurate **collision** simulation.

The Mesh **collider** builds its collision geometry to match an assigned [Mesh](class-Mesh.md), including its shape, position and scale. The benefit of this is that you can make the shape of the collider exactly the same as the shape of the visible Mesh for the **GameObject**, which creates more precise and realistic collisions.

The precision of a Mesh collider comes with a higher processing overhead than [primitive colliders](primitive-colliders.md) (such as the [Sphere](class-SphereCollider.md), [Box](class-BoxCollider.md), and [Capsule](class-CapsuleCollider.md) colliders). For this reason, it is best practice to only use Mesh colliders for colliders that do not otherwise require a high amount of processing power, or for collisions where primitive colliders or a [compound collider](compound-colliders.md) would have a greater overhead. A good approach is to use Mesh colliders for static **scene** geometry, and compound colliders for moving GameObjects.

## Benefits and limitations of Mesh colliders

In most cases, Mesh colliders offer a similar solution to [compound colliders](compound-colliders.md): their primary purpose is to provide accurate collisions for items with complex shapes. When considering the benefits and limitations of Mesh colliders, you are usually comparing them to compound colliders.

The main benefits of Mesh colliders are:

* Mesh colliders are extremely accurate, because they perfectly match the shape of the item.
* Mesh colliders require less manual development time than compound colliders, because Unity handles their shape automatically based on the shape of the Mesh.
* In some cases, Mesh colliders can be less computationally demanding than compound colliders. This is often the case for very complex shapes which would require a high number of compound colliders to approximate. One Mesh collider might be more efficient than several primitive colliders.

However, Mesh colliders also have some significant limitations:

* Mesh colliders cannot provide accurate collisions between concave shapes. For details, see [Concave and convex Mesh colliders](#concave-convex-mesh-colliders), below.
* In some cases, compound colliders can be less computationally demanding than Mesh colliders. This is often the case for simpler shapes which only require a few colliders to approximate, or for shapes that don’t need to be too accurate. For example, a Mesh collider might generate hundreds of polygons in order to precisely match a complex shape, but an approximation with primitive colliders might require far fewer.
* Mesh colliders are not a good option for Meshes that change shape at runtime. For details, see [Meshes that change shape at runtime](#meshes-that-change-shape-runtime), below.

The decision is always unique to your project, so you should test each configuration and use the Physics Profiler to understand the efficiency of your collider setup.

## Concave and convex Mesh colliders

Mesh colliders behave differently based on whether they are marked as concave or convex. By default, PhysX considers Mesh colliders to be concave.

In mathematics, “concave” and “convex” are terms used to describe shapes:

A convex shape only has lines that curve outward (for example, a ball). Any line segment connecting two points on its boundary remains entirely within the shape.
A concave shape has at least one “cave” or indentation where the boundary curves inward (for example, a banana). Line segments connecting two points on its boundary might sometimes cross space outside the shape.

Concave colliders have some limitations:
Concave Mesh colliders can only be static (that is, they have no physics body) or kinematic (they have a kinematic physics body). See [Collider types](collider-types-introduction.md) for more details.
Concave Mesh colliders can only collide with convex colliders. If two concave colliders make contact, no collision occurs.

If you have two concave Mesh colliders that need to collide, you can do one of the following, depending on how accurate your collision needs to be:

If you do not need accurate collisions to take place in the concave parts of the shape, mark one of the Mesh colliders as convex in the **Inspector** (enable **Is Convex**). This generates a new convex collider shape called a “hull”, which is like the assigned concave Mesh but with any concave lines filled in, to make them convex.

![A concave Mesh collider on a banana-shaped GameObject.](../uploads/Main/food-pack-banana-concave-mesh.png)

A concave Mesh collider on a banana-shaped GameObject.

![A convex Mesh collider on the same banana-shaped GameObject.](../uploads/Main/food-pack-banana-convex-mesh.png)

A convex Mesh collider on the same banana-shaped GameObject.

If you need accurate collisions to take place in the concave parts of the shape, use a [compound collider](compound-colliders.md) made of convex colliders.

### Meshes that change shape at runtime

The Mesh that you have assigned to the Mesh collider should ideally not change shape at runtime.

Every time the Mesh changes shape, the **physics engine** needs to recalculate the Mesh collider geometry, which causes a substantial performance overhead. For this reason, you should not modify the geometry of a Mesh that a Mesh collider is using. If a Mesh needs to both collide and change shape at runtime, it is usually better to approximate the Mesh shape with [primitive colliders](primitive-colliders.md) or [compound colliders](compound-colliders.md).

### Directional mesh facing

Faces in collision meshes are one-sided. This means GameObjects can pass through them from one direction, but collide with them from the other.

For details about the underlying algorithms and data structures that Mesh colliders use, see the Nvidia PhysX documentation on [Geometry](https://docs.nvidia.com/gameworks/content/gameworkslibrary/physx/guide/Manual/Geometry.html).