# Custom Collider 2D component reference

The Custom **Collider** 2D is a [Collider 2D](../collider-2d-landing.md) that interacts with the 2D physics system. Unlike other colliders, you don’t configure this collider in the Unity Editor, instead you configure the collider by assigning [PhysicsShape2D](../../../../ScriptReference/PhysicsShape2D.md) geometry to it via the [PhysicsShapeGroup2D](../../../../ScriptReference/PhysicsShapeGroup2D.md) API.

You can define the collider’s shape by adding, removing, and modifying the `PhysicsShape2D` shapes. Refer to the [PhysicsShape2D](../../../../ScriptReference/PhysicsShapeGroup2D.md) API documentation for more information. This also means that a Custom Collider 2D can contain an unlimited number of low-level `PhysicsShape2D` and form any shape, or emulate other types of Collider 2D.

| **Property** | **Function** |
| --- | --- |
| **Material** | Select the [Physics Material 2D](../../physics-material-2d-reference.md) that determines properties of **collisions**, such as friction and bounce. |
| **Is Trigger** | Enable this if you want this Collider 2D to behave as a trigger. The physics system ignores this Collider when this is enabled. |
| **Used by Effector** | Enable this if you want the Collider 2D to be used by an attached Effector 2D. |
| **Offset** | Set the local offset values of the Collider 2D geometry. |
| **Custom Shape Count** (Read only) | Indicates how many `PhysicsShape2D` the Collider is using. |
| **Custom Vertex Count** (Read Only) | Indicates how many vertices all `PhysicsShape2D` in the Collider are using. |
| **Layer Overrides** | Expand for the Layer override settings. |
| **Layer Override Priority** | Assign the decision priority that this Collider2D uses when resolving conflicting decisions on whether a contact between itself and another Collision2D should happen or not. Refer to its [API](../../../../ScriptReference/Collider2D-layerOverridePriority.md) page for more information. |
| **Include Layers** | Select the additional Layers that this Collider 2D should include when deciding if a contact with another Collider2D should happen or not. Refer to its [API](../../../../ScriptReference/Collider2D-includeLayers.md) documentation for more information. |
| **Exclude Layers** | Select the additional Layers that this Collider 2D should exclude when deciding if a contact with another Collider2D should happen or not. Refer to its [API](../../../../ScriptReference/Collider2D-excludeLayers.md) documentation for more information. |
| **Force Send Layers** | Select the Layers that this Collider 2D is allowed to send forces to during a Collision contact with another Collider2D. Refer to its [API](../../../../ScriptReference/Collider2D-forceSendLayers.md) documentation for more information. |
| **Force Receive Layers** | Select the Layers that this Collider 2D can receive forces from during a Collision contact with another Collider2D. Refer to its [API](../../../../ScriptReference/Collider2D-forceReceiveLayers.md) documentation for more information. |
| **Contract Capture Layers** | Select the Layers of other Collider 2D, involved in contacts with this Collider2D, that will be captured. Refer to its [API](../../../../ScriptReference/Collider2D-contactCaptureLayers.md) documentation for more information. |
| **Callback Layers** | Select the Layers that this Collider 2D, during a contact with another Collider2D, will report collision or trigger callbacks for. Refer to its [API](../../../../ScriptReference/Collider2D-callbackLayers.md) documentation for more information. |

## Additional resources

* [Collider 2D API documentation](../../../../ScriptReference/Collider2D.md)
* [Rigidbody 2D](../../rigidbody/rigidbody-2d-landing.md)

CustomCollider2D