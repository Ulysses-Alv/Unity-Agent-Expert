# Configuring global particle properties

The **Particle** System’s [Main](PartSysMainModule.md) module contains global properties that affect the whole system. Most of these properties control the initial state of newly created particles.

The system emits particles for a specific duration, and can be set to emit continuously using the **Looped** property. This allows you to set particles to be emitted intermittently or continuously; for example, an object may emit smoke in short puffs or in a steady stream.

The **Start** properties (**lifetime**, **speed**, **size**, **rotation** and **color**) specify the state of a particle on emission. You can specify a particle’s width, height and depth independently, using the **3D Start Size** property (see [Non-uniform particle scaling](#scaling), below).

All **Particle Systems** use the same gravity vector specified in the **Physics** settings. The Gravity Multiplier value can be used to scale the gravity, or switch it off if set to zero.

## Non-uniform particle scaling

The 3D Start Size property allows you to specify a particle’s width, height and depth independently. In the Particle System **Main** module, check the **3D Start Size** checkbox, and enter the values for the initial x (width), y (height) and z (depth) of the particle. Note that z (depth) only applies to 3D **Mesh** particles. You can also set randomised values for these properties, in a range between two constants or curves.

You can set the particle’s initial size in the Particle System **Main** module, and its size over the particle’s lifetime using the **Separate Axes** option in the **Size over Lifetime** module. You can also set the particle’s size in relation to its speed using the **Separate Axes** option in the **Size by Speed** module.

## Simulation Space

The **Simulation Space** property determines whether the particles move with the Particle System parent object, a custom object, or independently in the game world. For example, systems like clouds, hoses and flamethrowers need to be set independently of their parent **GameObject**, as they tend to leave trails that persist in the world space even if the object producing them moves around. On the other hand, if particles are used to create a spark between two electrodes, the particles should move along with the parent object. For more advanced control over how particles follow their Transform, see documentation on the [Inherit Velocity module](PartSysInheritVelocity.md).

When set to Custom, particles no longer move relative to their own **Transform component**. Instead, they all move relative to the movement of the specified Transform component. The Particle System uses the Custom Transform to calculate emitter velocity, which the Inherit Velocity module and **Rate over Distance** property of the Emission module use to control particle velocity and emission.