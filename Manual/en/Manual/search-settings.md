# Search settings and preferences

Use the Settings search provider to find settings in the [Project Settings](comp-ManagerGroup.md) and [Preferences](Preferences.md) windows.

## Query syntax

Provider token: `set:`

Query example: Searches for all light-related settings and preferences.

```
set: light
```

## Provider filters

The Settings provider has additional filters in the [visual query builder](search-visual-query.md). This table lists the filters and their textual query equivalents.

| Filter | Search token | Query example | Description |
| --- | --- | --- | --- |
| **Scope: Project** | `set: scope:project` | `set: scope:project physics` | Searches for physics settings in the Project Settings window. |
| **Scope: User** | `set: scope:user` | `set: scope:user physics` | Searches for physics settings in the Preferences window. |

## Results

**Search** window tab: **Settings**.

## Actions

Default action: Opens the page in the Project Settings or Preferences window.

## Additional resources

* [Project Settings](comp-ManagerGroup.md)
* [Preferences](Preferences.md)
* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)