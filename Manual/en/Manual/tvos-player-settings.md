# tvOS Player Settings

tvOS shares many **Player Settings** with [iOS](class-PlayerSettingsiOS.md) but there are certain settings that are specific only to tvOS.

![tvOS Player Settings](../uploads/Main/tvos-player-settings.png)

tvOS Player Settings

## Icon

Use the **Icon** settings to customize the branding for your Apple TV app.

Apple TV images consist of between two and five layers. Unity only provides two layers for Apple TV icons. For more information on layering images for Apple TV, see the Apple Developer documentation on [Layered Images](https://developer.apple.com/design/human-interface-guidelines/tvos/icons-and-images/layered-images/).

| **Setting** | **Function** |
| --- | --- |
| **App icons** | Build the custom icon that you would like to appear on your [AppleTV home screen](https://developer.apple.com/design/human-interface-guidelines/tvos/overview/home-screen/) for each resolution (1280x768, 800x480, and 400x240). |
| **Top Shelf icons** | Build the custom icon that you would like to appear on the [Apple Top Shelf](https://developer.apple.com/design/human-interface-guidelines/tvos/overview/top-shelf/) for each aspect and resolution (4640x1440, 2320x720, 3840x1440, and 1920x720). |

## Splash Image

In addition to the [common Splash Screen settings](class-PlayerSettingsSplashScreen.md), there are two additional settings for the tvOS platform:

![tvOS Splash Settings](../uploads/Main/tvos-splash-settings.png)

tvOS Splash Settings

Use the **AppleTV (1x)** and **AppleTV (2x)** properties to set [Apple TV static splash screens](https://developer.apple.com/design/human-interface-guidelines/tvos/app-architecture/loading/).

## Configuration

tvOS has the same Configuration settings as iOS, except for one - the **Require Extended Game Controller** setting.

If your application requires a game controller, enable the **Require Extended Game Controller** setting. For more information, see the Apple Developer documentation on [Game Controllers](https://developer.apple.com/design/human-interface-guidelines/tvos/remote-and-controllers/game-controllers/).