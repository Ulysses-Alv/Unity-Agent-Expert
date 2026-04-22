# Troubleshooting motion vectors in URP

Solve common issues with motion vectors in the Universal **Render Pipeline** (URP).

## Fix motion vectors that are too large

If a **camera** is locked to an object that moves, for example, a model of a car in a racing game, select the **Per Object Motion** option in the **Motion Vectors** property of that object. If you don’t select that option, the object will have incorrectly large motion vectors. This happens because Unity calculates camera motion vectors assuming that the geometry of the object is static in the world, and that the camera is moving relative to it. This might cause significant TAA or motion blur artifacts.

## Fix visual artefacts

When the motion vector texture is used by full-screen **post-processing** effects, such as TAA and motion blur, any screen regions with incorrect motion vectors (either missing, or inaccurate) will likely exhibit visual artifacts. The artifacts can include: texture blurring, movement ghosting, improper anti-aliasing, non-realistic or inappropriate motion blur, and so on.

If you are experiencing artifacts that you suspect are caused by incorrect motion vectors, check if the affected objects have object motion vector rendering enabled. In the **Frame Debugger**, the object should be present in the **MotionVectors** pass. To troubleshoot missing or incorrect motion vectors for a particular object, refer to the section [Cases when Unity renders motion vectors](motion-vectors.md#cases-when-motion-vectors).