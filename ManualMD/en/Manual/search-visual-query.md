# Visual query builder reference

The **Search** window’s visual query builder supports the same search tokens and syntax as the textual search, but relies on a visual interface.

## Toggle the visual query builder

To use the visual query builder, in the **Search** window, select **Toggle Query Builder Mode**.

![The Toggle Query Builder Mode button next to a textual query](../uploads/Main/search-toggle-query-builder-mode.png)

The Toggle Query Builder Mode button next to a textual query

The search field changes to a visual query builder:

![The search field with a visual query](../uploads/Main/search-visual-query-material.png)

The search field with a visual query

To return to the textual search, select **Toggle Query Builder Mode** again.

## Convert a textual query to a visual query

If you have a saved textual search query, use the **Search** window to convert it to a visual query:

* Paste the textual query into the search field, or load it from your [saved searches](search-manage-queries.md).
* Select **Toggle Query Builder Mode**.

The visual query builder converts the textual query to a visual query.

## Build a query

To build a query:

1. Select the **+** dropdown menu in the search field.
2. If you want to limit your search to a specific search provider rather than use the default providers, or if you want to use a search provider that must be manually invoked (such as the **Asset Store** provider), select the search provider from the dropdown menu. The search provider appears in the search field, and a new dropdown menu appears next to it.
3. Select or search for filters in the dropdown menu. The filter you select appears in the search field, and a new dropdown menu appears next to it.
4. Continue adding filters to build your query.

   Each filter you add narrows the filters for the next element. For example, if you select a [search provider](search-by-provider.md), the next element only shows filters relevant to that provider.

## Use operators

Use operators to connect filters and filter values in your query.

### And/Or between filters

By default, filters have an `and` operator to connect them. To use the `or` operator:

1. Select the **+** dropdown menu in the search field.
2. From the **Operators** category, select the operator you want to use. The operator appears in the search field, and a new dropdown menu appears to continue building your query.

### Filter value operators

Each filter also has an operator that specifies how to treat the value of the filter, such as `Equals`, `Contains`, or `Greater Than`. To change the operator:

1. Right-click on the filter.
2. From the **Operator** menu, select the operator you want to use.

For a full list of operators, refer to [Search query operators](search-query-operators.md).

## Exclude (Not) by value

You can use a search filter to exclude results that match the filter’s value:

1. Right-click on the filter.
2. Select **Exclude**.

## Enable or disable a filter

You can disable a filter in your query without removing it. To disable a filter:

1. Right-click on the filter.
2. Deselect **Enable**.

To re-enable the filter, right-click on the filter and select **Enable**.

## Use text in a visual query

You can still use a plain-text search query in the visual query builder. To add a plain-text element to your query, type the text in the search field.

![A visual query with a single plain-text element.](../uploads/Main/search-visual-query-with-text.png)

A visual query with a single plain-text element.

## Additional resources

* [Focus searches with search providers](search-filter-by-provider.md)
* [Textual query references](search-textual-query.md)
* [Launch and use the Search window](search-window-reference.md)
* [Search query operators](search-query-operators.md)