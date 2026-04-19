# Introduction to linear color space in Unity

Working in linear color space gives more accurate rendering than working in gamma color space.

You can work in linear color space if your Textures were created in linear or gamma color space. Gamma color space Texture inputs to the linear color space **Shader** program are supplied to the Shader with gamma correction removed from them.

For information on how to set the color space of your project, refer to [Set a project’s color space](set-project-color-space.md).

**Note**: The **Texture preview** window in [Texture Import Settings](../ScriptReference/TextureImporter.md) displays Textures using gamma blending even if you are working in linear color space.

## Linear Textures

Selecting **Color Space: Linear** assumes your Textures are in gamma color space. Unity uses your GPU’s sRGB sampler by default to crossover from gamma to linear color space. If your Textures are authored in linear color space, you need to bypass the sRGB sampling. For more information, refer to [Disable sRGB sampling for a texture](disable-srgb-sampling-textures.md).

## Gamma Textures

Crossing over from gamma color space to linear color space requires some tweaking. For more information, refer to [Gamma Textures in linear color space](gamma-textures-linear-color-space.md).

### Gamma to linear conversion

For colors, this conversion is applied implicitly, because the Unity Editor already converts the values to floating point before passing them to the GPU as constants. When sampling Textures, the GPU automatically removes the gamma correction, converting the result to linear space.

These inputs are then passed to the Shader, with lighting calculations taking place in linear space as they normally do. When writing the resulting value to a framebuffer, it is either gamma-corrected straight away or left in linear space for later gamma correction - this depends on the current rendering configuration. For example, in high dynamic range (HDR), rendering results are left in linear space and gamma corrected later.