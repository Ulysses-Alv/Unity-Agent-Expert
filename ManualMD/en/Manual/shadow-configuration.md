# Enable shadows

Each Light in a **scene** can cast the following types of shadows:

* [Real-time shadows](shadow-realtime.md). Unity stores the shadows cast from each Light in [shadow map textures](shadow-mapping.md). The textures update each frame; shadows move when the lights move.
* Baked shadows. The Unity [lightmapper](Lightmappers.md) stores the shadows received by each **GameObject** in static **lightmap** textures, shadow mask textures, or **Light Probes**. Shadows don’t move when the lights move.

To configure which type of shadows a Light casts, use the **Inspector** window of the [Light component](class-Light.md). For more information, refer to the following:

* [Setting a light as real-time or baked](LightModes-landing.md).
* [Lighting Mode](lighting-mode.md), if you set the ****Light Mode**** of a Light to Mixed.

## Set a GameObject to receive shadows

GameObjects receive real-time shadows by default.

To set a static GameObject so it receives shadows from baked shadows, follow these steps:

1. Select the GameObject.
2. In the Inspector window, in the [Mesh Renderer](class-MeshRenderer.md) component, enable **Contribute **Global Illumination****.
3. Set **Receive Global Illumination** to **Lightmaps**.

## Set a GameObject to shadow other GameObjects

Follow these steps:

1. Select the GameObject.
2. In the [Mesh Renderer](class-MeshRenderer.md) component of the Inspector window, set **Cast shadows** to **On**.

## Configure shadows

To configure shadows for your whole project or a specific scene, refer to the following:

* [Lighting window reference](lighting-window.md)
* [Graphics settings](class-GraphicsSettings.md)
* [Universal Render Pipeline (URP) asset reference](urp/universalrp-asset.md)
* [Universal Renderer asset](urp/urp-universal-renderer.md)

## Additional resources

* [Shadows in URP](urp/Shadows-in-URP.md)
* [Shadows in the High Definition Render Pipeline (HDRP)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.3/manual/shadows.html)
* [Use shadows in a custom URP shader](urp/use-built-in-shader-methods-shadows.md)