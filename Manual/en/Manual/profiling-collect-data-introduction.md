# Collect performance data introduction

Collect performance data related to your application with the Profiler.

To collect data, you must connect the Profiler to a data source. You also must enable any [Profiler modules](profiler-modules-introduction.md) that you want to collect data for.

You can collect data with the [Profiler](profiler-introduction.md) from the following sources:

* Profile your application in a Player on a target platform.
* Profile your application in Play mode.
* Profile the Unity Editor.

The Profiler collects data only for the Profiler modules you have enabled. For more information, refer to [Activating Profiler modules](profiler-modules-activate.md).

## Collect performance data on a target platform

The best way to get accurate timings about your application is to [profile it on the end platform](profiling-target-device.md) you intend to publish it on. This gives you accurate timings about what impacts the performance of your application.

## Collect performance data in Play mode

It can be time-consuming to build your application every time you want to improve elements of its performance. To quickly assess the performance of your application you can [profile it directly in Play mode](profiling-play-mode.md) in the Editor. The Profiler’s default target is Play mode.

Play mode runs in the same application and main thread as the Editor which means that the Editor’s systems such as the UI, **Inspectors**, **Scene** view rendering, and asset management affect the performance and memory measurements you get when profiling in Play mode.

Profiling in Play mode doesn’t give you an accurate reflection of what the performance of your application looks like on a real device. It’s useful to test out changes without having to rebuild a player, and to identify areas to investigate further.

To effectively profile in Play mode, build your application often and deploy it to a range of target devices, and test and profile your application on these devices. If you identify performance issues with your application on these devices, narrow down the area that needs the most attention.

You can then profile your application in Play mode and quickly iterate over any changes you make to your application. Once you’re satisfied with your changes, build and deploy your application to the target devices again to validate your changes.

## Collect performance data about the Unity Editor

The Editor might affect the performance of your application, because it uses the same resources as your application when it’s running in Play mode. You can [profile the Editor](profiling-edit-mode.md) separately to find out what resources it uses. This is useful if you have designed your application to work only in Play mode, such as for movie making.

You can also profile in Play mode or profile the Editor to identify issues unrelated to the performance of your application. For example, to investigate if long load times or an unresponsive Editor slows down iteration time, or if your application’s performance is poor in Play mode.

## Additional resources

* [Connecting the Profiler to your application](profiler-profiling-applications.md)
* [Profiler modules introduction](profiler-modules-introduction.md)
* [Profiler window reference](ProfilerWindow.md)
* [Adding profiling information to your code](profiler-adding-information-code.md)