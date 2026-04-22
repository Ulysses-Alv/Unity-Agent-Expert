# Set up a project for 2D games

**Note**: For this guide, Unity recommends and assumes that you choose the [Universal Render Pipeline (URP)](universal-render-pipeline.md) and not the [Built-in Render Pipeline](built-in-render-pipeline.md).

1. Install Unity version 2019 LTS or a later version; see [Installing Unity](GettingStartedInstallingUnity.md).
2. Create a new project with the [2D template](https://docs.unity3d.com/hub/manual/Templates.html).
3. In the Package Manager, install the latest [URP package](urp/urp-introduction.md) version; see [Installing the Universal Render Pipeline into an existing Project](urp/InstallURPIntoAProject.md).
4. Install any optional packages you need; see [Installing packages](#InstallingPackages).

## Installing packages

Most packages you need to make a 2D game in Unity are included in the Unity Editor. The following table lists the packages included by default when you choose the [2D template](https://docs.unity3d.com/hub/manual/Templates.html):

| **Package** | **Description** |
| --- | --- |
| [2D Animation](https://docs.unity3d.com/Packages/com.unity.2d.animation@latest) | 2D Animation provides the necessary tooling and runtime components for applying skeletal animation to your **Sprites**. |
| [2D Pixel Perfect](urp/2d-pixelperfect.md) | The 2D **Pixel** Perfect package contains the Pixel Perfect **Camera** component, which ensures your pixel art remains crisp and clear at different resolutions, and stable in motion. A standalone [2D Pixel Perfect package](https://docs.unity3d.com/Packages/com.unity.2d.pixel-perfect@latest) that doesn’t require the URP is available if you want to use this package in a project using the [Built-in Render Pipeline](built-in-render-pipeline.md) instead. |
| [2D PSD Importer](https://docs.unity3d.com/Packages/com.unity.2d.psdimporter@latest) | The 2D PSD Importer package allows you to import multilayered PSD files from Photoshop. You can use this for your Sprites, or to rig your characters. |
| [2D Sprite](sprite/sprite-landing.md) | The Sprite Editor provides an in-Editor environment to create and edit Sprite assets. Sprite Editor lets you add custom behavior for editing Sprite-related data. |
| [2D SpriteShape](https://docs.unity3d.com/Packages/com.unity.2d.spriteshape@latest) | 2D Sprite Shape allows you to create organic shapes and **terrains**, similar to a vector drawing tool. For example, you can choose the fill texture and border Sprites. |
| [2D Tilemap Editor](tilemaps/work-with-tilemaps/tilemap-reference.md) | 2D **Tilemap** Editor allows you to create grid-based worlds with square, hexagonal or isometric tiles. Add your Tiles to the Tile Palette, and paint and fill Tile Grids using different settings and brushes. Extra tools let you add smart drawing, randomization or animation to the Tile assets. |

The following table lists some optional packages you can install that might be particularly useful for 2D game development:

| **Package** | **Description** |
| --- | --- |
| [Shader Graph](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest) | **Shader** Graph lets you build your shaders visually. |
| [Cinemachine](https://docs.unity3d.com/Packages/com.unity.cinemachine@latest) | The Cinemachine package is a suite of modules that provide advanced functionality for operating the Unity Camera. |
| [2D Tilemap Extras](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest) | The 2D Tilemap Extras package contains reusable 2D and Tilemap Editor **scripts** that you can use for your own Projects. |