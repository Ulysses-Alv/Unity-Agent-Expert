# Build for QNX from the command line

To build a Unity project for the QNX system on the command line interface (CLI), you must have the Unity Editor installed on the build host. The build host must be running a Linux operating system and also have the QNX SDK installed. For more information, refer to [Install the Unity Editor for QNX](qnx-install-editor.md).

The following example uses the Linux variant of the executable (`Unity`), but you can replace it with an equivalent executable for your preferred build host’s operating system.

## Command line arguments

**Prerequisites**

Make sure you’ve followed these instructions before building the project using the command line:

* Set up `QNX_TARGET` and `QNX_HOST` as per [Install the Unity Editor for QNX](qnx-install-editor.md).

To instruct Unity to start in CLI mode and build the project at `<path-to-unity-project-root>` for QNX, run the following command:

```
Unity -quit -batchmode -nographics -buildTarget QNX -executeMethod Builder.Build -projectPath <path-to-unity-project-root>
```

The build process also calls the function `Builder.Build` to continue with the build setup.

## Build script

You can add the example build script into `Assets/Editor/` for the project to you’re building from the command line. Use the `-executeMethod` option to call the `Build()` method of this class, which sets up the build options and triggers the build.

For more information, refer to the `BuildPipeline.BuildPlayer` [documentation](../ScriptReference/BuildPipeline.BuildPlayer.md).

```
using UnityEditor;
using UnityEditor.Build.Reporting;
using UnityEngine;

public class Builder
{
    private static void BuildQNX(QNXOsVersion qnxVersion, QNXArchitecture architecture)
    {
        // Set QNX version in BuildSettings
        EditorUserBuildSettings.selectedQnxOsVersion = qnxVersion;

        // Set architecture in BuildSettings
        EditorUserBuildSettings.selectedQnxArchitecture = architecture;

        // Setup build options (e.g. scenes, build output location)
        var options = new BuildPlayerOptions
        {
            // Change to scenes from your project
            scenes = new[]
            {
                "Assets/Scenes/SampleScene.unity",
            },
            // Change to location the output should go
            locationPathName = "../QNXPlayer/",
            options = BuildOptions.CleanBuildCache | BuildOptions.StrictMode,
            target = BuildTarget.QNX
        };

        var report = BuildPipeline.BuildPlayer(options);

        if (report.summary.result == BuildResult.Succeeded)
        {
            Debug.Log($"Build successful - Build written to {options.locationPathName}");
        }
        else if (report.summary.result == BuildResult.Failed)
        {
            Debug.LogError($"Build failed");
        }
    }

    // This function will be called from the build process
    public static void Build()
    {
        // Build QNX 7.1 ARM64 Unity player
        BuildQNX(QNXOsVersion.Neutrino71, QNXArchitecture.Arm64);
    }
}
```

## Additional resources

* [Install the Unity Editor for QNX](qnx-install-editor.md)
* [QNX build settings reference](qnx-build-settings.md)