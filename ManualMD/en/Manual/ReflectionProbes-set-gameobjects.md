# Add GameObjects to reflections

## Add GameObjects to baked Reflection Probes

The reflections captured by baked probes can only include **scene** objects marked as ****Reflection Probe** Static** (using the **Static** menu at the top left of the inspector panel for all objects). You can further refine the objects that get included in the reflection **cubemap** using the ****Culling Mask**** and ****Clipping Planes**** properties, which work the same way as for a [Camera](class-Camera.md) (the probe is essentially like a camera that is rotated to view each of the six cubemap faces).

## Add GameObjects to realtime Reflection Probes

You don’t need to mark objects as **Reflection Probe Static** to capture their reflections (as you would with a baked probe). However, you can selectively exclude objects from the reflection cubemap using the **Culling Mask** and **Clipping Planes** properties, which work the same way as for a [Camera](class-Camera.md) (the probe is essentially like a camera that is rotated to view each of the six cubemap faces).

## Update Reflection Probes

To update probes, do one of the following:

* In a probe’s ****Inspector**** window, select **Bake**.
* In the Lighting window, select **Generate Lighting**. This updates all probes.

The bake process will take place asynchronously while you continue to work in the editor. If you move any static objects, change their materials or otherwise alter their visual appearance during the baking process, you must rebake to see the updated results.

**Note**: Currently, real-time probes will only update their reflections in the **Scene view** when **Reflection Probe Static** objects are moved or change their appearance. This means that moving dynamic objects won’t cause an update even though those objects appear in the reflection. You should choose the **Bake Reflection Probes** option from the **Generate Lighting** button dropdown in the [Lighting window](lighting-window.md) to update reflections when a dynamic object is changed.