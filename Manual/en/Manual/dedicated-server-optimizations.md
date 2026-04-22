# Dedicated Server optimizations

Dedicated Server build target performs optimizations for networked applications. The build target applies some optimizations automatically by default, while others are optional because their impact depends on the game’s implementation.

## Automatic optimizations

By default, the Dedicated Server build target performs the following optimizations:

* [Deactivates the Audio Subsystem](#AudioSubsystem)
* [Removes lighting threads](#LightingThreads)
* [Removes some Player loop callbacks](#PlayerLoopCallbacks)
* [Removes GPU-only assets](#GpuOnlyAssets)

### Audio Subsystem

The Dedicated Server build target deactivates the [Audio](AudioOverview.md) [Subsystem](../ScriptReference/Subsystem.md) because builds don’t need audio support when running as a dedicated server. Disabling the Audio Subsystem reduces CPU load.

### Lighting threads

The Dedicated Server build target removes process threads related to lighting because there’s no need to render lighting on a server build.

### Player loop callbacks

The Dedicated Server build target disables the following [PlayerLoop callbacks](../ScriptReference/LowLevel.PlayerLoop.md) because they aren’t necessary for a server build.

* Player update loop registration of [`SendMouseEvents`](../ScriptReference/PlayerLoop.PreUpdate.SendMouseEvents.md)
* Player update loop registration of [`UpdateAllRenderers`](../ScriptReference/PlayerLoop.PostLateUpdate.UpdateAllRenderers.md)
* Player update loop registration of [`PlayerUpdateCanvases`](../ScriptReference/PlayerLoop.PostLateUpdate.PlayerUpdateCanvases.md)
* Player update loop registration of all audio callbacks such as [`UpdateAudio`](../ScriptReference/PlayerLoop.PostLateUpdate.UpdateAudio.md)
* Player update loop registration of all UI callbacks such as [`UIElementsUpdatePanels`](../ScriptReference/PlayerLoop.PreLateUpdate.UIElementsUpdatePanels.md)
* Player update loop registration of all input callbacks such as [`UpdateInputManager`](../ScriptReference/PlayerLoop.EarlyUpdate.UpdateInputManager.md)

### GPU-only assets

The Dedicated Server build target removes GPU-only assets such as texture **pixel** data for textures and **mesh** vertex data, that the server doesn’t need. The build target preserves assets with CPU Read/Write access and assets in the [Resource Folders](LoadingResourcesatRuntime.md).

Textures and meshes imported with CPU Read/Write access disabled are only accessible by the GPU, whereas the CPU can’t access them. Because the Dedicated Server build target doesn’t initialize a graphics device, there’s no need to include this data. Excluding this data reduces the memory usage of the executable.

Refer to the following lists to learn more about which assets the Dedicated Server build target removes and preserves.

**Removed**:

* Textures with CPU Read/Write access disabled.
* Vertex data for meshes with CPU Read/Write access disabled.

**Preserved**:

* Textures with CPU Read/Write access enabled.
* Vertex data for meshes with CPU Read/Write access enabled.
* Assets in the protected [Resource Folders](LoadingResourcesatRuntime.md).
* Texture [metadata](AssetMetadata.md) (such as the texture size value).
* Mesh data that internal systems that run on the CPU require (such as [physics](PhysicsSection.md)), even if CPU Read/Write is disabled.

**Note**: To learn more about CPU Read/Write access, refer to [Texture.isReadable](../ScriptReference/Texture-isReadable.md) and [Mesh.isReadable](../ScriptReference/Mesh-isReadable.md).

## Additional optimizations

You also have an option to apply additional optimizations through [Enable Dedicated Server Optimizations](dedicated-server-player-settings.md#Optimization) property in the **Player Settings** window. These optimizations remove assets such as ****Shaders**** and **Fonts**. However, applying these optimizations might cause issues if the removed assets are referenced in the game’s implementation.

**Assets removed**:

* Shaders
* Fonts

In addition to the automatic optimizations applied through the Dedicated Server build target, you can apply the following implementation-specific optimizations manually.

* Use [conditional compilation](platform-dependent-compilation.md) to selectively include and exclude code depending on the build target.
* Separate player-specific and server-specific code through class implementations.
* Remove additional items from the PlayerLoop in server builds. Refer to [PlayerLoop](../ScriptReference/LowLevel.PlayerLoop.md) and [PlayerLoopSystem](../ScriptReference/LowLevel.PlayerLoopSystem.md).

## Known limitations

Dedicated server optimizations don’t apply to content built with the Scriptable Build Pipeline (SBP). Systems that rely on SBP, such as the Addressables system and Entities subscenes, don’t benefit from these optimizations.

## Additional resources

* [Optimization section](analysis.md)
* [Optimizations tips for Unity UI](https://unity.com/how-to/unity-ui-optimization-tips)
* [Best practices for performance optimization in Unity](https://unity.com/how-to/advanced-programming-and-code-architecture)