# Build your application for Dedicated Server

You can create a Dedicated Server build in either of the following ways:

* [Unity Editor](#unity-editor)
* [Scripting](#scripting)
* [Command line](#command-line)

## Unity Editor

To create a Dedicated Server build in the Unity Editor, use the following steps:

1. Open the [Build Profiles](BuildSettings.md) window from **File** > **Build Profiles**.
2. Select **Add Build Profile** to open the **Platform Browser** window.
3. Select the type of server to build from the list of available platforms. For example, select **Linux Server** to build a Linux server.
4. If the server type isn’t available, select **Install with Unity Hub** and follow the installation instructions. For information on how to install modules, refer to [Add modules](https://docs.unity3d.com/hub/manual/AddModules.html).
5. Select **Switch Profile** to set the new build profile as the active profile.
6. Click **Build**.

**Tip**: You can further configure the Dedicated Server build in the [Player settings](dedicated-server-player-settings.md).

## Scripting

To create a Dedicated Server build using a script, set `buildPlayerOptions.subtarget` to `(int)StandaloneBuildSubtarget.Server`.

```
buildPlayerOptions.target = BuildTarget.StandaloneWindows;
// SubTarget expects an integer.
buildPlayerOptions.subtarget = (int)StandaloneBuildSubtarget.Server;
```

## Command line

To create a Dedicated Server build through the command line, use the `-standaloneBuildSubtarget Server` [argument](CommandLineArguments.md).

```
-buildTarget Linux64 -standaloneBuildSubtarget Server
```

## Code sign macOS Dedicated Server builds

Dedicated Server builds that aren’t code signed might display security warnings when deployed on macOS systems. To avoid such warnings, make sure you code sign the build before distribution. For more information, refer to the documentation on [Code sign and notarize your macOS application](macos-building-notarization.md).

## Additional resources

* [macOS build settings reference](macosbuildsettings.md)
* [Windows build settings reference](WindowsStandaloneBinaries.md)
* [Linux build settings reference](Buildsettings-linux.md)
* [Creating and Using Scripts](creating-scripts.md)