# Sort renderers within a sorting group

Unity sorts all Renderers within the same Sorting Group by their individual **Sorting Layer** and **Order in Layer** [Renderer properties](../renderer/sprite-renderer-reference.md). Unity does not consider each Renderer’s individual **Distance to **Camera**** property during this sorting process. Instead, it sets a Distance to Camera value for the whole Sorting Group (including all its child Renderers), based on the position of the root **GameObject** that contains the Sorting Group component.

The Sorting Group’s internal sort order remains constant when Unity sorts the Sorting Group among other Renderers and Sorting Groups in the **Scene**.

The following diagram shows the sorting process.

![Internal Sorting Group sorting process.](../../../uploads/Main/SG_diagram1.png)

Internal Sorting Group sorting process.

Unity treats all Renderers that belong to the same Sorting Group as a single layer, and sorts non-grouped Renderers based on their **Sorting Layer** and **Order in Layer** property settings.

### Particle System

The Editor treats a [Particle System](../../class-ParticleSystem.md) that is a child of a Sorting Group as another Renderer within that Sorting Group, and sorts it internally among other Renderers based on its **Sorting Layer** and **Order in Layer** property settings.

When Unity sorts the **Particle** System with the other Renderers within the Sorting Group, it ignores the Particle System’s **Sorting Fudge** value.

### Nested Sorting Groups

A nested Sorting Group is a Sorting Group that has a parent Sorting Group. Unity sorts Renderers within a nested Sorting Group first, before their parent (refer to [Sorting Renderers within a Sorting Group](sort-renderers-within-sorting-group.md)).

After Unity determines the internal sort order of the nested Sorting Group, it sorts the nested Sorting Group with other Renderers or Sorting Groups within the parent Sorting Group. A nested Sorting Group can have a nested Sorting Group as a child. Unity sorts the innermost child Groups first, before their respective parents.

The following diagram gives you an example of the nested Sorting Group sorting process.

![Nested Sorting Group sorting process.](../../../uploads/Main/SG_diagram2.png)

Nested Sorting Group sorting process.

### Sort At Root

In certain situations, you may have a nested Sorting Group that is ordered based on the Scene [Hierarchy](../../Hierarchy.md). However, you may want this nested Sorting Group to be rendered separately from its parent Sorting Group without changing its position in the Scene Hierarchy.

You can enable this option to allow this Sorting Group to ignore its parent Sorting Groups, which allows this Sorting Group to be sorted against other Renderers and Sorting Groups globally without requiring the GameObject to be reparented to another Transform. All child Renderers and Sorting Groups (which have not ignored their parents) will be sorted under this Sorting Group.

![Nested Sorting Group sorting process with the Sort At Root property enabled.](../../../uploads/Main/SG_diagram3.png)

Nested Sorting Group sorting process with the Sort At Root property enabled.

SortingGroup