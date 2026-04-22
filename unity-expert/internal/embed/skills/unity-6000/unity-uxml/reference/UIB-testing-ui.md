# Test UI in UI Builder

You can test your UI inside the UI Builder, and use debugging tools to investigate styles.

## Test UI in Preview mode

To test your UI inside the UI Builder, directly inside the **Canvas**, you can enable **Preview** mode, found in the ****Viewport****’s **toolbar**. **Preview** mode removes the UI Builder-specific picking overlay and handles from the **Canvas**. You can tell if you have **Preview** mode enabled by looking for an orange border around the entire **Viewport**:

![PreviewMode](../uploads/Main/UIBuilder/PreviewMode.png)

PreviewMode

With **Preview** mode enabled, you can test the following:

* State-based controls like `Foldout`, which you can click on to expand and see how the rest of the UI reacts.
* Input fields like `IntegerField`, where you can test input validation.
* Large containers like `ScrollView`, where you can test scrolling up and down.
* `:hover` pseudo states to check hover-only styles.

In **Preview** mode, you can still work on your UI. However, **Canvas** picking and manipulators are turned off in **Preview** mode.

## Debug styles

If you don’t know where a style comes from, you can find the styles for an element in the **Matching Selectors** section in the UI Builder’s **Inspector**.

1. In the **Canvas** of the UI Builder, select the element.
2. In the Inspector window, expand **StyleSheet** > **Matching Selectors**.

![An example of Matching Selectors](../uploads/Main/UIBuilder/MatchingSelectorsFoldout.png)

The **Matching Selectors** section displays the following:

* USS selectors from your own style sheets
* USS selectors from the default Unity theme

**Note**: USS selectors that appear lower in the list always override the same style properties in higher USS selectors.

If the **Matching Selectors** section doesn’t provide enough information, you can use the [UI Toolkit Debugger](UIE-ui-debugger.md). The UI Toolkit Debugger is a tool that you can use to inspect and debug your UI elements in real-time. It provides a visual representation of your UI hierarchy. You use it to examine the state and properties of each UI element.

## Additional resources

* [Style UI with USS](UIE-USS.md)
* [USS selectors](UIE-USS-Selectors.md)
* [UI Toolkit Debugger](UIE-ui-debugger.md)