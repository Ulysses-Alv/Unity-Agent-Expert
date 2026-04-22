# Enable and use Advanced Text Generator

Advanced Text Generator (ATG) is not enabled by default, and you need to enable it in your **project settings**.

With ATG, you can use a wide array of languages and **scripts**, such as right-to-left (RTL) languages like Arabic and Hebrew.

## Enable ATG

1. From the menu, select **Edit** > **Project Settings** > **UI Toolkit**.
2. Select the **Enable Advanced Text Generator** checkbox.

## Use ATG

You can use ATG in UI Builder, USS, and C# scripts. It’s best practice to set it in the root element of your UI hierarchy, so that it cascades to all child elements. If you want to use standard text rendering for specific elements, you can override the text generator type for those elements individually.

### In UI Builder

To use Advanced Text Generator in UI Builder, do the following:

1. Select the (root) **visual element** that you want to apply to.
2. In the **Inspector** panel, select **Inlined Styles** > **Text**.
3. From the **Text Generator Type** dropdown, select **Advanced**.

### In USS

To use Advanced Text Generator in USS, set `-unity-text-generator` to `advanced`. The following USS example sets it for the root element:

```
:root {
    -unity-text-generator: advanced;
}
```

### In C# scripts

To use Advanced Text Generator in C# scripts, set [`TextGeneratorType`](../../ScriptReference/TextGeneratorType.md) to `Advanced`. For example:

```
textElement.style.unityTextGenerator = new StyleEnum<TextGeneratorType>(TextGeneratorType.Advanced);
```

## Limitations

ATG has the following limitations:

* Doesn’t support static font assets. If your project uses static font assets, you must [migrate them for ATG](migrate-static-font-assets.md).
* Customization of glyph metrics is not available. The recommended best practice is using font editing tools to adjust metrics or trim the font as needed.