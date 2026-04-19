# Profiler command line arguments

Set how the Profiler starts from the command line.

If you start your built player or the Unity Editor from the command line (such as the Command Prompt on Windows, Terminal on macOS, Linux shell, or adb for Android), you can pass command line arguments to configure some Profiler settings. For more information about starting Unity from the command line, refer to [Command line arguments](CommandLineArguments.md).

| **Command line argument** | **Description** |
| --- | --- |
| `-profiler-enable` | Profile the start-up of a player or the Editor.   When you use this argument with a player, it has the same effect as building the player with the [Autoconnect Profiler](profiler-build-settings-reference.md) setting enabled.   When you use this argument with the Editor, it starts collecting and displaying Profiler information in the Profiler window on start-up of the Editor. |
| `-profiler-log-file <Path/To/Log/File.raw>` | Stream profile data to a .raw file on startup. This argument works for both players and the Editor. |
| `-profiler-capture-frame-count <NumberOfFrames>` | Set how many frames to capture in a profile when streaming to a .raw file on start-up. This argument only works on players. |
| `-profiler-maxusedmemory` | Set a maximum amount of memory in bytes for the Profiler to use at start up (for example, `-profiler-maxusedmemory 16777216`). By default, `maxUsedMemory` is 16MB for players and 256MB for the Editor. |
| `-connection-id` | Set a custom player name set when you launch a player from the command line. You can also set this name in the [Profiler Preferences window](profiler-preferences-reference.md). |

## Additional resources

* [Profiling your application](profiler-profiling-applications.md)
* [Profiler Preferences reference](profiler-preferences-reference.md)
* [Profiler Build Profiles settings reference](profiler-build-settings-reference.md)