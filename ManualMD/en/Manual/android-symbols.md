# Android symbols

To help you debug your application, Unity can generate a package that contains symbol files with debug metadata for native Unity libraries. Symbol files contain an executable file section called symbol table (`.symtab`) that translates active memory addresses into information you can use, like a method name. The translation process is called symbolication. You can upload a symbols package to the Google Play Console to see a human-readable stack trace on the [Android Vitals](https://developer.android.com/topic/performance/vitals) dashboard.

For more information about executable file (ELF) section, refer to [Wikipedia](https://en.wikipedia.org/wiki/Executable_and_Linkable_Format).

Unity generates symbol files for the following libraries:

* `libmain`: Responsible for the initial Unity engine loading logic. The symbol file is precompiled.
* `libunity`: Unity’s engine code:

  + If **Strip Engine Code** property is disabled, the symbol file is precompiled.
  + If **Strip Engine Code** property is enabled, Unity compiles the symbol file during build process.
* `libil2cpp`: Contains C# **scripts**:

  + If you don’t export your project, Unity compiles the symbol file during build process.
  + If you export your project, **Gradle** compiles the symbol file.

Gradle generates symbol files for any additional shared libraries by stripping the `.symtab` and `.debug*` executable file sections from the shared libraries. If a shared library doesn’t contain these sections, Gradle will not generate a symbol file for it.

Depending upon your application build format, Unity generates symbol files in two ways :

* As a zip file that can be generated for both `apk` or `aab`.
* Directly embeds symbol files in the `aab`. Unity doesn’t embed symbol files into an `apk`.

Use [UserBuildSettings.DebugSymbols.format](../ScriptReference/Android.UserBuildSettings.DebugSymbols-format.md) API to set the format of the symbol package.

## Types of symbol files

Unity uses [debugSymbolLevel](https://developer.android.com/reference/tools/gradle-api/7.3/com/android/build/api/dsl/Ndk#debugSymbolLevel()) property of Gradle to generate symbol files. There are two types of symbol files:

* Public: A small file that contains a symbol table section. For more information, refer to [Public symbols](#public-symbols).
* Debugging: Contains everything that a public symbol file contains along with full debugging information that you can use for more in-depth debugging. For more information, refer to [Debugging symbols](#debugging-symbols).

Use [UserBuildSettings.DebugSymbols.level](../ScriptReference/Android.UserBuildSettings.DebugSymbols-level.md) to generate the required type of symbol file.

**Note:** By default, Gradle generates symbol files with `.so.sym` or `.so.dbg` extensions. Some digital distribution services don’t recognize these extensions and require files with `.so` extension. In such cases, use [Unity.Android.Types.DebugSymbolFormat.LegacyExtensions](../ScriptReference/Unity.Android.Types.DebugSymbolFormat.LegacyExtensions.md) API to create symbol files with `.so` extension.

```
UserBuildSettings.DebugSymbols.format = DebugSymbolFormat.IncludeInBundle | DebugSymbolFormat.Zip | DebugSymbolFormat.LegacyExtensions;
```

### Public symbols

A public symbols file contains information that resolves function addresses to human-readable strings. These files don’t contain debugging information. This makes public symbol files smaller than [debugging symbols](#debugging-symbols) files.

### Debugging symbols

A debugging symbols file contains full debugging information and a symbol table section. Use it to:

* Resolve stack traces and to debug applications that you have source code available for.
* Attach a native debugger to the application and debug the code.

**Note:** If debugging symbols aren’t available, Unity places a [public symbols](#public-symbols) file in your project at build time. For the `libmain` and `libunity` libraries, debugging symbols aren’t available and Unity always generates public files.

### Custom symbols

You can instruct Unity to include additional symbol files. This is useful if you use shared libraries and want your local debugger, and Google Play, to resolve the shared library stack traces if the application crashes.

To make Unity include a custom symbols file:

1. In the **Project window**, select a **plug-in** that has a `.so` file extension.
2. In the **Inspector**, find the **Platform settings** section.
3. Set **CPU** to the CPU architecture that the symbols file is compatible with.
4. Set **Shared Library Type** to **Symbol**.

Whenever Unity generates a symbols package, it adds the additional symbol files to the symbols package.

**Important:** If you set the **Shared Library Type** to **Executable** and the shared library contains the `.symtab` and `.debug*` executable file sections, Gradle generates a symbol file that overwrites your custom symbol file.

If you want to make Unity include a custom symbols file from a C# script, the **UnityEditor.Android** namespace includes the following APIs to set the **CPU** and **Shared Library Type** respectively:

* [PluginImporter.SetAndroidCPU](../ScriptReference/Android.AndroidPluginImporterUtilities.SetAndroidCPU.md)
* [PluginImporter.SetAndroidSharedLibraryType](../ScriptReference/Android.AndroidPluginImporterUtilities.SetAndroidSharedLibraryType.md)

**Note:** The symbols file name must match the name of the shared library that the symbols file is for. For example, if a shared library is called **mylibrary.so**, the symbols file must also be named **mylibrary.so**. To avoid file name **collisions**, the symbols file and the shared library must be in separate directories.

**Important:** Ensure the symbols file is up to date and compatible with the shared library that contains the executable code. If you don’t, your local debugger and Google Play will fail to resolve stack traces for code in the shared library.

![Custom plug-in in the Inspector.](../uploads/Android/android_custom_symbol.png)

Custom plug-in in the Inspector.

## Generating a symbols package

There are two ways to enable symbols package generation for your application:

* Through the ****Build Profiles**** window.
* Using the [UserBuildSettings.DebugSymbols.level](../ScriptReference/Android.UserBuildSettings.DebugSymbols-level.md) and [UserBuildSettings.DebugSymbols.format](../ScriptReference/Android.UserBuildSettings.DebugSymbols-format.md) APIs.

To enable symbols package generation through the **Build Profiles** window:

1. Open the **Build Profiles** window (menu: **File** > **Build Profiles**).
2. From the list of platforms in the **Platforms** panel, select **Android** or [create a build profile](create-build-profile.md) for the **Android** platform.
3. Set **Debug Symbols** to one of the following:

   * [Public](#public-symbols)
   * [Debugging](#debugging-symbols)
4. Set **Symbols output options** to **.zip**.

After you enable symbols package generation, building your project generates a `.zip` file that contains symbol files for the `libmain` and `libunity` library. If you set your [scripting backend](scripting-backends.md) to **IL2CPP**, the `.zip` also contains a symbol file for the `libil2cpp` library. Unity places this symbols package within the output directory. The name of the symbols package file is in this format: `<appbuildfilename>-<version>-v<bundle version code>-IL2CPP.symbols.zip`. For example, `MyApp-0.1-v100-IL2CPP.symbols.zip` where `MyApp` is the name of your application build file, `0.1` is the version, and `100` is the bundle version code. The version and bundle version code are set in the [Android Player settings](class-PlayerSettingsAndroid.md#Identification).

If you enable **Export Project** in the Android build settings, Unity doesn’t build the project. Instead, it exports the project for Android Studio, generates symbols for `libmain` and `libunity`, and places them within `unityLibrary/symbols/<architecture>/` in the output directory. When you build your exported project from Android Studio, Gradle generates the `libil2cpp` symbol file and places it within the `unityLibrary/symbols/<architecture>/` directory alongside the `libmain` and `libunity` symbol file.

## Using symbols in the Google Play console

### Embedded symbols

If you’re producing an Android App bundle (aab), you can embed symbols directly into an `aab` and upload it to Google Play.

**Note:** Unity doesn’t embed symbols into an `apk`. In this case, you must upload a zip file with symbols separately.

```
UserBuildSettings.DebugSymbols.level = DebugSymbolLevel.SymbolTable;
UserBuildSettings.DebugSymbols.format = DebugSymbolFormat.IncludeInBundle;
```

### Zipped symbols

To produce a zipped symbols package, use the following code:

```
UserBuildSettings.DebugSymbols.level = DebugSymbolLevel.SymbolTable;
UserBuildSettings.DebugSymbols.format = DebugSymbolFormat.Zip | DebugSymbolFormat.LegacyExtensions;
```

After you upload your application to Google Play, you can upload a [public symbols](#public-symbols) zip package. For information on how to do this, refer to Google’s documentation: [Deobfuscate or symbolicate crash stack traces](https://support.google.com/googleplay/android-developer/answer/9848633).

**Note:** Google Play doesn’t symbolicate crashes that your application received before you uploaded the symbols package.