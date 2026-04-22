# Display camera views on multiple monitors

You can use multi-display to display up to eight different Camera views of your application on up to eight different monitors at the same time. You can use this for setups such as PC games, arcade game machines, or public display installations.

Unity supports multi-display on:

* Desktop platforms (Windows, macOS X, and Linux)
* Android (OpenGL ES and Vulkan)
* iOS

Some features work only on some platforms. See the [Display](../ScriptReference/Display.md), [Screen](../ScriptReference/Screen.md) and [FullScreenMode](../ScriptReference/FullScreenMode.md) APIs for more information about compatibility.

## Activating multi-display support

Unity’s default display mode is one monitor only. When you run your application, you need to use `Display.Activate()` to explicitly activate additional displays. Once you activate a display, you can’t deactivate it.

The best time to activate additional displays is when your application creates a new **Scene**. A good way to do this is to attach a script component to the default Camera. Make sure you call `Display.Activate()` only once during startup. As a best practice, you might find it helpful to create a small initial Scene to test your script.

### Example script

```
using UnityEngine;
using System.Collections;

public class ActivateAllDisplays : MonoBehaviour
{
    void Start ()
    {
        Debug.Log ("displays connected: " + Display.displays.Length);
            // Display.displays[0] is the primary, default display and is always ON, so start at index 1.
            // Check if additional displays are available and activate each.
    
        for (int i = 1; i < Display.displays.Length; i++)
            {
                Display.displays[i].Activate();
            }
    }
    
    void Update()
    {

    }
}
```

## Previewing multiple displays in your Project

To preview different Camera views, follow these steps:

1. In the Camera’s **Inspector**, select a **Target Display** for that Camera.

![Camera Inspector with Target Display option](../uploads/Main/InspectorCamera35.png)

Camera Inspector with Target Display option

1. Make sure you’re in the Game view.
2. From the **Display** menu in the top-left corner, select the *Display* to Preview.

![Display preview in the top left corner of the Game View](../uploads/Main/TargetDisplay.png)

Display preview in the top left corner of the Game View

## API support

Unity supports the following **UnityEngine.Display** API methods:

| **Method** | **Description** |
| --- | --- |
| `public void Activate()` | Activates a specific display on the current monitor’s width and height. This call must be made once on starting a new Scene. It can be called from a user script attached to a **Camera** or dummy **GameObject** in a new Scene. |
| `public void Activate(int width, int height, int refreshRate)` | Windows only. Activates a specific display of custom width and height. On Linux and macOS X, secondary displays always use the display’s current resolution, if available. |