# Configure the custom control name and visibility in UI Builder

Use the [`UxmlElement`](../../ScriptReference/UIElements.UxmlElementAttribute.md) attribute to change how your custom controls appear in UXML and in the UI Builder library.

## Customize UXML tag name

By default, the tag name in UXML for your custom control is the C# class name. While it’s not recommended to use a different tag name, you can customize it if needed.

To customize a UXML tag name, add a [`name`](../../ScriptReference/UIElements.UxmlElementAttribute-name.md) argument to the [`UxmlElement`](../../ScriptReference/UIElements.UxmlElementAttribute.md) attribute.

**Note**: The tag name must be unique and you must reference the classes’ namespace in UXML.

For example, if you create the following custom button:

```
using UnityEngine.UIElements;

namespace MyNamespace
{
    [UxmlElement("MyButton")]
    public partial class CustomButtonElement : Button
    {
    }
}
```

You can then reference the custom button in UXML with the custom name or the C# class name:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements">
    <MyNamespace.MyButton />
    <MyNamespace.CustomButtonElement />
</ui:UXML>
```

## Configure path and visibility

By default, custom controls appear in the UI Builder library under a path that matches their C# namespace. You can use the [`libraryPath`](../../ScriptReference/UIElements.UxmlElementAttribute-libraryPath.md) argument to customize the path of the `UxmlElement` attribute.

By default, namespaces that start with `Unity`, `UnityEngine`, or `UnityEditor` are reserved, so elements in these namespaces don’t appear in the UI Builder library. You can use the [`visible`](../../ScriptReference/UIElements.UxmlElementAttribute-visibility.md) argument to override this and make your controls visible.

The following example shows how to make a custom control visible in the UI Builder library under a custom path, and how to hide another custom control:

```
using UnityEngine.UIElements;

namespace Unity.MyGame
{
    // This element is visible in the UI Builder library under "My Game/Inventory".
    [UxmlElement(libraryPath = "My Game/Inventory", visibility = LibraryVisibility.Visible)]
    public partial class VisibleItem : VisualElement
    {
        [UxmlAttribute]
        public string description { get; set; }
    }
}

// This element is hidden from the UI Builder library.
// However, you can still use it in UXML or C# code.
[UxmlElement(visibility = LibraryVisibility.Hidden)]
public partial class HiddenItem : VisualElement
{
    [UxmlAttribute]
    public string description { get; set; }
}
```

## Additional resources

* [Create custom controls](../UIE-create-custom-controls.md)