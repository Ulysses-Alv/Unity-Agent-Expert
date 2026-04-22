# Shadows Midtones Highlights Volume Override reference for URP

The **Shadows Midtones Highlights** effect separately controls the shadows, midtones, and highlights of the render. Unlike [Lift, Gamma, Gain](Post-Processing-Lift-Gamma-Gain.md), you can use this effect to precisely define the tonal range for shadows, midtones, and highlights.

## Properties

| **Property** | **Description** |
| --- | --- |
| **Shadows** | Use this to control the shadows.Use the trackball to select the color URP should shift the hue of the shadows to.Use the slider to offset the color lightness of the trackball color. |
| **Midtones** | Use this to control the midtones.Use the trackball to select the color URP should shift the hue of the midtones to.Use the slider to offset the color lightness of the trackball color. |
| **Highlights** | Use this to control the highlights.Use the trackball to select the color URP should shift the hue of the highlights to.Use the slider to offset the color lightness of the trackball color. |

### Graph view

This graph shows the overall contribution of the **Shadows** (blue), **Midtones** (green), and **Highlights** (yellow). This is useful to visualize the transitions between the tonal regions.

### Shadow Limits

| **Property** | **Description** |
| --- | --- |
| **Start** | Set the start point of the transition between the shadows and the midtones of the render. |
| **End** | Set the end point of the transition between the shadows and the midtones of the render. |

### Highlight Limits

| **Property** | **Description** |
| --- | --- |
| **Start** | Set the start point of the transition between the midtones and the highlights of the render. |
| **End** | Set the end point of the transition between the midtones and the highlights of the render. |