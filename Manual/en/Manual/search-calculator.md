# The calculator

The Calculator search provider performs mathematical calculations that you can copy to the clipboard. You can also save the formula as a reusable search query, so it’s available to anyone in the project.

For more information about mathematical expressions, refer to [Use numeric field expressions](InspectorNumericFields.md).

**Note:** You can’t implicitly use this search provider [as part of a default search](search-filter-by-provider.md). You must explicitly use its search token or its [visual query](search-visual-query.md) filter.

## Query syntax

Provider token: `=`

**Tip:** If you leave only the Calculator search provider [active](search-manage-providers.md), you don’t have to use the provider token.

Supported operators: `+ - * / % ^ ( )`

Query examples:

* `=cos(3 * (2 * pi) * 5)`
* `=42/34 + (56 % 6)`
* `=23 * 9 ^ 3`

## Provider filters

The Calculator search provider doesn’t have any filters.

## Results

**Search** window tab: **Calculator**.

## Actions

To copy the result to the clipboard, double-click it, or right-click and select **Exec**.

## Additional resources

* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)
* [Use numeric field expressions](InspectorNumericFields.md)