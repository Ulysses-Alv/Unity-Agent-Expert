# Capturing frame rate

A special case of time management is where you want to record gameplay as a video. Since the task of saving screen images takes considerable time, the game’s normal frame rate is reduced, and the video doesn’t reflect the game’s true performance.

To improve the video’s appearance, use the [`Time.captureFramerate`](../ScriptReference/Time-captureFramerate.md) property. The property’s default value is 0, for unrecorded gameplay. When you set the property’s value to anything other than zero, game time is slowed and the frame updates are issued at precise regular intervals for the purposes of recording. The interval between frames is equal to `1 / Time.captureFramerate`, so if you set the value to 5.0 then updates occur every fifth of a second.

With the demands on frame rate effectively reduced, you have time in the `Update` or `LateUpdate` functions to save screenshots or take other actions. The following code example demonstrates this:

```
//C# script example
using UnityEngine;
using System.Collections;

public class ExampleScript : MonoBehaviour {
    // Capture frames as a screenshot sequence. Images are
    // stored as PNG files in a folder - these can be combined into
    // a movie using image utility software (eg, QuickTime Pro).
    // The folder to contain our screenshots.
    // If the folder exists we will append numbers to create an empty folder.
    string folder = "ScreenshotFolder";
    int frameRate = 25;
        
    void Start () {
        // Set the playback frame rate (real time will not relate to game time after this).
        Time.captureFramerate = frameRate;
        
        // Create the folder
        System.IO.Directory.CreateDirectory(folder);
    }

    // Use LateUpdate to capture the screen after all other updates have been processed.
    void LateUpdate()
    {
        // Append filename to folder name (format is '0005 shot.png"')
        string name = string.Format("{0}/{1:D04} shot.png", folder, Time.frameCount );

        // Capture the screenshot to the specified file.
        ScreenCapture.CaptureScreenshot(name);
    }
}
```

Using this technique improves the video, but can make the game much harder to play. Try different values of `Time.captureFramerate` to find a good balance.

## Additional resources

* [Time.captureFramerate API reference](../ScriptReference/Time-captureFramerate.md)