# Create a custom provider

Create and configure a custom provider for Adaptive Performance.

Adaptive Performance relies on [providers](providers.md) to act as a bridge between the Unity Engine and a device’s hardware sensors. Creating a custom provider involves building a self-contained package that’s discoverable by the Adaptive Performance system.

Follow these steps to create a new provider:

1. [Create a loader.](#create-loader)
2. [Configure build and runtime settings.](#configure-settings)
3. [Provide package metadata so Unity can recognize your provider.](#package-metadata)

## Create a loader

The [`AdaptivePerformanceManagerSettings`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceManagerSettings.md) class manages the lifecycle of your provider’s subsystems. Your provider needs a loader, which acts as the main driver for these management features.

To create a loader:

1. Create a subclass of [`AdaptivePerformanceLoader`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceLoader.md). This class is a [`ScriptableObject`](../class-ScriptableObject.md) that defines the subsystems your provider requires and their load order. To handle subsystem management in a type-safe manner, derive from the [`AdaptivePerformanceLoaderHelper`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceLoaderHelper.md) class instead.
2. Implement the lifecycle methods. `AdaptivePerformanceManagerSettings` calls these methods to initialize, start, stop, and deinitialize (clean up) your loader’s subsystems. For example:

   ```
   using UnityEngine;
   using UnityEngine.AdaptivePerformance;

   public abstract class AdaptivePerformanceLoader : ScriptableObject
   {
       public virtual bool Initialize() { return false; }

       public virtual bool Start() { return false; }

       public virtual bool Stop() { return false; }

       public virtual bool Deinitialize() { return false; }

       public abstract T GetLoadedSubsystem<T>() where T : class, ISubsystem;

       public abstract ISubsystem GetDefaultSubsystem();

       public abstract IAdaptivePerformanceSettings GetSettings();
   }
   ```
3. Add the loader to the manager settings. In the Unity Editor, add your `AdaptivePerformanceLoader` instance to the **Loaders** list in the **Adaptive Performance Manager Settings**. You can add multiple loaders and arrange them by priority. At runtime, the manager attempts to initialize each loader in order. The first one that initializes successfully becomes the active loader.
4. (Optional) Query the static [`AdaptivePerformanceManagerSettings.ActiveLoader`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceManagerSettings-activeLoader.md) instance to access the active loader. If all loaders fail to initialize, Unity sets `activeLoader` to null.

## Configure build and runtime settings

A provider might need additional settings to help manage build issues or runtime configuration.

To configure build and runtime settings through **Project Settings**:

1. Add the [`AdaptivePerformanceConfigurationData`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceConfigurationDataAttribute.md) attribute to a `ScriptableObject`.
2. Define properties in the `ScriptableObject` to control the provider’s configuration. Unity displays these options in the **Adaptive Performance** section of the **Project Settings** window.
     
   **Note**: Unity manages the lifecycle of one instance of the class marked with the attribute through the [`EditorBuildSettings`](../../ScriptReference/EditorBuildSettings.md) config object API. If you don’t provide a dedicated interface, configuration settings display in the **Project Settings** window using the standard **Scriptable Object** **Inspector**. You can create a custom **Editor** for your configuration settings type to replace the standard Inspector in the **Project Settings** window.
3. Inherit from [`AdaptivePerformanceBuildHelper<T>`](../../ScriptReference/AdaptivePerformance.Editor.AdaptivePerformanceBuildHelper_1.md) in a build script to pass settings from the Editor into your runtime loader. The following is an example of a simple build script:

   ```
   using UnityEngine.AdaptivePerformance.Editor;
   using UnityEngine;

   public class MySettings : ScriptableObject { }

   public class MyBuildProcessor : AdaptivePerformanceBuildHelper<MySettings>
   {
       public override string BuildSettingsKey { get { return "MyPackageSettingsKey"; } }
   }
   ```

### Customize the build process

You can also override build processing steps like [`OnPreprocessBuild`](../../ScriptReference/AdaptivePerformance.Editor.AdaptivePerformanceBuildHelper_1.OnPreprocessBuild.md) and [`OnPostprocessBuild`](../../ScriptReference/AdaptivePerformance.Editor.AdaptivePerformanceBuildHelper_1.OnPostprocessBuild.md). Make sure to call the base class implementation to ensure your settings transfer to the built application. For example:

```
using UnityEditor.Build.Reporting;
using UnityEngine.AdaptivePerformance.Editor;
using UnityEngine;

public class MySettings : ScriptableObject { }
public class MyBuildProcessor : AdaptivePerformanceBuildHelper<MySettings>
{
    public override string BuildSettingsKey { get { return "MyPackageSettingsKey"; } }

    public override void OnPreprocessBuild(BuildReport report)
    {
        base.OnPreprocessBuild(report);
        // Do your work here
    }

    public override void OnPostprocessBuild(BuildReport report)
    {
        base.OnPreprocessBuild(report);
        // Do your work here
    }
}
```

If you want to support different settings per platform at build time, you can override [`AdaptivePerformanceBuildHelper<T0>.SettingsForBuildTargetGroup`](../../ScriptReference/AdaptivePerformance.Editor.AdaptivePerformanceBuildHelper_1.SettingsForBuildTargetGroup.md) and use the [`buildTargetGroup`](../../ScriptReference/BuildTargetGroup.md) attribute to retrieve the appropriate platform settings. By default, this method uses the key associated with the settings instance to copy the entire settings object from [`EditorUserBuildSettings`](../../ScriptReference/EditorUserBuildSettings.md) to [`PlayerSettings`](../../ScriptReference/PlayerSettings.md).

```
using UnityEditor;
using UnityEngine;
using UnityEngine.AdaptivePerformance.Editor;

public class MySettings : ScriptableObject { }
public class MyBuildProcessor : AdaptivePerformanceBuildHelper<MySettings>
{
    public override string BuildSettingsKey { get { return "MyPackageSettingsKey"; } }

    public override UnityEngine.Object SettingsForBuildTargetGroup(BuildTargetGroup buildTargetGroup)
    {
        // Get platform specific settings and return them. Use something like the following
        // for simple settings data that isn't platform specific.
        UnityEngine.Object settingsObj = null;
        EditorBuildSettings.TryGetConfigObject(BuildSettingsKey, out settingsObj);
        if (settingsObj == null || !(settingsObj is T))
            return null;

        return settingsObj;
    }
}
```

## Provide package metadata

For the Adaptive Performance provider management system to use your provider, you must supply metadata.

To do this, create a class that implements the following interfaces:

* [`IAdaptivePerformancePackage`](../../ScriptReference/AdaptivePerformance.Editor.Metadata.IAdaptivePerformancePackage.md)
* [`IAdaptivePerformancePackageMetadata`](../../ScriptReference/AdaptivePerformance.Editor.Metadata.IAdaptivePerformancePackageMetadata.md)
* [`IAdaptivePerformanceLoaderMetadata`](../../ScriptReference/AdaptivePerformance.Editor.Metadata.IAdaptivePerformanceLoaderMetadata.md)

The system uses .NET reflection to find all types implementing the `IAdaptivePerformancePackage` interface. It then attempts to instantiate each one and populate the metadata store with the information provided by each instance.

### Package initialization

Implementing the package metadata allows the Adaptive Performance provider management system to automatically create and initialize your loaders and settings instances. The system passes any new instances of your settings to the [`PopulateNewSettingsInstance`](../../ScriptReference/AdaptivePerformance.Editor.Metadata.IAdaptivePerformancePackage.PopulateNewSettingsInstance.md) method to allow your provider to initialize the new instance data after it’s created, if needed.

### Example of a minimal package setup with metadata

```
using System.Collections.Generic;
using UnityEditor;
using UnityEngine;
using UnityEngine.AdaptivePerformance;
using UnityEngine.AdaptivePerformance.Editor.Metadata;

// The following are minimal placeholder classes required for this snippet to compile.
// In your project, 'MyPackageSettings' and 'MyLoader' would contain your actual custom logic and definitions.

public class MyPackageSettings : ScriptableObject { }

public class MyLoader : AdaptivePerformanceLoader
{
    public override bool Initialize() { return true; }
    public override bool Start() { return true; }
    public override bool Stop() { return true; }
    public override bool Deinitialize() { return true; }
    public override T GetLoadedSubsystem<T>() { return null; }
    public override ISubsystem GetDefaultSubsystem() { return null; }
    public override IAdaptivePerformanceSettings GetSettings() { return null; }
}

class MyPackage : IAdaptivePerformancePackage
{
    private class MyLoaderMetadata : IAdaptivePerformanceLoaderMetadata
    {
        public string loaderName { get; } = "My Loader";
        public string loaderType { get; } = typeof(MyLoader).AssemblyQualifiedName;
        public List<BuildTargetGroup> supportedBuildTargets { get; } = new List<BuildTargetGroup>()
        {
            BuildTargetGroup.Standalone,
            BuildTargetGroup.Android,
            BuildTargetGroup.iOS
        };
    }

    private class MyPackageMetadata : IAdaptivePerformancePackageMetadata
    {
        public string packageName { get; } = "My AdaptivePerformance Provider";
        public string packageId { get; } = "my.adaptiveperformance.package";
        public string settingsType { get; } = typeof(MyPackageSettings).AssemblyQualifiedName;
        public List<IAdaptivePerformanceLoaderMetadata> loaderMetadata { get; } = new List<IAdaptivePerformanceLoaderMetadata>()
        {
            new MyLoaderMetadata()
        };
    }

    private static IAdaptivePerformancePackageMetadata s_Metadata = new MyPackageMetadata();

    public IAdaptivePerformancePackageMetadata metadata => s_Metadata;

    public bool PopulateNewSettingsInstance(ScriptableObject obj)
    {
        MyPackageSettings packageSettings = obj as MyPackageSettings;
        if (packageSettings != null)
        {
            return true;
        }
        return false;
    }
}
```

## Provider settings in the Unity Editor

You can display provider settings under both **Project Settings** and ****Build Profiles****, so you can edit the settings in both locations. Provide the custom UI code by inheriting from the [`ProviderSettingsEditor`](../../ScriptReference/AdaptivePerformance.Editor.ProviderSettingsEditor.md) and adding a custom attribute.

The following example demonstrates how provider settings function if only the default UI is necessary.

**Note**: [`ShowTargetGroupSelection`](../../ScriptReference/AdaptivePerformance.Editor.ProviderSettingsEditor.ShowTargetGroupSelection.md) hides the platform selection group tab if you’re using the default [`DisplayProviderSettings()`](../../ScriptReference/AdaptivePerformance.Editor.ProviderSettingsEditor.DisplayProviderSettings.md) function.

```
using UnityEditor;
using UnityEditor.Localization;
using UnityEngine;
using UnityEngine.AdaptivePerformance.Editor;

public class BasicProviderSettings : ScriptableObject { }

[CustomEditor(typeof(BasicProviderSettings))]
internal class BasicProviderSettingsEditor : ProviderSettingsEditor
{
    protected override BuildTargetGroup CurrentTargetGroup => BuildTargetGroup.Unknown;
    public override bool ShowTargetGroupSelection => false;
    public override string UnsupportedInfo => L10n.Tr("Adaptive Performance Basic provider is not supported on this platform");

    protected override bool IsAutoPerformanceModeAvailable => false;
    protected override bool IsBoostAvailable => false;
    protected override bool IsThermalActionDelayAvailable => false;

    public override void OnInspectorGUI()
    {
        DisplayProviderSettings();
    }
}
```

## Additional resources

* [Introduction to providers](providers-introduction.md)
* [Provider settings reference](provider-settings-reference.md)
* [Provider lifecycle management](provider-lifecycle.md)
* [Basic provider](basic-provider.md)