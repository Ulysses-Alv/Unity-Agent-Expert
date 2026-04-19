# Cubemap texture Import Settings window reference

You can further refine **Cubemap** shape textures with the following properties:

Here’s your fixed Markdown table with properly aligned rows:

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Mapping** | - | Use **Mapping** to specify how the Texture is projected onto your **GameObject**. This is set to **Auto** by default. |
| **Mapping** | **Auto** | Unity tries to create the layout from the Texture information. |
| **Mapping** | **6 Frames Layout (Cubic Environment)** | The Texture contains six images arranged in one of the standard cubemap layouts: cross, or sequence (+x -x +y -y +z -z). The images can be orientated either horizontally or vertically. |
| **Mapping** | **Latitude Longitude (Cylindrical)** | Maps the Texture to a 2D Latitude-Longitude representation. |
| **Mapping** | **Mirrored Ball (Sphere Mapped)** | Maps the Texture to a sphere-like cubemap. |
| **Convolution Type** | - | Choose the type of pre-convolution (filtering) that you want to use for this Texture. The result of pre-convolution is stored in mipmaps.  This property is only available for the [Default](texture-type-default.md) Texture type. |
| **Convolution Type** | **None** | The Texture has no pre-convolution (no filtering). This is the default. |
| **Convolution Type** | **Specular (Glossy Reflection)** | Select this to use cubemaps as **Reflection Probes**. The Texture mipmaps are pre-convoluted (filtered) with the engine BRDF. For more information, see Wikipedia’s page on [Bidirectional reflectance distribution function](https://en.wikipedia.org/wiki/Bidirectional_reflectance_distribution_function). |
| **Convolution Type** | **Diffuse (Irradiance)** | The Texture is convoluted (filtered) to represent irradiance. This is useful if you use the cubemap as a **Light Probe**. |
| **Fixup Edge Seams** | - | This option is only available with the **None** or **Diffuse** convolution (filter). Use this on low-end platforms as a work-around for filtering limitations, such as cubemaps incorrectly filtered between faces. |

## Additional resources

* [Cubemaps](class-Cubemap-landing.md)