# Enter Play mode with scene reload disabled

Scene reload on entering Play mode is enabled by default. This means that when you enter Play mode, Unity destroys all existing **scene** **GameObjects** and reloads the scene from disk. As your project gets more complex, the time between pressing the Play button and the scene fully loading in the Editor increases.

When you disable scene reloading, the process takes less time. Instead of reloading the scene from disk, Unity only resets the scene’s modified contents. Unity still calls the same [event functions](event-functions.md) (such as `OnEnable`, `OnDisable` and `OnDestroy`) as if the scene were freshly loaded.

## Effects of disabling scene reload when entering Play mode

When you [disable scene reload](configurable-enter-play-mode.md#configure-play-mode), the time it takes to start your application in the Editor is no longer representative of the startup time in the built version. If you want to debug or profile exactly what happens during your project’s startup, you should enable scene reloading to more accurately represent the true loading time and processes that happen in the built version of your application.

Otherwise, disabling scene reload should have minimal effects on your project. However, because scene reloading is closely connected to domain reload, there are a few important differences:

* Unity doesn’t recreate existing objects or call constructors, which means non-serialized fields keep the values assigned to them during Play mode on returning to Edit mode. This applies for fields of all script types, including MonoBehaviours, ScriptableObjects, and your own custom C# types. For detailed information on what is and isn’t serialized in different contexts, refer to [Serialization rules](script-serialization-rules.md).
  + **Note**: `private` fields are not serialized as part of the regular build pipeline but **are** serialized as part of the Editor’s [hot reloading of scripts](script-serialization-how-unity-uses.md#hot-reload). This is why `private` fields you modify in Play mode might reset to their original values on exiting Play mode even when scene and domain reload on enter Play mode are disabled.
* Unity converts null `private` and `internal` fields of array and `List` type to an empty array or `List` object during domain reload and they stay non-null for runtime (non-Editor) **scripts**.
* Scripts decorated with [`[ExecuteInEditMode]`](../ScriptReference/ExecuteInEditMode.md) or [`[ExecuteAlways]`](../ScriptReference/ExecuteAlways.md) don’t receive `OnDestroy` or `Awake` calls. While in Edit mode, these scripts might modify their own fields or those of other runtime scripts. To mitigate this, you can initialize any affected fields in an `OnEnable` callback with code inside a condition that checks the value of [`EditorApplication.isPlaying`](../ScriptReference/EditorApplication.isPlaying.md). For an example of this and more context on the importance of separating Play mode and Edit mode code, refer to the [`[ExecuteAlways]`](../ScriptReference/ExecuteAlways.md) API description.

For more details on the events skipped with scene reload disabled, refer to [Details of disabling domain and scene reload](configurable-enter-play-mode-details.md).

## Additional resources

* [Configuring code reload on entering Play mode](configurable-enter-play-mode.md)
* [Enter Play mode with domain reload disabled](domain-reloading.md)
* [Details of disabling domain and scene reload](configurable-enter-play-mode-details.md)