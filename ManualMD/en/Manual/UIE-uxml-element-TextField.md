# TextField

A TextField lets users input text data into an application. It accepts and displays text input. You can set placeholder text to provide hints or instructions to the user on what to enter. You can also add validation functions to ensure that the entered data meets certain requirements.

**Note**: To align a TextField with other fields in an **Inspector** window, simply apply the `.unity-base-field__aligned` USS class to it. For more information, refer to [`BaseField`](../ScriptReference/UIElements.BaseField_1.md).

## Create a TextField

You can create a TextField with UI Builder, UXML, and C#. The following C# example creates a TextField that allows multiple lines text input:

```
var textField = new TextField();
textField.label = "Enter your text here:";
textField.multiline = true;
textField.maxLength = 140;
```

## Set a placeholder text

You can set a placeholder text for the element. You can also hide the placeholder text on focus.

**Note**: The placeholder text won’t display if you set a value for the element. To unset a value in UI Builder, right-click the **Value** field in the element’s Inspector tab and select **Unset**.

In C#, use the [`placeholder`](../ScriptReference/UIElements.ITextEdition-placeholder.md) and the [`hidePlaceholderOnFocus`](../ScriptReference/UIElements.ITextEdition-hidePlaceholderOnFocus.md) properties through [`textEdition`](../ScriptReference/UIElements.TextInputBaseField_1-textEdition.md):

```
myElement.textEdition.placeholder = "Enter your value here";
myElement.textEdition.hidePlaceholderOnFocus = true;
```

To style the placeholder text, use the `.unity-base-text-field__input--placeholder` USS selector.

## Customize the input text selection

Input text is selectable by default. You can customize the selection behaviors such as [`selectAllOnMouseUP`](../ScriptReference/UIElements.ITextSelection-selectAllOnMouseUp.md) and [`selectAllOnFocus`](../ScriptReference/UIElements.ITextSelection-selectAllOnFocus.md).

In C#, set them through [`textSelection`](../ScriptReference/UIElements.TextInputBaseField_1-textSelection.md):

```
myElement.textSelection.selectAllOnMouseUp = false;
myElement.textSelection.selectAllOnFocus = false;
```

You can also style the selected text using the `--unity-selection-color` USS custom property. For example, the following USS sets the color of the selected input text to yellow:

```
.unity-base-text-field__input {
    --unity-selection-color: yellow;
}
```

## Enable a vertical scrollbar

If you have enabled multiline text on the TextField, you can display a vertical scrollbar for the **text input field**.

In C#, set the [`verticalScrollerVisibility`](../ScriptReference/UIElements.TextInputBaseField_1-verticalScrollerVisibility.md) property to [`ScrollerVisibility.Auto`](../ScriptReference/UIElements.ScrollerVisibility.Auto.md) or [`ScrollerVisibility.AlwaysVisible`](../ScriptReference/UIElements.ScrollerVisibility.AlwaysVisible.md):

```
// Set the multiline property to true to enable multiple lines of text
textField.multiline = true;

// Set the vertical scrollbar visibility to AlwaysVisible
textField.verticalScrollerVisibility = ScrollerVisibility.AlwaysVisible;
```

## Control the behavior of the input text

You can use [events](UIE-Events.md) to control the behavior of the input text. The following example uses event callback functions to handle navigation and submission events for a text field. The example registers [navigation events](UIE-Navigation-Events.md) to the [TrickleDown](UIE-Events-Dispatching.md) phase before they’re handled by the `TextElement` base class:

```
// A callback function that executes submit logic to handle user input submission. 
// It stops the propagation of the event to prevent it from reaching other event handlers.
myField.RegisterCallback<NavigationSubmitEvent>((evt) => {

    // Submit logic
    evt.StopPropagation();
}, TrickleDown.TrickleDown);

// A callback function that executes focus logic to handle navigation between different UI elements. 
// If the navigation direction is up, down, left, or right, it stops the propagation of the event and instructs 
// the focusController to ignore the event.
myField.RegisterCallback<NavigationMoveEvent>((evt) => {

    if (evt.direction == NavigationMoveEvent.Direction.Up || evt.direction == NavigationMoveEvent.Direction.Down || evt.direction == NavigationMoveEvent.Direction.Left || evt.direction == NavigationMoveEvent.Direction.Right)
    {
        // Focus logic
        evt.StopPropagation();
        focusController.IgnoreEvent(evt);
    }

}, TrickleDown.TrickleDown);
```

## Example

The following UXML example creates a TextField:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <TextField label="UXML Field" name="the-uxml-field" value="It's snowing outside." />
</UXML>
```

The following C# example illustrates some of the customizable functionalities of the TextField:

```
/// <sample>
// Get a reference to the field from UXML and append a value to it.
var uxmlField = container.Q<TextField>("the-uxml-field");
uxmlField.value += "..";

// Create a new field, disable it, and give it a style class.
var csharpField = new TextField("C# Field");
csharpField.value = "It's snowing outside...";
csharpField.SetEnabled(false);
csharpField.AddToClassList("some-styled-field");
csharpField.value = uxmlField.value;
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

**C# class**: [`TextField`](../ScriptReference/UIElements.TextField.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`TextInputBaseField_1`](../ScriptReference/UIElements.TextInputBaseField_1.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `multiline` | `boolean` | Set this to true to allow multiple lines in the textfield and false if otherwise. |
| `value` | `string` | The string currently being exposed by the field. |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `auto-correction` | `boolean` | Determines if the touch screen keyboard auto correction is turned on or off. |
| `binding-path` | `string` | Path of the target property to be bound. |
| `emoji-fallback-support` | `boolean` | Specifies the order in which the system should look for Emoji characters when rendering text. If this setting is enabled, the global Emoji Fallback list will be searched first for characters defined as Emoji in the Unicode 14.0 standard. |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `hide-mobile-input` | `boolean` | Hides or shows the mobile input field. |
| `hide-placeholder-on-focus` | `boolean` | Hides the placeholder on focus. |
| `is-delayed` | `boolean` | If set to true, the value property isn’t updated until either the user presses Enter or the text field loses focus. |
| `keyboard-type` | `TouchScreenKeyboardType` | The type of mobile keyboard that will be used. |
| `label` | `string` | The string representing the label that will appear beside the field. If the string is empty, the label element is removed from the hierarchy. If the string is not empty, the label element is added to the hierarchy. |
| `mask-character` | `System.Char` | The character used for masking in a password field. |
| `max-length` | `int` | Maximum number of characters for the field. |
| `multiline` | `boolean` | multiline |
| `password` | `boolean` | Returns true if the field is used to edit a password. |
| `placeholder-text` | `string` | A short hint to help users understand what to enter in the field. |
| `readonly` | `boolean` | Returns true if the field is read only. |
| `select-all-on-focus` | `boolean` | Controls whether the element’s content is selected upon receiving focus. |
| `select-all-on-mouse-up` | `boolean` | Controls whether the element’s content is selected when you mouse up for the first time. |
| `select-line-by-triple-click` | `boolean` | Controls whether triple clicking selects the entire line under the mouse pointer or not. |
| `select-word-by-double-click` | `boolean` | Controls whether double-clicking selects the word under the mouse pointer. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |
| `value` | `string` | The value associated with the field. |
| `vertical-scroller-visibility` | [`UIElements.ScrollerVisibility`](../ScriptReference/UIElements.ScrollerVisibility.md) | Option for controlling the visibility of the vertical scroll bar in the `TextInputBaseField_1`. |

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
| `ussClassName` | `.unity-text-field` | USS class name of elements of this type. |
| `labelUssClassName` | `.unity-text-field__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-text-field__input` | USS class name of input elements in elements of this type. |
| `ussClassName` | `.unity-base-text-field` | USS class name of elements of this type. |
| `labelUssClassName` | `.unity-base-text-field__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-base-text-field__input` | USS class name of input elements in elements of this type. |
| `singleLineInputUssClassName` | `.unity-base-text-field__input--single-line` | USS class name of single line input elements in elements of this type. |
| `multilineInputUssClassName` | `.unity-base-text-field__input--multiline` | USS class name of multiline input elements in elements of this type. |
| `placeholderUssClassName` | `.unity-base-text-field__input--placeholder` | USS class name of input elements when placeholder text is shown |
| `ussClassName` | `.unity-base-field` | USS class name of elements of this type. |
| `labelUssClassName` | `.unity-base-field__label` | USS class name of labels in elements of this type. |
| `inputUssClassName` | `.unity-base-field__input` | USS class name of input elements in elements of this type. |
| `noLabelVariantUssClassName` | `.unity-base-field--no-label` | USS class name of elements of this type, when there is no label. |
| `labelDraggerVariantUssClassName` | `.unity-base-field__label--with-dragger` | USS class name of labels in elements of this type, when there is a dragger attached on them. |
| `mixedValueLabelUssClassName` | `.unity-base-field__label--mixed-value` | USS class name of elements that show mixed values |
| `alignedFieldUssClassName` | `.unity-base-field__aligned` | USS class name of elements that are aligned in a inspector element |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## USS custom properties

The following table outlines the custom properties that are available exclusively for the TextField element in USS:

| Property | Type | Description |
| --- | --- | --- |
| `--unity-cursor-color` | string | The color code of the text caret. |
| `--unity-selection-color` | string | The color code of the selected text. |

## Additional resources

* [Work with text](UIE-work-with-text.md)
* [Label](UIE-uxml-element-Label.md)
* [DoubleField](UIE-uxml-element-DoubleField.md)
* [LongField](UIE-uxml-element-LongField.md)
* [FloatField](UIE-uxml-element-FloatField.md)
* [IntegerField](UIE-uxml-element-IntegerField.md)
* [PropertyField](UIE-uxml-element-PropertyField.md)