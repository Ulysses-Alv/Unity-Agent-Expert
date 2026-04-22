# Introduction to post-processing in URP

The Universal **Render Pipeline** (URP) includes an integrated implementation of [post-processing](https://docs.unity3d.com/Manual/PostProcessingOverview.html) effects. This implementation uses the [volume](Volumes.md) framework for post-processing effects.

URP is not compatible with the [Post Processing Stack v2](https://docs.unity3d.com/Packages/com.unity.postprocessing@latest/index.html) package.

The images below show a **scene** with and without URP post-processing:

![A scene without post-processing](../../uploads/urp/AssetShots/Beauty/SceneWithoutPost.png)

A scene without post-processing

![A scene with post-processing.](../../uploads/urp/AssetShots/Beauty/SceneWithPost.png)

A scene with post-processing.

**Note**: URP does not support Post-processing on OpenGL ES 2.0.

## Post-processing in URP for mobile devices

Post-processing effects can take up a lot of frame time. If you’re using URP for mobile devices, these effects are the most “mobile-friendly” by default:

* Bloom (with **High Quality Filtering** disabled)
* Chromatic Aberration
* Color Grading
* Lens Distortion
* Vignette

**Note**: For depth-of field, Unity recommends that you use Gaussian **Depth of Field** for lower-end devices. For console and desktop platforms, use Bokeh Depth of Field.

## Post-processing in URP for VR

In **VR** apps and games, certain post-processing effects can cause nausea and disorientation. To reduce motion sickness in fast-paced or high-speed apps, use the Vignette effect for VR, and avoid the effects Lens Distortion, Chromatic Aberration, and Motion Blur for VR.