# Command events

Command events are sent to allow the Unity Editor to forward top-level menu actions to the Editor UI, as well as their equivalent keyboard shortcuts.

The following are the available commands:

* `Copy`
* `Cut`
* `Paste`
* `Delete`
* `DeselectAll`
* `SoftDelete`
* `Duplicate`
* `FrameSelected`
* `FrameSelectedWithLock`
* `SelectAll`
* `Find`
* `FocusProjectWindow`

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [ValidateCommandEvent](../ScriptReference/UIElements.ValidateCommandEvent.md) | The Editor sends this event when determining whether an element in the panel handles the command. | ✔ | ✔ | ✔ |
| [ExecuteCommandEvent](../ScriptReference/UIElements.ExecuteCommandEvent.md) | The Editor sends this event when an element in the panel executes a command. | ✔ | ✔ | ✔ |

## Unique properties

**`target`**: The element with keyboard focus. This value is `null` if no element has focus.

**`commandName`**: The command to validate or execute.

## Event List

### ValidateCommandEvent

The `ValidateCommandEvent` event asks an EditorWindow if it can execute a command. For example, the Editor can enable or disable a menu item based on the results from the validation command event.

To verify if the Editor can execute a command:

1. Register a callback for `ValidateCommandEvent`.
2. Test the `commandName` property of the event.
3. Call the `Event.StopPropogation()` method on the event if the command can be executed.

### ExecuteCommandEvent

The `ExecuteCommandEvent` event asks an Editor Window to execute a command.

Even if this event follows a validation event, it’s best practice to ensure the action is possible first, independently of any previous validation.

To respond to the command:

1. Register a callback for `ExecuteCommandEvent`.
2. Test the `commandName` property of the event.
3. Call the `Event.StopPropogation()` method on the event before executing the actual logic of the command, so the Editor knows that the comment has been executed.

## Example

The following example uses command events to support copy and paste in a custom Editor Window. The example displays a list of fruits in a custom Editor window. Users can use keyboard shortcuts to copy and paste any fruits.

```
using System;
using System.Collections.Generic;
using System.Linq;
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;

public class CopyPasteExample : EditorWindow
{
    [MenuItem("Window/UI Toolkit Examples/CopyPasteExample")]
    public static void Show()
    {
        GetWindow<CopyPasteExample>();
    }

    readonly List<string> fruits = new ()
    {
        "Banana",
        "Apple",
        "Lime",
        "Orange"
    };
    
    ListView m_ListView;

    public void CreateGUI()
    {
        Func<VisualElement> makeItem = () => new Label();

        Action<VisualElement, int> bindItem = (e, i) => (e as Label).text = fruits[i];

        m_ListView = new ListView();
        m_ListView.makeItem = makeItem;
        m_ListView.bindItem = bindItem;
        m_ListView.itemsSource = fruits;
        m_ListView.selectionType = SelectionType.Single;
        
        m_ListView.RegisterCallback<ValidateCommandEvent>(OnValidateCommand);
        m_ListView.RegisterCallback<ExecuteCommandEvent>(OnExecuteCommand);
        
        rootVisualElement.Add(m_ListView);
    }

    void OnExecuteCommand(ExecuteCommandEvent evt)
    {
        if (evt.commandName == "Copy" && m_ListView.selectedIndices.Count() > 0)
        {
            EditorGUIUtility.systemCopyBuffer = fruits[m_ListView.selectedIndex];
            evt.StopPropagation();
        }
        else if (evt.commandName == "Paste" && !string.IsNullOrEmpty(EditorGUIUtility.systemCopyBuffer))
        {
            fruits.Add(EditorGUIUtility.systemCopyBuffer);
            m_ListView.RefreshItems();
            evt.StopPropagation();
        }
    }

    void OnValidateCommand(ValidateCommandEvent evt)
    {
        if (evt.commandName == "Copy" && m_ListView.selectedIndices.Count() > 0)
        {
            evt.StopPropagation();
        }
        else if (evt.commandName == "Paste" && !string.IsNullOrEmpty(EditorGUIUtility.systemCopyBuffer))
        {
            evt.StopPropagation();
        }
    }
}
```