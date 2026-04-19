# Introduction to decals in URP

## Decal Renderer Feature

With the Decal Renderer Feature, Unity can project specific Materials (decals) onto other objects in the **scene**. The decals interact with the scene’s lighting and wrap around Meshes.

![Sample scene without decals](../../uploads/urp/decal/decal-sample-without.png)  
*Sample scene without decals*

![Sample scene with decals](../../uploads/urp/decal/decal-sample-with.png)  
*Sample scene with decals. The decals hide the seams between materials and add artistic details.*

For examples of how to use Decals, refer to the [Decals samples in URP Package Samples](package-sample-urp-package-samples.md#decals).

### Limitations

This feature has the following limitations:

* The decal projection does not work on transparent surfaces.

## Decal Projector

The Decal Projector component lets Unity project decals onto other objects in the scene. A Decal Projector component must use a Material with the [Decal Shader Graph](decal-shader.md) assigned (`Shader Graphs/Decal`).

## Performance

Decals do not support the **SRP Batcher** by design because they use Material property blocks. To reduce the number of draw calls, decals can be batched together using GPU instancing. If the decals in your scene use the same Material, and if the Material has the **Enable GPU Instancing** property turned on, Unity instances the Materials and reduces the number of draw calls.

To reduce the number of Materials necessary for decals, put multiple decal textures into one texture (atlas). Use the UV offset properties on the decal projector to determine which part of the atlas to display.

The following image shows an example of a decal atlas.

![Decal Atlas](../../uploads/urp/decal/decal-atlas.png)   
 *left: decal atlas with four decals. Right: a decal projector is projecting one of them. If the decal Material has GPU instancing enabled, any instance of the four decals is rendered in a single instanced draw call.*