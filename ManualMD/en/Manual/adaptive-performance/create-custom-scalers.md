# Create custom scalers

Create custom scalers to automatically adjust quality settings based on device performance.

Custom scalers allow you to extend the Adaptive Performance system with your own logic for controlling quality settings. For more information about scalers, refer to [Modifying asset quality with scalers](scalers-introduction.md).

There are two ways to create custom scalers:

* [Using a `ScriptableObject` in the Unity Editor interface](#editor-interface)
* [Using a script-only method](#script-only)

The method you choose determines how you configure the scaler and how it applies at runtime.

**Important:** These two workflows are mutually exclusive. If you add a custom scaler through the Editor interface, Unity disables the automatic runtime scan for script-only scalers. Choose one method for your project.

## Add a custom scaler in the Editor interface

Create a [`ScriptableObject`](../class-ScriptableObject.md) asset from your scaler script to configure the scaler’s properties directly in the Editor for each [scaler profile](scaler-profiles.md).

To add a custom scaler in the Editor, you need to create the custom scaler script and add the scaler asset to a scaler profile.

### Create the custom scaler script

To create the script:

1. Create a new C# script that inherits from [`AdaptivePerformanceScaler`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceScaler.md).
2. Add the [`CreateAssetMenu`](../../ScriptReference/CreateAssetMenuAttribute.md) attribute to the class. This exposes the script in the **Assets** > **Create** menu.
3. Define the scaler’s behavior by overriding methods like `OnLevel`, `OnEnabled`, and `OnDisabled`.

### Add the scaler asset to a scaler profile

To add and configure the asset:

1. In the Editor, from the main menu, go to **Assets** > **Create** and select the menu item you defined in the `CreateAssetMenu` attribute. This creates a new scaler asset file in your Project window.
2. Go to **Edit** > **Project Settings** > **Adaptive Performance**.
3. Select a scaler profile.
4. Add the asset to the custom scalers list.
5. Configure the scaler’s settings directly in the Editor. **Note**: Settings you define in the script override settings configured in the Editor. For this reason, it’s best practice to define scaler behavior in the script and all settings in the Editor.

After you complete these steps, you have a scaler asset in your **Project window** that you can add to any scaler profile and configure independently.

After you add the asset, be aware of the following behavior:

* Adaptive Performance creates a clone of your scaler asset and manages it internally within the profile. If you make changes to the original asset file in your Project window, you must remove the old scaler from the profile and add the updated asset again for the changes to take effect.
* The scaler’s name in the list is taken from the asset’s file name. You can’t add multiple scalers with the same name to a single profile.

### Example: UI-based scaler

This example demonstrates a script for the Editor interface workflow. All settings are configured in the Editor.

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;

// This attribute allows you to create an asset from this script
// via Assets > Create > Scriptable Objects > Texture Quality Scaler UI.
[CreateAssetMenu(fileName = "TextureQualityScaler_UI", menuName = "Scriptable Objects/Texture Quality Scaler UI")]
public class TextureQualityScalerUI : AdaptivePerformanceScaler
{
    private int m_DefaultTextureLimit;

    protected override void OnEnabled()
    {
        // Store the original texture quality setting when the scaler is enabled.
        m_DefaultTextureLimit = QualitySettings.globalTextureMipmapLimit;
    }

    protected override void OnDisabled()
    {
        // Restore the original setting when the scaler is disabled.
        m_DefaultTextureLimit = QualitySettings.globalTextureMipmapLimit;
    }

    // This method is called when the performance level changes.
    protected override void OnLevel()
    {
        // The base class calculates the new `Scale` value based on the current performance level.
        // Apply the new scale to the global mipmap limit.
        // A higher scale (better quality) maps to a lower mipmap limit (better quality).
        if (ScaleChanged())
        {
            // Adjust the global texture mipmap limit based on the new scale.
            Debug.Log($"TextureQualityScalerUI new scale: {Scale}");
            QualitySettings.globalTextureMipmapLimit = (int)MaxBound - ((int)(MaxBound * Scale));
        }
    }
}
```

## Create a scaler with a script only

In this workflow, you define the scaler’s settings directly in the script. At runtime, Unity scans your project for any **scripts** that inherit from `AdaptivePerformanceScaler` and adds them automatically. The settings defined in the script apply globally to all scaler profiles.

To create a script-only custom scaler:

1. Create a new C# script that inherits from [`AdaptivePerformanceScaler`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceScaler.md).
2. Make sure the class name matches the [`AdaptivePerformanceScalerSettingsBase`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceScalerSettingsBase.md) name.
3. Define the scaler’s behavior.

After you complete these steps, you have a custom scaler script in your project. When you enter Play mode, Adaptive Performance automatically detects and applies this scaler, provided you haven’t added any scalers using the Editor interface.

### Example: Script-only scaler

This example includes all settings defined in the script.

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;

public class TextureQualityScalerScriptOnly : AdaptivePerformanceScaler
{
    // Define the default settings for the scaler.
    private AdaptivePerformanceScalerSettingsBase m_AdaptiveTextureQualityScaler = new AdaptivePerformanceScalerSettingsBase
    {
        // This name must exactly match the class name.
        name = "TextureQualityScalerScriptOnly",
        enabled = false,
        scale = 1.0f,
        visualImpact = ScalerVisualImpact.High,
        target = ScalerTarget.GPU,
        minBound = 0,
        maxBound = 4,
        maxLevel = 4
    };

    private int m_DefaultTextureLimit;

    protected override void Awake()
    {
        base.Awake();
        // Apply the default settings defined above.
        ApplyDefaultSetting(m_AdaptiveTextureQualityScaler);
    }

    protected override void OnEnabled()
    {
        m_DefaultTextureLimit = QualitySettings.globalTextureMipmapLimit;
    }

    protected override void OnDisabled()
    {
        QualitySettings.globalTextureMipmapLimit = m_DefaultTextureLimit;
    }

    protected override void OnLevel()
    {
        if (ScaleChanged())
        {
            Debug.Log($"TextureQualityScalerScriptOnly new scale: {Scale}");
            QualitySettings.globalTextureMipmapLimit = (int)MaxBound - ((int)(MaxBound * Scale));
        }
    }
}
```

## Additional resources

* [Use scaler profiles](scaler-profiles.md)
* [Adaptive Performance provider settings reference](provider-settings-reference.md)
* [Apply performance optimizations](performance-optimization-strategies.md)