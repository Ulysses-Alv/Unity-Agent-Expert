# New in Unity 6.1

Unity 6.1 (6000.1) introduces several new features and improvements.

To read about all of the changes and known issues in this version of Unity, refer to the [Unity 6.1 Release Notes](https://unity.com/releases/editor/whats-new/6000.1.0#notes).

To find the release notes for other releases, refer to the [Unity download archive](https://unity.com/releases/editor/archive).

To read about changes made to packages that were updated as part of this release, refer to the [documentation of those packages](upm-docs.md).

If you’re upgrading existing projects from Unity 6.0 (6000.0), refer to the [Upgrade Guide to Unity 6.1](UpgradeGuideUnity61.md) for information about how your project might be affected.

Review the new features and improvements for Unity 6.1 in these areas:

* [2D](#2d)
* [Bug Reporter](#bug-reporter)
* [Editor and Workflow](#editor-and-workflow)
* [Graphics](#graphics)
* [Raytracing API](#raytracing-api)
* [Platforms](#platforms)
* [Project Auditor](#project-auditor)
* [SRP Core](#srp-core)
* [UI Toolkit](#ui-toolkit)
* [XR](#xr)

## 2D

Enabled auto-generation of **Tilemap** assets in PSD Importer and Aseprite Importer, so that you can start authoring 2D worlds more quickly.

The new Tile Set asset provides generative capabilities for existing **Sprite** Texture assets, where you only need to set up the asset initially. Changes to the source propagate down automatically.

To learn more, refer to [Tile Set documentation](tilemaps/tile-palettes/tile-set-landing.md).

## Bug Reporter

Updated the Bug Reporter for Unity 6.1 to be more stable, have a smoother UI, and a new dark mode. The Bug Reporter uses a new framework to make it more robust and future proof in native systems. The Bug Reporter now has an Apple Silicon native client, meaning we have removed our reliance on Rosetta for support.

The Bug Reporter now also links to the Hub. If you are logged in through the Hub, all of your credentials are forwarded automatically, so it is easier than ever to submit any issues you come across in the Unity Editor.

## Editor and workflow

This section outlines recent updates to the Editor and its general workflows.

### Choose a save location for scripts created from the Inspector

Creating a new script from the **Add Component** menu now opens a dialog where you can choose a save location for the new script.

### Updates to the main menu

Reorganized the submenus in the following main menu items to be more navigable:

* **Edit**
* **Assets**
* **Component**
* **Window**
* **Help**

In addition, made the following changes in these menus:

* **Edit**: Moved **Undo History** from **Edit** to **Window** > **General**.
* **Assets**: Renamed **Extract from **Prefab**** to **Extract Material Subasset**. This menu item is disabled if you do not have a material asset selected in the **Project** window.
* **Component**: Moved most items from the **Miscellaneous** submenu to more appropriate submenus that better indicate their type. This move affects the **Add Component** **Inspector** interface as well.
* **Window**:
  + Added a new **Package Management** submenu that contains the Package Manager, **Asset Store**, Services, and My Assets windows.
  + Moved the ****Version Control**** menu item to a **Version Control** submenu and renamed it to **External Provider**. This might affect user-defined keyboard shortcuts or **scripts** that rely on menu item paths.
* **Help**: On macOS, **Check For Updates** is now in the **Unity Menu**.

## Graphics

This section outlines recent updates to Unity’s graphics systems.

### Reduce shader variants in URP

If your project uses the Universal **Render Pipeline** (URP), you can now reduce the number of [shader variants](shader-variants.md) Unity compiles for fog and **level of detail** meshes (LODs). This helps reduce build times, file size, **shader** loading times, and runtime memory usage.

For more information, refer to the following:

* [Strip fog shader variants in URP](urp/shader-stripping-fog.md)
* [Strip level of detail (LOD) shader variants in URP](urp/shader-stripping-lod)

### Reduce shader stuttering with pipeline state object (PSO) tracing

To reduce stuttering caused by shader compilation, you can now use the experimental `GraphicsStateCollection` API. Use the API to trace the pipeline state objects (PSOs) the graphics API needs and compile them ahead-of-time. For more information, refer to [Prewarm shaders](shader-prewarm.md).

### Bicubic lightmap sampling in URP and HDRP

Added a **Use Bicubic **Lightmap** Sampling** setting to [URP](urp/urp-global-settings.md) and [HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/Default-Settings-Window.html) graphics settings to improve the visual fidelity of lightmaps by smoothing sharp or jagged edges, especially at the edges of shadows. For more information, refer to [Improve the visual fidelity of lightmaps](urp/lighting/lightmapping-improve-visual-fidelity.md).

### DirectX 12 as the default graphics API

DirectX 12 is now the default graphics API for new projects targeting the Windows platform. You can still use DirectX 11 for your project if needed.

### Read/write texture support in custom HLSL

Added support for read/write graphics buffers and textures as input parameters to the Custom HLSL Block and Operator. This allows VFX to write into custom buffers without relying on global bindings. For more information, refer to [Custom HLSL Nodes (block and operator)](https://docs.unity3d.com/Packages/com.unity.visualeffectgraph@17.0/manual/CustomHLSL-Common.html).

### Runtime and compilation optimizations

Added VFX performance optimizations for both the Editor and runtime:

* Optimized **particle** data layout for reduced bandwidth usage and smaller memory footprint.
* Improved parallelization in the VFX runtime for faster CPU execution and less main thread computation.
* Reduced garbage creation during shader source generation, resulting in faster import and authoring times.

### Instancing support for GPU events

Added instancing support for GPU events, so VFX using GPU events no longer disables instancing by force. This simulates and renders many instances of VFX with GPU events more efficiently as batches.

### Deferred+ rendering path

Added Deferred+ as a supported **rendering path** in URP. The Deferred+ rendering path, like the Deferred rendering path, has no limit on the number of lights that can affect an opaque **GameObject**, but uses [Forward+](urp/rendering/forward-rendering-paths.md) instead of Forward for the transparent and forward only opaque pass. Refer to [Deferred+ rendering paths in URP](urp/rendering/deferred-rendering-path-landing.md) for more information.

### HDRP

This section outlines recent updates to the High Definition Render Pipeline (HDRP).

#### Enhanced performance with Variable Rate Shading

Added Variable Rate Shading (VRS) support for custom passes to optimize shading performance. You can now dynamically adjust the shading resolution in specific areas of the screen. This helps reduce GPU workload while maintaining visual quality. For more information, refer to the [ShadingRateImage](../ScriptReference/Rendering.ShadingRateImage.md) and the [ShadingRateInfo](../ScriptReference/Rendering.ShadingRateInfo.md) API documentation.

#### 3D deformation for Water System

Added support for horizontal deformations in the Water Decal system’s deformation output. Previously, you could only modify the vertical position of Water Surface vertices. With this feature, you can create more complex effects like rolling waves. For more information, refer to the Rolling Wave **scene** in the HDRP Water Samples package and [Deform a water surface vertically](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.1/manual/water-deform-a-water-surface.html).

![A rolling wave using the 3D deformation water system.](../uploads/Main/61-hdrp-deformed-water.png)

A rolling wave using the 3D deformation water system.

#### New water sample scene with caustics and deformation effects

Added a new water sample cave scene that showcases a Custom Pass and a Local Volumetric Fog sampling the Water System caustics buffer to demonstrate advanced water rendering techniques. For more information, refer to [Caustics in the water system](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.1/manual/water-caustics-in-the-water-system.html).

![A HDRP sample scene of an icy cave partially submerged in water, with purple light rays shining through a gap and a glowing treasure chest on a ledge.](../uploads/Main/61-hdrp-water-sample.png)

A HDRP sample scene of an icy cave partially submerged in water, with purple light rays shining through a gap and a glowing treasure chest on a ledge.

## Raytracing API

This section outlines recent updates to Unity’s raytracing API.

### Manual control over ray-tracing acceleration updates

Added the `RayTracingAccelerationStructure.UpdateInstanceGeometry` API. You can now manually trigger updates to the Bottom-Level Acceleration Structure (BLAS). This gives you greater control over when and how geometry updates occur in ray-traced scenes. As a result, you can manage resources more efficiently and improve performance in dynamic environments. For more information, refer to the [API documentation](../ScriptReference/Rendering.RayTracingAccelerationStructure.UpdateInstanceGeometry.md).

### Efficient instance management

Added the new `RayTracingAccelerationStructure.AddInstancesIndirect` API. You can now specify ray-tracing instance transformation matrices using a GraphicsBuffer. As a result, you can manage large numbers of ray-traced instances more efficiently, particularly in dynamic or procedurally generated scenes. For more information, refer to the [API documentation](../ScriptReference/Rendering.RayTracingAccelerationStructure.AddInstancesIndirect.md).

### Flexible ray-tracing instance removal

Added the new `RayTracingAccelerationStructure.RemoveInstance` API. You can now remove ray-tracing instances from an acceleration structure based on their layer or ray-tracing mode. As a result, you can manage dynamic scenes, especially those requiring frequent instance updates, more flexibly and efficiently. For more information, refer to the [API documentation](../ScriptReference/Rendering.RayTracingAccelerationStructure.RemoveInstance.md).

### More efficient ray-tracing instance management

Added the `RayTracingAccelerationStructure.CullInstances` API. Using various filtering and culling parameters, you can now populate the acceleration structure with ray-tracing instances associated with renderers in the scene. As a result, you can more easily update the acceleration structure based on scene visibility and rendering criteria. For more information, refer to the [API documentation](../ScriptReference/Rendering.RayTracingAccelerationStructure.CullInstances.md).

## Platforms

This section outlines recent updates to platform-specific tools and settings.

### Android

This section outlines recent updates to Android-specific tools and settings.

#### Large screen and foldable device support

You can now use the following Android **Player settings** and APIs to support your application across various device screen sizes, including large screen and foldable devices:

* [Resizeable Activity](class-PlayerSettingsAndroid.md#Resolution) Player setting
* [Run Without Focus](class-PlayerSettingsAndroid.md#Resolution) Player setting
* [AndroidApplication](../ScriptReference/Android.AndroidApplication.md) and [AndroidConfiguration](../ScriptReference/Android.AndroidConfiguration.md) APIs

These settings and APIs support multi-window capabilities in your application, detect device configuration changes, and provide access to the application runtime data.

For more information, refer to [Large screen and foldable device support](android-large-screen-and-foldable-support.md).

#### Vulkan device filter asset

Added a new Player setting Vulkan Device Filter asset allowing you to specify the preferred graphics jobs mode at runtime for specific Android devices. You can use this asset to configure the use of Vulkan API on specific Android devices. The asset uses filter lists to identify the Android devices for Vulkan API and graphics jobs mode configuration.

For more information, refer to [Vulkan API and graphics jobs mode configuration](vulkanapi-graphics-jobs-configuration.md).

### Web platform

This section outlines recent updates to Unity’s Web platform.

#### Facebook Instant Games platform

You can now create Web builds tailored for Facebook Instant Games. Key features include:

* Preconfigured settings for mobile and desktop builds.
* Integration with the Facebook Instant Games SDK and Web Stripping Tool package.
* A streamlined publishing workflow.

For more information, refer to [Facebook Instant Games](instant-games.md).

#### Instant publish to Unity Play

You can now publish Web builds directly to [Unity Play](https://play.unity.com/en). Use the new **Publish to Play** option for the Web platform at **File** > **Build Profiles**.

### WebGPU

You can now use the WebGPU graphics API to include modern graphics features in your Web projects. Key features include:

* Compute shaders
* Indirect rendering
* GPU **skinning**
* VFX Graph

**Note**: WebGPU is currently experimental and not yet suitable for production usage.  
For more information, refer to [WebGPU (Experimental)](WebGPU.md).

### Build Profiles

This section outlines recent updates to Unity’s build profiles.

#### Override graphics and quality settings

You can now override graphics and quality settings for a build profile from the **Build Profiles** window. Enable **Override Global Graphics Settings** or **Override Global Quality Settings** to customize settings for a specific build target. These overrides allow you to customize performance and visual quality for different platforms while preserving global settings. For more information, refer to [Override settings with build profiles](build-profiles-override-settings.md).

## Project Auditor

Use the Project Auditor package to optimize your Unity Projects. The Project Auditor can detect problems in your project early in development and give you guidance on how to resolve those problems. This static analysis tool scans your project’s codebase and files and creates a detailed report that gives you insight into potential issues related to assets, scripts, **project settings**, builds, Domain Reload, and more. For more information, refer to [Project Auditor package](https://docs.unity3d.com/Packages/com.unity.project-auditor@1.0/manual/index.html).

## SRP Core

You can now use a render graph debug setting to temporarily disable pass merging to isolate issues or investigate performance changes.

You can use the new variable rate shading (VRS) API to control the shading rate of Scriptable Renderer Features, and balance between GPU performance and image quality. This API is supported on DirectX 12, Vulkan, and compatible consoles. For more information about **VRS**, refer to the Unity Discussions post [Optimize Renderer Features with Variable Rate Shading in Unity 6.1](https://discussions.unity.com/t/optimize-renderer-features-with-variable-rate-shading-in-unity-6-1/1605893/1).

## UI Toolkit

This section outlines recent updates to Unity’s UI Toolkit.

### New Variables field in UI Builder

You can now create USS variables directly in UI Builder.

This feature streamlines the workflow for managing styles. It lets you define and modify USS variables without switching to an external text editor. It improves efficiency, reduces context-switching, and ensures consistency in your styling. For more detail, refer to [Manage USS variables in UI Builder](UIB-styling-ui-using-uss-variables.md).

### New Mask64Field UI control

You can now use the new **Mask64Field** control to create and edit 64-bit bitmask fields, allowing for precise bitwise flag management and streamlined data handling in your UI Toolkit workflows. For more detail, refer to [Mask64Field](UIE-uxml-element-Mask64Field.md).

### Updated UI Toolkit profiler markers

The top-level UI Toolkit **profiler markers** have been renamed to be more accurate and descriptive. The user manual now includes detailed descriptions of these markers and the conditions under which they appear. These improvements help developers identify bottlenecks and optimize UI Toolkit workflows. For more detail, refer to [UI Toolkit profiler marker](UIE-profiler-markers.md).

## XR

This section outlines recent updates to Unity’s extended reality support.

### Meta Quest build profile

Added a Meta Quest build profile to override project settings specific to the Meta Quest build target.

To learn more about developing for Meta Quest, visit [Develop for Meta Quest workflow](xr-meta-quest-develop.md). You can learn more about build profiles in the [Build profiles](BuildSettings.md) documentation.

### Android XR

Unity now supports Google’s Android XR platform with the Unity OpenXR Android XR package. For more information about Unity’s support for Android XR, refer to the [Android XR support is now available](https://discussions.unity.com/t/android-xr-support-is-now-available/1567695) announcement post, and the [Unity OpenXR Android XR](https://docs.unity3d.com/Packages/com.unity.xr.androidxr-openxr@latest) documentation.

#### Android XR build target

Added a new build target for Android XR that you can use to create a [Build profile](BuildSettings.md) to manage your project settings for the Android XR platform.

### Mixed Reality multiplayer tabletop template

Added a new **Mixed Reality** multiplayer tabletop template to the Unity Hub. This template leverages Unity multiplayer services, and configures the necessary project settings to create a starting point for tabletop Mixed Reality multiplayer experiences.

For more information, refer to the [Mixed Reality multiplayer tabletop template](https://docs.unity3d.com/Packages/com.unity.template.mr-multiplayer@1.0/manual/index.html) documentation.

## Additional resources

* [New in Unity 6.0](WhatsNewUnity6.md)
* [Upgrade Unity](UpgradeGuides.md)
* [System requirements](system-requirements.md)
* [Install Unity](GettingStartedInstallingUnity.md)