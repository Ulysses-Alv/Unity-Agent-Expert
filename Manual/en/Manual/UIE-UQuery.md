# Find visual elements with UQuery

You can use [UQuery](../ScriptReference/UIElements.UQuery.md) to find elements from a [visual tree](UIE-VisualTree.md). UQuery was inspired by JQuery and Linq, and is designed to limit dynamic memory allocation. This allows for optimal performance on mobile platforms.

## Query methods

You can use UQuery through the following extension methods:

* [`Q`](../ScriptReference/UIElements.UQueryExtensions.Q.md)
* [`Query`](../ScriptReference/UIElements.UQueryExtensions.Query.md)

Internally, the `Q` and `Query` methods use [`UQueryBuilder`](../ScriptReference/UIElements.UQueryBuilder_1.md) to construct a query. These extension methods reduce the verbosity of creating a `UQueryBuilder`.

To use UQuery to find elements, you must [load](UIE-manage-asset-reference.md) and [instantiate](UIE-LoadingUXMLcsharp.md) the UXML first, and then use `Query` or `Q` to construct selection rules on a root **visual element**.

`Query` returns a list of elements that match the selection rules. You can filter the return results of `Query` with the public methods of [`UQueryBuilder`](../ScriptReference/UIElements.UQueryBuilder_1.md), such as [First](../ScriptReference/UIElements.UQueryBuilder_1.First.md), [Last](../ScriptReference/UIElements.UQueryBuilder_1.Last.md), [AtIndex](../ScriptReference/UIElements.UQueryBuilder_1.AtIndex.md), [Children](../ScriptReference/UIElements.UQueryBuilder_1.Children.md), and [Where](../ScriptReference/UIElements.UQueryBuilder_1.Where.md).

`Q` is the shorthand for `Query<T>.First()`. It returns the first element that matches the selection rules.

## Query elements

You can query elements by their [name](#name), their [USS class](#class), or their [element type (C# type)](#type). You can also query with a [predicate](#predicate) or make [complex hierarchical queries](#complex).

The following sections use this example UXML to demonstrate how to find elements:

```
<UXML xmlns="UnityEngine.UIElements">
    <VisualElement name="container1">
      <Button name="OK" text="OK" />
      <Button name="Cancel" text="Cancel" />
    </VisualElement>
     <VisualElement name="container2">
      <Button name="OK" class="yellow" text="OK" />
      <Button name="Cancel" text="Cancel" />
    </VisualElement>
    <VisualElement name="container3">
      <Button name="OK" class="yellow" text="OK" />
      <Button name="Cancel" class="yellow" text="Cancel" />
    </VisualElement>
</UXML>
```

### Query by name

To find elements by their [name](UIE-USS-Selectors-name.md), use `Query(name: "element-name")` or `Q(name: "element-name")`. You can omit the `name` as it’s the first argument. For example:

The following example finds a list of elements named `OK`:

```
List<VisualElement> result = root.Query("OK").ToList();
```

The following example uses `Query` to find the first element named `OK`:

```
VisualElement result = root.Query("OK").First(); //or VisualElement result = root.Q("OK");
```

The following example uses `Q` to find the first element named `OK`:

```
VisualElement result = root.Q("OK");
```

The following example finds the second element named `OK`:

```
VisualElement result3 = root.Query("OK").AtIndex(1);
```

The following example finds the last element named `OK`:

```
VisualElement result4 = root.Query("OK").Last();
```

### Query by USS class

To find elements by a [USS class](UIE-USS-Selectors-class.md), use `Query(className: "class-name")` or `Q(className: "class-name")`.

The following example finds all the elements that have the class `yellow` and assigns them to a list:

```
List<VisualElement> result = root.Query(className: "yellow").ToList();
```

The following example finds the first element that has the class `yellow`:

```
VisualElement result = root.Q(className: "yellow");
```

### Query by element type

To find elements by their [element type](UIE-USS-Selectors-type.md)(C# type), use `Query<Type>` or `Q<Type>`.

The following example finds the first button and adds a tooltip for it:

```
VisualElement result = root.Q<Button>();
result.tooltip = "This is a tooltip!";
```

The following example finds the third button:

```
VisualElement result = root.Query<Button>().AtIndex(2);
```

**Note**: You can only query by the actual type of the element, not base classes.

### Query with a predicate

Other than to query elements by name, class, and type, you can also use the `Where` method to select all elements that satisfy a predicate. The predicate must be a function callback that takes a single `VisualElement` argument.

The following example finds all the elements with the `yellow`USS class that have no tooltips:

```
List<VisualElement> result = root.Query(className: "yellow").Where(elem => elem.tooltip == "").ToList();
```

### Complex hierarchical queries

You can combine name, class, and type to make complex hierarchical queries.

The following example finds the first button named `OK` that has a class of `yellow`:

```
VisualElement result = root.Query<Button>(className: "yellow", name: "OK").First();
```

The following example finds the child cancel button of the `container2`:

```
VisualElement result = root.Query<VisualElement>("container2").Children<Button>("Cancel").First();
```

**Note**: UQuery includes the queried element itself in the search results if it matches the criteria. For example:

```
var list = myElement.Query<VisualElement>("nameOrClass").ToList();
```

It might return `myElement` itself if it fits the selector. This behavior differs from `Element.querySelectorAll()` on the web, which explicitly returns only descendant nodes.

### Operate on results

You can use the [ForEach](../ScriptReference/UIElements.UQueryBuilder_1.ForEach.md) method to operate directly on the query results.

The following example adds a tooltip for any elements that have no tooltips:

```
root.Query().Where(elem => elem.tooltip == "").ForEach(elem => elem.tooltip="This is a tooltip!");
```

## Best practices

Consider the following when you use UQuery:

* UQuery traverses through the hierarchy to find elements by name, class or type. Cache results from UQuery at initialization.
* To find ancestor elements or check parent-child relationships, manually traverse up the `.parent` chain until you find the target element or reach null.
* If you need to retrieve multiple elements, use the `QueryState` struct (returned by the `element.Query()` method) and enumerate it to avoid creating lists. You can also construct a query once and execute it on different elements.
* UI Toolkit doesn’t destroy visual elements that are no longer needed, it uses C# garbage collector to collect them. Be mindful to not accidentally retain references to visual elements in a class that outlives the UIDocuments or Window where the elements came from.
* When you create or release lots of elements, enable [incremental garbage collection](performance-incremental-garbage-collection.md) to avoid garbage collector spikes.
* Capture `VisualElement` variables inside closures. When you use event callbacks or delegates with visual elements, capture the specific elements you need in local variables before creating the closure. This prevents the closure from capturing a larger scope (like the entire `this` reference) which can lead to memory leaks. For example:

```
  // Good: Capture only what you need.
  var button = root.Q<Button>("myButton");
  button.clicked += () => Debug.Log($"Button {button.name} clicked");
  
  // Avoid: This captures the entire 'this' reference.
  button.clicked += () => Debug.Log($"Button {this.myButton.name} clicked");
```

## Additional resources

* [USS selectors](UIE-USS-Selectors.md)
* [Introduction to visual elements and visual tree](UIE-VisualTree.md)
* [Load UXML and USS from C# scripts](UIE-manage-asset-reference.md)
* [Instantiate UXML with C#](UIE-LoadingUXMLcsharp.md)