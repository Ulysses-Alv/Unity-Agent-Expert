# Lift Gamma Gain Volume Override reference for URP

This effect allows you to perform three-way color grading. The **Lift Gamma Gain** trackballs follow the [ASC CDL](https://en.wikipedia.org/wiki/ASC_CDL) standard. When you adjust the position of the point on the trackball, it shifts the hue of the image towards that color in the given tonal range. Use the different trackballs to affect different ranges within the image. Adjust the slider under the trackball to offset the color lightness of that range.

## Properties

| **Property** | **Description** |
| --- | --- |
| **Lift** | Use this to control the dark tones. This has a more exaggerated effect on shadows.  * Use the trackball to select which color URP should shift the hue of the dark tones to. * Use the slider to offset the color lightness of the trackball color. |
| **Gamma** | Use this to control the mid-range tones with a power function.  * Use the trackball to select which color URP should use to shift the hue of the mid-tones to. * Use the slider to offset the color lightness of the trackball color. |
| **Gain** | Use this to increase the signal and make highlights brighter.  * Use the trackball to select which color that URP uses to shift the hue of the highlights to. * Use the slider to offset the color lightness of the trackball color. |