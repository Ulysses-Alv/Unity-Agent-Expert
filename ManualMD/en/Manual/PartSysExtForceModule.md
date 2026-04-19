# External Forces module reference

This property modifies the effect of **wind zones** and [Particle System Force Fields](class-ParticleSystemForceField.md) on **particles** emitted by the system.

## Properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| ***Property*** | ***Function*** |
| --- | --- |
| **Multiplier** | Scale value applied to wind zone forces. |
| **Influence Filter** | Choose whether to include Force Fields based on a **Layer Mask**, or via an explicit **List**. |
| **List** | Define an explicit list of Force Fields that can affect this **Particle System**. This appears when the **Influence Filter** is set to **List**. |
| **Influence Mask** | Use a Layer Mask to determine which Force Fields affect this Particle System. This appears when the **Influence Filter** is set to **Layer Mask**.   This property has the following options:  * **Nothing**: Automatically unticks all other options, turning them off * **Everything** (default): Automatically ticks all other options, turning them on * **Default** * **TransparentFX** * **Ignore [Raycast](https://docs.unity3d.com/Packages/com.unity.ugui@latest?subfolder=/manual/Raycasters.html)** * **Water** * **UI** * **[PostProcessing](PostProcessingOverview.md)** |

## Additional resources

* [Particle physics forces](particle-physics-forces.md)