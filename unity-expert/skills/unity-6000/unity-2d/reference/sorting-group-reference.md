# Sorting group reference

Unity uses a **Sorting Group**’s [Sorting Layer](https://docs.unity3d.com/Manual/class-TagManager.html#SortingLayers) and **Order in Layer** values to determine its priority in the rendering queue among other Sorting Groups and **GameObjects** in the **Scene**.

| Property | Function |
| --- | --- |
| **Sorting Layer** | Select or add a [Sorting Layer](https://docs.unity3d.com/Manual/class-TagManager.html#SortingLayers) from this drop-down menu to determine the Sorting Group’s position in the render queue. Unity determines the Sorting Layer order by its place in the Sorting Layer settings; it renders Sorting Layers in the order they appear in the list. See [Tags and Layers](https://docs.unity3d.com/Manual/class-TagManager.html#SortingLayers) for information on setting up Sorting Layers. |
| **Order in Layer** | Set the render order of this Sorting Group within its **Sorting Layer**. Unity queues Renderers with lower values first in the render queue, so they appear before Renderers with higher values. |
| **Sort At Root** | Enable this option to ignore all parent Sorting Groups if this Sorting Group is nested. This allows this Sorting Group to be sorted against other Renderers and Sorting Groups at the root level. |

See [2D Sorting](../../2d-renderer-sorting.md) for more information on using Sorting Layers to sort **sprites**, and Unity’s Renderer sorting criteria.