# Introduction to motion vectors in URP

URP calculates the frame-to-frame screen-space movement of surface fragments using a [motion vector render pass](motion-vectors-sample.md). URP stores the movement data in a full-screen texture and the stored per-pixel values are called [motion vectors](#definition).

Unity runs the motion vector render pass only when there are active features in the frame with render passes that request it. For example, the following URP features request the motion vector pass: [temporal anti-aliasing](../anti-aliasing.md#taa) and [motion blur](../Post-Processing-Motion-Blur.md). For information on how to request the motion vector pass in a custom passes, refer to section [Using the motion vector texture in your passes](motion-vectors-sample.md).

Incorrect or missing motion vectors can result in [visual artifacts](motion-vectors-troubleshooting.md) in effects that rely on them. Follow the instructions on this page to ensure that your object renderers, Materials, and **shaders** are set up correctly to support motion vectors.

URP supports motion vectors only for opaque materials, including alpha-clipped materials. URP does not support motion vectors for transparent materials.

## Implementation details

This section describes how URP implements motion vectors.

### Motion vector definition

A motion vector is a 2D vector representing a surface fragment’s motion relative to the **camera** since the last frame, projected onto the camera’s near **clipping plane**. The motion vector texture uses 2 channels (R and G). Each texel stores the UV offset of each visible surface fragment. If you subtract the motion vector for a given texel from its current UV coordinate, you will get the UV coordinate of where this texel would have been on the screen last frame. The computed last frame UV coordinate can be outside the screen bounds.

### Object motion vectors and camera-only motion vectors

There are two categories of motion vectors:

* **Camera motion vectors**: motion vectors caused only by the camera’s own motion.
* **Object motion vectors**: motion vectors caused by a combination of the camera’s motion and the world-space motion of the object the fragment belongs to.

Given only the motion vector texture it’s impossible to infer whether a fragment’s motion on screen is due to only camera motion, object motion or a combination of both.

A single full-screen pass is enough to calculate camera motion vectors. Such pass has a fixed per-frame computation load independent of **scene** complexity. It’s only necessary to know the current 3D positions of all **pixels** on screen and how the camera moved, which can be inferred from the **depth buffer** and the camera matrices for the current and the previous frames.

Object motion vectors have computation load which depends on the number and complexity of the moving objects in the scene because a draw per-object is required to account for each object’s motion. Each draw needs to apply the camera’s motion contribution too.

### Cases when Unity renders per-object motion vectors

Unity renders object motion vectors for a **mesh** in a frame when the following three conditions are met:

1. The shader associated with the mesh has a [MotionVectors pass](motion-vectors-shader-support.md) in an active SubShader block.
2. The mesh is being rendered through any of the following renderers:

   1. **SkinnedMeshRenderer**.
   2. **MeshRenderer**, when its **Motion Vectors** property is not set to `Camera Motion Only`.
   3. Using the following APIs: [Graphics.RenderMesh](https://docs.unity3d.com/ScriptReference/Graphics.RenderMesh.html), [Graphics.RenderMeshInstanced](https://docs.unity3d.com/ScriptReference/Graphics.RenderMeshInstanced.html) or [Graphics.RenderMeshIndirect](https://docs.unity3d.com/ScriptReference/Graphics.RenderMeshIndirect.html), with the `MotionVectorMode` member of the `RenderParams` struct not set to `Camera`.
3. If any of the following conditions is true:

   1. The `MotionVectorGenerationMode` property is set to `ForceNoMotion` on the `MeshRenderer` or the `RenderParams.MotionVectorMode` struct member of the `Grphics.Render...` APIs.

      **Note:** The `ForceNoMotion` option requires a per-object motion vector pass to be rendered every frame so that the camera motion vectors for such objects can be overwritten with zeros.
   2. The **MotionVectors** pass [is enabled on the material](https://docs.unity3d.com/ScriptReference/Material.SetPass.html) (for example, when a material has a vertex animation in Shader Graph or alembic animation).
   3. The **MotionVectors** pass [is disabled on the material](https://docs.unity3d.com/ScriptReference/Material.SetPass.html) but the model matrices for the previous and the current frame don’t match, or if the object has skeletal animation. Stationary renderers without skeletal animation are not rendered with an object motion vector pass if their shader has a `MotionVectors` pass but it’s disabled on their material.

### MotionVectors pass in URP pre-built shaders

When the **MotionVectors** pass is enabled for the pre-built shaders, Unity renders object motion vectors for **mesh renderers** even if they don’t move.

Unity enables the **MotionVectors** pass in the pre-built URP shaders when the following conditions are true:

* URP **ShaderLab** shaders: on a Material, the **Alembic Motion Vectors** property is enabled in the **Advanced Options** section.
* URP Lit and Unlit Shader Graph shaders: in the **Graph **Inspector****, any of the following properties is set:

  + **Alembic Motion Vectors** is enabled.
  + **Additional Motion Vectors** property is set to **TimeBased** or **Custom**.