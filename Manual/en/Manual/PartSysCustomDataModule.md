# Custom Data module reference

View and edit the settings you can use to define custom data formats for **particles**.

Use the Custom Data module to define two additional channels (`Custom1` and `Custom2`) of data in the Unity Editor for each particle. You can use custom data to specify additional particle properties to control gameplay effects (for example, the health of a player) or particle appearance in a custom **shader** (for example, the temperature of a surface). For more information, refer to [Define custom data formats for particles](define-custom-data-formats-particles.md).

**Note**: If you want to set your custom data from a script, disable the **Custom Data** module.

## Custom Data module properties

Each channel in the Custom Data module contains the following settings.

| **Property** | **Description** |
| --- | --- |
| **Mode** | Specifies the format of your custom data. The following options are available:  * **Disabled:** Disables the Custom Data module. * **Vector:** Defines up to four [MinMaxCurve](../ScriptReference/ParticleSystem.MinMaxCurve.md) components. * **Color:** Defines an HDR-enabled [MinMaxGradient](../ScriptReference/ParticleSystem.MinMaxGradient.md) using a fixed color or gradient. |
| **Number of Components** | Defines how many components to include in the vector. The minimum value is 1 and the maximum value is 4.    This property is available only when **Mode** is **Vector**. |
| **X** | Defines how the `x` component of the custom data changes over the particle’s lifetime. This corresponds with the `x` channel in a `Vector4`. For the X component dropdown options, refer to [X, Y, Z, and W component dropdown](#vector-component-dropdown).  This property is available only when **Mode** is **Vector** and **Number of Components** is **1** or higher. |
| **Y** | Defines how the `y` component of the custom data changes over the particle’s lifetime. This corresponds with the `y` channel in a `Vector4`. For the Y component dropdown options, refer to [X, Y, Z, and W component dropdown](#vector-component-dropdown).   This property is available only when **Mode** is **Vector** and **Number of Components** is **2** or higher. |
| **Z** | Defines how the `z` component of the custom data changes over the particle’s lifetime. This corresponds with the `z` channel in a `Vector4`. For the Z component dropdown options, refer to [X, Y, Z, and W component dropdown](#vector-component-dropdown).    This property is available only when **Mode** is **Vector** and **Number of Components** is **3** or higher. |
| **W** | Defines how the `w` component of the custom data changes over the particle’s lifetime. This corresponds with the `w` channel in a `Vector4`. For the W component dropdown options, refer to [X, Y, Z, and W component dropdown](#vector-component-dropdown).   This property is available only when **Mode** is **Vector** and **Number of Components** is **4**. |
| **Color** | Defines an HDR [color or gradient](InspectorColorPicker.md) for the particle. The following options are available:  * **Color:** Uses a fixed HDR color value. * **Gradient:** Applies a color gradient over the particle’s lifetime. The color changes from the start to the end of the particle’s life according to the gradient. * **Random Between Two Colors:** Uses a random HDR color between two color values. * **Random Between Two Gradients:** Applies a random gradient between two color gradients. The color changes from the start to the end of the particle’s life according to the gradient.  This property is available only when **Mode** is **Color**. |

#### X, Y, Z, and W component dropdown

| **Property** | **Description** |
| --- | --- |
| **Constant** | Uses a fixed value throughout the particle’s lifetime. |
| **Curve** | Uses a single [curve](EditingCurves.md) to change the value over the particle’s lifetime. |
| **Random Between Two Constants** | Uses a random fixed value between two constants throughout the particle’s lifetime. |
| **Random Between Two Curves** | Uses a random [curve](EditingCurves.md) between two curves to change the value over the particle’s lifetime. |

## Additional resources

* [Define custom data formats for particles](define-custom-data-formats-particles.md)
* [Particle System custom vertex streams](PartSysVertexStreams.md)
* [ParticleSystem-customData API](../ScriptReference/ParticleSystem-customData.md)
* [ParticleSystem.CustomDataModule API](../ScriptReference/ParticleSystem.CustomDataModule.md)
* [ParticleSystem.SetCustomParticleData API](../ScriptReference/ParticleSystem.SetCustomParticleData.md)