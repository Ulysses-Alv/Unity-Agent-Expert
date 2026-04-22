# Create a sprite atlas

Create a [sprite atlas](atlas-introduction.md) and pack **sprites** into it.

**Note:** You can also create multiple sprites from a single texture instead. For more information, refer to [Import a sprite or spritesheet texture](../import-images-sprites/import-images-sprites-landing.md) and [Cut out sprites from a texture](../sprite-editor/use-editor.md).

## Prepare sprites for packing

To optimize the performance of sprite atlases and the memory they use, do the following for each texture you want to pack:

* Disable **Read/Write** unless you read or write to the texture in C# **scripts**.
* Enable **Tight Packing** to reduce the number of transparent **pixels** around the sprite.

Select **Open Sprite Editor** and check the following:

* The **Custom Outline** tab, to make sure the outline fits closely to the sprite. For more information, refer to [Crop a sprite](../sprite-editor/generate-outline.md).
* The **Secondary Textures** tab, to make sure the sprite has the same number of secondary textures as the other sprites you want to pack it with. Otherwise the combined secondary textures in the sprite atlas might contain a lot of empty space. For more information, refer to [Add normal map and mask map textures to a sprite](../../urp/SecondaryTextures.md).

## Create a sprite atlas

To create a sprite atlas, from the main menu, go to **Assets** > **Create** > **2D** > **Sprite Atlas**.

Unity creates a file with a `.spriteatlasv2` file extension in the **Project** window. Unity no longer uses Sprite Atlas V1 format by default. For more information, refer to [Convert Sprite Atlas V1 assets](v2/sprite-atlas-v2.md).

## Add sprites to a sprite atlas

For more information about assigning sprites to different sprite atlases, refer to [Sprite atlases](atlas-introduction.md).

1. In the **Project** window, select the sprite atlas asset.
2. To add a sprite to the sprite atlas, drag a sprite from the **Project** window onto the **Objects for Packing** label in the ****Inspector**** window.

   You can drag sprites, textures, or folders onto the **Objects for Packing** label. You can also select the **Add** (**+**) button in the **Objects for Packing** section.
3. Select **Pack Preview** to display the packed texture in the **Preview** window at the bottom of the **Inspector** window.

If your textures have **normal maps** or mask maps as secondary textures, select the dropdown at the top-right of the **Preview** window to display the combined secondary textures.

Sprites now render using their sprite atlas texture in the **Scene** view, the Game view, and in Play mode. Unity also includes sprite atlases in your builds by default, and automatically loads and uses them at runtime.

**Note:** If you add a sprite to multiple sprite atlases, Unity randomly chooses which sprite atlas texture to use at runtime. To avoid this, refer to [Load sprite atlases manually at runtime](distribution/load-sprite-atlas-spriteatlasmanageratlasrequested.md).

## Use the original textures while editing

To use the original unpacked textures in the **Scene view** and Game view, but keep using the sprite atlas textures in Play mode and at runtime, follow these steps:

1. From the main menu, select **Edit** > **Project Settings**.
2. Select **Editor**.
3. In the **Sprite Atlas** section, set **Mode** to **Sprite Atlas V2 - Enabled for Builds**. Unity now builds and uses sprite atlases only when you enter Play mode or build your project.

For information on other **Mode** values, refer to [Editor settings](../../class-EditorManager.md).

## Additional resources

* [Sprite atlas best practices in Unity 6](https://www.youtube.com/watch?v=hXlpnwD-TgY) on the Unity YouTube channel
* [2D game art, animation, and lighting for artists](https://unity.com/resources/2d-game-art-animation-lighting-unity-6-3-lts?isGated=false)