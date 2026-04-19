# Sort 3D GameObjects with 2D sprites

Sort 3D **GameObjects** with 2D **sprites** in the same **scene**.

When you enable the **Sort as 2D** property, if 3D meshes have the same depth, Unity renders them like overlapping 2D sprites instead of intersecting meshes.

Follow these steps:

1. Add a [Sorting Group](../sprite/sorting-group/sorting-group-landing.md) component to the 3D GameObject.
2. Enable **Sort as 2D** in the **Sorting Group** property settings.
3. Assign [Sorting Layers and Order in Layer](../2d-renderer-sorting.md#sortlayer) values to determine the 3D renderer’s position in the render queue.

**Important:** Enabling the **Sort as 2D** setting clears the object’s depth information which means you cannot use it with depth-based **post-processing** effects. If you want to use these kind of effects, use Unity’s default behavior that renders sprites and meshes based on their [world-space depth](../2d-renderer-sorting.md#cameradistance).

## Additional resources

* [2D Renderer sorting](../2d-renderer-sorting.md)
* [Sorting groups](../sprite/sorting-group/sorting-group-landing.md)