# Configure settings with the URP Config package

You can use the Universal **Render Pipeline** (URP) Config package to control some of the settings of URP. Unity automatically adds the package files in the package cache as they are a dependency of URP, but you must make a copy of them in your project before you can use the package.

The URP Config Package currently only changes one setting which is the maximum number of visible lights URP renders when you use the Forward+ **Rendering Path**. For more information, refer to [Change the maximum number of visible lights](rendering/forward-plus-rendering-path-limitations.md).

## Embed the URP Config package source code in your project

To prepare the URP Config package source code for modifications:

1. In the root folder of your project, go to `Library/PackageCache`.
2. Copy the following folder into the `Packages` folder at the root of your project:

   ```
   com.unity.render-pipelines.universal-config
   ```

The URP Config package is now ready for modification in your project.

**Note:** After you copy the package source code into the `Packages` folder, this package code becomes embedded in your project and is no longer part of a Unity install. If you upgrade the Unity version later, Unity doesn’t update the **embedded package** source automatically. You need to apply the changes to the source code manually. For more information about embedded packages, refer to [Embedded dependencies](../upm-embed.md).

## Configure URP with the URP Config package

You can edit the `ShaderConfig.cs` file to configure the properties of your URP project. If you edit this file, you must also update the equivalent `ShaderConfig.cs.hlsl` header file so that it mirrors the definitions you set in `ShaderConfig.cs`.

You can update the `ShaderConfig.cs.hlsl` file in two ways:

* Manually edit the `ShaderConfig.cs.hlsl` file to mirror the `ShaderConfig.cs` file. This method is faster but more likely to result in an error due to a mistake.
* Use the Editor to generate the `ShaderConfig.cs.hlsl` file from the `ShaderConfig.cs` file, which might take longer than a manual edit but ensures that the two files are synchronized.

To use the Editor to generate the `ShaderConfig.cs.hlsl` file, follow these steps:

1. In the **Project** window, go to **Packages** > **Universal RP Config** > **Runtime** and open **ShaderConfig.cs**.
2. Edit the values of the properties you want to change and then save and close the file.
3. In the Editor, select **Edit** > **Rendering** > **Generate Shader Includes**.
4. Unity automatically configures your project and **shaders** to use the new configuration.

### Update the URP Config package

When you use the Package Manager to update your URP package, the Package Manager downloads the latest version of the URP Config package to the `/Library/PackageCache/` folder, but doesn’t automatically update the files of the URP Config package in your `Packages` folder. Instead, you need to manually update your copy of the URP Config package in your `Packages` folder and reapply your changes. To do this, use the following steps:

1. Make a copy of the `com.unity.render-pipelines.universal-config` from your `Packages` folder. You can reference this later when you reapply your changes.
2. Delete the `com.unity.render-pipelines.universal-config` folder in your `Packages` folder.
3. Copy the `com.unity.render-pipelines.universal-config` folder again from the `/Library/PackageCache/` folder to your `Packages` folder, as shown above in the [Set up the URP Config package](#set-up-the-urp-config-package) section.
4. Manually reapply your modifications to the updated copy of the URP Config package.