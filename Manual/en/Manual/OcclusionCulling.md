# Occlusion culling

Occlusion culling is a process which prevents Unity from performing rendering calculations for **GameObjects** that are completely hidden from view (occluded) by other GameObjects.

Every frame, **Cameras** perform culling operations that examine the Renderers in the **Scene** and exclude (cull) those that do not need to be drawn. By default, Cameras perform frustum culling, which excludes all Renderers that do not fall within the Camera’s [view frustum](UnderstandingFrustum.md#view-frustum). However, frustum culling does not check whether a Renderer is occluded by other GameObjects, and so Unity can still waste CPU and GPU time on rendering operations for Renderers that are not visible in the final frame. **Occlusion culling** stops Unity from performing these wasted operations.

![Regular frustum culling renders all Renderers within the Cameras view.](../uploads/Main/OcclusionFrustumCulling.jpg)

Regular frustum culling renders all Renderers within the Camera’s view.

![Occlusion culling removes Renderers that are entirely obscured by nearer Renderers.](../uploads/Main/OcclusionFullCulling.jpg)

Occlusion culling removes Renderers that are entirely obscured by nearer Renderers.

## When to use occlusion culling

To determine whether occlusion culling is likely to improve the runtime performance of your Project, consider the following:

* Preventing wasted rendering operations can save on both CPU and GPU time. Unity’s built-in occlusion culling performs runtime calculations on the CPU, which can offset the CPU time that it saves. Occlusion culling is therefore most likely to result in performance improvements when a Project is GPU-bound due to overdraw.
* Unity loads occlusion culling data into memory at runtime. You must ensure that you have sufficient memory to load this data.
* Occlusion culling works best in Scenes where small, well-defined areas are clearly separated from one another by solid GameObjects. A common example is rooms connected by corridors.
* You can use occlusion culling to occlude Dynamic GameObjects, but Dynamic GameObjects cannot occlude other GameObjects. If your Project generates Scene geometry at runtime, then Unity’s built-in occlusion culling is not suitable for your Project.

## How occlusion culling works

Occlusion culling generates data about your Scene in the Unity Editor, and then uses that data at runtime to determine what a Camera can see. The process of generating data is known as baking.

**Note:** If your project uses the Universal **Render Pipeline** (URP) or the High Definition Render Pipeline (HDRP), you can instead [Enable GPU occlusion culling in URP](urp/gpu-culling.md) or [Enable GPU occlusion culling in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.4/manual/gpu-culling.html).

When you bake occlusion culling data, Unity divides the Scene into cells and generates data that describes the geometry within cells, and the visibility between adjacent cells. Unity then merges cells where possible, to reduce the size of the generated data. To configure the baking process, you can change parameters in the [Occlusion Culling window](occlusion-culling-window.md), and use [Occlusion Areas](class-OcclusionArea.md) in your Scene.

At runtime, Unity loads this baked data into memory, and for each Camera that has its Occlusion Culling property enabled, it performs queries against the data to determine what that Camera can see. Note that when occlusion culling is enabled, Cameras perform both frustum culling and occlusion culling.