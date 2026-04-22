# Auto sizing text elements

You can enable auto sizing for text with **Best Fit**. The Best Fit feature allows you to set a minimum and maximum font size for the text element. The text element adjust its font size dynamically based on the content and available space within the range.

Best Fit also supports the following features:

* **Word wrapping**: Automatically wraps text to fit within the specified width.
* **Ellipsis truncation**: Automatically truncates text with an ellipsis (…) if it exceeds the available space.
* **Multiple alignments**: Supports left, center, and right alignment.
* **Newline**: Automatically inserts line breaks when necessary.

You can enable auto sizing in UI Builder, USS, or C# **scripts**.

## In UI Builder

1. Select the text element you want to modify.
2. In the **Inspector** panel, select **Text** > **Auto Size**.
3. Enter the minimum and maximum font sizes. The text element automatically adjusts its font size within this range to fit the text content.

## In USS

In USS, set the `-unity-text-auto-size` property to `best fit` with a minimum and maximum font size. For example:

```
.labelText {
    -unity-text-generator: advanced; 
    -unity-text-auto-size: best-fit 5px 100px;
}
```

## In C# scripts

To enable auto sizing in C# scripts, use the [`TextAutoSizeMode`](../../ScriptReference/UIElements.TextAutoSizeMode.md) property. For example:

```
textElement.style.unityTextAutoSize = new StyleTextAutoSize(new TextAutoSize(TextAutoSizeMode.BestFit, minSize: 10, maxSize: 100));
```

## Additional resources

* [`TextAutoSizeMode`](../../ScriptReference/UIElements.TextAutoSizeMode.md)