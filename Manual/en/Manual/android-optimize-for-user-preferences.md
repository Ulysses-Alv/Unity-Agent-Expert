# Optimize for user preferences

Android’s [game mode](https://developer.android.com/games/optimize/adpf/gamemode/about-API-and-interventions) feature indicates how the user wants to optimize your application. It enables users to say whether they want your application to run normally, optimize for battery life, or optimize for best performance.

## Requirements and compatibility

Android’s game mode feature requires Android version 13.

## Use game mode in Unity

Unity provides the [AndroidGame.GameMode](../ScriptReference/Android.AndroidGame.GameMode.md) property that you can use to retrieve the current game mode for your application.

Depending on the current game mode, you should adapt your Unity application to accommodate the user’s preference. For example, if the game mode is [battery saving](../ScriptReference/Android.AndroidGameMode.Battery.md), the user expects your application to run for as long as possible. In this case, you should reduce the resource intensity of effects and calculations to reduce the impact your application has on the device’s battery life. This includes things like:

* Decrease the frame rate or resolution of the application.
* Reduce the **LOD** bias to make Unity display lower detail LODs closer to the **camera**.
* Reduce the cull distance of particular layers.
* Reduce or possible batch network calls and/or sensor usage.

If the game mode is set to [maximize performance](../ScriptReference/Android.AndroidGameMode.Performance.md), the user expects your application to look and play as well as possible. In this case, you can enable more effects and not be as wary of battery consumption.

**Tip**: If your application is resource intensive and you want to enhance its overall performance, consider using [Adaptive Performance](adaptive-performance/adaptive-performance.md). It provides feedback about the thermal and power state of a mobile device and enables you to react appropriately to prevent thermal throttling or excessive battery consumption.