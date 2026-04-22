# Standard Shader Material Inspector window reference for the Built-In Render Pipeline

The standard **shader** presents you with a list of Material Properties. These Properties vary slightly depending on whether you have chosen to work in the Metallic workflow mode or the Specular workflow mode. Most of the Properties are the same across both modes, and this page covers all the Properties for both modes.

These Properties can be used together to recreate the look of almost any real-world surface.

![A Standard Shader material with default Properties and no values or textures assigned](../uploads/Main/StandardShaderNewEmptyMaterial.png)

A Standard Shader material with default Properties and no values or textures assigned

## Rendering Mode

The first Material Parameter in the Standard Shader is **Rendering Mode**. This allows you to choose whether the object uses transparency, and if so, which type of blending mode to use.

For more information, refer to [Change Rendering Mode in the Standard Shader using a script](StandardShaderMaterialParameterRenderingMode.md).

* **Opaque** - Is the default, and suitable for normal solid objects with no transparent areas.
* **Cutout** - Allows you to create a transparent effect that has hard edges between the opaque and transparent areas. In this mode, there are no semi-transparent areas, the texture is either 100% opaque, or invisible. This is useful when using transparency to create the shape of materials such as leaves, or cloth with holes and tatters.
* **Transparent** - Suitable for rendering realistic transparent materials such as clear plastic or glass. In this mode, the material itself will take on transparency values (based on the texture’s alpha channel and the alpha of the tint colour), however reflections and lighting highlights will remain visible at full clarity as is the case with real transparent materials.
* **Fade** - Allows the transparency values to entirely fade an object out, including any specular highlights or reflections it may have. This mode is useful if you want to animate an object fading in or out. It is not suitable for rendering realistic transparent materials such as clear plastic or glass because the reflections and highlights will also be faded out.

## Albedo

The **Albedo** property controls the base color and transparency of the material. For more information, refer to [Albedo](StandardShaderMaterialParameterAlbedoColor.md).

## Specular

The **Specular** Property is only visible when using the **Specular setup**. For more information, refer to [Metallic and specular workflows](StandardShaderMetallicVsSpecular.md).

## Metallic

When working in the Metallic workflow (as opposed to the Specular workflow), the reflectivity and light response of the surface are modified by the **Metallic** level and the **Smoothness** level. For more information, refer to [Metallic and specular workflows](StandardShaderMetallicVsSpecular.md).

### Smoothness texture map

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Smoothness source** | Select the texture channel where the smoothness value is stored. |
| *Specular/Metallic Alpha* | Use the Alpha channel of the Specular or Metallic map (depending which of these two modes you are using) as the Smoothness value. Because the smoothness of each point on the surface is a single value, only a single channel of an image texture is required for the data. Unity assumes the smoothness data is in the Alpha Channel of the texture used for the Metallic or Specular texture map . |
| *Albedo Alpha* | Use the Alpha channel of the Albedo map as the Smoothness value. This lets you reduce the total number of textures, or use textures of different resolutions for the Smoothness value and Specular/Metallic map. |

You can toggle highlights and reflections with the settings found in the ****Forward Rendering** Options** section of the Material when it is open in the **Inspector**.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Highlights** | Check this box to disable highlights. This is an optional performance optimization for mobile. It removes the calculation of highlights from the Standard Shader. How this affects the appearance mainly depends on the Specular/Metallic value and the Smoothness. |
| **Reflections** | Check this box to disable environment reflections. This is an optional performance optimization for mobile. It removes the calculation of highlights from the Standard Shader. Instead of sampling the environment map, an approximation is used. How this affects the appearance depends on the smoothness. |

## Normal Map

Normal maps are a type of **Bump Map**. They are a special kind of texture that allow you to add surface detail such as bumps, grooves, and scratches to a model which catch the light as if they are represented by real geometry.

For more information, refer to [Normal Map (Bump mapping)](StandardShaderMaterialParameterNormalMap.md).

## Height Map

Height mapping (also known as parallax mapping) is a similar concept to normal mapping, however this technique is more complex - and therefore also more performance-expensive. For more information, refer to [Heightmap](StandardShaderMaterialParameterHeightMap.md).

## Occlusion

The occlusion map is used to provide information about which areas of the model should receive high or low indirect lighting. For more information, refer to [Occlusion Map](StandardShaderMaterialParameterOcclusionMap.md).

## Emission

For more information, refer to [Emission](StandardShaderMaterialParameterEmission.md).

| **Property** | **Description** |
| --- | --- |
| **Color** | Specifies the color and intensity of the emission. Click the **Color** box to open the ****HDR** Color** picker. Here you can alter the color of the illumination and the **Intensity** of the emission. To specify which areas of the Material emit light, you can assign an emission map to this property. If you do this, Unity uses the full color values of the map for the emission color and brightness. You can also use the HDR color picker to tint the map and change the emission intensity. |
| **Global Illumination** | Specifies how the light that this Material emits affects the contextual lighting of other nearby GameObjects. There are three options: • **Realtime**: Unity adds the emissive light from this Material to the **Realtime** Global Illumination calculations for the Scene. This means that this emissive light affects the illumination of nearby GameObjects, including ones that are moving. • **Baked**: Unity bakes the emissive light from this Material into the static Global Illumination lighting for the Scene. This Material affects the lighting of nearby static GameObjects, but not dynamic GameObjects. However, Light Probes still affect the lighting of dynamic GameObjects. • **None**: The emissive light from this Material does not affect Realtime lightmaps, Baked lightmaps, or Light Probes in the Scene. It does not illuminate or affect other GameObjects. The Material itself does have the emission color. |

## Secondary Maps

Secondary Maps (or Detail maps) allow you to overlay a second set of textures on top of the main textures listed above. For more information, refer to [Secondary Maps (Detail Maps) & Detail Mask](StandardShaderMaterialParameterDetail.md).