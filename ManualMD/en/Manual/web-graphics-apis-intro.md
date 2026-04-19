# Recommendations for both WebGL2 and WebGPU APIs

If you use either of the WebGL2 or WebGPU APIs, there are recommendations and restrictions you need be aware of:

* [Recommendations for how to use fonts in your Web application](#recommendations-for-how-to-use-fonts-in-your-web-application).
* [Restrictions on global illumination in Web](#restrictions-on-global-illumination-in-web).
* [Recommendations on how to use video in Web](#recommendations-on-how-to-use-video-in-web).
* [Restrictions on shader variants in Web](#information-about-shader-variants-in-web).

**Note:** There are further restrictions and limitations that are unique to each API that you need to be aware of. For more information, refer to [WebGL2](WebGL2.md) or [Limitations of WebGPU](WebGPU-limitations.md).

## Recommendations for how to use fonts in your Web application

Unity Web supports dynamic font rendering similar to other Unity platforms. However, Unity Web doesn’t have access to the fonts installed on the user’s machine, so if you want to use any fonts, include them in the project folder (including any fallback fonts for international characters, or bold/italic versions of fonts), and set as fallback font names.

## Restrictions on global illumination in Web

Unity Web only supports [Baked Global Illumination](LightingInUnity.md#globalIllumination). Realtime **Global Illumination** isn’t currently supported in Web. Also, Unity Web supports non-directional **lightmaps** only.

## Recommendations on how to use video in Web

You can’t use `VideoClipImporter` to import video clips to your Unity project because it might increase the initial asset data download size and prevent network streaming. To use video playback in Web:

1. Use the `URL` option in the VideoPlayer component.
2. Place the asset in the `StreamingAssets` directory to use the built-in network streaming of your browser.

## Recommendations for shader variants in Web

Due to limited available memory in Web, don’t include unwanted **shader** variants because it can lead to unnecessary memory usage. Therefore, it’s recommended to familiarize yourself with [shader variants](shader-variants.md) and [shader stripping](shader-variant-stripping.md). Also, take extra care to ensure that you don’t add shaders with too many variants (for example, Unity’s [Standard Shader](shader-StandardShader.md)) to the [Always-included Shaders](class-GraphicsSettings.md#Always) section in [Graphics Settings](class-GraphicsSettings.md).

## Additional resources

* [WebGL2](WebGL2.md)
* [WebGPU (Experimental)](WebGPU.md)