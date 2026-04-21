# Packing sprites into atlas textures

To reduce the number of [draw calls](../../DrawCallBatching.md) Unity sends to the GPU, create a **sprite** atlas. A **sprite atlas** combines multiple textures into a single texture. Unity only needs to create one draw call for all the sprites in a sprite atlas.

| **Page** | **Description** |
| --- | --- |
| [Sprite atlases](atlas-introduction.md) | Learn about sprite atlases and how to use them to optimize performance. |
| [Create a sprite atlas](create-sprite-atlas.md) | Create a sprite atlas, add sprites and textures to it, and analyze and optimize it. |
| [Create lower resolution versions of sprite atlases](master-variant/master-variant-sprite-atlases.md) | To create different versions of the same sprite atlas for different platforms, create sprite atlas variants. |
| [Load sprite atlases manually at runtime](distribution/load-sprite-atlas-spriteatlasmanageratlasrequested.md) | To avoid Unity loading sprite atlases when your project starts, load the sprite atlas yourself at runtime instead. |
| [Upgrade Sprite Atlas V1 assets](v2/sprite-atlas-v2.md) | Upgrade Sprite Atlas V1 assets from Unity version 2022.2 and earlier to Sprite Atlas V2 assets. |
| [Sprite Atlas Inspector window reference](sprite-atlas-reference.md) | Explore the properties and settings you can use to customize a sprite atlas. |

## Additional resources

* [Sprite atlas best practices in Unity 6](https://www.youtube.com/watch?v=hXlpnwD-TgY) on the Unity YouTube channel
* [2D game art, animation, and lighting for artists](https://unity.com/resources/2d-game-art-animation-lighting-unity-6-3-lts?isGated=false)