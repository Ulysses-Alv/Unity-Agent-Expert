# Android thread configuration

Unity configures [thread affinity](#thread-affinity) and [thread priority](#thread-priority) based on the CPU topology of the device. Unity’s default thread configuration works well for most projects, but in some situations, you might want to change the thread configuration (for example, if you want to optimize for specific devices, or if your application aims for low power consumption and doesn’t require high frame rates).

**Important**: Use Unity’s default settings when possible. If you change the thread configuration to optimize for specific devices, any changes you make can have a negative performance impact on other devices or even the same devices in future operating system versions.

## Thread affinity

The affinity of a thread controls which CPU cores the thread should run on. For Android devices, cores usually fit into two categories:

* Big: Powerful cores that are battery power-intensive.
* Little: Slower cores that are less battery power-intensive than big cores.

**Important**: For typical applications that target a variety of Android devices, don’t bind threads to specific cores.

### Identifying CPU cores

Unity usually receives the capacity and big/little assignment for every CPU core from the device’s operating system. For older operating system versions, this information isn’t always available. In this case, Unity calculates the capacity of the CPU cores and uses that information to assign each core as either big or little. To assign a core, Unity compares the CPU capacity of each core to a threshold. By default, a core is big if it has at least double the CPU capacity of the slowest core, and is little otherwise.

For more control over which cores Unity assigns as big or little, you can provide a custom threshold. To provide a custom threshold, use the `-platform-android-cpucapacity-threshold [value]` command-line argument with a value between 0 and 1024 where 0 represents the lowest capacity core and 1024 represents the highest capacity core. For example, a value of `870`, which is approximately 85% of 1024, means that a core is big if its CPU capacity is in the top 15% of those on the chip. For information on how to add start-up command-line arguments to Unity on Android devices, see [Specify Unity startup arguments from a custom UnityPlayerActivity file](android-custom-activity-command-line.md).

**Note**: Unity always categorizes a core as either big or little. For some complex CPU topologies (for example, those that have medium cores), Unity still categorizes its cores as either big or little.

### Thread affinity aliases

Unity provides the following aliases that you can use as the thread affinity value:

* `any`: Allows the thread to run on any core.
* `little`: Allows the thread to run on any little core.
* `big`: Allows the thread to run on any big core.

You can also use hexadecimal and binary values to specify thread affinity. In binary, the index of the bit references a specific CPU core. Typically, little cores start at index 0 and big cores come directly after them.

For example, if the CPU contains four little cores and four big cores:

* The binary value `0b11110000` and hexadecimal value `0xf0` allow the thread to run on big cores.
* The binary value `0b1111` and hexadecimal value `0xf` allow the thread to run on little cores.

## Thread priority

The priority of a thread controls how the device’s operating system allocates CPU time to the thread. The operating system allocates more CPU time to threads with a higher priority relative to threads with a lower priority.

The thread priority values you can use in Unity are in the range of –20 to 19, where –20 indicates the highest priority and 19 indicates the lowest priority.

## Configurable Unity threads

This section contains information about the Unity threads you can configure, and which command-line arguments you use to configure them.

For information on how to add startup command-line arguments to Unity on Android devices, see [Specify Unity startup arguments from a custom UnityPlayerActivity file](android-custom-activity-command-line.md).

**Important**: On some devices and Android OS versions, the Android OS might ignore command-line parameters and throw an error. If this occurs, you application still runs normally without issue, but Android doesn’t apply the thread affinity or thread priority settings that you specified.

### Unity main thread

Unity’s main thread executes all **scripts** and usually has a high CPU load. The command-line arguments that affect this thread are:

* `-platform-android-unitymain-priority [value]` where `[value]` is the [priority](#thread-priority) of the thread.
* `-platform-android-unitymain-affinity [value]` where `[value]` is the [affinity](#thread-affinity) of the thread.

### Unity JobWorker threads

Unity’s worker threads execute jobs from both the core engine and those your application dispatches via the [C# Job System](job-system.md). The command-line arguments that affect these threads are:

* `-platform-android-jobworker-priority [value]` where `[value]` is the [priority](#thread-priority) of the threads.
* `-platform-android-jobworker-affinity [value] [value1 value2 value3]` where `[value]` is the [affinity](#thread-affinity) of the thread and `[value1 value2 value3]` are optional values you can use to specify the affinity of worker threads per thread. If you only set a single value, Unity uses the same value for all worker threads. You can use ’-job-worker-count’ to specify the number of worker threads.

### Unity render thread

Unity’s render thread interacts with the Graphics API if your project uses [multithreaded rendering](../ScriptReference/Rendering.RenderingThreadingMode.MultiThreaded.md).

* `-platform-android-gfxdeviceworker-priority [value]` where `[value]` is the [priority](#thread-priority) of the thread.
* `-platform-android-gfxdeviceworker-affinity [value]` where `[value]` is the [affinity](#thread-affinity) of the thread.

**Note**: If you use GraphicsJobs, JobWorker threads also interact with the Graphics API.