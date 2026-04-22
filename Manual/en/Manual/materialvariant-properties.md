# Control material variant properties

By default, when you create a material variant, it inherits all the properties of its parent material. When you update a property of the parent material, Unity propagates the change down to all its child materials.

It takes a short while for Unity to update material variants. In the **Project** window, material variants display as a striped icon until Unity finishes updating them.

## Override a parent property

To override a parent property so the child material uses a different value, follow these steps:

1. Select the child material.
2. In the ****Inspector**** window, modify the property you want to override.

Unity displays a line to the left of the property to indicate that it’s an override.

To move the override further up the hierarchy, right-click on the property name and select the **Apply as Override in…** item that has the name of the parent or ancestor.

To apply the override value to the material at the top of the hierarchy, right-click on the property name and select **Apply to Material**. This removes any overrides for this property in the hierarchy.

## Remove an override

To remove an override, right-click the name of the property and select **Revert**. The child material reverts to inheriting the value from its parent.

You can also select **Revert all Overrides** to remove all the overrides on the material.

## Prevent overrides

To prevent child materials from overriding a parent property, lock the property. Follow these steps:

1. Select the parent material.
2. In the **Inspector** window, right-click on the name of the property you want to lock.
3. Select **Lock in children**.

Unity applies the value from the parent property to all its child materials, removes any overrides, and displays a padlock in the **Inspector** window of the child materials.

To unlock the property, follow the same steps again. Unity restores any overrides it removed when you locked the property.

## Additional resources

* [Material variants](materialvariant-concept.md)