# Normal Shader Family

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces these **shaders**.

These shaders are the basic shaders in Unity. They are not specialized in any way and should be suitable for most opaque objects. They are not suitable if you want your object to be transparent, emitting light etc.

## [Vertex Lit](shader-NormalVertexLit.md)

![shader-NormalVertexLit](../uploads/Shaders/Thumb-NormalVertex.jpg)

shader-NormalVertexLit

**Assets needed:**

* One **Base** texture, no alpha channel required

## [Diffuse](shader-NormalDiffuse.md)

![shader-NormalDiffuse](../uploads/Shaders/Thumb-NormalDiffuse.jpg)

shader-NormalDiffuse

**Assets needed:**

* One **Base** texture, no alpha channel required

## [Specular](shader-NormalSpecular.md)

![shader-NormalSpecular](../uploads/Shaders/Thumb-NormalSpec.jpg)

shader-NormalSpecular

**Assets needed:**

* One **Base** texture with alpha channel for Specular Map

## [Normal mapped](shader-NormalBumpedDiffuse.md)

![shader-NormalBumpedDiffuse](../uploads/Shaders/Thumb-NormalBump.jpg)

shader-NormalBumpedDiffuse

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Normal map**

## [Normal mapped Specular](shader-NormalBumpedSpecular.md)

![shader-NormalBumpedSpecular](../uploads/Shaders/Thumb-NormalBumpSpec.jpg)

shader-NormalBumpedSpecular

**Assets needed:**

* One **Base** texture with alpha channel for Specular Map
* One **Normal map**

## [Parallax](shader-NormalParallaxDiffuse.md)

![shader-NormalParallaxDiffuse](../uploads/Shaders/Thumb-NormalParallaxBump.jpg)

shader-NormalParallaxDiffuse

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Normal map**
* One **Height** texture with Parallax Depth in the alpha channel

## [Parallax Specular](shader-NormalParallaxSpecular.md)

![shader-NormalParallaxSpecular](../uploads/Shaders/Thumb-NormalParallaxBumpSpec.jpg)

shader-NormalParallaxSpecular

**Assets needed:**

* One **Base** texture with alpha channel for Specular Map
* One **Normal map**
* One **Height** texture with Parallax Depth in the alpha channel

## [Decal](shader-NormalDecal.md)

![shader-NormalDecal](../uploads/Shaders/Thumb-NormalDecal.jpg)

shader-NormalDecal

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Decal** texture with alpha channel for Decal transparency

## [Diffuse Detail](shader-NormalDiffuseDetail.md)

![shader-NormalDiffuseDetail](../uploads/Shaders/Thumb-NormalDiffuseDetail.jpg)

shader-NormalDiffuseDetail

**Assets needed:**

* One **Base** texture, no alpha channel required
* One **Detail** grayscale texture; with 50% gray being neutral color

NormalFamily