# Decals

Resources and techniques for projecting materials to act as decals that decorate the surface of other materials.

| **Topic** | **Description** |
| --- | --- |
| [Introduction to decals and projection](introduction-decals-projection.md) | Understand what decals are, how Unity uses projection to create them, and what you could use projection for. |
| [Decals in the Universal Render Pipeline](urp/renderer-feature-decal-landing.md) | Techniques for using a Decal Renderer Feature or a Decal Projector in the Universal **Render Pipeline** (URP). |
| [Decals in the Built-In Render Pipeline](decals-birp.md) | Techniques for using a Projector component in the Built-In Render Pipeline. |

## Render pipeline information

How you work with decals and projectors depends on the render pipeline you use.

| **Feature name** | **Built-in Render Pipeline** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** |
| --- | --- | --- | --- |
| **Decal and projectors** | Yes  Use the [Projector component](class-Projector.md). | Yes  Use the [Decal Renderer Feature](urp/renderer-feature-decal.md). | Yes  Use the [Decal Projector](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/decals.html). |