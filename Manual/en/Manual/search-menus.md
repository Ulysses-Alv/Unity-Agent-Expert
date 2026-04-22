# Search the Unity Editor main menu

Use the Menus search provider to access the Unity Editor main menu commands.

## Query syntax

Provider token: `m:`

Query example: List all menu commands related to splines, from all submenus.

```
m: splines
```

## Provider filters

The Menus provider has additional filters in the [visual query builder](search-visual-query.md). This table lists the filters and their textual query equivalents.

| Filter | Search token | Query example | Description |
| --- | --- | --- | --- |
| **Menu Path** | `m: id:<"path">` | `m: id:"2d/tiles"` | Searches only within the **2D** > **Tiles** menu. |

## Results

**Search** window tab: **Menus**.

## Actions

Default action: Executes the menu command.

Notes:

* For commands like **Create**, make sure you’re in the correct folder in the **Project** window before executing the command.
* For commands that interact with GameObjects in the **Hierarchy** window, make sure to select the correct **GameObject** in the **Hierarchy** window before executing the command, or the command might work at the **scene** level. For example, if you execute **Create Empty Child** while your focus is in the **Project** window, the Unity Editor creates the GameObject at the scene level rather than as a child of a GameObject.
* Some commands, like **Duplicate**, don’t execute from the **Search** window.

## Additional resources

* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)