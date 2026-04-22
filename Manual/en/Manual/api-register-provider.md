# Registering a Search Provider

To add a new Search Provider, refer to the examples in the [SearchProvider](../ScriptReference/Search.SearchProvider.md) class.

These examples create a function and tag it with the [SearchItemProvider](../ScriptReference/Search.SearchItemProviderAttribute.md) attribute:

* The function must return a new `SearchProvider` instance.
* The `SearchProvider` instance must have the following:
  + A unique `type`. For example, **Asset**, **Menu**, or ****Scene****.
  + A `displayName` to use in the [Filters pane](search-filters#persistent-search-filters).
* The optional `filterId` provides a search token for [text-based filtering](search-filters#search-tokens). For example, `p:` is the filter ID for [Asset searches](search-assets.md).

## Registering a Search Provider shortcut

To register a shortcut for a new provider, use:

```
[UsedImplicitly, Shortcut("Help/Quick Search/Assets")]
private static void PopQuickSearch()
{
    // Open Search with only the "Asset" provider enabled.
    SearchService.ShowContextual("asset");
}
```

You can map shortcuts to keys or key combinations using the [shortcuts manager](https://docs.unity3d.com/Manual/ShortcutsManager.html).