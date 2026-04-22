# Create lower resolution versions of sprite atlases

To create lower resolution versions of the same **sprite** atlas for different platforms, create variant sprite atlases. A variant **sprite atlas** combines the same sprites and textures as the original sprite atlas, but the combined texture it creates has a lower resolution.

Follow these steps:

1. [Create a sprite atlas](../create-sprite-atlas.md) as normal. Make sure its **Type** is set to **Master**. This sprite atlas is the parent sprite atlas.
2. Create another sprite atlas, then in its ****Inspector**** window set **Type** to **Variant**.
3. In the **Inspector** window of the variant sprite atlas, set **Master Atlas** to the parent sprite atlas you created in step 1. Either drag the parent sprite atlas from the **Project** window, or select the picker (**⊙**).

   A variant sprite atlas combines the same sprites and textures as the parent asset. It doesn’t have its own **Objects for Packing** list, and Unity doesn’t create a sprite atlas asset for it in the **Project** window.
4. To select the size of the combined texture of the variant sprite atlas, set the **Scale** property. For example, set **Scale** to 0.5 to create a sprite atlas texture half the resolution of the parent sprite atlas.
5. Select **Pack Preview** to display the packed texture in the **Preview** window at the bottom of the **Inspector** window.

## Use the variant sprite atlas

By default, Unity includes both the parent sprite atlas and the variant sprite atlas in your built project. This means Unity randomly chooses which sprite atlas to render with at runtime.

To avoid this, use any of the following approaches:

* To manually select a sprite atlas for a sprite, use the [SpriteAtlas.GetSprite](../../../../ScriptReference/U2D.SpriteAtlas.GetSprite.md) API.
* To manually control which sprite atlas Unity loads and uses, refer to [Load sprite atlases manually at runtime](../distribution/load-sprite-atlas-spriteatlasmanageratlasrequested.md).
* To exclude the parent sprite atlas from the build, disable **Include in Build** in the **Inspector** window of the parent sprite atlas.

**Important**: If you disable **Include in Build** for both the parent sprite atlas and the variant sprite atlas, Unity includes no textures for the sprites and textures, and the sprites that use the texture are invisible.

## Additional resources

* [Sprite atlas best practices in Unity 6](https://www.youtube.com/watch?v=hXlpnwD-TgY) on the Unity YouTube channel