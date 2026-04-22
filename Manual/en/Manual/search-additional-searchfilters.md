# Additional Search filters

This page reviews additional search filters that you can use to narrow down your search results.

## Combine search tokens for complex searches

You can combine search tokens to create more complex queries.

* The queries are written on one line with one character space between tokens.
* The character space between each new token is an “And” operation, so both filters must be true for the query to return a result. Add another [operator](search-query-operators.md) (or, <, >) to return different results.
* If a Search Provider filter token (h:, p:) is used, it must be the first component in the query.

Here are a few examples:

| Query | Description |
| --- | --- |
| `h: t:meshrenderer p(castshadows)!="Off"` | Searches all static meshes in a **Scene** that cast a shadow. |
| `h: t:light p(color)=#FFFFFF p(intensity)>7.4` | Searches all lights in a Scene with a specific color with brightness higher than 7.4. |
| `h: path:/Collectables t:collectable` | Find all objects with a component `Collectable` located in the path `/Collectables.` |

## Cross-reference across multiple providers with search expressions

Search expressions allow you to add to the search query language to express complex queries that cross-reference multiple providers, for example, to search for all objects in a scene that use a **shader** that doesn’t compile. Refer to [Search expressions](search-expressions.md) for more information.

## Type Filters

These filters are available if the index uses the Types Indexing option (See [Index Manager](search-index-manager.md)).

| Filter | Search token | Description |
| --- | --- | --- |
| **Type** | `t:<Asset Type>` | `t:texture`  Searches all Assets containing texture in their type name (ex: Texture2D, Texture).  `t:prefab`  Searches all prefab assets. |
| **Type** | `<Asset Type>` | You can search Assets by type without using the t: filter above.  `texture`  Searches all Assets containing texture in their type name (ex: Texture2D, Texture) or in their name (ex: myTexture.png).  `prefab`  Searches all Prefab Assets or Assets with “prefab” in their name. |
| **File** | `t:file` | `t:file level1`  Searches all file Assets containing the word level1. |
| **Folder** | `t:folder` | `t:folder`  Searches all folder Assets. |

## Indexed Property Search

Searching properties is available if the index has been specified with the Properties Indexing option (see [Index Manager](search-index-manager.md)). To view the list of all indexed properties check the Index Manager Keywords tab.
All property values are converted to a string or number.
The name of the property has to be complete and not partial (case does not matter though).
Unity indexes properties at the top level object of a **Prefab** Asset. If you want all Prefab hierarchies to be indexed, create a Prefab Index (see [Index Manager](search-index-manager.md)).
For .unity files Unity indexes the properties of the SceneAsset and not the Scene content. If you want all Scene contents to be indexed, create a Scene Index (see [Index Manager](search-index-manager.md)).

| Filter | Search token | Description |
| --- | --- | --- |
| **Type** | `t:<type>` | When using the indexed Property, you can use `t:` to search for a component type for an Asset type.  `t:collider`  Searches all Prefabs containing a component with the word `collider`.  `t:texture`  Searches all Assets with a type containing the word `texture` (ex: Texture or Texture2D). |
| **Has Component** | `t:<component type>` | `t:collider`  Searches all Prefabs containing a component with the word collider.  `t=BoxCollider`  Searches all Prefabs containing a component called BoxCollider. |
| **Label** | `l:<label name>` | `l:archi`  Searches all Assets with a label containing the word `archi` (e.g.: Architecture).  `l=Wall`  Searches all Assets with a label that is exactly `Wall`. |

All properties of an Asset (Prefab or other types) are indexed and searchable. Here are a few examples of Property query:

| Filter | Search token | Description |
| --- | --- | --- |
| **Number** | `property:value` | `bounciness>0.1`  Searches all assets with a property named `bounciness` (e.g.: a PhysicsMaterial2D) higher than 0.1.  `health=2`  Searches all Assets with a property named `health` (e.g.: HealthSystem Component of a Prefab) with of a value of exactly 2.  `t:texture filtermode!=0`  Searches all textures with a `filtermode` property different than 0 (i.e different than Point). |
| **Boolean** | `property:value` | `t:Dungeon generatePath=true`  Searches all Dungeon ScriptableObjects where the property `generatePath` is true.  `isStunned=false`  Searches all GameObjects containing a property `isStunned` that is false. |
| **String** | `property:string value` | `t:Character trait:indestru`  Searches all Prefab with a Character component whose trait property contains the word `indestru` (ex: indestructible).  `t:Character trait="tough but fair"`  Searches all Prefab with a Character component whose trait property is exactly `tough but fair`. |
| **Enum** | `property:<enum value>` | `characterclass:rog`  Searches all GameObjects with with a property named `characterclass` whose value contains the word rog (e.g.: value is rogue).  `characterclass=FighterMage`  Searches all GameObjects with a property named `characterclass` with an exact value of `FighterMage`. |
| **Color** | `property:<html color value>` | `color:ADA`  Searches all GameObjects with a property named `color` where the color value starts with ADA (like ADADAD00).  `color=ADADAD00`  Searches all GameObjects with a property named `color` where the color value is exactly ADADAD00.  `color=ADADAD`<br/> Searches all GameObjects with a property named color where the color value is exactly ADADAD and alpha value is 1. |
| **Vector** | `property.[xyzw]:value` | `bounds.x>1`  Searches all GameObjects with a property named `bounds` where the x value is larger than 1.  `acceleration.z=2`  Searches all GameObjects with a property named acceleration where the z value is equal to 2 |
| **Object** | `sprite:<object exact name>` | `sprite:CharacterBody`  Searches all Assets with a `sprite` property (e.g.: Image Component of a Prefab) that references a GameObject named CharacterBody. |

## Dependency filters

If you are using the Dependencies Indexing option (See [Index Manager](search-index-manager.md)) Unity indexes **direct dependencies** of all Assets using [AssetDatabase.GetDependencies](../ScriptReference/AssetDatabase.GetDependencies.md).

| Filter | Search  token | Description |
| --- | --- | --- |
| **Reference Path** | `ref:<asset full path>` | `ref:assets/images/particles/p_smoke.png` Searches all Assets that have direct dependencies on the **exact** Asset path: `assets/images/particles/p_smoke.png` . |
| **Reference Name** | `ref:<asset name>` | `ref:p_smo` Searches all Assets that have direct dependencies on an Asset whose name contains the word `p_smo`.   `ref:p_smoke.png` Searches all Assets that have direct dependencies on an Asset named `p_smoke.png`. |

## Search for null values

Search any object property for a `null` value using the `none` keyword:

```
h:#MeshFilter.m_mesh=none
```

## Additional resources

* [Search provider references](search-by-provider.md)
* [Create a custom search provider](create-search-provider.md)
* [Visual query builder reference](search-visual-query.md)
* [Launch and use the Search window](search-window-reference.md)