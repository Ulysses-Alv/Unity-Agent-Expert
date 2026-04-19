# Self-Illuminated Shader Family

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces these **shaders**.

The **Self-Illuminated** shaders will emit light only onto themselves based on an attached alpha channel. They do not require any Lights to shine on them to emit this light. Any vertex lights or **pixel** lights will simply add more light on top of the self-illumination.

This is mostly used for light emitting objects. For example, parts of the wall texture could be self-illuminated to simulate lights or displays. It can also be useful to light power-up objects that should always have consistent lighting throughout the game, regardless of the lights shining on it.

## [Self-Illuminated Vertex-Lit](shader-SelfIllumVertexLit.md)

![shader-SelfIllumVertexLit](../uploads/Shaders/Thumb-IllumVertex.jpg)

shader-SelfIllumVertexLit

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Illumination** texture with alpha channel for Illumination Map

[» More details](shader-SelfIllumVertexLit.md)

## [Self-Illuminated Diffuse](shader-SelfIllumDiffuse.md)

![shader-SelfIllumDiffuse](../uploads/Shaders/Thumb-IllumDiffuse.jpg)

shader-SelfIllumDiffuse

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Illumination** texture with alpha channel for Illumination Map

[» More details](shader-SelfIllumDiffuse.md)

## [Self-Illuminated Specular](shader-SelfIllumSpecular.md)

![shader-SelfIllumSpecular](../uploads/Shaders/Thumb-IllumSpec.jpg)

shader-SelfIllumSpecular

**Assets needed:**

* One **Base** texture with alpha channel for Specular Map
* One **Illumination** texture with alpha channel for Illumination Map

[» More details](shader-SelfIllumSpecular.md)

## [Self-Illuminated Bumped](shader-SelfIllumBumpedDiffuse.md)

![shader-SelfIllumBumpedDiffuse](../uploads/Shaders/Thumb-IllumBump.jpg)

shader-SelfIllumBumpedDiffuse

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Normal map** normal map with alpha channel for Illumination

[» More details](shader-SelfIllumBumpedDiffuse.md)

## [Self-Illuminated Bumped Specular](shader-SelfIllumBumpedSpecular.md)

![shader-SelfIllumBumpedSpecular](../uploads/Shaders/Thumb-IllumBumpSpec.jpg)

shader-SelfIllumBumpedSpecular

**Assets needed:**

* One **Base** texture with alpha channel for Specular Map
* One **Normal map** normal map with alpha channel for Illumination Map

[» More details](shader-SelfIllumBumpedSpecular.md)

## [Self-Illuminated Parallax](shader-SelfIllumParallaxDiffuse.md)

![shader-SelfIllumParallaxDiffuse](../uploads/Shaders/Thumb-IllumParallaxBump.jpg)

shader-SelfIllumParallaxDiffuse

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Normal map** normal map with alpha channel for Illumination Map & Parallax Depth combined

**Note:**
One consideration of this shader is that the **Bumpmap** texture’s alpha channel doubles as a Illumination and the Parallax Depth.

[» More details](shader-SelfIllumParallaxDiffuse.md)

## [Self-Illuminated Parallax Specular](shader-SelfIllumParallaxSpecular.md)

![shader-SelfIllumParallaxSpecular](../uploads/Shaders/Thumb-IllumParallaxBumpSpec.jpg)

shader-SelfIllumParallaxSpecular

**Assets needed:**

* One **Base** texture with alpha channel for Specular Map
* One **Normal map** normal map with alpha channel for Illumination Map & Parallax Depth combined

**Note:**
One consideration of this shader is that the **Bumpmap** texture’s alpha channel doubles as a Illumination and the Parallax Depth.

[» More details](shader-SelfIllumParallaxSpecular.md)

SelfIllumFamily