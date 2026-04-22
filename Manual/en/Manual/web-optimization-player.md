# Recommended Player settings to optimize your Web build

Use the following recommended **Player** settings to optimize your builds for the Unity Web platform.

## Player settings quick reference

Find these settings under **Edit** > **Project settings** > **Player**. For more information on each setting, refer to the details in [Player settings](class-PlayerSettingsWebGL.md).

| **Setting** | **Recommended Setting** | **Description** |
| --- | --- | --- |
| [API Compatibility Level](#api-compatibility-level) | .NET Standard 2.1 | Produces smaller builds. |
| [IL2CPP Code Generation](#il2cpp-code-generation) | Optimize for code size and build time | Generates code optimized for build size and iteration. |
| [Managed Stripping Level](#managed-stripping-level) | High | Unity does a high level of managed stripping to create a smaller build. |

Configure the following recommended settings in the **Publishing Settings** section:

| **Setting** | **Recommended Setting** | **Description** |
| --- | --- | --- |
| [Compression Format](#compression-format) | Brotli | Files compressed by Brotli are smaller. |
| [Data Caching](#data-caching) | Enabled | Runs faster because cached data doesn’t need to be downloaded again on subsequent runs (unless the contents have changed). |
| [Debug Symbols](#debug-symbols) | Off | Debug symbols can slow down your application. |
| [Enable Exceptions](#enable-exceptions) | None | Enable **Target WebAssembly 2023** instead for exceptions without overhead. |
| [Target WebAssembly 2023](#wasm-2023) | Enable | The Web Assembly 2023 **feature set** contains features that help optimize your project. |
| [Maximum Memory Size (MB)](#maximum-memory-size) | Enter the amount of memory you need in MB. | Unity Web supports up to 4096 MB of memory for your build. |

## API Compatibility Level

Use the **API Compatibility Level** setting to choose which .NET APIs you can use in your project. The recommended setting is **.Net Standard 2.1** because this setting produces smaller builds and has cross-platform support. However, you need to check if your platform fully supports .Net Standard 2.1. For other options, refer to [Web player settings](class-PlayerSettingsWebGL.md).

### Change the API Compatibility Level

To change the **API Compatibility Level** via script instead, add this code to one of your **scripts**:

```
PlayerSettings.SetApiCompatibilityLevel(namedBuildTarget, ApiCompatibilityLevel.NET_2_0);
```

## IL2CPP Code generation

Use the ****IL2CPP** Code generation** setting to configure how Unity manages IL2CPP code generation (if your project uses the IL2CPP scripting back end).

The **Optimize for code size and build time** option is recommended because it creates a smaller build and generates less code which results in faster build times. Faster build times are vital in Web applications. However, this setting can reduce runtime performance.

For more information, refer to [IL2CPP overview](./scripting-backends-il2cpp.md).

### Change the IL2CPP Code generation via C

To enable this setting via script instead, add this code to one of your scripts:

```
PlayerSettings.SetIl2CppCodeGeneration(namedBuildTarget,         
                                       Il2CppCodeGeneration.OptimizeSize);
```

## Managed stripping level

Use **Managed stripping level** to configure how much the Unity linker process strips unused code from the managed DLLs your project uses. The recommended setting is **High** because stripping code can make your executable significantly smaller, which is important in Web applications. However, possible side effects might include:

* Managed code debugging of some methods might not work.
* You might need to maintain a custom `link.xml` file.
* Some reflection code paths might not behave the same.

For more information, refer to [Managed code stripping](managed-code-stripping.md).

### Change the Managed stripping level via C

To change this setting via script instead, add this code to one of your scripts:

```
PlayerSettings.SetManagedStrippingLevel(namedBuildTarget,     
                                        ManagedStrippingLevel.High);
```

## Compression Format

Choose what **compression** format to use for release build files. **Brotli** is the recommended setting because it has the best compression ratios and Brotli-compressed files are smaller than gzip. Smaller files are best for Web applications. However, Chrome and Firefox only support Brotli if the user accesses the website over HTTPS. Also, although most modern servers support Brotli, some servers don’t support it. Make sure that your server supports Brotli.

For more information on compression formats, refer to [Deploy a Web application](webgl-deploying.md).

### Change Compression Format via C# script

To enable this setting via script instead, add this code to one of your scripts:

```
PlayerSettings.WebGL.compressionFormat = WebGLCompressionFormat.Brotli;
```

## Data caching

Enable **Data caching** so that Unity caches your contents asset data on the user’s machine. Data caching can make the application run faster. For more information, refer to [Cache behavior in Web](webgl-caching.md).

### Enable data caching via C# script

To enable this setting via script instead, add this code to one of your scripts:

```
PlayerSettings.WebGL.dataCaching = true;
```

## Debug symbols

The **Debug symbols** setting preserves debug symbols and displays original function names of the stack trace when an error occurs, so it’s easier to identify the source of an error. The recommended best practice is to disable **Debug symbols** on your final release build because it can make your build unnecessarily large and slow down your application. Enable this setting during development and testing to make it easier to identify the source of an error.

### Deactivate Debug symbols via C# script

To deactivate this setting via script instead, add this code to one of your scripts:

```
PlayerSettings.WebGL.debugSymbolMode = WebGLDebugSymbolMode.Off;
```

## Enable exceptions

The **Enable exceptions** setting lets you choose how to handle errors at runtime. The recommended best practice is to choose the **None** option in your final release build because this setting gives the best performance and smallest builds.

However, with this option, any exception thrown causes your content to stop with an error, so it’s best to enable exceptions during development and testing.
For more information, refer to [Web Player settings](class-PlayerSettingsWebGL.md#wasm-language-features).

### Deactivate exceptions via C# script

To deactivate exceptions via script instead, add this code to one of your scripts:

```
PlayerSettings.WebGL.exceptionSupport = WebGLExceptionSupport.None;
```

## Target WebAssembly 2023

WebAssembly 2023 is a selection of WebAssembly language features that can improve performance and reduce code size. Therefore, it is recommended that you enable this setting.

### Enable WebAssembly 2023 via C# script

To enable this setting via script instead, add this code to one of your scripts:

```
PlayerSettings.WebGL.wasm2023 = true;
// PlayerSettings.WebGL.webAssemblyTable = true;
// PlayerSettings.WebGL.webAssemblyBigInt = true;
```

To enable the `WebAssembly.Table` or `BigInt` features without WebAssembly 2023, remove the first line of code and remove the comment lines.

## Maximum Memory Size (MB)

Unity supports up to 4096 MB of memory for your build. For most applications, 2048 MB of memory is enough. But for more complex real-time 3D experiences, especially for desktop browsers, it’s better to have more memory. With more memory, you can load in richer, larger **scenes** more dynamically.

However, there are known bugs for builds over 2048 MB in Firefox and Chrome browsers before version 119.

### Change the maximum memory via C# script

To change the maximum memory via script instead, add this code to your script and change the number to your preferred memory size:

```
PlayerSettings.WebGL.maximumMemorySize = 4096;
```

## Additional resources

* [Optimize your Web build](web-optimization.md)
* [Recommended Graphics settings to optimize your Web build](web-optimization-graphics.md)
* [Recommended Quality settings to optimize your Web build](web-optimization-quality.md)
* [Remove unused resources from your Web build](web-optimization-remove-resources.md)
* [Optimize Web platform for mobile](web-optimization-mobile.md)