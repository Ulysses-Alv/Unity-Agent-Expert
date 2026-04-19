# Modify URP source code

For advanced customizations, it might not be enough to extend or override APIs defined in URP. You might need to modify URP source code to change certain methods.

To prepare URP source code for modifications:

1. In the root folder of your project, go to `Library/PackageCache`.
2. Copy the following folders into the `Packages` folder at the root of your project:

   ```
   com.unity.render-pipelines.universal
   com.unity.render-pipelines.universal-config
   ```

Now Unity uses the URP source code from the `Packages` folder of your project, and you can make changes to that code.

**Note:** After you copy the package source code into the `Packages` folder, this package code becomes embedded in your project and is no longer part of a Unity install. If you upgrade the Unity version later, Unity doesn’t update the **embedded package** source automatically. You need to apply the changes to the source code manually. For more information about embedded packages, refer to [Embedded dependencies](../../upm-embed.md).

## Document changes to embedded package

To make it easier to upgrade to future Unity versions, document all the changes you make to the embedded URP package code.

## Additional resources

* [Custom lighting in URP](../lighting/custom-lighting-landing.md)