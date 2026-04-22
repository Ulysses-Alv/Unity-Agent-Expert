# Stop and restart Unity Accelerator

Unity **Accelerator** runs as a background process for each platform. To stop or restart Unity Accelerator for each platform, perform the following steps:

## Windows

1. Open the **Services Panel** by searching for the term `Services` in the Settings menu, or running `services.msc` in the **Run** dialog (**Win** + **R**).
2. Locate the `Unity Accelerator` service in the list. The option to **Stop the service** or **Restart the service** appears on the left panel.

Alternatively you can use the `net stop` and `net start` console commands from a terminal to stop and start the service.
Examples:
`C:\> net stop "Unity Accelerator" # Stops the service
C:\> net start "Unity Accelerator" # Starts the service`

## macOS

Run the `launchctl` command from the terminal to control the `com.unity3d.accelerator` service from the LaunchControl utility. For more information, refer to <https://www.launchd.info/>.

Examples:

```
$ sudo launchctl bootout system /Library/LaunchDaemons/com.unity3d.accelerator.plist    # Stops the service
$ sudo launchctl bootstrap system /Library/LaunchDaemons/com.unity3d.accelerator.plist  # Starts the service
$ sudo launchctl kickstart -k system/com.unity3d.accelerator                            # Restarts the service
```

## Linux

Use the `systemctl` console utility to control the `unity-accelerator` service. For more information, refer to <https://manpages.ubuntu.com/manpages/bionic/en/man1/systemctl.1.html>.

Examples:

```
$ sudo systemctl stop unity-accelerator.service     # Stops the service
$ sudo systemctl start unity-accelerator.service    # Starts the service
$ sudo systemctl restart unity-accelerator.service  # Restarts the service
```

## Additional resources

* [Monitor Unity Accelerator](accelerator-monitor.md)
* [Use Unity Accelerator on the command line](accelerator-command-line.md)
* [Command line arguments reference](EditorCommandLineArguments.md#accelerator)