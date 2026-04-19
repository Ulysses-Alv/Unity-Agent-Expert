# Color Adjustments Volume Override reference for URP

Use this effect to tweak the overall tone, brightness, and contrast of the final rendered image.

![Scene without Color Adjustments effect.](../../uploads/urp/post-proc/color-adjustments-off.png)

Scene without Color Adjustments effect.

![Scene with Color Adjustments effect.](../../uploads/urp/post-proc/color-adjustments.png)

Scene with Color Adjustments effect.

## Properties

| **Property** | **Description** |
| --- | --- |
| **Post Exposure** | Adjusts the overall exposure of the scene in EV (not EV100). URP applies this after the HDR effect and before tonemapping, which means that it does not affect previous effects in the chain. |
| **Contrast** | Use the slider to expand or shrink the overall range of tonal values. Larger positive values expand the tonal range and lower negative values shrink the tonal range. |
| **Color Filter** | Use the color picker to select which color the Color Adjustment effect should use to multiply the render and tint the result. |
| **Hue Shift** | Use the slider to shift the hue of all colors. |
| **Saturation** | Use the slider to push the intensity of all colors. |