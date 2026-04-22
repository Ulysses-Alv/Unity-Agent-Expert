# Scripting with Adaptive Performance

Control and monitor device performance at runtime using the Adaptive Performance scripting API.

After you [set up Adaptive Performance](initial-setup.md), you can use its scripting API to access and control device performance metrics directly from your code. The central point of access is the [`IAdaptivePerformance`](../../ScriptReference/AdaptivePerformance.IAdaptivePerformance.md) interface, which Unity automatically creates and manages on a **GameObject** at runtime.

For a code-free approach using visual graphs, refer to the [Adaptive Performance package documentation about visual scripting](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@latest?subfolder=/manual/visual-scripting.html).

## Accessing the instance

To access the instance, use [`AdaptivePerformance.Holder.Instance`](../../ScriptReference/AdaptivePerformance.Holder.Instance.md).

To check if your device supports Adaptive Performance, use the [`IAdaptivePerformance.Active`](../../ScriptReference/AdaptivePerformance.IAdaptivePerformance.Active.md) property.

## Debug logging

To get detailed information during runtime, enable **Logging** in the [provider settings](provider-settings-reference.md) or via [`IDevelopmentSettings.Logging`](../../ScriptReference/AdaptivePerformance.IDevelopmentSettings.Logging.md) during runtime or using boot time flags from [`IAdaptivePerformanceSettings`](../../ScriptReference/AdaptivePerformance.IAdaptivePerformanceSettings.md):

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;

static class AdaptivePerformanceConfig
{
    // This method runs automatically after assemblies are loaded, before any scene starts.
    [RuntimeInitializeOnLoadMethod(RuntimeInitializeLoadType.AfterAssembliesLoaded)]
    static void Setup()
    {
        // Retrieve the Adaptive Performance settings.
        IAdaptivePerformanceSettings settings = AdaptivePerformanceGeneralSettings.Instance.Manager.activeLoader.GetSettings();
        // Check if settings were successfully retrieved before modifying.
        if (settings != null)
        {
            // Enable Adaptive Performance logging.
            settings.logging = true;
        }
    }
}
```

## Additional resources

* [Set up Adaptive Performance](initial-setup.md)
* [Track timing and thermal data](track-timing-thermal-data.md)