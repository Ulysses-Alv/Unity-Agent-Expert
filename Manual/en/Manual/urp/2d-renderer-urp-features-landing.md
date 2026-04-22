# 3D GameObjects in 2D URP scenes

Render and sort 3D **GameObjects** as if they were 2D **sprites** in your 2D URP project.

When Unity renders 3D GameObjects in a 2D **scene**, it usually renders them clipping into each other on the same z-position. By using the following properties and features, Unity renders the 3D GameObjects as if they were 2D sprites, so that they would overlap each other rather than intersect.

You can use features that usually only work with 2D sprites with these 3D GameObjects, such as allowing these 3D GameObjects to receive lighting from [2D Lights](2d-index.md) by using a compatible **shader**, or interact with [Sprite Masks](../sprite/mask/mask-landing.md). These 3D GameObjects can also be sorted with sprites in the [2D renderer sorting queue](../2d-renderer-sorting.md) when **Sort 3D As 2D** is enabled in a [Sorting Group](../sprite/sorting-group/sorting-group-landing.md).

| Topic | Description |
| --- | --- |
| [Render 3D GameObjects in 2D scenes](2d-renderer-urp-shader-compatibility.md) | Create custom shaders and materials to render 3D GameObjects in 2D scenes. |
| [Set 3D GameObjects to interact with Sprite masks](2d-renderer-urp-sprite-mask-interaction.md) | Use Sprite masks with 3D GameObjects. |
| [Sort 3D GameObjects with 2D sprites](2d-renderer-urp-sorting-workflows.md) | Sort 3D GameObjects with 2D GameObjects in the same scene and control the order they’re rendered. |