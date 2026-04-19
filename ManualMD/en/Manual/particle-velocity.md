# Particle velocity

Understand how the **Particle** System can change a particle’s velocity.

## Limiting velocity over lifetime

The [Limit Velocity over Lifetime](PartSysLimitVelOverLifeModule.md) module is very useful for simulating air resistance that slows the particles, especially when a decreasing curve is used to lower the speed limit over time. For example, an explosion or firework initially bursts at great speed but the particles emitted from it rapidly slow down as they pass through the air.

The **Drag** option offers a more physically accurate simulation of air resistance by offering options to apply varying amounts of resistance based on the size and speed of the particles.

## Inheriting velocity

Use the [Inherit Velocity](PartSysInheritVelocity.md) module on subemitters. Each particle in the parent system can spawn particles in the subemitter. This module reads the velocity from the parent particle and controls how the speed of the subemitter particles reacts to that velocity over time.

## Details

This effect is very useful for emitting particles from a moving object, such as dust clouds from a car, smoke from a rocket, steam from a steam train’s chimney, or any situation where the particles should initially be moving at a percentage of the speed of the object they appear to come from. This module only has an effect on the particles when **Simulation Space** is set to **World** in the [Main module](PartSysMainModule.md).

It is also possible to use curves to influence the effect over time. For example, you could apply a strong attraction to newly created particles, which reduces over time. This could be useful for steam train smoke, which would drift off slowly over time and stop following the train it was emitted from.

Unity calculates the emitter’s velocity in one of two ways:
\* Based on the velocity of an attached **Rigidbody** component
\* Based on how far the **Particle System**’s **Transform component** travelled during the current frame

To specify the method Unity uses, see the Main module’s [Emitter Velocity](PartSysMainModule.md) property: