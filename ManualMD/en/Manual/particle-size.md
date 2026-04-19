# Particle size

Understand how the **Particle** System can change a particle’s size based on its speed or lifetime.

## Changing particle size based on the particle’s speed

The [Size By Speed Module](PartSysSizeBySpeedModule.md) can create particles that change size based on their speed in distance units per second.

Some situations will require particles which vary in size depending on their speed. For example, you would expect small pieces of debris to be accelerated more by an explosion than larger pieces. You can achieve effects like this using **Size By Speed** with a simple ramp curve that proportionally increases the speed as the size of the particle decreases. Note that this should not be used with the **Limit Velocity Over Lifetime** module, unless you want particles to change their size as they slow down.

**Speed Range** specifies the range of values that the X (width), Y (height) and Z (depth) shapes apply to. The Speed Range is only applied when the size is in one of the curve modes. Fast particles will scale using the values at the right end of the curve, while slower particles will use values from the left side of the curve. For example, if you specify a Speed Range between 10 and 100:

* Speeds below 10 will set the Particle size corresponding with the left-most edge of the curve.
* Speeds above 100 will set the Particle size corresponding with the right-most edge of the curve.
* Speeds between 10 and 100 will set the Particle size determined by the point along the curve corresponding to the Speed. In this example, a Speed of 55 would set the size according to the midpoint of the curve.

### Non-uniform particle scaling by speed

You can specify how a particle’s width, height and depth size changes by speed independently. In the **Size by Speed** module, check the **Separate Axes** checkbox, then choose how the X (width), Y (height) and Z (depth) of the particle is affected by the speed of the particle. Remember that Z will only be used for **Mesh** particles.

## Changing particle size over the particle’s lifetime

The [Size Over Lifetime](PartSysRotOverLifeModule.md) module can change a particle’s size based on how long it has existed.

Some particles will typically change in size as they move away from the point of emission, such as those that represent gases, flames or smoke. For example, smoke will tend to disperse and occupy a larger volume over time. You can achieve this by setting the curve for the smoke particle to an upward ramp, increasing with the particle’s age. You can also further enhance this effect using the **Color Over Lifetime** module to fade the smoke as it spreads.

For fireballs created by burning fuel, the flame particles will tend to expand after emission but then fade and shrink as the fuel is used up and the flame dissipates. In this case, the curve would have a rising “hump” that rises and then falls back down to a smaller size.

The values specified in the curves are multiplied by the [Start Size](PartSysMainModule.md) to get the final particle size.

### Non-uniform particle scaling over lifetime

You can specify how a particle’s width, height and depth changes over lifetime independently. In the **Size over Lifetime** module, check the **Separate Axes** checkbox, then change the X (width), Y (height) and Z (depth). Remember that Z will only be used for Mesh particles.