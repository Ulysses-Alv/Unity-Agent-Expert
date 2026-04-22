# Search expressions

Search expressions allow you to add to the search query language to express complex queries that cross-reference multiple providers, for example, to search for all objects in a **scene** that use a **shader** that doesn’t compile.

Search expressions can be chained together to transform or perform set manipulations on Search Items.

## How to use search expressions

Enter a search expression in the **Search** window to return a result:

![Results for the search expression count{a:assets}](../uploads/Main/search-expression-tab.png)

Results for the search expression `count{a:assets}`

If you’re using the [visual query builder](search-visual-query.md), toggle it off to enter a search expression. The toggle is located next to the search field.

![The Toggle Query Builder Mode button next to a textual query](../uploads/Main/search-toggle-query-builder-mode.png)

The Toggle Query Builder Mode button next to a textual query

## Create a search expression

A search expression starts with a root expression that can contain inner expressions.

A simple query like `t:shader` is an expression. It returns all of the Search Items corresponding to shaders in your project.

Use the search expression language to combine *multiple queries* into a single expression to create more flexible searches.

Search expressions support the following:

* query: `t:shader`
* query with a nested expression: `t:prefab ref={t:texture}`
* function: `count{t:shader}`
* literal: `3` or `"this is a string literal"`
* property selector: `@path or @mesh` (any identifier beginning with `@` is assumed to be a selector that extracts a value from a Search Item. Selectors can find properties of an Object or be used to compute results dynamically)

All search expressions return a set of Search Items (`IEnumerable<SearchItem>`). As a general guideline, the search expression language uses curly braces `{}` to denote a set of items. Even literal expressions such as `3` are evaluated internally as a set : `{3}`.

Search expressions can be chained together to make more complex queries:

* `t:prefab ref={t:texture}`
  + The `t:texture` part of the query finds all textures in your project
  + The `t:prefab ref=` part of the query runs a query for each texture returned in the t:texture set to see which **prefabs** refer to this texture.
* `t:prefab ref={t:texture size>4000}`: finds all prefabs in a project with texture size larger than 4000 bytes.

This is the equivalent to:

* Running `t:texture size>4000`to return a list of all 4K textures in the project (for example, `armor.png`, `treasure.png`).
* Running `t:prefab ref=<one of the 4K textures>` (for example, `t:prefab ref=armor.png` and then `t:prefab ref=treasure.png`).
* Aggregating all results and returning this as a single list of Search Items.
* `t:[shader, material, texture]`: finds all objects whose types are `shader`, `material`, or `texture`. Using a normal query it could be expressed as:
* `t:shader or t:material or t:texture`

## Functions

Unity has a library of search expression functions to manipulate and transform sets of Search Items. Each of these functions can take multiple arguments and returns a set of items. Unity Search Expressions use curly braces to denote function calls.

The complete list of all functions is [here](search-expression-functions-ref.md).

Example: The function [count](search-expression-functions-ref.md) counts the items of each set passed as parameters and returns a number.

* `count{t:shader}` : returns a set containing the number of shaders in the project. Example: `{34}`.
* `count{t:shader, t:texture}` : returns a set containing the number of shaders and textures in the project. Example: `{34, 2048}`.

Functions can be chained together, similar to [s-expression](https://en.wikipedia.org/wiki/S-expression) used in the [LISP](https://en.wikipedia.org/wiki/Lisp_(programming_language)) programming language.

> **Note:** There can only be a single root expression.

Example of a chain of functions being evaluated, in order of operation:

`print{"path=@path", first{10, sort{t:texture, @size}}}`: [print](search-expression-functions-ref.md) the path of the 10 smallest textures in the project

1. `t:texture`: finds all texture in the project and then selects the size property.
2. `sort{ t:texture, @size}`: sorts all these textures according to their size property.
3. `first{10`: selects the first 10 sorted textures.
4. `print{ "path=@path"`: prints this list in the console according to a format string where Unity Search extracts the path property of each resulting item `path=@path`.

Functions can have multiple signatures (name of the method and the type and kind (value, reference, or output) of each of its formal parameters, similar to C#) and can support **optional** (parameters that don’t need to be passed to the function) and [variadic](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/params) (parameters that can take a variable number of arguments) parameters.

## Literals in search expressions

Literals are actual words or numbers that you would like to search for or number amounts you want to retrieve, as opposed to the query string. For example, t:texture searches for Assets with texture in the type name (e.g.,Texture2D), but adding quotation marks to make it a literal, “t:texture” searches for an Asset named t:texture.

| Expression | Description |
| --- | --- |
| [**Number**](#number) | A literal number (`1,2,3, and so on`) |
| [**Set**](#set) | Square brackets (`[ ]`) |
| [**String**](#string) | Single or double quotes (`''` or `""`) |

### Number

Number literals can used as parameters to functions (such as `first`).

`first{10, t:shader}` -> `{the first 10 shaders return by the 't:shader' query}`

### Set

Express a set of values using square brackets `[]`. A set can contain any type of expression, but the search expression parser will assume the elements of a set are literals instead of a search query.

Example:

`[t:shader, hello, 3]` -> `["t:shader", "hello", 3]`

If using curly brackets (`{}`), the parser would treat this as three queries:

`{t:shader, hello, 3}` -> {<run a query for t:shader>, <run a query for hello>, <run a query for 3>}

### String

String literals can be used as parameters for some functions (such as `format`). You can specify a string literal using single or double quotes:

“hello” or ‘hello’

The format function taking a format string as a parameter:

`format{t:texture, '@path (@size)'}`

## Property selectors and search item values

A selector is an identifier denoted using the `@` prefix. Use selectors to access the property of an item in order to compute, filter, or format. Selectors can map to [serialized](https://docs.unity3d.com/ScriptReference/SerializedProperty.html) properties of a UnityEngine.Object or to a custom function to allow access to dynamic data.

The base selectors that any Search Items supports are:

* `id`: Unique ID of this item according to its Search Provider.
* `value`: Internal value of an item. By default it is its ID, but functions can override the value of an item.
* `label`: item label as shown in the **Search** window
* `desc` or `description` : item description (second line of the **Search** window)

Unity also defines generic selectors for a Search Item, for example:

* `@name`: UnityEngine.Object.name
* `@size`: File size on disk for an asset
* `@path`: Asset path
* `@extension`: Asset file extension
* `@provider`: Search Provider that has yielded this item

To access specific properties of a Search Item in order to perform an operation, use selectors:

* `count{t={distinct{select{a:assets, @type}}}}`
  + `a:assets, @type` finds all of the assets in a project and selects the `type` property for those objects.
  + `distinct` returns a list of all types that exist.
  + `t={list of types}` returns the multiple lists of each asset of each type.
  + `count` counts how many assets of each type are in your project.
* `avg{@size,t:shader}`
  + `t:shader, @size`: finds all the shaders and then selects the size property.
  + `avg`: computes the average size for all shaders in your project.
* `print{@value, count{t:texture, t:material}}`
  + This will print the count value extracted from the results of the count function in the **Console window**.

The `@#` selector finds serialized properties and material properties:

* `sort{t:texture, @#height}`: sorts all textures in order of their height serialized properties.

## Name search expressions (alias)

You can name search expressions to make them easier to display in the **Search** window.

For example, entering the expression `sort{count{t:audio, t:texture}, @value,desc}` in the **Search** window, it may be difficult to read which count corresponds to which type:

![Two results, both called Count](../uploads/Main/search-expression-alias1.png)

Two results, both called Count

Using an alias: `sort{count{t:audio as Audio, t:texture as Texture}, desc}` yields a more readable result:

![One result is called Texture, and the other is called Audio](../uploads/Main/search-expression-alias2.png)

One result is called Texture, and the other is called Audio

To dynamically generate alias names, use the `alias` function. For example:

`alias{[1, 2, 3], 'Title value'}`

will yield items with these labels:

`{Title 1, Title 2, Title 3}`

![The alias function gave the labels Title 1, Title 2, and Title 3 to the results](../uploads/Main/search-expression-alias-dynamic-generation.png)

The alias function gave the labels Title 1, Title 2, and Title 3 to the results

## Expand operator

The expand operator `...` allows a set of items to be grouped into multiple sets of items instead of just one set.

`...{expandable expression}` -> `{sub expr 1} {sub expr 2} {sub expr N}`

Example:

A small project has *3 prefabs, 4 textures and 5 shaders*. The following expression provides a count for all:

`count{t:[prefab, texture, shader]}` -> `{12}`

The result of the search expression `t:[prefab, texture, shader]` is a unified list of 12 items of type prefabs, textures, and shaders. This could also be expressed as: `t:prefab or t:texture or t:shader`.

Use the expand operator to count the number of each type of asset separately:

`count{...t:[prefab, texture, shader]}` -> `count{t:prefab, t:texture, t:shader}` -> `{3, 4, 5}`

The groupBy function can be used with the expand operator to return sets of items grouped by a common key.

Example:

A project has three types of assets (prefabs, textures, and shaders). For a selector `@type` that returns the type of Search Item, the expression would expand to:

`count{t:prefab or t:texture or t:shader}`

You can use the `groupBy` function to group items by their type and then count each number of items per type:

`count{...groupBy{t:prefab or t:texture or t:shader, @type}}`

## Search expression examples

These examples demonstrate how search expressions can be used for complex requests in a project.

* Count prefab usage in scenes and sort by highest usage:

  `sort{select{p: t:prefab *.prefab, @path, count{t:scene ref:@path} as count}, @count, desc}`
* Sort and count all asset types

  `sort{count{...groupby{a:assets, @type}}, @value, desc}`
* Find the **mesh** with the most vertices (assuming the @vertices selector is available in the project).

  `first{sort{t:mesh, @vertices, desc}}`
* Sort all meshes by their vertex count

  `sort{select{h: t:mesh, @path, @vertices}, @vertices, desc}`
* Count the number of vertices for all meshes

  `sum{select{h:t:mesh, @vertices}}`
* Find all assets referencing a mesh

  `ref=select{p:t:mesh, @path}`
* List prefab scene references count

  `select{p: *.prefab, @path, count{p: a:sceneIndex ref="@path"} as count}`

## Expression evaluation

Search expression evaluation is based on the C# [iterator](https://docs.microsoft.com/en-us/dotnet/csharp/iterators) pattern and uses of the [yield](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/yield) keyword. This ensures that the evaluation is exposed as a fully asynchronous operation and is as non-blocking as possible. It allows each search expression to start its computation before the initial sets of search items are fully computed.

By default all search expression evaluation is done in a worker thread in order to not block the Editor. For operations that need to rely on non-thread safe Unity API (like some selector accessors), we have a utility API to queue those operations on the main thread and maintain the pattern of *yielding* search items.

## Customize the search expression language

The search expression language has been designed to be customizable. APIs to customize all parts of the framework will be available in a future release.

## Additional resources

* [Search provider references](search-by-provider.md)
* [Create a custom search provider](create-search-provider.md)
* [Visual query builder reference](search-visual-query.md)
* [Launch and use the Search window](search-window-reference.md)