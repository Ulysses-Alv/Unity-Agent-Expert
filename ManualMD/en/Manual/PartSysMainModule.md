# Main module reference

The **Particle** System module contains global properties that affect the whole system. Most of these properties control the initial state of newly created particles.

The name of the module appears in the **inspector** as the name of the **GameObject** that the [Particle System component](class-ParticleSystem.md) is attached to.

## Properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| **Property** | **Function** |
| --- | --- |
| **Duration** | The length of time the system runs. |
| **Looping** | If enabled, the system starts again at the end of its duration time and continues to repeat the cycle. |
| **Prewarm** | If enabled, the system is initialized as though it had already completed a full cycle (only works if **Looping** is also enabled). |
| **Start Delay** | Delay in seconds before the system starts emitting once enabled. |
| **Start Lifetime** | The initial lifetime for particles. |
| **Start Speed** | The initial speed of each particle in the appropriate direction. |
| **3D Start Size** | Enable this if you want to control the size of each axis separately. |
| **Start Size** | The initial size of each particle. |
| **3D Start Rotation** | Enable this if you want to control the rotation of each axis separately. |
| **Start Rotation** | The initial rotation angle of each particle. |
| **Flip Rotation** | Causes some particles to spin in the opposite direction. |
| **Start Color** | The initial color of each particle. |
| **Gravity Modifier** | Scales the gravity value set in the [Physics](class-PhysicsManager.md) window. A value of zero switches gravity off. |
| **Simulation Space** | Controls whether particles are animated in the parent object’s local space (therefore moving with the parent object), in the world space, or relative to a custom object (moving with a custom object of your choosing). |
| **Simulation Speed** | Adjust the speed at which the entire system updates. |
| **Delta Time** | Choose between **Scaled** and **Unscaled**, where **Scaled** uses the **Time Scale** value in the [Time](class-TimeManager.md) window, and **Unscaled** ignores it. This is useful for **Particle Systems** that appear on a Pause Menu, for example. |
| **Scaling Mode** | Choose how to use the scale from the transform. Set to **Hierarchy**, **Local** or **Shape**. Local applies only the Particle System transform scale, ignoring any parents. Shape mode applies the scale to the start positions of the particles, but does not affect their size. |
| **Play on Awake** | If enabled, the Particle System starts automatically when the object is created. |
| **Emitter Velocity** | Choose how the Particle System calculates the velocity used by the Inherit Velocity and Emission modules. The system can calculate the velocity using a **Rigidbody** component, if one exists, or by tracking the movement of the **Transform component**. If no Rigidbody component exists, the system uses its Transform component by default. |
| **Max Particles** | The maximum number of particles in the system at once. If the limit is reached, some particles are removed. |
| **Auto Random Seed** | If enabled, the Particle System looks different each time it is played. When set to false, the system is exactly the same every time it is played. |
| **Random Seed** | When disabling the automatic random seed, this value is used to create a unique repeatable effect. |
| **Stop Action** | When all the particles belonging to the system have finished, it is possible to make the system perform an action. A system is determined to have stopped when all its particles have died, and its age has exceeded its Duration. For looping systems, this only happens if the system is stopped via script. |
| Disable | The GameObject is disabled. |
| Destroy | The GameObject is destroyed. |
| Callback | The OnParticleSystemStopped callback is sent to any **scripts** attached to the GameObject. |
| **Culling Mode** | Choose whether to pause Particle System simulation when particles are offscreen. Culling when offscreen is most efficient, but you may want to continue simulation for off-one effects. |
| Automatic | Looping systems use **Pause**, and all other system use **Always Simulate**. |
| Pause And Catch-up | The system stops simulating while offscreen. When re-entering the view, the simulation performs a large step to reach the point where it would have been had it not paused. In complex systems, this option can cause performance spikes. |
| Pause | The system stops simulating while offscreen. |
| Always Simulate | The system processes its simulation on each frame, regardless of whether it is on screen or not. This can be useful for one-shot effects such as fireworks, where during the simulation would be obvious. |
| **Ring Buffer Mode** | Keeps particles alive until they reach the **Max Particles** count, at which point new particles recycle the oldest ones, instead of removing particles when their lifetimes elapse. |
| Disabled | Disable **Ring Buffer Mode**, so the system removes particles when their lifetime elapses. |
| Pause Until Replaced | Pauses old particles at the end of their lifetime until the **Max Particle** limit is reached, at which point the system recycles them, so they reappear as new particles. |
| Loop Until Replaced | At the end of their lifetime, particles rewind back to the specified proportion of their lifetime until the **Max Particle** limit is reached, at which point the system recycles them, so they reappear as new particles. |