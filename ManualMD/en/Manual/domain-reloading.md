# Enter Play mode with domain reload disabled

The scripting domain, also known as the [application domain](https://learn.microsoft.com/en-us/dotnet/framework/app-domains/application-domains), or simply domain, is a core feature of Unity’s managed scripting environment. The domain is an isolated section of memory dedicated to a particular application, which contains the compiled types required by the application, grouped into logical units called [assemblies](assembly-definition-files.md). It also contains the data that represents the current application state, such as the values of variables and object references in the various sections of [managed memory](performance-managed-memory.md).

By default Unity reloads the domain on entering Play mode to reset the application state. Resetting state before entering Play mode is often desirable so your application starts up as it would at the beginning of a new build. For example, static counters that were incremented in a previous Play mode session should begin from zero again in the next one. However, domain reload is also a time-consuming operation that negatively impacts iteration times when you frequently switch between Edit and Play mode. For faster iteration times, you can [disable domain reload](configurable-enter-play-mode.md#configure-play-mode) on entering Play mode but you must then manually [reset state in your code](#reset-state).

**Note**: Unity also performs domain reload as part of an asset database refresh when it detects changes to **scripts**. This still happens even when domain reload on entering Play mode is disabled. For more information on when and how asset database refreshes happen, refer to [Refreshing the Asset Database](AssetDatabaseRefreshing.md).

## Effects of disabling domain reload when entering Play mode

When you disable domain reload:

* Non-serialized fields keep the values assigned to them during Play mode on returning to Edit mode. This applies for fields of all script types, including MonoBehaviours (including those on prefab assets), ScriptableObjects, and your own custom C# types. For detailed information on what is and isn’t serialized in different contexts, refer to [Serialization rules](script-serialization-rules.md).
* Static variables keep their values between Play mode sessions.
* Static events keep their registered subscribers between Play mode sessions.
* There are no additional `OnDisable` or `OnEnable` calls for scripts marked with the [`[ExecuteInEditMode]`](../ScriptReference/ExecuteInEditMode.md) or [`[ExecuteAlways]`](../ScriptReference/ExecuteAlways.md).

To compensate for this persistence of data between Play mode sessions and enter Play mode with a fresh application state, you must [reset state in your code](#reset-state).

For more information on the effects of disabling domain and **scene** reload, refer to [Details of disabling domain and scene reload](configurable-enter-play-mode-details.md).

## Resetting state from code

When domain reloading is disabled, the values of [static fields](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/static-classes-and-static-class-members#static-members) and the handlers assigned to static events persist between Play mode runs. The following code example has a static counter that increments on a press of any keyboard key. The code also registers a method to handle the static event [`EditorApplication.playModeStateChanged`](../ScriptReference/EditorApplication-playModeStateChanged.md).

With domain reload enabled, Unity reinitializes this code on entering Play mode, erasing the state from the previous Play mode run, including the counter value and the registration of the event handler. With domain reload disabled, the counter value and the event handler registration are both preserved from the previous run. On the next run of Play mode, the counter begins with the value it had at the end of the previous run and the event handler method is called multiple times on one occurrence of the event, generating multiple “Exiting Play mode!” messages in the console.

```
// Copy-paste this code into a MonoBehaviour script attached to a GameObject in your project.
// Run it with domain reload enabled and then with domain reload disabled and note the different behavior.

using UnityEngine;
#if UNITY_EDITOR
using UnityEditor;
#endif

public class StateResetExample : MonoBehaviour
{
    // With domain reload disabled this counter won't reset to zero on exiting Play mode
    static int counter = 0;

    void Start()
    {
    // Register handler
#if UNITY_EDITOR
        EditorApplication.playModeStateChanged += OnExitPlayMode;
#endif    
    }

    void Update()
    {
        if (Input.anyKeyDown)
        {
            counter++;
            Debug.Log("Counter: " + counter);
        }
    }
#if UNITY_EDITOR
    private static void OnExitPlayMode(PlayModeStateChange state)
    {
        if(state == PlayModeStateChange.ExitingPlayMode)
        {
            // With domain reload disabled this message prints multiple times after the first Play mode run
            Debug.Log("Exiting Play mode!");
        }
    }
#endif
}
```

You can fix the problem behavior with code that explicitly resets the counter and unregisters the event handler between Play mode runs. You can either do this on entering Play mode or on exiting Play mode.

### Resetting state on exiting Play mode

It’s often most efficient to reset state on exiting Play mode rather than on entering. You can use the [`EditorApplication.playModeStateChanged`](../ScriptReference/EditorApplication-playModeStateChanged.md) event and its [`ExitingPlayMode`](../ScriptReference/PlayModeStateChange.ExitingPlayMode.md) enum value to catch the Play mode exit event and reset state at that point. In the following example the event handler for Play mode exit is used to unregister itself, but you can also use this method to unregister any other static event handlers at this point:

```
using UnityEngine;
#if UNITY_EDITOR
using UnityEditor;
#endif

public class StateResetOnExit : MonoBehaviour
{
    static int counter = 0;

    void Start()
    {
#if UNITY_EDITOR
        EditorApplication.playModeStateChanged += OnExitPlayMode;
#endif        
    }

    // Update is called once per frame
    void Update()
    {
        if (Input.anyKeyDown)
        {
            counter++;
            Debug.Log("Counter: " + counter);
        }
    }

#if UNITY_EDITOR
    private static void OnExitPlayMode(PlayModeStateChange state)
    {
        if(state == PlayModeStateChange.ExitingPlayMode)
        {
            Debug.Log("Exiting Play mode!");
            Debug.Log("Unregistering handler.");
            // Unregister the handler so it doesn't affect the next Play mode run
            EditorApplication.playModeStateChanged -= OnExitPlayMode;
            Debug.Log("Resetting counter.");
            // Reset the counter so it starts from 0 on the next Play mode run
            counter = 0;
        }
    }
#endif
}
```

[`EditorApplication.playModeStateChanged`](../ScriptReference/EditorApplication-playModeStateChanged.md) is in the `UnityEditor` namespace so it only works in Play mode within the Unity Editor and not in a standalone Player build.

If your code executes in Edit mode in addition to Play mode, you can’t rely on resetting state on exiting Play mode. Your code might modify a static variable while in Edit mode, so you must reset the variable [on entering Play mode](#reset-on-enter) instead.

**Note**: For scripts that execute in Edit mode, disabling domain reload skips [`MonoBehaviour.OnDisable`](../ScriptReference/MonoBehaviour.OnDisable.md) and disabling scene reload skips [`MonoBehaviour.OnDestroy`](../ScriptReference/MonoBehaviour.OnDestroy.md), which makes these methods inappropriate for resetting state in such scripts. For more information, refer to [Details of disabling domain and scene reload](configurable-enter-play-mode-details.md).

### Resetting state on entering Play mode

You can also reset state on entering Play mode. This can be a useful alternative if you encounter any platform-specific problems with capturing the exiting Play mode event, or if you have code that also executes in Edit mode. For the purposes of illustration, the following example demonstrates both resetting static variables and unregistering static event handlers on entering Play mode. Best practice is to always unregister static event handlers [on exiting Play mode](#reset-on-exit). Waiting until the next Play mode session can cause object reference issues that lead to unregistering the wrong handler.

Unity has custom initialization attributes in both the `UnityEngine` and `UnityEditor` namespaces, which you can use to perform initialization work, including manual state reset. Which attributes you should use depends on whether your code runs in Edit mode or Play mode. To reset state on entering Play mode for your Player (runtime) scripts, you can use the [`[RuntimeInitializeOnLoadMethod]`](../ScriptReference/RuntimeInitializeOnLoadMethodAttribute.md) attribute with the [`RuntimeInitializeLoadType.SubsystemRegistration`](../ScriptReference/RuntimeInitializeLoadType.SubsystemRegistration.md) parameter:

```
using UnityEngine;
#if UNITY_EDITOR
using UnityEditor;
#endif

public class StateResetOnEnter : MonoBehaviour
{
    static int counter = 0;

    [RuntimeInitializeOnLoadMethod(RuntimeInitializeLoadType.SubsystemRegistration)]
    static void Init()
    {
        Debug.Log("Unregistering handler.");
        // Unregister the handler so it doesn't affect the next Play mode run
        EditorApplication.playModeStateChanged -= OnExitPlayMode;
        Debug.Log("Resetting counter.");
        // Reset the counter so it starts from 0 on the next Play mode run
        counter = 0;  
    }

    void Start()
    {
#if UNITY_EDITOR
        EditorApplication.playModeStateChanged += OnExitPlayMode;
#endif
    }

    // Update is called once per frame
    void Update()
    {
        if (Input.anyKeyDown)
        {
            counter++;
            Debug.Log("Counter: " + counter);
        }
    }
#if UNITY_EDITOR
    private static void OnExitPlayMode(PlayModeStateChange state)
    {
        if(state == PlayModeStateChange.ExitingPlayMode)
        {
            Debug.Log("Exiting Play mode!");
        }
    }
#endif
}
```

[`[RuntimeInitializeOnLoadMethod]`](../ScriptReference/RuntimeInitializeOnLoadMethodAttribute.md) is in the `UnityEngine` namespace so it’s only available in Play mode. For Editor-only scripts such as custom Editor windows or **Inspectors** that use statics, you can do one of the following:

* Use [`[InitializeOnLoadAttribute]`](../ScriptReference/InitializeOnLoadAttribute.md) to register a handler for [`EditorApplication.playModeStateChanged`](../ScriptReference/EditorApplication-playModeStateChanged.md) and then perform state cleanup on either entering or exiting Play mode. Refer to the [`EditorApplication.playModeStateChanged`](../ScriptReference/EditorApplication-playModeStateChanged.md) API reference for an example.
* Use the [`[InitializeOnEnterPlayMode]`](../ScriptReference/InitializeOnEnterPlayModeAttribute.md) attribute to perform cleanup specifically on entering Play mode.

**Note**: While the previous code examples use MonoBehaviour classes, you can also use these initialization attributes on scripts of any kind, including those that inherit from ScriptableObject or your own custom C# types.

## Additional resources

* [Project Settings window](comp-ManagerGroup.md)
* [Editor Project Settings](class-EditorManager.md)
* [Toolbar](Toolbar.md)
* [Scene Reloading](scene-reloading.md)
* [Configurable Enter Play Mode](configurable-enter-play-mode.md)