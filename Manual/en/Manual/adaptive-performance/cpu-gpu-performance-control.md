# CPU and GPU performance control

Use performance control modes in Adaptive Performance to control how your application uses CPU and GPU resources to balance performance, power use, and device temperature on mobile devices.

## CPU and GPU power management

The CPU and GPU are the main sources of power consumption on mobile devices, especially during gameplay. Running these components at maximum clock speeds can be inefficient, consuming excess power and generating heat without a proportional gain in performance. This can cause overheating, prompting the operating system to throttle CPU and GPU frequency to cool down the device. Typically, the operating system controls CPU and GPU clock speeds.

Adaptive Performance provides performance control modes to manage CPU and GPU performance levels to optimize power consumption and thermal behavior.

## Performance control modes

Adaptive Performance has three modes for controlling CPU and GPU performance: automatic, manual, and system. Typically, automatic mode provides the best results over a wide set of conditions.

To check which mode your application currently runs in, use [`IDevicePerformanceControl.PerformanceControlMode`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.PerformanceControlMode.md).

**Note**: Automatic and manual performance control modes are only available on Android providers and in the [Device Simulator](device-simulator-reference.md). All other providers use system performance control.

### System performance control

In system performance control, the operating system manages CPU and GPU levels. This mode is the default on non-Android providers. It’s also the recommended mode if you want to use [temporary performance boosts](temporary-performance-boosts.md).

### Automatic performance control

Automatic performance control uses real-time feedback to adjust CPU and GPU levels according to your application’s needs and the device’s thermal state. It considers whether:

* The application reached the target frame rate in the previous frames.
* Device temperatures are rising.
* The device is close to thermal throttling.
* The device is GPU or CPU bound.

For best results, use the automatic performance control mode and set the [`targetFrameRate`](../../ScriptReference/Application-targetFrameRate.md) for your application. Adaptive Performance then adjusts CPU and GPU levels to maintain this frame rate while minimizing power consumption and heat generation. You can change the target frame rate based on the content that your app is playing to further optimize performance. For example, you might lower the frame rate on menu or pause screens and restore it during game play.

The following example demonstrates how to reduce thermal pressure and power consumption automatically, where you need only set the `targetFrameRate`:

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;

public class AutomaticPerformance : MonoBehaviour
{
    // Declare 'ap' as a private class-level field to hold the IAdaptivePerformance instance.
    private IAdaptivePerformance ap;

    void Start()
    {
        // Get the Adaptive Performance instance.
        ap = Holder.Instance;
    }

    public void EnterMenu()
    {
        // Check if Adaptive Performance is active.
        if (ap == null || !ap.Active)
        {
            return;
        }

        // Set the target frame rate for the application.
        Application.targetFrameRate = 30;
        // Get the device performance control interface.
        var ctrl = ap.DevicePerformanceControl;
        // Enable automatic CPU and GPU level regulation by Adaptive Performance.
        ctrl.AutomaticPerformanceControl = true;
    }
}
```

### Manual performance control

Manual performance control allows direct, fine-grained adjustment of CPU and GPU levels. However, it’s not recommended for most applications. This mode requires careful balancing with the device’s thermal state and battery consumption, as incorrect settings can lead to overheating, thermal throttling, and excessive battery drain. Use it only for specific scenarios such as benchmarking or highly optimized custom loops.

Note: For optimal results in most cases, use automatic performance control and set only a `targetFrameRate`.

To enter manual performance control, set [`IDevicePerformanceControl.AutomaticPerformanceControl`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.AutomaticPerformanceControl.md) to `false`.

In manual mode, you can change these properties to optimize CPU and GPU performance:

* [`IDevicePerformanceControl.CpuLevel`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.CpuLevel.md)
* [`IDevicePerformanceControl.GpuLevel`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.GpuLevel.md)

The following example demonstrates how to enter manual performance control and set CPU and GPU levels to their maximum values:

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;

public class ManualPerformance : MonoBehaviour
{
    // Declare 'ap' as a private class-level field to hold the IAdaptivePerformance instance.
    private IAdaptivePerformance ap;

    void Start()
    {
        // Get the Adaptive Performance instance.
        ap = Holder.Instance;
    }

    public void EnableManualMode()
    {
        // Check if Adaptive Performance is active.
        if (ap == null || !ap.Active)
        {
            return;
        }

        // Get the device performance control interface.
        var ctrl = ap.DevicePerformanceControl;
        // Disable automatic performance control to enable manual mode.
        ctrl.AutomaticPerformanceControl = false;
        // Set CPU level to maximum performance.
        ctrl.cpuLevel = ctrl.MaxCpuPerformanceLevel;
        // Set GPU level to maximum performance.
        ctrl.gpuLevel = ctrl.MaxGpuPerformanceLevel;
    }
}
```

**Notes:**

* Changing GPU and GPU levels only has an effect when the device isn’t in a thermal throttling state, meaning the [`WarningLevel`](../../ScriptReference/AdaptivePerformance.WarningLevel.md) isn’t [`WarningLevel.Throttling`](../../ScriptReference/AdaptivePerformance.WarningLevel.Throttling.md).
* In some situations, the device might take control over the CPU and GPU levels. The system logs this event and changes the value of [`IDevicePerformanceControl.PerformanceControlMode`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.PerformanceControlMode.md)
  to [`PerformanceControlMode.System`](../../ScriptReference/AdaptivePerformance.PerformanceControlMode.System.md).

## Additional resources

* [Application.targetFrameRate](../../ScriptReference/Application-targetFrameRate.md)
* [Provider settings reference](provider-settings-reference.md)
* [Track timing and thermal data](track-timing-thermal-data.md)