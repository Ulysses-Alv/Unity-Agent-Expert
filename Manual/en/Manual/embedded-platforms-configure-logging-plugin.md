# Configure a logging plug-in

Embedded platforms might have limitations on storing the log output to a file. As a result, these platforms collect log events in special logging environments, such as DLT and slogger2. To implement this functionality in Unity, you can create a logging plug-in using the API provided in a header file located at `<UnityInstallPath>/Editor/Data/PlaybackEngines/<EmbeddedLinux|QNX>/PluginAPI/unity_logging_plugin.h`.

```
#pragma once

#include <stdint.h>

#if __GNUC__
    #define UNITY_LOGGING_PLUGIN_API extern "C" __attribute__ ((visibility ("default")))
#else
    #error "Unsupported compiler/platform"
#endif

/// Increments on changes to the UnityLoggingPluginAPI structure
#define UNITY_LOGGING_PLUGIN_VERSION 1

enum UnityLoggingLevel
{
    kUnityLoggingLevelError = 0,
    kUnityLoggingLevelAssert = 1,
    kUnityLoggingLevelWarning = 2,
    kUnityLoggingLevelLog = 3,
    kUnityLoggingLevelException = 4,
    kUnityLoggingLevelDebug = 5,
    kUnityLoggingLevelNumLevels
};

enum UnityLoggingStatus
{
    kUnityLoggingStatusSuccess = 0,
    kUnityLoggingStatusUnsupportedVersion = 1,
    kUnityLoggingStatusLoadFailure = 2,
    kUnityLoggingStatusUnloadFailure = 3,
    kUnityLoggingStatusDependencyFailure = 4,
};

typedef struct {
    const char* (*const GetString)(const char* key, const char* defaultValue);
    uint64_t (*const GetUInt64)(const char* key, uint64_t defaultValue);
    int64_t (*const GetInt64)(const char* key, int64_t defaultValue);
    uint32_t (*const GetUInt32)(const char* key, uint32_t defaultValue);
    int32_t (*const GetInt32)(const char* key, int32_t defaultValue);
} UnityLoggingPluginConfig;

typedef struct {
    typedef UnityLoggingStatus (*UnityLoggingLoad)(UnityLoggingPluginConfig* config);
    typedef UnityLoggingStatus (*UnityLoggingUnload)();
    typedef void (*UnityLoggingLog)(UnityLoggingLevel, const char*, uint32_t);

    UnityLoggingLoad Load;
    UnityLoggingUnload Unload;
    UnityLoggingLog Log;
} UnityLoggingPluginAPI;

UNITY_LOGGING_PLUGIN_API UnityLoggingStatus unity_logging_plugin_get_api(UnityLoggingPluginAPI* api, uint32_t version);
```

You can enable the logging **plug-in** by providing a Player argument `-platform-hmi-log-plugin` and specifying the name of the plug-in’s shared library to be loaded from the `Plugins` folder. The shared library must be accessible to the dynamic loader allowing it to be loaded using [`dlopen`](https://learn.microsoft.com/en-us/dotnet/api/objcruntime.dlfcn.dlopen?view=xamarin-ios-sdk-12#objcruntime-dlfcn-dlopen\(system-string-system-int32\)) and initialized. For more information, refer to [Command line arguments for logging](embedded-platforms-command-line-arguments.md).

If you set the `-logfile` argument together with the logging plug-in argument, the plug-in initializes, logs a single test message, and always results in a plug-in error. These steps are logged to help identify possible issues.

If the plug-in fails to initialize, Unity returns an error exit code by default. In this case, you can use the `-platform-hmi-log-disable-on-plugin-failure` argument to keep the Player running without any log output.

**Note**: Logging plug-in samples are available on request with embedded platforms support contract. To request access, contact the embedded platforms support team.

## Integrate the plug-in with Logging Player setting

You can configure the logging plug-in to integrate with the **Logging** Player setting available for the [embedded Linux](embedded-linux-player-settings.md) and [QNX](qnx-player-settings.md) platforms. To do so, consider the following steps and refer to the following code example:

1. Use the [SettingsProvider](../ScriptReference/SettingsProvider.md) class to add the **Logging** setting.
2. Use a class name that ends with `PluginBuildPostProcessor`.
3. Make sure the signature of the `ProcessConfiguration` method matches the format mentioned in the following code example.
4. Set the plug-in library using the `platform-hmi-log-plugin` argument.

```
namespace MyPluginNamespace
{
    public static class MyPluginBuildPostProcessor
    {
        private const string LoggingPluginKey = "platform-hmi-log-plugin";
        public static void ProcessConfiguration(Action<string, string> setValue)
        {
            // Set the shared library name for the plugin
            setValue(LoggingPluginKey, "mylogplugin.so");

            // (Optional) Set specific key values for the plugin
            setValue("my-plugin-setting", "42");
        }
    }

}
```

## Additional resources

* [Command line arguments for logging](embedded-platforms-command-line-arguments.md)
* [Standard log output overview](embedded-platforms-standard-logoutput.md)