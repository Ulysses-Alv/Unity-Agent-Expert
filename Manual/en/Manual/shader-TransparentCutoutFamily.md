# Transparent Cutout Shader Family

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces these **shaders**.

The Transparent Cutout shaders are used for objects that have fully opaque and fully transparent parts (no partial transparency). Things like chain fences, trees, grass, etc.

## [Transparent Cutout Vertex-Lit](shader-TransCutVertexLit.md)

![shader-TransCutVertexLit](../uploads/Shaders/Thumb-TransCutoutVertex.jpg)

shader-TransCutVertexLit

**Assets needed:**

* One **Base** texture with alpha channel for Transparency Map

[» More details](shader-TransCutVertexLit.md)

## [Transparent Cutout Diffuse](shader-TransCutDiffuse.md)

![shader-TransCutDiffuse](../uploads/Shaders/Thumb-TransCutoutDiffuse.jpg)

shader-TransCutDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for Transparency Map

[» More details](shader-TransCutDiffuse.md)

## [Transparent Cutout Specular](shader-TransCutSpecular.md)

![shader-TransCutSpecular](../uploads/Shaders/Thumb-TransCutoutSpec.jpg)

shader-TransCutSpecular

**Assets needed:**

* One **Base** texture with alpha channel for combined Transparency Map/Specular Map

**Note:**
One limitation of this shader is that the **Base** texture’s alpha channel doubles as a Specular Map for the Specular shaders in this family.

[» More details](shader-TransCutSpecular.md)

## [Transparent Cutout Bumped](shader-TransCutBumpedDiffuse.md)

![shader-TransCutBumpedDiffuse](../uploads/Shaders/Thumb-TransCutoutBump.jpg)

shader-TransCutBumpedDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for Transparency Map
* One **Normal map** normal map, no alpha channel required

[» More details](shader-TransCutBumpedDiffuse.md)

## [Transparent Cutout Bumped Specular](shader-TransCutBumpedSpecular.md)

![shader-TransCutBumpedSpecular](../uploads/Shaders/Thumb-TransCutoutBumpSpec.jpg)

shader-TransCutBumpedSpecular

**Assets needed:**

* One **Base** texture with alpha channel for combined Transparency Map/Specular Map
* One **Normal map** normal map, no alpha channel required

**Note:**
One limitation of this shader is that the **Base** texture’s alpha channel doubles as a Specular Map for the Specular shaders in this family.

[» More details](shader-TransCutBumpedSpecular.md)

TransparentCutoutFamily