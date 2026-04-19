# Distribution size and code stripping

When publishing for Web, it is important to keep your build size low so users get reasonable download times before the content starts. For generic tips on reducing asset sizes, see documentation on [Reducing the file size of the build](ReducingFilesize.md).

## Hints and tips specific to Web

* Specify the **Crunch** texture **compression** format for all your compressed textures in the [Texture Importer](class-TextureImporter.md).
* Don’t deploy **development builds**; they’re not compressed or [minified](https://en.wikipedia.org/wiki/Minification_%28programming%29), and so have much larger file sizes.
* In the **Player** settings window, (click **Edit** > **Project Settings** > **Web**) expand **Publishing Settings**, and set **Enable Exceptions** to **None** if you don’t need exceptions in your build.
* Enable **Strip Engine Code** in the **Player** settings > **Other Settings** panel, to ensure an efficient build.
* When using third-party managed dlls, be aware that it might come with dependencies that increase the generated code size.

If you make a release build, Unity compresses the build output files according to the **Compression Format** selected in the **Publishing Settings** panel of the Web **Player** settings.

For more information on how to publish compressed builds, see [Deploying compressed builds](webgl-deploying.md).

## Code stripping

Unity removes all unused code from your build by default. You can change this via the **Player** settings (menu: **Edit** > **Project Settings**, then select the **Player** category): Select the **Other Settings** panel to access the **Strip Engine Code** option. It is better to build with stripping enabled.

With code stripping, Unity scans your project for any `UnityObject`-derived classes used (either by being referenced in your script code, or in the serialized data in your Scenes). It then removes from the build any Unity subsystems which have none of their classes used. This makes your build have less code, resulting in both smaller downloads and less code to parse (so code runs faster and uses less memory).

### Issues with code stripping

Code stripping might cause issues with your project if it strips code which is actually necessary. This can be the case when you load AssetBundles at run time which contain classes that are not included in the main build, and have therefore been stripped from the project. Error messages appear in your browser’s JavaScript console when this happens (possibly followed by more errors). For example:

`Could not produce class with ID XXX`

To troubleshoot these errors, look up the ID (such as `XXX` in the example above) in the [Class ID Reference](ClassIDReference.md) to see which class it is trying to create an instance of. In such cases, you can force Unity to include the code for that class in the build, either by adding a reference to that class to your **scripts** or to your **Scenes**, or by adding a `link.xml` file to your project.

Below is an example which makes sure that the **Collider** class and the Physics module gets preserved in a project. Add this XML code to a file called `link.xml`, and put that file into your *Assets* folder.

```
<linker>
    <assembly fullname="UnityEngine">
        <type fullname="UnityEngine.Collider" preserve="all"/>
    </assembly>
</linker>
```

If you suspect that stripping is causing problems with your build, you can also try disabling the **Strip Engine Code** option during testing.

Unity does not provide a convenient way to see which modules and classes are included in a build, which would allow you to optimize your project to strip well. However, to get an overview of included classes and modules, you can look at the generated file `Temp/StagingArea/Data/il2cppOutput/UnityClassRegistration.cpp` after making a build.

Note that the **Strip Engine Code** option only affects Unity engine code. **IL2CPP** always strips byte code from your managed dlls and scripts. This can cause problems when you need to reference managed types dynamically through reflection rather than through static references in your code. If you need to access types through reflection, you may also need to set up a `link.xml` file to preserve those types. See the documentation page on [iOS Build size optimization](iphone-playerSizeOptimization.md) for more information on `link.xml` files.

## Moving build output files

To change the location of your `Build` folder, modify the buildUrl variable in the Web Template index.html file.

To change the location of the files inside the `Build` folder, change their URLs (that is, `dataUrl`, `wasmCodeUrl`, `wasmMemoryUrl`, and `wasmFrameworkUrl`) in the parameters to the config variable in the index.html file.

You can specify URLs on external servers for these if you want to host your files on a content distribution network (CDN), but you need to make sure that the hosting server has enabled Cross Origin Resource Sharing (CORS) for this to work. See the manual page on [Web networking](webgl-networking.md) for more information about CORS.

## Incremental builds

The C++ code generated for your project by IL2CPP is compiled incrementally; that is, only generated C++ code that has changed since the last build is compiled again. Unchanged source code re-uses the same object files generated for the previous build. The object files used for incremental C++ builds are stored in the `Library\Bee\artifacts\WebGL` directory in your Unity project.

To perform a clean, from-scratch build of the generated C++ code which doesn’t use incremental compiling, delete the `Library\Bee\artifacts\WebGL` directory in your Unity project. Note that if the Unity Editor version differs from the one used for the previous Web build, Unity does a clean, from-scratch build automatically.