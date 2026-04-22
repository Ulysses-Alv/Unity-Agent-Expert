# Check pseudo-state of a control

**Version:** 6.3+

You can check the pseudo-state of a control in the UI Toolkit. This is useful for debugging and ensuring that your styles are applied correctly based on the state of the control.

## Example overview

This example demonstrates how to check the pseudo-state of a button and a toggle control.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/pseudo-state-checker).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [Pointer events](../UIE-Pointer-Events.md)
* [Focus events](../UIE-Focus-Events.md)
* [hasHoverPseudoState](../../ScriptReference/UIElements.VisualElement-hasHoverPseudoState.md)
* [hasActivePseudoState](../../ScriptReference/UIElements.VisualElement-hasActivePseudoState.md)
* [hasFocusPseudoState](../../ScriptReference/UIElements.VisualElement-hasFocusPseudoState.md)
* [hasCheckedPseudoState](../../ScriptReference/UIElements.VisualElement-hasCheckedPseudoState.md)
* [hasDisabledPseudoState](../../ScriptReference/UIElements.VisualElement-hasDisabledPseudoState.md)

## Create the UXML file

Create a UXML file to define the UI structure. This file contains a button and a toggle control that you can interact with to check their pseudo states.

1. Create a Unity project with any template.
2. In the **Project window**, create a UXML file named `ButtonAndToggle.uxml`.
3. Double-click the `ButtonAndToggle.uxml` file to open it in the UI Builder.
4. Add a Button to the **Hierarchy** panel and name it as `my-button`.
5. Add a Toggle to the **Hierarchy** panel and name it as `my-toggle`.
6. Save your changes.

The finished `ButtonAndToggle.uxml` file looks like this:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements">
    <ui:Button text="Button" name="my-button"/>
    <ui:Toggle label="Toggle" name="my-toggle"/>
</ui:UXML>
```

## Create the C# script

Create a C# script to check the pseudo-states of the button and toggle control. This script registers callbacks to log the pseudo-states when the user interacts with the controls.

**Note**: Some callbacks occur before the pseudo-state changes, while others occur after. This timing can make the logged results seem unexpected. For example, during [`FocusInEvent`](../../ScriptReference/UIElements.FocusInEvent.md), [hasFocusPseudoState](../../ScriptReference/UIElements.VisualElement-hasFocusPseudoState.md) is still `false`, and during [`FocusOutEvent`](../../ScriptReference/UIElements.FocusOutEvent.md), it is still `true`. The property updates after the event finishes. To get the actual state, check the property in the next Update:

* [hasFocusPseudoState](../../ScriptReference/UIElements.VisualElement-hasFocusPseudoState.md) is `true` only when the element has focus.
* [hasActivePseudoState](../../ScriptReference/UIElements.VisualElement-hasActivePseudoState.md) is `true` only when the pointer is down and over the element.
* [hasHoverPseudoState](../../ScriptReference/UIElements.VisualElement-hasHoverPseudoState.md) is `true` only when the pointer is over the element.

Create a C# script named `PseudoStateChecker.cs` with the following content:

```
using UnityEngine;
using UnityEngine.UIElements;

public class PseudoStateChecker : MonoBehaviour
{
    public UIDocument uiDocument;

    void OnEnable()
    {
        var root = uiDocument.rootVisualElement;

        // Query for the elements. 
        var myButton = root.Q<Button>("my-button");
        var myToggle = root.Q<Toggle>("my-toggle");

        if (myButton == null || myToggle == null)
        {
            Debug.LogError("Button or Toggle not found in the UI Document!");
            return;
        }

        // Register callbacks to check states on a Button. 
        myButton.RegisterCallback<PointerEnterEvent>(evt =>
            Debug.Log($"Button PointerEnterEvent: hasHoverPseudoState = {myButton.hasHoverPseudoState}"));

        myButton.RegisterCallback<PointerLeaveEvent>(evt =>
            Debug.Log($"Button PointerLeaveEvent: hasHoverPseudoState = {myButton.hasHoverPseudoState}"));

        // Use TrickleDown to ensure the callback receives the event before it is stopped by the Button.
        myButton.RegisterCallback<PointerDownEvent>(
            evt => Debug.Log($"Button PointerDownEvent: hasActivePseudoState = {myButton.hasActivePseudoState}"),
            TrickleDown.TrickleDown);

        myButton.RegisterCallback<PointerUpEvent>(evt =>
            Debug.Log($"Button PointerUpEvent: hasActivePseudoState = {myButton.hasActivePseudoState}"));

        myButton.RegisterCallback<FocusInEvent>(evt =>
            Debug.Log($"Button FocusInEvent: hasFocusPseudoState = {myButton.hasFocusPseudoState}"));

        myButton.RegisterCallback<FocusOutEvent>(evt =>
            Debug.Log($"Button FocusOutEvent: hasFocusPseudoState = {myButton.hasFocusPseudoState}"));

        // Register a callback to check the state on a Toggle. 
        myToggle.RegisterValueChangedCallback(evt =>
            Debug.Log($"Toggle ValueChangedEvent: hasCheckedPseudoState = {myToggle.hasCheckedPseudoState}"));

        // Example of checking the disabled state.
        // You can un-comment this to check the disabled state after 3 seconds.
        // Invoke(nameof(DisableTheButton), 3f);
    }

    private void DisableTheButton()
    {
        var myButton = uiDocument.rootVisualElement.Q<Button>("my-button");

        if (myButton != null)
        {
            myButton.SetEnabled(false);
            Debug.Log($"Button has been disabled: hasDisabledPseudoState = {myButton.hasDisabledPseudoState}");
        }
    }
}
```

## Test the example

You can test the example in the SampleScene. The script logs the pseudo-state changes of the button and toggle control to the **Console window** when you interact with them.

1. In the SampleScene, select **+** > **UI Toolkit** > **UI Document** to create a new UI Document GameObject.
2. Assign the `ButtonAndToggle.uxml` file to the **Source Asset** field of the UI Document.
3. Select **Add Component** and add the `PseudoStateChecker` script.
4. Save the **scene**.
5. Enter Play mode.
6. Click the button and toggle the toggle control to check the pseudo-state changes in the Console window.

## Additional resources

* [Pseudo-classes](../UIE-USS-Selectors-Pseudo-Classes.md)