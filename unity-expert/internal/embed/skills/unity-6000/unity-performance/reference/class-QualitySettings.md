# Quality settings tab reference

[Switch to Scripting](../ScriptReference/QualitySettings.md "Go to QualitySettings page in the Scripting Reference")

To configure the levels of graphical quality that Unity uses to render your project for different platforms, go to **Edit** > **Project Settings** > **Quality**.

**Note**: You can access **Quality** settings from the **Build Profiles** window (menu: **File** > **Build Profiles**). With [build profiles](build-profiles.md), you can customize the Quality settings per build profile to set different values for each platform. For more information, refer to [Override settings with build profiles](build-profiles-override-settings.md).

A higher quality usually results in a lower frame rate. It’s best to use lower quality on mobile devices and older platforms, to avoid having a detrimental effect on gameplay.

![Edit the settings for a specific Quality level](../uploads/Main/quality-settings-panel.png)

The **Quality** tab contains the following sections:

* **A**: The matrix of quality levels and build platforms in this project.
* **B**: The active build platform.
* **C**: The active quality level.
* **D**: The configuration of the active quality level.

## Quality levels matrix

| **Property** | **Description** |
| --- | --- |
| **Levels** | Lists the quality levels in the project, and the platforms they apply to when Unity builds your project. To make a quality level active in the Editor and the configuration section, select it to highlight it. To apply a quality level to a build platform, enable the checkbox under the platform.  To rename a level, under **Current Active Quality Level**, select **Name**.  To delete a quality level, select the trashcan icon next to the quality level. |
| **Default** | Sets the default quality level for each platform. The checkbox for the default platform displays green. |
| **Add Quality Level** | Adds a quality level by duplicating the highlighted quality level. |

## Configuration section

The configuration section contains the following sections:

* [Rendering](#Rendering)
* [Textures](#Textures)
* [Particles](#Particles)
* [Terrain](#Terrain)
* [Shadows](#Shadows)
* [Async Asset Upload](#AsyncAssetUpload)
* [Level of Detail](#LevelOfDetail)
* [Meshes](#Meshes)

### Rendering

| **Property** | **Description** |
| --- | --- |
| ****Render Pipeline**** | Sets the Render Pipeline Asset to use at this quality level. |
| ****Pixel** Light Count** | Sets the maximum number of per-pixel lights if you use a forward **rendering path**. For more information, refer to [Per-pixel and per-vertex lights](PerPixelLights.md). |
| **Anti Aliasing** | Smooths edges using multisample anti-aliasing (MSAA). The higher the **antialiasing** level, the smoother the appearance of the edges of polygons, but the more processing time needed on the GPU. MSAA is supported only in [forward rendering paths](rendering-paths-introduction.md#forward). For the dropdown options, refer to [Anti Aliasing dropdown](#anti-aliasing-dropdown). |
| **Realtime Reflection Probes** | Indicates whether to update [reflection probes](ReflectionProbes.md) at runtime. |
| **Resolution Scaling Fixed DPI Factor** | Increases or decreases the screen resolution of the device. For more information, refer to [Android Player settings](class-PlayerSettingsAndroid.md#Scaling) and [iOS Player settings](class-PlayerSettingsiOS.md#Scaling). |
| ****VSync** Count** | Synchronizes rendering with the refresh rate of the display device to avoid [tearing artifacts](CameraTroubleshooting.md#tearing). For the dropdown options, refer to [VSync Count dropdown](#vsync-count-dropdown). |
| **Realtime GI CPU Usage** | Controls how much CPU capacity Unity can use to calculate [Enlighten Realtime Global Illumination](realtime-gi-using-enlighten-landing.md). Higher settings make the system react faster to changes in lighting. For more information, refer to [Realtime GI CPU Usage dropdown](#realtime-gi-cpu-usage-dropdown)  This property is available in the **Quality** window only if your project uses the Universal Render Pipeline (URP) or the High Definition Render Pipeline (HDRP). For the Built-In Render Pipeline, refer to [graphics settings](class-GraphicsSettings.md). |

#### Anti Aliasing dropdown

| **Value** | **Description** |
| --- | --- |
| **Disabled** | Disables MSAA. |
| **2x Multi Sampling** | Samples twice per pixel. |
| **4x Multi Sampling** | Samples four times per pixel. |
| **8x Multi Sampling** | Samples eight times per pixel. |

#### VSync Count dropdown

| **Value** | **Description** |
| --- | --- |
| **Don’t Sync** | Disables synchronizing rendering with the display device. |
| **Every V Blank** | Synchronizes rendering so that Unity switches frames every time the display isn’t updating. |
| **Every Second V Blank** | Synchronizes rendering so that Unity switches frames every second time the display isn’t updating. |

#### Realtime GI CPU Usage dropdown

| **Value** | **Description** |
| --- | --- |
| **Low** | Limits **Enlighten** Realtime **Global Illumination** to a low amount of CPU capacity. |
| **Medium** | Limits Enlighten Realtime Global Illumination to a medium amount of CPU capacity. |
| **High** | Limits Enlighten Realtime Global Illumination to a high amount of CPU capacity. |
| **Unlimited** | Uses no CPU capacity limit for Enlighten Realtime Global Illumination. **Note:** Some platforms have their own limit on how many CPU threads or cores Unity can use. For example, Android might limit Unity to one less than the total number of CPUs, or to only little CPUs on devices with big/little CPUs. For more information, refer to [Android thread configuration](android-thread-configuration.md) |

### Textures

| **Property** | **Description** |
| --- | --- |
| **Global Mipmap Limit** | Limits the mipmap resolution level that Unity uses when it renders textures. To use lower resolutions that require less GPU memory and processing time, set this property to a higher mipmap level. For the dropdown options, refer to [Global Mipmap Limit dropdown](#global-mipmap-limit-dropdown).  This property only affects textures with a [texture shape](class-TextureImporter.md#textureshape) of 2D or 2D Array. |
| **Mipmap Limit Groups** | Lists the custom groups that you can add textures to so they override the **Global Mipmap Limit**. Use Mipmap Limit Groups to allocate more memory for important textures, and less memory for less important textures. For more information, refer to [Mipmap Limit Groups properties](#mipmap-limit-groups-properties)  This property affects only textures with a [texture shape](class-TextureImporter.md#textureshape) of 2D or 2D Array. |
| **Anisotropic Textures** | Controls which textures use anisotropic filtering, which improves the visual quality when you view a texture at a steep angle. Anisotropic filtering increases rendering time. For dropdown options, refer to [Anisotropic Textures dropdown](#anisotropic-textures-dropdown). |
| **Mipmap Streaming** | Limits the size of textures in GPU memory by using [mipmap streaming](TextureStreaming.md). Disable this setting to reduce processing time. For dropdown options, refer to [Mipmap Streaming dropdown](#mipmap-streaming-dropdown). |

#### Global Mipmap Limit dropdown

For more information about mipmaps and mipmap limits, refer to [Mipmaps](texture-mipmaps-introduction.md).

| **Value** | **Description** |
| --- | --- |
| **0: Full Resolution** | Uses no limit for texture resolution. Unity can use mipmap level 0, which is the full resolution of the texture. |
| **1: Half Resolution** | Limits the highest resolution to mipmap level 1, which is half resolution. |
| **2: Quarter Resolution** | Limits the highest resolution to mipmap level 2, which is quarter resolution. |
| **3: Eighth Resolution** | Limits the highest resolution to mipmap level 3, which is eighth resolution. |

#### Mipmap Limit Groups properties

To create a new group, select the **Add** (**+**) button. To delete a group, select the **Remove** (**−**) button. If you delete a group, a dialog appears that checks if you want to remove the textures from the group and reimport them. You can’t undo this.

The dropdown for each group contains the following options.

| **Group property** | **Description** |
| --- | --- |
| **Offset Global Mipmap Limit: –3** | Allows the textures in this group to use eight times the resolution of **Global Mipmap Limit**. This results in Unity using higher resolution textures and more memory. For example, if **Global Mipmap Limit** is **3: Eighth Resolution**, Unity limits this group of textures to **0: Full Resolution**. |
| **Offset Global Mipmap Limit: –2** | Allows the textures in this group to use four times the resolution of **Global Mipmap Limit**. |
| **Offset Global Mipmap Limit: –1** | Allows the textures in this group to use two times the resolution of **Global Mipmap Limit**. |
| **Use Global Mipmap Limit** | Uses **Global Mipmap Limit** for the textures in this group. |
| **Offset Global Mipmap Limit: +1** | Limits the textures in this group to half the resolution of **Global Mipmap Limit**. This results in Unity using lower resolution textures and less memory. For example, if **Global Mipmap Limit** is **1: Half Resolution**, Unity limits this group of textures to **2: Quarter Resolution**. |
| **Offset Global Mipmap Limit: +2** | Limits the textures in this group to a quarter of the resolution of **Global Mipmap Limit**. |
| **Offset Global Mipmap Limit: +3** | Limits the textures in this group to an eighth of the resolution of **Global Mipmap Limit**. |
| **Override Global Mipmap Limit: Full Resolution** | Limits the textures in this group to mipmap level 0, which is full resolution. |
| **Override Global Mipmap Limit: Half Resolution** | Limits the textures in this group to mipmap level 1, which is half resolution. |
| **Override Global Mipmap Limit: Quarter Resolution** | Limits the textures in this group to mipmap level 2, which is quarter resolution. |
| **Override Global Mipmap Limit: Eighth Resolution** | Limits the textures in this group to mipmap level 3, which is eighth resolution. |

Open the **More** (**⋮**) menu for additional properties.

| **Additional property** | **Description** |
| --- | --- |
| **Identify textures** | Selects all the textures that belong to the group in the **Project** window. For more information about adding textures to a Mipmap Limit Group, refer to [Texture Import Settings window reference](class-TextureImporter.md). |
| **Duplicate Group** | Duplicate the group. |
| **Rename Group** | Rename the group. If you rename a group, a dialog appears that checks if you want to reassign the textures from the old group to the new group, and reimport them. You can’t undo this. |

#### Anisotropic Textures dropdown

For more detail about anisotropic texture filtering levels, refer to the [`Texture-anisoLevel`](../ScriptReference/Texture-anisoLevel.md) API.

| **Value** | **Description** |
| --- | --- |
| **Disabled** | Disables anisotropic filtering. |
| **Per Texture** | Uses the ****Aniso Level**** each texture is set to in its [texture import settings](class-TextureImporter.md). |
| **Forced On** | Uses anisotropic filtering for all textures. |

### Mipmap Streaming dropdown

| **Property** | **Description** |
| --- | --- |
| **Add All Cameras** | Indicates whether to use mipmap streaming for all active cameras in the project. If you disable this setting, Unity calculates mipmap streaming only for cameras that have a **Streaming Controller** component.  For more information, refer to [Configure mipmap streaming](TextureStreaming-configure.md). |
| **Memory Budget** | Limits the total amount of memory for loaded textures, in MB. The default is 512 MB. For more information, refer to [Set the memory budget for textures](TextureStreaming-configure.md#memory-budget). |
| **Renderers Per Frame** | Limits how many renderers mipmap streaming processes per frame. A lower number decreases the processing time on the CPU, but increases texture loading times. The default is 512. |
| **Max Level Reduction** | Sets the number of mipmap levels that the mipmap streaming system can discard if it reaches the **Memory Budget**. The default is 2.  This value is also the mipmap level that the mipmap streaming system initially loads at startup. For example, when **Max Level Reduction** is set to 2, Unity skips mipmap level 0 and 1 on first load.  For more information, refer to [Set the memory budget for textures](TextureStreaming-configure.md#memory-budget). |
| **Max IO Requests** | Limits the maximum number of file requests the mipmap streaming system makes at any one time. Lower values avoid the system trying to load too many textures if the **scene** changes quickly, but reduce how quickly the system reacts to texture changes. The default is 1024, which is high enough to prevent any limit outside the [Async Upload pipeline](https://unity.com/blog/engine-platform/understanding-the-async-upload-pipeline) and the file system itself. |

### Particles

| **Property** | **Description** |
| --- | --- |
| **Soft Particles** | Indicates whether to fade particles as they approach the edges of opaque **GameObjects**. For more information, refer to [Soft particles](particle-color.md#soft-particles). This property is only available if your project uses the Built-In Render Pipeline. For the Universal Render Pipeline, refer to [Universal Render Pipeline asset reference for URP](urp/universalrp-asset.md). |
| **Particle Raycast Budget** | Sets the maximum number of raycasts to use for **particle system** collisions if ****Collision** Quality** is set to **Medium** or **Low**. For more information, refer to [Particle collisions](particle-collisions.md). |

### Terrain

| **Property** | **Description** |
| --- | --- |
| ****Billboards** Face **Camera** Position** | Enable this option to force billboards to face the camera while rendering instead of the camera plane. This produces a better, more realistic image, but is more expensive to render. |
| **Use Legacy Details Distribution** | Enable this option to use the previously supported scattering algorithm that often resulted in overlapping details. Included for backward compatibility with Terrains authored in Unity 2022.1 and earlier. |
| **Terrain Setting Overrides** | Various override settings that, when enabled, override the value of all active terrains (except those with the “Ignore Quality Settings” setting enabled). For more information about these settings, see [Terrain Settings](terrain-OtherSettings.md). |
| Pixel Error | Value set to Terrain Pixel Error. See [Terrain Settings](terrain-OtherSettings.md). |
| Base Map Dist. | Value set to Terrain Basemap Distance. See [Terrain Settings](terrain-OtherSettings.md). |
| Detail Density Scale | Value set to Terrain Density Scale. See [Terrain Settings](terrain-OtherSettings.md). |
| Detail Distance | Value set to Terrain Detail Distance. See [Terrain Settings](terrain-OtherSettings.md). |
| Tree Distance | Value set to Terrain Tree Distance. See [Terrain Settings](terrain-OtherSettings.md). |
| Billboard Start | Value set to Terrain Billboard Start. See [Terrain Settings](terrain-OtherSettings.md). |
| Fade Length | Value set to Terrain Fade Length. See [Terrain Settings](terrain-OtherSettings.md). |
| Max **Mesh** Trees | Value set to Terrain Max Mesh Trees. See [Terrain Settings](terrain-OtherSettings.md). |

### Shadows

| **Property** | **Description** |
| --- | --- |
| ****Shadowmask** Mode** | Controls when Unity uses baked shadows and real-time shadows if you set **Lighting Mode** to **Shadowmask**. For the dropdown options, refer to [Shadowmask mode dropdown](#shadowmask-mode-dropdown). |

The following properties are only available if your project uses the Built-In Render Pipeline. For the Universal Render Pipeline, refer to [Universal Render Pipeline asset reference for URP](urp/universalrp-asset.md).

| **Property** | **Description** |
| --- | --- |
| **Shadows** | Identifies whether to render soft or hard shadows, or no shadows. |
| **Shadow Resolution** | Controls the visual fidelity of shadows. The higher the resolution, the greater the processing overhead, and the memory used on the GPU. |
| **Shadow Projection** | Controls the quality and stability of shadows from the directional light. For the dropdown settings, refer to [Shadows Projection dropdown](#shadows-projection-dropdown). |
| **Shadow Distance** | Sets the distance away from the camera where Unity no longer renders shadows, in meters. |
| **Shadow Near Plane Offset** | Sets how far to pull back the near clip plane of a shadow map. Use this setting to fix distorted shadows cast by large triangles. For more information, refer to [Shadow pancaking](ShadowPerformance.md#shadow-pancaking). |
| **Shadow Cascades** | Choose the number of shadow cascades to use. A higher number of shadow cascades results in higher-quality shadows but a longer processing time. For more information, refer to [Shadow cascades](shadow-cascades-landing.md). |
| **Cascade splits** | Controls the distance where each shadow cascade starts and ends. To adjust the distances, select and drag the vertical lines between each pair of cascades. This property is only available if you set **Shadow Cascades** to **Two Cascades** or **Four Cascades**. For more information, refer to [Shadow cascades](shadow-cascades-landing.md). |

#### Shadowmask mode dropdown

For more information, refer to [Lighting Mode](lighting-mode.md).

| **Value** | **Description** |
| --- | --- |
| **Shadowmask** | Casts baked shadows from static GameObjects at all distances. |
| ****Distance Shadowmask**** | Casts baked shadows only from static GameObjects beyond the distance set in **Shadow Distance**. Other GameObjects cast real-time shadows. |

#### Shadows Projection dropdown

| **Value** | **Description** |
| --- | --- |
| **Close Fit** | Renders higher resolution shadows that might sometimes wobble slightly if the camera moves. |
| **Stable Fit** | Renders lower resolution shadows that don’t wobble. |

### Async Asset Upload

For more information about the asynchronous upload pipeline, refer to [Texture and mesh loading](LoadingTextureandMeshData.md).

| **Property** | **Description** |
| --- | --- |
| **Time Slice** | Sets the amount of CPU time in ms per frame Unity spends uploading buffered texture and mesh data to the GPU. |
| **Buffer Size** | Sets the size in MB of the asynchronous upload buffer Unity uses to stream texture and mesh data to the GPU. |
| **Persistent Buffer** | Indicates whether the upload buffer persists when there’s nothing left to upload. |

### Level of Detail

For more information, refer to [Optimize mesh rendering using level of detail (LOD)](lod-landing.md).

| **Property** | **Description** |
| --- | --- |
| **LOD Bias** | Adjusts the detail level of GameObjects by scaling the distances where Unity transitions between different level of detail (LOD) meshes. A **LOD Bias** value between 0 and 1 results in Unity selecting lower-quality LODs at closer distances than normal. A value of 1 or more results in Unity selecting lower-quality LODs at farther distances than normal, so GameObjects retain higher quality for longer. For example, if you set **LOD Bias** to 2, a transition to a lower-quality LOD level that usually happens at 50% distance now happens at 25% distance (50% / 2 = 25%). |
| **Maximum LOD Level** | Sets the lowest LOD level the project uses. Unity removes LOD meshes below the **Maximum LOD level** from the build, which makes the build smaller, and reduces memory use at runtime. If you have different quality levels that use different **Maximum LOD Level** values, Unity uses the smallest value. For example, if any quality level uses a **Maximum LOD Level** of 0, Unity includes all LOD levels in the build. **Note:** If a model is included in [Addressables](http://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html) group, Unity adds the entire model including all LOD meshes in the build, regardless of the **Maximum LOD level** property. |
| **Mesh LOD Threshold** | Affects how Unity selects a LOD index to render. Increasing the value makes Unity favor less detailed LODs in the evaluation process. For more information, refer to [Mesh LOD runtime quality](lod/mesh-lod-quality.md#project-wide-quality-setting). |
| **LOD Cross Fade** | Smooths the transition between LOD meshes by fading between two LOD levels using a dithering pattern. This property is only available if your project uses the Built-In Render Pipeline. For the Universal Render Pipeline, refer to [Universal Render Pipeline asset reference for URP](urp/universalrp-asset.md). |

### Meshes

| **Property** | **Description** |
| --- | --- |
| **Skin Weights** | Sets the number of bones that can affect a vertex during an animation. For more information, refer to [Skinned Mesh Renderer component reference](class-SkinnedMeshRenderer.md). |

## Additional resources

* [Create and manage build profiles](create-build-profile.md)