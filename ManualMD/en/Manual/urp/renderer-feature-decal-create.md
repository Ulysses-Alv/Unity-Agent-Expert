# Create a decal via a Decal Renderer Feature in URP

To add decals to your **scene**:

1. [Add the Decal Renderer Feature](urp-renderer-feature.md) to the URP Renderer.
2. Create a Material, and assign it the `Shader Graphs/Decal` **shader**. In the Material, select the Base Map and the **Normal Map**.

   ![Example decal Material](../../uploads/urp/decal/decal-example-material.png)

   Example decal Material
3. Create a new Decal Projector **GameObject**, or add a [Decal Projector component](renderer-feature-decal.md#decal-projector) to an existing GameObject.

The following illustration shows a Decal Projector in the scene.

![Decal Projector in the scene.](../../uploads/urp/decal/decal-projector-selected-with-inspector.png)

Decal Projector in the scene.

For more information, refer to the [Decal Projector component](renderer-feature-decal.md#decal-projector).

An alternative way to add decals to a scene:

1. Create a **Quad** GameObject.
2. Assign a Decal Material to the GameObject.
3. Position the Quad on the surface where you want the decal to be. If necessary, adjust the [mesh bias](prebuilt-shader-graphs-urp-decal.md) value to prevent z-fighting.