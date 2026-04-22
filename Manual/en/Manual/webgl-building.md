# Web Build folder

When you build your Web application, Unity creates a `Build` folder with all files necessary to run your application in a web browser.

## Build folder structure

The `Build` folder has the name you specify in the ****Build Profiles**** window. The folder contains the following files, where `[ExampleBuild]` represents the name of the target build folder.

| File name | Description |
| --- | --- |
| `[ExampleBuild].loader.js` | The JavaScript code that the web page needs to load the Unity content. |
| `[ExampleBuild].framework.js` | JavaScript runtime and **plug-ins**. |
| `[ExampleBuild].wasm` | WebAssembly binary. |
| `[ExampleBuild].mem` | A binary image to initialize the heap memory for your Player. Unity generates this file for multithreaded WebAssembly builds only. |
| `[ExampleBuild].data` | Asset data and **Scenes**. |
| `[ExampleBuild].symbols.json` | Debug symbol names necessary to demangle an error stack trace. This file is only generated for Release builds when you enable the Debug Symbols option (**File** > **Build Profiles** > **Player Settings**.) |
| `[ExampleBuild].jpg` | A background image, which displays while the build is loading. This file is only generated when a Background Image is available in the Player Settings (**File** > **Build Profiles** > **Player Settings** > **Splash Image**). For more information, refer to [Splash Screen](class-PlayerSettingsSplashScreen.md). |

### File extensions

If you enable a ****Compression** Method** for your build, Unity identifies the extension that corresponds with the compression method and adds this extension to the names of the files inside the Build sub folder. If you enable **Decompression Fallback**, Unity appends the extension `.unityweb` to the build file names. Otherwise, Unity appends the extension `.gz` for the Gzip compression method, or `.br` for the Brotli compression method.

For more information, refer to [Compressed builds and server configuration](webgl-deploying.md).

### Name files as hashes

If you enable **Name Files As Hashes** in the ****Player Settings****, Unity uses the hash of the file content instead of the default file name. This applies to each file in the build folder. This option allows you to upload updated versions of the game builds into the same folder on the server, and only upload the files which have changed between build iterations.

**Note**: Opening a Player directly from the file system might not work in some browsers. This is due to security restrictions applied to local file URLs.

## Additional resources

* [Web Player settings](class-PlayerSettingsWebGL.md)
* [Web build settings](web-build-settings.md)
* [Build and distribute a Web application](webgl-building-distribution.md)
* [Deploy a Web application](webgl-deploying.md)
* [Web](webgl.md)