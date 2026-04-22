# Search GameObjects with the Hierarchy search provider

The Hierarchy search provider searches all [loaded scenes](setupmultiplescenes.md), including inactive ones.

The search uses progressive caching, not an index, so you can use it even if your project has no indices.

## Query syntax

Provider token: `h:` (for `Hierarchy`)

## Query examples

| Query example | Description |
| --- | --- |
| `h: camera` | Find all **GameObjects** in the **scene** that have the word `camera` in their name. |
| `h: t=MeshFilter vertices>4096` | Find all meshes in the scene that are bigger than 4096 vertices. |
| `h: t=MeshRenderer ref="Assets/Art/Environment/Materials/Architect/M_Metal01.mat"` | Find all renderers in the scene using M\_Metal01 material. |
| `h: t:Camera #fieldofview>70` | Find all cameras in the scene whose field of view property is greater than 70. |
| `h: t:Light #m_layer=5 #m_Intensity>0.5` | Find lights in the scene on layer `5` with an intensity property greater than `0.5`. |
| `h: t:MeshRenderer #StaticShadowCaster=True` | Find all **mesh** renderers in the scene that are static shadow casters. |
| `h: tag=Player layer=9` | Find objects in the scene by layer and tag. |
| `h: t:SpriteRenderer #DrawMode=Sliced` | Find all **sprites** in the scene with the `Sliced` draw mode. |
| `h: t:SpriteRenderer -#DrawMode=Sliced` | Find all sprites in the scene drawn as a normal sprite (not sliced). |
| `h: t:SpriteRenderer ref="Assets/Prefabs/Crystal.prefab"` | Find all GameObjects in the scene referencing a specific prefab. |
| `h: -t:LODGroup t:MeshRenderer vertices>2000` | Find meshes in the scene that have no Level of Detail (LOD) and a triangle count higher than 2000. |
| `h: ref=”Assets/Prefabs/Models/Buildings/Fence.prefab”` | Find all assets used in a selected scene prefab. |

## Provider filters

The Hierarchy search provider supports the following filters.

### Components and properties

The syntax for components and properties in the Hierarchy search provider:

* Use the `t` prefix to search for a component. The prefix can require an exact match if you use `=`, or a partial match if you use `:`.

  `h: t:Came`
  `h: t=Camera`
* Use the `component.property` or `p(component.property)` notation to find properties with a specified value.

  `h: Camera.Orthographic=true`
  `h: p(Camera.Orthographic)=true`
  `h: Orthographic=true`
  `h: p(Orthographic)=true`
* If a property is unique to a component, you don’t have to specify the component in the query. For example, the following queries are equivalent if only the **Camera** component has a field of view property:

  `h: fieldofview=60`
  `h: Camera.fieldofview=60`
* Words in filters can’t have spaces, so if a property or component name has spaces, remove them. For example, `Field of view` becomes `fieldofview`.

Unity indexes properties according to their internal name, which might be different than the display name in the ****Inspector**** window. To search by property value, do one of the following:

* Use the visual query builder, which uses the property’s display name.
* Check the **Inspector** window in Debug mode to find the internal name of a property.

**Tip:** To search for GameObjects that share a value with a GameObject you’re editing in the **Inspector** window: right-click the property in the **Inspector** window and select **Search Same Property Value**.

For more information about components and properties in the Hierarchy search provider, and how they differ from the Project search provider, refer to [Search your project and scenes](search-project-and-scenes-overview.md).

### Layers

Find GameObjects by their built-in or user-defined layer.

```
h: layer=<layer ID>
```

For example: Find all GameObjects in the Water layer, whose ID is `4`:

```
h: layer=4
```

The visual query builder shows all available layers, including the built-in ones. If you want to use a textual query, you can find the layer ID and name in the **Tags & Layers** window. For more information, refer to [Layers](Layers.md).

### Rendering Layers

Find GameObjects by their built-in or user-defined rendering layer.

```
h: renderinglayer:NameOfRenderingLayer
```

For example: Find all GameObjects in the `Default` rendering layer:

```
h: renderinglayer:Default
```

The visual query builder shows all available rendering layers, including the built-in ones. If you want to use a textual query, you can find the layer name in the **Tags & Layers** window. For more information, refer to [Introduction to Rendering Layers in URP](rendering-layers-introduction)

### Tags

Find GameObjects by their built-in or user-defined tag.

```
h: tag=<string>
```

For example, to find all GameObjects with the tag `Player`:

```
h: tag=Player
```

The visual query builder shows all available tags, including the built-in ones. If you want to use a textual query, you can find the tags in the **Tags & Layers** window. For more information, refer to <!-[Tags and Layers](class-TagManager.md).]

### Reference

Find GameObjects by the asset they reference.

| Filter option | Query example | Description |
| --- | --- | --- |
| **Reference By Asset Expression** | `h: ref={p:terrain}` | This is a nested query. It finds GameObjects that reference the result of the search for `terrain` in the Project provider. |
| **Reference By Instance ID (Number)** | `h: ref=1000` | Finds GameObjects that reference a component with the specified Instance ID. To see a component’s instance IDs, set the **Inspector** window to [Debug Mode](InspectorOptions.md). |
| **Referencing Asset (ID, path, or name)** | `h:ref:Assets/Trees/Prefab/Pine.prefab` | Find GameObjects referencing an asset by its name, exact path, or [`GlobalObjectId`](ScriptReference:GlobalObjectId). |
| **Referencing GameObject (ID, path, or name)** | `h: ref=<$object:GlobalObjectId_V1-4-0000000000000000d000000000000000-2807357573292280742-0,UnityEngine.GameObject$>` | Find GameObjects referencing another GameObject by its name, exact path, or [`GlobalObjectId`](ScriptReference:GlobalObjectId). |

If you don’t know the asset’s ID, use this filter as part of a [visual query](search-visual-query.md), where you’re able to choose the referenced asset from a list. If you need a textual query, you can convert the visual query to a textual one. For more information, refer to [Launch and use the search window](search-window-reference.md#search-mechanisms).

To get an asset’s exact path, right-click the asset in the **Project** window and select **Copy Path**.

Another way to find references is with the context (right-click) menu:

* **Project** window: right-click an asset and select **Find References In Scene** or **Find References In Project**.
* **Hierarchy** window: right-click a GameObject and select **Find References In Scene**.
* **Inspector** window: right-click a component and select **Find References In Scene**.

### Missing

Find GameObjects that are missing references such as materials or **scripts**.

Available options:

* **Any**: Any missing reference.
* **Asset**: Missing asset reference.
* **Prefab**: Missing prefab reference.
* **Script**: Missing script reference.

```
h: missing:<option>
```

For example: Find all GameObjects in the scene that are missing a prefab reference:

```
h: missing:Prefab
```

### Prefabs

Find [prefabs](Prefabs.md) or their GameObjects.

| Filter | Query example | Description |
| --- | --- | --- |
| **Any** | `prefab:any` | Find any asset (Project search provider) or GameObject (Hierarchy search provider) that is [part of a prefab](../ScriptReference/PrefabUtility.IsPartOfAnyPrefab.md). If you search for `prefab` instead of `prefab:any`, the results include anything with the string `prefab` in its name or path. |
| **Altered** | `prefab:altered` | Find [prefab instances](../ScriptReference/PrefabUtility.HasPrefabInstanceAnyOverrides.md) or variants with overrides, including default overrides. |
| **Asset** | `prefab:asset` | Find [prefab asset](../ScriptReference/PrefabUtility.IsPartOfPrefabAsset.md). |
| **Base** | `prefab:base` | Find base prefabs. |
| **Instance** | `prefab:instance` | Find objects that are part of a prefab instance. |
| **Model** | `prefab:model` | Find GameObjects that are part of a model prefab. |
| **Modified** | `prefab:modified` | Find GameObjects that are a prefab instance with overrides. |
| **Non asset** | `prefab:nonasset` | Find prefabs that aren’t prefab asset. |
| **Regular** | `prefab:regular` | Find regular prefab instances or assets. |
| **Root prefab** | `prefab:root` | Find [prefabs that aren’t children of other prefabs](../ScriptReference/PrefabUtility.IsAnyPrefabInstanceRoot.md). |
| **Top (root)** | `prefab:top` | Find [top-level prefab roots](../ScriptReference/PrefabUtility.IsPartOfPrefabInstance.md). |
| **Variant** | `prefab:variant` | Find prefab variants (Project search provider) or GameObjects from a variant (Hierarchy search provider). |

For details on the different prefab options, refer to [Prefabs](Prefabs.md).

### Scene Filters

Find GameObjects based on properties that impact their behavior in the scene.

| Filter | Query example | Description |
| --- | --- | --- |
| **Active** | `h: active=true` | Find all active GameObjects. |
| **Child** | `h: is:child` | Fine all child GameObjects. |
| **Components count** | `h: components>5` | Find GameObjects with more than five components. |
| **Hidden** | `h: is:hidden` | Find all hidden GameObjects. |
| **Instance ID** | `h: id=33714` | Find GameObjects by the instance ID of their components. To display a component’s instance IDs, set the **Inspector** window to [Debug Mode](InspectorOptions.md). |
| **Leaf** | `h: is:leaf` | Find GameObjects that don’t have children. |
| **Path** | `h: path="/ParentGameObject/ChildGameObject"` | Find GameObjects nested under the ChildGameObject. |
| **Prefab** | `h: is:prefab` | Find all prefabs. |
| **Root** | `h: is:root` | Find GameObjects that don’t have a parent. |
| **Static** | `h: is:static` | Find all static GameObjects. |
| **Visible** | `h: is:visible` | Find all visible GameObjects. |
| **Volume size** | `h: size>5` | Find GameObjects whose volume (based on their Transform Scale values) is greater than 5. |

### Custom Scene Filters

Find GameObjects by their [serialised properties](script-serialization-how-unity-uses.md).

The available filters depend on the serialized properties in your project.

### Fuzzy search

Use a fuzzy search to find items that are similar to your query, but not an exact match. For example, where the string’s characters aren’t in the same order, or have other characters between them.

Query example: To find two GameObjects, one called `score` and one called `scr`:

```
h: +fuzzy scr
```

When using a visual query, the Fuzzy Search filter is part of the **Options** filter group.

## Results

**Search** window tab: **Hierarchy**.

## Actions

The context menu for the Hierarchy search provider includes the following actions:

| Action | Description |
| --- | --- |
| **Select** | Open the GameObject in the **Inspector** window and **Scene** view. This is the default double-click action. To change the default action, refer to [Preferences](Preferences.md#search). |
| **Open** | Open the GameObject in the **Project** and **Inspector** windows. |
| **Hide** or **Show** | Toggle the visibility of the GameObject in the **Scene** view. |

## Additional resources

* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Search your project and scenes](search-project-and-scenes-overview.md)
* [How Unity uses serialization](script-serialization-how-unity-uses.md)