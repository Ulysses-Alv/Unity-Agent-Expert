# Optimize performance of moving elements at runtime

When you move a large amount of elements at runtime, performance optimization is crucial, especially for applications requiring smooth and responsive interactions.

Certain CSS properties, such as top, left, width, height, or flex, can trigger layout recalculations when updated. These recalculations dirty the layout and might cascade updates across multiple elements, significantly impacting rendering performance. Use these properties sparingly and only when absolutely necessary.

## Best practices for moving elements

The most performance-efficient way to move elements is to use the [`style.translate`](UIE-Transform.md) property. This property adjusts an element’s position dynamically without recalculating the layout. Unlike other positioning methods, `style.translate` operates at the transform stage. It minimizes the computational overhead and improves rendering performance by reducing dependency on CPU-intensive processes.

## Use usage hints to reduce draw calls and geometry regeneration

To further enhance performance, leverage usage hints to optimize how elements are processed during runtime. Usage hints help:

* Reduce draw calls.
* Avoid unnecessary geometry regeneration.

If the transform of the element, such as position, rotation, or scale, changes frequently, set usage hint to `DynamicTransform`. This pushes transform updates directly to the GPU, bypassing CPU **mesh** updates and improving performance.

You can set the usage hints in UI Builder, UXML, or C#. The following examples set the usage hints to `DynamicTransform`:

**UXML**:

```
<ui:VisualElement usage-hints="DynamicTransform" />
```

**C#**:

```
visualElement.usageHints = UsageHints.DynamicTransform;
```

The supported usage hints are:

* [`DynamicTransform`](../ScriptReference/UIElements.UsageHints.DynamicTransform.md)
* [`GroupTransform`](../ScriptReference/UIElements.UsageHints.GroupTransform.md)
* [`MaskContainer`](../ScriptReference/UIElements.UsageHints.MaskContainer.md)
* [`DynamicColor`](../ScriptReference/UIElements.UsageHints.DynamicColor.md)

## Flow of updating and rendering a visual element

The following diagram illustrates the flow of updating and rendering a **visual element** at runtime:

![Flow of updating and rendering a visual element](../uploads/Main/uitk/flow-of-rendering-elements.png)

Flow of updating and rendering a visual element

For an example that optimizes rendering and updates for frequently changing elements at runtime, refer to [Move elements at runtime](UIE-move-elements-at-runtime.md).

## Additional resources

* [`usageHints`](../ScriptReference/UIElements.UsageHints.md)
* [`VisualElement.usageHints`](../ScriptReference/UIElements.VisualElement-usageHints.md)
* [Move elements at runtime](UIE-move-elements-at-runtime.md)