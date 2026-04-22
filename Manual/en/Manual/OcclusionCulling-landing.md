# Excluding hidden objects with occlusion culling

Resources about preventing Unity doing rendering calculations for hidden **GameObjects**.

**Note:** If your project uses the Universal **Render Pipeline** (URP) or the High Definition Render Pipeline (HDRP), you can instead [Enable GPU occlusion culling in URP](urp/gpu-culling.md) or [Enable GPU occlusion culling in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.4/manual/gpu-culling.html).

| **Page** | **Description** |
| --- | --- |
| [Occlusion culling](OcclusionCulling.md) | Understand how occlusion culling checks for hidden objects, and when to use it. |
| [Set up a scene for occlusion culling](occlusion-culling-getting-started.md) | Set up a **scene** for occlusion culling, bake occlusion culling data, and check the results. |
| [Set up multiple scenes for occlusion culling](occlusion-culling-scene-loading.md) | Prepare occlusion culling data differently if you load multiple scenes at a time. |
| [Cull moving GameObjects](occlusion-culling-dynamic-gameobjects.md) | Enable or disable Dynamic Occlusion. |
| [Create high-precision occlusion areas](class-OcclusionArea.md) | Use the Occlusion Area component to define an area where Unity calculates culling more precisely. |
| [Control occlusion in areas with Occlusion Portals](class-OcclusionPortal.md) | Turn occlusion on and off through an object, for example a door that opens and closes. |
| [Occlusion Culling window reference](occlusion-culling-window.md) | Explore the properties and settings in the Occlusion Culling window to customize how culling works. |
| [Configure culling with the CullingGroup API](CullingGroupAPI-landing.md) | Integrate your own systems into Unity’s culling and **level of detail** (LOD) pipelines. |
| [Troubleshooting occlusion culling](occlusion-culling-troubleshooting.md) | Solve common problems when using occlusion culling. |

## Additional resources

* [Graphics performance and profiling](graphics-performance-profiling.md)
* [Enable GPU occlusion culling in URP](urp/gpu-culling.md)