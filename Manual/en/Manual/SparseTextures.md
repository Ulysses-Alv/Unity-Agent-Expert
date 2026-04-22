# Sparse Textures

**Sparse textures** (also known as “tiled textures” or “mega-textures”) are textures that are too large to fit in graphic memory in their entirety. To handle them, Unity breaks the main texture down into smaller rectangular sections known as “tiles”. Individual tiles can then be loaded as necessary. For example, if the **camera** can only see a small area of a sparse texture, then only the tiles that are currently visible need to be in memory.

Aside from the tiling, a sparse texture behaves like any other texture in usage. **Shaders** can use them without special modification and they can have mipmaps, use all texture filtering modes, etc. If a particular tile cannot be loaded for some reason then the result is undefined; some GPUs show a black area where the missing tile should be but this behaviour is not standardised.

Not all hardware and platforms support sparse textures. For example, on DirectX systems they require [DX11.2](http://msdn.microsoft.com/en-us/library/windows/desktop/dn312084.aspx) (Windows 8.1) and a fairly recent GPU. On OpenGL they require [ARB\_sparse\_texture](http://www.opengl.org/registry/specs/ARB/sparse_texture.txt) extension support. Sparse textures only support non-compressed **texture formats**.

See the [SparseTexture](../ScriptReference/SparseTexture.md) script reference page for further details about handling sparse textures from **scripts**.

## Example

The following example loads only the texture for the tile directly under the camera. The example only works if your hardware and platform support sparse textures. To use the example:

1. Create a **scene** with a large plane **GameObject**. For more information, refer to [Primitive and placeholder objects](PrimitiveObjects.md).
2. Attach the script to the GameObject.
3. Enter Play mode.
4. Change the x and y position of the camera. The **Scene view** shows the tiles loading and unloading.

```
using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class SparseTextureExample : MonoBehaviour {

    private int tileCountX;
    private int tileCountY;
    private Bounds bounds;
    private SparseTexture texture;
    private Color32[] colorForSingleTile;
    private Vector2 lastTileVisible;

    void Awake()
    {
        if (!SystemInfo.supportsSparseTextures)
        {
            Debug.Log("Sparse textures are not supported on this platform.");
        }

        bounds = GetComponent<Renderer>().bounds;
        lastTileVisible = new Vector2(0f, 0f);

        // Create the sparse texture
        DestroyImmediate(texture);
        texture = new SparseTexture(1024, 1024, TextureFormat.RGBA32, 1, false);

        // Set the texture on the material
        GetComponent<Renderer>().sharedMaterial.mainTexture = texture;

        // Get the number of tiles from the sparse texture
        tileCountX = (texture.width + texture.tileWidth - 1) / texture.tileWidth;
        tileCountY = (texture.height + texture.tileHeight - 1) / texture.tileHeight;
        Debug.Log("Tile size: " + texture.tileWidth + " x " + texture.tileHeight + " pixels");
        Debug.Log("Tile count: " + tileCountX + " x " + tileCountY);

        // Fill the color array for a single tile with red
        colorForSingleTile = new Color32[texture.tileWidth * texture.tileHeight];
        for(int i = 0; i < colorForSingleTile.Length; i++)
        {
            colorForSingleTile[i] = new Color32(255, 0, 0, 255);
        }
    }

    void LateUpdate()
    {
        // Get camera position
        Vector3 camPos = Camera.main.transform.position;

        // Get tile index x and y under the camera position
        float px = Mathf.Clamp(Mathf.InverseLerp (bounds.min.x, bounds.max.x, camPos.x), 0.0f, 0.9999f);
        float py = Mathf.Clamp(Mathf.InverseLerp (bounds.min.z, bounds.max.z, camPos.z), 0.0f, 0.9999f);
        int tileX = (int)(px * 1024 / texture.tileWidth);
        int tileY = (int)(py * 1024 / texture.tileHeight);

        // Update the tile with the red color
        // In a real project, you might load the texture from disk or generate it procedurally
        texture.UpdateTile(tileX, tileY, 0, colorForSingleTile);

        // Use the SparseTexture API to unload the texture of the previous tile under the camera
        Vector2 tileVisible = new Vector2(tileX, tileY);
        if(lastTileVisible != tileVisible)
        {           
            Debug.Log("Tile under camera: " + tileX + ", " + tileY);

            texture.UnloadTile((int)lastTileVisible.x, (int)lastTileVisible.y, 0);
            Debug.Log("Unloaded: " + (int)lastTileVisible.x + ", " + (int)lastTileVisible.y);

            lastTileVisible = tileVisible;
        }
    }

    void OnDisable ()
    {
        GetComponent<Renderer>().sharedMaterial.mainTexture = null;
        DestroyImmediate(texture);
    }
}
```