# Particle collisions

The [Collisions](PartSysCollisionModule.md) module controls how **particles** collide with **GameObjects** in the **Scene**.

When other objects surround a **Particle System**, the effect is often more convincing when the particles interact with those objects. For example, water or debris should be obstructed by a solid wall rather than simply passing through it. With the **Collision** module enabled, particles can collide with objects in the Scene.

A Particle System can be set so its particles collide with any **Collider** in the scene by selecting **World** mode from the pop-up. Colliders can also be disabled according to the layer they are on by using the **Collides With** property. The pop-up also has a **Planes** mode option which allows you to add a set of planes to the Scene that don’t need to have Colliders. This option is useful for simple floors, walls and similar objects, and has a lower processor overhead than **World** mode.

When **Planes** mode is enabled, a list of transforms (typically empty GameObjects) can be added via the **Planes** property. The planes extend infinitely in the objects’ local XZ planes, with the positive Y axis indicating the planes’ normal vectors. To assist with development, the planes will be shown as **Gizmos** in the Scene, regardless of whether or not the objects have any visible **Mesh** themselves. The Gizmos can be shown as a wireframe grid or a solid plane, and can also be scaled. However, the scaling only applies to the visualization - the collision planes themselves extend infinitely through the Scene.

When collisions are enabled, the size of a particle is sometimes a problem because its graphic can be clipped as it makes contact with a surface. This can result in a particle appearing to “sink” partway into a surface before stopping or bouncing. The **Radius Scale** property addresses this issue by defining an approximate circular radius for the particles, as a percentage of its actual size. This size information is used to prevent clipping and avoid the sinking-in effect.

The **Dampen** and **Bounce** properties are useful for when the particles represent solid objects. For example, gravel will tend to bounce off a hard surface when thrown but a snowball’s particles might lose speed during a collision. **Lifetime Loss** and **Min Kill Speed** can help to reduce the effects of residual particles following a collision. For example, a fireball might last for a few seconds while flying through the air but after colliding, the separate fire particles should dissipate quickly.

You can also detect particle collisions from a script if **Send Collision Messages** is enabled. The script can be attached to the object with the particle system, or the one with the Collider, or both. By detecting collisions, you can use particles as active objects in gameplay, for example as projectiles, magic spells and power-ups. See the script reference page for [MonoBehaviour.OnParticleCollision](../ScriptReference/MonoBehaviour.OnParticleCollision.md) for further details and examples.

### World Collision Quality

The World Collision module has a **Collision Quality** property, which you can set to **High**, **Medium** or **Low**. When **Collision Quality** is set to **Medium (Static Colliders)** or **Low (Static Colliders)**, collisions use a grid of voxels (values on a 3D grid) to cache previous collisions, for fast re-use in later frames.

This cache consists of a plane in each voxel, where the plane represents the collision surface at that location. On each frame, Unity checks the cache for a plane at the position of the particle, and if there is one, Unity uses it for **collision detection**. Otherwise, it asks the physics system. If a collision is returned, it is added to the cache for fast querying on subsequent frames.

This is an approximation, so some missed collisions might occur. You can reduce the Voxel Size value to help with this; however, doing so uses extra memory, and is less efficient.

The only difference between **Medium** and **Low** is how many times per frame the system is allowed to query the physics system. Low makes fewer queries per frame than Medium. Once the per-frame budget has been exceeded, only the cache is used for any remaining particles. This can lead to an increase in missed collisions, until the cache has been more comprehensively populated.

## Additional resources

* [Quality settings tab reference](class-QualitySettings.md#particles)