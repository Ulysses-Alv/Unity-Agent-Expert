# Adaptive Probe Volume Inspector reference

Select an Adaptive Probe Volume and open the **Inspector** to view its properties.

## Mode

| **Option** | **Description** |
| --- | --- |
| **Mode** | Sets the size of the Adaptive Probe Volume. The options are:  * **Global**: URP sizes this Adaptive Probe Volume to include all renderers in the scene or Baking Set that have **Contribute Global Illumination** enabled in their [Mesh Renderer component](https://docs.unity3d.com/Manual/class-MeshRenderer.html). URP recalculates the volume size every time you save or generate lighting. * **Scene**: URP sizes this Adaptive Probe Volume to include all renderers in the same scene as this Adaptive Probe Volume. URP recalculates the volume size every time you save or generate lighting. * **Local**: Set the size of this Adaptive Probe Volume manually. |
| **Size** | Set the size of this Adaptive Probe Volume. This setting only appears when you set **Mode** to **Local**. |

## Subdivision Override

| **Property** | **Description** |
| --- | --- |
| **Override Probe Spacing** | Override the Probe Spacing set in the **Baking Set** for this Adaptive Probe Volume. This cannot exceed the **Min Probe Spacing** and **Max Probe Spacing** values in the [Adaptive Probe Volumes panel in the Lighting window](probevolumes-lighting-panel-reference.md). |

## Geometry Settings

| **Property** | **Description** |
| --- | --- |
| **Override Renderer Filters** | Enable filtering by Layer which **GameObjects** URP considers when it generates probe positions. Use this to exclude certain GameObjects from contributing to Adaptive Probe Volume lighting. |
| ****Layer Mask**** | Filter by Layer which GameObjects URP considers when it generates probe positions. |
| **Min Renderer Size** | The smallest [Renderer](https://docs.unity3d.com/ScriptReference/Renderer.html) size URP considers when it generates probe positions. |
| **Fill Empty Spaces** | Enable URP filling the empty space between and around Renderers with bricks. Bricks in empty spaces always use the **Max Probe Spacing** value. |