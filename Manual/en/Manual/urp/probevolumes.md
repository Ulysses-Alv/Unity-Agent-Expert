# Adaptive Probe Volumes (APV) in URP

Adaptive Probe Volumes make [Light Probes](https://docs.unity3d.com/Manual/LightProbes.html) easier to use by automating placement. They also provide higher quality, more accurate lighting, because they light per-pixel not per-object.

| Topic | Description |
| --- | --- |
| [Introduction to Adaptive Probe Volumes](probevolumes-concept.md) | The purpose of Adaptive Probe Volumes and what you can do with them. |
| [Use Adaptive Probe Volumes](probevolumes-use.md) | Add Adaptive Probe Volumes to your project and configure them. |
| [Display Adaptive Probe Volumes](probevolumes-showandadjust.md) | Visualize the structure of Adaptive Probe Volumes. |
| [Configure the size and density of Adaptive Probe Volumes](probevolumes-changedensity.md) | Change the size of an Adaptive Probe Volume, or increase the density of Light Probes. |
| [Bake multiple scenes together with Baking Sets](probevolumes-usebakingsets.md) | Add **scenes** to a Baking Set so you can bake the lighting for all the scenes together. |
| [Light management with rendering layer masks](features/rendering-layer-masks-apv-landing.md) | Learn how to prevent light leaks across boundaries, even with a low light probe density. |
| [Optimize loading Adaptive Probe Volume data](probevolumes-streaming.md) | Stream lighting data to provide lighting for large open worlds, or load data from AssetBundles or Addressables. |
| [Changing lighting at runtime](probe-volumes-change-lighting-at-runtime.md) | Use Lighting Scenarios or sky occlusion to change how objects use the data in Adaptive Probe Volumes at runtime. |
| [Troubleshooting Adaptive Probe Volumes](probevolumes-fixissues.md) | Solve common issues with Adaptive Probe Volumes, such as light leaks and seams. |
| [Adaptive Probe Volume Inspector window reference](probevolumes-inspector-reference.md) | Reference for the Adaptive Probe Volume **Inspector** window. |
| [Adaptive Probe Volumes panel reference](probevolumes-lighting-panel-reference.md) | Reference for the Adaptive Probe Volumes panel in the Lighting settings. |
| [Probe Volumes Options Override reference](probevolumes-options-override-reference.md) | Reference for the Probe Volumes Options Override. |
| [Probe Adjustment Volume component reference](probevolumes-adjustment-volume-component-reference.md) | Reference for the Probe Adjustment Volume component. |

## Additional resources

* [Light Probes](https://docs.unity3d.com/Manual/LightProbes.html)
* [Light Probes for moving objects](https://docs.unity3d.com/Manual/LightProbes-MovingObjects.html)
* [Light Probe Group](https://docs.unity3d.com/Manual/class-LightProbeGroup.html)
* [Rendering Debugger](features/rendering-debugger.md)