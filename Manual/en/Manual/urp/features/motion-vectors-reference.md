# Motion vectors settings reference for URP

To specify how a GameObject contributes to the motion vector buffer, use the **Motion Vectors** property: **Mesh Renderer** >> **Additional Settings** >> **Motion Vectors**. This property lets you disable motion vector rendering for a specific object, or fill the motion vector texture with zeros for the visible fragments of the object’s renderer.

The following table describes the available **Motion Vectors** property options.

| **Motion Vectors** option | Description |
| --- | --- |
| **Camera** Motion Only | Unity treats the object as stationary in the world when rendering camera motion vectors. Unity does not draw a per-object motion vector pass for this MeshRenderer. If motion vector rendering is a GPU bottleneck, you can use this option as an optimization for objects which move slowly. |
| Per Object Motion | Unity renders the per-object **Motion Vectors** pass for this object. |
| Force No Motion | Unity renders the per-object **Motion Vectors** pass every frame for this object, but sets a special shader uniform variable to tell the pass to skip calculations and to write zero values. A per-object pass is still necessary to overwrite any non-zero camera motion vectors from the full-screen pass. You can use this option to avoid artefacts from the camera motion blur on a 3D HUD, or a third person character, or a race car, or if you have some other artefacts related to incorrect motion vectors on an object that you would like to avoid. |