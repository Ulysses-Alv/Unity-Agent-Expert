# Particle System component reference

[Switch to Scripting](../ScriptReference/ParticleSystem.md "Go to ParticleSystem page in the Scripting Reference")

A **Particle System** component simulates fluid entities such as liquids, clouds and flames by generating and animating large numbers of small 2D images in the **scene**. For a full introduction to **particle** systems and their uses, see further documentation on [Particle Systems](ParticleSystems.md).

| **Property** | **Function** |
| --- | --- |
| **Simulate Layers** | Allows you to preview Particle Systems that are not selected. By default, only selected Particle Systems play in the **Scene View**. However, when you set Simulate Layers to anything other than **Nothing**, effects that match the **Layer Mask** play automatically, without you needing to select them. This is particularly useful for previewing environmental effects. |
| **Resimulate** | When this property is enabled, the Particle System immediately applies property changes to particles it has already generated. When disabled, the Particle System leaves existing particles as they are, and only applies property changes to new particles. |
| **Show Bounds** | When this property is enabled, Unity displays the **bounding volume** around the selected Particle Systems. These bounds determine whether a Particle System is currently on screen or not. |
| **Show Only Selected** | When this property is enabled, Unity hides all non-selected Particle Systems, allowing you to focus on producing a single effect. |

## Additional resources

* [Particle System module component reference](ParticleSystemModules.md)

ParticleSystem