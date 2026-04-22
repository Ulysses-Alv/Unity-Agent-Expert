# Light Blend Styles component reference for URP

Find the **Light Blend Styles** settings in the [2D Renderer Data asset](2DRendererData-overview.md) to customize the way that the light interacts with **sprites** in the **scene**.

The asset can contain a total of four different **Light Blend Styles**, each with a different combination of Blend Style settings by default. All lights in the scene must pick from one of these available **Light Blend Styles**. When you create a light, it’s automatically assigned the first Blend Style listed in the asset.

| **Property** | **Description** |
| --- | --- |
| **Name** | The name that appears when choosing a Blend Style for a Light 2D. |
| **Mask Texture Channel** | Select the channel(s) that the mask texture affects in this **Blend Style**. These channels comprise the red green blue alpha (RGBA) color model.  * **R** - The red color channel. * **G** - The green color channel. * **B** - The blue color channel. * **A** - The alpha channel. |
| **Blend Mode** | A **Blend Mode** controls the way a sprite is lit by light. Select how a Light 2D is blended when using this Blend Style. Refer to [these examples](2d-light-blending.md) of the different **Blend Modes**.  * **Additive** - Select to set the **Blend Mode** to Additive. * **Multiply** - Select to set the **Blend Mode** to Multiply. * **Subtractive** - Select to set the **Blend Mode** to Subtractive. |

## Additional resources

* [Types of 2D lights](LightTypes.md)
* [Light 2D component reference](2DLightProperties.md)