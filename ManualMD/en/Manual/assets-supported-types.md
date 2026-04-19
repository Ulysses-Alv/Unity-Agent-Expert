# Supported asset type reference

Unity supports many different types of assets and most common image file types. It uses [asset importers](#asset-importers) to process and convert external files into assets you can use in your project.

Unity supports the following kinds of files:

* **3D model files**: Unity supports standard and proprietary **model file** formats. Internally, Unity uses the FBX file format as its importing chain. For a list of supported files, refer to [3D model importers](#3d-model-importers).
* **Image files**: Unity imports image files as textures and supports most common image file types. For a list of supported files, refer to [Image importers](#image-importers).
* **Audio and video files**: Unity supports many audio and video file formats. For a list of supported files, refer to [Audio and video importers](#audio-and-video-importers).
* ****Shader** files**: Unity supports different shader file types depending on the graphics pipeline you’re using. For a list of supported files, refer to [Shader importers](#shader-importers).
* **Text and arbitrary data**: Unity can import arbitrary data from files such as .html, .xml, .json files, which you can use to store and use data from external sources. For a list of supported files, refer to [Text and arbitrary data importers](#text-and-arbitrary-data-importers).
* **Plug-ins and code related assets**: You can add managed and native [plug-ins](plug-ins.md) files into your project to expand the functionality of your application, and [assembly definitions](assembly-definition-files.md) to create and organize your **scripts** into assemblies. For a list of supported files, refer to [Plug-ins and code importers](#plug-ins-and-code-importers)
* **Native assets**: There are a range of asset types that are native to the Unity Editor. For a list of supported files, refer to [Native asset importers](#native-asset-importers).

## Asset importers

Unity uses importers to process and convert external files into assets that can be used in your project. Each file type has a corresponding importer that handles its specific requirements.

Unity supports many asset file types via its collection of built-in importers. Most of these importers are implemented in the Unity Editor’s native code. They provide the import functionality for most of Unity’s basic asset types, such as 3D models, textures, and audio files.

Most importers have a corresponding class in the `UnityEditor` namespace that exposes the same configurable import settings as the **Inspector** window for the asset type. They also have corresponding pre-process and post-process callbacks on the [`AssetPostprocessor`](../ScriptReference/AssetPostprocessor.md) class so you can define custom behavior to run before or after asset import. For example, the import settings for the `AudioImporter` are configurable in the [Audio Clip Inspector window](class-AudioClip.md) or from code with the [`AudioImporter` class](../ScriptReference/AudioImporter.md). You can also create custom pre-import or post-import behavior for audio assets with [`AssetPostprocessor.OnPreprocessorAudio`](../ScriptReference/AssetPostprocessor.OnPreprocessAudio.md) and [`AssetPostprocessor.OnPostprocessAudio`](../ScriptReference/AssetPostprocessor.OnPostprocessAudio.md) respectively. This pattern applies for most major asset types.

### 3D model importers

Unity uses the FBX file format as its importing chain. For a list of 3D modeling software that Unity supports, refer to [Model file formats](3D-formats.md).

| **Importer** | **Description** | **Supported file formats** |
| --- | --- | --- |
| **FBXImporter** | Imports 3D model files. For more information, refer to [Importing a model](ImportingModelFiles.md). | * `.blend` * `.c4d` * `.dae` * `.dxf` * `.fbx` * `.jas` * `.lxo` * `.ma` * `.mb` * `.max` * `.obj` |
| **Mesh3DSImporter** | Imports 3D Studio Max files. For more information, refer to [Importing a model](ImportingModelFiles.md). | `.3ds` |
| **SketchUpImporter** | Imports SketchUp files. For more information, refer to [SketchUp Import Settings](class-SketchUpImporter.md) and [`SketchUpImporter`](../ScriptReference/SketchUpImporter.md). | `.skp` |
| **SpeedTreeImporter** | Imports SpeedTree files. For more information, refer to [SpeedTree Import Settings window](class-SpeedTreeImporter.md) and [`SpeedTreeImporter`](../ScriptReference/SpeedTreeImporter.md). | * `.spm` * `.st` |
| **SubstanceImporter** | Imports Substance files. | `.sbsar` |

### Audio and video importers

| **Importer** | **Description** | **Supported file formats** |
| --- | --- | --- |
| **AudioImporter** | Imports audio files. For more information, refer to [Audio files](AudioFiles.md) and [`AudioImporter`](../ScriptReference/AudioImporter.md). | * `.aif` * `.aiff` * `.flac` * `.it` * `.mod` * `.mp3` * `.ogg` * `.s3m` * `.wav` * `.xm` |
| **VideoClipImporter** | Imports video files. For more information, refer to [Video sources](video-sources.md) and [`VideoClipImporter`](../ScriptReference/VideoClipImporter.md). | * `.asf` * `.avi` * `.dv` * `.m4v` * `.mov` * `.mp4` * `.mpg` * `.mpeg` * `.ogv` * `.vp8` * `.webm` * `.wmv` |

### Image importers

Unity imports image files as textures. Unity supports most common image file types, such as `.bmp`, `.tif`, `.tga`, `.jpg`, `.svg`, and `.psd`. For more information, refer to [Import a texture](ImportingTextures.md).

| **Importer** | **Description** | **Supported file formats** |
| --- | --- | --- |
| **IHVImageFormatImporter** | Imports specialized image formats. For more information, refer to [`IHVImageFormatImporter`](../ScriptReference/IHVImageFormatImporter.md). | * `.astc` * `.dds` * `.ktx` * `.pvr` |
| **TextureImporter** | Imports texture files. For more information, refer to [Import a texture](ImportingTextures.md) and [`TextureImporter`](../ScriptReference/TextureImporter.md). | * `.bmp` * `.exr` * `.gif` * `.hdr` * `.iff` * `.jpeg` * `.jpg` * `.pct` * `.pic` * `.pict` * `.png` * `.psd` * `.tga` * `.tif` * `.tiff` * `.svg` |

### Native asset importers

There are a range of asset types that are native to the Unity Editor. You can create assets of these types using Editor features. When you create these, Unity saves the files which represent them as asset files in the `Assets` folder of your project. These include [animations](animeditor-CreatingANewAnimationClip.md), [curves](EditingCurves.md), [gradients](PartSysUsage.md), [masks](class-AvatarMask.md), [materials](Materials.md), and [presets](Presets.md).

| **Importer** | **Description** | **Supported file formats** |
| --- | --- | --- |
| **NativeFormatImporter** | Imports Unity’s native asset formats. | * `.anim` * `.asset` * `.blendtree` * `.brush` * `.buildreport` * `.colors` * `.controller` * `.cubemap` * `.curves` * `.curvesNormalized` * `.flare` * `.fontsettings` * `.giparams` * `.gradients` * `.guiskin` * `.ht` * `.mask` * `.mat` * `.mesh` * `.mixer` * `.overrideController` * `.particleCurves` * `.particleCurvesSigned` * `.particleDoubleCurves` * `.particleDoubleCurvesSigned` * `.physicMaterial` * `.physicsMaterial2D` * `.playable` * `.preset` * `.renderTexture` * `.shadervariants` * `.signal` * `.spriteatlas` * `.state` * `.statemachine` * `.terrainlayer` * `.texture2D` * `.transition` * `.webCamTexture` |
| **PrefabImporter** | Imports **prefab** files. For more information, refer to [Creating prefabs](CreatingPrefabs.md). | `.prefab` |
| **VisualEffectImporter** | Imports visual effect files. | * `.vfx` * `.vfxblock` * `.vfxoperator` |

### Plug-ins and code importers

You can add managed and native [plug-ins](plug-ins.md) such as `.dll` files into your project to expand the functionality of your application. Unity also supports [assembly definitions](assembly-definition-files.md) to help you create and organize your scripts into assemblies.

| **Importer** | **Description** | **Supported file formats** |
| --- | --- | --- |
| **AssemblyDefinitionImporter** | Imports assembly definition files. For more information, refer to [Introduction to assemblies in Unity](assembly-definitions-intro.md). | `.asmdef` |
| **AssemblyDefinitionReferenceImporter** | Imports assembly definition reference files. For more information, refer to [Introduction to assemblies in Unity](assembly-definitions-intro.md). | `.asmref` |
| **DefaultImporter** | Imports system files. | * `.rsp` * `.unity` |
| **PackageManifestImporter** | Imports package manifest files. For more information, refer to [Package manifest](upm-manifestPkg.md). | `.json` |
| **PluginImporter** | Imports plug-in files. For more information, refer to [Import and configure plug-ins](plug-in-inspector.md) and [`PluginImporter`](../ScriptReference/PluginImporter.md). | * `.a` * `.aar` * `.bc` * `.bundle` * `.c` * `.cc` * `.config` * `.cpp` * `.dylib` * `.h` * `.jar` * `.java` * `.jslib` * `.jspre` * `.kt` * `.m` * `.mm` * `.prx` * `.rpl` * `.so` * `.suprx` * `.swift` * `.winmd` * `.xib` |

### Shader importers

| **Importer** | **Description** | **Supported file formats** |
| --- | --- | --- |
| **ComputeShaderImporter** | Imports compute shader files. For more information, refer to [Writing compute shaders](class-ComputeShader.md) and [`ComputeShader`](../ScriptReference/ComputeShaderImporter.md). | `.compute` |
| **RayTracingShaderImporter** | Imports **ray tracing** shader files. For more information, refer to [Introduction to shaders](shader-introduction.md). | `.raytrace` |
| **ShaderImporter** | Imports shader files. For more information, refer to [Introduction to shaders](shader-introduction.md) and [`ShaderImporter`](../ScriptReference/ShaderImporter.md). | * `.cg` * `.cginc` * `.glslinc` * `.hlsl` * `.shader` |

### Text and arbitrary data importers

| **Importer** | **Description** | **Supported file formats** |
| --- | --- | --- |
| **LocalizationImporter** | Imports localization files. | `.po` |
| **TextScriptImporter** | Imports text and script files. For more information, refer to [Text assets](class-TextAsset.md). | * `.boo` * `.bytes` * `.csv` * `.fnt` * `.htm` * `.html` * `.js` * `.json` * `.manifest` * `.md` * `.rsp` * `.txt` * `.xml` * `.yaml` |
| **TrueTypeFontImporter** | Imports font files. For more information, refer to [Font assets](UIE-font-asset-landing.md) and [`TrueTypeFontImporter`](../ScriptReference/TrueTypeFontImporter.md). | * `.dfont` * `.otf` * `.ttc` * `.ttf` |

### Built-in scripted importers

[Scripted importers](ScriptedImporters.md) allow you to write your own custom importers for formats that Unity doesn’t natively support. Some of Unity’s own built-in importers are implemented as scripted importers because they are written in C# in core packages, rather than within the Unity Editor’s native code itself. Unity imports scripted importer assets after native importer assets.

| **Importer** | **Description** | **File formats** |
| --- | --- | --- |
| **SpeedTree9Importer** | Imports SpeedTree 9 files. For more information, refer to [`SpeedTree9Importer`](../ScriptReference/SpeedTree.Importer.SpeedTree9Importer.md). | `.st9` |
| **StyleSheetImporter** | Imports Unity style sheet files. For more information, refer to [Introduction to USS](UIE-about-uss.md). | `.uss` |
| **UIElementsViewImporter** | Imports Unity XML files. For more information, refer to [Structure UI with UXML](UIE-UXML.md). | `.uxml` |

## Additional resources

* [Scripted importers](ScriptedImporters.md)
* [Introduction to importing assets](ImportingAssets.md)
* [Asset metadata](AssetMetadata.md)
* [`AssetPostprocessor` API reference](../ScriptReference/AssetPostprocessor.md)