# Detect Game Controllers

Unity includes the **Game Controller** framework in the project only if a script in the project references [`Input.GetJoystickNames`](../ScriptReference/Input.GetJoystickNames.md). If it’s available, Unity iOS Runtime loads the framework dynamically.

To get the list of all available controllers, call [`Input.GetJoystickNames`](../ScriptReference/Input.GetJoystickNames.md). You can re-check this list at any time to detect if controllers have been attached or detached.

You can call this API to detect the type of controller that’s attached. Names follow the pattern: `[$profile_type,$connection_type] joystick $number by $model`. `$profile_type` can be either **basic** or **extended**, and `$connection_type` is either **wired** or **wireless**. When Unity detects at least one controller, you can either disable on-screen touch controls or amend them to supplement controller input.

## Example: Detect attached Game Controllers

The following code sample checks to see if any controllers are connected to the device.

```
using System.Collections;
using UnityEngine;

public class GameControllers : MonoBehaviour
{
    private bool connected = false;

    IEnumerator CheckForControllers() {
        while (true) {
            var controllers = Input.GetJoystickNames();

            if (!connected && controllers.Length > 0) {
                connected = true;
                Debug.Log("Connected");
            
            } else if (connected && controllers.Length == 0) {         
                connected = false;
                Debug.Log("Disconnected");
            }

            yield return new WaitForSeconds(1f);
        }
    }

    void Awake() {
        StartCoroutine(CheckForControllers());
    }
}
```

## Additional resources:

* [Input class](../ScriptReference/Input.md)