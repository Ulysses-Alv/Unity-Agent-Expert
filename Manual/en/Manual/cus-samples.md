# Create samples for your package

You can create samples to include in a Unity Package Manager (UPM) package you develop.

Although optional, including samples helps users learn how to use your package. A sample might be a piece of example code, some **shaders** and textures, some animation, or any other files that you typically find under the project’s `Assets` folder.

When you open the Package Manager window and select a package containing samples, an **Import** button appears in the package’s [details panel](upm-ui-details.md) for each sample in the package. When you select **Import**, the Package Manager copies the whole subfolder structure for that sample under the project’s `Assets` folder.

To add samples to your package:

1. Put the asset files or C# code files [under the Samples folder](#sample-subfolder). You can have more than one sample in a package. Each subfolder of the `Samples` folder has one sample.
2. Open the **package manifest** file (`package.json`) for editing. To locate the file, refer to [Locate the manifest file](cus-edit-manifest.md#locate-manifest).
3. If the file doesn’t contain a JSON array called `samples`, add it.
4. For each sample you want to include, add a JSON object in the `samples` array.
5. Include the following keys and values for each JSON object in the `samples` array:

   | **Key** | **Description** |
   | --- | --- |
   | `displayName` | The name of the sample as it appears in the package details panel of the Package Manager window. |
   | `description` (optional) | A brief description of what the sample demonstrates or contains. This description displays in the **Samples** tab of the Package Manager window’s package details panel. |
   | `path` | The path to the sample’s folder, starting with `Samples`. |
6. Make sure that `package.json` contains valid JSON. You can check its validity using an online JSON validator or directly in your code editor, if it supports JSON syntax checking.
7. Save the file.

## Location of sample files

You can add your sample assets under subfolders of the `Samples` folder of your package. For example, a package with shader samples might look something like this:

```
MyPackage
  ├── package.json
  └── Samples
        ├── SamplesHDRP
        │    ├── Textures
        │    |     ├── MossyRock.bmp
        │    |     └── SandyRock.bmp
        │    └── Shader
        │          ├── Lit Texture Blend HDRP.ShaderGraph
        │          └── Lit Vertex Color HDRP.ShaderGraph
        └── SamplesStandard
        │    ├── Textures
        │    |     ├── MossyRock.bmp
        │    |     └── SandyRock.bmp
        │    └── Shader
        │          ├── StandardTextureBlend.shader
        │          └── StandardVertexColor.shader
        └── SamplesUniversalRP
             ├── Textures
             |     ├── MossyRock.bmp
             |     └── SandyRock.bmp
             └── Shader
                   ├── Lit Texture Blend URP.ShaderGraph
                   └── Lit Vertex Color URP.ShaderGraph
```

## Example of a samples array in the manifest file

Using the same structure as the example for [Location of sample files](#sample-subfolder), the `samples` array in `package.json` looks similar to this:

```
{
    "samples": [
        {
            "displayName": "HDRP Shaders",
            "description": "Contains sample shaders for the High Definition render pipeline",
            "path": "Samples/SamplesHDRP"
        },
        {
            "displayName": "Standard RP Shaders",
            "description": "Contains sample shaders for the Standard render pipeline",
            "path": "Samples/SamplesStandard"
        },
        {
            "displayName": "URP Shaders",
            "description": "Contains sample shaders for the Universal render pipeline",
            "path": "Samples/SamplesUniversalRP"
        }
    ]
}
```

## Samples folder naming convention

When creating your package with the **Create package** function, keep the samples path in your `package.json` pointing to `Samples`. Don’t append a trailing tilde (`~`). During the [export process](cus-export.md), Unity will rename the `Samples` folder to `Samples~` automatically so that it’s hidden in the `Packages` folder of the **Project** window. You don’t need to adjust the path values for samples in your manifest file. The Package Manager will resolve the sample path correctly. After developers import your sample, they can find it in the `Assets` folder of the **Project** window.

## Additional resources

* [Package creation](cus-pkg-lp.md)
* [Package development workflow](CustomPackages.md)
* [Edit the package manifest](cus-edit-manifest.md)