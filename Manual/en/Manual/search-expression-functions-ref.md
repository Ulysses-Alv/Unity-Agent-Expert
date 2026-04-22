# Functions reference

This page explains the functions you can use in search expressions. Functions are used to manipulate search results and perform operations on them. The functions are listed by category and in alphabetical order:

* [Functions by category](#functions-by-category)
* [Functions in alphabetical order](#functions-in-alphabetical-order)

## Functions by category

### Set manipulation

| Expression | Description |
| --- | --- |
| [**distinct**](#distinct)  {Variadic | AnyExpression} | Produces the unique set union of multiple expressions. |
| [**except**](#except)  {Iterable, Variadic | Iterable} | Produces the set difference of two sequences. |
| [**first**](#first)  {Number, Variadic | Iterable} | Returns the first result for each expression. |
| [**groupBy**](#groupby)  {Iterable, Optional | Selector} | Group search results by a specified `@selector` | Hints: ExpandSupported |
| [**intersect**](#intersect)  {No signature} | Produces the set intersection of two sequences. |
| [**last**](#last)  {Number, Variadic | Iterable} | Returns the last result for each expression. |
| [**map**](#map)  {Iterable, AnyValue} | Maps an expression result to form a new expression using variables. |
| [**select**](#select)  {No signature} | Creates new results by selecting which value and property to take. |
| [**sort**](#sort)  {No signature} | Sorts expression results based on a criteria. |
| [**union**](#union)  {Variadic | AnyExpression} | Produces a unique set union of multiple expressions. |

### Set filtering

| Expression | Description |
| --- | --- |
| [**eq**](#eq)  {Iterable, Selector, Literal | QueryString} | Keeps search results that have an equal value. |
| [**gt**](#gt)  {Iterable, Selector, Literal | QueryString} | Keeps search results that have a greater value. |
| [**gte**](#gte)  {Iterable, Selector, Literal | QueryString} | Keeps search results that have a greater or equal value. |
| [**lw**](#lt)  {Iterable, Selector, Literal | QueryString} | Keeps search results that have a lower value. |
| [**lwe**]((#lwe))  {Iterable, Selector, Literal | QueryString} | Keeps search results that have a lower or equal value. |
| [**neq**](#neq)  {Iterable, Selector, Literal| QueryString} | Keeps search results that have a different value. |
| [**where**](#where)  {Iterable, Text | Selector | QueryString} | Excludes search results for which the expression is not valid. |

### Math

| Expression | Description |
| --- | --- |
| [**avg**](#avg)  {Selector, Variadic | Iterable} | Finds the average value for each expression. |
| [**count**](#count)  {Variadic | Iterable} | Counts the number of results in an expression. |
| [**max**](#max)  {Selector, Variadic | Iterable} | Finds the maximum value for each expression. |
| [**min**](#min)  {Selector, Variadic | Iterable} | Finds the minimal value for each expression. |
| [**sum**](#sum)  {Selector, Variadic | Iterable} | Computes the sum value for each expression. |

### Utility

| Expression | Description |
| --- | --- |
| [**alias**](#alias)  {Iterable, Text | Selector | Iterable} | Assigns an alias to an expression. | Hints: AlwaysExpand |
| [**constant**](#constant)  {Literal | Function} | Returns the first result for each expression. |
| [**format**](#format)  {Text | Selector, Variadic | AnyValue} | Converts arguments to a string allowing you to format the result. |
| [**print**](#print)  {Text | Selector, Variadic | Iterable} | Prints expression results in the console. |
| [**query**](#query)  {No signature} | Hints: ExpandSupported |
| [**random**](#random)  {Variadic | Iterable} | Randomly picks a result for each expression. |
| [**set**](#set)  {Variadic | Iterable} | Returns a set of elements from any expression. |
| [**text**](#text)  {No signature} | Converts the text of any expression to a literal string. |

## Functions in alphabetical order

### alias

`alias{set, aliasName} -> {each element in set gets assigned an aliasName}`

`alias{set, formatString} -> {each element in set gets assigned an alias computed from the formatString}`

This function effectively assigns an alias to an expression. This is equivalent to using the “as” keyword. See also [alias usage](search-expressions.md#alias-usage).

Example:

`count{t:material as material}` is equivalent to:

`count{alias{t:material, material}}`

Example:

Assign the alias “number” to the elements: 1, 2 and 3.

`alias{[1, 2, 3], 'number'}`

Example:

Assign and alias dynamically computed from a format string to 1,2 and 3:

`alias{[1,2,3], 'number @value}'}`

### avg

`avg{selector, sets...} -> {Average of selector value in each set}`

`avg{set...} -> {Average of @value in each set}`

Find the average value of each item in each set passed as parameters.

Example:

`avg{[1,2,3,4, 5,6], [1, 1, 1], []} -> {3.5, 1, 1}`

Example:

`avg{@size, t:texture, t:material} -> {1100576, 3618}`

### constant

`constant{value} -> {constant value}`

Attempts to convert the value to a constant literal value: either a number, string, or Boolean. This can use used to disambiguate parameter types in some functions.

Example:

`first{constant{5}, t:material}`

### count

`count{sets...} -> {count of each set}`

Counts the number of results in each set passed as parameters.

Example:

`count{t:texture, t:material} -> {359, 194}`

Example:

From the index named “assets”, group assets by type and count each of those asset groups. Take the five largest groups.

`first{5, sort{count{...groupBy{a:assets, @type}}, @value, desc}}`

### distinct

`distinct{sets...} -> {Unified set of all unique elements in all sets}`

Distinct creates a set of all elements found in all sets passed as parameters. Like union, no duplicate elements are kept.

Example:
This example shows that duplicates are removed.

`distinct{[1,2,3], [3,4,5]} -> {1,2,3,4,5}`

Example:

`distinct{*.mat, t:shader} -> {all materials and shaders in the project}`

Example:

Select the `@type` of each asset in the index named “project” and keep a single asset of each type since duplicates are computed according to the `@value` of an item.

`distinct{select{a:project, @type}}`

### eq

`eq{set, value} -> {all elements equal to value}`

`eq{set, selector, value} -> {all selector value equal to value}`

Filters a set of element by keeping those equal to a given value.

Example:

`eq{[2,4,5,6,4,3], 4} -> {4,4}`

this is equivalent to:

`where{[2,4,5,6,4,3], “@value=4”}`

Example:

Find all textures with a `#width` serialized property of 256.

`eq{t:texture, #width, 256}`

this is equivalent to

`t:texture #width=256`

### except

`except{set, sets...} -> {New set of elements}`

Except creates a set of all elements that are not in any of the sets passed as parameters.

Example:

`except{[1,2,3,4,5], [2,3], [5], [6]} -> {1,4}`

Example:

`except{t:prefab, t:prefab ref=select{t:texture, @path}}`

### first

`first{sets...} -> {all first element in each sets}`

`first{count, sets...} -> {first count elements in each sets}`

First returns a set of the first element found in each set passed as a parameter. If `count` is passed as parameter, will take the first count element in each set.

Example:

`first{[3, 4,5], [9, 28, 3]} -> {3,9}`

Example:

`first{3, [9, 28, 3, 4, 6, 7, 8, 9], [1, 2, 3, 4, 5, 6, 7]} -> {9, 28, 3, 1, 2, 3}`

Example:

Sort all .fbx files in the project by descending order of size and show the three largest.

`first{3,sort{*.fbx, @size, desc}}`

### format

`format{expression} -> {set of items where the value is converted to a string}`

`format{formatString, expression} -> {set of items where the value is set by applying a format string}`

Format can be used in two ways. When only a single expression argument is used it will try to convert the current `@value` of an item to a string representation.

If format is used with a format string, it will replace all selectors in the format string with the selected value of the item and set this formatted string in the `@value` of the item.

Example:

`print{format{“value=@value”, [1,2,3,4]}}` would print [“value=1”,”value=2”,”value=3”,”value=4”] in the **Console window**.

### groupBy

`groupBy{set, selector}` -> `{multiple sets of items grouped by their selector}`

Example:

To count the number of assets for each type in a project, the `groupBy` function coupled with the Expand Operator allows you to automatically group items according to a selector. In the example below it creates a set of items of the same types. The Expand Operator allows you to expand this set of items into multiple sets of the same type.

`count{...groupBy{a:assets, @type}}` -> `{numberOfAssetType1, numberOfAssetType2, .. numberOfAssetTypeN}`

### gt

`gt{set, value} -> {all elements greater than value}`

`gt{set, selector, value} -> {all elements with selector value greater than value}`

Filters a set of elements by keeping those greater than a given value.

Example:

`gt{[2,4,5,6,4,3], 4} -> {5,6}`

### gte

`gte{set, value} -> {all elements greater or equal than value}`

`gte{set, selector, value} -> {all with a selector value greater or equal than value}`

Filters a set of element by keeping those greater than or equal to a given value.

Example:

`gte{[2,4,5,6,4,3], 4} -> {4,5,6,4}`

### intersect

`intersect{sets...} -> {elements that are in all sets}`

`intersect{sets…, selector} -> {elements whose selector value is contained in all sets}`

Intersect creates a new set of elements whose `@value` is contained in all sets. If a selector is passed as parameter the selector value is used to compare elements.

Example:

`intersect{[1,2,3,4,5], [3, 5], [6, 3]} -> {3}`

Example:

Find all textures whose size is bigger than 4000 bytes and intersect with all textures whose path contains the word “rock”.

`intersect{t:texture size>4000, t:texture @path:rock}`

### last

`last{sets...} -> {all last element in each sets}`

`last{count, sets...} -> {last count elements in each sets}`

Last returns a set of the last element found in each set passed as parameter. If `count` is passed as parameter, will take the last count elements in each set.

Example:

`last{[3, 4,5], [9, 28, 3]} -> {5, 3}`

Example:

`last{3, [9, 28, 3, 4, 6, 7, 8, 9], [1, 2, 3, 4, 5, 6, 7]} -> {7,8,9,5,6,7}`

Example:

Sort all .fbx files in the project by ascending order of size and take the last three.

`last{3,sort{*.fbx, @size}}`

### lw

`lw{set, value} -> {all elements less than value}`

`lw{set, selector, value} -> {all elements with selector value less than value}`

Filters a set of elements by keeping those less than a given value.

Example:

`lw{[2,4,5,6,4,3], 4} -> {2,3}`

### lwe

`lwe{set, value} -> {all elements less or equal than value}`

`lwe{set, selector, value} -> {all with a selector value less or equal than value}`

Filters a set of elements by keeping those less than or equal to a given value.

Example:

`lwe{[2,4,5,6,4,3], 4} -> {2,4,3,4}`

### map

`map{set, <string containing selectors to replace>}`

The map operator is a functional way of expressing a [search expression](#query). It maps an expression result to form a new expression using variables.

The following example:

`map{t:texture, t:prefab ref=@id}` -> `{t:prefab ref=<textureId1>} {t:prefab ref=<textureId2>} ... {t:prefab ref=<textureIdN>}`

will run the query `t:prefab ref=@id` for each of the textures returned by `t:texture`.

This is equivalent to running:

`{t:prefab ref={t:texture}}`

### max

`max{sets...} -> {max value element in each set}`

`max{selector, sets...} -> {max selector value element in each set}`

Finds the maximum `` @value` for each set passed as parameters. If a selector is passed as a parameter, find the element with maximum ``@selector` value.

Example:

`max{[1,2,3,4,5,6], [2,4,5]} -> {6, 5}`

Example:

Find the largest .png file and the largest .jpg file.

`max{@size, *.png, *.jpg} -> {<largest .png file>, <largest .jpg file>}`

### min

`min{sets...} -> {min value element in each set}`

`min{selector, sets...} -> {min selector value element in each set}`

Finds the minimum `` @value` for each set passed as parameters. If a selector is passed as parameter, find the element with minimum ``@selector` value.

Example:

`min{[1,2,3,4,5,6], [2,4,5]} -> {1, 2}`

Example:

Find the smallest .png file and the largest .jpg file.

`min{@size, *.png, *.jpg} -> {<smallest .png file>, <smallest .jpg file>}`

### neq

`neq{set, value} -> {all elements not equal to value}`

`neq{set, selector, value} -> {all selector value not equal to value}`

Filters a set of elements by keeping those not equal to a given value.

Example:

`neq{[2,4,5,6,4,3], 4} -> {2,5,6,3}`

### print

`print(FormatString | selector, set)`

Print takes a format string (or a selector) and a set of items and prints the formatted results of each item in the console. This can be useful to debug the values of an item.

In a project, the following expression returns the textures and their paths in the Search window and prints them in the Console window:

`print{'path=@path', t:texture}`

Example:

`print{@value, [1,2,3,4]}` would print [1,2,3,4] in the Console window.

### query

`query{value} -> {query expression}`

Attempts to convert value to a query expression. This can use used to disambiguate parameters type in some specific functions.

Example:

`count{t:material}` is equivalent to

`count{query{t:material}}`

### random

`random{sets...} -> {random elements in each set}`

Creates a set out of random elements taken from each set of items passed as a parameter.

Example:

`random{[1,3,5], [6,4,2]} -> {3, 2}` returns a random results in each set.

### select

`select(<set>, selectorForLabel, selectorForDesc, otherSelectors...)`

Select is a function that creates a new set of search items and extracts or transforms properties from an original set of search items.

The second parameter (selectorForLabel) of the select function assigns the selected value to the label of the new item.

The third parameter (selectorForDesc) assigns the selected value to the description of the new item.

The last selector in a select function also specifies the value of the item

Example:

When searching for textures in a project, using `select`, generates a set of items where the label is the path and where the description becomes the size of the texture:

`select{t:texture, @path, @size}`

![Search results show the files path and size](../uploads/Main/search-expression-select-1.png)

Search results show the file’s path and size

The last selector in a select statement also sets the value of the item. For the example `select{t:texture, @path, @size}` this means `@size` will also become the `@value` of the search item. Switch to the Table view to see how the search expression replaces the label and description sets the value:

![Search results show the path in the label column](../uploads/Main/search-expression-select-2.png)

Search results show the path in the label column

Selectors of a `select` function can also be functions themselves. Selectors in these functions are evaluated according to the search item being processed.

Example:

`select{t:texture, @path, count{t:prefab ref=@id}}`

1. For each texture in the project, create a new search item (`t:texture`)
2. whose label is equal to the `@path` of the texture (`@path`)
3. and whose description is equal to how many **prefabs** refer to that texture (`count{t:prefab ref=@id}`). In the example, `@id` refers to the `SearchItem.id` of each texture item.

You can assign a property value to a search item field if you use the alias operator.

Example:

`print{"@pow", select{t:script, @path as pow}}`

This extracts the `@path` property from a search item and assigns it to the search item field named `pow` in the new search item created by the `select` function.

### set

`set{sets...} -> {set of all elements in all sets}`

Creates a set of all elements contained in all the sets passed as parameter. This is equivalent to using the “[set, of , element]” notation.

Example:

`set{[1, 2,3], [4,5,6]}` is equivalent to:

`[[1,2,3], [4,5,6]]` which is equivalent to:

`[1,2,3,4,5,6]`

### sort

`sort{set, selector} -> {set of items sorted in ascending order according to selector value}`

`sort{set, selector, desc} -> {set of items sorted in descending order according to selector value}`

Sorts a set of items according to a selector value in descending or ascending order.

Example:

Sort this set in ascending order (from smallest to largest).

`Sort{[4,2,7,5,8,1,9], @value} -> {1,2,4,5,7,8,9}`

Example:

Sort this set in descending order (from largest to smallest).

`sort{[4,2,7,5,8,1,9], @value, desc} -> {9,8,7,5,4,2,1}`

Example:

Sort all textures in my project according to their @size in descending order (from largest to smallest).

`sort{t:texture, @size, desc}`

### sum

`sum{sets...} -> {sum of all elements in each set}`

`sum{selector, sets...} -> {sum of all elements in each set}`

Find the sum ``@value` of each item in each set passed as parameters. If a selector is passed as parameter, find the sum `@selector` value of each element in each set.

Example:

`sum{[1,2,3,4,5,6], [2,4,5]} -> {21, 11}`

Example:

Find the sum of all textures `@size` in a project.

`sum{@size, t:texture}`

### text

`text{expression} -> {“text of the expression”}`

Creates a string out of an expression. This is equivalent to using the “” or ‘’ delimiters.

Example:

`[text{hello}]` is equivalent to

`[“hello”]` or `[‘hello’]`

### union

`union{sets...} -> {Unified set of all unique elements in all sets}`

Union creates a set of all elements found in all sets passed as a parameter. Like `distinct`, no duplicate of elements are kept.

Example:

This example shows that duplicates are removed.

`union{[1,2,3], [3,4,5]} -> {1,2,3,4,5}`

Example:

`union{*.mat, t:shader} -> {all materials and shaders in the project}`

### where

`where{set, filterString | selector} -> {filtered set of items}`

Where is a general filtering function that accepts a selector or a string containing selector usage and operators (>, ==, …) to filter elements of a set and returns the new filtered set.

Example:
Filtering out numbers according to their value:

`where{[1,2,3,4,5,6,7,8,9,10], '@value>4 and @value<9'}` -> `{5,6,7,8}`

Example:
Finding all **audio clips** containing the word `effect`:

`where{t:AudioClip, @path:effect}`

## Additional resources

* [Search provider references](search-by-provider.md)
* [Create a custom search provider](create-search-provider.md)
* [Visual query builder reference](search-visual-query.md)
* [Launch and use the Search window](search-window-reference.md)