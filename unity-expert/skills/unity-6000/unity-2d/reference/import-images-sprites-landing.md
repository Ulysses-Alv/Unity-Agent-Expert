# Import a sprite or spritesheet texture

Import an image and use it as a **sprite** or set of sprites in a 2D **scene**.

To use an imported image as one or more sprites, set its texture type to **Sprite (2D and UI)**. Follow these steps:

1. Make sure the 2D Sprite package is installed in your project. For more information, refer to [set up your project for 2D games](../../setup-project-2d-game.md).
2. Import your sprite image or spritesheet. For example, drag a spritesheet .png file into the **Project** window.
3. Select the imported image in the **Project** window.
4. In the ****Inspector**** window, set **Texture Type** to **Sprite (2D and UI)**.
5. Select **Apply**.

**Note:** In a 2D project, Unity automatically sets **Texture Type** to **Sprite (2D and UI)** when you import an image.

## Divide the texture into sprites

The **Sprite Mode** property of the texture determines whether Unity creates one sprite or multiple sprites from the texture.

For each sprite, Unity creates a sprite asset () as a child of the texture in the **Project** window.

Set **Sprite Mode** to the following:

* **Single** to use the texture as one sprite, either using the whole image or part of the image.
* **Multiple** to divide the texture into multiple sprites, for example if the texture is a spritesheet.

To set the shape and size of sprites, refer to [Create sprites from a texture](sprite/sprite-editor/use-editor.md).

## Add a sprite to a scene

To add a sprite to a scene, drag the sprite asset () from the **Project** window into the **Scene** window or **Hierarchy** window.

Unity creates a **GameObject** with a **Sprite Renderer** component. The component renders the sprite and controls how it appears in the scene.

### Add a sprite to an existing GameObject

To add a sprite to an existing GameObject, follow these steps:

1. Select the GameObject in the **Scene** or **Hierarchy** window.
2. In the **Inspector** window, select **Add Component**.
3. Select **Rendering** > **Sprite Renderer**.
4. In the **Sprite Renderer** component, set the **Sprite** property to the sprite you want to render.

## Additional resources

* [Create sprites from a texture](sprite/sprite-editor/use-editor.md)
* [Introduction to importing assets](../../ImportingAssets.md)
* [Sprite (2D and UI) import settings reference](../../texture-type-sprite.md)
* [Texture Import Settings](../../class-TextureImporter.md)