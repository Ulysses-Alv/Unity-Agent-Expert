# Panels

The panel is the parent object of a [visual tree](UIE-VisualTree.md). It owns the `rootVisualElement` but itself isn’t a **visual element**. A visual tree must connect to a panel for the visual elements inside a tree to render. All panels belong to either an [Editor Window](../ScriptReference/EditorWindow.md) or a runtime [UIDocument](../ScriptReference/UIElements.UIDocument.md). The panel also handles focus control and event dispatching for the visual tree.

Every element in a visual tree holds a direct reference to the panel that holds the visual tree. To verify the connection of a [`VisualElement`](../ScriptReference/UIElements.VisualElement.md) to a panel, you can test the [`panel`](../ScriptReference/UIElements.VisualElement-panel.md) property of this element. When the visual element isn’t connected, the test returns `null`.

## Additional resources

* [Dispatch events](UIE-Events-Dispatching.md)
* [The visual tree](UIE-VisualTree.md)
* [The Panel Settings asset](UIE-Runtime-Panel-Settings.md)
* [Relative and absolute positioning C# example](UIE-relative-absolute-positioning-example.md)