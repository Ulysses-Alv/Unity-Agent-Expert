# Texture types

## Albedo texture maps

Albedo texture maps contain the base colors of a material. Make sure albedo textures don’t contain any lighting or shadows, because Unity adds lighting based on the positions of the object and the lights.

![Two examples of albedo texture maps. On the left is a texture map for a character model, and on the right is a wooden crate. There are no shadows or lighting in the textures.](../uploads/Main/StandardShaderAlbedoTextureExamples.jpg)

Two examples of albedo texture maps. On the left is a texture map for a character model, and on the right is a wooden crate. There are no shadows or lighting in the textures.

## Normal maps

Normal maps are used by **normal map** **Shaders** to make low-polygon models look as if they contain more detail. Unity uses normal maps encoded as RGB images. You also have the option to generate a normal map from a grayscale height map image.

### Alpha maps

An alpha map is a Texture that contains only alpha information. You can use an alpha map to apply varying levels of transparency to a Material.

You can create an alpha map by creating a Texture with information in the alpha channel, or by creating a grayscale Texture and converting the grayscale values to alpha in Unity.

See the documentation on [Texture Import Settings](class-TextureImporter.md) for more information.

## Detail maps

If you want to make a **Terrain**, you normally use your main Texture to show areas of terrain such as grass, rocks and sand. If your terrain is large, it may end up very blurry. Detail Textures hide this fact by fading in small details as your main Texture gets closer.

When drawing Detail Textures, a neutral gray is invisible, white makes the main Texture twice as bright, and black makes the main Texture completely black.

See documentation on [Secondary Maps (Detail Maps)](StandardShaderMaterialParameterDetail.md) for more information.

## Terrain Heightmaps

Textures can even be used in cases where the image will never be viewed at all, at least not directly. In a greyscale image, each **pixel** value is simply a number corresponding to the shade of grey at that point in the image (this could be a value in the range 0..1 where zero is black and one is white, say). Although an image like this can be viewed, there is no reason why the numeric pixel values can’t be used for other purposes as well, and this is precisely what is done with **Terrain Heightmaps**.

A *terrain* is a **mesh** representing an area of ground where each point on the ground has a particular height from a baseline. The *heightmap* for a terrain stores the numeric height samples at regular intervals as greyscale values in an image where each pixel corresponds to a grid coordinate on the ground. The values are not shown in the **scene** as an image but are converted to coordinates that are used to generate the terrain mesh.

Interestingly, even though a **heightmap** is not viewed directly as an image, there are still common image processing techniques that are useful when applied to the height data. For example, adding noise to a heightmap will create the impression of rocky terrain while blurring will smooth it out to produce a softer, rolling landscape.

More information about terrains in Unity can be found in [this section](script-Terrain.md) of the manual.

## Additional resources

* [Textures reference](textures-reference.md)
* [Default texture Import Settings window reference](texture-type-default.md)
* [Normal Map texture Import Settings window reference](texture-type-normal-map.md)