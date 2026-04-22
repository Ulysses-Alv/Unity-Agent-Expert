# Polygon Collider 2D component reference

The Polygon **Collider** 2D component is a [Collider 2D](./collider-2d-landing.md) that interacts with the 2D physics system. This collider 2D’s shape is a freeform edge of line segments that you can adjust to fit the shape of a **sprite** or any other shape. **Note:** The edge must completely enclose an area for the collider to work.

| **Property** | **Function** |
| --- | --- |
| **Edit Collider** | Select this to edit the collider’s geometry by editing and moving its vertices. |
| **Material** | Select the [Physics Material 2D](../physics-material-2d-reference.md) that determines properties of **collisions**, such as friction and bounce. |
| **Is Trigger** | Enable this if you want this Collider 2D to behave as a trigger. The physics system ignores this collider when this is enabled. |
| **Used by Effector** | Enable this if you want an attached Effector 2D to use the Collider 2D. |
| **Auto Tiling** | Enable this if the [Sprite Renderer](../../sprite/renderer/sprite-renderer-reference.md) component for the selected sprite has its **Draw Mode** set to **Tiled**. This enables automatic updates to the shape of the [Collider 2D](./collider-2d-landing.md), allowing the shape to automatically readjust when the Sprite’s dimensions change. If you don’t enable **Auto Tiling**, the Collider 2D geometry doesn’t automatically repeat. |
| **Composite Operations** | Select the [composite operation](../../../ScriptReference/Collider2D.CompositeOperation.md) used by an attached [Composite Collider 2D](composite-collider/composite-collider-2d-reference.md) component.  **Note:** When you select any operation besides **None**, the following properties—**Material**, **Is Trigger**, **Used By Effector**, and **Edge Radius**—become controlled by the attached Composite Collider 2D component and are no longer available in this collider’s properties. |
| **None** | Select this to have no composite operation take place. |
| **Merge** | Select this to have this composite operation compose geometry using a Boolean OR operation. |
| **Intersect** | Select this to have this composite operation compose geometry using a Boolean AND operation. |
| **Difference** | Select this to have this composite operation compose geometry using a Boolean NOT operation. |
| **Flip** | Select this to have this composite operation compose geometry using a Boolean XOR operation. |
| **Offset** | Set the local offset values of the Collider 2D geometry. |
| **Use Delaunay **Mesh**** | Enable this property to include an additional Delaunay triangulation step to produce the collider mesh. |
| **Points** | Expand to view information about the complexity of the generated collider. |
| **Layer Overrides** | Expand for the Layer override settings. |
| **Layer Override Priority** | Assign the decision priority that this Collider 2D uses when resolving conflicting decisions on whether a contact between itself and another Collider 2D happens or not. Refer to its [API](../../../ScriptReference/Collider2D-layerOverridePriority.md) page for more information. |
| **Include Layers** | Select the additional Layers that this Collider 2D includes when deciding if a contact with another Collider 2D happens or not. Refer to its [API](../../../ScriptReference/Collider2D-includeLayers.md) documentation for more information. |
| **Exclude Layers** | Select the additional Layers that this Collider 2D excludes when deciding if a contact with another Collider 2D happens or not. Refer to its [API](../../../ScriptReference/Collider2D-excludeLayers.md) documentation for more information. |
| **Force Send Layers** | Select the Layers that this Collider 2D can send forces to during a Collision contact with another Collider 2D. Refer to its [API](../../../ScriptReference/Collider2D-forceSendLayers.md) documentation for more information. |
| **Force Receive Layers** | Select the Layers that this Collider 2D can receive forces from during a Collision contact with another Collider 2D. Refer to its [API](../../../ScriptReference/Collider2D-forceReceiveLayers.md) documentation for more information. |
| **Contract Capture Layers** | Select the Layers of other Collider 2D, involved in contacts with this Collider 2D, that will be captured. Refer to its [API](../../../ScriptReference/Collider2D-contactCaptureLayers.md) documentation for more information. |
| **Callback Layers** | Select the Layers that this Collider 2D, during a contact with another Collider 2D, will report collision or trigger callbacks for. Refer to its [API](../../../ScriptReference/Collider2D-callbackLayers.md) documentation for more information. |
| **Info** | Expand for read-only physics system related information about the collider. |

## Additional resources

* [Collider 2D API documentation](../../../ScriptReference/Collider2D.md)
* [Rigidbody 2D](../rigidbody/rigidbody-2d-landing.md)

PolygonCollider2D