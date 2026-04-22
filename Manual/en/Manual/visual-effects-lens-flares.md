# Lens flares

Resources and techniques for creating lens flares lighting effects, which can add atmosphere to your **scene**.

![Lens flare](../uploads/Main/srp-lens-flare.png)

Lens flare

| **Topic** | **Description** |
| --- | --- |
| [Introduction to lens flare effects](lens-flare-introduction.md) | Understand how Unity manages lens flares, which simulate the effect of lights refracting inside a **camera** lens. |
| [Lens flares in URP](urp/shared/lens-flare/lens-flare.md) | Resources and techniques for creating and configuring lens flares in the Universal **Render Pipeline** (URP). |
| [Lens flares in the Built-In Render Pipeline](lens-flare-birp.md) | Resources and techniques for creating lens flares lighting effects in the Built-In Render Pipeline. |

## Render pipeline information

How you work with lens flare depends on the render pipeline you use.

| **Feature name** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Built-in Render Pipeline** |
| --- | --- | --- | --- |
| **Lens flares** | Yes  Use a [Lens Flare (SRP) Data Asset](urp/shared/lens-flare/lens-flare-asset.md) and a [Lens Flare (SRP) component](urp/shared/lens-flare/lens-flare-component.md), or a [Screen Space Lens Flare override](urp/shared/lens-flare/post-processing-screen-space-lens-flare.md). | Yes  Use a [Lens Flare (SRP) Data Asset](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/shared/lens-flare/lens-flare-asset.html) and a [Lens Flare (SRP) component](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/shared/lens-flare/lens-flare-component.html), or a [Screen Space Lens Flare override](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/shared/lens-flare/Override-Screen-Space-Lens-Flare). | Yes  Use a [Flare asset](class-Flare.md) and, optionally, a [Lens Flare component](class-LensFlare.md). |