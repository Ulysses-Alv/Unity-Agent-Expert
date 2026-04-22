# Bind ListView to a list with runtime binding

**Version**: 6000.0+

The ListView control is the most efficient way to create lists. This example demonstrates how to use runtime binding to bind a ListView to a list without C# code.

## Example overview

This example creates a `GameSwitchListAsset` ScriptableObject that contains a list of `GameSwitch` objects. Each `GameSwitch` object has two properties: `enabled` and `name`. The ListView displays a list of toggles and text fields that are bound to the `enabled` and `name` properties of each `GameSwitch` object.

![A list of custom switches, displayed in the Game view.](../uploads/Main/uitk/runtime-binding-listview.png)

A list of custom switches, displayed in the Game view.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/runtime-binding-listview).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](UIBuilder.md)
* [UXML](UIE-UXML.md)
* [Runtime data binding](UIE-runtime-binding.md)
* [UXML element ListView](UIE-uxml-element-ListView.md)

## Create an object with a list

Create the data structure that holds your list items.

1. Create a Unity project with any template.
2. In your **Project window**, create a C# script named `GameSwitchListAsset.cs` with the following contents:

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

## Create the ListView item template in UI Builder

Create a reusable template that defines how each list item look and behave. ListView uses this template to generate **visual elements** for each data item. The template approach is efficient because:

* One template definition works for any number of list items
* UI Toolkit can recycle visual elements as users scroll
* Changes to the template automatically apply to all items

Each item contains a Toggle for the enabled state and a TextField for the name. You bind these UI elements to the corresponding properties of the `GameSwitch` object.

1. Create a UXML file named `ListViewItem.uxml`.
2. Double-click the `ListViewItem.uxml` file to open it in UI Builder.
3. In the **Hierarchy** panel, add a **VisualElement**.
4. Add a **Toggle** as a child of **VisualElement**.
5. Add a **TextField** as a child **VisualElement**.
6. Remove the label text for the **Toggle** and **TextField**.
7. Set the **Flex direction** of the **VisualElement** to `Row`.
8. In the ****Inspector**** panel of the **Toggle**, do the following:
9. Right-click **Value** and select **Add binding**.
10. In the **Add Binding** window, set **Data Source Path** to `enabled`.
11. Set **Binding Mode** to `Two Way`. This ensures changes in the Toggle update the data source, and changes in the data source update the Toggle.
12. In the **Inspector** panel of the **TextField**, do the following:

    * Right-click **Value** and select **Add binding**.
    * In the **Add Binding** window, set **Data Source Path** to `name`.
    * Set **Binding Mode** to `Two Way`. This ensures changes in the TextField update the data source, and changes in the data source update the TextField.
13. Save the UXML file. The finished UXML file looks like the following:

```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" editor-extension-mode="False">
       <ui:VisualElement name="" style="flex-direction: row;">
           <ui:Toggle>
               <Bindings>
                   <ui:DataBinding property="value" data-source-path="enabled" binding-mode="TwoWay"/>
               </Bindings>
           </ui:Toggle>
           <ui:TextField placeholder-text="filler text">
               <Bindings>
                   <ui:DataBinding property="value" data-source-path="name" binding-mode="TwoWay"/>
               </Bindings>
           </ui:TextField>
       </ui:VisualElement>
   </ui:UXML>
```

## Create the ListView UI in UI Builder

Create the asset that contains a list of items and create the main UI that hosts the ListView. This step configures how the ListView connects to your data and defines the user interaction capabilities.

1. From the menu, select **Assets** > **Create** > **UIToolkitExamples** > **GameSwitchList**. This creates an asset named **GameSwitchListAsset**.
2. Create a UXML file named `UIListView.uxml`.
3. Double-click the `UIListView.uxml` file to open it in UI Builder.
4. In the **Hierarchy** panel, add a **ListView**.
5. In the **Inspector** panel of the **ListView**, configure the following settings:
6. **Data Source**: Select **GameSwitchListAsset.asset**. This establishes the type of object the ListView expects to receive as its data source.
7. **Binding Source Selection Mode**: Set to **Auto Assign**. This tells the binding system to automatically determine the correct data context for each list item. When rendering item templates, each template receives the corresponding `GameSwitch` object as its binding source.
8. **Item Template**: Set to `ListViewItem.uxml`.
9. **Show Add Remove Footers**: Enable this option to allow users to add or remove items from the list.
10. **Reorder Mode**: Set to **Animated** to enable drag-and-drop reordering.
11. Save the UXML file.

## Set the `itemsSource` binding in UXML

The final crucial step is connecting the ListView’s [`itemsSource`](../ScriptReference/UIElements.BaseVerticalCollectionView-itemsSource.md) property to your data. The `itemsSource` determines what collection of objects the ListView displays. In this case, the `switches` list from your `GameSwitchListAsset`.

This property isn’t exposed in UI Builder’s Inspector. You need to manually edit the UXML to add this binding. The `itemsSource` binding is what makes the ListView populate with your actual data items.

To create this binding, add a `<Bindings>` tag inside `<ui:ListView>`:

```
<ui:ListView>
   <Bindings>
      <ui:DataBinding property="itemsSource" data-source-path="switches" />
   </Bindings>
</ui:ListView>
```

**Note**: This example sets up the `itemsSource` in UXML for simplicity. You can also set up the data source connection in your C# script like this:

```
var listView = rootVisualElement.Q<ListView>();
var gameSwitchListAsset = /* obtain reference to your GameSwitchListAsset instance */;
listView.dataSource = gameSwitchListAsset;
```

The finished `UIListView.uxml` file looks like the following:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" editor-extension-mode="False">
    <ui:ListView binding-source-selection-mode="AutoAssign" item-template="ListViewItem.uxml" data-source="GameSwitchListAsset.asset" show-foldout-header="true" virtualization-method="DynamicHeight" show-border="false" reorderable="true" horizontal-scrolling="false" selection-type="Single" reorder-mode="Animated" show-add-remove-footer="true">
        <Bindings>
            <ui:DataBinding property="itemsSource" data-source-path="switches"/>
        </Bindings>
    </ui:ListView>
</ui:UXML>
```

**Note**: In your UXML file, make sure to remove the `engine="UnityEngine.UIElements"` and `editor="UnityEditor.UIElements"` attributes from the `<ui:UXML>` tag if you have them. These attributes set the default namespace, which breaks the `<Bindings>` tag because `Bindings` isn’t a type, but a field.

## Test the binding in a scene

Test the complete system to check how runtime binding creates a dynamic, interactive interface.

1. In a sample scene, select **GameObject** > **UI Toolkit** > **UI Document** to add a UI Document.
2. In the **Inspector** panel of the **UI Document** component, set the **Source Asset** to `UIListView.uxml`.
3. Enter Play mode. The ListView displays the list of `GameSwitch` objects from the `GameSwitchListAsset`. You can interact with the toggles and text fields, add or remove items, and reorder the list. All changes reflect in the underlying data source, and vice versa.

## Additional resources

* [`ListView`](../ScriptReference/UIElements.ListView.md)
* [Bind to a list](UIE-bind-to-list.md)
* [Create a custom Editor window with C# script](UIE-HowTo-CreateEditorWindow.md)
* [Create custom binding to bind USS](UIE-runtime-custom-binding-bind-uss.md)