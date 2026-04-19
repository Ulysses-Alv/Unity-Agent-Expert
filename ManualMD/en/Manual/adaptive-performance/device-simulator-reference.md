# Adaptive Performance Device Simulator plug-in reference

Explore the Adaptive Performance settings and controls that simulate thermal and bottleneck conditions in the Device Simulator.

Find these settings in the [Simulator view](../device-simulator-view.md), under **Control Panel** > **Adaptive Performance**.

![Adaptive Performance settings in the Device Simulator.](../../uploads/Main/ap/simulator-extension-settings.png)

Adaptive Performance settings in the Device Simulator.

## Adaptive Performance plug-in settings

The Adaptive Performance **plug-in** contains the following settings.

### Thermal

The **Thermal** section contains the following settings:

| **Property** | **Description** |
| --- | --- |
| **Warning Level** | Simulates a thermal warning level.  * **No Warning**: Simulates that no thermal issues are detected. * **Throttling Imminent**: Simulates that the device is approaching thermal throttling. * **Throttling**: Simulates that the device is actively throttling performance. |
| **Level** | Simulates the skin temperature level of a device. Options range between nominal temperature (0) and critical throttling temperature (maximum level of 1.0). |
| **Trend** | Simulates a temperature trend. Values less than 0 simulate a temperature drop. Values higher than 0 simulate a temperature increase. |

### Performance

The **Performance** section contains the following settings:

| **Property** | **Description** |
| --- | --- |
| **Bottleneck** | Simulates the device bottleneck.  * **Unknown**: The system can’t determine a primary bottleneck. * **CPU**: Device is CPU bound. * **GPU**: Device is GPU bound. * **Target Frame Rate**: The application reaches its pre-defined frame rate limit and is intentionally capped by VSync (syncing to the monitor’s refresh rate) or the value set in [`Application.targetFrameRate`](../../ScriptReference/Application-targetFrameRate.md). |
| **Mode** | Simulates the device performance mode.  * **Unknown**: Mode not set or unknown. * **Standard**: Default mode that balances battery and CPU. * **Optimize**: Optimizes towards CPU, GPU, and battery consumption. * **CPU**: Device prioritizes CPU performance. * **GPU**: Device prioritizes GPU performance. * **Battery**: Device prioritizes battery consumption. |
| **Auto Control Mode** | Sets the CPU and GPU level automatically. When the device is throttling, the system’s own thermal management takes control and ignores **Auto Control Mode**. |
| **Target **FPS**** | Simulates the application target frame rate in **frames per second**. |
| **CPU Level** | Simulates the frequency cap of the CPU in discrete levels. |
| **GPU Level** | Simulates the frequency cap of the GPU in discrete levels. |
| **CPU Boost** | Boosts CPU performance for 10 s. Enabling **CPU Boost** disables CPU levels until the boost is over. If the system is throttling, the CPU boost has no effect. For more information, refer to [Temporary performance boosts](temporary-performance-boosts.md). |
| **GPU Boost** | Boosts GPU performance for 10 s. Enabling GPU boost disables GPU levels until the boost is over. If the system is throttling, the GPU boost has no effect. For more information, refer to [Temporary performance boosts](temporary-performance-boosts.md). |

### Indexer

The Indexer monitors the device’s thermal and performance state to decide when to adjust quality levels using [scalers](scalers.md).

The **Indexer** section contains settings that override the Indexer’s automatic behavior:

| **Property** | **Description** |
| --- | --- |
| **Thermal Action** | The action that the Indexer takes to adjust the visual quality level to regulate the device temperature. For more information about the individual scalers, refer to [Modifying asset quality with scalers](scalers-introduction.md).  * **Stale**: Doesn’t override the Indexer’s actions. * **Increase**: Increases the visual quality level. * **Decrease**: Decreases the visual quality level gradually. * **Fast Decrease**: Decreases the visual quality level quickly. |
| **Performance Action** | The action that the Indexer takes to adjust the visual quality level to regulate the application’s performance.  * **Stale**: Doesn’t override the Indexer’s actions. * **Increase**: Increases the visual quality level. * **Decrease**: Decreases the visual quality level gradually. * **Fast Decrease**: Decreases the visual quality level quickly. |

#### Scalers

Use this section to configure scalers for the Device Simulator. The settings are identical to those in the [Adaptive Performance provider settings](provider-settings-reference.md#scaler-profiles).

### Device Settings

The **Device Settings** section contains the following settings:

#### Cluster Info

The following settings simulate a device’s CPU clusters, which group different types of cores (big, medium, and little) to balance performance and power. For information about CPU clusters, refer to [Track timing and thermal data](track-timing-thermal-data.md#clusters).

| **Property** | **Description** |
| --- | --- |
| **Big Cores** | Simulates the number of big CPU cores available on the device. |
| **Medium Cores** | Simulates the number of medium CPU cores available on the device. |
| **Little Cores** | Simulates the number of little CPU cores available on the device. |

### Developer Options

Use this section to configure developer settings for the Device Simulator. These settings are identical to the **Development Settings** in the [Adaptive Performance provider settings](provider-settings-reference.md#development-settings).

## Additional resources

* [Apply performance optimizations](performance-optimization-strategies.md)
* [Collect performance metrics](performance-metrics.md)
* [Provider settings reference](provider-settings-reference.md)
* [Adaptive Performance Profiler integration](profiler-integration.md)
* [Profiler modules introduction](../profiler-modules-introduction.md)