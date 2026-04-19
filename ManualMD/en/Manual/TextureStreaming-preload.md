# Preload mipmap levels

If you enable a **camera** at runtime, mipmap streaming needs time to stream the mipmap levels into memory.

You can use texture preloading to prevent this. Do the following:

1. Add a **Streaming Controller** component to the disabled camera. For more information about this component, refer to [Configure mipmap streaming](TextureStreaming-configure.md).
2. Call the [StreamingController.SetPreloading](../ScriptReference/StreamingController.SetPreloading.md) API on the camera to preload the mipmap levels.

Use [StreamingController.CancelPreloading](../ScriptReference/StreamingController.CancelPreloading.md) to cancel preloading.

You can use the following APIs after you enable preloading:

* [StreamingController.IsPreloading](../ScriptReference/StreamingController.IsPreloading.md) to check if the camera is preloading.
* [Texture.streamingTextureLoadingCount](../ScriptReference/Texture-streamingTextureLoadingCount.md) and [Texture.streamingTexturePendingLoadCount](../ScriptReference/Texture-streamingTexturePendingLoadCount.md) to check how many textures Unity is still loading mipmap levels for.

If these APIs return values that indicate Unity has finished preloading, you might need to wait a few frames before you enable the camera to make sure preloading has finished.