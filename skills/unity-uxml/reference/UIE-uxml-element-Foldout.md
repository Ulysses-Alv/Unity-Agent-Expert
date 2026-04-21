# Foldout

A Foldout is a collapsible section of a UI. The Foldout hides or reveals the elements it contains when you click its header.

![Foldout](../uploads/Main/uxml/foldout.gif)

Foldout

## Create a Foldout

A Foldout has the following elements:

* **Label**: A label you can use as the name of the Foldout.
* **Toggle**: Click the toggle to expand or collapse the container sub-element. It is styled to look like an arrow rather than the default checkbox. When the container is collapsed, the arrow points right. When the container is expanded, the arrow points down.
* **Container**: A container contains the **visual elements** that you want to show or hide when you toggle the Foldout.

You can add elements to a Foldout’s container, and the show-or-hide function works automatically.

The following UI Document creates a Foldout with two buttons in its container. The two buttons show or hide when you toggle the Foldout.

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements" xsi="http://www.w3.org/2001/XMLSchema-instance" engine="UnityEngine.UIElements" editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
    <ui:Foldout text="Foldout" name="MyFoldout" value="true">
        <ui:Button text="First item" />
        <ui:Button text="Second item" />
    </ui:Foldout>
</ui:UXML>
```

## Respond to user actions

A Foldout and its Toggle sub-element respond to [Change events](UIE-Change-Events.md). The Toggle hides or shows the Foldout container.

You can bind a Foldout to a Boolean variable or access its value through the [`INotifyValueChanged<bool>`](../ScriptReference/UIElements.INotifyValueChanged_1.md) interface.

## Apply styles to the direct child of a Foldout

When you add elements to a Foldout, they’re added under a child element of the Foldout that acts as a container. To add a style to this container, do this in USS:

```
/* Style only direct children of Foldout */
Foldout > .unity-foldout__content {
    margin-left: 20px;
}
```

## Examples

The following C# code snippet creates a Foldout and checks whether the foldout is expanded:

```
// Create a new foldout, add two elements to it and add it to the container
var cSharpFoldout = new Foldout {text = "Elements"};
cSharpFoldout.Add(new Label("Indented Label"));
cSharpFoldout.Add(new Slider("Indented Slider", 0, 100));
container.Add(cSharpFoldout);

cSharpFoldout.RegisterValueChangedCallback(e =>
{
    // Check whether the foldout is expanded
    if (cSharpFoldout.value)
    {
        Debug.Log("Foldout is expanded");
    }
    else
    {
        Debug.Log("Foldout is collapsed");
    }
});
```

## C# base class and namespace

**C# class**: [`Foldout`](../ScriptReference/UIElements.Foldout.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`BindableElement`](../ScriptReference/UIElements.BindableElement.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `focusable` | `boolean` | focusable |
| `text` | `string` | The label text for the toggle. |
| `toggle-on-label-click` | `boolean` | Whether to toggle the foldout state when the user clicks the label. |
| `value` | `boolean` | This is the state of the Foldout’s toggle. It is true if the `Foldout` is open and its contents are visible, and false if the Foldout is closed, and its contents are hidden. |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `binding-path` | `string` | Path of the target property to be bound. |
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
| `ussClassName` | `.unity-foldout` | The USS class name for Foldout elements.  Unity adds this USS class to every instance of a `Foldout`. Any styling applied to this class affects every Foldout located beside, or below the stylesheet in the visual tree. |
| `toggleUssClassName` | `.unity-foldout__toggle` | The USS class name of Toggle sub-elements in Foldout elements.  Unity adds this USS class to the `Toggle` sub-element of every `Foldout`. Any styling applied to this class affects every Toggle sub-element located beside, or below the stylesheet in the visual tree. |
| `contentUssClassName` | `.unity-foldout__content` | The USS class name for the content element in a Foldout.  Unity adds this USS class to the `VisualElement` that contains the elements to be shown or hidden. Any styling applied to this class affects every foldout container located beside, or below the stylesheet in the visual tree. |
| `inputUssClassName` | `.unity-foldout__input` | The USS class name for the Label element in a Foldout.  Unity adds this USS class to the `VisualElement` that contains the `Toggle` input elements. Any styling applied to this class affects every foldout container located beside, or below the stylesheet in the visual tree. |
| `checkmarkUssClassName` | `.unity-foldout__checkmark` | The USS class name for the Label element in a Foldout.  Unity adds this USS class to the `VisualElement` that represents the checkmark of the `Toggle` sub-element of every `Foldout`. Any styling applied to this class affects every foldout container located beside, or below the stylesheet in the visual tree. |
| `textUssClassName` | `.unity-foldout__text` | The USS class name for the Label element in a Foldout.  Unity adds this USS class to the `Label` in the `Toggle` sub-element of every `Foldout`. Any styling applied to this class affects every foldout container located beside, or below the stylesheet in the visual tree. |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## Additional resource

* [Change events](UIE-Change-Events.md)
* [Toggle](UIE-uxml-element-Toggle.md)