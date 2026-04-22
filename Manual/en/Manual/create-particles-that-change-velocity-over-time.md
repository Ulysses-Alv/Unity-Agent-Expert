# Create particles that change velocity over time

To create **particles** that drift in a particular direction, use the [Velocity over Lifetime](PartSysVelOverLifeModule.md) module’s Linear X, Y and Z curves.

To create effects with particles that spin around a center position, use the use the [Velocity over Lifetime](PartSysVelOverLifeModule.md) module’s **Orbital** velocity values. Additionally, you can propel particles towards or away from a center position using the **Radial** velocity values. You can define a custom center of rotation for each particle by using the **Offset** value.

You can also use the [Velocity over Lifetime](PartSysVelOverLifeModule.md) module to adjust the speed of the particles in the **Particle System** without affecting their direction, by leaving all the above values at zero and only modifying the **Speed Modifier** value.