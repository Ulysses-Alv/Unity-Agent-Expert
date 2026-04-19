# Particle System Force Field component reference

[Switch to Scripting](../ScriptReference/ParticleSystemForceField.md "Go to ParticleSystemForceField page in the Scripting Reference")

The ****Particle** System Force Field** component applies forces to particles in [Particle Systems](class-ParticleSystem.md).

## Shape

| **Property** | **Function** |
| --- | --- |
| Shape | Select the shape of the area of influence. |
| Start Range | Set the value for the inner point within the shape where the area of influence begins. |
| End Range | Set the value for the outer point of the shape where the area of influence ends. |
| Direction X, Y and Z | Set a linear force to apply to particles along the x-axis, y-axis and z-axis. The higher the value, the greater the force. You can specify a **constant force** or vary the force over time. See [Varying properties over time](varying-particle-system-properties-over-time.md) documentation for more information. |

## Gravity

| **Property** | **Function** |
| --- | --- |
| Strength | Set the amount of attraction that particles have towards the focal point within the shape. The higher the value, the greater the strength. You can specify a constant strength or vary the strength over time. For more information, see [Varying properties over time](varying-particle-system-properties-over-time.md) documentation. |
| Gravity Focus | Set the focal point for gravity to pull particles towards. A value of 0 attracts particles to the center of the shape, and a value of 1 attracts particles to the outer edge of the shape. |

## Rotation

| **Property** | **Function** |
| --- | --- |
| Speed | Set the speed for the Particle System to propel particles around the vortex, which is the center of the force field. The higher the value, the faster the speed. You can specify a constant speed or vary the speed over time. For more information, see the [Varying properties over time](varying-particle-system-properties-over-time.md) documentation. |
| Attraction | Set the strength that particles are dragged into the vortex motion. A value of 1 applies the maximum attraction, and a value of 0 applies no attraction. You can specify a constant attraction or vary the attraction over time. For more information, see the [Varying properties over time](varying-particle-system-properties-over-time.md) documentation. |
| Rotation Randomness | Set a random axes of the shape to propel particles around. A value of 1 applies maximum randomness, and a value of 0 applies no randomness. |

## Drag

| **Property** | **Function** |
| --- | --- |
| Strength | Set the strength of the drag effect which slows particles down. The higher the value, the greater the strength. You can specify a constant strength or vary the strength over time. For more information, see the [Varying properties over time](varying-particle-system-properties-over-time.md) documentation. |
| Multiply Drag by Size | Enable this checkbox to adjust the drag based on the size of the particles. |
| Multiply Drag by Velocity | Enable this checkbox to adjust the drag based on the velocity of the particles. |

## Vector Field

| **Property** | **Function** |
| --- | --- |
| Volume Texture | Select the texture of the vector field. |
| Speed | Set the speed of the multiplier to apply to particles traveling through the vector field. The higher the value, the faster the speed. You can specify a constant strength or vary the strength over time. See [Varying properties over time](varying-particle-system-properties-over-time.md). |
| Attraction | Set the strength at which Unity drags particles into the vector field motion. The higher the value, the greater the attraction. You can specify a constant attraction or vary the attraction over time. See [Varying properties over time](varying-particle-system-properties-over-time.md). |

ParticleSystemForceField