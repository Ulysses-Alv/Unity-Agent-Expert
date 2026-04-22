# Enable mipmap streaming

To enable mipmap streaming, follow these steps:

1. Go to **Edit** > **Project Settings** > **Quality**.
2. In the **Textures** section, enable **Mipmap Streaming**.

By default, this enables mipmap streaming for all the **cameras** in your project.

In the Editor, mipmap streaming is active in both Edit mode and Play mode by default. For more information about configuring mipmap streaming in the Editor, refer to [Editor settings](class-EditorManager.md)

## Make a texture work with mipmap streaming

Follow these steps:

1. Select a texture asset in the **Project** window.
2. In the **Texture Import Settings** window, select **Advanced** to open advanced settings.
3. Ensure that **Generate Mipmap** is enabled.
4. Enable **Stream Mipmap Levels**.

If you build your project for Android, you must also follow these steps:

1. Open the [**Build Profiles**](build-profiles.md) window.
2. Set ****Compression** Method** to **LZ4** or **LZ4HC**.

### Lightmaps

To enable mipmap streaming for **lightmaps**, follow these steps:

1. Go to **Edit** > **Project Settings** > **Player**.
2. In the **Other Settings** section, enable **Lightmap Streaming**.

You can edit the mipmap streaming settings of the lightmap assets in the same way that you can edit settings of any other texture, but they reset to their default values when Unity regenerates the lightmaps. To solve this, you can tell Unity to apply these values when it generates the lightmaps. Use the **Lightmap Streaming Enabled** and **Streaming Priority** in the [Player settings window](class-PlayerSettings.md).

### Custom meshes

If you create a custom **mesh** in a Unity script, Unity doesn’t automatically calculate the UV density of the mesh. This means Unity loads the wrong mipmap levels.

To prevent this, do either of the following:

* Use the [Mesh.RecalculateUVDistributionMetrics](../ScriptReference/Mesh.RecalculateUVDistributionMetrics.md) API to manually calculate the UV density of the mesh after you create it.
* [Import the mesh as a model file](ImportingModelFiles.md) instead.

### Objects that don’t have a standard Renderer component

Unity can’t calculate the correct mipmap level for the following objects, because they don’t use a standard [Renderer](../ScriptReference/Renderer.md) component:

* [Decals](visual-effects-decals.md) textures.
* [Reflection Probe](ReflectionProbes.md) textures, because they use mipmap levels to store lookup tables for material roughness.
* **Sprites**, which use a [Sprite Renderer](sprite/renderer/sprite-renderer-reference.md) component.
* **Shaders** that use UV texture coordinates in a channel other than [Mesh.uv](../ScriptReference/Mesh-uv.md) (also called UV0).
* Shaders that change texture coordinates other than scale and translation.

For these objects, use the [Texture2D.requestedMipmapLevel](../ScriptReference/Texture2D-requestedMipmapLevel.md) API to override mipmap streaming and set mipmap levels manually. If you don’t do this, Unity uses low-resolution mipmap levels, which might display as blurry depending on the distance from the camera.

## Additional resources

* [Texture2D.streamingMipmaps](../ScriptReference/Texture2D-streamingMipmaps.md)
* [SystemInfo.supportsMipStreaming](../ScriptReference/SystemInfo-supportsMipStreaming.md)