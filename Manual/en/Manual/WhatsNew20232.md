# New in Unity 2023.2

Find out what’s changed in Unity 2023.2 since 2023.1 and view the documentation for the affected areas.

## Release Notes

To find out more about the new features, changes, and improvements to this Unity version, refer to the [2023.2 Release Notes](https://unity3d.com/unity/whats-new/2023.2.0).

To find the release notes for other releases, refer to the [Unity download archive](https://unity.com/releases/editor/archive).

## Upgrade Guide

If you are upgrading existing projects from 2023.1, read the [Upgrade Guide to 2023.2](UpgradeGuide20232.md) for information about how your project might be affected.

## Accessibility

Added a new Editor window, the Accessibility Hierarchy Viewer, that displays the active accessibility hierarchy and its nodes. To access the Accessibility Hierarchy viewer, go to **Window** > **Accessibility** > **Hierarchy viewer** in the main menu.

## Audio

* Added the [Audio Random Container](Create-randomized-playlist.md) to randomize audio and ensure that volume, pitch, time and triggers can be set to non-repetitive intervals, so your game never sounds the same twice.
* Added a VU meter to the Audio Random Container.

## Authoring Tools

### 2D

Added overlay support to the **Tile Palette** editor window.

### 2D Physics

Added `ArticulationBody.jointPosition` pointer lines to the ****Joint** Angular Limits** **gizmo** to show the exact position of the joint in the **Scene** view.

### Editor and Workflow

* Added the [Scene view context menu](SceneViewContextMenu.md). You can now access new context menus in the **Scene view** by right-clicking or using a customizable keyboard shortcut. These menus, created with the UI Toolkit and extensible in C#, offer easy access to frequently used commands.
* Added the **Refresh the Scene view only when the Editor is in focus** option to the [Scene view preferences](Preferences.md#scene-view). Enable this option to refresh the Scene view only when the Editor is in focus.
* Added the ability to bind a keyboard shortcut to make transitions between Animator states.
* Added basic OpenType font feature support. Currently, only kerning is enabled.
* Added a new [Cameras overlay](cameras-overlay.md) to replace the **Camera** preview. You can use the Cameras overlay to take [first-person control of cameras](control-camera.md) and manage cameras in the Scene view.
* Replaced most OS contextual menus with the UI Toolkit version.
* Added the Color Checker, which is a tool used to calibrate lighting and post process. The Color Checker is an object that the user can add through **GameObject** > **Rendering** > **Color Checker Tool**. The tool is meant only as a production tool for lighting artists and won’t be saved in Build.
* Enabled the Editor to show different license notification modals.
* Added a `PropertyCollectionAttribute`, which you can use to implement custom drawers for collections.
* Added a new tool for light placement using the pan, zoom, and orbit controls of the Camera. For more information, refer to [View and control a light from its perspective](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/lights-placement-tool.html).

### UI Toolkit

* You can now use a new and flexible [runtime binding system](UIE-runtime-binding.md) to effortlessly connect data values to UI element properties when you create the Editor or runtime UI. You can configure data binding in either UI Builder or C# code.
* Added new controls such as the [ToggleButtonGroup](UIE-uxml-element-ToggleButtonGroup.md), [Tab](UIE-uxml-element-Tab.md), and [TabViews](UIE-uxml-element-TabView.md), and made improvements to existing controls.
* Added icon support to the [Button](UIE-uxml-element-Button.md), ListView, and TreeView controls.
* Added new UxmlElement and UxmlAttribute attributes. These attributes replace the existing UxmlFactory and UxmlTraits when you [create custom controls](UIE-create-custom-controls.md). This streamlines the creation of custom controls, offering an efficient alternative through C# attributes, and eliminating the need for extensive code writing. Additionally, you can now create [custom property drawers](ui-systems/custom-control-customize-uxml-attributes.md#add-custom-property-drawer-in-ui-builder) for fields in the same manner as the ****Inspector****.
* Enhanced UI Builder with improvements to display what drives style properties, refined canvas manipulation, and made other enhancements. The updated UI Builder now also supports the authoring of UXML Objects, enabling the editing of the MultiColumnTreeView and MultiColumnListView.
* Added the Emojis Fallback Support field to [TextElement](UIE-uxml-element-TextElement.md) and [TextField](UIE-uxml-element-TextField.md) to control the ordering of where to search for the glyph in the emoji range (primary font vs global fallback). Additionally, added basic support for the OpenType font features, with the current focus on enabling kerning.
* Added a new Spacing field with a box model widget in the UI Builder’s Inspector.

### TextMeshPro

* Added basic Emoji support.
* Added basic OpenType font feature support, only kerning has been enabled for now.

## Graphics

* Batched Compute **Skinning**: Unity’s Skinned **Mesh** Renderer component uses compute dispatches for parallelizing vertex transformations on the GPU. Unity 2023.2 introduces optimizations to Skinned **Mesh Renderers** across all supported platforms. These optimizations aim to batch compute skinning and blendshape dispatches. This can increase the amount of vertices deformed in parallel and improve the GPU performance of character and skinned mesh rendering.
* The Progressive GPU **Lightmapper** is now out of preview and fully supported.
* Added APIs to move **Light Probe** positions at runtime.
* Removed the **Auto Generate** setting in the Lighting window. Related APIs are now obsolete. To check **lightmaps** while you’re editing, you can now select a Scene View Draw Mode and set **Lighting Data** to **Preview**. This displays a preview of the baked lighting. The preview lightmaps are non-destructive, and you can use them after you’ve baked the scene.
* Added `BatchCullingContext.cullingFlags` to specify whether to cull lightmapped shadow casters.
* Added `rendererPriority` support for `BatchRendererGroup`.
* Added support for building **ray tracing** acceleration structures asynchronously on a compute queue. AsyncCompute CommandBuffers can now run `CommandBuffer.BuildRayTracingAccelerationStructure` commands. Added support for `RayTracingAccelerationStructure` to RenderGraph and Render Graph Viewer.
* Added mipmap limit support for Texture2DArrays.
* Added mipmap stripping support for Texture2DArrays.
* Added support for providing tiled EXR images to LoadImage.
* Enabled exposing raytracing acceleration structure build flags for balancing build times versus ray tracing speed and memory consumption on the GPU. You can customize flags from C# when creating and building a `RayTracingAccelerationStructure`, or via the Renderer settings.
* Added Native Render Pass support for Vulkan for Android, Metal for iOS/MacOS (Apple Sillicon), and DirectX12 for Windows on ARM.
* Added support for GPU batched skinning for D3D12 (Windows and XBox platforms).
* Added the Custom Vertex Streams feature for **particle** trails.
* Added `BakeTexture` and `BakeTrailsTexture` scripting methods.

### Universal Render Pipeline (URP)

For a complete description of new features and improvements in URP, refer to [What’s new in URP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@15.0/manual/whats-new/urp-whats-new.html).

* Added cross-platform **HDR** display support. HDR displays are capable of reproducing images in the higher range of difference in luminance closer to natural lighting conditions. HDR output preserves the contrast and quality of the linear lighting renders and HDR images displayed on these devices. The Editor and Standalone Players now provide full HDR tone mapping and display support, across all rendering pipelines and capable platforms, including mobile and **XR**.
* Added per-vertex quality levels for indirect lighting from Probe Volumes so that URP renders lit environments more efficiently. You might need to do further optimization on mobile because URP doesn’t support Lighting Scenario blending or Lighting Normalization for **Reflection Probes**.
* Added a new option in the Motion Blur volume component called **Camera And Objects** that uses motion vectors to blur objects that move faster than the camera’s exposure time.
* Added support for additional directional light cookies.
* Added the Default Volume Profile field to URP Global Settings.
* Added the Volume Profile field to the Universal **Render Pipeline** Asset.
* Added Alembic velocity motion vector support for URP materials.
* Added automatic **TimeBased** motion vector generation for ShaderGraphs with vertex animation based only on the **Time** node. All other data which affects position must be constant between frames.
* Added support for XR rendering and cameras using orthographic projection to Forward+ **rendering path**.
* Added support for foveated rendering in the Forward+ rendering path.

### High Definition Render Pipeline (HDRP)

For a complete description of new features and improvements in URP, refer to [What’s new in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@16.0/manual/whats-new-16.html).

* Added HDRP path tracer support for Decals projectors. The path tracer does not support emission from decals.
* Added the ability to stream Probe Volume data from disk at runtime. You can enable or disable disk streaming for different quality levels.
* Made HDRP **Shader** Graph decals compatible with transparent objects. This means you can use decals created with Shader Graph to affect transparent objects to build procedural effects like rain drops, ripples, custom engravings, dirt effects on glass.
* Added a night sky with stars and celestial bodies to the HDRP physical sky.
* Added a Shader Graph output for the physically based sky and added controls to create a moon.
* Added visualization of async compute passes and their synchronization points to the Render Graph Viewer.
* Added beer shadow maps for volumetric clouds.
* Added a Volume Profile field to HD Render Pipeline Asset.
* Added a material type for thin objects with colored transmission.
* Added the option to disable clear coat on the material for Lit ShaderGraphs.
* Added the Global Pass API that you can use to inject custom passes in the render pipeline without any **GameObjects** in a scene.
* Added the volumetric fog fullscreen debug mode output for Arbitrary Output Variables (AOV).
* Added Adaptive Probe Volumes (APV) ability to stream data directly from the disk. This feature is only available on devices with compute shader compatibility.
* Added the ability to use the baking API to bake Adaptive Probe Volumes (APV) independently from lightmaps or reflection probes.

## Multiplayer

### Netcode for GameObjects

* Added the **Refresh In-Scene **Prefab** Instances** option in the Network Object component context menu. Use this property to update the `GlobalObjectIdHash` value for prefab instances in the scene that existed before they had a `NetworkObject` component. This property affects all scenes included in the build list.
* Added the `NetworkManager` methods `SetPeerMTU` and `GetPeerMTU` to give full control over Maximum Transmission Unit (MTU) sizes. Use this for custom protocols in Netcode for GameObjects.
* Improved network prefab identification generation (for example, `GlobalObjectIdHash`) to fix issues where Unity assigned invalid values.
* Improved the serialization API and codegen pipelines in the following ways:
  + Added the `GenerateSerializationForTypeAttribute` method that you can apply to any class or method to ensure the specific type is included in the codegen serialization process.
  + Added the `GenerateSerializationForGenericParameterAttribute` method that you can use to include generic type(s) wrapped by a `NetworkVariable` type in the codegen process.
  + Exposed additional `NetworkVariableSerialization` methods to improve custom `NetworkVariable` creation without any boxing cost. You can use the `NetworkVariableBase.MarkNetworkBehaviourDirty` method to mark NetworkVariables as dirty.
* RPCs in generic `NetworkBehaviour` types can now serialize parameters of the class’s generic types.
* Improved support for 32-bit ARMv7 in Netcode.
* Added additional `NetworkManager` events `OnServerStarted`, `OnServerStopped`, `OnClientStarted`, and `OnClientStopped` to provide better notification of the `NetworkManager` initialization process.

### Dedicated server platform

* Added the **Dedicated Server Optimizations** option in the **player settings** that, when enabled, strips all shaders from server builds. Enable this option to considerably reduce the build time.

## Unity Transport

* Unity Transport 2.X is the now the default and recommended version of [Unity Transport](https://docs.unity3d.com/Packages/com.unity.transport@latest).
* It is now possible to configure the maximum message size that the transport sends through a new `maxMessageSize` parameter in `NetworkSettings.WithNetworkConfigParameters`. This is useful for environments where network equipment mishandles larger packets (like some mobile networks or VPNs). The value excludes IP and UDP headers, but includes headers added by the transport itself (for example, reliability headers). The default value is 1400. Note that both client and server should use the same value.

## Package Manager

### Changes to package cache management

* Changed the default location and structure of the [global cache](upm-cache.md).
* The registry data cache is now limited to a default maximum size of 10 GB. You can override the default by [customizing the global cache](upm-config-cache.md).

### Changes to the Package Manager window

* Added a [navigation panel](upm-ui-nav.md):
  + Added a panel to improve the navigation between views of package subsets. In earlier versions, these contexts were stored in the **Packages** dropdown menu.
  + Added **Services** as a dedicated entry in the Package Manager navigation panel.
  + Added **Updates** as a nested entry beneath **In Project** in the navigation panel. This nested view lists all the packages in your project that have updates available.
  + Added individual scoped registries as nested entries beneath My Registries in the navigation panel.
* Enhancements to the [list panel](upm-ui-list.md):
  + Added a **Packages - **Asset Store**** expander in the list panel of the **In Project** context, so you can view and manage Asset Store packages from the **In Project** view.
  + Relocated the [search box](upm-ui-search.md) and changed its behavior so it recalls your search terms on a per-context basis.
* Enhancements to the [details panel](upm-ui-details.md):
  + Improved the flow for managing packages in the **My Assets** context by streamlining the action buttons in the details panel. The most common action appears as the default action in a menu button, with additional actions listed in the menu. Also added an **In Project** label to clearly indicate when an Asset Store package is already imported into your project and up-to-date.
  + Updated the label of the **Import** button when viewing an Asset Store package. After you [download an updated Asset Store package](upm-ui-update2.md), the main action button label is **Import update #.# to project**, so it’s clear that you need to import the update to your project.
  + Updated the behavior of the [Documentation](upm-docs.md#Editor) link for packages that you installed from a registry. You can now right-click the link and choose to **Open in browser** or **Open locally**.
  + Added a button to sign in with your Unity ID when you’re logged out and try to view the details of an Asset Store package.
* Other enhancements:
  + Updated the values in the [Filters](upm-ui-filter2.md) menu when you’re viewing the My Assets context.
  + Updated the label for the [Filters](upm-ui-filter2.md) control so it shows the selected value rather than the parent category.

## Android

* Added Addressables for Android package (`com.unity.addressables.android`) to provide Play Asset Delivery support for Addressables.
* Added texture **compression** format targeting support through Addressables for Android package to pack multiple **texture compression** formats to Android asset packs. At install-time, an **APK** is built only using the optimal texture compression formats based on the device’s mobile GPU to load and **render textures** quicker and with less memory.
* Added `ApplicationExitInfo` API to provide information on application crashes and application not responding (ANR) errors.
* Made these improvements:
  + Made `GameActivity` as the default application entry point to provide more control over the interaction between Android and your application.
  + Raised the minimum supported Android version to 6.0 (API level 23).

## Profiler

* Added metadata support for AudioClip and Shader in Memory Profiler.
* Added a Highlights module to the Profiler window.
* Added a new Inverted Hierarchy view for the CPU Profiler.

## Raytracing API

* Ray Tracing Acceleration Structure Build Flags: New Raytracing Acceleration Structure build flags were introduced to Unity’s renderers and RTAS API. These flags allow you to control the tradeoffs between ray tracing memory usage, RTAS build time and ray tracing performance. Developers and artists can use the new flags to fine tune the ray tracing performance in their scenes and renderers.
* Inline Ray Tracing in Shaders: You can now use in-line ray tracing in **rasterization** and compute shaders when targeting DXR1.1-capable Windows platforms, Xbox Series X/S, and Playstation 5. You can issue ray queries from within shaders to traverse the bound Raytracing Acceleration Structure and perform intersection testing. This allows you to implement all kinds of hardware accelerated raytracing effects and simulations.

## Shader Graph

* Added UGUI support for Shader Graph. Shader Graph now has subtargets for Canvas in both URP and HDRP. UI artists can create custom shaders for their UI elements using Shader Graph. Define the overall look and style of UI elements, create animated UI effects, and define custom button states while using less texture memory.
* Enabled Shader Graph Canvas Master Node to allow users to create UI shaders for Canvas in HDRP, URP, and Built-in.

## SpeedTree

Improved SpeedTree visual quality in HDRP using a Transmission Mask to apply subsurface scattering only on leaves. HDRP/Nature/SpeedTree8.shadergraph to use its Subsurface Map for the Transmission Mask node to remove the unintended light transmission from tree barks and twigs. This also fixes the overly bright **billboard** lighting not matching the 3D geometry’s lighting.

## Splines

You can now store personalized data on Spline objects. Additionally, the interface for editing points in the Inspector has been enhanced and certain APIs have been made available to the public. When you work on splines, you can now use the new [Scene view context menu](SceneViewContextMenu.md).

## Test Framework

* Enabled retrying and repeating tests based on test level. Soon as the test finishes running its first iteration, the Editor now retries or repeats the test. Pass the command line arguments to the Editor to:
  + Repeat x runs the test x amount of times or until it fails. This is useful for testing unstable tests.
  + Retry x if a test fails. This run the test x amount of times or until it succeeds.
* You can now run the tests in a randomized order, by using the Editor command line new argument `-randomOrderSeed x`, where `x` is an integer different from 0. If a new test is added to the project, the Test Framework keeps the random order passing the same seed, and places the new test in the random list accordingly.
* Added `TestFileReferences.json` to generate on a build step of the player, so that **Test runners** can consume it to enrich data for the run step.
* The UTF version now automatically updates for SRP tests.

## Version control

Added a project option to support tracking packages that exist on disk outside of the project’s root folder.

## VFX Graph

* Added VFX Graph assets with a predefined effect to the template window. You can use these templates as a starting point for your own effects.
* Added Custom HLSL blocks and operators. Custom HLSL nodes let you execute your own code during particle simulation. You can use an operator for horizontal flow or a block for vertical flow in the VFX Graph context.
* Added URP Decals with VFX Graph. URP Lit decal output are now supported in VFX Graph, making it possible to create stunning decal effects on URP that matches correctly your scene lighting and materials
* Added lighting for Shader Graph smoke effects in HDRP and URP. This means that you can create custom smoke shaders that use six-way lighting in lit shaders.
* Extended the Camera Depth and Color buffer behavior in URP. This means you can use the **depth buffer** for depth-based **collisions** or sample the color buffer to create custom effects.
* Added support for motion vectors in URP.
* Added support for motion vectors with Shader Graph.
* Added the ability to use VFX Graph’s Shade Graph integration to use the material variant workflow to override settings in the VFX Output.
* Enabled VFX instancing with exposed textures, meshes, or graphic buffers.

## Web platform (previously WebGL)

* Unity web builds now take advantage of the latest size and performance optimizations in the Emscripten toolchain.
* Added support to specify the browser type and its executable path that you want your application to launch at runtime. You can specify this either using the command line or the GUI setting in the **Build Settings** window.

## XR

* Added Hololens Automation Support.
* Extended Unity’s integrated support for tone-mapping and outputting to HDR Displays in URP, HDRP and the built-in render pipeline to provide support for XR devices that have a HDR display.

## Additional resources

* [New in Unity 6.0 Preview](WhatsNewUnity6Preview.md)
* [New in Unity 2023.1](WhatsNew20231.md)