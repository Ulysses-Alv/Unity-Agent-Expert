# Material variants

A material variant is a [material](Materials.md) that inherits the properties of a parent material. Use material variants to create a hierarchy of materials, to help you manage and maintain groups of similar materials more efficiently than creating duplicate copies.

For example, use material variants to manage outfits with a variety of color schemes, damaged and undamaged versions of scenery, or shiny and weathered instances of scenery.

You can do the following more efficiently with material variants compared to duplicate copies:

* Change a parent material and automatically propagate those changes down to all the child materials.
* Copy property values from a child material to its parent.
* Prevent child materials modifying specific properties.
* Change the parent of a child material.

## Limitations

You can’t use material variants to do the following:

* Improve performance or scalability.
* Change materials at runtime.

## Similarities to prefab variants

Material variants are similar to [prefab variants for GameObjects](https://docs.unity3d.com/Manual/PrefabVariants.html), but with the following key differences:

* You can change the parent of a material variant.
* You can lock properties so child materials can’t override them.

## Additional resources

* [Materials introduction](materials-introduction.md)
* [Prefab variants](PrefabVariants.md)