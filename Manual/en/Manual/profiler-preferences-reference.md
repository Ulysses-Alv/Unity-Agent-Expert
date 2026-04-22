# Profiler Preferences reference

The Preferences window contains additional Profiler window settings. To open the Preferences window go to menu: **Unity > Settings > Analysis > Profiler**.

## Profiler settings

Adjust the settings to customize how the Profiler collects and displays data.

| **Setting** | **Description** |
| --- | --- |
| **Profiler Capture Storage Path** | Set the default folder where Unity saves [Profiler captures](ProfilerWindow.md#captures-list). |
| **Frame count** | Set the maximum number of frames for the Profiler to capture. You can set this number between 600 and 4,000. Higher number of loaded frames increases memory usage of Unity Editor. |
| **Automatic memory management** | Automatically manage Profiler memory usage. When the system memory usage is at a critical level and the Profiler uses more than 75% of the Editor memory, Unity discards data until the system memory usage is back to a normal state, or the Profiler uses less than 25% of Editor memory. |
| **Show stats for ‘current frame’** | Set whether to display annotations on the frame line when you select the Current Frame button (⏭) on the Profiler controls.   This setting is disabled by default, because the annotations might make it difficult to view data in real-time. To display the annotations, enable this setting. |
| **Default recording state** | Select which recording state the Profiler opens in:   * **Enabled**: Keeps the Record button enabled between profiling sessions. * **Disabled**: Disables the Record button, regardless of whether you turn it on or off during your profiling session. * **Remember**: Remembers whether you enabled or disabled the Record button during your session and keeps it at its last state the next time you open the Profiler window. |
| **Default editor target mode on start** | Select what mode the Target Selection dropdown targets by default:   * **Play mode** (default): Targets [Play mode profiling](profiling-play-mode.md) by default. * **Edit mode**: Targets [Edit mode profiling](profiling-edit-mode.md) by default. |
| **Custom connection ID** | Change the name of your application as it appears in the Target Selection dropdown. |
| **Target **Frames Per Second** (Highlights Module)** | Set the target frames per second that the [Highlights Profiler module](ProfilerHighlights.md) uses. The Highlights Profiler module flags any frames which exceed this target frame rate. |

## Additional resources

* [Profiler window reference](ProfilerWindow.md)
* [Profiler modules introduction](profiler-modules-introduction.md)
* [Connecting the Profiler to a data source](profiler-profiling-applications.md)