# Toggle

A Toggle includes an image and a label. Just like other standard UI Toolkit controls, such as a [Button](UIE-uxml-element-Button.md), a Toggle has a Clickable manipulator attached to it that registers to [`MouseUpEvent`](../ScriptReference/UIElements.MouseUpEvent.md) and [`PointerUpEvent`](../ScriptReference/UIElements.PointerUpEvent.md). When one of those events triggers the manipulator, the Toggle’s value changes from `true` to `false` or from `false` to `true`. You can read or set the current values of a Toggle. You can also bind a Toggle to a `Boolean` variable.

You can display or hide elements based on other conditions, such as a selection from a drop-down, or an enabled option. An example of a conditional UI is the Unity **Camera** **Inspector**. It displays additional options when you activate **Physical Camera** mode.

**Note**: To align an element with other fields in an Inspector window, simply apply the `.unity-base-field__aligned` USS class to it. For more information, refer to [`BaseField`](../ScriptReference/UIElements.BaseField_1.md).

## Create a Toggle

You can create a Toggle with UI Builder, UXML, or C#.

The following C# example creates a [Toggle](../ScriptReference/UIElements.Toggle.md) with a specific label:

```
Toggle myToggle = new Toggle("Click me");
```

## Style a Toggle

By default, a Toggle control appears as a checkbox. You can apply styles to make it look like a typical toggle switch.

To style the checkbox of a toggle, use the `.unity-toggle__input` and `.unity-toggle__checkmark` style properties. For example, if you have applied a class named `className` to the Toggle. The following USS changes the background images of the checkbox when you select or clear the toggle:

```
/* Set the background image when the checkbox is selected. */
.className:checked > .unity-toggle__input > .unity-toggle__checkmark {
    background-image: url("path/to/image-file");
}

/* Set the background image When the checkbox is not selected. */
.className > .unity-toggle__input > .unity-toggle__checkmark {
    background-image: url("path/to/image-file"); 
}
```

## Examples

The following UXML example creates a Toggle:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <Toggle label="UXML Field" name="the-uxml-field" />
</UXML>
```

The following C# example illustrates some of the customizable functionalities of the Toggle:

```
/// <sample>
// Get a reference to the field from UXML and assign a value to it.
var uxmlField = container.Q<Toggle>("the-uxml-field");
uxmlField.value = true;

// Create a new field, disable it, and give it a style class.
var csharpField = new Toggle("C# Field");
csharpField.value = false;
csharpField.SetEnabled(false);
csharpField.AddToClassList("some-styled-field");
csharpField.value = uxmlField.value;
container.Add(csharpField);

// Mirror the value of the UXML field into the C# field.
uxmlField.RegisterCallback<ChangeEvent<bool>>((evt) =>
{
    csharpField.value = evt.newValue;
});
/// </sample>
```

To try this example live in Unity, go to **Window** > **UI Toolkit** > **Samples**.

For more examples, refer to the following:

* [Create a conditional UI](UIE-create-a-conditional-ui.md).

## C# base class and namespace

**C# class**: [`Toggle`](../ScriptReference/UIElements.Toggle.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`BaseBoolField`](../ScriptReference/UIElements.BaseBoolField.md)

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `binding-path` | `string` | Path of the target property to be bound. |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `label` | `string` | The string representing the label that will appear beside the field. If the string is empty, the label element is removed from the hierarchy. If the string is not empty, the label element is added to the hierarchy. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |
| `text` | `string` | Optional text that appears after the BaseBoolField.  Unity creates a `Label` automatically if one does not exist. |
| `toggle-on-label-click` | `boolean` | Whether to activate the toggle when the user clicks the label. |
| `value` | `boolean` | The value associated with the field. |

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
| `ussClassName` | `.unity-toggle` | USS class name for Toggle elements.  Unity adds this USS class to every instance of the Toggle element. Any styling applied to this class affects every Toggle located beside, or below the stylesheet in the visual tree. |
| `labelUssClassName` | `.unity-toggle__label` | USS class name for Labels in Toggle elements.  Unity adds this USS class to the `Label` sub-element of the `Toggle` if the Toggle has a Label. |
| `inputUssClassName` | `.unity-toggle__input` | USS class name of input elements in Toggle elements.  Unity adds this USS class to the input sub-element of the `Toggle`. The input sub-element provides responses to the manipulator. |
| `noTextVariantUssClassName` | `.unity-toggle--no-text` | USS class name of Toggle elements that have no text.  Unity adds this USS class to the `Toggle` if the Toggle does not have a label. |
| `checkmarkUssClassName` | `.unity-toggle__checkmark` | USS class name of Images in Toggle elements.  Unity adds this USS class to the Image sub-element of the `Toggle` that contains the checkmark image. |
| `textUssClassName` | `.unity-toggle__text` | USS class name of Text elements in Toggle elements.  Unity adds this USS class to Text sub-elements of the `Toggle`. |
| `mixedValuesUssClassName` | `.unity-toggle__mixed-values` | USS class name of Toggle elements that have mixed values  Unity adds this USS class to checkmark of the `Toggle` when it has mixed values. |
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

* [Foldout](UIE-uxml-element-Foldout.md)
* [Pseudo-classes](UIE-USS-Selectors-Pseudo-Classes.md)