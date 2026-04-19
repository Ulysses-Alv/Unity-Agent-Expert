# Provider lifecycle management

Learn how Adaptive Performance manages the lifecycle of its providers automatically by integrating with Unity’s **standard event** functions.

The [`AdaptivePerformanceManagerSettings`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceManagerSettings.md) `ScriptableObject` uses an [`AdaptivePerformanceLoader`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceLoader.md) to automatically initialize, start, stop, and deinitialize provider subsystems. This process requires no manual intervention and operates in both scene-based and application-based contexts.

## Scene-based lifecycle management

Scene-based automatic lifecycle management ties the provider to a specific [`MonoBehaviour`](../../ScriptReference/MonoBehaviour.md) in a **scene**, making it active only when that scene loads.

Scene-based management hooks into the following `MonoBehaviour` callback points:

| **Callback** | **Lifecycle step** |
| --- | --- |
| `OnEnable` | Find the first loader that initialized successfully and set `ActiveLoader`. |
| `Start` | Start all subsystems. |
| `OnDisable` | Stop all subsystems. |
| `OnDestroy` | Deinitialize all subsystems and remove the `ActiveLoader` instance. |

## Application-based lifecycle management

Application-based automatic lifecycle management uses a persistent, global instance to keep the provider active for the entire duration of the application, from launch until it’s closed.

Application-based management hooks into the following callback points:

| **Callback** | **Lifecycle step** |
| --- | --- |
| Runtime initialization after assemblies loaded | Find the first loader that succeeds initialization and set `ActiveLoader`. |
| Runtime initialization before splash screen displays | Start all subsystems. |
| `OnDisable` | Stop all subsystems. |
| `OnDestroy` | Deinitialize all subsystems and remove the `ActiveLoader` instance. |

## Additional resources

* [Introduction to providers](providers-introduction.md)
* [Create a custom provider](create-custom-provider.md)
* [Provider settings reference](provider-settings-reference.md)
* [Basic provider](basic-provider.md)