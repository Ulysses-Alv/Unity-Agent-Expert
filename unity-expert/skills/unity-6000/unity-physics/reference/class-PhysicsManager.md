# Physics

To apply global settings for 3D physics, use the **Physics** settings (main menu: **Edit** > **Project Settings**, then select the **Physics** category).

**Note:** To manage global settings for 2D physics, use the [Physics 2D](class-Physics2DSettings.md) settings instead.

The Physics settings define limits on the accuracy of the 3D physical simulation. In most cases, a more accurate simulation requires more processing overhead, so these settings offer a way to trade off accuracy against performance.

The **Physics** panel contains settings to define which Physics SDKs you want to use in your project.

![Physics Inspector settings.](../uploads/Main/Physics_gameObject_sdk.png)

Physics Inspector settings.

| **Setting** | **Description** |
| --- | --- |
| ****GameObject** SDK** | Define the SDK you want to use to manage physics simulation on GameObjects in your project. The default setting is **PhysX**. Set to **None** to disable physics simulation for GameObjects in your project. |

The Physics **project settings** section also contains a panel called **Settings**.

## Settings

The Settings panel contain settings for all Physics SDKs. The following tabs are available:

* **Shared**: Settings that apply to all Physics SDKs.
* **GameObject**: Settings that apply to all GameObject physics in your project.
* **Cloth**: Settings that apply to all [Cloth](class-Cloth.md) physics.

### Shared settings

The **Shared** tab contains global settings that apply to all physics SDKs.

![Shared Physics Settings window.](../uploads/Main/Physics_shared_settings.png)

Shared Physics Settings window.

| **Setting** | **Description** |
| --- | --- |
| **Gravity** | Use the x, y and z axes to set the amount of gravity applied to all Rigidbody components. For realistic gravity settings, apply a negative number to the y axis. Gravity is defined in world units per seconds squared.  **Note**: If you increase the gravity, you might need to also increase the **Default Solver Iterations** value to maintain stable contacts. |
| **Layer **Collision** Matrix** | Define how the [layer-based collision](LayerBasedCollision.md) detection system behaves. To select which layers on the Collision Matrix interact with the other layers, check their respective checkboxes. |

### GameObject settings

The **GameObject** tab contains global physics settings that apply to all GameObjects in your project.

![The Physics Settings window.](../uploads/Main/Physics_gameObject_settings.png)

The Physics Settings window.

| **Setting** | **Description** |
| --- | --- |
| **Default Material** | Set a reference to the default **Physics Material** to use if none has been assigned to an individual **Collider**. |
| **Bounce Threshold** | Set a velocity value. If two colliding objects have a relative velocity below this value, they do not bounce off each other. This value also reduces jitter, so it is not recommended to set it to a very low value. |
| **Default Max Depenetration Velocity** | Define the default value for the maximum depenetration velocity (the velocity that the solver can set to a body while trying to pull it out of overlap with the other bodies). |
| **Sleep Threshold** | Set a global energy threshold, below which a non-kinematic **Rigidbody** (that is, one that is not controlled by the physics system) may go to sleep. When a Rigidbody is sleeping, it is not updated every frame, making it less resource-intensive. If a Rigidbody’s kinetic energy divided by its mass is below this threshold, it is a candidate for sleeping. |
| **Default Contact Offset** | Set the distance the **collision detection** system uses to generate collision contacts. The value must be positive, and if set too close to zero, it can cause jitter. This is set to 0.01 by default. Colliders only generate collision contacts if their distance is less than the sum of their contact offset values. |
| **Default Solver Iterations** | Define how many solver processes Unity runs on every physics frame. Solvers are small physics engine tasks which determine a number of physics interactions, such as the movements of joints or managing contact between overlapping Rigidbody components.  This affects the quality of the solver output, and it is advisable to change the property in case non-default `Time.fixedDeltaTime` is used, or the configuration is extra demanding. Typically, it is used to reduce the jitter resulting from joints or contacts. |
| **Default Solver Velocity Iterations** | Set how many velocity processes a solver performs in each physics frame. The more processes the solver performs, the higher the accuracy of the resulting exit velocity after a Rigidbody bounce. If you experience problems with jointed Rigidbody components or Ragdolls moving too much after collisions, try increasing this value. |
| **Queries Hit Backfaces** | Enable this option if you want physics queries (such as `Physics.Raycast`) to detect hits with the backface triangles of MeshColliders. By default, this setting is disabled. |
| **Queries Hit Triggers** | Enable this option if you want physics hit tests (such as Raycasts, SphereCasts and SphereTests) to return a hit when they intersect with a Collider marked as a Trigger. Individual raycasts can override this behavior. By default, this setting is enabled. |
| **Enable Adaptive Force** | Enable this option to enable the adaptive force. The adaptive force affects the way forces are transmitted through a pile or stack of objects, to give more realistic behaviour. By default, this setting is disabled. |
| **Simulation Mode** | Choose when Unity executes the physics simulation from the following options:   * **Fixed Update** (default): Execute the physics simulation immediately after [`MonoBehaviour.FixedUpdate`](../ScriptReference/MonoBehaviour.FixedUpdate.md). * **Update**: Execute the physics simulation immediately after [`MonoBehaviour.Update`](../ScriptReference/MonoBehaviour.Update.md). * **Script**: Execute the physics simulation when a script calls [`Physics.Simulate`](../ScriptReference/Physics.Simulate.md). |
| **Auto Sync Transforms** | Enable this option to automatically sync transform changes with the physics system whenever a **Transform component** changes. By default, this setting is disabled. |
| **Reuse Collision Callbacks** | Enable to re-use collision callbacks. The physics system only creates a single instance of the Collision type, and reuses it for each individual callback. This reduces waste for the garbage collector to handle, and improves performance. This property is enabled by default. |
| **Invoke Collision Callbacks** | Send [MonoBehaviour](../ScriptReference/MonoBehaviour.md) collision messages for [`OnCollisionEnter`](../ScriptReference/MonoBehaviour.OnCollisionEnter.md), [`OnCollisionStay`](../ScriptReference/MonoBehaviour.OnCollisionStay.md), and [`OnCollisionExit`](../ScriptReference/MonoBehaviour.OnCollisionExit.md) to the corresponding **scripts** that implement the `OnCollision` methods. This property is enabled by default. |
| **Contact Pairs Mode** | Choose the type of contact pair generation to use:   * **Default Contact Pairs**: Receive collision and trigger events from all contact pairs except kinematic-kinematic and kinematic-static pairs. * **Enable Kinematic Kinematic Pairs**: Receive collision and trigger events from kinematic-kinematic contact pairs. * **Enable Kinematic Static Pairs**: Receive collision and trigger events from kinematic-static contact pairs. * **Enable All Contact Pairs**: Receive collision and trigger events from all contact pairs. |
| **Broadphase Type** | Choose which [broad-phase algorithm](https://docs.nvidia.com/gameworks/content/gameworkslibrary/physx/guide/3.3.4/Manual/RigidBodyCollision.html#broad-phase-algorithms) to use in the physics simulation. For more information, refer to NVIDIA’s documentation on [PhysX SDK](https://docs.nvidia.com/gameworks/content/gameworkslibrary/physx/apireference/files/structPxBroadPhaseType.html) and [Rigid Body Collision](https://docs.nvidia.com/gameworks/content/gameworkslibrary/physx/guide/Manual/RigidBodyCollision.html). Choose from the following options:   * **Sweep and Prune Broadphase**: Use the sweep-and-prune broad phase collision method which sorts objects along a single axis to rule out having to check pairs that are far apart. * **Multibox Pruning Broadphase**: Multibox pruning uses a grid, and each grid cell performs sweep-and-prune. This usually helps improve performance if, for example, there are many GameObjects in a flat world. * **Automatic Box Pruning**: This algorithm is similar to the **Multibox Pruning**, except that it can also automatically compute the world boundaries and number of subdivisions. It maintains the set of grid cells and uses the regular sweep-and-prune approach to work out potentially overlapping pairs of colliders. It usually helps with big scenes where a single Sweep and Prune would produce lots of extra false positives. |
| **World Bounds** | Define the 2D grid that encloses the world to prevent far away objects from affecting each other when using the **Multibox Pruning Broadphase** algorithm.  This option is only used when you set **Broadphase Type** to **Multibox Pruning Broadphase**. |
| **World Subdivisions** | The number of cells along the x and z axis in the 2D grid algorithm.  This option is only used when you set **Broadphase Type** to **Multibox Pruning Broadphase**. |
| **Friction Type** | Choose the friction algorithm used for simulation:  * **Patch Friction Type**: A basic strong friction algorithm which typically leads to the most stable results at low solver iteration counts. This method uses only up to four scalar solver constraints per pair of touching objects. * **One Directional Friction Type**: A simplification of the Coulomb friction model, in which the friction for a given point of contact is applied in the alternating tangent directions of the contact’s normal. This requires more solver iterations than patch friction but is not as accurate as the two-directional model. For [Articulation bodies](class-ArticulationBody.md) to work with this friction type, set the **Solver Type** to **Temporal Gauss Seidel**. * **Two Directional Friction Type**: Like the one-directional model, but applies friction in both tangent directions simultaneously. This requires more solver iterations but is more accurate. More expensive than patch friction for scenarios with many contact points because it is applied at every contact point. For [Articulation bodies](class-ArticulationBody.md) to work with this friction type, set the **Solver Type** to **Temporal Gauss Seidel**. |
| **Enable Enhanced Determinism** | Simulation in the **scene** is consistent regardless the actors present, provided that the game inserts the actors in a deterministic order. This mode sacrifices some performance to ensure this additional determinism. |
| **Enable Unified **Heightmaps**** | Enable this option to process **Terrain** collisions in the same way as **Mesh** collisions. |
| **Improved Patch Friction** | Optimize patch friction so that static and **dynamic friction** do not exceed analytical results. This property is deactivated by default. |
| **Solver Type** | Choose the PhysX solver type to use for the simulation:   * **Projected Gauss Seidel**: The default PhysX solver. * **Temporal Gauss Seidel**: This solver offers a better convergence and a better handling of high-mass ratios, minimizes energy introduced when correcting penetrations and improves the resistance of joints to overstretch. It usually helps when you experience some erratic behavior during simulation with the default solver. |
| **Default Max Angular Speed** | Set the project-wide default maximum angular speed of all dynamic Rigidbody GameObjects, in radian. The default value is `50`. |
| **Scratch Buffer Chunk Count** | The number of 16KB memory chunks to allocate to the physics system for temporary allocations. The default value is 4, which provides a 64KB buffer. |
| **Fast Motion Threshold** | The linear motion threshold that sweep-based CCD algorithms use to determine whether a fast-moving body moved this frame or not. Must be greater than zero. The default value is infinity, represented by the value `3.402823e+38`. |

### Cloth settings

The settings in the **Cloth** tab apply to Cloth physics only.

![The Cloth Physics Settings window.](../uploads/Main/Physics_cloth_settings.png)

The Cloth Physics Settings window.

| **Setting** | **Description** |
| --- | --- |
| **Cloth Gravity** | Set the gravity value on each axis for Cloth components. By default, X is set to `0`, Y is set to `-9.81`, and Z is set to `0`. |
| **Enable Cloth Inter-Collision** | Enable the ability for Cloth particles to collide with each other. Refer to [Cloth: Self collision and intercollision](class-Cloth.md#intercollision) for details. When you enable this setting, you can adjust the following additional settings:   * **Inter-Collision Distance**: Define the diameter of a sphere around each inter-colliding Cloth particle, in world units. Unity ensures that these spheres do not overlap during simulations. The distance should be smaller than the smallest distance between two particles in the configuration. If the distance is larger, cloth collision can violate distance constraints and result in jittering. This value is set to `0.1` by default. * **Inter-Collision Stiffness**: Define the strength of the separating force between inter-colliding Cloth particles when they are closer than the **Distance** value. The value is the factor by which PhysX multiplies the separation force; the cloth solver then calculates the stiffness based on the value you provide. This value is set to `0.2` by default. |