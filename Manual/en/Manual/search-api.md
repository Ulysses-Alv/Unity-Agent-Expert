# Execute Static API methods

Use the Static API search provider to execute Unity’s parameterless public static API methods and properties.

**Note:** You can’t implicitly use this search provider [as part of a default search](search-filter-by-provider.md). You must explicitly use its search token or its [visual query](search-visual-query.md) filter.

## Query syntax

Provider token: `api:`

Query example: Searches for methods in the [`Application`](../ScriptReference/Application.md) class, which provides access to application runtime data.

```
api:Application
```

## Provider filters

The Static API search provider doesn’t have any filters.

## Results

**Search** window tab: **Static API**.

## Actions

Default action: Executes the API method.

## Additional resources

* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)
* [The Unity scripting reference](../ScriptReference/index.md)