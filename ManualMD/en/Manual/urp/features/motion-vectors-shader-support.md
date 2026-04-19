# Built-in shader support for motion vectors in URP

This section describes which types of motion different **shader** types support for calculating motion vectors.

URP supports motion vectors only for opaque materials, including alpha-clipped materials. URP does not support motion vectors for transparent materials.

## URP ShaderLab shaders

URP `Lit`, `Unlit`, `Simple Lit`, `Complex Lit`, and `Baked Lit` shaders support per-object motion vectors for the following motion types:

* Rigid transform motion.
* Skeletal animation.
* Blend shape animation.
* Alembic animation.

To enable alembic motion vectors for particular material, enable the **Alembic Motion Vectors** checkbox in the **Advanced** section of the material **inspector**.

**Note:**: Use materials with the **Alembic Motion Vectors** checkbox enabled only on alembic vertex animation caches rendered with a **PlayableDirectors** component. When using such materials with regular draw calls and MeshRenderers, the materials cannot read the correct motion vector attribute stream, which results in incorrect motion vectors.

## URP Lit and Unlit Shader Graph targets

The URP Shader Graph `Lit` and `Unlit` targets support all of the motion vector features described in the [URP ShaderLab shaders](#urp-shaderlab-shaders) section. In addition to that, they have the **Additional Motion Vectors** setting with the following options:

* **Time-Based**: motion vectors for a Shader Graph vertex animation are generated automatically by running the vertex position subgraph twice with the current and the previous frame values for the **Time** node. This mode only works correctly for vertex animations that are computed procedurally based on the **Time** node, and that use only user-defined parameters (constants, attributes, buffers, textures) that do not change between frames.
* **Custom**: selecting this option adds an extra **Motion Vector** vertex output. The output lets you specify a custom object motion vector: a 3D local object space offset for each vertex from where it was in the previous frame. If you know how to compute the position of a vertex in the previous frame, you can author custom motion vectors for your vertex animations. Unity applies the custom motion vectors additively to transform motion, skeletal animation, and alembic animation before projecting the motion vector to the final 2D screen space vector.
* **None**: the default value. Use this option when the shader either does not modify the vertex output or if it does not animate the vertex deformation.

The options **Time-Based** and **Custom** enable the **MotionVectors** pass on the materials. The object motion vector pass is rendered every frame for them, even when the object’s transform is stationary between frames, unless the **MotionVectorGenerationMode** property is set to `Camera`.