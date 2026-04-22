# Search your project and scenes

You can search your project and **scenes** with the Hierarchy, Project, and Asset Database search providers:

* [**Hierarchy**](search-scene.md): Search **GameObjects** in the loaded scenes.
* [**Project**](search-assets.md): Search the project’s assets. This is a faster search than the Asset Database search, and supports more filters. The Project provider includes more information about each asset than the Asset Database provider, such as the asset’s properties.
* [**Asset Database**](search-adb.md): Search the Asset Database. The Asset Database provider includes some asset states that the Project provider doesn’t, such as the asset’s **version control** state.

**Note:** By default, the Asset Database provider is inactive. To enable it, refer to [Activate and deactivate search providers](search-manage-providers.md).

## Results examples

To demonstrate the differences between the providers, consider the following example: In a scene with a single GameObject named `Dry5818`, based on a **prefab** of the same name:

| Provider | Query example | Results |
| --- | --- | --- |
| **Hierarchy** | `h: Dry5818` | GameObjects in the scene that have `Dry5818` in their name. |
| **Project** | `p: Dry5818` | Assets in the Project folder that have `Dry5818` in their name, including **LOD** files. |
| **Asset Database** | `adb: Dry5818` | Assets in the Asset Database with `Dry5818` in their name, including `fbx` and `prefab`, but not including the LOD files, which aren’t saved to the Asset Database. |

## Filter examples

The Hierarchy, Project, and Asset Database use similar filters, but the details differ. For example, all three can search by component, but the Asset Database can’t search by parameter values.

| Provider | Example description | Query |
| --- | --- | --- |
| **Asset Database** | Finds assets with a Sphere **Collider** component, such as presets for that component, but can’t filter on parameter values within that component. | `adb: t:SphereCollider` |
| **Project** | Finds assets with a Sphere Collider component, such as presets for that component, and can focus the search on Sphere Colliders with a radius greater than 0 in assets, such as presets for that component. | `p: t:SphereCollider`    `p: spherecollider.radius>0` |
| **Hierarchy** | Finds GameObjects in the loaded scenes that have a Sphere Collider component, and can focus the search on GameObjects in loaded scene with a Sphere Colliders component with a radius greater than 0. | `h: t:SphereCollider`    `h: #SphereCollider.m_Radius>0` |

For all filters, refer to:

* [Search GameObjects with the Hierarchy search provider](search-scene.md)
* [Search assets with the Project search provider](search-assets.md)
* [Search the Asset Database](search-adb.md).

## Deep dive: Hierarchy and Project search providers

The Hierarchy and Project search providers are often used together to quickly add project assets to GameObjects in the scene. Because of this, it’s important to understand how the providers’ filters complement each other:

* The **Hierarchy** window handles only GameObjects, therefore the Hierarchy search provider doesn’t have a Type filter. It has a Components filters and a Properties filter. The Properties filter is always in the context of a component, but you can search for components without limiting them by a specific property. If you select a component in the Hierarchy search provider, the Properties filter shows only properties possible for that component.
* The Project search has a Types filter and a Properties filter. It doesn’t explicitly have a Components filter, but properties are always in the context of a component.
* Assets can have properties that GameObjects don’t have, such as import properties.

### Drill-down examples

The types, components, and properties filters allow you to drill down based on what you know about the asset or GameObject you’re looking for.

For example:

* To find all material assets in the project, both used and unused, use the Project’s Type filter:

  + `p: t:Material`
  + `p: t:Material marble` (narrow the search by the name `marble`)
* To find all assets or GameObjects that include a specific component, use the Project or Hierarchy Component filter:

  + `p: t:BoxCollider`
  + `h: t:BoxCollider`
* To find an asset or GameObject that don’t have an assigned material for their PhysicsMaterial property, use the Project or Hierarchy Properties filter:

  + `p: boxcollider.material=none`
  + `h: boxcollider.material=none`
* To find all assets or GameObjects that use a specific material for a specific component, use the Project or Hierarchy Properties filter. For example, to filter for the BoxCollider components that use the PhysicsMaterial `NewPhysicsMaterial`:

  + `p: boxcollider.material=NewPhysicsMaterial`
  + `h: #BoxCollider.m_Material=NewPhysicsMaterial`
* To find a GameObject in the scene that uses a material, but not limited to the scope of a single component, use:

  + The Hierarchy’s Custom Scene filter for material:

    `h: material:NewPhysicsMaterial`
  + The Project’s References filter with a path and file name:

    `p: ref=path/to/my/material.mat`
* To find all GameObjects that are missing a reference to a script, use the Hierarchy’s Missing filters:

  `h: missing:Any`
* To find assets or GameObjects using specific references, use the Project or Hierarchy References filter:

  `p: ref="New Auto Tile"`

## Additional resources

* [Search for files](search-files.md)
* [Search presets](search-presets.md)
* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)