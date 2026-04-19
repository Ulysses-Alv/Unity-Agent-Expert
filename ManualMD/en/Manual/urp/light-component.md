# Light component Inspector window reference for URP

This page contains information on Light components in the Universal **Render Pipeline** (URP). For a general introduction to lighting in Unity and examples of common lighting workflows, refer to [Lighting](../Lighting.md).

When using a preset of a **Light** component, only a subset of properties are supported. Unsupported properties are hidden.

## Properties

The **Light **Inspector**** window includes the following groups of properties:

* [General](#General)
* [Shape](#Shape)
* [Emission](#Emission)
* [Rendering](#Rendering)
* [Shadows](#Shadows)

### General

| Property: | Function: |
| --- | --- |
| **Type** | The current type of light. Possible values are **Directional**, **Point**, **Spot** and **Area**. |
| **Mode** | Specify the [**Light Mode**](https://docs.unity3d.com/Manual/LightModes.html) used to determine if and how a light is “baked”.  Options:  * **Realtime** * **Mixed** * **Baked**   **Note**: If **Type** is set to **Area**, this property is automatically set to **Baked**. |
| **Rendering Layers** | Set which rendering layers the light applies to.  **Note**: This property is only available if **Mode** is set to **Realtime** or **Mixed**. |

### Shape

| Property: | Function: |
| --- | --- |
| **Inner/Outer Spot Angle** | The inner and outer angles (in degrees) at the base of a spot light’s cone (spot light only). Avoid setting the **Outer Spot Angle** property to excessively high values (for example, higher than 160 degrees). Very high outer spot angle values cause Unity to spread the shadow map over a large area, which reduces shadow quality. |
| **Shape** | The shape of the area light.  Available options:  * **Rectangle** * **Disc** |
| **Width** | The width of the area light.  **Note**: This property is only available if **Shape** is set to **Rectangle**. |
| **Height** | The height of the area light.  **Note**: This property is only available if **Shape** is set to **Rectangle**. |
| **Radius** | The radius of the area light  **Note**: This property is only available if **Shape** is set to **Disc**. |

### Emission

| Property: | Function: |
| --- | --- |
| **Light Appearance** | Select the method used to create the color of the light.  Available options:  * **Color** * **Filter and Temperature** |
| **Color** | The color of the emitted light. Set this property with the color slider.  **Note**: This property is only available if **Light Apperance** is set to **Color**. |
| **Filter** | The color of the tint for the light source. Set this property with the color slider.  **Note**: This property is only available if **Light Apperance** is set to **Filter and Temperature**. |
| **Temperature** | The temperature (in Kelvin) of the light. Set this property with the slider or enter a specific value.  **Note**: This property is only available if **Light Apperance** is set to **Filter and Temperature**. |
| **Intensity** | Set the brightness of the light. The default value for a **Directional** light is 0.5. The default value for a **Point**, **Spot** or **Area** light is 1. |
| **Indirect Multiplier** | Use this value to vary the intensity of indirect light. Indirect light is light that has bounced from one object to another. The **Indirect Multiplier** defines the brightness of bounced light calculated by the **global illumination** (GI) system. If you set **Indirect Multiplier** to a value lower than **1,** the bounced light becomes dimmer with every bounce. A value higher than **1** makes light brighter with each bounce. This is useful, for example, when a dark surface in shadow (such as the interior of a cave) needs to be brighter in order to make detail visible. |
| **Range** | Define how far the light emitted from the center of the object travels (**Point** and **Spot** lights only). |
| **Cookie** | The RGB texture this Light projects into the scene. Use cookies to create silhouettes or patterned illumination. The texture format to use depends on the type of Light:  • Directional: 2D texture  • Spot: 2D texture  • Point: [cubemap texture](https://docs.unity3d.com/Manual/class-Cubemap.html)  **Note**: URP doesn’t support light cookies for Area lights.  For more information about light cookies, refer to [Cookies](https://docs.unity3d.com/Manual/Cookies.html). |
| **Cookie Size** | The per-axis scale Unity applies to the cookie texture. Use this property to set the size of the cookie.  **Note**: This property is available only if you set **Type** to **Directional** and assign a texture to **Cookie**. |
| **Cookie Offset** | The per-axis offset Unity applies to the cookie texture. Use this property to move the cookie without moving the light itself. You can also animate this property to scroll the cookie.   **Note**: This property is available only if you set **Type** to **Directional** and assign a texture to **Cookie**. |

## Rendering

| Property: | Function: |
| --- | --- |
| ****Culling Mask**** | Use this to selectively exclude groups of objects from being affected by the light. For more information, refer to [Layers](https://docs.unity3d.com/Manual/Layers.html). |

## Shadows

| Property: | Function: |
| --- | --- |
| **Shadow Type** | Determine whether this light casts Hard Shadows, Soft Shadows, or no shadows at all. For information on hard and soft shadows, refer to documentation on [lights](https://docs.unity3d.com/Manual/class-Light.html). |
| **Baked Shadow Angle** | If **Type** is set to **Directional** and **Shadow Type** is set to **Soft Shadows**, this property adds some artificial softening to the edges of shadows and gives them a more natural look.  **Note**: This property is only available if **Mode** is set to **Mixed** or **Baked**. |
| **Baked Shadow Radius** | If **Type** is set to **Point** or **Spot** and **Shadow Type** is set to **Soft Shadows**, this property adds some artificial softening to the edges of shadows and gives them a more natural look.  **Note**: This property is only available if **Mode** is set to **Mixed** or **Baked**. |
| **Realtime Shadows** | These properties are available when **Shadow Type** is set to **Hard Shadows** or **Soft Shadows**. Use these properties to control real-time shadow rendering settings. |
| **Strength** | Use the slider to control how dark the shadows cast by this light are. The range is between 0 and 1. Default value: 1. |
| **Bias** | Controls whether to use shadow bias settings from the URP Asset, or whether to define custom shadow bias settings for this light. Possible values are **Use Pipeline Settings** or **Custom**. |
| **Depth** | Controls the distance at which the shadows will be pushed away from the light. Useful for avoiding false self-shadowing artifacts. This property is visible only when **Bias** is set to **Custom**. |
| **Normal** | Controls the distance at which the shadow casting surfaces will be shrunk along the surface normal. Useful for avoiding false self-shadowing artifacts. This property is visible only when **Bias** is set to **Custom**. |
| **Near Plane** | Use the slider to control the value for the near clip plane when rendering shadows, defined as a value between 0.1 and 10. This value is clamped to 0.1 units or 1% of the light’s **Range** property, whichever is lower. This is set to 0.2 by default. |
| **Soft** **Shadows** **Quality** | Select the soft shadows quality. With the **Use Pipeline Settings** option selected Unity uses the value from the URP Asset. Options **Low**, **Medium**, and **High** let you specify the soft shadow quality value for this light. For more information on the values, refer to the [Soft Shadows](universalrp-asset.md#soft-shadows) section. |
| **Custom Shadow Layers** | Enable to specify the layer for shadows from the light separately to the layer for the light itself.  **Note**: This property is only available if **Mode** is set to **Mixed** or **Baked**, and **Shadow Type** is set to **Hard Shadows** or **Soft Shadows**. |
| **Layer** | The layer for shadows from the light. |

## Additional resources

* [Light component Inspector window reference for the Built-In Render Pipeline](../class-Light.md)