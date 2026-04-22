# Profiler introduction

Analyze the performance of your application with the Profiler.

The Profiler records multiple areas of your application’s performance, and displays that information to you. You can use this information to decide what you might need to optimize in your application, and to test the performance of changes you make.

To open the Profiler window go to **Window** > **Analysis** > **Profiler**.

![Profiler window with a frame in the CPU Usage Profiler module selected. The Timeline view is selected in the details pane.](../uploads/Main/profiler-cpu-module.png)

Profiler window with a frame in the CPU Usage Profiler module selected. The Timeline view is selected in the details pane.

You can inspect script code and view how your application uses certain assets and resources that might be slowing it down. You can also compare how your application performs on different devices. The Profiler has several different [Profiler modules](profiler-modules-introduction.md) which display performance data in areas such as rendering, memory, and audio.

The Profiler is an instrumentation-based profiler, which means that the Profiler uses markers in your application’s code to record detailed timing information about how long the code in each marker takes to execute. The Unity API has profiler markers built in so you can explore the performance of your code, locate performance issues, and identify areas to optimize.

You can also use custom [Profiler markers](profiler-markers.md) to monitor specific data, or use [Deep Profiling](profiler-deep-profiling.md) to further customize the data you gather.

## Additional resources

* [Collect performance data](profiler-collect-data)
* [Profiler modules](profiler-visualizing-data.md)
* [Using the Profiler window](ProfilerWindow.md)