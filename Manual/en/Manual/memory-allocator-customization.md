# Customizing native memory allocators

To customize allocator settings you can either edit the configurable values through the [Unity Editor](#use-the-editor) or supply them as [command line arguments](#use-command-line-arguments).

Unity stores allocator settings in `MemorySettings.asset` and applies them during the build process. This means new settings take effect at every build, for both Editor and Player builds.

The allocator settings are stored in the `ProjectSettings` folder and are updated every time Unity imports or changes `MemorySettings.asset`. New values for the Editor only take effect on the next Editor startup.

## Using the Editor

1. Open the **Project Settings** window (**Edit** > **Project Settings**).
2. Select the **Memory Settings** panel.
3. Select the lock icon next to the value you want to edit.

![Project Settings > Memory Settings, showing a selection of Player memory settings](../uploads/Main/memory-native-settings.png)

Project Settings > Memory Settings, showing a selection of Player memory settings

For more information on how the values affect each allocator, refer to [Native memory allocator examples](performance-native-memory-allocator-examples.md).

## Using command line arguments

You can use [command line arguments](CommandLineArguments.md) to set the size of each allocator. To find the name of the allocator parameters you want to change, check the list of allocator settings the Editor and players print when they start up.

For example, to change the block size of the main heap allocators, use the following:

`-memorysetup-main-allocator-block-size=<new_value>`

For a full list of command line arguments, refer to [Native memory allocators reference](performance-native-memory-allocator-reference.md).

## Measuring the performance of changes

To ensure your settings improve performance, [profile your application](profiler-collect-data) before and after making changes. For more information, refer to the [Profiler overview page](Profiler.md).

You can also use the [Profiler Analyzer](https://docs.unity3d.com/Packages/com.unity.performance.profile-analyzer@latest) package to measure changes. The Profiler Analyzer enables you to compare multiple frames, and two different Profiler captures to each other. Comparing two captures is useful for highlighting differences in allocator performance between two different runs with different settings.

You can also check the memory usage reports, which are available in the log when you close the Player or Editor. To find your log files, follow the instructions on the [log files page](log-files.md).

## Additional resources

* [Native memory allocator examples](performance-native-memory-allocator-examples.md)
* [Native memory allocators](performance-native-allocators.md)
* [Native memory allocators reference](performance-native-memory-allocator-reference.md)