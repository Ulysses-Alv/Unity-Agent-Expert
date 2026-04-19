# Platform and mesh considerations

Things to consider when building UI for different platforms and **mesh** tessellation.

## Consider device capabilities

Some Android devices and the Web platform can’t partially patch index buffers. If your audience uses such devices or if you target the Web platform, UI Toolkit rendering is still functional but performance may be degraded. To avoid performance degradation, don’t animate too many elements at the same time and [profile on device](profiler-profiling-applications.md).

## Avoid mesh tessellation

It’s computationally expensive to build mesh tessellation for **visual elements**. Any time the size (width/height) of the element changes, its geometry re-builds, which can be an issue with animations or frequent resizing.

Generally speaking, [transforms](UIE-Transform.md) and texture aren’t good choices in terms of flexibility and re-use. However, when you animate, to get better performance, do the following:

* Use [transforms](UIE-Transform.md) instead of width or other layout properties
* Use textures or 2D **sprites** instead of rounded corners and borders

## Additional resources

* [`Transform`](../ScriptReference/Transform.md)
* [Profiler](profiler-profiling-applications.md)