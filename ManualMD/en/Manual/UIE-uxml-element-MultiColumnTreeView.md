# MultiColumnTreeView

A `MultiColumnTreeView` is a UI control commonly used to display tabular or grid-like data with multiple columns. Use a `MultiColumnTreeView` to present data in a structured format, where each row represents an item or entry, and each column represents a specific attribute or property of that item.

![MultiColumnTreeView example](../uploads/Main/planets-multicolumntreeview.png)

MultiColumnTreeView example

## Create a MultiColumnTreeView

You can create a MultiColumnTreeView with UXML and C#. The following C# example creates a MultiColumnTreeView with two columns:

```
var multiColumnTreeView = new MultiColumnTreeView();
multiColumnTreeView.columns.Add(new Column { title = "Name", width = 20 });
multiColumnTreeView.columns.Add(new Column { title = "Power", width = 80 });
```

## Refresh the collection view

To refresh the collection view, you can call either [`RefreshItems`](../ScriptReference/UIElements.BaseVerticalCollectionView.RefreshItems.md)/[`RefreshItem`](../ScriptReference/UIElements.BaseVerticalCollectionView.RefreshItem.md) or [`Rebuild`](../ScriptReference/UIElements.BaseVerticalCollectionView.Rebuild.md).

In general, call the `RefreshItems` or `RefreshItem` method—these methods only refresh the collection view and are less expensive. If you call `Rebuild`, the collection view is completely rebuilt, which can be expensive.

However, you must call `Rebuild` in the following cases:

* You change the type of the [data source](../ScriptReference/UIElements.BaseVerticalCollectionView-itemsSource.md), such as changing from a `List<float>` to a `List<Vector3>`.
* You make changes to [`Column.makeCell`](../ScriptReference/UIElements.Column-makeCell.md) or [`Column.destroyCell`](../ScriptReference/UIElements.Column-destroyCell.md).

## Bind the MultiColumnTreeView to data

You can bind the `MultiColumnTreeView` to a data source, and automatically populate the rows and columns based on the provided data. Refer to [Create list and tree views](UIE-ListView-TreeView.md) for an example.

## Sort the MultiColumnTreeView

You can enable sorting for MultiColumnTreeView using a default algorithm, or implementing a custom algorithm if the default sorting operation isn’t fast enough to meet your needs.

Once the sorting is enabled, to sort by a single column, click the column headers. To sort multiple columns, hold down the **Ctrl** key (macOS: **Cmd**) and click. To disable sorting, hold down the **Shift** key and click.

## Default sorting

To enable sorting using the default algorithm, set the [`sortingMode`](../ScriptReference/UIElements.MultiColumnTreeView-sortingMode.md) attribute to [`ColumnSortingMode.Default`](../ScriptReference/UIElements.ColumnSortingMode.Default.md), and implement the [`Column.comparison`](../ScriptReference/UIElements.Column-comparison.md) to compare two items by their index in the source.

For an example, refer to [a MultiColumnTreeView default sorting example](../ScriptReference/UIElements.Column-comparison.md).

## Custom sorting

To implement custom sorting, set the [`sortingMode`](../ScriptReference/UIElements.MultiColumnTreeView-sortingMode.md) attribute to [`ColumnSortingMode.Custom`](../ScriptReference/UIElements.ColumnSortingMode.Custom.md), and implement the [`columnSortingChanged`](../ScriptReference/UIElements.MultiColumnTreeView-columnSortingChanged.md) callback. In the callback, use [`sortColumns`](../ScriptReference/UIElements.MultiColumnTreeView-sortedColumns.md) to reorder your list accordingly and trigger a [`RefreshItems()`](../ScriptReference/UIElements.BaseVerticalCollectionView.RefreshItems.md) once the items’ source has been updated.

The `sortColumns` attribute is the current state of sorted columns: it lists the columns in the order in which they’re sorted. If you sort by column 1, then column 2, then column 3, the `sortedColumns` list is [1, 2, 3]. If you sort by column 3, then column 2, then column 1, the `sortedColumns` list is [3, 2, 1].

## Implement drag-and-drop operations for MultiColumnTreeView

Drag-and-drop is a common feature in UI design. To implement the drag-and-drop operations, override the following methods:

* To enable dragging items, override [`canStartDrag`](../ScriptReference/UIElements.BaseVerticalCollectionView-canStartDrag.md).
* To set which items are dragged, override [`setupDragAndDrop`](../ScriptReference/UIElements.BaseVerticalCollectionView-setupDragAndDrop.md).
* To get item status changes, override [`dragAndDropUpdate`](../ScriptReference/UIElements.BaseVerticalCollectionView-dragAndDropUpdate.md). You can perform certain actions based on the drag position or other conditions.
* To control the drag-and-drop behavior, override [`handleDrop`](../ScriptReference/UIElements.BaseVerticalCollectionView-handleDrop.md).

During the drag-and-drop operation, you can enable the reordering of items by dragging. To enable, set the `reorderable` attribute to `true` in UI Builder, UXML, and C#.

Refer to [Create a drag-and-drop list and tree views between windows](UIE-create-drag-and-drop-list-treeview-between-windows.md) for an example.

## MultiColumnTreeView FAQs

The following are some frequently asked questions about the MultiColumnTreeView control:

**Can I get the list of rows that are visible in the view?**

There are no dedicated APIs for this. You can use [UQuery](UIE-UQuery.md) to retrieve the elements of interest.

**Is it mandatory for any of the overridden functions of a view controller to call `base.Method`?**

Call this method only if you want to extend its default behavior.

**I’ve added a Toggle to my row. Why doesn’t the selection jump to that row when the user selects it?**

By default, the row is selected only if the mouse down event is not consumed by the row’s contents. In this case, your Toggle stops the event propagation. To fix this, register a [`PointerDownEvent`](../ScriptReference/UIElements.PointerDownEvent.md) callback with [`TrickleDown`](../ScriptReference/UIElements.TrickleDown.md) on the Toggle to call [`SetSelection`](../ScriptReference/UIElements.BaseVerticalCollectionView.SetSelection.md).

**How do I receive a callback when a user changes their selection in the view?**

It is recommended to use the [`selectedIndicesChanged`](../ScriptReference/UIElements.BaseVerticalCollectionView-selectedIndicesChanged.md) callback to retrieve data by index when needed. While you can also use [`selectionChanged`](../ScriptReference/UIElements.BaseVerticalCollectionView-selectionChanged.md), be aware that it returns a list of objects, which can cause boxing allocations when used with value types.

**Can I get the indices of the rows that are visible on the screen?**

There are no dedicated APIs for this. You can use `column.bindCell` and `column.unbindCell` to track the indices on one of the columns.

**How do I convert from ID to index and vice versa?**

Use [`BaseTreeViewController.GetIndexForId`](../ScriptReference/UIElements.BaseTreeViewController.GetIndexForId.md) and [`BaseTreeViewController.GetIdForIndex`](../ScriptReference/UIElements.BaseTreeViewController.GetIdForIndex.md).

**Can I have a horizontal MultiColumnTreeView?**

The MultiColumnTreeView control doesn’t support horizontal layout and virtualization. It’s recommended to use a ScrollView with `flex-direction: row` to layout elements horizontally.

## Examples

* [Create list and tree views](UIE-ListView-TreeView.md): Use ListView, TreeView, MultiColumnListView, and MultiColumnTreeView to create list and tree views.

## C# base class and namespace

**C# class**: [`MultiColumnTreeView`](../ScriptReference/UIElements.MultiColumnTreeView.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`BaseTreeView`](../ScriptReference/UIElements.BaseTreeView.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `columns` | [`UIElements.Columns.UxmlSerializedData`](../ScriptReference/UIElements.Columns.UxmlSerializedData.md) | The collection of columns for the multi-column header. |
| `sort-column-descriptions` | [`UIElements.SortColumnDescriptions.UxmlSerializedData`](../ScriptReference/UIElements.SortColumnDescriptions.UxmlSerializedData.md) | The collection of sorted columns by default. |
| `sorting-enabled` | `boolean` | Whether or not sorting is enabled in the multi-column header. This is deprecated. Use `sortingMode` instead. |
| `sorting-mode` | [`UIElements.ColumnSortingMode`](../ScriptReference/UIElements.ColumnSortingMode.md) | Indicates how to sort columns. To enable sorting, set it to `ColumnSortingMode.Default` or `ColumnSortingMode.Custom`. The `Default` mode uses the sorting algorithm provided by `MultiColumnController`, acting on indices. You can also implement your own sorting with the `Custom` mode, by responding to the `columnSortingChanged` event.  **Note**: If there is at least one sorted column, reordering is temporarily disabled. |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `auto-expand` | `boolean` | When true, items are automatically expanded when added to the TreeView. |
| `binding-path` | `string` | Path of the target property to be bound. |
| `fixed-item-height` | `float` | The height of a single item in the list, in pixels.  This property must be set when using the `virtualizationMethod` is set to `FixedHeight`, for the collection view to function. If set when `virtualizationMethod` is `DynamicHeight`, it serves as the default height to help calculate the number of items necessary and the scrollable area, before items are laid out. It should be set to the minimum expected height of an item. |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `horizontal-scrolling` | `boolean` | This property controls whether the collection view shows a horizontal scroll bar when its content does not fit in the visible area. |
| `reorderable` | `boolean` | Gets or sets a value that indicates whether the user can drag list items to reorder them.  The default value is `false` which allows the user to drag items to and from other views when you implement `canStartDrag`, `setupDragAndDrop`, `dragAndDropUpdate`, and `handleDrop`. Set this value to `true` to allow the user to reorder items in the list. |
| `selection-type` | [`UIElements.SelectionType`](../ScriptReference/UIElements.SelectionType.md) | Controls the selection type.  The default value is `SelectionType.Single`. When you set the collection view to disable selections, any current selection is cleared. |
| `show-alternating-row-backgrounds` | [`UIElements.AlternatingRowBackground`](../ScriptReference/UIElements.AlternatingRowBackground.md) | This property controls whether the background colors of collection view rows alternate. Takes a value from the `AlternatingRowBackground` enum. |
| `show-border` | `boolean` | Enable this property to display a border around the collection view.  If set to true, a border appears around the ScrollView that the collection view uses internally. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |
| `virtualization-method` | [`UIElements.CollectionVirtualizationMethod`](../ScriptReference/UIElements.CollectionVirtualizationMethod.md) | The virtualization method to use for this collection when a scroll bar is visible. Takes a value from the `CollectionVirtualizationMethod` enum.  The default value is `FixedHeight`. When using fixed height, specify the `fixedItemHeight` property. Fixed height is more performant but offers less flexibility on content. When using `DynamicHeight`, the collection will wait for the actual height to be computed. Dynamic height is more flexible but less performant. |

This element also inherits the following attributes from [`VisualElement`](UIE-uxml-element-VisualElement.md):

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `content-container` | `string` | Logical container where child elements are added. If a child is added to this element, the child is added to this element’s content container instead.  When iterating over the `VisualElement.Children` of an element, the element’s content container hierarchy is used instead of the element itself. This can lead to unexpected results, such as elements being ignored by the navigation events if they are not directly in the content container’s hierarchy. Refer to `IFocusRing` for more information.    If the content container is the same as the element itself, child elements are added directly to the element. This is true for most elements but can be overridden by more complex types.   The `ScrollView`, for example, has a content container that is different from itself. In that case, child elements added to the scroll view are added to its content container element instead. While the physical parent (`VisualElement.Hierarchy.parent`) of the child elements is the scroll view’s content container element, their logical parent (`VisualElement.parent`) still refers to the scroll view itself. Since some of the scroll view’s focusable children are not part of its logical hierarchy, like its `Scroller` elements, these focusable children are not considered by default when using sequential navigation. Refer to [How can I change what element is focused next](UIE-faq-event-and-input-system.md) for an example of a workaround solution if the default navigation rules don’t correspond to your needs. |
| `data-source` | `Object` | Assigns a data source to this VisualElement which overrides any inherited data source. This data source is inherited by all children. |
| `data-source-path` | `string` | Path from the data source to the value. |
| `data-source-type` | `System.Type` | The possible type of data source assignable to this VisualElement.    This information is only used by the UI Builder as a hint to provide some completion to the data source path field when the effective data source cannot be specified at design time. |
| `enabled` | `boolean` | Returns true if the `VisualElement` is enabled locally.  This flag isn’t changed if the VisualElement is disabled implicitly by one of its parents. To verify this, use `enabledInHierarchy`. |
| `language-direction` | [`UIElements.LanguageDirection`](../ScriptReference/UIElements.LanguageDirection.md) | Indicates the directionality of the element’s text. The value will propagate to the element’s children.  Setting `languageDirection` to `RTL` can only get the basic RTL support like text reversal. To get more comprehensive RTL support, such as line breaking, word wrapping, or text shaping, you must enable [Advance Text Generator](UIE-advanced-text-generator.md). |
| `name` | `string` | The name of this VisualElement.  Use this property to write USS selectors that target a specific element. The standard practice is to give an element a unique name. |
| `picking-mode` | [`UIElements.PickingMode`](../ScriptReference/UIElements.PickingMode.md) | Determines if this element can be the target of pointer events or picked by `IPanel.Pick` queries.  Elements can not be picked if:  - They are invisible- Their `style.display` is set to `DisplayStyle.None`   Elements with a picking mode of `PickingMode.Ignore` never receive the hover pseudo-state. |
| `style` | `string` | Sets the style values on a `VisualElement`.  The returned style data, computed from USS files or inline styles written to this object in C#, doesn’t represent the fully resolved styles, such as the final height and width of a VisualElement. To access these fully resolved styles, use `resolvedStyle`.      For information about how to use this property and all the supported USS properties, refer to the [Apply styles in C# scripts](UIE-apply-styles-with-csharp.md) and [USS properties reference](UIE-USS-Properties-Reference.md) manual pages. |
| `tooltip` | `string` | Text to display inside an information box after the user hovers the element for a small amount of time. This is only supported in the Editor UI. |
| `usage-hints` | [`UIElements.UsageHints`](../ScriptReference/UIElements.UsageHints.md) | A combination of hint values that specify high-level intended usage patterns for the `VisualElement`. This property can only be set when the `VisualElement` is not yet part of a `Panel`. Once part of a `Panel`, this property becomes effectively read-only, and attempts to change it will throw an exception. The specification of proper `UsageHints` drives the system to make better decisions on how to process or accelerate certain operations based on the anticipated usage pattern. Note that those hints do not affect behavioral or visual results, but only affect the overall performance of the panel and the elements within. It’s advised to always consider specifying the proper `UsageHints`, but keep in mind that some `UsageHints` might be internally ignored under certain conditions (e.g. due to hardware limitations on the target platform). |
| `view-data-key` | `string` | Used for view data persistence, such as tree expanded states, scroll position, or zoom level.  This key is used to save and load the view data from the view data store. If you don’t set this key, the persistence is disabled for the associated `VisualElement`. For more information, refer to [View data persistence in the Unity Manual](UIE-ViewData.md). |

## USS classes

The following table lists all the C# public property names and their related USS selector.

| **C# property** | **USS selector** | **Description** |
| --- | --- | --- |
| `ussClassName` | `.unity-tree-view` | The USS class name for TreeView elements.  Unity adds this USS class to every instance of the TreeView element. Any styling applied to this class affects every TreeView located beside, or below the stylesheet in the visual tree. |
| `itemUssClassName` | `.unity-tree-view__item` | The USS class name for TreeView item elements.  Unity adds this USS class to every item element of the TreeView. Any styling applied to this class affects every item located beside, or below the stylesheet in the visual tree. |
| `itemToggleUssClassName` | `.unity-tree-view__item-toggle` | The USS class name for TreeView item toggle elements.  Unity adds this USS class to every item toggle element of the TreeView. Any styling applied to this class affects every item located beside, or below the stylesheet in the visual tree. |
| `itemIndentsContainerUssClassName` | `.unity-tree-view__item-indents` | The USS class name for TreeView indent container elements.  Unity adds this USS class to every indent container element of the TreeView. Any styling applied to this class affects every item located beside, or below the stylesheet in the visual tree. |
| `itemIndentUssClassName` | `.unity-tree-view__item-indent` | The USS class name for TreeView indent element.  Unity adds this USS class to every indent element of the TreeView. Any styling applied to this class affects every item located beside, or below the stylesheet in the visual tree. |
| `itemContentContainerUssClassName` | `.unity-tree-view__item-content` | The USS class name for TreeView item container elements.  Unity adds this USS class to every item container element of the TreeView. Any styling applied to this class affects every item located beside, or below the stylesheet in the visual tree. |
| `ussClassName` | `.unity-collection-view` | The USS class name for BaseVerticalCollectionView elements.  Unity adds this USS class to every instance of the BaseVerticalCollectionView element. Any styling applied to this class affects every BaseVerticalCollectionView located beside, or below the stylesheet in the visual tree. |
| `borderUssClassName` | `.unity-collection-view--with-border` | The USS class name for BaseVerticalCollectionView elements with a border.  Unity adds this USS class to an instance of the BaseVerticalCollectionView element if the instance’s `BaseVerticalCollectionView.showBorder` property is set to true. Any styling applied to this class affects every such BaseVerticalCollectionView located beside, or below the stylesheet in the visual tree. |
| `itemUssClassName` | `.unity-collection-view__item` | The USS class name of item elements in BaseVerticalCollectionView elements.  Unity adds this USS class to every item element the BaseVerticalCollectionView contains. Any styling applied to this class affects every item element located beside, or below the stylesheet in the visual tree. |
| `dragHoverBarUssClassName` | `.unity-collection-view__drag-hover-bar` | The USS class name of the drag hover bar.  Unity adds this USS class to the bar that appears when the user drags an item in the list. Any styling applied to this class affects every BaseVerticalCollectionView located beside, or below the stylesheet in the visual tree. |
| `dragHoverMarkerUssClassName` | `.unity-collection-view__drag-hover-marker` | The USS class name of the drag hover circular marker used to indicate depth.  Unity adds this USS class to the bar that appears when the user drags an item in the list. Any styling applied to this class affects every BaseVerticalCollectionView located beside, or below the stylesheet in the visual tree. |
| `itemDragHoverUssClassName` | `.unity-collection-view__item--drag-hover` | The USS class name applied to an item element on drag hover.  Unity adds this USS class to the item element when it’s being dragged. Any styling applied to this class affects every BaseVerticalCollectionView item located beside, or below the stylesheet in the visual tree. |
| `itemSelectedVariantUssClassName` | `.unity-collection-view__item--selected` | The USS class name of selected item elements in the BaseVerticalCollectionView.  Unity adds this USS class to every selected element in the BaseVerticalCollectionView. The `BaseVerticalCollectionView.selectionType` property decides if zero, one, or more elements can be selected. Any styling applied to this class affects every BaseVerticalCollectionView item located beside, or below the stylesheet in the visual tree. |
| `itemAlternativeBackgroundUssClassName` | `.unity-collection-view__item--alternative-background` | The USS class name for odd rows in the BaseVerticalCollectionView.  Unity adds this USS class to every odd-numbered item in the BaseVerticalCollectionView when the `BaseVerticalCollectionView.showAlternatingRowBackgrounds` property is set to `ContentOnly` or `All`. When the `showAlternatingRowBackgrounds` property is set to either of those values, odd-numbered items are displayed with a different background color than even-numbered items. This USS class is used to differentiate odd-numbered items from even-numbered items. When the `showAlternatingRowBackgrounds` property is set to `None`, the USS class is not added, and any styling or behavior that relies on it’s invalidated. |
| `listScrollViewUssClassName` | `.unity-collection-view__scroll-view` | The USS class name of the scroll view in the BaseVerticalCollectionView.  Unity adds this USS class to the BaseVerticalCollectionView’s scroll view. Any styling applied to this class affects every BaseVerticalCollectionView scroll view located beside, or below the stylesheet in the visual tree. |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## Additional resources

* [MultiColumnListView](UIE-uxml-element-MultiColumnListView.md)
* [TreeView](UIE-uxml-element-TreeView.md)
* [ScrollView](UIE-uxml-element-ScrollView.md)
* [ListView](UIE-uxml-element-ListView.md)