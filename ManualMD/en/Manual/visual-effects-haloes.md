# Light halos

Strategies for creating glowing halos around light sources, to give the impression of small dust **particles** in the air, and add atmosphere to your **scene**.

![A Light with a separate Halo Component](../uploads/Main/HaloWindow.jpg)

A Light with a separate Halo **Component**

| **Topic** | **Description** |
| --- | --- |
| [Create and configure a halo light effect](create-configure-halo.md) | Create and customize a halo around a light source. |
| [Halo component reference](class-Halo.md) | Explore the properties for the Halo component to customize the appearance of a halo. |

## Render pipeline information

How you work with light halos depends on the **render pipeline** you use.

| **Feature name** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Built-in Render Pipeline** |
| --- | --- | --- | --- |
| **Halos** | Yes  Use a [Lens Flare (SRP) Data Asset](urp/shared/lens-flare/lens-flare-asset.md) and a [Lens Flare (SRP) component](urp/shared/lens-flare/lens-flare-component.md). | Yes  Use a [Lens Flare (SRP) Data Asset](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/shared/lens-flare/lens-flare-asset.html) and a [Lens Flare (SRP) component](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/shared/lens-flare/lens-flare-component.html), or a [Screen Space Lens Flare override](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/shared/lens-flare/Override-Screen-Space-Lens-Flare.html). | Yes  Use a [Halo component](class-Halo.md). |