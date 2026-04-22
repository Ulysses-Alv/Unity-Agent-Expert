# Motion vectors render pass in URP

Understand how the **MotionVectors** render pass renders the motion vector texture.

## Location in the frame loop

URP renders motion vectors at the `BeforeRenderingPostProcessing` event. Before that event the motion vector texture might not be set or might contain previous frame’s motion vector data.

## MotionVectors pass structure

URP renders the motion vector texture in 2 steps:

1. URP renders the **camera** motion vectors in the **MotionVectors** full-screen pass. This pass uses the depth texture and the camera matrices for the current and the previous frames to calculate the camera motion vectors. This pass has a fixed per-camera computation load and does not require special motion vector support from renderers or materials.
2. URP draws a per-object motion vector **shader** pass for [each renderer and material combination that supports motion vectors](motion-vectors.md#cases-when-motion-vectors).