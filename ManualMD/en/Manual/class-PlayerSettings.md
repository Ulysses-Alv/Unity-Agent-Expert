# Player

[Switch to Scripting](../ScriptReference/PlayerSettings.md "Go to PlayerSettings page in the Scripting Reference")

The **Player** settings window (menu: **Edit** > **Project Settings** > **Player**) contain settings that determine how Unity builds and displays your final application. You can use the [PlayerSettings](../ScriptReference/PlayerSettings.md) API to control most of the settings available in this window.

**Note**: You can access **Player** settings from the **Build Profiles** window (menu: **File** > **Build Profiles**).

## General settings

The **Player settings** differ between the [platform modules](https://docs.unity3d.com/hub/manual/AddModules.html) that you’ve installed. Each [platform](PlatformSpecific.md) has its own Player settings which you’ll need to set for each version of your application you want to build. To navigate between them, click on the tabs with the platform operating system icon on.

![Player settings window](../uploads/Main/player-settings-window.png)

Player settings window

There are some general settings that all platforms share by default, unless you use [build profiles](build-profiles.md).

**Note**: With build profiles, you can customize the Player settings per build profile to set different values for each platform. For more information, refer to [Override settings with build profiles](build-profiles-override-settings.md).

| **Property** | **Function** |
| --- | --- |
| **Company Name** | Enter the name of your company. Unity uses this to locate the preferences file. |
| **Product Name** | Enter the name that appears on the menu bar when your application is running. Unity also uses this to locate the preferences file. |
| **Version** | Enter the version number of your application. |
| **Default Icon** | Pick the Texture 2D file that you want to use as a default icon for the application on every platform. You can override this for specific platforms. |
| **Default Cursor** | Pick the Texture 2D file that you want to use as a default cursor for the application on every supported platform. |
| **Cursor Hotspot** | Set the **pixel** offset value from the top left of the default cursor to the location of the cursor hotspot. The cursor hotspot is the point in the cursor image that Unity uses to trigger events based on cursor position. |

## Platform-specific settings

The platform-specific settings are divided into the following sections:

* **Icon**: the game icon(s) as shown on the desktop. You can choose icons from 2D image assets in the Project, such as **sprites** or imported images.
* **Resolution and Presentation**: settings for screen resolution and other presentation details such as whether the game should default to fullscreen mode.
* **Splash Image**: the image shown while the game is launching. This section also includes common settings for creating a Splash Screen. For more information, refer to the [Splash Image](class-PlayerSettingsSplashScreen.md) documentation.
* **Other Settings**: any remaining settings specific to the platform.
* **Publishing Settings**: details of how the built application is prepared for delivery from the app store or host webpage.
* ****XR** Settings**: settings specific to [Virtual Reality, Augmented Reality, and Mixed Reality](XR.md) applications.

You can find information about the settings specific to individual platforms in the [platform’s own manual section](PlatformSpecific.md):

* **Android:** [Android Player settings](class-PlayerSettingsAndroid.md)
* **Dedicated Server:** [Dedicated Server Player settings](dedicated-server-player-settings.md)
* **Embedded Linux:** [Embedded Linux Player settings](embedded-linux-player-settings.md)
* **iOS:** [iOS Player settings](class-PlayerSettingsiOS.md)
* **Linux:** [Linux Player settings](PlayerSettings-linux.md)
* **macOS:** [macOS Player settings](PlayerSettings-macOS.md)
* **QNX:** [QNX Player settings](qnx-player-settings.md)
* **tvOS:** [tvOS Player settings](tvos-player-settings.md)
* **Universal Windows Platform:** [UWP Player settings](class-PlayerSettingsWSA.md)
* **Web and Facebook Instant Games:** [Web Player settings](class-PlayerSettingsWebGL.md)
* **Windows:** [Windows Player settings](playersettings-windows.md)

You can find details of **closed platform** Player settings in their respective documentation.

## Additional resources

* [Create and manage build profiles](create-build-profile.md)

PlayerSettings