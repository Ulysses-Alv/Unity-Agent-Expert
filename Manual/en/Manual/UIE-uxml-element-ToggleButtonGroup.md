# ToggleButtonGroup

Use ToggleButtonGroup to enable single or multiple selections from a logical group of [Button](UIE-uxml-element-Button.md) elements. The ToggleButtonGroup consists of a label and a set of interactable Button elements.

**Note**: To align an element with other fields in an **Inspector** window, simply apply the `.unity-base-field__aligned` USS class to it. For more information, refer to [`BaseField`](../ScriptReference/UIElements.BaseField_1.md).

## Create a ToggleButtonGroup

You can create a ToggleButtonGroup with UI Builder, UXML, and C#.

**UI Builder**

To create a ToggleButtonGroup in UI Builder:

1. Drag **ToggleButtonGroup** from the **Library** panel into the **Hierarchy** panel.
2. Add a set of **Button** elements to the **ToggleButtonGroup** element as its child elements.

To create a ToggleButtonGroup in UXML, add a ToggleButtonGroup element to a UXML document and add a set of Button elements as its child elements. For example:

**C#**

To create a ToggleButtonGroup in C#, create a new instance of a [ToggleButtonGroup](../ScriptReference/UIElements.ToggleButtonGroup.md) object and add a set of [Buttons](../ScriptReference/UIElements.Button.md).

By default, you can select one button at a time. To enable multiple selections, set the `is-multiple-selection` to `true`.

To enable empty selection, set the `allow-empty-selection` to `true`.

## Example

The following UXML example creates a ToggleButtonGroup with Buttons:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <ToggleButtonGroup label="UXML Toggle Button Group">
        <Button text="A"/>
        <Button text="B"/>
        <Button text="C"/>
        <Button text="D"/>
        <Button text="E"/>
    </ToggleButtonGroup>
</UXML>
```

The following C# example illustrates some customizable functionalities of the ToggleButtonGroup and its Buttons:

```
/// <sample>
// Create custom buttons with Text value only.
var csharpToggleButtonGroupWithButtonTextOnly = new ToggleButtonGroup("C# Toggle Button Group Buttons with Text");
csharpToggleButtonGroupWithButtonTextOnly.Add(new Button() { text = "one", tooltip = "custom button one" });
csharpToggleButtonGroupWithButtonTextOnly.Add(new Button() { text = "two", tooltip = "custom button two" });
container.Add(csharpToggleButtonGroupWithButtonTextOnly);

// Create custom buttons with IconImage value only.
var csharpToggleButtonGroupWithButtonIconOnly = new ToggleButtonGroup("C# Toggle Button Group Buttons with Icons");
csharpToggleButtonGroupWithButtonIconOnly.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_PlayButton"), tooltip = "Play" });
csharpToggleButtonGroupWithButtonIconOnly.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_PauseButton"), tooltip = "Pause" });
csharpToggleButtonGroupWithButtonIconOnly.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_StepButton"), tooltip = "Step" });
container.Add(csharpToggleButtonGroupWithButtonIconOnly);

// Create custom buttons with IconImage and Text.
var csharpToggleButtonGroupWithButtonIconAndText = new ToggleButtonGroup("C# Toggle Button Group Buttons with Icons and Text");
csharpToggleButtonGroupWithButtonIconAndText.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_PlayButton"), text = "Play", tooltip = "Play" });
csharpToggleButtonGroupWithButtonIconAndText.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_PauseButton"), text = "Pause", tooltip = "Pause" });
csharpToggleButtonGroupWithButtonIconAndText.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_StepButton"), text = "Step", tooltip = "Step" });
container.Add(csharpToggleButtonGroupWithButtonIconAndText);

// Create custom buttons with IconImage, Text and AllowEmptySelection.
var csharpToggleButtonGroupSingleSelectionAndAllowEmptySelection = new ToggleButtonGroup("C# Toggle Button Group Buttons with Allow Empty Selection") { allowEmptySelection = true };
csharpToggleButtonGroupSingleSelectionAndAllowEmptySelection.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_PlayButton"), text = "Play", tooltip = "Play" });
csharpToggleButtonGroupSingleSelectionAndAllowEmptySelection.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_PauseButton"), text = "Pause", tooltip = "Pause" });
csharpToggleButtonGroupSingleSelectionAndAllowEmptySelection.Add(new Button() { iconImage = EditorGUIUtility.FindTexture("d_StepButton"), text = "Step", tooltip = "Step" });
container.Add(csharpToggleButtonGroupSingleSelectionAndAllowEmptySelection);

// Create a ToggleButtonGroup that allows multiple active selections.
var csharpToggleButtonGroupAllowMultiSelect = new ToggleButtonGroup("C# Toggle Button Group with Multiple Selection Enabled") { isMultipleSelection = true };
csharpToggleButtonGroupAllowMultiSelect.Add(new Button() { text = "X", tooltip = "tooltip text for X" });
csharpToggleButtonGroupAllowMultiSelect.Add(new Button() { text = "Y", tooltip = "tooltip text for Y" });
csharpToggleButtonGroupAllowMultiSelect.Add(new Button() { text = "Z", tooltip = "tooltip text for Z" });
container.Add(csharpToggleButtonGroupAllowMultiSelect);

// Create a ToggleButtonGroup that allows multiple active selections and allow empty selection.
var csharpToggleButtonGroupAllowMultiSelectWithAllowEmptySelection = new ToggleButtonGroup("C# Toggle Button Group with Multiple Selection and Allow Empty Selection") { isMultipleSelection = true, allowEmptySelection = true };
csharpToggleButtonGroupAllowMultiSelectWithAllowEmptySelection.Add(new Button() { text = "X", tooltip = "tooltip text for X" });
csharpToggleButtonGroupAllowMultiSelectWithAllowEmptySelection.Add(new Button() { text = "Y", tooltip = "tooltip text for Y" });
csharpToggleButtonGroupAllowMultiSelectWithAllowEmptySelection.Add(new Button() { text = "Z", tooltip = "tooltip text for Z" });
container.Add(csharpToggleButtonGroupAllowMultiSelectWithAllowEmptySelection);

// Create ToggleButtonGroup with a custom class that sets the text's font style to Bold.
var csharpToggleButtonGroupWithCustomClass = new ToggleButtonGroup("C# Toggle Button Group with a Custom Class");
csharpToggleButtonGroupWithCustomClass.AddToClassList("my-custom-style");
csharpToggleButtonGroupWithCustomClass.Add(new Button() { text = "Button A", tooltip = "Bolded font Button A" });
csharpToggleButtonGroupWithCustomClass.Add(new Button() { text = "Button B", tooltip = "Bolded font Button B" });
container.Add(csharpToggleButtonGroupWithCustomClass);
/// </sample>
```

To try this example live in Unity, go to **Window** > **UI Toolkit** > **Samples**.

## C# base class and namespace

**C# class**: [`ToggleButtonGroup`](../ScriptReference/UIElements.ToggleButtonGroup.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`BaseField_1`](../ScriptReference/UIElements.BaseField_1.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `allow-empty-selection` | `boolean` | Allows having all buttons to be unchecked when set to true.  When the property value is false, the control will automatically set the first available button to checked. |
| `is-multiple-selection` | `boolean` | Whether all buttons can be selected. |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `binding-path` | `string` | Path of the target property to be bound. |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `label` | `string` | The string representing the label that will appear beside the field. If the string is empty, the label element is removed from the hierarchy. If the string is not empty, the label element is added to the hierarchy. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |
| `value` | [`UIElements.ToggleButtonGroupState`](../ScriptReference/UIElements.ToggleButtonGroupState.md) | The value associated with the field. |

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
| `ussClassName` | `.unity-toggle-button-group` | USS class name of elements for this type. |
| `containerUssClassName` | `.unity-toggle-button-group__container` | USS class name of container element of this type. |
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

* [Button](UIE-uxml-element-Button.md)