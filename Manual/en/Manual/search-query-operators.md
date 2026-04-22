# Search query operators

The Scene, Asset, Objects, and Resource providers use a [`QueryEngine`](../ScriptReference/Search.QueryEngine.md) ([Scene](#scene-provider), [Asset](#asset-provider), [Objects](#objects-provider) and [Resource](#resources-provider) providers) to parse and resolve their queries. They support a basic set of Query Operators that allow for more complex queries using Boolean operators and parentheses grouping. The table on this page lists the supported Query Operators.

**Case sensitivity**: Most Search queries ignore case. For example, `Stone` or `stone` or `sToNe` will yield the same results.

| **Filter** | **Search token** | **Description** |
| --- | --- | --- |
| **Basic search** | `<any partial name>` | `main` Searches matching the word `Main` |
| **And** | `and` | `Main and t:camera` Search where name contains `Main` and is of type name containing `camera`  `t:texture and jpg` Searches all of type `texture` containing `jpg` in their filename.   Note: Since `and` **is the default operator of the QueryEngine** the last query is equivalent to: `t:texture jpg` |
| **or** | `or` | `Player or Monster` Searches containing the word `Player` or `Monster`. |
| **Group** | `(<group content>)` | `t:Character and (status=Poison or status=Stunned)` Searches for a `Character` component where the value of `status` property is either `Poison` or `Stunned`. |
| **Exclusion** | `-<Expression to exclude>` | `p: dep:door -t:Scene` Searches all Assets with a dependency on an Asset containing the word `door` and that are not of type `Scene`.   `p: dep:door -stone` Searches all Assets with a dependency on an Asset containing the word `door` and that do not contain the word `stone`. |
| **Exact Operator** | `!<something>` | Most of the string matching in Search returns partial matches. Use the `!` operator to return **exact** matches.  `p: stone` Searches all Assets containing the word stone (`stone_hammer.png`, `stone_door.prefab`).  `p: !stone` Searches all Assets with the exact name `stone` (e.g.: `stone.png`) |
| **Partial Value match (:)** | `property:<partial value>` | `ref:aster` Because `:` is used, searches all Assets having an Asset containing the word `aster` (ex: `asteroid2`, `asteroids`) as a dependency. |
| **Exact Value (=)** | `property=exactValue` | `ref=asteroid` Because `=` is used, searches all Assets having an Asset with the exact name `asteroid` as a dependency. |
| **>** | `property>number` | `t:texture size>256` Searches all textures with a size bigger than 256 bytes. |
| **<** | `property<number` | `t:texture size<256` Searches all textures with a size smaller than 256 bytes. |
| **!=** | `property!=number` | `t:texture size!=256` Searches all textures with a size different than 256 bytes. |
| **>=** | `property>=number` | `t:texture size>=256` Searches all textures with a size bigger or equal than 256 bytes. |
| **<=** | `property>number` | `t:texture size<=256` Searches all textures with a size smaller or equal than 256 bytes. |

## Additional resources

* [Focus searches with search providers](search-filter-by-provider.md)
* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Launch and use the Search window](search-window-reference.md)