# Transparent Shader Family

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces these **shaders**.

The Transparent shaders are used for fully- or semi-transparent objects. Using the alpha channel of the **Base** texture, you can determine areas of the object which can be more or less transparent than others. This can create a great effect for glass, HUD interfaces, or sci-fi effects.

## [Transparent Vertex-Lit](shader-TransVertexLit.md)

![shader-TransVertexLit](../uploads/Shaders/Thumb-TransVertex.jpg)

shader-TransVertexLit

**Assets needed:**

* One **Base** texture with alpha channel for Transparency Map

[» More details](shader-TransVertexLit.md)

## [Transparent Diffuse](shader-TransDiffuse.md)

![shader-TransDiffuse](../uploads/Shaders/Thumb-TransDiffuse.jpg)

shader-TransDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for Transparency Map

[» More details](shader-TransDiffuse.md)

## [Transparent Specular](shader-TransSpecular.md)

![shader-TransSpecular](../uploads/Shaders/Thumb-TransSpec.jpg)

shader-TransSpecular

**Assets needed:**

* One **Base** texture with alpha channel for combined Transparency Map/Specular Map

**Note:**
One limitation of this shader is that the **Base** texture’s alpha channel doubles as a Specular Map for the Specular shaders in this family.

[» More details](shader-TransSpecular.md)

## [Transparent Normal mapped](shader-TransBumpedDiffuse.md)

![shader-TransBumpedDiffuse](../uploads/Shaders/Thumb-TransBump.jpg)

shader-TransBumpedDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for Transparency Map
* One **Normal map** normal map, no alpha channel required

[» More details](shader-TransBumpedDiffuse.md)

## [Transparent Normal mapped Specular](shader-TransBumpedSpecular.md)

![shader-TransBumpedSpecular](../uploads/Shaders/Thumb-TransBumpSpec.jpg)

shader-TransBumpedSpecular

**Assets needed:**

* One **Base** texture with alpha channel for combined Transparency Map/Specular Map
* One **Normal map** normal map, no alpha channel required

**Note:**
One limitation of this shader is that the **Base** texture’s alpha channel doubles as a Specular Map for the Specular shaders in this family.

[» More details](shader-TransBumpedSpecular.md)

## [Transparent Parallax](shader-TransParallaxDiffuse.md)

![shader-TransParallaxDiffuse](../uploads/Shaders/Thumb-TransParallaxBump.jpg)

shader-TransParallaxDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for Transparency Map
* One **Normal map** normal map with alpha channel for Parallax Depth

[» More details](shader-TransParallaxDiffuse.md)

## [Transparent Parallax Specular](shader-TransParallaxSpecular.md)

![shader-TransParallaxSpecular](../uploads/Shaders/Thumb-TransParallaxBumpSpec.jpg)

shader-TransParallaxSpecular

**Assets needed:**

* One **Base** texture with alpha channel for combined Transparency Map/Specular Map
* One **Normal map** normal map with alpha channel for Parallax Depth

**Note:**
One limitation of this shader is that the **Base** texture’s alpha channel doubles as a Specular Map for the Specular shaders in this family.

[» More details](shader-TransParallaxSpecular.md)

TransparentFamily