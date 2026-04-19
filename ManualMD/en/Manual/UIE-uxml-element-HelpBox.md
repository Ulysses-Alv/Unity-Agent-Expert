# HelpBox

A HelpBox displays a message to the user. You can set the type of message to display, such as a warning or an error. You can also set the message to display as rich text.

## Create a HelpBox

You can create a HelpBox with UXML and C#. The following C# example creates a HelpBox with the message “Help text” and the information icon:

```
HelpBox helpBox = new HelpBox("Help text", HelpBoxMessageType.Info);
```

## Example

The following UXML example creates a HelpBox with a warning message:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <HelpBox text="UXML Help Box" message-type="Info" name="the-uxml-help-box" />
</UXML>
```

The following C# example illustrates some of the customizable functionalities of the HelpBox:

```
/// <sample>
// Get a reference to the help box from UXML and update its text.
var uxmlHelpBox = container.Q<HelpBox>("the-uxml-help-box");
uxmlHelpBox.text += " (Updated in C#)";

// Create a new help box and give it a style class.
var csharpHelpBox = new HelpBox("This is a help box", HelpBoxMessageType.Warning);
csharpHelpBox.AddToClassList("some-styled-help-box");
container.Add(csharpHelpBox);
/// </sample>
```

To try this example live in Unity, go to **Window** > **UI Toolkit** > **Samples**.

## C# base class and namespace

**C# class**: [`HelpBox`](../ScriptReference/UIElements.HelpBox.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`VisualElement`](../ScriptReference/UIElements.VisualElement.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `message-type` | [`UIElements.HelpBoxMessageType`](../ScriptReference/UIElements.HelpBoxMessageType.md) | The type of message. |
| `text` | `string` | The message text. |

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
| `ussClassName` | `.unity-help-box` | The USS class name for Elements of this type. |
| `labelUssClassName` | `.unity-help-box__label` | The USS class name for labels in Elements of this type. |
| `iconUssClassName` | `.unity-help-box__icon` | The USS class name for images in Elements of this type. |
| `iconInfoUssClassName` | `.unity-help-box__icon--info` | The USS class name for the `HelpBoxMessageType.Info` state in Elements of this type. |
| `iconwarningUssClassName` | `.unity-help-box__icon--warning` | The USS class name for the `HelpBoxMessageType.Warning` state in Elements of this type. |
| `iconErrorUssClassName` | `.unity-help-box__icon--error` | The USS class name for the `HelpBoxMessageType.Error` state in Elements of this type. |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## Additional resources

* [Work with text](UIE-work-with-text.md)