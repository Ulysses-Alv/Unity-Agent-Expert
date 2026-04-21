# Set up a sorting group

To place a **GameObject** into a Sorting Group, add the Sorting Group component to it:

1. Select the GameObject and go to **Component > Rendering > Sorting Group**, or select the [Add Component](../../UsingComponents.md) button in the Inspector window of the GameObject.

When you add a Sorting Group component to a GameObject, Unity applies the same [Sorting Group](sorting-group-reference.md) to all child GameObjects of the GameObject the component is attached to.

Unity uses the Sorting Group’s settings to determine how to sort its Renderers among other Renderers and Sorting Groups in the **Scene**. See [2D sorting](../../2d-renderer-sorting.md) for more information.

To sort Renderers within a Sorting Group, Unity uses the individual sort settings of the Renderers in the Sorting Group. See [Sorting Renderers within a Sorting Group](sort-renderers-within-sorting-group.md) for more information.

SortingGroup