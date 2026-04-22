# Instantiate UXML from C# scripts

To build UI from a UXML file:

1. [Load the file into a `VisualTreeAsset`](UIE-manage-asset-reference.md),.
2. Use either:
   * [`Instantiate()`](../ScriptReference/UIElements.VisualTreeAsset.Instantiate.md) to instantiate without a parent, which creates a new `TemplateContainer`.
   * [`CloneTree(parent)`](../ScriptReference/UIElements.VisualTreeAsset.CloneTree.md) to clone inside a parent.

Once the UXML is instantiated, you can retrieve specific elements from the **visual tree** with [UQuery](UIE-UQuery.md).

The following example creates a custom Editor window and loads a UXML file as its content:

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;
using UnityEditor.UIElements;

public class MyWindow : EditorWindow  {
    [MenuItem ("Window/My Window")]
    public static void  ShowWindow () {
        EditorWindow w = EditorWindow.GetWindow(typeof(MyWindow));

        VisualTreeAsset uiAsset = AssetDatabase.LoadAssetAtPath<VisualTreeAsset>("Assets/MyWindow.uxml");
        VisualElement ui = uiAsset.Instantiate();

        w.rootVisualElement.Add(ui);
    }
}
```

To load UXML assets for runtime, [set up `VisualTreeAsset` references in your MonoBehaviour scripts](UIE-manage-asset-reference.md) and assign the UXML assets from the **Inspector**. For more information and an example, refer to [Support for runtime UI](UIE-support-for-runtime-ui.md) and [Create a list view runtime UI](UIE-HowTo-CreateRuntimeUI.md).

## Additional resources

* [Load UXML and USS from C# scripts](UIE-manage-asset-reference.md)
* [UQuery](UIE-UQuery.md)