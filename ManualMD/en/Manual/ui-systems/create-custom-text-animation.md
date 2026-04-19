# Create custom text animation

**Version**: 6000.3+

The [`TextElement.PostProcessTextVertices`](../../ScriptReference/UIElements.TextElement.PostProcessTextVertices.md) API lets you modify the **mesh** vertices of each glyph immediately before UI Toolkit renders them. You can use this callback to customize the position, tint, and UV coordinates of the text at a low level.

## Example overview

This example creates a custom Editor window that animates text by fading glyphs in and out when you press **Spacebar**. It uses the `TextElement.PostProcessTextVertices` API to modify the vertex data of the text glyphs.

![Text animation demo](../../uploads/Main/uitk/text-animation-demo.gif)

Text animation demo

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/text-animation-example).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [`TextElement`](../../ScriptReference/UIElements.TextElement.md)

## Create a custom Editor window with text animation

Create a C# script for a custom Editor window. Use a scheduled job to track animation time. In the `OnPostProcessTextVertices` method, iterate through `GlyphsEnumerable` and adjust each vertex tint’s alpha channel until the fade completes.

1. Create a Unity project with any template.
2. Create a folder named `Editor` if you don’t have one.
3. Inside the `Editor` folder, create a C# script named `TextAnimation.cs` with the following content:

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;
using TextElement = UnityEngine.UIElements.TextElement;

public class TextAnimation : EditorWindow
{
    [MenuItem("Window/UI Toolkit/TextAnimation")]
    public static void ShowExample()
    {
        TextAnimation wnd = GetWindow<TextAnimation>();
        wnd.titleContent = new GUIContent("TextAnimation");
    }

    Label label;
    float animationDuration = 1f; // in seconds
    float elapsed = 0f;
    float startRealtime;
    IVisualElementScheduledItem animationJob;
    bool isTextVisible = true;

    public void CreateGUI()
    {
        VisualElement root = rootVisualElement;

        var container = new VisualElement()
        {
            style =
            {
                flexGrow = 1,
                top = 0,
                bottom = 0,
                right = 0,
                left = 0
            },
            focusable = true // We can only receive key events on a window with a focusable element.
        };
        label = new Label("Hello ❤️ World!") { style = { flexGrow = 1, fontSize = 24, unityTextAlign = TextAnchor.MiddleCenter } };
        container.Add(label);
        root.Add(container);

        rootVisualElement.RegisterCallback<KeyDownEvent>(evt => OnSpacebarPressed(evt), TrickleDown.TrickleDown);
        label.PostProcessTextVertices += OnPostProcessTextVertices;

        const int targetHz = 60;
        animationJob = label.schedule.Execute(UpdateTime).Every(1000 / targetHz);
        animationJob.Pause(); // Pause the job until the animation starts
    }

    private void UpdateTime()
    {
        elapsed = Mathf.Min(Time.realtimeSinceStartup - startRealtime, animationDuration);
        if (elapsed >= animationDuration)
        {
            elapsed = animationDuration; // Cap at max duration
            animationJob.Pause();
        }

        label.MarkDirtyRepaint();
    }

    public void OnSpacebarPressed(KeyDownEvent evt)
    {
        if (evt.keyCode != KeyCode.Space || animationJob.isActive)
            return;

        elapsed = 0f;
        startRealtime = Time.realtimeSinceStartup;
        animationJob.Resume();
        isTextVisible = !isTextVisible;
    }

    void OnPostProcessTextVertices(TextElement.GlyphsEnumerable glyphs)
    {
        int glyphsToToggle = (int)(elapsed * glyphs.Count / animationDuration);

        int toggled = 0;
        foreach (TextElement.Glyph glyph in glyphs)
        {
            if (toggled++ >= glyphsToToggle)
                break;

            var verts = glyph.vertices;
            for (int i = 0; i < verts.Length; i++)
            {
                var v = verts[i];
                var tint = v.tint;
                tint.a = isTextVisible ? (byte)255 : (byte)0;
                v.tint = tint;
                verts[i] = v;
            }
        }
    }
}
```

## Test the sample

Test the text animation sample in the custom Editor window.

1. From the menu, select **Window** > **UI Toolkit** > **TextAnimation** to open the sample window.
2. Press the **Spacebar** to start the animation.

## Additional resources

* [Style text with rich text tags](../UIE-rich-text-tags.md)