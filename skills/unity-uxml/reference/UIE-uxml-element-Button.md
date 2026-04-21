# Button

Use the Button element to create clickable buttons in a UI. For example, when a user clicks or taps on a Button element, it triggers an action or event, such as opening a new window, submitting a form, or playing a sound effect.

## Create a Button

You can create a Button with UI Builder, UXML, or C#. The following C# example creates a Button with a label:

```
var button = new Button(() => { Debug.Log("Button clicked"); }) { text = "Click me" };
```

## Add an icon to a Button

You can enhance the visual appeal of your Button by incorporating an icon, such as a Texture, **Render Texture**, **Sprite**, or Vector image asset that exists in your project. For information on how to reference an image asset, refer to [Assets](UIE-USS-PropertyTypes.md#assets).

To add an icon for a Button with UI Builder:

* From the Button’s **Inspector** tab, choose the desired icon from the **Icon Image** dropdown list.
* Alternatively, drag the icon from the Assets window to the **Icon Image** field located in the Button’s Inspector tab.

To add an icon for a Button with UXML, provide the image source to the `icon-image` attribute:

```
<ui:Button name="Button" text="Button text" icon-image="/path/to/image-file.png" />
```

To add an icon for a Button with C#, assign the image source to the [`iconImage`](../ScriptReference/UIElements.Button-iconImage.md) property:

```
Button myButton = new Button();
var buttonIconImage = Resources.Load<Texture2D>("image-file");

myButton.text = "Button text";
myButton.iconImage = buttonIconImage;
```

## Change the Button icon position

When you add an icon to a Button, it adds two [read-only](UIB-structuring-ui-elements.md#read-only-elements) sub-elements to the Button:

* An [Image](UIE-uxml-element-Image.md) element to define the icon image.
* A [TextElement](UIE-uxml-element-TextElement.md) to define the button text.

You can’t edit the sub-elements but you can change their layout using the USS `flex-direction` property for the Button element. By default, the icon is added on the left of the button text.

To reposition the button icon, adjust the Button’s `flex-direction` with the following:

* To shift the icon to the right of the button text, set it to `row-reverse`.
* To move the icon above the button text, set it to `column`.
* To place the icon below the button text, set it to `column-reverse`.

## Use sub-elements of a Button

The Button element allows you to provide additional information to the user with the `text`, `iconImage` properties, and the `background-image` style property. As with any **visual element**, you can also add sub-elements in a Button’s hierarchy, such as a Label or Image, to provide additional information to the user if you want to have more fine-grained control over the appearance and behavior of those elements.

In general, use sub-elements in the following situations:

* Customization: Using sub-elements allows you to customize the appearance and behavior of each individual element separately. For example, you may want to use a custom font or color for the label text or add a specific animation to the image. When you add an Image element to your Button, you can add the element from your **scene** directly. In addition, an Image element also supports scale mode and repeat mode.
* Dynamic content: If the content of the label or image is dynamic, using sub-elements allows you to update the content of each element separately without affecting the other properties of the button.
* Interaction: If you want to add interactivity to a specific element within the button, such as allowing the user to click only the label (disregarding clicks on other elements of the button) to trigger an action, using sub-elements makes this possible.

Use properties for the following purposes:

* Simple designs: If the design of the button is simple and doesn’t require customization or dynamic content, using properties is recommended.
* Performance: If you have many buttons in your UI, using sub-elements for each one may negatively impact performance. In this case, it’s more efficient to use properties to set the appearance of the button.
* Consistency: If you want to maintain consistency across your UI, using properties can ensure that all buttons have a consistent appearance and behavior.

## Examples

The following UXML example creates a Button:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <Button text="UXML Button" name="the-uxml-button" />
</UXML>
```

The following C# example illustrates some of the customizable functionalities of the Button:

```
/// <sample>
// Action to perform when button is pressed.
// Toggles the text on all buttons in 'container'.
Action action = () =>
{
    container.Query<Button>().ForEach((button) =>
    {
        button.text = button.text.EndsWith("Button") ? "Button (Clicked)" : "Button";
    });
};

// Get a reference to the Button from UXML and assign it its action.
var uxmlButton = container.Q<Button>("the-uxml-button");
uxmlButton.RegisterCallback<MouseUpEvent>((evt) => action());

// Create a new Button with an action and give it a style class.
var csharpButton = new Button(action) { text = "C# Button" };
csharpButton.AddToClassList("some-styled-button");
container.Add(csharpButton);
/// </sample>
```

To try this example live in Unity, go to **Window** > **UI Toolkit** > **Samples**.

## C# base class and namespace

**C# class**: [`Button`](../ScriptReference/UIElements.Button.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`TextElement`](../ScriptReference/UIElements.TextElement.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `icon-image` | `Object` | The Texture, Sprite, or VectorImage that will represent an icon within a Button element. |
| `text` | `string` | text |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `binding-path` | `string` | Path of the target property to be bound. |
| `display-tooltip-when-elided` | `boolean` | When true, a tooltip displays the full version of elided text, and also if a tooltip had been previously provided, it will be overwritten. |
| `double-click-selects-word` | `boolean` | Controls whether double-clicking selects the word under the mouse pointer. |
| `emoji-fallback-support` | `boolean` | Specifies the order in which the system should look for Emoji characters when rendering text. If this setting is enabled, the global Emoji Fallback list will be searched first for characters defined as Emoji in the Unicode 14.0 standard. |
| `enable-rich-text` | `boolean` | When false, rich text tags will not be parsed. |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `parse-escape-sequences` | `boolean` | Determines how escape sequences are displayed. When set to `true`, escape sequences (such as `\n`, `\t`) are parsed and transformed into their corresponding characters. For example, ‘\n’ will insert a new line. When set to `false`, escape sequences are displayed as raw text (for example, `\n` is shown as the characters ’' followed by ‘n’). |
| `selectable` | `boolean` | Returns true if the field is selectable. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |
| `text` | `string` | The text to be displayed.  Changing this value will implicitly invoke the `INotifyValueChanged_1.value` setter, which will raise a `ChangeEvent_1` of type string. |
| `triple-click-selects-line` | `boolean` | Controls whether triple-clicking selects the entire line under the mouse pointer. |

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
| `ussClassName` | `.unity-button` | USS class name of elements of this type.  Unity adds this USS class to every instance of the Button element. Any styling applied to this class affects every button located beside, or below the stylesheet in the visual tree. |
| `iconUssClassName` | `.unity-button--with-icon` | The USS class name for Button elements with an icon.  Unity adds this USS class to an instance of the Button element if the instance’s `Button.iconImage` property contains a valid Texture. Any styling applied to this class affects every button with an icon located beside, or below the stylesheet in the visual tree. |
| `iconOnlyUssClassName` | `.unity-button--with-icon-only` | The USS class name for Button elements with an icon only, no text.  Unity adds this USS class to an instance of the Button element if the instance’s `Button.iconImage` property contains a valid Texture and no text is set. Any styling applied to this class affects every button with an icon located beside, or below the stylesheet in the visual tree. |
| `ussClassName` | `.unity-text-element` | USS class name of elements of this type. |
| `selectableUssClassName` | `.unity-text-element__selectable` | USS class name of selectable text elements. |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## Additional resources

* [Image](UIE-uxml-element-Image.md)
* [TextElement](UIE-uxml-element-TextElement.md)
* [ToggleButtonGroup](UIE-uxml-element-ToggleButtonGroup.md)