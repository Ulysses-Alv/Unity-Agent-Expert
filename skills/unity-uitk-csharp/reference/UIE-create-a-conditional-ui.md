# Use Toggle to create a conditional UI

This example uses a Toggle to create a conditional UI.

## Example overview

The example creates a custom Editor window with two toggles. You can use the toggles to do the following:

* Show or hide a label
* Activate or deactivate a button

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/ToggleExample).

## Create the example

To create the example:

1. Create a Unity project with any template.
2. In the **Project window**, create a folder named `Editor`.
3. In the `Editor` folder, create a C# script file named `ToggleExample` with the following content:

   ```
   using UnityEngine;
   using UnityEditor;
   using UnityEngine.UIElements;
   namespace Samples.Editor.Controls
   {
       public class ToggleExample : EditorWindow
       {
           private Toggle showToggle;
           private Toggle activateToggle;
           private Label labelToShow;
           private Button buttonToActivate;
           [MenuItem("Window/ToggleExample")]
           public static void OpenWindow()
           {
               var window = GetWindow<ToggleExample>("Controls: Toggle Sample");
               window.minSize = new Vector2(200, 170);
           }
           public void CreateGUI()
           {
               showToggle = new Toggle("Show label")
               {
                   value = true
               };
               activateToggle = new Toggle("Active button")
               {
                   value = true
               };
               labelToShow = new Label("This label is shown when the above toggle is set to On");
               buttonToActivate = new Button(() => Debug.Log("Button pressed!"))
               {
                   text = "Active if above toggle is On"
               };
               rootVisualElement.Add(showToggle);
               rootVisualElement.Add(labelToShow);
               rootVisualElement.Add(activateToggle);
               rootVisualElement.Add(buttonToActivate);
               showToggle.RegisterValueChangedCallback(evt => labelToShow.visible = evt.newValue);
               activateToggle.RegisterValueChangedCallback(evt => buttonToActivate.SetEnabled(evt.newValue));
           }
       }
   }
   ```
4. To try the example, from the menu, select **Window** > **ToggleExample**.

## Additional resources

* [Toggle](UIE-uxml-element-Toggle.md)