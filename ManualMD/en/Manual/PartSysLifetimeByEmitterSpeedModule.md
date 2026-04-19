# Lifetime by Emitter Speed module reference

This module controls the initial lifetime of each **particle** based on the speed of the emitter when the particle spawns.

## Properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| **Property** | **Function** |
| --- | --- |
| **Multiplier** | The multiplier to apply to the particle’s initial lifetime. The module uses this value differently depending on the curve mode you set. The curve modes are: • **Constant**: Uses the constant multiplier value you set for this property. Using this curve mode ignores the **Speed Range** property. • **Curve**: Takes the emitter’s speed, maps it to a value between 0 and 1 based on the **Speed Range**, then uses the normalized value to sample the curve. • **Random Between Two Constants**: Sets a random multiplier for each particle between the two values you set for this property. Using this curve mode ignores the **Speed Range** property. • **Random Between Two Curves**: Takes the emitter’s speed, maps it to a value between 0 and 1 based on the **Speed Range**, then uses the normalized value to sample each curve. For each particle, the module sets the multiplier to a random value between the two samples. |
| **Speed Range** | The minimum and maximum emitter speeds the Particle System maps to a value along the **Multiplier** curve. If the emitter’s speed is equal to the first value, the multiplier is the value at the start of the curve. If the emitter’s speed is equal to the second value, the multiplier is the value at the end of the curve. This property is only relevant if the curve mode for **Multiplier** is set to **Curve** or **Random Between Two Curves**. |

## Additional resources

* [Particle emissions and emitters](particle-emissions-emitters.md)