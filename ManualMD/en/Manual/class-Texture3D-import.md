# Create a 3D texture

A 3D texture is a bitmap image that contains information in three dimensions rather than the standard two. 3D textures are commonly used to simulate volumetric effects such as fog or smoke, to approximate a volumetric 3D **mesh**, or to store animated textures and blend between them.

To create a 3D texture, import a 2D texture with the following properties:

* Flipbook format - a single 2D texture arranged as regular cells.
* Maximum size of 2048 × 2048 × 2048 **pixels**.

The size of a 3D texture on disk and in memory increases quickly as its resolution increases. For example, a 3D texture with [`RGBAFloat`](../ScriptReference/TextureFormat.RGBAFloat.md) format, no mipmaps, and a size of 32 × 32 × 32 pixels has a size of 512 KB (32 × 32 × 32 × 16 bytes) in memory. If you increase the resolution to 256 × 256 × 256 pixels, the size in memory becomes 256 MB (256 × 256 × 256 × 16 bytes).

Follow these steps to import:

1. Import the texture into your project.
2. In the **Project** window, select the texture asset. Unity displays the texture import settings in the ****Inspector**** window.
3. In the **Inspector** window, set **Texture Shape** to **3D**.
4. Set **Columns** and **Rows** to the appropriate values for your flipbook texture.
5. Select **Apply**.

Unity adds a 3D texture slice for each cell in the texture.

![Flipbook image as a 2D texture](../uploads/Main/TextureImporter-Flipbook-2D.jpg)
  
An imported flipbook texture with 8 columns and 8 rows.

![Flipbook image as a 3D texture](../uploads/Main/TextureImporter-Flipbook-3D.jpg)
  
The texture after you set **Texture Shape** to **3D**, **Columns** to **8**, and **Rows** to **8**.

Refer to [Importing textures](ImportingTextures.md) for more information.

## Additional resources

* [Free VFX image sequences and flipbooks](https://blog.unity.com/engine-platform/free-vfx-image-sequences-flipbooks)