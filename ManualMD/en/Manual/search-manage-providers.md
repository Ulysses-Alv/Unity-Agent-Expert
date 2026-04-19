# Activate and deactivate search providers

You can manage search providers at two levels:

* At the **Search** window level, you can toggle search providers to be active or inactive. You can open different **Search** windows with different active search providers at the same time.
* At the project level, you can add custom search providers.

## Toggle providers on and off in the Search window

You can temporarily toggle search providers off in the **Search** window. Toggling a search provider off:

* Reduces the number of results in the **All** tab, because the search provider doesn’t return results.
* Reduces the number of tabs in the **Search** window, because the inactive search provider’s tab isn’t displayed.
* Removes the inactive search provider from the [visual query builder’s](search-visual-query.md) filter dropdowns.

  **Tip:** If you leave only one provider active, you don’t have to use the search provider’s token to run queries. For example, if you leave only the Calculator search provider active, you can run calculations without prefixing them with `=`.

You can also toggle on the search providers that aren’t on by default. For more information, refer to [Use search providers](search-use-provider.md).

### Toggle providers

To toggle a search provider in the **Search** window:

1. In the **Search** window > **Results** tab, select **More** (**⋮**).
2. Select a search provider to toggle it on or off.

![The More menu with the list of search providers available in the project. A provider with a check next to its name is active in the window.](../uploads/Main/search-providers-list.png)

The More menu with the list of search providers available in the project. A provider with a check next to its name is active in the window.

**Note:** If you load a saved search query that uses an inactive provider, the search provider becomes active.

### Usage example

The toggle impacts only the current **Search** window. That means you can have multiple **Search** windows open, each with different search providers active. For example, you can:

1. Open a **Search** window for the Hierarchy search provider, showing **GameObjects**.
2. Have a second **Search** window open for the Project search provider, showing your project assets.
3. Use the Project provider to find textures and materials, and drag them to the Hierarchy search window to apply them to GameObjects.

## Add a search provider to the project

To create your own search providers, refer to [Create a custom search provider](create-search-provider.md).

## Toggle search indices

If you create a search index, you can toggle it on or off in the **Search** window. To toggle an index, from the **Search** window’s **View** (![The View menu](../uploads/Main/view_menu.png)) menu, select your index.

To create an index, refer to [Manage search indices](search-index-manager.md).

## Additional sources

* [Use Search providers](search-use-provider.md)
* [Search providers references](search-by-provider.md)
* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Manage search queries](search-manage-queries.md)
* [Create a custom search provider](create-search-provider.md)