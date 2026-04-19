# Native plug-in API for profiling

You can use the low-level native **plug-in** profiler API to extend [the Profiler](profiler-introduction.md) and collect performance data of native plug-in code, or prepare profiling data to send to third-party profiling tools such as Razor (PS4), PIX (Xbox, Windows), Chrome Tracing, ETW, ITT, Vtune, or Telemetry.

The low-level native plug-in Profiler APIs provide the following interfaces for communication between the Unity Profiler and external tooling:

* **IUnityProfiler**: Use this interface to add instrumentation events to the Unity Profiler from native plug-in code.
* **IUnityProfilerCallbacks**: Use this interface to intercept Unity Profiler events and store or redirect them to other tools.

## IUnityProfiler API reference

Use the `IUnityProfiler` plug-in API to add instrumentation to your native plug-ins. The plug-in API is represented by `IUnityProfiler` interface, which is declared in the `IUnityProfiler.h` header, located in the [PluginAPI folder](native-plugin-interface-introduction.md#plugin-api-folder).

| **Method** | **Description** |
| --- | --- |
| `CreateMarker` | Creates a **Profiler marker** which represents a named instrumentation scope, which you can then use to generate instrumentation samples. |
| `SetMarkerMetadataName` | Specifies custom parameter names which can be passed along with an instrumentation sample for a Profiler marker. |
| `BeginSample` | Begins an instrumentation section of a code named after the Profiler marker. |
| `EndSample` | Ends an instrumentation section. |
| `EmitEvent` | Emits generic events with metadata. |
| `IsEnabled` | Returns 1 if the Profiler is capturing the data. |
| `IsAvailable` | Returns 1 for Editor or Development Players where the Profiler is available, and 0 for Release Players. |
| `RegisterThread` | Registers the current thread under the specified name. |
| `UnregisterThread` | Unregisters the current thread from Profiler. |

### IUnityProfiler example

The following example generates Profiler events that the Profiler window can display:

```
#include <IUnityInterface.h>
#include <IUnityProfiler.h>

static IUnityProfiler* s_UnityProfiler = NULL;
static const UnityProfilerMarkerDesc* s_MyPluginMarker = NULL;
static bool s_IsDevelopmentBuild = false;

static void MyPluginWorkMethod()
{
    if (s_IsDevelopmentBuild)
        s_UnityProfiler->BeginSample(s_MyPluginMarker);

    // Code I want to see in Unity Profiler as "MyPluginMethod".
    // ...

    if (s_IsDevelopmentBuild)
        s_UnityProfiler->EndSample(s_MyPluginMarker);
}

extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginLoad(IUnityInterfaces* unityInterfaces)
{
    s_UnityProfiler = unityInterfaces->Get<IUnityProfiler>();
    if (s_UnityProfiler == NULL)
        return;
    s_IsDevelopmentBuild = s_UnityProfiler->IsAvailable() != 0;
    s_UnityProfiler->CreateMarker(&s_MyPluginMarker, "MyPluginMethod", kUnityProfilerCategoryOther, kUnityProfilerMarkerFlagDefault, 0);
}

extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginUnload()
{
    s_UnityProfiler  = NULL;
}
```

## IUnityProfilerCallbacks API callbacks

The native Profiler plug-in API provides an interface between Unity’s subsystems and third-party profiling APIs so you can use an external profiling tool to profile your Unity application. The `IUnityProfilerCallbacks` header exposes the API, which Unity stores in the `<UnityInstallPath>\Editor\Data\PluginAPI` folder of your Unity installation. (On macOS, right-click on the Unity application, and select **Show Package Contents**. The header is in `Contents\PluginAPI`).

The following Unity Profiler features help capture instrumentation data so you can analyze the performance of your application:

| **Profiler feature** | **Description** |
| --- | --- |
| **Categories** | Unity groups profile data into categories (such as Rendering, Scripting and Animation), and assigns a color to each category. Color-coded categories help you visually distinguish the types of data in the Profiler window. The Profiler native plug-in API retrieves these colors so you can use them in an external profiling tool. |
| **Usage flags** | Usage flags act as a filter that reduce the amount of data that Unity sends to an external profiling tool. You can use usage flags to remove unnecessary information from the profiling data before Unity sends it to the external tool. The Profiler applies the following usage flags to event markers so that you can filter the data:  **Availability flags** Flags whether a marker is available in the Unity Editor, a development player, or a release player.  **Verbosity levels** Relates to the type of the task you’re doing in the Editor, and the level of information that task requires (for example, internal, debug, or user level). |
| **Frame events** | You can use the Profiler native plug-in API to perform frame-time analysis in an external profiling tool. |
| **Thread profiling** | Unity does a significant amount of work on threads (for example, the main thread, render thread and job system worker thread). You can use the Profiler native plug-in API to enable profiling on any thread. |

To use the instrumentation data that the Unity Profiler generates in an external profiler, you can use these minimal set of callbacks in your C/C++ plug-in code that integrates the third-party profiler:

| **Callback** | **Function** |
| --- | --- |
| `RegisterCreateCategoryCallback` | Registers a `IUnityProfilerCreateCategoryCallback` callback to get the **Profiler category** name and color whenever Unity creates a category. |
| `RegisterCreateMarkerCallback` | Registers a `IUnityProfilerCreateMarkerCallback` callback which is called whenever Unity creates a marker. Use it to get the name, Profiler category, and usage flags of a marker. The const `UnityProfilerMarkerDesc* markerDesc` parameter of the callback function represents a persistent pointer to the marker description which you can use to filter markers in `RegisterMarkerEventCallback`. |
| `RegisterMarkerEventCallback` | Registers a IUnityProfilerMarkerEventCallback callback that Unity invokes when single-shot, scoped, memory allocation, or garbage collection events happen. You can then use this callback to invoke the relevant functions in an external profiling tool. **Note:** Unity represents memory allocation events with the `GC.Alloc` marker, and garbage collection events with the `GC.Collect` markers. |
| `RegisterFrameCallback` | Encapsulates samples into logical frames so external profiling tools that don’t use frames can use these samples. Also registers a callback that the Unity Profiler runs when Unity starts the next logical CPU frame. |
| `RegisterCreateThreadCallback` | Registers a callback to get the internal thread name whenever Unity registers a thread for profiling. |

### IUnityProfilerCallbacks example

This example shows you how to pass Unity Profiler events to another profiler which has push/pop semantics. It provides two functions:

* `void MyProfilerPushMarker(const char* name)`: pushes a named marker.
* `void MyProfilerPopMarker()`: pops instrumentation marker.

The following example provides the minimal implementation required to pass begin and end instrumentation events from the Unity Profiler to the external profiler:

```
#include <IUnityInterface.h>
#include <IUnityProfilerCallbacks.h>

static IUnityProfilerCallbacks* s_UnityProfilerCallbacks = NULL;

static void UNITY_INTERFACE_API MyProfilerEventCallback(const UnityProfilerMarkerDesc* markerDesc, UnityProfilerMarkerEventType eventType, unsigned short eventDataCount, const UnityProfilerMarkerData* eventData, void* userData)
{
    switch (eventType)
    {
        case kUnityProfilerMarkerEventTypeBegin:
        {
            MyProfilerPushMarker(markerDesc->name);
            break;
        }
        case kUnityProfilerMarkerEventTypeEnd:
        {
            MyProfilerPopMarker();
            break;
        }
    }
}

static void UNITY_INTERFACE_API MyProfilerCreateMarkerCallback(const UnityProfilerMarkerDesc* markerDesc, void* userData)
{
    s_UnityProfilerCallbacks->RegisterMarkerEventCallback(markerDesc, MyProfilerEventCallback, NULL);
}

extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginLoad(IUnityInterfaces* unityInterfaces)
{
    s_UnityProfilerCallbacks = unityInterfaces->Get<IUnityProfilerCallbacks>();
    s_UnityProfilerCallbacks->RegisterCreateMarkerCallback(&MyProfilerCreateMarkerCallback, NULL);
}

extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginUnload()
{
    s_UnityProfilerCallbacks->UnregisterCreateMarkerCallback(&MyProfilerCreateMarkerCallback, NULL);
    s_UnityProfilerCallbacks->UnregisterMarkerEventCallback(NULL, &MyProfilerEventCallback, NULL);
}
```

**Note:** To unregister the given callback from all markers, run `UnregisterEventCallback` with the first parameter set to `null`.

### UnitySystracePlugin example

You can register and unregister marker callbacks dynamically, once every frame. The following example minimizes profiling overhead by enabling and disabling callbacks, depending on a third party profile state.

```
| static void UNITY_INTERFACE_API SystraceFrameCallback(void* userData)
{
    bool isCapturing = ATrace_isEnabled();
    if (isCapturing != s_isCapturing)
    {
        s_isCapturing = isCapturing;
        if (isCapturing)
        {
            s_UnityProfilerCallbacks->
              RegisterCreateMarkerCallback(SystraceCreateEventCallback, NULL);
        }
        else
        {
            s_UnityProfilerCallbacks->
              UnregisterCreateMarkerCallback(SystraceCreateEventCallback, NULL);
            s_UnityProfilerCallbacks->
              UnregisterMarkerEventCallback(NULL, SystraceEventCallback, NULL);
        }
    }
}
```

**Note:** To unregister the given callback from all markers, run `UnregisterEventCallback` with the first parameter set to `null`.

## Special markers

Unity has the following special markers that contain useful metadata:

* `Profiler.DefaultMarker`
* `GC.Alloc`

### Profiler.DefaultMarker

`Profiler.DefaultMarker` is a marker that Unity reserves for [Profiler.BeginSample](../ScriptReference/Profiling.Profiler.BeginSample.md) and [Profiler.EndSample](../ScriptReference/Profiling.Profiler.EndSample.md) events.

In the previous example, `kUnityProfilerMarkerEventTypeBegin eventType` corresponds to the `Profiler.BeginSample` event and has the following data:

* Int32: The `UnityEngine.Object` instance ID. This is 0 if the object isn’t specified.
* UInt16 array: The UTF16 string that’s passed to `Profiler.BeginSample`. The size is in bytes.
* UInt32: The category index.

### GC.Alloc

`GC.Alloc` is a marker that corresponds to garbage collection allocation. It has the following data:

* Int64: The size of the allocation.

## Additional resources

* [Native plug-ins](plug-ins-native.md)
* [Profiler introduction](profiler-introduction.md)
* [Adding profiling information to your code](profiler-adding-information-code.md)