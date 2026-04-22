# Adaptive Performance Profiler integration

Visualize Adaptive Performance metrics in the Unity Profiler.

In the [Unity Profiler](../profiler-introduction.md), you can use the Adaptive Performance [Profiler module](../profiler-modules-introduction.md) to display status information about scalers, bottlenecks, and thermal warning throttling.

This feature is available in the Unity Editor and for **development builds** on devices.

![The Adaptive Performance Profiler module running in the Unity Profiler and showing a CPU bottleneck, thermal throttling, and all enabled scalers.](../../uploads/Main/ap/profiler.png)

The Adaptive Performance Profiler module running in the Unity Profiler and showing a CPU bottleneck, thermal throttling, and all enabled scalers.

## Scaler status in the Profiler

The scalers section displays all available [scalers](scalers.md). Each bar represents a scaler and displays its current quality level and the maximum level that it can reach.

### Scaler status colors

The color of the scaler bar indicates its current state:

| **Color** | **State** |
| --- | --- |
| **Gray** | Scaler is disabled. |
| **Bright green** | Scaler is enabled and applied. The Adaptive Performance system is actively adjusting the scaler at this specific frame. |
| **Faded green** | Scaler is disabled and applied. The Adaptive Performance system is not actively adjusting the scaler at this specific frame, but it has been adjusted in the past and is currently at a non-default level. |
| **Blue** | Scaler is enabled but not applied. |

## Additional resources

* [Collect performance metrics](performance-metrics.md)
* [Apply performance optimizations](performance-optimization-strategies.md)
* [Unity Profiler](../profiler-introduction.md)
* [Profiler modules introduction](../profiler-modules-introduction.md)