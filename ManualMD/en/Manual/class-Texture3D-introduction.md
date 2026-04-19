# Introduction to 3D textures

A 3D texture is a bitmap image that contains information in three dimensions rather than the standard two. 3D textures are commonly used to simulate volumetric effects such as fog or smoke, to approximate a volumetric 3D **mesh**, or to store animated textures and blend between them.

The source texture files for **2D Array** and **3D** Textures are divided into cells; these textures are called flipbook textures. When Unity imports flipbook textures, it places the contents of each cell into its own 2D array layer or 3D texture slice.

For example, an image with 8x8 cells of smoke effect frames looks like this as a default 2D texture:

![Flipbook image as a 2D shape](../uploads/Main/TextureImporter-Flipbook-2D.jpg)

Flipbook image as a 2D shape

But when you correctly import is as a 3D texture with 8 Columns and 8 Rows, it looks like this:

![Flipbook image as a 3D shape](../uploads/Main/TextureImporter-Flipbook-3D.jpg)

Flipbook image as a 3D shape