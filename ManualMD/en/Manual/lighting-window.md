# Lighting window reference

The Lighting window (menu: **Window** > **Rendering** > **Lighting**) is the main control point for Unity’s lighting features. You can use the Lighting window to adjust settings related to the lighting in your Scene, and to optimise your precomputed lighting data for quality, bake time, and storage space.

## Related APIs

You can perform many of the functions available in the Lighting window in **scripts**, using the [LightingSettings](../ScriptReference/LightingSettings.md) and [Lightmapping](../ScriptReference/Lightmapping.md) APIs.

## Lighting window layout

The Lighting window contains the following elements:

* The [Scene tab](#Scene)
* The [Adaptive Probe Volumes tab](#ProbeVolumes) - Universal **Render Pipeline** (URP) and High-Definition Render Pipeline (HDRP) only
* The [Environment tab](#Environment)
* The [Realtime Lightmaps tab](#RealtimeLightmaps)
* The [Baked Lightmaps tab](#BakedLightmaps)
* The [control area](#ControlArea), at the bottom of the window

## The Scene tab

The **Scene** tab displays information about the [Lighting Settings Asset](class-LightingSettings.md) that is assigned to the active Scene. If no Lighting Settings Asset is assigned to the active Scene, it displays information about the [default LightingSettings object](class-LightingSettings.md).

The Scene tab is divided into the following sections:

* [Lighting Settings Asset controls](#AssetControls)
* [Lighting settings](#LightingSettingsProperties) - Realtime Lighting, Mixed Lighting and Lightmapping settings

### Lighting Settings Asset controls

Use the controls in this section to assign a Lighting Settings Asset to the active Scene, or to create a new Lighting Settings Asset.

| **Property:** | **Function:** |
| --- | --- |
| **Lighting Settings** | The Lighting Settings Asset assigned to the active Scene. |
| **New Lighting Settings** | Click this button to generate a new Lighting Settings Asset in your Project, and automatically assign the new Lighting Settings Asset to the active Scene. |

### Lighting Settings

Use this section to view and edit the properties of the Lighting Settings Asset or `LightingSettings` object assigned to the current Scene. See [Lighting Settings Asset](class-LightingSettings.md).

## The Adaptive Probe Volumes tab - URP and HDRP only

The Adaptive Probe Volumes tab contains settings related to Adaptive Probe Volumes (APV). This tab only appears if you use URP or HDRP in your project. Refer to the following for more information:

* [Adaptive Probe Volumes in URP](urp/probevolumes.md)
* [Adaptive Probe Volumes in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/probevolumes.html)

## The Environment tab

The Environment tab contains settings related to environmental lighting effects for the current Scene. The contents depend on the render pipeline that your Project uses.

The Environment tab is divided into two sections:

* [Environment](#EnvironmentSection)
* [Other settings](#OtherSection)

### Environment

The Environment section contains lighting-related settings and controls that apply to the environmental lighting in the current scene, such as the Skybox, diffuse lighting and reflections.

| **Property** | **Description** |
| --- | --- |
| **Skybox Material** | A [Skybox](sky-landing.md) is a Material that appears behind everything else in the Scene to simulate the sky or other distant background. Use this property to choose the Skybox you want to use for the Scene. The default value is the built-in Default Skybox. |
| **Sun Source** | Select a Light to use as the sun in your Scene. Unity uses this Light to simulate the effect of sun position and intensity on the Skybox and your Scene. If you set this to **None**, Unity considers the brightest directional light in the Scene the sun. Lights whose **Render Mode** property is set to **Not Important** do not affect the Skybox.   For more information about the **Render Mode** setting, see the Additional settings section of [Lights](https://docs.unity3d.com/Manual/class-Light.html). |
| **Realtime Shadow Color** | Define the color that Unity uses to render real-time shadows in Subtractive Lighting Mode. |
| **Environment Lighting** | This section contains settings that affect [ambient light](lighting-ambient-light.md) in the current Scene. |
| **Source** | Use this to define a source color for ambient light in the Scene. The default value is **Skybox**. The options are:  * **Skybox**: Use the colors of the Skybox set in **Skybox Material** to determine the ambient light coming from different angles. This allows for more precise effects than **Gradient**. * **Gradient**: Choose separate colors for ambient light from the sky, horizon and ground, and blend smoothly between them. * **Color**: Use a flat color for all ambient light. |
| **Intensity Multiplier** | Use this to set the brightness of the ambient light in the Scene, defined as a value between 0 and 8. The default value is 1. |
| **Environment Reflections** | This section contains global settings for [Reflection Probe](class-ReflectionProbe.md) baking, and settings that affect global reflections. |
| **Source** | Use this setting to specify whether you want to use the Skybox for reflection effects, or a Cubemap of your choice. The default value is **Skybox**. The options are:  * **Skybox**: Select this to use the Skybox as the reflection source. * **Custom**: Select this to use either a Cubemap asset or a [RenderTexture of type cube](class-RenderTexture.md) for reflections. |
| **Resolution** | Use this to set the resolution of the Skybox for reflection purposes. This property is visible only when **Source** is set to **Skybox**. |
| **Cubemap** | Use this to specify the Cubemap to use for reflection purposes. This property is visible only when **Source** is set to **Cubemap**. |
| **Compression** | Use this to define whether or not reflection textures are compressed. The default setting is **Auto**. The options are:  * **Auto**: The reflection texture is compressed if the compression format is suitable. * **Uncompressed**: The reflection texture is stored in memory uncompressed. * **Compressed**: The texture is compressed. |
| **Intensity Multiplier** | The degree to which the reflection source is visible in reflective objects. |
| **Bounces** | A reflection bounce occurs when a reflection from one object is then reflected by another object. Use this property to set how many times the Reflection Probes evaluate bounces back and forth between objects. If this is set to 1, then Unity only takes the initial reflection (from the skybox or cube map specified in the **Reflection Source** property) into account. |

### Other Settings

The Other Settings section contains settings for fog, [Halos](class-Halo.md), [Flares](class-Flare.md) and [Cookies](Cookies.md).

| **Property** | **Description** |
| --- | --- |
| **Fog** | Enables or disables fog in the Scene. Note that fog is not available with the [Deferred rendering path](RenderingPaths.md). For more information, refer to [Fog settings](#fogsettings). |
| **Halo Texture** | Set the Texture you want to use for drawing a [Halo](class-Halo.md) around lights. |
| **Halo Strength** | Define the visibility of Halos around Lights, from a value between 0 and 1. |
| **Flare Fade Speed** | Define the time (in seconds) over which [lens flares](class-LensFlare.md) fade from view after initially appearing. This is set to 3 by default. |
| **Flare Strength** | Define the visibility of lens flares from lights, from a value between 0 and 1. |
| **Spot Cookie** | Set the [Cookie](Cookies.md) texture you want to use for [Spot Lights](Lighting.md). The default is ‘Soft’. To revert to ‘Soft’, select **None**. |

#### Fog settings

| **Property** | **Description** |
| --- | --- |
| **Color** | Use the color picker to set the color Unity uses to draw fog in the Scene. |
| **Fog Mode** | Define the way in which the fogging accumulates with distance from the camera. For more information, refer to [Fog Mode dropdown](#fogmodedropdown). The options are:  * **Linear**: Fog density increases linearly with distance. **Start** sets the distance from the Camera at which the fog starts. **End** sets the distance from the Camera at which the fog completely obscures Scene GameObjects. * **Exponential**: Fog density increases exponentially with distance. **Density** controls the density of the fog. The fog appears more dense as the **Density** increases. * **Exponential Squared**: Fog density increases faster with distance (exponentially and squared). **Density** controls the density of the fog. The fog appears more dense as the **Density** increases. |

## The Realtime Lightmaps tab

The Realtime **Lightmaps** tab shows a list of all lightmaps generated by the **Enlighten** Realtime **Global Illumination** system in the current Scene. If Enlighten Realtime Global Illumination is not enabled in your Project, this tab will be empty.

## The Baked Lightmaps tab

This tab displays a list of all the lightmaps generated by the [lightmapper](Lightmappers.md) for the current scene, and the [Lighting Data Asset](LightmapSnapshot.md).

If you [use Scene view Draw Modes to preview lightmapping](Lightmapping-preview.md), the tab also contains the temporary lightmaps Unity generates for the preview.

If Baked Global Illumination is not enabled in your Project, the tab is empty.

## Control area

Controls for precomputing lighting data are at the bottom of the Lighting window.

| **Property:** | **Description** |
| --- | --- |
| **GPU Baking Device** | Use this to change the GPU that Unity uses for precomputing lighting data. This property is visible only when you use the [GPU Progressive Lightmapper](progressive-lightmapper.md) backend. |
| **GPU Baking Profile** | The profile you select in this property defines how the GPU Lightmapper breaks lightmaps into smaller tiles to reduce GPU memory usage. The options are:  * **Automatic**: Unity chooses the tile size based on the size of the largest lightmap while still aiming for good GPU utilization. * **Highest Performance** and **High Performance**: Force a higher fixed tile size for all lightmaps. * **Low Memory Usage** and **Lowest Memory Usage**: Force a lower fixed tile size for all lightmaps. Small tiles take up less GPU memory at the expense of lower GPU utilization, leading to longer bake times.  This property is visible only when you use the [GPU Progressive Lightmapper](progressive-lightmapper.md). |
| **Bake on Scene Load** | Set whether Unity automatically generates precomputed lighting data. The options are:  * **Never**: Prevent Unity from automatically generating precomputed lighting data when you open a scene. To generate lighting data manually, select **Generate Lighting**. This is the default mode in a new Unity 6 project. * **If Missing Lighting Data**: Enable Unity automatically generating precomputed lighting data when you open a scene, if the data doesn’t exist or is invalid. This is the default mode for a project that was created in Unity 2023.2 or earlier. Refer to [Upgrade to Unity 6.0](UpgradeGuideUnity6.md) for more information. |
| **Generate Lighting** | Select the **Generate Lighting** button to precompute lighting data. This data includes lightmaps for the Baked Global Illumination system, lightmaps for the Enlighten Realtime Global Illumination system, Light Probes, Adaptive Probe Volumes, and Reflection Probes. Edits you make after starting the bake process using the **Generate Lighting** button do not affect baked lighting.  **Note:** On MacOS Arm64 (Apple Silicon), make sure Rosetta 2 is installed when using the Enlighten Realtime Global Illumination system.  The dropdown options are:  * **Bake Probe Volumes**: Bake only the Adaptive Probe Volumes in the scene or Baking Set. * **Bake Reflection Probes**: Bake only the Reflection Probes for all open Scenes. **Note:** If you select **Bake Reflection Probes** but you haven’t baked the scene before, Unity also bakes lightmaps, Light Probes, and Adaptive Probe Volumes. * **Clear Baked Data**: Clear all precomputed lighting data from all open scenes without clearing the [GI Cache](GICache.md). |

---

* [Lighting Settings Asset] added in [2020.1](https://docs.unity3d.com/2020.1/Documentation/Manual/30_search.html?q=newin20201) NewIn20201