# Use search providers

All searches use at least one search provider:

* The default providers are Project, Hierarchy, Menu, and Settings. If you run a search without specifying a provider, Search uses these default providers.
* Alternatively, you can focus the search to one provider. This can be one of the default search providers, or any other search provider available in your project.

  For example, if you need to add assets to your project, you can focus your search to the **Asset Store** search provider. This focus excludes the project’s current assets from your search results.

**Tip:** You can have multiple **Search** windows open, each with different providers active. This makes it easier to perform queries, because you don’t need to explicitly specify your provider in each query. For example, you can have a **Search** window that uses just the Calculator search provider, so you don’t have to prefix your queries with `=`.

![The Search windows More menu with the list of search providers available in the project](../uploads/Main/search-providers-list.png)

The Search window’s More menu with the list of search providers available in the project

For information about individual search providers, refer to [Search provider references](search-by-provider.md).

## The purpose of search providers

Use search providers to narrow your search results to a certain type of item. Providers have filters that match the item type they search for, so selecting a provider also narrows your filter list to only the relevant ones.

You can divide the search providers by their purpose:

| Purpose | Examples |
| --- | --- |
| [**Search your project’s assets**](search-project-and-scenes-overview.md) | Project, Hierarchy, Asset Database. |
| [**Search files and presets**](search-files-presets-overview.md) | Files, Presets. |
| [**Search external sources**](search-external-sources-overview.md) | Asset Store, the Packages Registry. |
| [**Search performance and debugging**](search-performance-and-debugging-overview.md) | **Profiler Markers**, Performance Trackers, Logs. |
| [**Search the Unity Editor interface and actions**](search-interface-and-actions-overview.md) | Settings, Menus, Static API, Calculator. |

Some search providers fit into more than one set. For example,:

* The Packages search provider can fetch items from the Package Manager registry, but it also searches your project for installed packages.
* The Menus search provider focuses on the Unity Editor’s menus, but it presents its results as actions that you can execute with a double-click.
* You can include the Files search provider in other search providers.

## Custom search providers

You can create your own search providers. For more information, refer to [Create a custom search provider](create-search-provider.md).

## Search with all active providers at once

By default, a search includes results from four providers:

* Project: Searches the project’s assets.
* Hierarchy: Searches **GameObjects** in the loaded **scenes** (a loaded scene is displayed in the **Hierarchy** window).
* Menu: Search the Unity Editor main menu.
* Settings: Search all Unity Editor settings.

You can add and remove providers from the default search by toggling them on and off. For more information, refer to [Activate and deactivate search providers](search-manage-providers.md)

## Search a single provider

You can limit your search to a single provider for faster queries and a shorter results list.

You can add and remove providers from this search by toggling them on and off. For more information, refer to [Activate and deactivate search providers](search-manage-providers.md)

### Use a provider filter in visual queries

When you use the visual query builder, the first filter’s dropdown includes all available search providers.

**Note:** The search provider selection must be the first filter; you can’t use it anywhere else in the query.

If you select a search provider, all other filters in the query match that provider. For example, if you select the Menus search provider, the next filter dropdown doesn’t include filters for Profiler Markers or Logs.

For more information, refer to [Visual query builder](search-visual-query.md).

### Use a provider token in a textual query

To use a search provider in a textual query, prefix the query with the provider’s search token.

**Note:** The search provider’s token must be the first token; you can’t use it anywhere else in the query.

For example, `m:` is the search token for the Menus search provider. When you enter `m:light` in the search field, search returns all menu items that contain the word `light`.

For more information, refer to [Textual query references](search-textual-query.md).

## Additional sources

* [Search providers references](search-by-provider.md)
* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Manage search queries](search-manage-queries.md)
* [Create a custom search provider](create-search-provider.md)