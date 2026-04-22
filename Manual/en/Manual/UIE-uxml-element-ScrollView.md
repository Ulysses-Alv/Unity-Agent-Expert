# ScrollView

ScrollView displays its content inside a scrollable area. When you add content to a **ScrollView**, the content is added to the content container (`#unity-content-container`) of the ScrollView.

## Create a ScrollView

You can create a ScrollView with UI Builder, UXML, or C#. The following C# example creates a ScrollView with both horizontal and vertical scroll capabilities, that contains a title Label, and a number of Toggle elements. It also adds a button that, when clicked, scrolls the list down by one page using the vertical scroller.

```
// Create a ScrollView with vertical and horizontal scrolling enabled
// and set its width and height.
var scrollView = new ScrollView(ScrollViewMode.VerticalAndHorizontal);
scrollView.style.width = 250;
scrollView.style.height = 400;

// Add a label and 100 toggle elements to the ScrollView.
scrollView.Add(new Label("List of checkboxes:"));

for (int i = 0; i < 100; ++i)
{
    var toggle = new UnityEngine.UIElements.Toggle()
    { text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua." };
    scrollView.Add(toggle);
}

// Access the vertical scroller.
var verticalScroller = scrollView.verticalScroller;

// Add a button to scroll down one page when clicked.
var scrollDownButton = new Button(() =>
{
    verticalScroller.ScrollPageDown(50f);
})
{ text = "Scroll Down One Page" };

uiDocument.rootVisualElement.Add(scrollView);
uiDocument.rootVisualElement.Add(scrollDownButton);
```

## Manage the scroll bar

You can specify the scroll mode of the scroll bar movement, whether the horizontal or vertical scroll bars are visible, and control the speed of the scroll bars.

### Scroll bar mode and visibility

To set the scroll mode of the scroll bar movement in UI Builder, in the **Inspector** window of the ScrollView, select one of the following options for **Mode**:

* Vertical (default)
* Horizontal
* Vertical and Horizontal

To set the visibility of the scroll bar in UI Builder, in the Inspector window of the ScrollView, select one of the following options for **Horizontal Scroller Visibility** or **Vertical Scroller Visibility**:

* Auto (default)
* Always Visible
* Hidden

### Control the scroll speed

To control the scroll speed, adjust the values of the following properties in UI Builder, UXML, or C# code. The higher the value, the faster the scrolling speed.

* `horizontal-page-size` controls the speed of horizontal scrolling when using a keyboard or the on-screen scrollbar buttons (arrows and handle). The speed is calculated by multiplying the page size by the specified value. For example, a value of `2` means that each scroll movement covers a distance equal to twice the width of the page.
* `vertical-page-size` controls the speed of vertical scrolling when using a keyboard or the on-screen scrollbar buttons (arrows and handle). The speed is calculated by multiplying the page size by the specified value. For example, a value of `2` means that each scroll movement covers a distance equal to twice the length of the page.
* `mouse-wheel-scroll-size` controls the speed of scrolling when using the mouse wheel. The speed is calculated by dividing the specified value by `18`, which corresponds to the default line height of `18px`. For example, if you set the value to `36`, each scroll movement covers a distance equivalent to two lines of content.

**Note**: You can also override the [USS built-in variable](UIE-uss-built-in-variable-reference.md) `–unity-metrics-single_line-height` to control the speed of scrolling when using the mouse wheel. The `mouse-wheel-scroll-size` attribute takes precedence over the `–unity-metrics-single_line-height` USS variable.

### Scroll to the end

To scroll to the end of a ScrollView, set the [scrollOffset](scriptRef:UIElements.ScrollView-scrollOffset) property to the maximum value of the vertical scroller. This scrolls the content to the end when you initialize the UI.

```
// Set the scroll offset to the bottom of the ScrollView.
// This scrolls to the bottom when the UI is initialized.
scrollView.scrollOffset = new Vector2(0, scrollView.verticalScroller.highValue);
```

## Wrap elements inside ScrollView

You can wrap **visual elements** inside a ScrollView so that the elements display in a row. If the row is full, the elements continue to display on the next line.

To wrap visual elements inside a ScrollView, set the ScrollView’s content container `flex-direction` to `row` and `flex-wrap` to `wrap`.

**In USS**:

```
.unity-scroll-view__content-container {
    flex-direction: row;
    flex-wrap: wrap;
}
```

**In C#**:

```
scrollview.contentContainer.style.flexDirection = FlexDirection.Row;
scrollview.contentContainer.style.flexWrap = Wrap.Wrap;
```

## Wrap text of elements inside ScrollView

To wrap the text of an element inside a ScrollView, for example, the text of a Label element, do the following:

1. Style the Label element with any of the following methods:
   * In UI Builder, in the Inspector window of the Label, select **Text** > **Wrap** > **normal**.
   * Apply style `white-space: normal;` to the Label element in USS, UXML, or C#.
2. Create a VisualElement as a container inside the ScrollView. The original container element inside the ScrollView does not have bounds set (is infinite in size), so text won’t properly wrap. This container VisualElement provides the finite size for text to wrap within.
3. Add the Label to the VisualElement container.

## Apply styles to the direct child of a ScrollView

When you add elements to a ScrollView, they’re added under a child element of the ScrollView that acts as a container. To add a style to this container, do this in USS:

```
ScrollView > .unity-scroll-view__content-container {
    margin: 10px;
}
```

## Examples

* [Wrap content inside a ScrollView](UIE-wrap-content-inside-scrollview.md): This example demonstrates how to use styles to wrap content inside a scroll view.

## C# class and namespace

**C# class**: [`ScrollView`](../ScriptReference/UIElements.ScrollView.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`VisualElement`](../ScriptReference/UIElements.VisualElement.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `elastic-animation-interval-ms` | `long` | The minimum amount of time, in milliseconds, between executions of elastic spring animation. |
| `elasticity` | `float` | The amount of elasticity to use when a user tries to scroll past the boundaries of the scroll view.  Elasticity is only used when `touchScrollBehavior` is set to Elastic. |
| `horizontal-page-size` | `float` | This property controls the speed of the horizontal scrolling when using a keyboard or the on-screen scrollbar buttons (arrows and handle), based on the size of the page.  When scrolling, the page size is applied to the offset for each scroll step, so the final offset will be the page size multiplied by the number of steps.\ |
| `horizontal-scroller-visibility` | [`UIElements.ScrollerVisibility`](../ScriptReference/UIElements.ScrollerVisibility.md) | Specifies whether the horizontal scroll bar is visible. |
| `mode` | [`UIElements.ScrollViewMode`](../ScriptReference/UIElements.ScrollViewMode.md) | Controls how the ScrollView allows the user to scroll the contents. `ScrollViewMode`  The default is `ScrollViewMode.Vertical`. Writing to this property modifies the class list of the ScrollView according to the specified value of `ScrollViewMode`. When the value changes, the class list matching the old value is removed and the class list matching the new value is added. |
| `mouse-wheel-scroll-size` | `float` | This property controls the scrolling speed only when using a mouse scroll wheel, based on the size of the page.  This property takes precedence over the `--unity-metrics-single_line-height` USS variable. If both the property and the variable are set, the property value is used.   The following example demonstrates how to use the `mouseWheelScrollSize` property to control the “speed” of a scroll when using the mouse wheel. Notice the difference in feel when selecting different values: |
| `nested-interaction-kind` | [`UIElements.ScrollView.NestedInteractionKind`](../ScriptReference/UIElements.ScrollView.NestedInteractionKind.md) | The behavior to use when scrolling reaches limits of a nested `ScrollView`. |
| `scroll-deceleration-rate` | `float` | Controls the rate at which the scrolling movement slows after a user scrolls using a touch interaction.  The deceleration rate is the speed reduction per second. A value of 0.5 halves the speed each second. A value of 0 stops the scrolling immediately. |
| `show-horizontal-scroller` | `boolean` | Obsolete. Use `ScrollView.horizontalScrollerVisibility` instead. |
| `show-vertical-scroller` | `boolean` | Obsolete. Use `ScrollView.verticalScrollerVisibility` instead. |
| `touch-scroll-type` | [`UIElements.ScrollView.TouchScrollBehavior`](../ScriptReference/UIElements.ScrollView.TouchScrollBehavior.md) | The behavior to use when a user tries to scroll past the boundaries of the ScrollView content using a touch interaction. |
| `vertical-page-size` | `float` | This property controls the speed of the vertical scrolling when using a keyboard or the on-screen scrollbar buttons (arrows and handle), based on the size of the page.  The speed is calculated by multiplying the page size by the specified value. For example, a value of `2` means that each scroll movement covers a distance equal to twice the width of the page.\ |
| `vertical-scroller-visibility` | [`UIElements.ScrollerVisibility`](../ScriptReference/UIElements.ScrollerVisibility.md) | Specifies whether the vertical scroll bar is visible. |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |

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
| `ussClassName` | `.unity-scroll-view` | USS class name of elements of this type. |
| `viewportUssClassName` | `.unity-scroll-view__content-viewport` | USS class name of **viewport** elements in elements of this type. |
| `horizontalVariantViewportUssClassName` | `.unity-scroll-view__content-viewport--horizontal` | USS class name that’s added when the Viewport is in horizontal mode. `ScrollViewMode.Horizontal` |
| `verticalVariantViewportUssClassName` | `.unity-scroll-view__content-viewport--vertical` | USS class name that’s added when the Viewport is in vertical mode. `ScrollViewMode.Vertical` |
| `verticalHorizontalVariantViewportUssClassName` | `.unity-scroll-view__content-viewport--vertical-horizontal` | USS class name that’s added when the Viewport is in both horizontal and vertical mode. `ScrollViewMode.VerticalAndHorizontal` |
| `contentAndVerticalScrollUssClassName` | `.unity-scroll-view__content-and-vertical-scroll-container` | USS class name of content elements in elements of this type. |
| `contentUssClassName` | `.unity-scroll-view__content-container` | USS class name of content elements in elements of this type. |
| `horizontalVariantContentUssClassName` | `.unity-scroll-view__content-container--horizontal` | USS class name that’s added when the ContentContainer is in horizontal mode. `ScrollViewMode.Horizontal` |
| `verticalVariantContentUssClassName` | `.unity-scroll-view__content-container--vertical` | USS class name that’s added when the ContentContainer is in vertical mode. `ScrollViewMode.Vertical` |
| `verticalHorizontalVariantContentUssClassName` | `.unity-scroll-view__content-container--vertical-horizontal` | USS class name that’s added when the ContentContainer is in both horizontal and vertical mode. `ScrollViewMode.VerticalAndHorizontal` |
| `hScrollerUssClassName` | `.unity-scroll-view__horizontal-scroller` | USS class name of horizontal scrollers in elements of this type. |
| `vScrollerUssClassName` | `.unity-scroll-view__vertical-scroller` | USS class name of vertical scrollers in elements of this type. |
| `horizontalVariantUssClassName` | `.unity-scroll-view--horizontal` | USS class name that’s added when the ScrollView is in horizontal mode. `ScrollViewMode.Horizontal` |
| `verticalVariantUssClassName` | `.unity-scroll-view--vertical` | USS class name that’s added when the ScrollView is in vertical mode. `ScrollViewMode.Vertical` |
| `verticalHorizontalVariantUssClassName` | `.unity-scroll-view--vertical-horizontal` | USS class name that’s added when the ScrollView is in both horizontal and vertical mode. `ScrollViewMode.VerticalAndHorizontal` |
| `scrollVariantUssClassName` | `.unity-scroll-view--scroll` | |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## Additional resources

* [ListView](UIE-uxml-element-ListView.md)
* [Scroller](UIE-uxml-element-Scroller.md)