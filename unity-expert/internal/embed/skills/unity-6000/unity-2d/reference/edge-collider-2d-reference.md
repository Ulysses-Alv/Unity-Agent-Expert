# Edge Collider 2D component reference

The Edge **Collider** 2D component is a [Collider 2D](./collider-2d-landing.md) that interacts with the 2D physics system. The collider’s shape is an edge made of line segments that you can adjust to fit the shape of a **sprite** or any other shape. The collider’s start and end points don’t need to meet or enclose an area to function (unlike the [Polygon Collider 2D](./polygon-collider-2d-reference.md)), and can form a straight line or other single edge shapes.

| **Property** | **Function** |
| --- | --- |
| **Edit Collider** | Select Edit Collider  to make the collider outline editable. Refer to [Edit Collider mode reference](edit-collider-mode-reference.md) for the actions and shortcuts available when **Edit Collider** is enabled. |
| **Material** | Select the [Physics Material 2D](../physics-material-2d-reference.md) that determines properties of **collisions**, such as friction and bounce. |
| **Is Trigger** | Enable this if you want this Collider 2D to behave as a trigger. The physics system ignores this Collider when this is enabled. |
| **Used by Effector** | Enable this if you want the Collider 2D to be used by an attached Effector 2D. |
| **Offset** | Set the local offset values of the Collider 2D geometry. |
| **Edge Radius** | Set a value that forms a radius around the edge of the Collider. This results in a larger Collider 2D with rounded convex corners. The default value is 0 (no radius). |
| **Points** | Expand to view read-only information about the complexity of the generated Collider. |
| **Use Adjacent Start Point** | Enable this property to calculate the collision response using the **Adjacent Start Point** to form the collision normal when a collision occurs at the Edge Collider’s start point. |
| **Adjacent Start Point X/Y** | Set the x and y-coordinates of the **Adjacent Start Point**. |
| **Use Adjacent End Point** | Enable this property to calculate the collision response using the **Adjacent End Point** to form the collision normal when a collision occurs at the Edge Collider’s end point. |
| **Adjacent End Point X/Y** | Set the x and y-coordinates of the Adjacent End Point. |
| **Layer Overrides** | Expand for the Layer override settings. |
| **Layer Override Priority** | Assign the decision priority that this Collider2D uses when resolving conflicting decisions on whether a contact between itself and another Collision2D should happen or not. Refer to its [API](../../../ScriptReference/Collider2D-layerOverridePriority.md) page for more information. |
| **Include Layers** | Select the additional Layers that this Collider 2D should include when deciding if a contact with another Collider2D should happen or not. Refer to its [API](../../../ScriptReference/Collider2D-includeLayers.md) documentation for more information. |
| **Exclude Layers** | Select the additional Layers that this Collider 2D should exclude when deciding if a contact with another Collider2D should happen or not. Refer to its [API](../../../ScriptReference/Collider2D-excludeLayers.md) documentation for more information. |
| **Force Send Layers** | Select the Layers that this Collider 2D is allowed to send forces to during a Collision contact with another Collider2D. Refer to its [API](../../../ScriptReference/Collider2D-forceSendLayers.md) documentation for more information. |
| **Force Receive Layers** | Select the Layers that this Collider 2D can receive forces from during a Collision contact with another Collider2D. Refer to its [API](../../../ScriptReference/Collider2D-forceReceiveLayers.md) documentation for more information. |
| **Contract Capture Layers** | Select the Layers of other Collider 2D, involved in contacts with this Collider2D, that will be captured. Refer to its [API](../../../ScriptReference/Collider2D-contactCaptureLayers.md) documentation for more information. |
| **Callback Layers** | Select the Layers that this Collider 2D, during a contact with another Collider2D, will report collision or trigger callbacks for. Refer to its [API](../../../ScriptReference/Collider2D-callbackLayers.md) documentation for more information. |

## Additional resources

* [Collider 2D API documentation](../../../ScriptReference/Collider2D.md)
* [Rigidbody 2D](../rigidbody/rigidbody-2d-landing.md)

EdgeCollider2D