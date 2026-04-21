# SliderInt

The SliderInt lets users select an integer value from a range of values. You can use a SliderInt to adjust a value, such as the number of lives in a game. To set the value, you drag a thumb along a track, or click the track itself. A slider can also include a **text input field**, where you can modify the value directly.

**Note**: To align an element with other fields in an **Inspector** window, simply apply the `.unity-base-field__aligned` USS class to it. For more information, refer to [`BaseField`](../ScriptReference/UIElements.BaseField_1.md).

## Create a SliderInt

You can create a SliderInt with UI Builder, UXML, or C#. The following C# example creates a SliderInt with a default value, low value, and high value, and sets the direction of the slider:

```
var sliderInt = new SliderInt();
sliderInt.value = 42;
sliderInt.lowValue = 0;
sliderInt.highValue = 100;
sliderInt.direction = SliderDirection.Horizontal;
```

## Set page size

The SliderInt control has a page-size property that controls how much the thumb moves, and how much the value changes, when you click the track to either side of the thumb.

If the page size value is `0`, clicking the track moves the thumb to the cursor or pointer position, and updates the value accordingly.

If the page size value is greater than `0`, clicking the track moves the thumb toward the cursor or pointer position by a specific amount. SliderInt controls treat page-size as a percentage of the total range.

For example, if the total range is `200`, and the page size is `20`, then the page size amount is `20%` of `200`, or `40`. Each time you click the track, the thumb moves, and the value increases or decreases by `40`.

## Keyboard Shortcuts

You can use keyboard shortcuts to navigate and adjust the SliderInt control. The following table lists the keyboard shortcuts:

| **Shortcut** | **Action** |
| --- | --- |
| **Arrow Keys** | Move by `1/100` of the entire range. |
| **Arrow Keys + Shift** | Move by `1/10` of the entire range. |
| **Home** | Jump to the start of the slider. |
| **End** | Jump to the end of the slider. |
| **Page Up** | Move up by `(1/100 of the range) × page size`. |
| **Page Down** | Move down by `(1/100 of the range) × page size`. |

## Example

The following UXML example creates a SliderInt:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <SliderInt
        label="UXML Slider"
        name="the-uxml-slider"
        low-value="0"
        high-value="100"
        page-size="1"
        value="42" />
</UXML>
```

The following C# example illustrates some of the customizable functionalities of the SliderInt:

```
/// <sample>
// Get a reference to the slider from UXML and assign it its value.
var uxmlSlider = container.Q<SliderInt>("the-uxml-slider");
uxmlSlider.value = 42;

// Create a new slider, disable it, and give it a style class.
var csharpSlider = new SliderInt("C# Slider", 0, 100);
csharpSlider.SetEnabled(false);
csharpSlider.AddToClassList("some-styled-slider");
csharpSlider.value = uxmlSlider.value;
container.Add(csharpSlider);

// Mirror value of uxml slider into the C# field.
uxmlSlider.RegisterCallback<ChangeEvent<int>>((evt) =>
{
    csharpSlider.value = evt.newValue;
});
/// </sample>
```

To try this example live in Unity, go to **Window** > **UI Toolkit** > **Samples**.

## C# base class and namespace

**C# class**: [`SliderInt`](../ScriptReference/UIElements.SliderInt.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`BaseSlider_1`](../ScriptReference/UIElements.BaseSlider_1.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `page-size` | `float` | The value to add or remove to the SliderInt.value when it is clicked.  This is casted to int. |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `binding-path` | `string` | Path of the target property to be bound. |
| `direction` | [`UIElements.SliderDirection`](../ScriptReference/UIElements.SliderDirection.md) | This is the actual property to contain the direction of the slider. |
| `fill` | `boolean` | Enables fill to set the color and shape of a slider. |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `high-value` | `int` | This is the maximum value that the slider encodes. |
| `inverted` | `boolean` | This indicates whether or not this slider is inverted. For an inverted horizontal slider, high value is located to the left, low value is located to the right For an inverted vertical slider, high value is located to the bottom, low value is located to the top. |
| `label` | `string` | The string representing the label that will appear beside the field. If the string is empty, the label element is removed from the hierarchy. If the string is not empty, the label element is added to the hierarchy. |
| `low-value` | `int` | This is the minimum value that the slider encodes. |
| `page-size` | `float` | Represents the value that should be applied to the calculated scroll offset while scrolling the slider, such as when clicking within the track or clicking the slider arrows. |
| `show-input-field` | `boolean` | The visibility of the optional field inside the slider control.  Set this property to true to display a numerical text field that provides another way to edit the slider value. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |
| `value` | `int` | The actual value of the slider. |

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
| `ussClassName` | `.unity-slider-int` | USS class name of elements of this type. |
| `labelUssClassName` | `.unity-slider-int__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-slider-int__input` | USS class name of input elements in elements of this type. |
| `ussClassName` | `.unity-base-slider` | USS class name of elements of this type. |
| `labelUssClassName` | `.unity-base-slider__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-base-slider__input` | USS class name of input elements in elements of this type. |
| `horizontalVariantUssClassName` | `.unity-base-slider--horizontal` | USS class name of elements of this type, when they are displayed horizontally. |
| `verticalVariantUssClassName` | `.unity-base-slider--vertical` | USS class name of elements of this type, when they are displayed vertically. |
| `dragContainerUssClassName` | `.unity-base-slider__drag-container` | USS class name of container elements in elements of this type. |
| `trackerUssClassName` | `.unity-base-slider__tracker` | USS class name of tracker elements in elements of this type. |
| `draggerUssClassName` | `.unity-base-slider__dragger` | USS class name of dragger elements in elements of this type. |
| `draggerBorderUssClassName` | `.unity-base-slider__dragger-border` | USS class name of the dragger border element in elements of this type. |
| `fillUssClassName` | `.unity-base-slider__fill` | USS class name of fill element in elements of this type. |
| `movableUssClassName` | `.unity-base-slider--movable` | USS class name on the dragger that indicates it is currently controlled by `NavigationMoveEvent`. When the slider detects move events aligned with the slider’s direction, it adjusts the slider’s value. If it detects a navigation submit event, it removes the style, causing all navigation events to revert to their default behavior. A second navigation submit event re-applies the style to the dragger and restores the previous customized behavior. |
| `ussClassName` | `.unity-base-field` | USS class name of elements of this type. |
| `labelUssClassName` | `.unity-base-field__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-base-field__input` | USS class name of input elements in elements of this type. |
| `noLabelVariantUssClassName` | `.unity-base-field--no-label` | USS class name of elements of this type, when there is no label. |
| `labelDraggerVariantUssClassName` | `.unity-base-field__label--with-dragger` | USS class name of labels in elements of this type, when there is a dragger attached on them. |
| `mixedValueLabelUssClassName` | `.unity-base-field__label--mixed-value` | USS class name of elements that show mixed values |
| `alignedFieldUssClassName` | `.unity-base-field__aligned` | USS class name of elements that are aligned in a inspector element |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## Additional resources

* [Slider](UIE-uxml-element-Slider.md)
* [MinMaxSlider](UIE-uxml-element-MinMaxSlider.md)