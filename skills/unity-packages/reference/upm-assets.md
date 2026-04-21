# Access package assets using scripts

This section explains how to access or refer to assets that are defined inside a package:

* [Referring to package paths](#Paths)
* [Loading a Texture inside a package](#Textures)
* [Resolving absolute paths](#GetFullPath)

**Note**: Package Manager doesn’t support streaming assets in packages. Use the [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest) package instead.

## Referring to package paths

To refer to assets that are defined inside a package, use this path scheme:

```
"Packages/<package-name>/..."
```

The path of the asset inside a package begins with `Packages/` and the package [name](upm-manifestPkg.md#name) (not the [display name](upm-manifestPkg.md#displayName)).

By contrast, you access project assets using this scheme:

```
"Assets/..."
```

For example, the path for the file `image.png` in the package subfolder `/Example/Images` of the **com.unity.images-library** package is:

```
"Packages/com.unity.images-library/Example/Images/image.png"
```

To get the absolute path of an item in your `Packages` folder, you can use the partial path as a parameter to the [Path.GetFullPath()](https://docs.microsoft.com/en-us/dotnet/api/system.io.path.getfullpath?view=netframework-4.7.2) method. For an example, refer to [Resolving absolute paths](#GetFullPath).

## Loading a Texture inside a package

To load a Texture stored inside a package, use the [LoadAssetAtPath](../ScriptReference/AssetDatabase.LoadAssetAtPath.md) method, which requires the `using UnityEditor` directive. Specify the path following the `Packages/<package-name>/` path scheme as demonstrated in this example:

```
using UnityEditor;
// ...
Texture2D texture = (Texture2D)AssetDatabase.LoadAssetAtPath("Packages/com.unity.images-library/Example/Images/image.png", typeof(Texture2D));
```

## Resolving absolute paths

To get the absolute path of a packaged asset, use the [Path.GetFullPath()](https://docs.microsoft.com/en-us/dotnet/api/system.io.path.getfullpath?view=netframework-4.7.2) method, which requires the `using System.IO` directive. For example:

```
using System.IO;
// ...
string absolute =   Path.GetFullPath("Packages/com.unity.images-library/Example/Images/image.png");
```

## Additional resources

* [PackageInfo.name](../ScriptReference/PackageManager.PackageInfo-name.md) (Scripting API)
* [Search for files](search-files.md)