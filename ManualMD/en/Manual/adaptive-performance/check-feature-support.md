# Check for feature support

Verify if your device supports a specific Adaptive Performance capability.

Adaptive Performance comes with different features. You can check which features a device supports using [`IAdaptivePerformance.SupportedFeature`](../../ScriptReference/AdaptivePerformance.IAdaptivePerformance.SupportedFeature.md). For available provider features, refer to the [`Feature`](../../ScriptReference/AdaptivePerformance.Provider.Feature.md) enumeration.

The following example checks for the [`Feature.ClusterInfo`](../../ScriptReference/AdaptivePerformance.Provider.Feature.ClusterInfo.md):

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;
using UnityEngine.AdaptivePerformance.Provider;

public class AdaptivePerformanceFeatureChecker : MonoBehaviour
{
    // Declare 'ap' as a private class-level field to hold the IAdaptivePerformance instance.
    private IAdaptivePerformance ap;

    void Start()
    {
        // Get the Adaptive Performance instance.
        ap = Holder.Instance;

        // Check if Adaptive Performance is active.
        if (ap == null || !ap.Active)
        {
            Debug.Log("[AP ClusterInfo] Adaptive Performance not active.");
            return;
        }
        // Check if the ClusterInfo feature is supported by the active provider.
        if (!ap.SupportedFeature(Feature.ClusterInfo))
        {
            Debug.Log("[AP ClusterInfo] Feature not supported.");
        }
        // Access ClusterInfo metrics if the feature is supported.
        var clusterInfo = ap.PerformanceStatus.PerformanceMetrics.ClusterInfo;
    }
}
```

## Additional resources

* [Scripting with Adaptive Performance](scripting-api.md)
* [Introduction to providers](providers-introduction.md)
* [Track timing and thermal data](track-timing-thermal-data.md)