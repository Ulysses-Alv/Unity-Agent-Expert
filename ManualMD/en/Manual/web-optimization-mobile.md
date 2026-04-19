# Optimize Web platform for mobile

Fast load times are crucial for Web-based applications, especially on mobile devices which might need to download application resources on a slow mobile network. Slow load times can result in a poor user experience and high bounce rate. Therefore, it’s important to optimize your Web build for mobile.

Web-based mobile applications work best when you have a small build and performant code because there’s less data to download, less data to store on a user’s device, and less data to load during initialization, which speeds up load times.

For Web optimizations not specific to mobile, refer to [Optimize your Web build](web-optimization.md).

## Optimization quick reference

Use these recommended settings to make optimizations specific to Unity Web platform for mobile:

| **Recommendation** | **Description** |
| --- | --- |
| [Optimize for size](#optimize-for-size) | Disable **development build** settings and optimize your build for disk size (Disk Size with LTO). |
| [Use Brotli compression](#use-brotli-compression) | Use the Brotli **compression** method to compress your Player build. |
| [Use the Unity Addressables system](#use-the-unity-addressables-system) | Use the Unity Addressables system for assets. |
| [Optimize your audio files](#optimize-your-audio-files) | Reduce the amount of disk space your audio takes up. |
| [Optimize the graphics in your project](#optimize-the-graphics-in-your-project) | Reduce the amount of disk space your graphics take up. |
| [Change the graphical quality level](#lower-graphical-quality-level) | Lower the quality level to make a faster build. |

## Optimize for size

A smaller build size is better for mobile because there’s less data to download, which usually makes load times shorter and uses less storage on the user’s device.

To make your build as small as possible, disable development build settings and optimize your build for disk size:

1. Go to **File** > **Build Profiles**.
2. Select **Web**.
3. Disable **Development Build** if enabled.
4. Set **Code Optimization** to **Disk Size with LTO**.

**Note**: Optimized builds with LTO can take a long time. Only use this optimization for final release or if you need to test performance. Use a faster build option during development.

For other settings that impact build size, refer to [Recommended Player settings to optimize your Web build](web-optimization-player.md) and [Remove unused resources from your Web build](web-optimization-remove-resources.md).

### Use C# to enable the optimize for size setting

If you have a script that edits settings and you want to enable the **Disk Size with LTO** setting with C#, add the following code to your script:

```
// Set Platform Settings to optimize for disk size (LTO)
UnityEditor.WebGL.UserBuildSettings.codeOptimization = UnityEditor.WebGL.WasmCodeOptimization.DiskSizeLTO;
```

## Use Brotli compression

Brotli compression is a setting where Unity pre-compresses for you while it builds. Brotli-compressed files are smaller than gzip-compressed files, which can reduce your build size.

However, Brotli takes a longer time to compress. Also, Chrome and Firefox only natively support Brotli compression over HTTPS. If that’s not suitable for your application, consider Gzip compression instead.

To use the Brotli compression setting:

1. Make sure your web server is configured to serve Brotli files with the correct encoding. For more details, refer to [Deploy Web application](webgl-deploying.md).
2. Access the **Player** settings (Menu: **Edit** > **Project Settings** > **Player**).
3. Click on the **Web settings** tab.
4. Expand the **Publishing Settings** section.
5. Set **Compression Format** to **Brotli**.

### Use C# to change the compression format

If you have a **project settings** script and want to change the compression format with C#, add the following code to your script:

```
// Set the compression format to Brotli
PlayerSettings.WebGL.compressionFormat = WebGLCompressionFormat.Brotli;
```

## Use the Unity addressables system

Mobile browser applications need short load times to decrease user bounce rates. To keep initialization times short, instead of loading all assets at startup, use addressables to load assets only when your application needs them. Defer the loading of certain assets until after your game loads.

For these changes, make sure to update `StreamingAssets/aa/catalog.json` with any new addressable files you make.

For more information about addressables and how to set them up, refer to [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest).

To optimize your addressables further, try the following:

### Sort addressables into groups

Sort your addressables into groups so that you reduce the number of asset bundles. The smaller the amount of asset bundles, the smaller the build size. For more information about addressable groups and how to create them, refer to [Manage and create groups](https://docs.unity3d.com/Packages/com.unity.addressables@2.0/manual/groups-create.html).

## Optimize your audio files

If you have a lot of audio files in your project, it’s best to compress your files to lower the size of the audio files. However compressed audio can result in lower quality audio. For information about compression formats, refer to [Audio Clip](class-AudioClip.md).

## Optimize the graphics in your project

Complex graphics can use up resources and make your build larger and less performant. But there are ways to optimize your graphics to reduce your build size and improve performance.
To optimize the graphics in your build, try the following:

### Use ASTC compressed textures where possible

ASTC compressed texture types give small texture sizes and save some download time. Where ASTC isn’t supported, consider ETC2. For more information, refer to [texture formats](texture-compression-formats.md).

To use ASTC compressed textures:

1. Select your texture asset in your project to open the **Texture Import Settings** window.
2. Select **Web settings**.
3. Enable **Override For Web**.
4. Set **Format** to one of the ASTC block options.

### Use larger block sizes for ASTC compressed textures

If you use ASTC compressed textures, try to use a larger block size for the compression. ASTC supports block sizes between 4x4 and 12x12 texels. Larger block sizes result in lower quality textures but smaller builds. For a good balance between quality and size that optimizes for download time over a mobile network, consider using 8x8. If download times are still too slow, try increasing the block size. For more information, refer to [Recommended, default, and supported texture formats, by platform](texture-choose-format-by-platform.md).

To set the block size of an ASTC compressed texture:

1. Select your texture asset in your project. The **Texture Import Settings** show.
2. Select **Web settings**.
3. Enable **Override For Web**.
4. Set **Format** to one of the larger ASTC block options, for example, **RGB(A) Compressed ASTC 8x8 block**.

## Lower the graphical quality level

The recommended best practice is to set the graphical quality level to the fastest option. The faster option results in a smaller build.

To change the graphical quality level:

1. Access the **Quality** settings (Menu: **Edit** > **Project Settings** > **Quality**).
2. Select the **Low** or **Very Low** quality level.

However, the fastest setting can impact the visuals of your application, so make sure that your application looks how you need it to look.

### Use C# to change the quality level

If you have a project settings script and want to change the quality level of your project in the script, add the following code:

```
// Set the quality level to Very Low (index 0)
QualitySettings.SetQualityLevel(0, true);
```

The `SetQualityLevel()` function takes the index of the **Quality Level** matrix as a value. `SetQualityLevel(0, true)` in this case is the **Very Low** setting, or the first option in the **Quality Level** matrix. To change it to **Low** or the second option, use `SetQualityLevel(1, true)` instead.

## Additional resources

* [Web](webgl.md)
* [Optimize your Web build](web-optimization.md)
* [Web browser compatibility](webgl-browsercompatibility.md)