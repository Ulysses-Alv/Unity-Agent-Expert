# Tilemap Collider 2D component reference

The **Tilemap **Collider** 2D** component generates collider shapes for [tiles](../tiles-for-tilemaps/tile-asset-reference.md) on a [Tilemap](./tilemap-reference.md). You control this collider generation process with the properties found in the component.

## Properties

| Property | Description |
| --- | --- |
| **Max Tile Change Count** | Set the maximum number of tile changes (such as adding/removing tiles to the tilemap) to accumulate before doing a full collider rebuild instead of an incremental rebuild. **Note:** A high number of accumulated changes can cause the incremental rebuild of the **Tilemap Collider 2D** to be slower than a full rebuild. Decrease this value to resolve this issue. |
| **Extrusion Factor** | Set the amount (in Unity world space units) to extrude the collider shape of each tile. This minimizes the gaps between the collider shapes of neighboring tiles and brings them to within the minimum **Vertex Distance** set in the [Composite Collider 2D](../../2d-physics/collider/composite-collider/composite-collider-2d-reference.md) component., which can then compose the tile colliders together.  **Note:** This property isn’t available by default. It becomes available when you select a [Composite Operation](#composite) and attach a Composite Collider 2D to the same GameObject. |
| **Use Delaunay **Mesh**** | Enables Unity using a Delaunay triangulation algorithm to generate the collider shapes. Enabling this property can improve collider shapes for tiles with complex shapes, but can reduce performance. |
| **Material** | Select the [Physics Material 2D](../../2d-physics/physics-material-2d-reference.md) that determines properties of **collisions**, such as friction and bounce. |
| **Is Trigger** | Enable this if you want this Collider 2D to behave as a trigger. The physics system ignores this collider when this is enabled. |
| **Used by Effector** | Enable this if you want the Collider 2D to be used by an attached [Effector 2D](../../2d-physics/effectors/effectors-2d-landing.md). |
| **Composite Operations** | Select the [composite operation](../../../ScriptReference/Collider2D.CompositeOperation.md) used by an attached [Composite Collider 2D](../../2d-physics/collider/composite-collider/composite-collider-2d-reference.md) component.  **Note:** When you select any operation besides **None**, the following properties—**Material**, **Is Trigger**, **Used By Effector**, and **Edge Radius**—become controlled by the attached Composite Collider 2D component and are no longer available in this collider’s properties. |
| **None** | Select this to have no composite operation take place. |
| **Merge** | Select this to have this composite operation compose geometry using a Boolean OR operation. |
| **Intersect** | Select this to have this composite operation compose geometry using a Boolean AND operation. |
| **Difference** | Select this to have this composite operation compose geometry using a Boolean NOT operation. |
| **Flip** | Select this to have this composite operation compose geometry using a Boolean XOR operation. |
| **Offset** | Set the local offset values of the Collider 2D geometry. |
| **Layer Overrides** | Expand for the Layer override settings. |
| **Layer Override Priority** | Assign the decision priority that this Collider2D uses when resolving conflicting decisions on whether a contact between itself and another Collision2D should happen or not. Refer to its [API](../../../ScriptReference/Collider2D-layerOverridePriority.md) page for more information. |
| **Include Layers** | Select the additional Layers that this Collider 2D should include when deciding if a contact with another Collider2D should happen or not. Refer to its [API](../../../ScriptReference/Collider2D-includeLayers.md) documentation for more information. |
| **Exclude Layers** | Select the additional Layers that this Collider 2D should exclude when deciding if a contact with another Collider2D should happen or not. Refer to its [API](../../../ScriptReference/Collider2D-excludeLayers.md) documentation for more information. |
| **Force Send Layers** | Select the Layers that this Collider 2D is allowed to send forces to during a Collision contact with another Collider2D. Refer to its [API](../../../ScriptReference/Collider2D-forceSendLayers.md) documentation for more information. |
| **Force Receive Layers** | Select the Layers that this Collider 2D can receive forces from during a Collision contact with another Collider2D. Refer to its [API](../../../ScriptReference/Collider2D-forceReceiveLayers.md) documentation for more information. |
| **Contract Capture Layers** | Select the Layers of other Collider 2D, involved in contacts with this Collider2D, that will be captured. Refer to its [API](../../../ScriptReference/Collider2D-contactCaptureLayers.md) documentation for more information. |
| **Callback Layers** | Select the Layers that this Collider 2D, during a contact with another Collider2D, will report collision or trigger callbacks for. Refer to its [API](../../../ScriptReference/Collider2D-callbackLayers.md) documentation for more information. |

## Additional resources

* [Physics 2D Reference](../../2d-physics/2d-physics.md)
* [Collider 2D](../../2d-physics/collider/collider-2d-landing.md)
* [Composite Collider 2D](../../2d-physics/collider/composite-collider/composite-collider-2d-reference.md)

---

* New properties added in Unity 2020.1
* Tilemaps added in [2017.2](https://docs.unity3d.com/2017.2/Documentation/Manual/30_search.html?q=newin20172) NewIn20172
  TilemapCollider2D