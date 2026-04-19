# Add hyperlinks in text

You can use rich text tags to add hyperlinks to text for the Editor UI or at runtime. For demonstration purposes, this example shows how to add hyperlinks in a custom Editor window.

## Example overview

This example creates a custom Editor window with three hyperlinks. The first two links are created with `<link>` tags that use link IDs. The third link uses an `<a>` tag with an `href` attribute to define the destination.

The third link is created with an `<a>` tag that uses the `href` attribute to define the link’s destination.

The `<link>` tag is a rich text tag that you can use to create links in UI Toolkit. Use the `<link>` tag if you need to create many links or want to define custom link behavior with C# scripts.

This example uses the following events to define the `<link>` tag logic:

**Note**: Those events are experimental, so they’re still in the process of becoming stable enough to release.

* [PointerDownLinkTagEvent](../../ScriptReference/UIElements.Experimental.PointerDownLinkTagEvent.md)
* [PointerUpLinkTagEvent](../../ScriptReference/UIElements.Experimental.PointerUpLinkTagEvent.md)
* [PointerMoveLinkTagEvent](../../ScriptReference/UIElements.Experimental.PointerMoveLinkTagEvent.md)
* [PointerOutLinkTagEvent](../../ScriptReference/UIElements.Experimental.PointerOutLinkTagEvent.md)
* [PointerOverLinkTagEvent](../../ScriptReference/UIElements.Experimental.PointerOverLinkTagEvent.md)

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/link-tag-example).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](../UIBuilder.md)
* [UXML](../UIE-UXML.md)
* [USS](../UIE-USS.md)
* [Rich text tags](../UIE-rich-text-tags.md)

## Create the example

To create a custom Editor window with hyperlinks with UI Toolkit, follow these steps:

1. Create a project in Unity with any template.
2. Right-click in the `Assets` folder, and select **Create** > **UI Toolkit** > **Editor Window**.
3. In **UI Toolkit Editor Window Creator**, enter `LinkTag` in the **C#** field.
4. Select **Confirm**. This creates an `Editor` folder with the C# script, UXML, and USS files for your custom window.
5. Double-click `LinkTag.uxml` to open it in the UI Builder.
6. In the UI Builder, select the first label.
7. Replace the **Text** field with the following content:

   ```
   Link to <link="1"><color=#40a0ff><u>Unity</u></color></link>
   Link to <link="2"><color=#40a0ff><u>UITK Discussions</u></color></link>!
   ```
8. Enable **Enable Rich Text**.
9. Select **Add Style Class to List**.
10. Enter `link` and press the **Enter** key. This adds a `link` class to the first label.
11. Select the second label, and replace the **Text** field with the following content:

    ```
    Link to <a href="https://www.google.com/"><u>Google</u></a> using <noparse><a></noparse> tag.
    ```
12. Enable **Enable Rich Text**.
13. Save and close the UI Builder.
14. Replace the content of `LinkTag.uss` with the following:

    ```
    Label {
        font-size: 75px;
    }
        
    /*Turn the mouse pointer into a hand pointer when you hover over the link text.
      Note that the cursor override for the link value doesn’t work at runtime.*/
        
    .link-cursor {
        cursor: link;
    }
    ```
15. Replace the content of `LinkTag.cs` with the following:

    ```
    using System.Collections.Generic;
    using UnityEditor;
    using UnityEngine;
    using UnityEngine.UIElements;
    using UnityEngine.UIElements.Experimental;
        
    public class LinkTag : EditorWindow
    {
        [SerializeField]
        private VisualTreeAsset m_VisualTreeAsset = default;
        
        [MenuItem("Window/LinkTag Sample")]
        public static void ShowExample()
        {
            LinkTag wnd = GetWindow<LinkTag>();
            wnd.titleContent = new GUIContent("LinkTag Sample");
        }
        
        readonly string linkCursorClassName = "link-cursor";
        Dictionary<int, string> m_UrlLookup;
        Label linkLabel;
        
        public void OnEnable()
        {
            // Use the dictionary to map link IDs to URLs.
            // The link ID is the number in the link tag, and the URL is the destination.
            // This is a simple example, in a real application you can define a more complex structure.
            m_UrlLookup = new Dictionary<int, string>()
            {
                { 1, "https://www.google.com/" },
                { 2, "https://discussions.unity.com/" }
            };
        }
        
        public void CreateGUI()
        {
            VisualElement uxml = m_VisualTreeAsset.Instantiate();
            rootVisualElement.Add(uxml);
        
            linkLabel = rootVisualElement.Q<Label>(className: "link");
        
            linkLabel.RegisterCallback<PointerDownLinkTagEvent>(HyperlinkOnPointerDown);
            linkLabel.RegisterCallback<PointerUpLinkTagEvent>(HyperlinkOnPointerUp);
            linkLabel.RegisterCallback<PointerMoveLinkTagEvent>(HyperlinkPointerMove);
            linkLabel.RegisterCallback<PointerOverLinkTagEvent>(HyperlinkOnPointerOver);
            linkLabel.RegisterCallback<PointerOutLinkTagEvent>(HyperlinkOnPointerOut);
        }
        
        void HyperlinkOnPointerOver(PointerOverLinkTagEvent _)
        {
            linkLabel.AddToClassList(linkCursorClassName);
        }
        void HyperlinkPointerMove(PointerMoveLinkTagEvent _) { }
        void HyperlinkOnPointerOut(PointerOutLinkTagEvent _)
        {
            linkLabel.RemoveFromClassList(linkCursorClassName);
        }
        
        void HyperlinkOnPointerDown(PointerDownLinkTagEvent _) { }
        
        void HyperlinkOnPointerUp(PointerUpLinkTagEvent evt)
        {
            var linkID = int.Parse(evt.linkID);
            if (m_UrlLookup.TryGetValue(linkID, out var url))
                Application.OpenURL(url);
        }
        
    }
    ```

## Test the link

To verify that the link works, follow these steps:

1. From the manual, select **Window** > **LinkTag Sample**.
2. Select the links in the custom window to open the linked websites.

## Additional resources

* [Style text with rich text tags](../UIE-rich-text-tags.md)
* [Supported rich text tags](../UIE-supported-tags.md)