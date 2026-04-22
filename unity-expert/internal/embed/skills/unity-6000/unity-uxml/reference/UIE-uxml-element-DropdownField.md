# DropdownField

A DropdownField lets users choose a value from a list. When clicking on the control, it opens a drawer that presents multiple options, allowing users to select one. If you want to always display all the choices, use [RadioButtonGroup](UIE-uxml-element-RadioButtonGroup.md).

**Note**: To align an element with other fields in an **Inspector** window, simply apply the `.unity-base-field__aligned` USS class to it. For more information, refer to [`BaseField`](../ScriptReference/UIElements.BaseField_1.md).

## Create a DropdownField

You can create a DropdownField with UI Builder, UXML, or C#.

**UI Builder**

To add options for a DropdownField in UI Builder, in the Inspector panel of the DropdownField, set the options in the **Choices** field.

To set the default value, in the Inspector of the DropdownField, enter the numbering sequence of the option in **Index**, starting at `0`.

**C#**

The following C# example create a DropdownField with three options. The DropdownField selects the first option by default but you can change the default value by setting the `index` property.

```
var dropdown = new DropdownField(new List<string> { "Option 1", "Option 2", "Option 3" }, 0);

// Add another option.
dropdown.choices.Add("Option 4");

// To return int value instead of string.
dropdown.index = 1; // Option 2
dropdown.value = "Option 3"; // Set to an existing value
// Assert that the index is set correctly.
Assert.IsTrue(dropdown.index == 2);

// Register to the value changed callback.
dropdown.RegisterValueChangedCallback(evt => Debug.Log(evt.newValue));

// Style the dropdown.
dropdown.style.width = 200;
dropdown.style.height = 50;
```

The UI looks like this:

![DropdownField Example](../uploads/Main/uxml/dropdown-example.png)

DropdownField Example

## Change styles for a DropdownField

You can override the default Panel Settings fields in a USS file to change the styles of a DropdownField:

1. In the **Hierarchy** window, select your UI Document (.uxml).
2. Set **Panel Settings** as your [Panel Settings](UIE-Runtime-Panel-Settings.md) asset. If you don’t have an existing Panel Settings asset, follow the instructions in the dialog to create one.
3. In your **Panel Settings** asset, [assign a Theme Style Sheet](UIE-tss.md#apply-a-tss).
4. In the **Project** window, select your TSS file.
5. Go to **Style Sheets** and select the **Add** (**+**) button.
6. Set your **Style Sheet** as your [USS file](UIE-about-uss.md).

## Examples

The following UXML example creates a DropdownField:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <DropdownField label="UXML Field" name="the-uxml-field" />
</UXML>
```

The following C# example illustrates some of the customizable functionalities of the DropdownField:

```
/// <sample>
// You can provide the list of choices by code, or by comma separated values in UXML
// <DropdownField .... choices="Option 1,Option 2,Option 3" .... />
var choices = new List<string> { "Option 1", "Option 2", "Option 3" };

// Get a reference to the dropdown field from UXML and assign a value to it.
var uxmlField = container.Q<DropdownField>("the-uxml-field");
uxmlField.choices = choices;
uxmlField.value = choices[0];

// Create a new dropdown with the same choices, disable it, and give it a style class.
var csharpField = new DropdownField("C# Field", choices, 0);
csharpField.SetEnabled(false);
csharpField.AddToClassList("some-styled-field");
container.Add(csharpField);

// Mirror the value of the UXML field into the C# field.
uxmlField.RegisterCallback<ChangeEvent<string>>((evt) =>
{
    csharpField.value = evt.newValue;
});
/// </sample>
```

To try this example live in Unity, go to **Window** > **UI Toolkit** > **Samples**.

## C# base class and namespace

**C# class**: [`DropdownField`](../ScriptReference/UIElements.DropdownField.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`PopupField_1`](../ScriptReference/UIElements.PopupField_1.md)

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `binding-path` | `string` | Path of the target property to be bound. |
| `choices` | `IList` | The list of choices to display in the popup menu. |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `index` | `int` | The currently selected index in the popup menu. Setting the index updates the `value` field and sends a property change notification. |
| `label` | `string` | The string representing the label that will appear beside the field. If the string is empty, the label element is removed from the hierarchy. If the string is not empty, the label element is added to the hierarchy. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |
| `value` | `int` | The currently selected value in the popup menu. |

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
| `ussClassName` | `.unity-popup-field` | USS class name of elements of this type. |
| `labelUssClassName` | `.unity-popup-field__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-popup-field__input` | USS class name of input elements in elements of this type. |
| `ussClassName` | `.unity-base-popup-field` | USS class name of elements of this type. |
| `textUssClassName` | `.unity-base-popup-field__text` | USS class name of text elements in elements of this type. |
| `arrowUssClassName` | `.unity-base-popup-field__arrow` | USS class name of arrow indicators in elements of this type. |
| `labelUssClassName` | `.unity-base-popup-field__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-base-popup-field__input` | USS class name of input elements in elements of this type. |
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

* [RadioButtonGroup](UIE-uxml-element-RadioButtonGroup.md)