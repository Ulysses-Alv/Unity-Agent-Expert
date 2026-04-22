# Control textures of the dynamic atlas

To reduce the number of batches broken by texture changes and achieve good performance, use an atlas to group together textures used at the same time. You can achieve this with either of the following methods:

* Use a [sprite atlas](sprite/atlas/atlas-landing.md). With this method, you have more control over the sprites but you need to manually create the **sprite** atlas.
* Use a dynamic atlas. UI Toolkit automatically adds textures to, or removes textures from, a dynamic atlas when **visual elements** reference them.

## Verify textures in dynamic atlas

When you use dynamic atlas to group together textures, to limit the number of [draw calls](DrawCallBatching.md), make sure that your textures enter the dynamic atlas. To verify, use the [Frame Debugger](FrameDebugger.md). Frame Debugger helps you to observe texture changes and deduce batch breaking.

The following example verifies that the dynamic atlas contains all the textures in runtime UI.

![The example Dynamic Atlas window contains all the textures in runtime UI.](../uploads/Main/uxml/dynamic-atlas.png)

The example Dynamic Atlas window contains all the textures in runtime UI.

## Apply built-in filters

The dynamic atlas texture starts from a specified minimum size, and grows as needed, doubling horizontally or vertically, up to a specified maximum size. You can define the [minimum and maximum atlas sizes](UIE-Runtime-Panel-Settings.md#dynamic-atlas-settings) in the Panel Settings asset. You can also use the filters in the dynamic atlas to decide whether to add a sub-texture to the atlas.

To enable or disable the filters, in the Panel Settings asset’s Inspector window, select the options from **Dynamic Atlas Settings** > [**Active Filters**](UIE-Runtime-Panel-Settings.md#active-filters) dropdown list.

## Use custom filters

You can assign a custom filter to [`PanelSettings.dynamicAtlasSettings.customFilter`](../ScriptReference/UIElements.DynamicAtlasSettings-customFilter.md) to add or relax constraints on a global or per-texture basis.

The following custom filter example bypasses a large texture from the [**Size**](UIE-Runtime-Panel-Settings.md#active-filters) filter while keeping the [**Size**](UIE-Runtime-Panel-Settings.md#active-filters) filter active for other textures:

```
using UnityEngine;
using UnityEngine.UIElements;

class MyCustomFilter : MonoBehaviour
{
    public PanelSettings panelSettings;
    public Texture2D largeTexture;

    void OnEnable() { panelSettings.dynamicAtlasSettings.customFilter = Filter; }

    void OnDisable() { panelSettings.dynamicAtlasSettings.customFilter = null; }

    bool Filter(Texture2D texture, ref DynamicAtlasFilters filtersToApply)
    {
        if (texture == largeTexture)
        {
            // Disable the Size check for this one.
            filtersToApply &= ~DynamicAtlasFilters.Size;
        }
        return true;
    }
}
```

## Manage dynamic atlas size

When new elements appear in the UI, their textures are added to the dynamic atlas until the atlas reaches or exceeds the Max Atlas Size specified in the Dynamic Atlas Settings.

On devices with limited memory, you might need to set a lower Max Atlas Size than the default, such as `2048` **pixels** instead of `4096` pixels.

## Manage dynamic atlas fragmentation

When textures are added or removed from the atlas, it can lead to fragmentation, creating small spaces where previous textures were, which are too small to reallocate to other textures.

To maximize the usage of the atlas and avoid fragmentation, you need to manually reset its contents by calling the [`RuntimePanelUtils.ResetDynamicAtlas()`](../ScriptReference/UIElements.RuntimePanelUtils.ResetDynamicAtlas.md) extension method, which requires accessing the `panel` property of a `VisualElement`.

Consider resetting the dynamic atlas based on your application’s behaviour, for example when you remove or add large amounts of visual elements.

The following example shows how to use the [`RuntimePanelUtils.ResetDynamicAtlas()`](../ScriptReference/UIElements.RuntimePanelUtils.ResetDynamicAtlas.md) extension method:

```
using UnityEngine;
using UnityEngine.UIElements;

[RequireComponent(typeof(UIDocument))]
class DynamicAtlasResetExample : MonoBehaviour
{
    VisualElement rootVisualElement => GetComponent<UIDocument>().rootVisualElement;

    public void OnEnable()
    {
        rootVisualElement.Add(new Button(ResetAtlas) { text = "Reset Dynamic Atlas"});
    }

    private void ResetAtlas()
    {
        rootVisualElement.panel.ResetDynamicAtlas();
    }

}
```

## Additional resources

* [Panel Settings asset](UIE-Runtime-Panel-Settings.md)
* [Frame Debugger](FrameDebugger.md)
* [Draw call batching](DrawCallBatching.md)
* [Sprite Atlas](sprite/atlas/atlas-landing.md)