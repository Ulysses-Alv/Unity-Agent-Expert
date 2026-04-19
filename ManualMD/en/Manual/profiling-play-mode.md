# Collect performance data in Play mode

Profile in Play mode to quickly test changes and monitor the performance of your application without having to rebuild your application every time.

To profile in Play mode:

1. Open the Profiler (**Window > Analysis > Profiler**)
2. Select the Target Selection dropdown, next to Record
3. Select **Play Mode**

![Profiler windows Target Selection dropdown](../uploads/Main/profiler-target-player.png)

Profiler window’s Target Selection dropdown

Unity Profiler minimizes the overhead of profiling the Editor itself and represents any Editor-only activity when profiling Play mode as **EditorLoop** markers.

![Timeline view of Play mode frame](../uploads/Main/profiler-playmode-timeline.png)

Timeline view of Play mode frame

**Tip:** Whenever you profile Play mode, open it in a maximized view to run your application at a resolution closer to that of your target device. This directly affects performance issues such as those related to fill rate.

To make sure that windows other than Play mode or the Profiler don’t use up time on the render thread and GPU, which affects performance data, reduce the amount of open Unity Editor windows.

## Additional resources

* [Collect performance data introduction](profiling-collect-data-introduction.md)
* [Collect performance data on a target platform](profiling-target-device.md)
* [Collect performance data about the Unity Editor](profiling-edit-mode.md)