# Command line arguments for logging

Use the following command line arguments to configure the standard log output and logging **plug-in** for embedded platforms.

| **Argument** | **Description** | **Use with** |
| --- | --- | --- |
| `-nolog` | Disables the log output completely. | standard log output |
| `-logfile [path/to/logfile]` | Routes the Player log output including `stderr` to the specified log file. You can use this parameter to get trace information when the logging plug-in initializes. The log file path set using this argument has a precedence over the path set using `-platform-hmi-player-datapath`.  **Note:** If the specified file path has write restrictions, the Unity Player exits with an error code. | standard log output and logging plug-in |
| `-logfile -` | Forces the Player to route the log output including `stderr` to the console through `stdout` instead of writing to the specified file. You can use this parameter to get trace information when the logging plug-in initializes. | standard log output and logging plug-in |
| `-logflush` | Synchronizes the log output with every log event, which can be helpful in crash investigations. | standard log output |
| `-platform-hmi-log-plugin [plugin.so]` | Specifies the logging plug-in’s shared library. Needs to be located in the Plugins folder. | logging plug-in |
| `-platform-hmi-log-plugin [" "]` | Disables the configured logging plug-in temporarily. | logging plug-in |
| `-platform-hmi-log-disable-on-plugin-failure` | Disables the log output when the logging plug-in fails to initialize. | logging plug-in |
| `-platform-hmi-player-datapath [path/to/writable/folder]` | Specifies the path where the Player log output is saved as persistent data. By default, Unity uses the value of `HOME` environment variable to set this path to your system’s home directory. You can configure this path through **Player Settings** > **Other Settings** > **Configuration** > **Player Data path**. | standard log output |

## Additional resources

* [Standard log output overview](embedded-platforms-standard-logoutput.md)
* [Configure a logging plug-in](embedded-platforms-configure-logging-plugin.md)