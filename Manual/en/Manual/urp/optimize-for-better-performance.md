# Adjust settings to improve performance in URP

If the performance of your Universal **Render Pipeline** (URP) project seems slow, you can adjust settings to increase performance.

Based on your [analysis](analyze-your-project.md), you can adjust the following settings in the [Universal Render Pipeline (URP) Asset](universalrp-asset.md) or the [Universal Renderer asset](urp-universal-renderer.md) to improve the performance of your project.

Depending on your project or the platforms you target, some settings might not have a significant effect. There might also be other settings that have an effect on performance in your project.

| **Setting** | **Where the setting is** | **What to do for better performance** |
| --- | --- | --- |
| **Accurate G-buffer normals** | [Universal Renderer](urp-universal-renderer.md) > **Rendering** | Disable if you use the Deferred rendering path |
| **Additional Lights** > **Cast Shadows** | [URP Asset](universalrp-asset.md) > **Lighting** | Disable |
| **Additional Lights** > **Cookie Atlas Format** | URP Asset > **Lighting** | Set to **Color Low** |
| **Additional Lights** > **Cookie Atlas Resolution** | URP Asset > **Lighting** | Set to the lowest you can accept |
| **Additional Lights** > **Per Object Limit** | URP Asset > **Lighting** | Set to the lowest you can accept. This setting has no effect if you use the Deferred or Forward+ rendering paths. |
| **Additional Lights** > **Shadow Atlas Resolution** | URP Asset > **Lighting** | Set to the lowest you can accept |
| **Additional Lights** > **Shadow Resolution** | URP Asset > **Lighting** | Set to the lowest you can accept |
| **Cascade Count** | URP Asset > **Shadows** | Set to the lowest you can accept |
| **Conservative Enclosing Sphere** | URP Asset > **Shadows** | Enable |
| **Technique** | [Decal Renderer Feature](renderer-feature-decal.md) | Set to **Screen Space**, and set **Normal Blend** to **Low** or **Medium** |
| **Fast sRGB/Linear conversion** | URP Asset > **Post Processing** | Enable |
| **Grading Mode** | URP Asset > **Post Processing** | Set to **Low Dynamic Range** |
| **LOD Cross Fade Dither** | URP Asset > **Quality** | Set to **Bayer Matrix** |
| **LUT size** | URP Asset > **Post Processing** | Set to the lowest you can accept |
| **Main Light** > **Cast Shadows** | URP Asset > **Lighting** | Disable |
| **Max Distance** | URP Asset > **Shadows** | Reduce |
| **Opaque Downsampling** | URP Asset > **Rendering** | If **Opaque Texture** is enabled in the URP Asset, set to **4x Bilinear** |
| **Render Scale** | URP Asset > **Quality** | Set to below 1.0 |
| **Soft Shadows** | URP Asset > **Shadows** | Disable, or set to **Low** |
| **Upscaling Filter** | URP Asset > **Quality** | Set to **Bilinear** or **Nearest-Neighbor** |

Refer to the following for more information on the settings:

* [Deferred Rendering Path in URP](rendering/deferred-rendering-path-landing.md)
* [Forward rendering paths in URP](rendering/forward-rendering-paths.md)
* [Decal Renderer Feature](renderer-feature-decal.md)
* [Universal Render Pipeline Asset](universalrp-asset.md)
* [Universal Renderer](urp-universal-renderer.md)

## Additional resources

* [Understand performance in URP](understand-performance.md)
* [Configure for better performance](configure-for-better-performance.md)
* [Graphics performance and profiling](https://docs.unity3d.com/Manual/graphics-performance-profiling.html)
* [Best practices for profiling game performance](https://unity.com/how-to/best-practices-for-profiling-game-performance)
* [Tools for profiling and debugging](https://unity.com/how-to/profiling-and-debugging-tools)
* [Graphics rendering: Getting the best performance with Unity 6](https://www.youtube.com/watch?v=Oc6T4hh5gaI)
* [Performance tips and tricks from a Unity consultant](https://www.youtube.com/watch?v=CmD8MVGkDxQ)