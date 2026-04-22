# Get started with UI Toolkit

Want to create your first UI with UI Toolkit? Use this basic UI Toolkit workflow example to get started.

**Note**: For demonstration purposes, this guide describes how to add UI controls for the Editor UI. However, the instructions on adding UI controls to a UI Document also apply to runtime UI. For more information, refer to [Get started with runtime UI](UIE-get-started-with-runtime-ui.md).

If you perform a specific task often, you can use UI Toolkit to create a dedicated UI for it. For example, you can create a custom Editor window. The example demonstrates how to create a custom Editor window and add UI controls into your custom Editor window with [UI Builder](UIBuilder.md), [UXML](UIE-UXML.md), and C# script.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/simple-ui-toolkit-workflow).

## Create a custom Editor window

Create a custom Editor window with two labels.

1. Create a project in Unity Editor with any template.
2. In the **Project** window, right-click in the `Assets` folder, and then select **Create** > **UI Toolkit** > **Editor Window**.
3. In **UI Toolkit Editor Window Creator**, enter `SimpleCustomEditor` in the **C**# box.
4. Keep the **UXML** checkbox selected and clear the **USS** checkbox.
5. Select **Confirm**.
6. To open the Editor window, select **Window** > **UI Toolkit** > **SimpleCustomEditor**.

You can find the source files for it in the `Assets/Editor` folder.

## Add UI controls to the window

You can add UI controls to your window in the following ways:

* [Use the UI Builder to visually add the UI controls](#use-ui-builder)
* [Use an XML-like text file (UXML) to add the UI controls](#use-uxml)
* [Use C# script to add the UI controls](#use-c#-script)

You can use any of these methods individually, or combine. The following examples create three sets of labels, buttons, and toggles with a combination of these methods.

### Use UI Builder to add UI controls

To visually add UI controls to your window, use [UI Builder](UIBuilder.md). The following steps add a button and a toggle into your custom Editor window in addition to the default label.

1. In the `Editor` folder, double-click `SimpleCustomEditor.uxml` to open the UI Builder.
2. In the UI Builder, drag **Button** and **Toggle** from **Library** > **Controls** into the **Hierarchy** or the window preview in the **Viewport**.
3. In the **Hierarchy** window, select **Label**.
4. In the ****Inspector**** window, change the default text to `These controls were created in UI Builder` in the **Text** field.
5. In the **Hierarchy** window, select **Button**.
6. In the **Inspector** window, enter `This is button1` in the **Text** field.
7. Enter `button1` in the **Name** field.
8. In the **Hierarchy** window, select **Toggle**.
9. In the **Inspector** window, enter `Number?` in the **Label** field.
10. Enter `toggle1` in the **Name** field.
11. [Save](UIB-interface-overview.md#open-and-save-ui-documents-uxml) and close the UI Builder window.
12. Close your custom Editor window if you haven’t done so.
13. In the **Project** window, select `SimpleCustomEditor.cs` and make sure the **Visual Tree Asset** is set to `SimpleCustomEditor (Visual Tree Asset)` in the **Inspector** window.
14. Select **Window** > **UI Toolkit** > **SimpleCustomEditor** to re-open your custom Editor window to see the button and the toggle you just added.

![Custom Editor Window with one set UI Controls](../uploads/Main/UIBuilder/CustomEditorWithUIControls1.png)

Custom Editor Window with one set UI Controls

### Use UXML to add UI controls

If you prefer to define your UI in a text file, you can edit the [UXML](UIE-UXML.md) to add the UI controls. The following steps add another set of label, button, and toggle into your window.

1. In the `Editor` folder, select **Assets** > **Create** > **UI Toolkit** > **UI Document** to create a UXML file called `SimpleCustomEditor_UXML.uxml`.
2. Select the arrow on `SimpleCustomEditor_UXML.uxml` in the **Project** window.
3. Double-click `inlineStyle` to open `SimpleCustomEditor_UXML.uxml` in a text editor.
4. Replace the contents of `SimpleCustomEditor_uxml.uxml` with the following:

```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements" xsi="http://www.w3.org/2001/XMLSchema-instance" engine="UnityEngine.UIElements" editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
       <ui:Label text="These controls were created with UXML." />
       <ui:Button text="This is button2" name="button2"/>
       <ui:Toggle label="Number?" name="toggle2"/>
   </ui:UXML>
```

1. Open `SimpleCustomEditor.cs`.
2. Import the UXML file you created manually by adding the following to the `CreateGUI` method:

   ```
   // Import UXML created manually.
   var visualTree = AssetDatabase.LoadAssetAtPath<VisualTreeAsset>("Assets/Editor/SimpleCustomEditor_uxml.uxml");
   VisualElement labelFromUXML = visualTree.Instantiate();
   root.Add(labelFromUXML);
   ```
3. Save your changes.
4. Select **Window** > **UI Toolkit** > **SimpleCustomEditor**. This opens your custom Editor window with three labels, two buttons, and two toggles.

![Custom Editor Window with two sets UI Controls](../uploads/Main/UIBuilder/CustomEditorWithUIControls2.png)

Custom Editor Window with two sets UI Controls

### Use C# script to add UI controls

If you prefer coding, you can add UI Controls to your window with a C# script. The following steps add another set of label, button, and toggle into your window.

1. Open `SimpleCustomEditor.cs`.
2. Unity uses `UnityEngine.UIElements` for basic UI controls like label, button, and toggle. To work with UI controls, you must add the following declaration if it’s not already present.

   ```
   using UnityEngine.UIElements;
   ```
3. Change the text of the existing label from `"Hello World! From C#"` to `"These controls were created using C# code."`.
4. The [EditorWindow](../ScriptReference/EditorWindow.md) class has a property called `rootVisualElement`. To add the UI controls to your window, first instantiate the element class with some attributes, and then use the `Add` methods of the `rootVisualElement`.

   Your finished `CreateGUI()` method looks like the following:

   ```
   public void CreateGUI()
   {
       // Each editor window contains a root VisualElement object
       VisualElement root = rootVisualElement;

       // VisualElements objects can contain other VisualElements following a tree hierarchy.
       Label label = new Label("These controls were created using C# code.");
       root.Add(label);

       Button button = new Button();
       button.name = "button3";
       button.text = "This is button3.";
       root.Add(button);

       Toggle toggle = new Toggle();
       toggle.name = "toggle3";
       toggle.label = "Number?";
       root.Add(toggle);

       // Instantiate UXML created automatically which is set as the default VisualTreeAsset.
       root.Add(m_VisualTreeAsset.Instantiate());

       // Import UXML created manually.
       var visualTree = AssetDatabase.LoadAssetAtPath<VisualTreeAsset>("Assets/Editor/SimpleCustomEditor_uxml.uxml");
       VisualElement labelFromUXML = visualTree.Instantiate();
       root.Add(labelFromUXML);
   }
   ```
5. Select **Window** > **UI Toolkit** > **SimpleCustomEditor** to open your custom Editor window to see three labels, three buttons, and three toggles.

![Custom Editor Window with three Controls](../uploads/Main/UIBuilder/CustomEditorWithUIControls.png)

Custom Editor Window with three Controls

## Define the behavior of your UI controls

You can set up event handlers for your UI controls so that when you click the button, and select or clear the toggle, your UI controls perform some tasks.

In this example, set up event handlers that:

* When a button is clicked, the Editor Console displays a message.
* When a toggle is selected, the Console shows how many times the buttons have been clicked.

Your finished `SimpleCustomEditor.cs` looks like the following:

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;

public class SimpleCustomEditor : EditorWindow
{
    [SerializeField]
    private VisualTreeAsset m_VisualTreeAsset = default;

    [MenuItem("Window/UI Toolkit/SimpleCustomEditor")]
    public static void ShowExample()
    {
        SimpleCustomEditor wnd = GetWindow<SimpleCustomEditor>();
        wnd.titleContent = new GUIContent("SimpleCustomEditor");
    }

    private int m_ClickCount = 0;

    private const string m_ButtonPrefix = "button";

    public void CreateGUI()
    {
        // Each editor window contains a root VisualElement object
        VisualElement root = rootVisualElement;

        // VisualElements objects can contain other VisualElement following a tree hierarchy.
        Label label = new Label("These controls were created using C# code.");
        root.Add(label);

        Button button = new Button();
        button.name = "button3";
        button.text = "This is button3.";
        root.Add(button);

        Toggle toggle = new Toggle();
        toggle.name = "toggle3";
        toggle.label = "Number?";
        root.Add(toggle);

        // Instantiate UXML created automatically which is set as the default VisualTreeAsset.
        root.Add(m_VisualTreeAsset.Instantiate());

        // Import UXML created manually.
        var visualTree = AssetDatabase.LoadAssetAtPath<VisualTreeAsset>("Assets/Editor/SimpleCustomEditor_uxml.uxml");
        VisualElement labelFromUXML = visualTree.Instantiate();
        root.Add(labelFromUXML);

        //Call the event handler
        SetupButtonHandler();
    }

    //Functions as the event handlers for your button click and number counts
    private void SetupButtonHandler()
    {
        VisualElement root = rootVisualElement;

        var buttons = root.Query<Button>();
        buttons.ForEach(RegisterHandler);
    }

    private void RegisterHandler(Button button)
    {
        button.RegisterCallback<ClickEvent>(PrintClickMessage);
    }

    private void PrintClickMessage(ClickEvent evt)
    {
        VisualElement root = rootVisualElement;

        ++m_ClickCount;

        //Because of the names we gave the buttons and toggles, we can use the
        //button name to find the toggle name.
        Button button = evt.currentTarget as Button;
        string buttonNumber = button.name.Substring(m_ButtonPrefix.Length);
        string toggleName = "toggle" + buttonNumber;
        Toggle toggle = root.Q<Toggle>(toggleName);

        Debug.Log("Button was clicked!" +
            (toggle.value ? " Count: " + m_ClickCount : ""));
    }
}
```

## Test the example

* From the menu, select **Window** > **UI Toolkit** > **SimpleCustomEditor**.

## Additional resources

* [UXML](UIE-UXML.md)
* [Visual Tree](UIE-VisualTree.md)
* [Label](UIE-uxml-element-Label.md)
* [Button](UIE-uxml-element-Button.md)
* [Toggle](UIE-uxml-element-Toggle.md)
* [Create a custom Editor window](UIE-HowTo-CreateEditorWindow.md)