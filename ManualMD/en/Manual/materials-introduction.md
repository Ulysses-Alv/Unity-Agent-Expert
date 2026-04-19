# Introduction to materials

To draw something in Unity, you must provide information that describes its shape, and information that describes the appearance of its surface. You use [meshes](class-Mesh.md) to describe shapes, and materials to describe the appearance of surfaces.

Materials and **shaders** are closely linked; you always use materials with shaders.

## Render pipeline compatibility

| Feature | Universal **Render Pipeline** (URP) | High Definition Render Pipeline (HDRP) | Custom Scriptable Render Pipeline (SRP) | Built-in Render Pipeline |
| --- | --- | --- | --- | --- |
| Materials | Yes | Yes | Yes | Yes |

## Material fundamentals

A material contains a reference to a [Shader object](shader-objects.md). If that Shader object defines [material properties](SL-Properties.md), then the material can also contain data such as colors or references to textures.

The [Material](../ScriptReference/Material.md) class represents a material in C# code. For information, see [Using Materials with C# scripts](MaterialsAccessingViaScript.md).

A material asset is a file with the `.mat` extension. It represents a material in your Unity project. For information on viewing and editing a material asset using the **Inspector** window, see [Material Inspector reference](class-Material.md).

## Material Variants

Unity supports functionality for creating variants of Materials. To learn more about this functionality, see [Material Variants](materialvariant-landingpage.md).