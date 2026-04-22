# Analyze mipmap streaming

The ****Scene**** view has a Debug Draw Mode that helps you visualize and debug [mipmap streaming](TextureStreaming.md) in your scene. To enable it, follow these steps:

1. Select **Debug Draw Mode** in the [Scene view Draw Modes overlay]((overlays-reference-draw-modes)).
2. Select **Texture Mipmap Streaming**.

The Debug Draw Mode tints **GameObjects** the following colours:

* Green if a texture uses fewer mipmap levels.
* Red if a texture uses fewer mipmap levels because mipmap streaming doesn’t have enough resources to load them all.
* Blue if a texture that doesn’t use mipmap streaming, or there’s no renderer calculating the mipmap levels.

**Note:** If you use the [MainTexture](../ScriptReference/Material-mainTexture.md) API to set a main texture, the texture won’t display in the Debug Draw Mode.

If your project uses a scriptable **render pipeline**, use the following to analyze mipmap streaming instead:

* [Rendering Debugger in URP](urp/features/rendering-debugger.md)
* [Rendering Debugger in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/rendering-debugger-window-reference)

## Get mipmap streaming information in a script

To get information about texture memory, use the following properties:

* [Texture.currentTextureMemory](../ScriptReference/Texture-currentTextureMemory.md)
* [Texture.desiredTextureMemory](../ScriptReference/Texture-desiredTextureMemory.md)
* [Texture.totalTextureMemory](../ScriptReference/Texture-totalTextureMemory.md)
* [Texture.targetTextureMemory](../ScriptReference/Texture-targetTextureMemory.md)
* [Texture.nonStreamingTextureMemory](../ScriptReference/Texture-nonStreamingTextureMemory.md)

To get information about the number of textures or renderers that mipmap streaming affects, use the following properties:

* [Texture.streamingMipmapUploadCount](../ScriptReference/Texture-streamingMipmapUploadCount.md)
* [Texture.nonStreamingTextureCount](../ScriptReference/Texture-nonStreamingTextureCount.md)
* [Texture.streamingTextureCount](../ScriptReference/Texture-streamingTextureCount.md)
* [Texture.streamingRendererCount](../ScriptReference/Texture-streamingRendererCount.md)

To get information about the mipmap levels for a texture, use the following properties:

* [Texture2D.desiredMipmapLevel](../ScriptReference/Texture2D-desiredMipmapLevel.md)
* [Texture2D.loadingMipmapLevel](../ScriptReference/Texture2D-loadingMipmapLevel.md)
* [Texture2D.loadedMipmapLevel](../ScriptReference/Texture2D-loadedMipmapLevel.md)