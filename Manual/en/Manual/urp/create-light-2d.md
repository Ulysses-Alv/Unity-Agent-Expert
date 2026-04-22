# Prepare your project for 2D lighting

Follow the steps below to prepare your project for 2D lighting and 2D lighting effects.

The following is the general workflow to include 2D lighting in your project.

1. Prepare your **Sprites** for lighting. For details, see [Prepare and upgrade sprites for 2D lighting in URP](PrepShader.md).
2. Set up **normal map** and mask Textures. 2D Lights can interact with normal map and mask Textures linked to Sprites to create advanced lighting effects, such as [normal mapping](https://en.wikipedia.org/wiki/Normal_mapping). See [Add normal map and mask textures to a sprite in URP](SecondaryTextures.md).
3. Create a 2D Light **GameObject**; see [Light 2D component reference for URP](2DLightProperties.md).
4. Configure the 2D Renderer Data asset; see [Configuring the 2D Renderer Asset](2DRendererData-overview.md).
5. To define the shape and properties that a Light uses to determine the shadows it casts, use the [Shadow Caster 2D component](2DShadows.md).
6. (Optional) if you want to apply 2D Light effects to a **pixel** art game, see [Precise pixel scaling and rotation via the Pixel Perfect Camera in URP](2d-pixelperfect.md).

## Create a Light 2D GameObject

Create a **2D Light** GameObject by going to **GameObject > Light** and selecting one of the four available types:

* Sprite Light 2D
* Spot Light 2D
* Global Light 2D
* Freeform Light 2D

## Additional resources

* [Light 2D component reference for URP](2DLightProperties.md)
* [Configure a 2D light](2d-light-properties-explained.md)
* [Creative Core Lighting tutorial in Unity Learn](https://learn.unity.com/project/creative-core-lighting)