# Add code to your package

Continue the development of your Unity Package Manager (UPM) package by adding the assets (including code) that make your package unique.

Before you begin this phase of development, consider how others will use your **UPM package**. Determine whether the package falls into one or both of the following categories:

* Code for tools that helps others as they develop a project. This code typically placed in the package’s `Editor` subfolder. Final builds of the project don’t include these packages. Examples include:
  + A custom debugger
  + A memory profiler
  + A level design tool for painting spawn points
* Code that runs during gameplay, whether that’s in the Editor’s Play mode or a built project. This code is typically placed in the package’s `Runtime` subfolder. Final builds of the project include these packages. Examples include:
  + A dialog system for branching conversations
  + A procedural generation engine
  + An input remapping system that let players customize controls at runtime

Follow these steps as you iterate on the development of your package:

1. Add Editor code to your package’s `Editor` subfolder and runtime code to the `Runtime` subfolder.

   **Note**: If your package also contains non-code assets, include them in the relevant folder.
2. Optional: You can organize your content in a more complex folder structure. However, complex structures might require you to update the package’s [assembly definition files](cus-asmdef.md).
3. Check the package’s [assembly definition files](cus-asmdef.md) to make sure they’re correct. They usually don’t need updating, unless you change the location of your **scripts**.
4. [Edit the package manifest](cus-edit-manifest.md), making sure all [required and recommended fields](upm-manifestPkg.md#required) have values. Consider supplying values for the relevant [optional fields](upm-manifestPkg.md#optional).
5. Test your package.

## Next steps

When your package is functioning as expected, you can proceed with other activities in the [package development workflow](CustomPackages.md). Typical activities at this stage of package development include:

* [Create or edit the assembly definition](cus-asmdef.md)
* [Add tests to your package](cus-tests.md)
* [Create samples for your package](cus-samples.md)

## Additional resources

* [Package creation](cus-pkg-lp.md)
* [Package development workflow](CustomPackages.md)