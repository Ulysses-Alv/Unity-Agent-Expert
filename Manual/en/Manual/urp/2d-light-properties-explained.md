# Configure a 2D light

Understand how the different [Light 2D component](2DLightProperties.md) properties interact and affect the appearance and behavior of a 2D light.

## Intensity

Light intensity are available to all types of Lights. Color adjusts the lights color, while intensity allows this color to go above 1. This allows lights which use multiply to brighten a **sprite** beyond its original color.

## Light Order

The **Light Order** value determines the position of the Light in the Render queue relative to other Lights that target the same sorting layer(s). Lower numbered Lights are rendered first, with higher numbered Lights rendered above those below. This especially affects the appearance of blended Lights when **Overlap Operation** is set to **Alpha Blend**.

## Overlap Operation

This property controls the way in the selected Light interacts with other rendered Lights. You can toggle between the two modes by enabling or disabling this property. The effects of both modes are shown in the examples below:

| Overlap Operation set to Additive | Overlap Operation set to Alpha Blend |
| --- | --- |
| **Overlap Operation** set to **Additive** | **Overlap Operation** set to **Alpha Blend** |

When **Overlap Operation** is set to **Additive**, the Light is blended with other Lights additively, where the **pixel** values of intersecting Lights are added together. This is the default Light blending behavior.

When **Overlap Operation** is set to **Alpha Blend**, Lights are blended together based on their alpha values. This can be used to completely overwrite one Light with another where they intersect, but the render order of the Lights is also dependent on the [Light Order](#light-order) of the different Lights.

## Use Normal Map

All lights except for global lights can be toggled to use the **normal maps** in the sprites material. When enabled, Distance and Accuracy will be visible as new properties. For more information, refer to [Add a normal map or a mask map to a sprite in URP](SecondaryTextures.md).

## Quality

Light quality allows the developer to choose between performance and accuracy. When choosing performance, artefacts may occur. Smaller lights and larger distance values will reduce the difference between fast and accurate.

## Distance

Distance controls the distance between the light and the surface of the Sprite, changing the resulting lighting effect. This distance does not affect intensity, or transform the position of the Light in the **scene**.

The following examples show the effects of changing the Distance values.

| Distance: 0.5 | Distance: 2 | Distance: 8 |
| --- | --- | --- |
| **Distance**: 0.5 | **Distance**: 2 | **Distance**: 8 |

## Volume Opacity

Volumetric lighting is available to all Light types. Use the **Volume Opacity** slider to control the visibility of the volumetric light. At a value of zero, no Light volume is shown while at a value of one, the Light volume appears at full opacity.

## Shadow Intensity

The Shadow Intensity property controls the amount of light that **Shadow Caster 2Ds** block from the Light source which affects the intensity of their shadows. This is available on all non global Light types. Use this slider to control the amount of light that Shadow Caster 2Ds block when they interact with or block this Light.

The slider ranges from 0 to 1. At 0, Shadow Caster 2Ds do not block any light coming from the Light source and they create no shadows. At the maximum value of 1, Shadow Caster 2Ds block all light from the Light source and create shadows at full intensity.

| Shadow Intensity = 0.0 | Shadow Intensity = 0.5 | Shadow Intensity = 1.0 |
| --- | --- | --- |
| Shadow Intensity = 0.0 | Shadow Intensity = 0.5 | Shadow Intensity = 1.0 |

## Shadow Volume Intensity

Shadow Volume Intensity determines the amount of volumetric light **Shadow Caster 2Ds** block from the Light source. It is available on all non global lights, and when Volume Opacity is above zero. Use this slider to control the amount of volumetric light that Shadow Caster 2Ds block when they interact with or block this Light.

The slider ranges from 0 to 1. At 0, Shadow Caster 2Ds do not block any light coming from the Light source and they create no shadows. At the maximum value of 1, Shadow Caster 2Ds block all light from the Light source and create shadows at full intensity.

## Target Sorting Layers

Lights only light up the Sprites on their targeted sorting layers. Select the desired sorting layers from the drop-down menu for the selected Light. To add or remove sorting layers, refer to the [Tag Manager - Sorting Layers](https://docs.unity3d.com/Manual/class-TagManager.html#SortingLayers) for more information.

## Additional resources

* [Light 2D component reference](2DLightProperties.md)
* [Light Blend Styles component reference for URP](LightBlendStyles.md)