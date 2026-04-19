# Blend Reflection Probes in the Built-In Render Pipeline

To enable Reflection Probe blending navigate to **Graphic Settings** > **Tier settings**. Tier settings are only available in Unity’s [Built-in Render Pipeline](built-in-render-pipeline.md). When blending is enabled, Unity gradually fades out one probe’s cubemap while fading in the other’s as the reflective object passes from one zone to the other. This gradual transition avoids the situation where a distinctive object suddenly “pops” into the reflection as an object crosses the zone boundary.

Blending is controlled using the **Reflection Probes** property of the [Mesh Renderer](class-MeshRenderer.md) component. Four blending options are available:

* **Off** - Reflection probe blending is disabled. Only the **skybox** will be used for reflection
* **Blend Probes** - Blends only adjacent probes and ignores the skybox. You should use this for objects that are “indoors” or in covered parts of the **scene** (eg, caves and tunnels) since the sky is not visible from these place and so should never appear in the reflections.
* **Blend Probes and Skybox** - Works like *Blend Probes* but also allows the skybox to be used in the blending. You should use this option for objects in the open air, where the sky would always be visible.
* **Simple** - Disables blending between probes when there are two overlapping reflection probe volumes.

When probes have equal **Importance** values, the blending weight for a given probe zone is calculated by dividing its intersection (volume) with the object’s bounding box by the sum of all probes’ intersections with the box. For example, if the box intersects probe A’s zone by 1.0 cubic units and intersects probe B’s zone by 2.0 cubic units then the blending values will be:

* Probe A: 1.0 / (1.0 + 2.0) = 0.33
* Probe B: 2.0 / (1.0 + 2.0) = 0.67

In other words, the blend will incorporate 33% of probe A’s reflection and 67% of probe B’s reflection.

The calculation must be handled slightly differently in the case where one probe is entirely contained within the other, since the inner zone overlaps entirely with the outer. If the object’s bounding box is entirely within the inner zone then that zone’s blending weight is 1.0 (ie, the outer zone is not used at all). When the object is partially outside the inner zone, the intersection volume of its bounding box with the inner zone is divided by the total volume of the box. For example, if the intersection volume is 1.0 cubic units and the bounding box’s volume is 4.0 cubic units, then the blending weight of the inner probe will be 1.0 / 4.0 = 0.25. This value is then subtracted from 1.0 to get the weight for the outer probe which in this case will be 0.75.

When one probe involved in the blend has a higher **Importance** value than another, the more important probe overrides the other in the usual way.

## Additional resources

* [Reflections](reflections-landing.md)