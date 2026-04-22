# Search for packages

The Packages search provider finds packages in your project and in the Unity Registry.

## Query syntax

Provider token: `pkg:`

**Note:** You can’t implicitly use this search provider [as part of a default search](search-filter-by-provider.md). You must explicitly use its search token or its [visual query](search-visual-query.md) filter.

Query example: Searches the Package Manager for the text `probu` to find the ProBuilder package.

```
pkg: probu
```

## Provider filters

The Packages provider has no additional filters.

## Results

**Search** window tab: **Logs**.

## Actions

The context menu for the Packages search provider includes the following actions:

| Action | Description |
| --- | --- |
| **Open** | Opens the package in the **Package Manager** window. This is the default double-click action. To change the default action, refer to [Preferences](Preferences.md#search). |
| **Install** | Install the package; the Unity Editor requests confirmation before installing. |
| **Remove** | Uninstalls the package from the Project. |

## Additional resources

* [Package Manager](class-PackageManager.md)
* [Visual query builder](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)