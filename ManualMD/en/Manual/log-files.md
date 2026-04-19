# Log files reference

Unity produces log files for the Editor, package manager, licensing, development players, and Hub. You can use these log files to understand where any problems happened in your application.

Unity adds all messages, warnings, and errors from the [Console window](Console.md) to the log files. To add your own messages to the Console window, and the logs, use the [Debug class](class-Debug.md).

Each operating system stores the log files in different locations. The default locations are outlined on this page, but you can also use certain command line arguments to control when and where Unity generates log files. For more information, see the [Command line arguments](EditorCommandLineArguments.md) documentation.

You can use the [Application.consoleLogPath](../ScriptReference/Application-consoleLogPath.md) property in your project code to obtain the log location for the currently running Editor or Player application.

**Note:** Not all platforms support this feature. For more information, refer to [Platform development](PlatformSpecific.md).

## Editor-related log locations

You can access the Editor logs from the console window. To do this, open a Console Window (main menu: **Window** > **General** > **Console**) and select **Open Editor Log** from the Console window menu.

You can access the other logs by using your operating system’s file manager application.

### Linux

| **Log type** | **Log location** |
| --- | --- |
| **Editor** | `~/.config/unity3d/Editor.log` |
| **Package manager** | `~/.config/unity3d/upm.log` |
| **Licensing client** | `~/.config/unity3d/Unity/Unity.Licensing.Client.log` |
| **Licensing audits** | `~/.config/unity3d/Unity/Unity.Entitlements.Audit.log` |

### macOS

On macOS, you can also access Unity’s logs via the Console.app utility

| **Log type** | **Log location** |
| --- | --- |
| **Editor** | `~/Library/Logs/Unity/Editor.log` |
| **Package manager** | `~/Library/Logs/Unity/upm.log` |
| **Licensing client** | `~/Library/Logs/Unity/Unity.Licensing.Client.log` |
| **Licensing audits** | `~/Library/Logs/Unity/Unity.Entitlements.Audit.log` |

### Windows

On Windows, the Package Manager and Editor logs are placed in folders which aren’t shown in the Windows Explorer by default. To view the AppData folder, you must enable the Hidden Items setting on Windows. For more information on how to do this, see Microsoft’s documentation on [View hidden files and folders in Windows](https://support.microsoft.com/en-us/windows/view-hidden-files-and-folders-in-windows-97fbc472-c603-9d90-91d0-1166d1d9f4b5).

On Windows, a standard out stream doesn’t exist by default, so you must launch the Editor with a valid configured `stdout` stream, as a child process from a CI system. If you specify `-` to send output` `to`stdout`, then you won’t see the output in the console window.

| **Log type** | **Log location** |
| --- | --- |
| **Editor** | `%LOCALAPPDATA%\Unity\Editor\Editor.log` |
| **Package manager** | User account: `%LOCALAPPDATA%\Unity\Editor\upm.log`  SYSTEM account: `%ALLUSERSPROFILE%\Unity\Editor\upm.log` |
| **Licensing client** | `%LOCALAPPDATA%\Unity\Unity.Licensing.Client.log` |
| **Licensing audits** | `%LOCALAPPDATA%\Unity\Unity.Entitlements.Audit.log` |
| **Crash files** | `%TMP%\Unity\Editor\Crashes`    **Note:** You can overwrite the location of the folder location with the `-crash-report-folder` command line argument. |

## Player-related log locations

To view the Player log, open a Console Window (main menu: **Window** > **General** > **Console**) and select **Open Player Log** from the Console window menu. You can also navigate to the following folder:

| **Operating system** | ****Player log** location** |
| --- | --- |
| **Android** | To access the Player log for an Android application, use [Android logcat](https://developer.android.com/studio/command-line/logcat). For more information, see [View Android logs](android-debugging-on-an-android-device.md#view-android-logs). |
| **iOS** | Use the GDB console, or the Organizer Console through XCode to access iOS device logs. For more information on device logs, see [Apple’s documentation](https://developer.apple.com/documentation/xcode/acquiring-crash-reports-and-diagnostic-logs). |
| **Linux** | `~/.config/unity3d/CompanyName/ProductName/Player.log` |
| **macOS** | `~/Library/Logs/Company Name/Product Name/Player.log`  **Note:** You can also use the Console.app utility to find the log file. |
| **Universal Windows Platform** | `%USERPROFILE%\AppData\Local\Packages\<productname>\TempState\UnityPlayer.log` |
| **Web** | Unity writes the log output to your [browser’s JavaScript console](webgl-debugging.md). |
| **Windows** | `%USERPROFILE%\AppData\LocalLow\CompanyName\ProductName\Player.log` |

### Player crash files (Windows only)

For the location of Player crash files on Windows, refer to [`CrashReporting.crashReportFolder`](../ScriptReference/Windows.CrashReporting-crashReportFolder.md).

## Unity Hub log locations

You can access the Hub logs by using your operating system’s file manager application. You can also access these logs from within the Hub. For more information, see [Hub documentation](https://docs.unity3d.com/hub/manual/Help.html).

| **Operating system** | **Player log location** |
| --- | --- |
| **Linux** | `~/.config/UnityHub/logs/info-log.json` |
| **macOS** | `~/Library/Application\ Support/UnityHub/logs/info-log.json` |
| **Windows** | `%UserProfile%\AppData\Roaming\UnityHub\logs\info-log.json` |

## Additional resources

* [Console Window](Console.md)
* [Stack trace logging](stack-trace.md)