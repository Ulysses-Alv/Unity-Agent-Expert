# Load and unload assets with the Resources system

To load assets with the Resources system:

1. Create a new folder called `Resources` in your project, and add assets to it. Unity then makes these assets available even if they’re not directly referenced in a **scene**. **Note:** You can have multiple `Resources` folders located in different subfolders within your `Assets` folder, and packages can also contain `Resources` folders.
2. Whenever you want to load an asset from one of these folders, call [`Resources.Load`](../ScriptReference/Resources.Load.md) in your code. Only assets in the `Resources` folder can be accessed in this way.

For example, you can apply the following script to load and apply a texture to it:

```
using UnityEngine;

public class LoadTexture : MonoBehaviour
{
    // Reference a material to apply the texture to
    public Renderer targetRenderer;

    void Start()
    {
        // Load a texture from the Textures folder inside the Resources folder
        Texture2D loadedTexture = Resources.Load<Texture2D>("Textures/MyTexture");

        if (loadedTexture != null)
        {
            // Apply the loaded texture to the material of the target renderer
            targetRenderer.material.mainTexture = loadedTexture;

            Debug.Log("Texture successfully loaded and applied.");
        }
        else
        {
            Debug.LogError("Failed to load texture from Resources.");
        }
    }
}
```

Unity stores all assets in the `Resources` folders and their dependencies in a file in the build output called `resources.assets`. If a scene in the build references an asset, Unity serializes that asset into a `sharedAssets*.assets` file instead.

Additional assets might end up in the `resources.assets` file if they’re dependencies. For example, a material in the `Resources` folder might reference a texture outside of the `Resources` folder. In that case the texture is also included in the `resources.assets` file, but isn’t available to load directly.

## Unload assets

If you want to destroy objects that [`Resources.Load`](../ScriptReference/Resources.Load.md) loaded before loading another scene, call [`Object.Destroy`](../ScriptReference/Object.Destroy.md) on them. To recover memory used by unreferenced objects, use [`Resources.UnloadUnusedAssets`](../ScriptReference/Resources.UnloadUnusedAssets.md).

## Additional resources

* [Introduction to the Resources system](LoadingResourcesatRuntime.md)
* [Introduction to asset management](assets-managing-introduction.md)
* [`Resources.Load` API](../ScriptReference/Resources.Load.md)