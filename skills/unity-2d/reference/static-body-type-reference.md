# Static Body Type reference

Due to limited behavior, **Rigidbody** 2D with a Static **body type** only have a very limited set of properties are available for this **Body Type**.

## Static Rigidbody 2D properties

| **Property** | **Function** |
| --- | --- |
| **Body Type** | Select to set the movement behavior and **Collider** 2D interaction of this Rigidbody 2D’s component settings. |
| **Dynamic** | Select to set this Rigidbody 2D to the [Dynamic](../dynamic/dynamic-body-type-reference.md) Body Type, which is designed to move under simulation and has all Rigidbody 2D properties available. The is the default Body Type for a Rigidbody 2D. |
| **Kinematic** | Select to set this Rigidbody 2D to the [Kinematic](../kinematic/kinematic-body-type-reference.md) Body Type, which is designed to move under simulation but only under very explicit user control. Refer to [Body Type: Kinematic](../kinematic/kinematic-body-type-reference.md) for more information. |
| **Static** | Select to set this Rigidbody 2D to the Static Body Type, which is designed to not move under simulation at all and behaves like an immovable object with infinite mass. |
| **Material** | Set a common [physics material](../../../physics-material-2d-reference.md) for all Collider 2Ds attached to this Rigidbody 2D. **Note:** A Collider 2D uses its own Material property if it has one set. If there is no Material specified here or in the Collider 2D, the default option is **None (Physics Material 2D)**. This uses a default Material which you can set in the [Physics 2D](../../../../class-Physics2DSettings.md) window.    **Note**: Use this to ensure that all Collider 2Ds attached to the same **Static** Body Type Rigidbody 2D can all use the same Material. |
| **Simulated** | Enable **Simulated** to have the Rigidbody 2D and any attached Collider 2Ds and **Joint** 2Ds to interact with the physics simulation during runtime. If this is disabled, these components do not interact with the simulation. Refer to [Rigidbody 2D properties: Simulated](../../rigidbody-2d-simulated-property.md) for more information. This property is enabled by default. |
| **Layer Overrides** | Expand for the Layer override settings. |
| **Include Layers** | Select the additional Layers that all Collider 2Ds attached to this Rigidbody 2D should include, when deciding if a **collision** with another Collider2D should occur or not. Refer to [Rigidbody2D-includeLayers](../../../../../ScriptReference/Rigidbody2D-includeLayers.md) for more information. |
| **Exclude Layers** | Select the additional Layers that all Collider 2Ds attached to this Rigidbody 2D should exclude, when deciding if a collision with another Collider 2D should occur or not. Refer to [Rigidbody2D-excludeLayers](../../../../../ScriptReference/Rigidbody2D-excludeLayers.md) for more information. |