# Reflective Shader Family

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces these **shaders**.

**Reflective** shaders will allow you to use a **Cubemap** which will be reflected on your **mesh**. You can also define areas of more or less reflectivity on your object through the alpha channel of the **Base** texture. High reflectivity is a great effect for glosses, oils, chrome, etc. Low reflectivity can add effect for metals, liquid surfaces, or video monitors.

## [Reflective Vertex-Lit](shader-ReflectiveVertexLit.md)

![shader-ReflectiveVertexLit](../uploads/Shaders/Thumb-ReflVertex.jpg)

shader-ReflectiveVertexLit

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas
* One **Reflection** Cubemap for Reflection Map

[» More details](shader-ReflectiveVertexLit.md)

## [Reflective Diffuse](shader-ReflectiveDiffuse.md)

![shader-ReflectiveDiffuse](../uploads/Shaders/Thumb-ReflDiffuse.jpg)

shader-ReflectiveDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas
* One **Reflection** Cubemap for Reflection Map

[» More details](shader-ReflectiveDiffuse.md)

## [Reflective Specular](shader-ReflectiveSpecular.md)

![shader-ReflectiveSpecular](../uploads/Shaders/Thumb-ReflSpec.jpg)

shader-ReflectiveSpecular

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas & Specular Map combined
* One **Reflection** Cubemap for Reflection Map

**Note:**
One consideration for this shader is that the **Base** texture’s alpha channel will double as both the reflective areas and the Specular Map.

[» More details](shader-ReflectiveSpecular.md)

## [Reflective Normal mapped](shader-ReflectiveBumpedDiffuse.md)

![shader-ReflectiveBumpedDiffuse](../uploads/Shaders/Thumb-ReflBump.jpg)

shader-ReflectiveBumpedDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas
* One **Reflection** Cubemap for Reflection Map
* One **Normal map**, no alpha channel required

[» More details](shader-ReflectiveBumpedDiffuse.md)

## [Reflective Normal Mapped Specular](shader-ReflectiveBumpedSpecular.md)

![shader-ReflectiveBumpedSpecular](../uploads/Shaders/Thumb-ReflBumpSpec.jpg)

shader-ReflectiveBumpedSpecular

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas & Specular Map combined
* One **Reflection** Cubemap for Reflection Map
* One **Normal map**, no alpha channel required

**Note:**
One consideration for this shader is that the **Base** texture’s alpha channel will double as both the reflective areas and the Specular Map.

[» More details](shader-ReflectiveBumpedSpecular.md)

## [Reflective Parallax](shader-ReflectiveParallaxDiffuse.md)

![shader-ReflectiveParallaxDiffuse](../uploads/Shaders/Thumb-ReflParallaxBump.jpg)

shader-ReflectiveParallaxDiffuse

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas
* One **Reflection** Cubemap for Reflection Map
* One **Normal map**, with alpha channel for Parallax Depth

[» More details](shader-ReflectiveParallaxDiffuse.md)

## [Reflective Parallax Specular](shader-ReflectiveParallaxSpecular.md)

![shader-ReflectiveParallaxSpecular](../uploads/Shaders/Thumb-ReflParallaxBumpSpec.jpg)

shader-ReflectiveParallaxSpecular

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas & Specular Map
* One **Reflection** Cubemap for Reflection Map
* One **Normal map**, with alpha channel for Parallax Depth

**Note:**
One consideration for this shader is that the **Base** texture’s alpha channel will double as both the reflective areas and the Specular Map.

[» More details](shader-ReflectiveParallaxSpecular.md)

## [Reflective Normal mapped Unlit](shader-ReflectiveBumpedUnlit.md)

![shader-ReflectiveBumpedUnlit](../uploads/Shaders/Thumb-ReflBumpUnlit.jpg)

shader-ReflectiveBumpedUnlit

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas
* One **Reflection** Cubemap for Reflection Map
* One **Normal map**, no alpha channel required

[» More details](shader-ReflectiveBumpedUnlit.md)

## [Reflective Normal mapped Vertex-Lit](shader-ReflectiveBumpedVertexLit.md)

![shader-ReflectiveBumpedVertexLit](../uploads/Shaders/Thumb-ReflBumpVertex.jpg)

shader-ReflectiveBumpedVertexLit

**Assets needed:**

* One **Base** texture with alpha channel for defining reflective areas
* One **Reflection** Cubemap for Reflection Map
* One **Normal map**, no alpha channel required

[» More details](shader-ReflectiveBumpedVertexLit.md)

ReflectiveFamily