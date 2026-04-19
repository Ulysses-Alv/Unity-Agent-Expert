# Bake data in multiple scenes

You can reduce the time taken to bake data and the size of the generated data if you bake your data into multiple **scenes**.

## Bake lightmaps with multiple scenes

To predetermine the brightness of surfaces in your scenes, you can bake [lightmaps](https://docs.unity3d.com/2022.2/Documentation/Manual/Lightmappers.html). There are two ways to bake lightmaps for multiple scenes at once:

* To automate your lightmap builds, use the [Lightmapping.BakeMultipleScenes](https://docs.unity3d.com/ScriptReference/Lightmapping.BakeMultipleScenes.html) function.
* To manually bake the lightmaps in multiple scenes, use the Unity Editor and perform the following steps:
  1. Open the scenes that you want to bake.
  2. Select **Generate Lighting**.

Shadows and GI light bounces work across all scenes, but the lightmaps and Realtime GI data loads and unloads separately for each scene. This means scenes don’t share and you can unload scenes with lightmaps safely. Scenes do share **Light Probe** data so all Light Probes for scenes baked together load at the same time.

## Bake occlusion culling data with multiple scenes

You can bake **occlusion culling** data for multiple scenes at once to determine which **GameObjects** render:
1. Open the scenes that you want to bake.
2. Go to the Occlusion Culling window (**Window** > **Rendering** > **Occlusion Culling**).
3. Select the Bake button.
4. Save the baked scenes to make the scene-to-occlusion-data reference persistent.

This saves the occlusion culling data into a single asset called `OcclusionCullingData.asset` in a folder matching the name of the current active scene.

If you load a scene additively and it has the same occlusion data reference as the active scene, then the static renderers and portals that cull information for that scene initialize from the occlusion data. This means the occlusion system performs as if static renderers and portals are baked into a single scene.

## Additional resources

* [Creating, loading, and saving Scenes](scenes-working-with.md)
* [Setup multiple scenes](setupmultiplescenes.md)
* [Use scripts to edit multiple scenes](scriptmultiplescenes.md)