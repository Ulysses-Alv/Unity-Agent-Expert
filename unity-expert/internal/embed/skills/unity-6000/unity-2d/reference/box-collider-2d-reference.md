# Box Collider 2D component reference

The Box **Collider** 2D component is a [Collider 2D](./collider-2d-landing.md) that interacts with the 2D physics system for **collision** detection. This collider 2D is a rectangle with a defined position, width, and height in the local coordinate space of a **sprite**. Adjust the component properties to change the shape and behavior of the collider 2D.

**Note**: The selection rectangle is axis-aligned, with its edges parallel to the x or y axes of local space.

| **Property** | **Function** |
| --- | --- |
| **Material** | Select the [Physics Material 2D](../physics-material-2d-reference.md) that determines properties of collisions, such as friction and bounce. |
| **Is Trigger** | Enable this if you want this Collider 2D to behave as a trigger. The physics system ignores this Collider when this is enabled. |
| **Used by Effector** | Enable this if you want the Collider 2D to be used by an attached Effector 2D. |
| **Composite Operations** | Select the [composite operation](../../../ScriptReference/Collider2D.CompositeOperation.md) used by an attached [Composite Collider 2D](composite-collider/composite-collider-2d-reference.md) component.  **Note:** When you select any operation besides **None**, the following properties—**Material**, **Is Trigger**, **Used By Effector**, and **Edge Radius**—become controlled by the attached Composite Collider 2D component and are no longer available in this collider’s properties. |
| **None** | Select this to have no composite operation take place. |
| **Merge** | Select this to have this composite operation compose geometry using a Boolean OR operation. |
| **Intersect** | Select this to have this composite operation compose geometry using a Boolean AND operation. |
| **Difference** | Select this to have this composite operation compose geometry using a Boolean NOT operation. |
| **Flip** | Select this to have this composite operation compose geometry using a Boolean XOR operation. |
| **Auto Tiling** | Enable this if the [Sprite Renderer](../../sprite/renderer/sprite-renderer-reference.md) component for the selected Sprite has the **Draw Mode** set to **Tiled**. This enables automatic updates to the shape of the [Collider 2D](./collider-2d-landing.md), allowing the shape to automatically readjust when the Sprite’s dimensions change. If you don’t enable **Auto Tiling**, the Collider 2D geometry doesn’t automatically repeat. |
| **Offset** | Set the local offset values of the Collider 2D geometry. |
| **Size** | Set the size of the box in local space units. |
| **Edge Radius** | Set a value that forms a radius around the edge of the Collider. This results in a larger Collider 2D with rounded convex corners. The default value is 0 (no radius). |
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

BoxCollider2D