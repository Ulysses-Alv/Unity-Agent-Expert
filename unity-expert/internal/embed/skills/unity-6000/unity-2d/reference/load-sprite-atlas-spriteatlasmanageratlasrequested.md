# Load sprite atlases manually at runtime

To avoid Unity loading **sprite** atlases and attaching them to your sprites when your project starts, load the **sprite atlas** yourself at runtime instead. This is called late binding.

For example, load a lower-resolution [variant sprite atlas](../master-variant/master-variant-sprite-atlases.md) depending on the hardware, or load a set of sprites into memory only when a later game level needs it.

**Note:** To get a list of the sprites packed into a sprite atlas, for example if you need to manually instantiate sprites at runtime, use the `GetSprites` method of the `SpriteAtlas` API. For more information, refer to [SpriteAtlas.GetSprites](SpriteRef:U2D.SpriteAtlas.GetSprites).

## Exclude a sprite atlas from the build

To prevent Unity automatically including your sprite atlas in your build, follow these steps:

1. In the **Project** window, select the sprite atlas.
2. In the ****Inspector**** window, disable **Include in Build**.

**Important**: By default, this doesn’t affect the **Scene** view and Game view, because they continue to reference the `.spriteatlasv2` file in the **Project** window. However, if you enter Play mode or build your project, the sprites disappear, because they now reference the missing sprite atlas in the build.

## Load the sprite atlas at runtime

When you disable **Include in Build**, your sprites still reference the sprite atlas when your project starts, but the sprite atlas texture is missing. You must manually bind the sprite atlas when sprites request it.

Use one of the following methods:

* Load the sprite atlas from the [Resources system](../../../assets-resources-system.md) and load it with the `Resources.Load` API.
* Load the sprite atlas from an [AssetBundle](../../../AssetBundlesIntro.md) and load it with the `AssetBundle` API.

### Load the sprite atlas from the Resources system

Follow these steps:

1. In the **Project** window, create a folder called **Resources**.
2. Move the sprite atlas asset into the **Resources** folder. Unity automatically includes assets from the **Resources** folder in the build, but doesn’t load them.
3. In a C# script, attach a new method to the `SpriteAtlasManager.atlasRequested` callback. The `atlasRequested` callback triggers when a sprite requests an unloaded sprite atlas. For example:

   ```
   void OnEnable()
   {
       SpriteAtlasManager.atlasRequested += MySpriteAtlasLoader;
   }
   ```
4. In the method, use the `Resources.Load` API to load the sprite atlas from the **Resources** folder and pass it back to Unity. For example:

   ```
   // Unity passes in the name of the sprite atlas that the sprite is requesting, and a callback to call
   void MySpriteAtlasLoader(string spriteAtlasName, System.Action<SpriteAtlas> callback)
   {
       // Load the sprite atlas
       var spriteAtlas = UnityEngine.Resources.Load<SpriteAtlas>(spriteAtlasName);

       // Pass the sprite atlas back to Unity
       callback(spriteAtlas);
   }
   ```

For more information, refer to [Use the Resources system to load assets at runtime](../../../assets-resources-system.md).

For a complete example, refer to the Sprite Atlas Examples in the 2D Common package. For more information about package samples, refer to [The Package Manager window](../../../upm-ui.md).

### Load the sprite atlas from an AssetBundle

Follow these steps:

1. Assign the sprite atlas to an AssetBundle. For more information, refer to [Assign assets to an AssetBundle](../../../assetbundles-assign-assets.md).

   **Important:** To avoid duplicate sprites in your project, assign the sprites to the same AssetBundle. For more information, refer to [Avoiding asset duplication](../../../assets-avoid-duplication.md).

   **Note:** If you assign a sprite atlas to an AssetBundle but enable **Include in Build**, you don’t need to load the sprite atlas yourself at runtime. Unity loads the sprite atlas automatically.
2. In a C# script, attach a new method to the `SpriteAtlasManager.atlasRequested` callback. The `atlasRequested` callback triggers when a sprite requests an unloaded sprite atlas. For example:

   ```
   void OnEnable()
   {
       SpriteAtlasManager.atlasRequested += MySpriteAtlasLoader;
   }
   ```
3. In the method, use the `AssetBundle` API to load the sprite atlas from the AssetBundle and pass it back to Unity. For example:

   ```
   // Unity passes in the name of the sprite atlas that the sprite is requesting, and a callback to call
   void MySpriteAtlasLoader(string spriteAtlasName, System.Action<SpriteAtlas> callback)
   {
       // Load the sprite atlas from the AssetBundle
       var bundle = AssetBundle.LoadFromFile("path/to/assetbundle");
       var spriteAtlas = bundle.LoadAsset<SpriteAtlas>(spriteAtlasName);

       // Pass the sprite atlas back to Unity
       callback(spriteAtlas);
   }
   ```

For more information, refer to [Loading assets from AssetBundles](../../../AssetBundles-Native.md).

For a complete example, refer to the Sprite Atlas Examples in the 2D Common package. For more information about package samples, refer to [The Package Manager window](../../../upm-ui.md).

## Additional resources

* [Create lower resolution versions of sprite atlases](../master-variant/master-variant-sprite-atlases.md)