# Navigating the Profiler window

To open the Profiler, go to **Window** > **Analysis** > **Profiler** or use the keyboard shortcut Ctrl+7 (Command+7 on macOS). Use the [Profiler module charts](profiler-modules-introduction.md) to inspect and analyze Profiler data.

## Selecting and inspecting a frame

To select and inspect a frame:

1. Click on the area of a Profiler module’s chart that you want to inspect. A white line appears, which highlights one frame of your application.
2. To navigate between frames, use the [transport controls](ProfilerWindow.md) in the **toolbar** of the Profiler window, or the arrow keys on your keyboard.

Unity then displays more details about the frame in the panel in the bottom half of the Profiler window. The type of detail in this window changes depending on which Profile module you select. For more information about the specific details that each module displays in this area, refer to the individual documentation for the [Profiler modules](ProfilerWindow.md).

Unity manages the vertical scale of the charts automatically, and the charts attempt to fill the vertical space of the window. To inspect a chart in more detail, drag the splitter between the charts and the details pane to increase the screen area of the charts.

## Customizing metric views

Hide or display metrics in a module’s chart, or reorder the metrics, to identify the causes of spikes, or make more prominent metrics visible. To customize the metrics displayed in a chart:

1. Click the colored square next to the metric’s label in its module to hide or display the metric.
2. Click and drag the handle icon (═) next to a metric to recorder it in stacked charts such as the [CPU Usage](ProfilerCPU.md) chart.

## Increasing frame count

By default, the Profiler records and keeps the last 2,000 frames of your application, and displays detailed information about every frame. To increase the default frame count:

1. Open the Preferences window (**Unity > Settings**)
2. Select the [Profiler preferences](profiler-preferences-reference.md) (**Analysis > Profiler**)
3. Enter a value for the **Frame Count**, between 600 and 4,000 frames.

**Note:** If you increase this setting to a large number, the Profiler’s overhead and memory usage might become more performance intensive. The **Automatic memory management** property is enabled by default to manage this behavior. When the system memory usage is at critical level and the Profiler uses more than 75% of Editor memory, then Unity discards frame data until system memory usage is back to a normal state (less than 25% of Editor memory).

## Additional resources

* [Profiler introduction](profiler-introduction.md)
* [Profiler modules introduction](profiler-modules-introduction.md)
* [Profiler window reference](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)