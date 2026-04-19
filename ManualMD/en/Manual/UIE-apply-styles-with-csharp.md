# Apply styles in C# scripts

You can write to [`style`](../ScriptReference/UIElements.VisualElement-style.md) to set style values to an element. However, to get the actually rendered styles of an element, read from [`resolvedStyle`](../ScriptReference/UIElements.VisualElement-resolvedStyle.md).

## Set styles

In a C# script, you can set styles directly to the `style` properties of a **visual element**. For example, the following code sets the background color of a button to red:

```
button.style.backgroundColor = Color.red
```

**Note**: Setting a property like `backgroundColor` is a direct instruction. Layout-related properties, such as `top`, `left`, `width`, and `height`, are calculated by the layout engine, which considers stylesheets, parent properties, and inheritance. To get the actual rendered style values for an element, you must read from the read-only `resolvedStyle` property.

You can also add a Unity style sheet (USS) to any visual element. Unity represents USS files as [`StyleSheet`](../ScriptReference/UIElements.StyleSheet.md) objects in C# **scripts**.

To add style sheets to a visual element:

1. Load `StyleSheet` objects with standard Unity APIs, such as [`AssetDatabase.Load()` or `Resources.Load()`](UIE-manage-asset-reference.md).
2. Use the `styleSheets` property of a visual element to add the `StyleSheet` object.

For example, given a style sheet in the local variable `styleSheet` and an element in the local variable `element`, the following example adds the style sheet to the element:

```
element.styleSheets.Add(styleSheet);
```

**Note**: Style rules apply to the visual element and all its descendants, but don’t apply to the parent or siblings of the element. Any change to the USS file automatically refreshes the UI that uses this style sheet.

## Get resolved styles

Style values on an element are computed from various sources, including multiple applied classes, inheritance from ancestors, and inline styles from UXML or C# code. These values might change from frame to frame. The `style` only holds the inline styles for the element and doesn’t reflect other sources. The `resolvedStyle` has the final calculated values, considering all sources on the current frame.

For example, when you use the inline style to set the height for an element, both the `style` and `resolvedStyle` start with the same value. When the element is added to the hierarchy, `resolvedStyle.height` can be `NaN` until the layout updates. If you define the height in a class as a percentage, the computed width relies on parent properties such as `border-height` and `padding`. Although `style.height` might give a relative value, such as for transitions that can change the value, `resolvedStyle.height` gives the actual rendered height.

To get the resolved style when the geometry changes, you can use the [`GeometryChangedEvent`](../ScriptReference/UIElements.GeometryChangedEvent.md) event. This event is triggered when the layout of a VisualElement changes, which includes changes in size and position. You can register a callback for this event and in the callback, you can access the `resolvedStyle` property of the VisualElement to get the final computed styles.

The following example creates a custom Editor window and logs the resolved height of an element. The height of the element changes if you resize the window:

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;

public class ResolvedStyles : EditorWindow
{
    [MenuItem("Window/UI Toolkit/ResolvedStyles")]
    public static void ShowExample()
    {
        GetWindow<ResolvedStyles>();
    }
    
    private void OnEnable()
    {
        titleContent = new GUIContent("Resolved Styles");
    }

    public void CreateGUI()
    {
        VisualElement root = rootVisualElement;
        
        // Element that is tracked.
        // When you resize the Editor window, the inner content is not necessarily updated
        // during the drag operation. The resolved height field is updated whenever the drag
        // operation is complete.
        var element = new VisualElement
        {
            style =
            {
                flexGrow = 1,
                backgroundColor = Color.red
            }
        };
        root.Add(element);

        // Register a callback for the GeometryChangedEvent
        element.RegisterCallback<GeometryChangedEvent>(OnGeometryChanged);
    }

    // Callback for the GeometryChangedEvent
    public void OnGeometryChanged(GeometryChangedEvent evt)
    {
        // Get the VisualElement that triggered the event
        VisualElement element = evt.target as VisualElement;

        // Get the resolved style of the VisualElement
        float height = element.resolvedStyle.height;

        // Log the resolved of the VisualElement
        Debug.Log("Resolved height: " + height);
    }
}
```

If the element’s geometry doesn’t change, you can add a scheduler to periodically check the resolved style of the element:

```
element.schedule.Execute(() =>
{
    Debug.Log(element.resolvedStyle.height);
}).Every(100);
```

## Additional resources

* [Manage UI asset references from C# scripts](UIE-manage-asset-reference.md)
* [Get custom styles](UIE-get-custom-styles.md)