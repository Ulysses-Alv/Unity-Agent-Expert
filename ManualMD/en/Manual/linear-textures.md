# Linear textures

sRGB sampling allows the Unity Editor to render **Shaders** in linear color space when Textures are in gamma color space. When you select to work in linear color space, the Editor defaults to using sRGB sampling. If your [Textures](Textures.md) are in linear color space, you need to work in linear color space and disable sRGB sampling for each Texture. To learn how to do this, see [Disable sRGB sampling for linear textures](disable-srgb-sampling-textures.md), below.

## Legacy GUI

Rendering of elements of the [Legacy GUI System](http://docs.unity3d.com/Manual/GUIScriptingGuide.html) is always done in gamma space. This means that for the legacy GUI system, Textures with their **Texture Type** set to **Editor GUI and Legacy GUI** do not have their gamma removed on import.

## Linear authored Textures

It is also important that lookup Textures, masks, and other textures with RGB values that mean something specific and have no gamma correction applied to them bypass sRGB sampling. This prevents values from the sampled Texture having non-existent gamma correction removed before they are used in the Shader, with calculations made with the same value as is stored on disk. Unity assumes that GUI textures and **normal map** textures are authored in a linear space.

For information on how to bypass sRGB sampling, refer to [Disable sRGB sampling for linear textures](disable-srgb-sampling-textures.md).