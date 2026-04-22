# Manage references

References are properties that use a **GameObject**, a component, or an asset as an input.

For example, on a 3D GameObject:

* The [Mesh Filter](class-MeshFilter.md) component, which gives the mesh its shape, refers to a [mesh](class-Mesh.md) asset.
* The **Mesh Renderer** component refers to a material.

Components can have default references, or you might need to add a reference manually. For example:

* If you create a new 3D Cube GameObject, its Mesh Filter and Mesh Renderer components have default references to the Cube mesh and your project’s default material. You can replace these, but you don’t have to.
* If you create an empty GameObject and add the Mesh Filter and Mesh Renderer components, they don’t refer to the default mesh or material. You must add those references manually.

## Assigning references

To assign a reference to a property, in the property field in the ****Inspector**** window, do one of the following:

* From the **Hierarchy** or **Project** windows, drag a compatible item into the field.
* To open a search window for the property field, select **Object Picker** (⊙). The search window filters for references that match the property type.

  There are two types of reference search windows: classic and advanced. The advanced search shows the built-in filter, so you can override it. The classic view doesn’t allow this override. For more information about the advanced search and how to toggle it on or off, refer to [Search Object Picker](search-advanced-object-picker.md).

## Component-dependent references

You can assign a GameObject to a property field that expects a component reference. When you assign this GameObject, Unity uses the first component of the required type from that GameObject. In the **Inspector** window, the first component is the one at the top of the list.

If the GameObject doesn’t have the right component type, you can’t assign the GameObject as the reference. For example, if the reference field requires a Sphere **Collider**, and you try to assign a GameObject that only has a **Box Collider**, the Unity Editor doesn’t accept that assignment.

## Additional references

* [Search Object Picker](search-advanced-object-picker.md)
* [The Hierarchy window](Hierarchy.md)