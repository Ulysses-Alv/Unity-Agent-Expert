# Add shadows from 2D lights in URP

To add shadows from 2D light in the Universal **Render Pipeline** (URP), add the **Shadow Caster 2D** component to a **GameObject**. Follow these steps:

**Note**: To see cast shadows, you must have a 2D light in your **scene** with its **Shadows** property enabled. For more information, refer to [2D lights in URP](../urp/2d-index.md).

1. Select the GameObject you want to cast shadows.
2. In the ****Inspector**** window, select **Add Component**.
3. Type **Shadow Caster 2D**, then select the **Shadow Caster 2D** item.

## Change the shape of a shadow

By default, the shadow is the shape defined in the **Custom Outline** window of the [Sprite Editor](../sprite/sprite-editor/sprite-editor-landing). Unity displays the shape as a white outline in the **Scene view**.

To set a specific shape for the shadow, do the following:

1. Select the GameObject with the **Shadow Caster 2D** component.
2. In the **Inspector** window, set the **Casting Source** property to **Shape Editor**.
3. Select **Edit Shape**.

Adjust the shape in the Scene view by doing the following:

* To move a point, click and drag the point. You can also select the point, then manually enter the new position in the **Position** property of the **Shadow Caster 2D** component.
* To add a point, click a line at the position you want to add the point.

## Merge shadows from multiple GameObjects

To merge shadows from multiple GameObjects into a single shadow, add a **Composite Shadow Caster 2D** to a parent GameObject. Follow these steps:

1. Create an empty GameObject.
2. In the **Inspector** window, select **Add Component**.
3. Type **Composite Shadow Caster 2D**, then select the **Composite Shadow Caster 2D** item.
4. Add GameObjects as children of the GameObject in the **Hierarchy** window. If the child GameObjects have Shadow Caster 2D components, their shadows are merged.

![Left: Two rocks cast shadows without a Composite Shadow Caster 2D component. The upper object casts its shadow onto the lower object. Right: The same rocks with a parent Composite Shadow Caster 2D component. The merged shadow is cast behind both objects.](../../uploads/urp/2D/composite_shadow-without-with.png)

Left: Two rocks cast shadows without a Composite Shadow Caster 2D component. The upper object casts its shadow onto the lower object. Right: The same rocks with a parent Composite Shadow Caster 2D component. The merged shadow is cast behind both objects.

## Additional resources

* [Shadow Caster 2D component reference for URP](ShadowCaster2D.md)