# Particle movement patterns

Use the [Noise](PartSysNoiseModule.md) module to add turbulence to **particle** movement.

Adding noise to your particles is a simple and effective way to create interesting patterns and effects. For example, imagine how embers from a fire move around, or how smoke swirls as it moves. Strong, high frequency noise could be used to simulate the fire embers, while soft, low frequency noise would be better suited to modeling a smoke effect.

For maximum control over the noise, you can enable the Separate Axes option. This allows you to control the strength and remapping on each axis independently.

The noise algorithm used is based on a technique called Curl Noise, which internally uses multiple samples of Perlin Noise to create the final noise field.

The settings on the [Quality](class-QualitySettings.md) window control how many unique noise samples are generated. When using Medium and Low, less samples of Perlin Noise are used, and those samples are re-used across multiple axes but combined in a way to try and hide the re-use. This means that the noise may look less dynamic and diverse when using lower quality settings. However, there is a significant performance benefit when using lower quality settings.