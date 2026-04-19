# Introduction to Adaptive Performance

Analyze and improve the performance of your application on a mobile device with Adaptive Performance.

Adaptive Performance tracks your device’s thermal and power state and uses that data to dynamically adjust your application’s workload. This proactive adjustment helps prevent thermal throttling and maintain a stable frame rate.

The Adaptive Performance system uses [providers](providers.md) to continuously monitor key performance metrics, including the following:

* Device thermal state
* CPU and GPU utilization
* Performance bottlenecks (such as CPU, GPU, or target frame rate limitations)

![Adaptive Performance capturing performance metrics in the Device Simulator.](../../uploads/Main/ap/ap-cover.jpg)

Adaptive Performance capturing performance metrics in the Device Simulator.

Using these insights, Adaptive Performance automatically prevents overheating and excessive power consumption through the following:

* Adjustments to application quality settings via [scalers](scalers.md)
* Dynamic control of [CPU and GPU performance levels](cpu-gpu-performance-control.md)
* [Temporary performance boosts](temporary-performance-boosts.md)

## Access sample projects

Sample projects provide practical demonstrations of Adaptive Performance features. Refer to [Adaptive Performance sample projects](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@latest?subfolder=/manual/samples-guide.html) for detailed examples and code.

## Additional resources

* [Get started with Adaptive Performance](get-started.md)
* [Collect performance metrics](performance-metrics.md)
* [Apply performance optimizations](performance-optimization-strategies.md)