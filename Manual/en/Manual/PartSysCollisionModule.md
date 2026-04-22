# Collision module reference

This module controls how [particles](class-ParticleSystem.md) collide with [GameObjects](GameObjects.md) in the [Scene](CreatingScenes.md).

| **Property** | **Function** |
| --- | --- |
| **Type** | Use the **Type**\* drop-down to define whether your **collision** settings apply to **Planes** or to the **World**. If you choose **World**, use the **Collision Mode** drop-down to define whether your collision settings apply for a [2D or 3D](2Dor3D.md) world. |

## Planes module properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| **Property** | **Function** |
| --- | --- |
| **Planes popup** | Select **Planes** mode. |
| **Planes** | An expandable list of Transforms that define collision planes. |
| **Visualization** | Selects whether the collision plane **Gizmos** will be shown in the **Scene view** as wireframe grids or solid planes. |
| **Scale Plane** | Size of planes used for visualization. |
| **Dampen** | The fraction of a particle’s speed that it loses after a collision. |
| **Bounce** | The fraction of a particle’s speed that rebounds from a surface after a collision. |
| **Lifetime Loss** | The fraction of a particle’s total lifetime that it loses if it collides. |
| **Min Kill Speed** | Particles travelling below this speed after a collision will be removed from the system. |
| **Max Kill Speed** | Particles travelling above this speed after a collision will be removed from the system. |
| **Radius Scale** | Allows you to adjust the radius of the particle collision spheres so it more closely fits the visual edges of the particle graphic. |
| **Send Collision Messages** | If enabled, particle collisions can be detected from **scripts** by the [OnParticleCollision](../ScriptReference/MonoBehaviour.OnParticleCollision.md) function. |
| **Visualize Bounds** | Renders the collision bounds of each particle as a wireframe shape in the Scene view. |

## World module properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| ***Property*** | ***Function*** |
| --- | --- |
| **World popup** | Select **World** mode. |
| **Collision Mode** | 3D or 2D. |
| **Dampen** | The fraction of a particle’s speed that it loses after a collision. |
| **Bounce** | The fraction of a particle’s speed that rebounds from a surface after a collision. |
| **Lifetime Loss** | The fraction of a particle’s total lifetime that it loses if it collides. |
| **Min Kill Speed** | Particles travelling below this speed after a collision will be removed from the system. |
| **Max Kill Speed** | Particles travelling above this speed after a collision will be removed from the system. |
| **Radius Scale** | Setting for 2D or 3D. |
| **Collision Quality** | Use the drop-down to set the quality of particle collisions. This affects how many particles can pass through a collider. At lower quality levels, particles can sometimes pass through colliders, but are less resource-intensive to calculate. |
| High | When **Collision Quality** is set to **High**, collisions always use the physics system for detecting the collision results. This is the most resource-intensive option, but also the most accurate. |
| Medium (Static Colliders) | When **Collision Quality** is set to **Medium (Static Colliders)**, collisions use a grid of voxels to cache previous collisions, for faster re-use in later frames. See [Particle collisions: World collisions](particle-collisions.md#WorldCollisions), below, to learn more about this cache.   The only difference between **Medium** and **Low** is how many times per frame the Particle System queries the physics system. **Medium** makes more queries per frame than **Low**.   Note that this setting is only suitable for static colliders that never move. |
| Low (Static Colliders) | When **Collision Quality** is set to **Low (Static Colliders)**, collisions use a grid of voxels to cache previous collisions, for faster re-use in later frames. See [Particle collisions: World collisions](particle-collisions.md#WorldCollisions), below, to learn more about this cache.   The only difference between **Medium** and **Low** is how many times per frame the Particle System queries the physics system. **Medium** makes more queries per frame than **Low**.   Note that this setting is only suitable for static colliders that never move. |
| **Collides With** | Particles will only collide with objects on the selected layers. |
| **Max Collision Shapes** | How many collision shapes can be considered for particle collisions. Excess shapes are ignored, and **terrains** take priority. |
| **Enable Dynamic Colliders** | Dynamic colliders are any collider not configured as Kinematic (see documentation on [colliders](CollidersOverview.md) for further information on collider types).   Check this option to include these collider types in the set of objects that the particles respond to in collisions. If you uncheck this option, the particles only respond to collisions against static colliders. |
| **Voxel Size** | A voxel represents a value on a regular grid in three-dimensional space. When using **Medium** or **Low** quality collisions, Unity caches collisions in a grid structure. This setting controls the grid size. Smaller values give more accuracy, but cost more memory, and are less efficient.   **Note**: You can only access this property when **Collision Quality** is set to **Medium** or **Low**. |
| **Collider Force** | Apply a force to Physics Colliders after a Particle collision. This is useful for pushing colliders with particles. |
| **Multiply by Collision Angle** | When applying forces to Colliders, scale the strength of the force based on the collision angle between the particle and the collider. Grazing angles will generate less force than a head-on collision. |
| **Multiply by Particle Speed** | When applying forces to Colliders, scale the strength of the force based on the speed of the particle. Fast-moving particles will generate more force than slower ones. |
| **Multiply by Particle Size** | When applying forces to Colliders, scale the strength of the force based on the size of the particle. Larger particles will generate more force than smaller ones. |
| **Send Collision Messages** | Check this to be able to detect particle collisions from scripts by the [OnParticleCollision](../ScriptReference/MonoBehaviour.OnParticleCollision.md) function. |
| **Visualize Bounds** | Preview the collision spheres for each particle in the Scene view. |

## Additional resources

* [Particle collisions](particle-collisions.md)
* [Particle triggers](particle-triggers.md)
* [Configure a particle trigger](configure-particle-trigger.md)