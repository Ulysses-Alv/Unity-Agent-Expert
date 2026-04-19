# Panoramic Skybox Shader Material Inspector Window reference

For information on how to create a Material that uses this **skybox** **Shader**, as well as details on how to render the skybox in your **Scene**, see [Using skyboxes](skyboxes-using.md).

## Render pipeline compatibility

| **Feature** | **Built-in **Render Pipeline**** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** |
| --- | --- | --- | --- |
| **Panoramic skybox** | Yes | Yes | No |

## Properties

| **Property** | **Description** |
| --- | --- |
| **Tint Color** | The color to tint the skybox to. Unity adds this color to the Textures to change their appearance without altering the base Texture files. |
| **Exposure** | Adjusts the skybox’s exposure. This allows you to correct tonal values in the skybox Textures. Larger values produce a more exposed, seemingly brighter, skybox. Smaller values produce a less exposed, seemingly darker, skybox. |
| **Rotation** | The rotation of the skybox around the positive y-axis. This changes the orientation of your skybox and is useful if you want a specific section of the skybox to be behind a particular part of your Scene. |
| **Spherical (HDR)** | The Texture this Material spherically wraps around the Scene to represent the sky.For information on how to create a **Cubemap** Asset from your input Textures, see [Cubemap Asset](class-Cubemap-landing.md). |
| **Mapping** | Specifies the method this Material uses to project the Texture to create the skybox. The options are:**6 sided**: Uses a net format to map the Texture to the skybox.**Latitude Longitude Layout**: Uses a cylindrical wrapping method to map the Texture to the skybox. |
| **Image Type** | Specifies the angle around the y-axis that this Material projects the skybox to. The options are:**180**: Draws the **Spherical** Texture as a hemisphere with the peak in the direction of the positive z-axis. To change which side of the Scene this Material draws the Texture to, modify the **Rotation** property. By default, the back of the skybox is black, but this Material can draw a duplicate of the **Spherical** Texture on the back instead. To do this, enable **Mirror on Back**.**360**: Draws the Texture as a full sphere representation that wraps around the entire Scene. |
| **- Mirror on Back** | Specifies whether the Material should duplicate the **Spherical** Texture on the back of the skybox instead of drawing it as black.This option only appears when **Image Type** is set to **180**. |
| **Render Queue** | Determines the order in which Unity draws **GameObjects**. For more information on **Render Queue**, see [SL-SubShaderTags](SL-SubShaderTags.md). |
| **Double Sided **Global Illumination**** | Specifies whether the **lightmapper** accounts for both sides of the geometry when it calculates Global Illumination. When `true`, if you use the Progressive Lightmapper, back faces bounce light using the same emission and albedo as front faces. |

## Additional resources

* [Configure a skybox with a Skybox Shader](skybox-shaders.md)