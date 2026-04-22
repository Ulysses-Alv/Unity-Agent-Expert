# Light Probe validity in Adaptive Probe Volumes in URP

Light Probes inside geometry are called invalid probes. The Universal **Render Pipeline** (URP) marks a **Light Probe** as invalid when the probe fires sampling rays to capture surrounding light data, but the rays hit the unlit backfaces inside geometry.

URP uses the following techniques to minimise incorrect lighting data from Light Probes:

* [Virtual Offset](probevolumes-troubleshoot-artefacts.md#virtualoffset) tries to make invalid Light Probes valid, by moving their capture points so they’re outside any [colliders](https://docs.unity3d.com/Documentation/Manual/CollidersOverview.html).
* [Dilation](probevolumes-troubleshoot-artefacts.md#dilation) detects Light Probes that remain invalid after Virtual Offset, and gives them data from valid Light Probes nearby.
* [Rendering Layers](probevolumes-troubleshoot-light-leaks.md#layers) prevent objects from sampling probes that are on another **Layer Mask**, reducing light leaking in certain scenarios.

You can check which Light Probes are invalid using the [Rendering Debugger](features/rendering-debugger.md)

![In the Scene on the left, Virtual Offset isnt active and dark bands are visible. In the Scene on the right, Virtual Offset is active.](../../uploads/urp/probe-volumes/probevolumes-virtualoffsetvsnot.png)

![In the Scene on the left, Dilation isnt active and some areas are too dark. In the Scene on the right, Dilation is active.](../../uploads/urp/probe-volumes/probevolumes-dilationvsnot.png)