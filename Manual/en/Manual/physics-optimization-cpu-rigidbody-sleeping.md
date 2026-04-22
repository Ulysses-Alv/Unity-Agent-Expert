# Use Rigidbody sleeping to improve physics performance

Reduce CPU load and improve physics performance by enabling **Rigidbody** sleeping for stationary objects.

Rigidbody sleeping can drastically reduce CPU load, especially in **scenes** with many physical objects that are often stationary or frequently come to rest. When a **Rigidbody** component moves at a slower speed than the **[Sleep Threshold](class-PhysicsManager.md)**, the physics system sets the **Rigidbody** component to a sleeping state. When a **Rigidbody** component is asleep, the physics system doesn’t include it in physics calculations. When a sleeping **Rigidbody** component receives a **collision** or force, the physics system wakes up the Rigidbody component and includes it in physics calculations.

In script, control Rigidbody sleeping with [`Rigidbody.Sleep`](../ScriptReference/Rigidbody.Sleep.md) and [`Rigidbody.WakeUp`](../ScriptReference/Rigidbody.WakeUp.md)

Rigidbody sleeping is highly effective in environments with many interactive props, destructible elements that settle, or physics-based puzzles that stabilize. While scenes with constant high-velocity motion benefit less, enabling sleeping is generally a good default behavior.

To enable Rigidbody sleeping, the recommended best practices are:

* Ensure that the **Rigidbody** component’s sleep threshold is set appropriately. Set the **Sleep Threshold** in the [Physics Project Settings](class-PhysicsManager.md) and in script with [`Rigidbody.sleepThreshold`](../ScriptReference/Rigidbody-sleepThreshold.md).
* Avoid calling `Rigidbody.WakeUp` unnecessarily on objects unless they need to be active in the simulation. Continuously waking objects negates the benefit of sleeping.
* Check if a **Rigidbody** component is sleeping with [`Rigidbody.IsSleeping`](../ScriptReference/Rigidbody.IsSleeping.md).
* Use the Physics Debugger (**Window > Analysis > Physics Debugger**) to visually inspect the sleep state of **Rigidbody** components in your scene. The Physics Debugger can help identify objects that are unexpectedly active and not sleeping, potentially due to persistent small contacts or incorrect sleep settings.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [`Rigidbody.Sleep`](../ScriptReference/Rigidbody.Sleep.md)