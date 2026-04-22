# Noise module reference

Add turbulence to **particle** movement using this module.

## Properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| Property | Function |
| --- | --- |
| **Separate Axes** | Control the strength and remapping independently on each axis. |
| **Strength** | A curve that defines how strong the noise effect is on a particle over its lifetime. Higher values will make particles move faster and further. |
| **Frequency** | Low values create soft, smooth noise, and high values create rapidly changing noise. This controls how often the particles change their direction of travel, and how abrupt those changes of direction are. |
| **Scroll Speed** | Move the noise field over time to cause more unpredictable and erratic particle movement. |
| **Damping** | When enabled, strength is proportional to frequency. Tying these values together means the noise field can be scaled while maintaining the same behaviour, but at a different size. |
| **Octaves** | Specify how many layers of overlapping noise are combined to produce the final noise values. Using more layers gives richer, more interesting noise, but significantly adds to the performance cost. |
| **Octave Multiplier** | For each additional noise layer, reduce the strength by this proportion. |
| **Octave Scale** | For each additional noise layer, adjust the frequency by this multiplier. |
| **Quality** | Lower quality settings reduce the performance cost significantly, but also affect how interesting the noise looks. Use the lowest quality that gives you the desired behavior for maximum performance. |
| **Remap** | Remap the final noise values into a different range. |
| **Remap Curve** | The curve that describes how the final noise values are transformed. For example, you could use this to pick out the lower ranges of the noise field and ignore the higher ranges by creating a curve that starts high and ends at zero. |
| **Position Amount** | A multiplier to control how much the noise affects particle positions. |
| **Rotation Amount** | A multiplier to control how much the noise affects particle rotations, in degrees per second. |
| **Size Amount** | A multiplier to control how much the noise affects particle sizes. |

## Additional resources

* [Particle movement patterns](particle-movement-patterns.md)