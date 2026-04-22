# Limit Velocity over Lifetime module reference

This module controls how the speed of **particles** is reduced over their lifetime.

## Properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| **Property** | **Function** |
| --- | --- |
| **Separate Axes** | Splits the axes up into separate X, Y and Z components. |
| **Speed** | Sets the speed limit of the particles. |
| **Space** | Selects whether the speed limitation refers to local or world space. This option is only available when **Separate Axes** is enabled. |
| **Dampen** | The fraction by which a particle’s speed is reduced when it exceeds the speed limit. |
| **Drag** | Applies linear drag to the particle velocities. |
| **Multiply by Size** | When enabled, larger particles are affected more by the drag coefficient. |
| **Multiply by Velocity** | When enabled, faster particles are affected more by the drag coefficient. |

## Additional resources

* [Particle velocity](particle-velocity.md)
* [Create particles that change velocity over time](create-particles-that-change-velocity-over-time.md)