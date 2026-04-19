# Connect the profiler to UWP

You can use the Unity profiler to get performance information about your application. For more information, refer to [Profiling your application](profiler-profiling-applications.md).

Due to restrictions with UWP, you won’t be able to connect the profiler if the Unity Editor is running on the same machine as UWP. Therefore, make sure to run the Unity Editor and UWP on separate machines. For example, if you’re running the Unity Editor and UWP on the same PC, you won’t be able to connect the profiler. The only exception to this rule is the **Autoconnect Profiler** build option, which makes the application connect to the Editor instead.

You must also ensure that the machine where the Unity Editor is running and the machine where the Universal Windows App is running are on the same subnet.

**Note:** The profiler doesn’t work on Master configuration.

## Connect the Unity profiler

To connect the Unity profiler to a running Universal Windows application, perform the following steps:

1. Go to **Edit** > **Project Settings** > **Player**.
2. Select the **Publishing Settings** > **Capabilities** section.
3. Enable **Private Networks Capability**.
4. Enable **Internet (Client & Server) Capability**.
5. If you’ve already enabled the **Autoconnect Profiler** in [Build Profiles](windowsstore-buildsettings.md), the profiler should connect automatically to the Universal Windows App. If not, you have to explicitly select it in Unity in **Window** > **Analysis** > **Profiler** > **Active Profiler**.
6. Build the Universal Windows App Visual Studio solution from Unity.
7. Select **Build and Run**.