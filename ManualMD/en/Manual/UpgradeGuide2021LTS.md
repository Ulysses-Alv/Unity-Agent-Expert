# Upgrade to Unity 2021 LTS

This page lists changes in the Unity 2021 LTS version which might affect existing projects when you upgrade from a 2020 version to 2021 LTS.

**Note**: 2021 LTS is also known as 2021.3.

### Page Outline

* [Render pipelines](#ren-pipe)
* [Device Simulator](#Device)
* [Environment Lighting](#Lighting)
* [Environment lighting: Ambient probe and skybox reflection probe are now automatically baked](#LightingProbe)
* [Enable Code Coverage preference](#CodeCoverage)
* [Particle System Force Fields](#Particles)
* [Particle System Start Delay + Rate over Distance emission](#ParticleDelay)
* [BuildReport - PackedAssets](#Build)
* [Terrain APIs out of experimental (WIP)](#Terrain)
* [Texture2D.Resize renamed to Reinitialize](#Texture)
* [Android changes](#Android)
* [UI Toolkit - Default scaleMode for Image changed](#UI)
* [Mono upgrade behavior changes](#Upgrade)
* [Adaptive Performance](#Performance)
* [RenderTexture DepthStencilFormat](#Render)
* [Graphics formats DepthAuto, ShadowAuto and VideoAuto are deprecated](#Graphics)
* [WebGL: Updated Emscripten to 2.0.19](#WebGL)
* [Progressive GPU Lightmapper dropped support for CPU OpenCL devices](#Lightmapper)
* [OnPostprocessAllAssets behavior changes](#Process)
* [Mixed point and spot lights with no shadows bake direct light in subtractive lighting mode](#BakeLighting)
* [Shader keyword system improvements](#Shader)
* [Unity supports .NET Standard 2.1 APIs](#NET)
* [Windows XR plugin removed](#Windows)

## Render pipelines

This upgrade guide describes how to upgrade to version 2021 of Unity’s built-in render pipeline.
To upgrade to other render pipelines to version 2021, see:

* [The URP upgrade guide](urp/upgrade-guide-2021-2.md)
* [The HDRP upgrade guide](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@14.0/manual/Upgrading-from-2021.1-to-2021.2.html)

To upgrade other packages, refer to the documentation for the packages you are using.

## Device Simulator

The [Device Simulator](https://docs.unity3d.com/Packages/com.unity.device-simulator@latest?subfolder=/manual/index.html) is now part of the Editor and is accessible from the Game window.
To set up the Device Simulator, add the `UnityEngine.Device` namespace to the Screen, Application, and SystemInfo classes:
`UnityEngine.Device.Screen;`
`UnityEngine.Device.Application;`
`UnityEngine.Device.SystemInfo;`

To switch to `UnityEngine.Device`, add the following logic to each script you want to use with the simulator:
`lang-cs
using Screen = UnityEngine.Device.Screen;
using Application = UnityEngine.Device.Application;
using SystemInfo = UnityEngine.Device.SystemInfo;`
The new namespace `UnityEngine.Device` transitions smoothly from Simulator (when in Editor) to actual device API with a runtime build.

## Environment Lighting

The Editor now automatically bakes the default [skybox](skyboxes-using.md) probe and [ambient probe](../ScriptReference/RenderSettings-ambientProbe.md) and retains that data until you manually bake the **Scene**. When you upgrade, Scenes with no **ambient light** contribution might visually change. To restore the original look of these Scenes, set the Environment lighting [intensity multiplier](lighting-window.md#EnvironmentSection) to 0. Alternatively, set the skybox to black, bake the scene, then reset the skybox to your preferred sky color.

## Environment lighting: Ambient probe and skybox reflection probe are now automatically baked

Unity’s Progressive ****Lightmapper**** now automatically generates the ambient probe and the skybox **reflection probe** for every scene by default. This means that a scene automatically receives environment lighting in accordance with the settings in the **Environment** tab in the **Lighting** settings panel. The Editor updates the ambient probe and skybox reflection probe every time the environment lighting changes, until you generate lighting. When you bake with the **Generate Lighting** control, the Editor stops updating the probes, and only updates them again at the next bake. When you enable the **Auto Generate** option, the Editor continues to update the probes every time the environment lighting changes. If you generate lighting and then delete this lighting data by removing the Lighting Data Asset from the project, the Editor automatically generates the ambient probe and skybox reflection probe again.

There is one situation which requires action when you upgrade a project. This is when you don’t wish to have any Environment lighting contribution in a project which also:

* Doesn’t have a Lighting Data Asset.
* Doesn’t have **Auto Generate** enabled.
* Has an Environment contribution set to a color other than black.

In this situation, navigate to **Window** > **Rendering** > **Lighting** > **Environment** and disable the environment contribution of the automatically generated ambient probe and skybox reflection probe by making one of the following changes:

* Option 1: Set the **Intensity Multiplier** to 0.
* Option 2: Use a black **Skybox Material**.
* Option 3: Use a black color in either **Color** or **Gradient** mode for the **Source**.

## Enable Code Coverage preference

The user interface for managing Code Coverage has moved from General Preferences to within the [Code Coverage package](https://docs.unity3d.com/Packages/com.unity.testtools.codecoverage@latest).

![In Unity 2021 the Enable Code Coverage checkbox was moved in the Code Coverage window.](../uploads/Main/upgrade211CodeCoverage.png)

In Unity 2021 the Enable Code Coverage checkbox was moved in the Code Coverage window.

The Code Coverage package is available as a released package via the Package Manager for Unity 2019.3 and above. The latest version is 1.0.0.

You can use one of the following methods to enable code coverage:

* [Install](https://docs.unity3d.com/Packages/com.unity.testtools.codecoverage@1.0/manual/InstallingCodeCoverage.html) the Code Coverage package, open the [Code Coverage window](https://docs.unity3d.com/Packages/com.unity.testtools.codecoverage@1.0/manual/CodeCoverageWindow.html) (**Window** > **Analysis** > **Code Coverage**) and select [**Enable Code Coverage**](https://docs.unity3d.com/Packages/com.unity.testtools.codecoverage@1.0/manual/Quickstart.html#4-enable-code-coverage-sub1-minsub).
* Use the [Command line argument](CommandLineArguments.md) `-enableCodeCoverage` in batchmode.
* Use the [`Coverage.enabled`](../ScriptReference/TestTools.Coverage-enabled.md) API. Example class:

```
// Create a new C# script called CodeCoverageMenuItem and place it
// under the Editor folder.
// This class creates a toggle menu item under Code Coverage > Enable
// Code Coverage. Use it to enable/disable Code Coverage.

using UnityEditor;
using UnityEngine.TestTools;

class CodeCoverageMenuItem
{
    const string EnableCodeCoverageItemName = "Code Coverage/Enable Code Coverage";

    [MenuItem(EnableCodeCoverageItemName, false)]
    static void EnableCodeCoverage()
    {
        Coverage.enabled = !Coverage.enabled;
    }

    [MenuItem(EnableCodeCoverageItemName, true)]
    static bool EnableCodeCoverageValidate()
    {
        Menu.SetChecked(EnableCodeCoverageItemName, Coverage.enabled);
        return true;
    }
}
```

## Particle System Force Fields

Previously, some Force Field properties behaved differently at different frame rates (or if using the Time Scale in the Time Manager settings)

The **Particle** System now uses a reference frame rate of 30fps to serve as the basis for the simulation. If your app runs at a different frame rate, the following settings might behave differently compared to earlier versions of Unity:

* Gravity
* Rotation
* Vector Fields

If these settings are affected, adjust the strengths of the affected areas to get the desired appearance.

## Particle System Start Delay + Rate over Distance emission

Previously, Rate over Distance emission ignored the Start Delay setting. Now, if the Start Delay setting is defined, it delays the start for the distance-based emission.

If this field was previously set, you might need to adjust it.

## BuildReport - PackedAssets

PackedAssets.file has been marked obsolete without a direct replacement. Previously this held an integer signifying a file id or index into BuildReport.files
To look up BuildReport files now, use PackedAssets.shortPath.

## Terrain APIs out of experimental (WIP)

The experimental **Terrain** APIs have been moved to non-experimental namespaces. There have also been some other minor changes to the Terrain API. If you used the experimental Terrain APIs, use the following APIs instead:

* `UnityEngine.TerrainTools`;
* `UnityEditor.TerrainTools`;
* `UnityEngine.TerrainUtils`;

Here is the full list of API changes:

* In most cases, `UnityEngine.Experimental.TerrainAPI` and `UnityEditor.Experimental.TerrainAPI` are now `UnityEngine.TerrainTools` and `UnityEditor.TerrainTools`, respectively. Some runtime APIs have moved to a new `UnityEngine.TerrainUtils` namespace.
* The `GetDesc()` of the `TerrainPaintTool<T>` class was renamed to `GetDescription()`.
* The `TerrainUtility` class has moved from `UnityEngine.Experimental.TerrainAPI` to `UnityEngine.TerrainUtils`.
* The `TerrainUtility.TerrainMap` class is no longer an inner-class and belongs to the `UnityEngine.TerrainUtils` namespace.
* The `TerrainMap.TileCoord` structure is no longer in the `TerrainMap` class, has been renamed to `TerrainTileCoord`, and is now also part of the `UnityEngine.TerrainUtils` namespace.
* The `UnityEditor.Experimental.TerrainAPI.BrushPreviewMode` enum has been renamed to `TerrainBrushPreviewMode` and moved to the `UnityEditor.TerrainTools` namespace.
* The `TerrainPaintUtilityEditor.BuiltinPaintMaterialPasses` enum has moved from the `TerrainPaintUtilityEditor` class and into the `UnityEditor.TerrainTools` namespace. It has also been renamed to `TerrainBuiltinPaintMaterialPasses`.
* The three `ShowBrushGUI` functions in `IOnInspectorGUI` have been merged into one function with default parameter values instead of the different overloaded functions.
* `TerrainFilter` has been removed. Use `System.Predicate<Terrain>` instead.

## Texture2D.Resize renamed to Reinitialize

`Texture2D.Resize` and its overloads have been renamed to `Texture2D.Reinitialize`.

The API Updater should rename this automatically. If not, change any usage of `Texture2D.Resize` to `Texture2D.Reinitialize`.

## Android changes

A large part of the Android build pipeline is now incremental and Unity removed the following features that were in the previous build pipeline:

* Unity no longer copies **Gradle** project assets located in **Assets**/**Plugins**/**Android**/**[res, assets]** to the Gradle project.
  + Previously you could place Gradle resources in this folder and Unity copied them into the Gradle project. You should now use [AAR or androidlib plug-ins](AndroidAARPlugins.md) to pass additional Gradle resources to the application.
  + If you place project assets in this folder, Unity shows a build error message.
* Unity no longer ignores files inside the **GENERATED BY UNITY. REMOVE THIS COMMENT TO PREVENT OVERWRITING WHEN EXPORTING AGAIN** comment.
  + Previously Unity wouldn’t overwrite files if you removed this comment. If you didn’t remove the comment, Unity always regenerated build.gradle, manifest, and UnityPlayerActivity files.
  + If you want your changes to persist using the new build pipeline, use templates.
* When you export an Android project, Unity no longer creates a symbols zip package. The symbols are now available in the **unityLibrary\symbols** directory which you can zip. The reason for this change is that not all symbols files are available when you export a project; Gradle produces the `libil2cpp.so` symbols when it builds the application.
* Changed how Unity checks to see if an obb is compatible with an **apk**. Both the apk and obb now have unity\_obb\_guid file inside them and if the contents match between them, Unity treats them as being compatible.
* For custom build **scripts** that use [PatchPackage](../ScriptReference/BuildOptions.PatchPackage.md), note that **Patch/Patch & Run** now works with all types of assets and doesn’t require a Script Only Build.

The default [mainTemplate.gradle](gradle-templates.md) file has changed. If you use a custom main template, you must regenerate it and reapply your changes on top. Otherwise, your application can experience performance regressions if it uses [Resources.Load](../ScriptReference/Resources.Load.md).

## UI Toolkit - Default scaleMode for Image changed

Default Image.scaleMode was changed from ScaleAndCrop to ScaleToFit.

The expected behavior for an image is to scale to the size of the element, so we changed the default value of Image.scaleMode to ScaleToFit.
If you didn’t override the Image scale mode, some cropped images might shrink to fit the size of the element.
If ScaleAndCrop was the expected mode for your images, you can override their style by adding the following value in your UXML file inline style:

`-unity-background-scale-mode: scale-and-crop;`

You can also create a style class with the override and apply it to images that require ScaleAndCrop.

## Mono upgrade behavior changes

The underlying C# runtime, Mono, has been upgraded in the latest version. This includes many fixes from the upstream version of Mono, and some notable changes in behavior.

* `Directory.GetFiles` is no longer guaranteed to return a sorted list.
  + Previously, this always returned an alphabetically sorted list of items. If your project needs items to be returned in the same order every time, sort the returned list. For example:
    var files = `Directory.GetFiles(dir).OrderBy(f => f)`;
* `Object.GetHashCode` now returns a different value and shouldn’t be relied on as a deterministic hashing algorithm between operating systems.
  + In general, you shouldn’t use the result of `GetHashCode` outside of the current process, that is don’t serialize it or expect it to be the same the next time the code runs in a new process. Unity recommends using a deterministic hashing algorithm such as MD5.
* Some bug fixes have resulted in new exceptions being thrown. Additionally, the content of some exception messages have changed.
  + This new behavior will be particularly noticeable in automated testing scenarios and may require changing the expected behavior if the tests are parsing the logs for specific exception messages.

## Adaptive Performance

Version 3.0 of the Adaptive Performance package is now available. For information on how to upgrade to version 3.0, see the [Adaptive Performance upgrade guides](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@3.0/manual/upgrade-guides.html).

## RenderTexture DepthStencilFormat

Previously, if you set the [`RenderTexture.depth`](https://docs.unity3d.com/2021.2/Documentation/ScriptReference/RenderTexture-depth.html) property to 32-bit, you could get D24\_S8 depending on the platform. Now if you set it to 32-bit, you get D32\_S8 with 32-bit for the depth component if that format is supported on the current platform. However, this doubles the memory usage for that **depth buffer**.

The new [`RenderTexture.depthStencilFormat`](https://docs.unity3d.com/2021.2/Documentation/ScriptReference/RenderTexture-depthStencilFormat.html) property returns the format that the graphics API uses to create the resource in video memory. You can also use this property to request a specific format. However, not all platforms support all depth stencil formats. When you set the `DepthStencilFormat` property to an unsupported format, Unity automatically selects a compatible format that has the equal or greater amount of bits for the depth and stencil components.

### Upgrading

The RenderTexture asset now serializes the depth stencil format that you selected. If you use an API that takes a number of bits instead of a format, these bits get mapped to a format and that format is serialized. RenderTexture assets from previous versions that had the depth set to more than 16-bits are automatically upgraded to use D24\_S8.

On some platforms (for example Windows) that use the DirectX graphics API, this results in a format with fewer depth bits because the graphics backend selects the D32\_S8 format internally if you set bits>16. To ensure consistent upgrades across all platforms, D24\_S8 is used on all platforms for the automatic upgrader. However, this might introduce visual artifacts in the render output of your project if you have RenderTexture assets in your project. Review these assets and change the depth stencil format to D32\_S8 when needed. The following issues might occur:

* The memory size might increase for the RenderTextures that you use. However, the memory reported is now correct.
* Now that setting the depth property to 32-bit gives you a 32-bit depth component, the memory usage increases. To avoid this, set the bits to 24 and use D24\_S8 if that is supported on your platform.
* If a platform doesn’t support D24\_S8, Unity falls back to the compatible format D32\_S8 by default. To prevent this, in the **Inspector** window, disable the [Enable Compatible Format](class-RenderTexture.md) property on an asset. If Unity doesn’t support the format D32\_S8 and it can’t fall back, you see the error message: **RenderTexture.Create failed: depth/stencil format unsupported. There is no compatible format on this platform or this fallback to a compatible format is disabled in the import inspector.**. In most cases, to fix the issue you can turn on **Enable Compatible Format**.

## Graphics formats DepthAuto, ShadowAuto and VideoAuto are deprecated

The following graphics formats have now been deprecated:

* DepthAuto
* ShadowAuto
* VideoAuto

These **Auto** formats are unclear about the exact format that’s used and might vary by platform.

The steps to remove the use of these deprecated formats depends on the format and use case.

### For VideoAuto

To get the current platform’s automatic video format, use `SystemInfo.GetGraphicsFormat(DefaultFormat.Video)`.

### For DepthAuto/ShadowAuto used to indicate a depth-only render texture.

The GraphicsFormat API often uses DepthAuto or ShadowAuto to create **render textures** with depth-only rendering, and no color buffer. Examples of this use case are:

* `renderTextureDescriptor.graphicsFormat = GraphicsFormat.ShadowAuto`
* `RenderTexture.GetTemporary(width, height, bits, GraphicsFormat.ShadowAuto)`

To indicate depth-only (non-color) rendering, use GraphicsFormat.None as the new color format.
`renderTextureDescriptor.graphicsFormat = GraphicsFormat.None;`

If you used ShadowAuto, set the shadowSamplingMode of your RenderTextureDescriptor to ShadowSamplingMode.CompareDepths to enable depth-compare sampling on the depth texture, and change the code to use an overload that takes a RenderTextureDescriptor.
`renderTextureDescriptor.shadowSamplingMode = ShadowSamplingMode.CompareDepths;`

### For DepthAuto/ShadowAuto in other situations

In some situations the DepthAuto/ShadowAuto formats represented an automatically selected depth format appropriate for the current platform. To replace the deprecated values in this case, use `SystemInfo.GetGraphicsFormat(DefaultFormat.Depth/Shadow)`

## WebGL: Updated Emscripten to 2.0.19

The asm.js Linker target that was available for advanced users is no longer available.

* In Unity 2021.2, the underlying Emscripten compiler used by the **WebGL** platform was updated to version 2.0.19. This upgraded the native code object file format, and therefore requires recompiling all native code plugins (C/C++ code plugins) in a project. If you are using closed source 3rd-party plugins from Unity **Asset Store**, for example, remember to ask the author for an updated version for Unity 2021.2.
* The Emscripten runtime JavaScript function `Pointer_stringify()`is now deprecated. Instead, call the function `UTF8ToString()` to marshal a UTF8-encoded null-terminated C string from the WebAssembly heap to a JavaScript string.

## Progressive GPU Lightmapper dropped support for CPU OpenCL devices

The Progressive GPU Lightmapper no longer supports CPU OpenCL devices. If no supported GPU is found but a CPU OpenCL device is detected, a warning message informs you that the device is skipped and falls back to the Progressive CPU Lightmapper. The Progressive CPU Lightmapper offers better performance for CPU-based computation of **lightmaps**.
This change of behavior will happen automatically and lightmaps should be computed as expected.

## [OnPostprocessAllAssets behavior changes](https://docs.unity3d.com/2021.2/Documentation/ScriptReference/AssetPostprocessor.OnPostprocessAllAssets.html)

When Unity calls `InitializeOnLoad` methods during the asset import process, asset loading can fail. During asset import, the asset database is in an update state and Unity can’t determine which assets have already been imported. `InitializeOnLoad` methods can’t load assets that haven’t been imported.

To improve the asset import process, the `OnPostprocessAllAssets` callback has been enhanced. In particular, the `OnPostprocessAllAssets` callback:

* Can perform initialization after a domain reload, when the asset database is guaranteed to be in a ready state.
* Contains the `didDomainReload` parameter, which is set to true if the domain has reloaded.
* Supports asset operations. This includes asset loading, which behaves in the same way as loading assets via a menu item.
* Has moved out of the asset database importing loop. Callback processing is postponed until asset import is complete.

Move any domain-related initialization logic that requires asset operations to the `OnPostprocessAllAsset` callback; do not perform asset operations inside `InitializeOnLoad` methods.

The following behavior change code examples show how asset operations were previously postponed.

Example 1:

```
public class AssetPostprocessorTester1 : AssetPostprocessor
{
    static void OnPostprocessAllAssets(string[] importedAssets, string[] deletedAssets, string[] movedAssets, string[] movedFromAssetPaths)
    {
        var assetPath = "Assets/hello.txt";

        if (File.Exists(assetPath))
        {
            var txtObj = AssetDatabase.LoadAssetAtPath<TextAsset>("Assets/hello.txt");
            AssetDatabase.DeleteAsset("Assets/hello.txt");

            if (txtObj == null)
                Debug.Log("New Behaviour: Asset object is unloaded");
            else
                Debug.Log("Old Behaviour: Asset is loaded for deleted asset!!");
        }
    }
}
```

Example 2:

```
 public class AssetPostprocessorTester2 : AssetPostprocessor
{
    static void OnPostprocessAllAssets(string[] importedAssets, string[] deletedAssets, string[] movedAssets, string[] movedFromAssetPaths)
    {
        var assetPath = "Assets/SomeText.txt";
        if (!File.Exists(assetPath))
        {
            File.WriteAllText(assetPath, "hello world");

            AssetDatabase.ImportAsset(assetPath);
            var txtObj = AssetDatabase.LoadAssetAtPath<TextAsset>(assetPath);

            if (txtObj == null)
                Debug.Log("Old Behaviour: Asset hasn't been imported yet");
            else
                Debug.Log("New Behaviour: Asset is imported and loaded");
        }
    }
}
```

The following example has the new `OnPostprocessAllAssets` variant with the `didDomainReload` parameter:

```
static void OnPostprocessAllAssets(string[] importedAssets, string[] deletedAssets, string[] movedAssets, string[] movedFromAssetPaths, bool didDomainReload)
{
    if (didDomainReload)
        Debug.Log("Domain has been reloaded");
    else
        Debug.Log("Domain did not reload during import");
}
```

All domain reloads are now processed inside the asset database.

`OnPostprocessAllAssets` now works better with asset operations, but any processing in this callback increases asset database refresh and domain reload time. `InitializeOnLoad` methods also add to domain reload time. It is best practice to minimize processing in these callbacks to improve Editor responsiveness between iterations.

## Mixed point and spot lights with no shadows bake direct light in subtractive lighting mode

[Mixed mode](lighting-mode-landing.md) point and spot lights now consistently contribute baked direct light in Scenes using [Subtractive lighting mode](lighting-mode.md#subtractive), regardless of their [Shadow Type](class-Light.md) setting. As a result, specular lighting for Static **GameObjects** may appear to be missing in affected Scenes. To resolve this issue, replace affected Mixed mode lights with [Realtime mode Lights](LightModes-introduction.md#realtime). Alternatively, use [Baked Indirect](lighting-mode.md#baked-indirect) or [Shadowmask](lighting-mode.md#shadowmask) lighting mode(s) with Mixed lights.

## Shader keyword system improvements

The **shader** keyword system now allows up to 65534 local keywords per shader or compute shader and 232–2 global keywords per project. All keywords you declare in a shader or compute shader are now local to this shader. Keywords you declare in a directive with the \_local suffix aren’t affected by the state of global keywords.

Example:

A pass in a shader declares the following keywords:

* `#pragma shader_feature FOO BAR`
* `#pragma shader_feature_local BOO BAZ`

When using this pass, keywords FOO and BAR are enabled if they’re enabled either globally or on the material. Keywords BOO and BAZ are only enabled if they’re enabled on the material.

## Unity supports .NET Standard 2.1 APIs

Unity now supports many additional APIs in the .NET base class libraries, including all APIs in the .NET Standard 2.1 API. Your project might fail to compile if any of your code has conflicts with new APIs.

To avoid errors when you upgrade projects created in previous versions of Unity, check and update your code to ensure that there are no conflicts with types and methods that are now available in .NET Standard 2.1.

Sources of conflict include the following:

* ambiguous references
* extension methods
* precompiled assemblies

### Resolve ambiguous references

If your code implements types or methods with a name that conflicts with types or methods that .NET Standard 2.1 adds, your code fails to compile. Name conflicts can lead to C# compiler errors due to ambiguous references.

For example, if you add a type named `MyCompany.MyCode.Range` to your project’s code, this might conflict with the existing `System.Range` type. Code that contains both `using System;` and `using MyCompany;` statements fails to compile.

To prevent errors, fully specify the namespace in the C# code for any types that have conflicting names.

### Resolve conflicting extension methods

If your existing code has an extension method that .NET Standard 2.1 now implements on a type directly, you can choose to rename your extension methods or to use the methods implemented by the base class library.

For example, code in your project might implement an extension method named `CopyTo` on the `ArraySegment` type. In .NET Standard 2.1, `ArraySegment` has a built-in `CopyTo` method. You can choose to rename your `CopyTo` extension method, or eliminate it entirely and use the built-in one.

### Resolve conflicting assemblies

If your project uses precompiled assemblies (that is, managed plugins) that implement types and methods that are now part of the base class library, then remove these assemblies from your project and use the built-in implementations.

For example, in previous versions of Unity, you needed to use the `System.Memory.dll` assembly from NuGet to access the `System.Span` value type. Now, .NET Standard 2.1 provides `System.Span` in the base class library. If you try to use the `System.Memory.dll` managed plugin in Unity 2021.2, the project fails to build.

## Windows XR plugin removed

Microsoft has deprecated the Windows **XR** plugin and now supports Windows **Mixed Reality** (WMR) features and devices through the OpenXR plugin.

* The Windows XR plugin has been removed as of Unity 2021.2.
* When you open an existing project in Unity 2021.2, the update process removes the Windows XR plugin, if present.
* To continue to use the project for Windows Mixed Reality, you must enable the Unity OpenXR plugin.
* Depending on which features you are using, you must also install Microsoft’s Mixed Reality OpenXR support plugin and possibly additional tools from the Microsoft Mixed Reality Toolkit.

To upgrade to the Unity OpenXR plugin:

1. Open the ****Project Settings**** window.
2. In the **XR **Plug-in** Management** section, enable **OpenXR** in the **Plug-in Providers** list.
   If needed, Unity downloads and installs the OpenXR plugin.

Once the OpenXR plugin is enabled, you can use the **Mixed Reality Feature Tool** provided by Microsoft to install the required supporting packages.

To install or update WMR features, tools, and samples:

1. Download and run the [Mixed Reality Feature Tool](https://docs.microsoft.com/en-us/windows/mixed-reality/develop/unity/welcome-to-mr-feature-tool).
2. Choose the Unity Project to update and click Discover Features.
3. Under Platform Support, select the Mixed Reality OpenXR Plugin.
4. Select any additional features you wish to add.
5. Click Get features.
6. Click Import and then Approve to complete the process.
7. Return to the XR Plug-in Management section of your Unity Project Settings and enable and configure the added features.

For more information about setting up new and updated Unity projects to use Windows Mixed Reality, see [Setting up your XR configuration](https://docs.microsoft.com/en-us/windows/mixed-reality/develop/unity/xr-project-setup?tabs=openxr) on Microsoft’s documentation site.

## Additional resources

* [Install Unity](GettingStartedInstallingUnity.md)
* [API updater](APIUpdater.md)