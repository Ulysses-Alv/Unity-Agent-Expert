# Add an Animation Event

Use an **Animation Event** to call a function at a specific point in time. This function can be in any script attached to the **GameObject** but must only accept a single parameter of type `float`, `int`, `string`, an `object` reference, or an `AnimationEvent` object.

For example, the following script accepts a string. It logs the time and the value of a string parameter when it is called.

```
// An example of C# function that can be called by an Animation Event
using UnityEngine;
using System.Collections;

public class ExampleClass : MonoBehaviour
{
    public void PrintEvent(string s)
    {
        Debug.Log("PrintEvent called at " + Time.time + " with a value of " + s);
    }
}
```

To add an Animation event to a clip at the current playhead location, click the **Event** button. To add an Animation event at any position, right-click the **Event** line where you want to add the Event and select **Add Animation Event** from the context menu. Once added, click and drag an Animation event to reposition it on the Event Line.

![Animation Events display on the Event Line](../uploads/Main/AnimationEditorEventLine.png)

Animation Events display on the **Event Line**

When you add an Event, the **Inspector** window displays the **Function** field. Use this field to select the method you want to call. Note that Animation Events only support methods with a single parameter. You cannot select a function that accepts more than one parameter.

However, you can use an `AnimationEvent` object to pass many parameters at the same time. An `AnimationEvent` object accepts a `float`, an `int`, a `string`, and an `object` reference as member values. The `AnimationEvent` object also provides information about the Animation Event that calls the function.

![The Inspector window with an Animation Event selected. The PrintEvent method is selected from ExampleClass.](../uploads/Main/AnimationEventInspector.png)

The Inspector window with an Animation Event selected. The `PrintEvent` method is selected from `ExampleClass`.

The Events added to a clip display as markers in the Event line. Hover the cursor over a marker to display a tooltip with the function name and parameter value.

![An example of hovering over an Animation Event to display a tooltip.](../uploads/Main/AnimationEditorEventTooltip.png)

An example of hovering over an Animation Event to display a tooltip.

You can select and manipulate multiple Events in the Event Line. To select multiple Events in the Event Line, hold the **Shift** key and click each Event marker one by one. To remove a marker from the selection, hold **Shift** and click a selected marker.

You can also use a selection box to select multiple Animation Events. To do this, click and drag within the Event Line:

![A selection box selects two consecutive Animation Events along the Event Line.](../uploads/Main/AnimationEditorMultipleEventSelection.png)

A selection box selects two consecutive Animation Events along the Event Line.

To delete an Animation Event, select it and press the **Delete** key. You can also right-click the Animation Event and choose **Delete Event** from the context menu.