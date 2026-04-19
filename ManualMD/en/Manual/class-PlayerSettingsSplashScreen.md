# Splash Image Player settings

Configure the splash screen that displays when your application starts.

Use the **Splash Image Player** settings to customize or disable the Unity logo, add your own logos, or disable the splash screen entirely.

To access the **Splash Image Player** settings, go to **Edit** > **Project Settings** > **Player** > **Splash Image**. You can also use the [SplashScreen API](../ScriptReference/PlayerSettings.SplashScreen.md) to control the splash screen settings.

**Note**: Across all platforms, the Unity Splash Screen displays while the first **scene** of your application loads in the background. Custom splash screens or animations might take longer to appear because Unity must load the engine and first scene before it can display a custom splash screen.

![Splash Image Player settings window](../uploads/Main/player-settings-splash-image.png)

Splash Image Player settings window

## Virtual Reality Splash Image

Set an image to use for applications that use **virtual reality**. The image must be a 2D [Texture](Textures.md).

## Splash Screen

Settings for displaying the splash screen.

| **Setting** | **Description** | |
| --- | --- | --- |
| **Show Splash Screen** | Enable or disable the splash screen. | |
| **Preview** | Load a preview of the splash screen in the Game view. The preview reflects the resolution and **aspect ratio** of the Game view. | |
| **Splash Style** | Choose the color style of the Unity logo. | |
|  | **Light on Dark** | A light version of the Unity logo best suited to display on dark colored backgrounds. This is the default setting. |
|  | **Dark on Light** | A dark version of the Unity logo best suited to display on light colored backgrounds. |
| **Animation** | Customize how the splash screen appears and disappears on screen. | |
|  | **Static** | No animation occurs. |
|  | **Dolly** | The logo and background zoom to create a visual dolly effect. This is the default setting. |
|  | **Custom** | Configure the background and logo zoom amounts to modify the dolly effect. |
|  | **Logo Zoom** | Set the target zoom (from 0 to 1) for the logo when it reaches the end of the logo animation’s total duration.   This setting is available only when **Animation** is **Custom**. |
|  | **Background Zoom** | The target zoom (from 0 to 1) for the background when it reaches the end of the splash screen animation’s total duration.   This setting is available only when **Animation** is **Custom**. |

## Logos

Settings to customize the logos in your splash screen.

| **Setting** | **Description** | |
| --- | --- | --- |
| **Show Unity Logo** | Enabled by default. Disable this option to remove the Unity logo from the splash screen. | |
| **Draw Mode** | Select the order in which the logos appear on the splash screen.   This setting is available only when **Show Unity Logo** is enabled. | |
|  | **Unity Logo Below** | Display the Unity logo underneath all logos in the **Logos** list. |
|  | **All Sequential** | Display the Unity logo sequentially with all logos in the **Logos** list. |
| **Logos** | Choose the logos to include on the splash screen. Each logo must be a [Sprite asset](sprite/sprite-landing.md).   Do the following to update logos:  * To use the entire sprite as the logo, set the **Sprite Mode** to **Single**. * To change the logo’s aspect ratio, set the **Sprite Mode** to **Multiple** and edit the dimensions. For more information, refer to [Create sprites from a texture](sprite/sprite-editor/use-editor.md). * To add and remove logos, use the plus (+) and minus (-) buttons. * To reorder logos, drag them in the list. * To create a delay between logos, add an entry to the Logos list with no Sprite Asset assigned. | |
|  | **Logo Duration** | Set the length of time each Sprite Asset appears on the screen. Use a value between 2 and 10 seconds.   This setting is available only when **Draw Mode** is **All Sequential**. |
| **Splash Screen Duration** | Displays the total duration of the splash screen sequence, in seconds. This is the total of all logos plus 0.5 seconds for fading out.   **Note:** Your application’s splash screen might display for longer if the first scene isn’t ready to play. In this case, the splash screen shows only the background image or color and then fades out when the first scene is ready to play. | |

## Background

Settings to customize the background of your application’s splash screen.

| **Setting** | **Description** |
| --- | --- |
| **Overlay Opacity** | Adjust how opaque the overlay appears to make the logos stand out. This affects the background color and image color of the logo, based on the **Splash Style**.  Set the opacity to a lower value to reduce this effect. You can also disable the effect by setting it to 0. For example, if the **Splash Style** is **Light on Dark**, with a white background, the background becomes gray if the **Overlay Opacity** is 1, and white if set to 0.  **Note**: The **Overlay Opacity** setting has an effect on the display of the **Background Color** and might not match the assigned color. |
| **Background Color** | Set a color for the splash screen. Unity uses this color if you don’t set a **Background Image**.  **Note**: The **Overlay Opacity** setting has an effect on the display of the **Background Color** and might not match the assigned color.  **Note**: On Android, Unity also uses this color to hide the default automatic splash screen which appears on Android 12+ devices. This happens even if the **Background Image** is set, so make sure this color matches with **Background Image**. |
| **Blur Background Image** | Enable this setting to blur the background image you set. If you disable this setting, it displays the background image with no blur effect. |
| **Background Image** | Set a reference to a sprite to use as a background. Unity adjusts the background image so that it fills the screen. It uniformly scales the image until it fits both the width and height of the screen. This means that parts of the image might extend beyond the screen edges in some aspect ratios. To adjust the background image’s response to aspect ratio, change the sprite’s **Position** values in the [Sprite Editor tab of the Sprite Editor window](sprite/sprite-editor/sprite-editor-window-reference.md).   **Note:** Splash screen backgrounds don’t support HDR images. Avoid using HDR image formats such as RGBA16 for splash screen background images. |
| **Alternate Portrait Image** | Set an alternative image to use with portrait aspect ratios (for example, a mobile device in portrait mode). If you don’t assign an **Alternate Portrait Image** sprite, the Unity Editor uses the sprite assigned as the **Background Image** for both portrait and landscape mode.  You can adjust the position and dimensions of the sprite in the **Sprite Editor** to control the aspect ratio and position of the background image on the splash screen. |

## Additional resources

* [Player](class-PlayerSettings.md)
* [Project settings](comp-ManagerGroup.md)
* [SplashScreen API](../ScriptReference/PlayerSettings.SplashScreen.md)

PlayerSettingsSplashScreen