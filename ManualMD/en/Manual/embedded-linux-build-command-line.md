# Build for Embedded Linux from the command line

To build a Unity project for the Embedded Linux system on the command line interface (CLI), you must have the Unity Editor installed on the build host. The build host can be a Linux, Windows, or macOS machine.

The following example uses the Linux variant of the executable (`Unity`), but you can replace it with an equivalent executable for your preferred build host’s operating system.

## Command line arguments

To instruct Unity to start in CLI mode and build the project at `<path-to-unity-project-root>` for Embedded Linux, run the following command:

```
Unity -quit -batchmode -nographics -buildTarget EmbeddedLinux -executeMethod Builder.Build -projectPath <path-to-unity-project-root>
```

The build process also calls the function `Builder.Build` to continue with the build setup process.

## Build script

You can put the example build script `Assets/Editor/` for the project to be built from the command line. Use the `-executeMethod` option to call the `Build()` method of this class, which sets up the build options and triggers the build.

For more information, refer to the [`BuildPipeline.BuildPlayer` API documentation](../ScriptReference/BuildPipeline.BuildPlayer.md).

```
using UnityEditor;
using UnityEditor.Build.Reporting;
using UnityEngine;

public class Builder
{
    private static void BuildEmbeddedLinux(EmbeddedLinuxArchitecture architecture)
    {
        // Set architecture in BuildSettings
        EditorUserBuildSettings.selectedEmbeddedLinuxArchitecture = architecture;

        // Setup build options (e.g. scenes, build output location)
        var options = new BuildPlayerOptions
        {
            // Change to scenes from your project
            scenes = new[]
            {
                "Assets/Scenes/SampleScene.unity",
            },
            // Change to location the output should go
            locationPathName = "../EmbeddedLinuxPlayer/",
            options =  BuildOptions.CleanBuildCache | BuildOptions.StrictMode,
            target = BuildTarget.EmbeddedLinux
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
        // Build EmbeddedLinux ARM64 Unity player
        BuildEmbeddedLinux(EmbeddedLinuxArchitecture.Arm64);
    }
}
```

## Additional resources:

* [BuildPipeline.BuildPlayer](../ScriptReference/BuildPipeline.BuildPlayer.md)