# Use scaler profiles

Use scaler profiles to quickly move between different performance configurations.

A scaler profile is a saved group of settings for your scalers. Using a profile lets you manage and move between different performance configurations with a single command, instead of changing each setting one-by-one in your code.

For examples of using scaler profiles, refer to the [Adaptive Performance samples](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@latest?subfolder=/manual/samples-guide.html#scaler-profiles-sample).

## Create scaler profiles

Create and configure scaler profiles in the [Adaptive Performance provider settings](provider-settings-reference.md#scaler-profiles). Only the **Default** profile is available by default.

To create a scaler profile:

1. Open the Adaptive Performance provider settings (menu: **Edit** > **Project Settings** > **Adaptive Performance**).
2. Select your provider.
3. In the **Scaler Profiles** section, select **Add New Scaler Profile** to create a new profile.
4. Your new profile appears in the list. Configure the settings as required.

## Delete scaler profiles

To delete a scaler profile:

1. Open the Adaptive Performance provider settings (menu: **Edit** > **Project Settings** > **Adaptive Performance**).
2. In the **Scaler Profiles** section, find the profile you want to delete.
3. Open the **More** (⋮) menu for the profile.
4. Select **Remove**.

## Load scaler profiles at runtime

The first profile in your list of scaler profiles loads automatically when the application starts. To change to a different profile during runtime, you must load it through a script.

**Important**: Loading scaler profiles is resource-intensive. To prevent performance issues, only load profiles when your application’s in a loading state.

To load scaler profiles at runtime:

1. Query for all profiles created and available during runtime using [`AdaptivePerformanceGeneralSettings.Instance`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceGeneralSettings.Instance.md).
2. Call [`LoadScalerProfile`](../../ScriptReference/AdaptivePerformance.IAdaptivePerformanceSettings.LoadScalerProfile.md) and provide the scaler profile name as a string to load.

For example:

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;
using System.Collections.Generic;

public class ScalerProfileLoader : MonoBehaviour
{
    public void LoadSpecificScalerProfile(int profileIndex)
    {
        // Retrieve the Adaptive Performance settings from the active loader.
        IAdaptivePerformanceSettings settings = AdaptivePerformanceGeneralSettings.Instance.Manager.activeLoader.GetSettings();
        // Check if settings were successfully retrieved.
        if (settings == null)
        {
            return;
        }

        // Get the list of all available scaler profiles.
        List<string> scalerProfiles = settings.GetAvailableScalerProfiles();

        // Check if any scaler profiles are available and the index is valid.
        if (profileIndex >= 0 && profileIndex < scalerProfiles.Count)
        {
            settings.LoadScalerProfile(scalerProfiles[profileIndex]);
        }
    }
}
```

## Additional resources

* [Modifying asset quality with scalers](scalers.md)
* [Create custom scalers](create-custom-scalers.md)
* [Adaptive Performance provider settings reference](provider-settings-reference.md)
* [Apply performance optimizations](performance-optimization-strategies.md)