# Introduction to textures

Normally, the **mesh** geometry of an object only gives a rough approximation of the shape while most of the fine detail is supplied by **Textures**. A texture is just a standard bitmap image that is applied over the mesh surface. You can think of a texture image as though it were printed on a rubber sheet that is stretched and pinned onto the mesh at appropriate positions. The positioning of the texture is done with the 3D modelling software that is used to create the mesh.

![Cylinder with tree bark](../uploads/Main/CylinderTreeBark.png)

Cylinder with tree bark

Unity can import textures from most common image file formats.

## Terminology

This page uses the following terminology:

* **Bits per **pixel** (bpp)** is the amount of storage required for a single texture pixel. Textures with a lower bpp value have a smaller size on disk and in memory. A lower bpp value also means that the GPU can store more pixels in its cache, which results in faster texture access.
* LDR (Low Dynamic Range) refers to most typical images where colors are conceptually between 0.0 (black) and 1.0 (white) values. The majority of image files (such as PNG and JPG) have low dynamic range.
* **HDR** (High Dynamic Range) refers to special image and **texture formats** where colors can have a higher range than 0 through 1. Image file formats like .exr or .hdr are often used for HDR image data. At runtime and on the GPU, there are several HDR formats, trading off accuracy, range and memory usage.
* **RGB** is a color model in which red, green and blue combine to reproduce an array of colors.
* **RGBA** is a version of **RGB** with an alpha channel, which supports blending and opacity alteration.
* **Variable bit rate (VBR)** means that bits per pixel is not a fixed value, and depends on the actual content instead. VBR only applies to **Crunch **compression****, and only texture size on disk. The size in memory is the same as when using the underlying texture format (for example, RGB Compressed DXT1 for RGB Crunched DXT1).

## Textures for use on 3D models

Textures are applied to objects using [Materials](Materials.md). Materials use specialised graphics programs called [Shaders](Shaders.md) to render a texture on the mesh surface. Shaders can implement lighting and colouring effects to simulate shiny or bumpy surfaces among many other things. They can also make use of two or more textures at a time, combining them for even greater flexibility.

You should make your textures in dimensions that are to the power of two (e.g. 32x32, 64x64, 128x128, 256x256, etc.) Simply placing them in your project’s Assets folder is sufficient, and they will appear in the Project View.

Once your texture has been imported, you should assign it to a [Material](class-Material.md). The material can then be applied to a mesh, **Particle System**, or **GUI Texture**. Using the **Import Settings**, it can also be converted to a **Cubemap** or **Normalmap** for different types of applications in the game. For more information about importing textures, please read the [Texture Component page](class-TextureImporter.md).

## 2D graphics

In 2D games, the **Sprites** are implemented using textures applied to flat meshes that approximate the objects’ shapes.

![Sprite from a 3D viewpoint](../uploads/Main/SpriteFrom3DViewPt.png)

Sprite from a 3D viewpoint

An object in a 2D game may require a set of related graphic images to represent animation frames or different states of a character. Special techniques are available to allow these sets of images to be designed and rendered efficiently. For more information, refer to [Create sprites from a texture](sprite/sprite-editor/use-editor.md).

## GUI

A game’s *graphic user interface* (GUI) consists of graphics that are not used directly in the game **scene** but are there to allow the player to make choices and see information. For example, the score display and the options menu are typical examples of game GUI. These graphics are clearly very different from the kind used to detail a mesh surface but they are handled using standard Unity textures nevertheless. See the manual chapter on [GUI Scripting Guide](GUIScriptingGuide.md) for further details about Unity’s GUI system.

## Particles

Meshes are ideal for representing solid objects but less suited for things like flames, smoke and sparkles left by a magic spell. This type of effect is handled much better by **Particle Systems**. A *particle* is a small 2D graphic representing a small portion of something that is basically fluid or gaseous, such as a smoke cloud. When many of these **particles** are created at once and set in motion, optionally with random variations, they can create a very convincing effect. For example, you might display an explosion by sending particles with a fire texture out at great speed from a central point. A waterfall could be simulated by accelerating water particles downward from a line high in the scene.

![Star particle system](../uploads/Main/ParticleSystemCone.png)

Star particle system

Unity’s particle systems have a wealth of options for creating all kinds of fluid effects. See the [manual chapter](ParticleSystems.md) on the subject for further information.

## Anisotropic filtering

Anisotropic filtering increases Texture quality when viewed from a grazing angle. This rendering is resource-intensive on the graphics card. Increasing the level of anisotropy is usually a good idea for ground and floor Textures. Use [Quality](class-QualitySettings.md) settings to force anisotropic filtering for all Textures or disable it completely. Although, if a texture has its ****Aniso level**** set to 0 in [Texture Import Settings](class-TextureImporter.md), forced anisotropic filtering does not appear on this texture.

![Anisotropy used on the ground Texture {No anisotropy (left) | Maximum anisotropy (right)} ](../uploads/Main/ImportingTextures-AnisoFiltering.png)

## Additional resources

* [Editor GUI and Legacy GUI texture Import Settings window reference](texture-type-editor-gui-and-legacy-gui.md)
* [Sprite (2D and UI) texture Import Settings window reference](texture-type-sprite.md)