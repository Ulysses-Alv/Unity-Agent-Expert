# Introduction to post-processing

Unity provides a number of ****post-processing**** effects and full-screen effects that can greatly improve the appearance of your application with little set-up time. You can use these effects to simulate physical **camera** and film properties, or to create stylised visuals.

This page contains the following information:

* [Render pipeline compatibility](#render-pipeline-compatibility)
* [Effect availability and location](#effect-availability-and-location)

The images below demonstrate a **Scene** with and without post-processing.

![Scene with no post-processing](../uploads/Main/PostProcessing-0.jpg)

Scene with no post-processing

![Scene with post-processing](../uploads/Main/PostProcessing-1.jpg)

Scene with post-processing

## Render pipeline compatibility

Which post-processing effects are available and how you apply them depend on which [render pipeline](render-pipelines.md) you are using. A post-processing solution from one render pipeline is not compatible with other render pipelines.

This table contains information about which of Unity’s post-processing solutions are compatible with each of Unity’s render pipelines.

| **Render pipeline** | **Post-processing support** |
| --- | --- |
| **[Universal Render Pipeline (URP)](urp/urp-introduction.md)** | URP includes its own post-processing solution, which Unity installs when you create a Project using a URP Template. For information on using post-processing effects in URP, see the [URP post-processing documentation](urp/integration-with-post-processing.md). |
| **[High Definition Render Pipeline (HDRP)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?preview=1)** | HDRP includes its own post-processing solution, which Unity installs when you create a Project using an HDRP Template. For information on using post-processing effects in HDRP, see the [HDRP post-processing documentation](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Post-Processing-Main.html). |
| [**Built-in Render Pipeline**](SL-RenderPipeline.md) | The Built-in Render Pipeline does not include a post-processing solution by default. To use post-processing effects with the Built-in Render Pipeline, download the [Post-Processing Version 2](https://docs.unity3d.com/Packages/com.unity.postprocessing@latest) package. For information on using post-processing effects in the Built-in Render Pipeline, see the [Post-Processing Version 2](https://docs.unity3d.com/Packages/com.unity.postprocessing@latest) documentation. |

**Note:** Post processing stack version 1 is now deprecated and should not be used.

## Effect availability and location

For information on the specific post-processing effects each render pipeline supports, refer to [Post-processing effect availability reference](post-processing-effect-availability-reference.md).