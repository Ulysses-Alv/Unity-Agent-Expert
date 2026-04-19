# Cubemap Skybox Shader Material Inspector Window reference

For information on how to create a Material that uses this **skybox** **Shader**, as well as details on how to render the skybox in your **Scene**, see [Using skyboxes](skyboxes-using.md).

## Render pipeline compatibility

| **Feature** | **Built-in **Render Pipeline**** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** |
| --- | --- | --- | --- |
| **Cubemap skybox** | Yes | Yes | No |

## Properties

| **Property** | **Description** |
| --- | --- |
| **Tint Color** | The color to tint the skybox to. Unity adds this color to the Textures to change their appearance without altering the base Texture files. |
| **Exposure** | Adjusts the skybox’s exposure. This allows you to correct tonal values in the skybox Textures. Larger values produce a more exposed, seemingly brighter, skybox. Smaller values produce a less exposed, seemingly darker, skybox. |
| **Rotation** | The rotation of the skybox around the positive y-axis. This changes the orientation of your skybox and is useful if you want a specific section of the skybox to be behind a particular part of your Scene. |
| **Cubemap (HDR)** | The Cubemap Asset this Material uses to represent the sky. For information on how to create a Cubemap Asset from your input Textures, see [Cubemap Asset](class-Cubemap-landing.md). |
| **Render Queue** | Determines the order in which Unity draws **GameObjects**. For more information on **Render Queue**, see [SL-SubShaderTags](SL-SubShaderTags.md). |
| **Double Sided **Global Illumination**** | Specifies whether the **lightmapper** accounts for both sides of the geometry when it calculates Global Illumination. When `true`, if you use the Progressive Lightmapper, back faces bounce light using the same emission and albedo as front faces. |

## Additional resources

* [Configure a skybox with a Skybox Shader](skybox-shaders.md)
* [Cubemaps](class-Cubemap-landing.md)