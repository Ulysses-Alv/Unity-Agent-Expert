# Troubleshooting common cross-platform issues

Most of Unity’s API and project structure is identical for all supported platforms and sometimes a project can be rebuilt to run on different devices. However, fundamental differences in the hardware and deployment methods mean that parts of a project may not port between platforms without change. This page lists some common cross-platform issues and suggestions for solving them.

## Input

The best example of different behaviour between platforms is in the input methods offered by the hardware.

### Keyboard and joypad

The **Input.GetAxis** function is convenient on desktop platforms to consolidate keyboard and joypad input. This function isn’t suitable for mobile platforms that rely on touchscreen input. The standard desktop keyboard input is only suitable for porting typed text to mobile devices. You can add a layer of abstraction to your input code if you are considering porting to other platforms in the future. For example, if you are making a driving game, then you might create your own input class and wrap the Unity API calls in your own functions:

```
using UnityEngine;

public class CarControls : MonoBehaviour
{
    // Returns values in the range of  -1.0 to 1.0, which correspond to left and right.
    public float Steering()
    {
        return Input.GetAxis("Horizontal");
    }

    // Returns values in the range of -1.0 to 1.0, which correspond to acceleration and braking.
    public float Acceleration()
    {
        return Input.GetAxis("Vertical");
    }

    private int currentGear = 0;

    // Returns an integer corresponding to the selected gear.
    public int Gears()
    {
        if (Input.GetKeyDown(KeyCode.P))
            currentGear++;
        else if (Input.GetKeyDown(KeyCode.L))
            currentGear--;

        return currentGear;
    }
}
```

Wrapping the API calls in a class puts them into a single source file and makes the calls easy to locate and replace. You should design your input functions according to the logical meaning of the inputs in your game. This helps isolate the rest of the game code from the specific method of input used with a particular platform.

For example, you could modify the Gears function above so that the actual input comes from touches on the screen of a mobile device. You can use an integer to represent the chosen gear works for all platforms, but if you mix platform-specific API calls with the rest of the code, you could cause problems. You can also use platform dependent compilation to combine the different implementation of the input functions in the same source file and avoid manual swaps.

### Touches and clicks

The **Input.GetMouseButtonXXX** functions are designed to have an obvious interpretation on mobile devices. The screen reports a simple touch as a left click and the **Input.mousePosition** property gives the position of the touch, as long as the finger is touching the screen. Games with simple mouse interactions can often work transparently between the desktop and mobile platforms.

Conversion is often much less straightforward than this. For example, a desktop game can make use of more than one mouse button and a mobile game can detect multiple touches on the screen at a time.
To help manage this, you can represent input with logical values that are then used by the rest of the game code.

For example, you might replace a pinch gesture to zoom on a mobile device with a +/- keystroke on the desktop; the input function could return a float value that specifies the zoom factor. Likewise, it might be possible to use a two-finger tap on a mobile to replace a right button click on the desktop. However, if the properties of the input device are an integral part of the game then it may not be possible to remodel them on a different platform. This might mean that you can’t port the game or that you need to modify the input or gameplay extensively.

### Accelerometer, compass, gyroscope and GPS

These inputs derive from the mobility of handheld devices and so may not have any meaningful equivalent on the desktop. However, some use cases simply mirror standard game controls and are ported easily. For example, a driving game might implement the steering control from the tilt of a mobile device (determined by the accelerometer). In cases like this, the input API calls are usually easy to change, so you can replace the accelerometer input with keystrokes.

However, it might be necessary to recalibrate inputs or vary the game difficulty to adjust to the different input method. Tilting a device is slower and more strenuous than pressing keys, which makes it harder to concentrate on the display. This might make the game more difficult to master on a mobile device, so it might be appropriate to slow down gameplay or allow more time per level. This requires designing the game code to adjust these factors.

## Memory, storage and CPU performance

Mobile devices have less storage, memory and CPU power available than desktop machines and so a game may be difficult to port simply because its performance isn’t acceptable on lower powered hardware. If you are pushing the limits of your desktop hardware, then the game probably isn’t a good candidate for porting to a mobile platform.

### Storage requirements

Video, audio, and textures can use a lot of storage space. You need to manage storage effectively if you want to port your game. Storage space (which often also corresponds to download time) is usually not an issue on desktop machines, but it can be limited on mobile devices. Mobile app stores often impose a limit on the maximum size of a submitted product. It might require some planning to address these concerns during your game development.

For example, you may need to provide cut-down versions of assets for mobiles to save space. Another possibility is that the game may need to be designed so that large assets can be downloaded on demand rather than being part of the initial download of the application.

### Automatic memory management

Unity automatically handles the recovery of unused memory from “dead” objects and often happens unnoticed on desktop machines. However, the lower memory and CPU power on mobile devices means that garbage collections can be more frequent, impacting performance and causing unwanted pauses in gameplay. Even if the game runs in the available memory, it might be necessary to optimize code to avoid garbage collection pauses.

For more information, refer to the [memory management page](performance-managed-memory.md).

### CPU power

A game that runs well on a desktop machine might experience poor frame rate on a mobile device because the mobile CPU struggles with the game complexity. Pay attention to code efficiency when a project is ported to a mobile platform.