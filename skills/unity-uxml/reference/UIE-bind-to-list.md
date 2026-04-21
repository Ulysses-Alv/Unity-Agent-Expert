# Bind to a list with ListView

**Version**: 2021.3+

The ListView control is the most efficient way to create lists. To bind to a list with ListView, set the binding path of the ListView to the name of the property that contains the list.

This example demonstrates how to bind to a list with ListView.

## Example overview

The example creates a list of toggles and binds the list to an underlying list of `GameSwitch` objects.

![A list of custom switches, displayed in the Unity Editor.](../uploads/Main/uie_binding_listview.png)

A list of custom switches, displayed in the Unity Editor.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/bind-to-list).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [USS](UIE-USS.md)
* [`ListView`](../ScriptReference/UIElements.ListView.md)

## Create an object with a list

Create a `GameSwitch` object and a serialized object that has a list of `GameSwitch` objects as a property.

1. Create a Unity project with any template.
2. In your **Project window**, create a folder named `bind-to-list` to store all your files.
3. Create a C# script named `GameSwitchListAsset.cs` and replace its contents with the following:

```
using System;
using System.Collections.Generic;
using System.Diagnostics;
using UnityEngine;
  
[CreateAssetMenu(fileName = "GameSwitchListAsset.asset", menuName = "GameSwitchListAsset")]
public class GameSwitchListAsset : ScriptableObject
{
    public List<GameSwitch> switches = new();

    public void Reset()
    {
        switches = new List<GameSwitch>{
            new() { name = "Use Local Server", enabled = false },
            new() { name = "Show Debug Menu", enabled = false },
            new() { name = "Show FPS Counter", enabled = true },
        };
    }

    public bool IsSwitchEnabled(string switchName) => switches.Find(s => s.name == switchName).enabled;

    [Serializable]
    public struct GameSwitch
    {
        public bool enabled;
        public string name;
    }
}
```

## Create the UXML template for the list items and ListView UI

Create the template of the list item. Each item contains a Toggle and a TextField. Create the ListView UI, and set the item template to the template of the list item.

1. Create a UXML file named `game_switch.uxml` with the following contents:

   ```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" editor-extension-mode="False">
       <ui:Box style="flex-direction: row;">
           <ui:Toggle binding-path="enabled"/>
           <ui:TextField binding-path="name" readonly="false" style="flex-grow: 1;"/>
       </ui:Box>
   </ui:UXML>
   ```
2. Create a UXML file named `game_switch_list_editor.uxml` and with the following contents:

   ```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" editor-extension-mode="False">
       <ui:ListView virtualization-method="DynamicHeight" reorder-mode="Simple" binding-path="switches" show-add-remove-footer="true" show-foldout-header="true" header-title="" reorderable="true" item-template="game_switch.uxml" allow-add="true">
       </ui:ListView>
   </ui:UXML>
   ```

## Create a custom Editor window and set the binding

Create a custom Editor window that can create an asset with a list of toggles. Bind the toggle list to the `GameSwitch` list by setting the `binding-path` attribute of the UI Document to the property name of the `GameSwitch` list, which is `switches`.

1. From the menu, select **Assets** > **Create** > **UIToolkitExamples** > **GameSwitchList**. This creates an asset named **New Game Switch List Asset**.
2. Create a folder named `Editor`.
3. In the **Editor** folder, create a C# script named `GameSwitchListEditor.cs` with the following contents:

   ```
   using System.Collections.Generic;
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;
   using UnityEditor.UIElements;
       
   public class ListViewBindingExample : EditorWindow
   {
       [SerializeField]
       private VisualTreeAsset m_VisualTreeAsset = default;
       
       [SerializeField]
       private GameSwitchListAsset gameSwitchList;
       
       [MenuItem("UI Toolkit Examples/ListView SerializedObject Binding Example")]
       public static void ShowExample()
       {
           ListViewBindingExample wnd = GetWindow<ListViewBindingExample>();
           wnd.titleContent = new GUIContent("ListView Binding SerializedObject Example");
       }
       
       public void CreateGUI()
       {
           VisualElement root = rootVisualElement;
           m_VisualTreeAsset.CloneTree(root);
       
           var listView = root.Q<ListView>();
           if (listView != null && gameSwitchList != null)
           {
               // Set the items source.
               listView.itemsSource = gameSwitchList.switches;
               // Bind the ListView to the GameSwitchListAsset serialized object.
               listView.Bind(new SerializedObject(gameSwitchList));
           }
       }
   }
   ```
4. Select `GameSwitchListEditor.cs` in the Project window.
5. In the ****Inspector**** panel, do the following:

   * Set **UXML Asset** to `game_switch_list_editor.uxml`
   * Set **Game Switch List Asset** to `GameSwitchListAsset.asset` .

## Test the binding

1. From the menu, select **UI Toolkit Examples** > **ListView SerializedObject Binding Example** to open the custom Editor window.
2. In the Project window, select **GameSwitchListAsset**. This shows a list of toggles in the Inspector. You can reorder the list, collapse it, add entries to or remove entries from the list, and change the number of entries in the list. If you make changes in the Custom Editor window, the `switches` property of the `GameSwitchListAsset` object changes, and vice versa.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)