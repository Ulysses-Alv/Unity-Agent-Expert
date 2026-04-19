# Native plug-in API for logging

Write to the Unity logs from your native **plug-in** code using the `IUnityLog` interface. Low-level native APIs for logging are provided in the `IUnityLog.h` header file, located in the [PluginAPI folder](native-plugin-interface-introduction.md#plugin-api-folder).

The file contains a single `Log` function with the following decalaration:

`void(UNITY_INTERFACE_API * Log)(UnityLogType type, const char* message, const char *fileName, const int fileLine);`

You can call this function directly as follows:

`s_UnityLog->Log(kUnityLogTypeLog, "Here is a regular log", __FILE__, __LINE__);`

However, for convenience the native logging API defines the following macros that wrap different log-level calls to the `Log` function:

| **Macro** | **Description** |
| --- | --- |
| `UNITY_LOG(PTR_, MSG_)` | Uses the log interface passed as a pointer (`PTR_`) to write the provided char\* `MSG_` as a regular log message. Equivalent to the managed API [Debug.Log](../ScriptReference/Debug.Log.md). |
| `UNITY_LOG_WARNING(PTR_, MSG_)` | Uses the log interface passed as a pointer (`PTR_`) to write the provided char\* `MSG_` as a warning-level log message. Equivalent to the managed API [Debug.LogWarning](../ScriptReference/Debug.LogWarning.md). |
| `UNITY_LOG_ERROR(PTR_, MSG_)` | Uses the log interface passed as a pointer (`PTR_`) to write the provided char\* `MSG_` as an error-level log message. Equivalent to the managed API [Debug.LogError](../ScriptReference/Debug.LogError.md). |

The following code example implements the `IUnityLog` interface in C++ and uses these predefined macros to write different levels of log output:

```
#include "IUnityLog.h"

static IUnityLog* s_UnityLog = NULL;

// Additional macros to include file and line number from the native code
#define UNITY_LOG_STRINGIZE_DETAIL(x) #x
#define UNITY_LOG_STRINGIZE(x) UNITY_LOG_STRINGIZE_DETAIL(x)
#define COMPOSE(MESSAGE) "[" __FILE__ ":" UNITY_LOG_STRINGIZE(__LINE__) "] " MESSAGE
#define NATIVE_LOG(PTR, MESSAGE) UNITY_LOG(PTR, COMPOSE(MESSAGE))
#define NATIVE_WARNING(PTR, MESSAGE) UNITY_LOG_WARNING(PTR, COMPOSE(MESSAGE))
#define NATIVE_ERROR(PTR, MESSAGE) UNITY_LOG_ERROR(PTR, COMPOSE(MESSAGE))

// Unity plugin load event
extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginLoad(IUnityInterfaces * unityInterfacesPtr)
{
    s_UnityLog = unityInterfacesPtr->Get<IUnityLog>();
}

// Unity plugin unload event
extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginUnload()
{
    s_UnityLog = nullptr;
}

extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API GenerateLog()
{
    // Output different log level messages to the Unity console
    UNITY_LOG(s_UnityLog, "Regular log message");
    UNITY_LOG_WARNING(s_UnityLog, "Warning log message");
    UNITY_LOG_ERROR(s_UnityLog, "Error log message");

    // Wrap log functions to provide native file and line number in output
    NATIVE_LOG(s_UnityLog, "Regular log with native file name and line number");
    NATIVE_WARNING(s_UnityLog, "Warning log with native file name and line number");
    NATIVE_ERROR(s_UnityLog, "Error log with native file name and line number");
}
```

**Note**: The additional macros to embed the native file and line number in the message are a workaround for a known issue in which Unity fails to include the information when the `Log` function is invoked via `UNITY_LOG` and equivalents.

## Additional resources

* [The Debug class](class-Debug.md)
* [Log files reference](log-files.md)