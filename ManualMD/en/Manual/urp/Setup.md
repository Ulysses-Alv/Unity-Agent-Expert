# Set up the 2D Renderer asset in URP

To use the Universal **Render Pipeline** (URP) for 2D in your Unity project, create a project with a 2D URP asset and a 2D Universal Renderer asset.

## Requirements

* Unity Editor version 2021.2.0b1 or higher.
* URP package version 10 or higher.

## Create 2D URP rendering assets

1. Open the Unity Hub.
2. Select the **Univeral 2D** project, then **Create project**.
3. When the new project opens, in the **Hierarchy** window, open the context menu (right-click) and select **Create** > **Rendering** > **URP Asset (with 2D Renderer)**.
4. Enter a name for the new assets.

URP creates a new [URP Asset](universalrp-asset.md), and assigns it a new [Universal Renderer](urp-universal-renderer.md) asset with `_Renderer` appended to the name.

2D rendering is now set up for your project.

**Note**: Some of the settings in the URP asset that relate to 3D rendering won’t affect your final application.

## Additional resources

* [Graphics quality settings in URP](urp-quality-settings-landing.md)