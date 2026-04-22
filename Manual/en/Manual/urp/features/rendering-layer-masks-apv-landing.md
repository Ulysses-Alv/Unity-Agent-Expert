# Prevent light leaks with rendering layer masks

Use rendering **layer masks** with [Adaptive Probe Volumes (APVs)](../probevolumes.md) to prevent light leaks across boundaries such as walls and ceilings. Restrict the probes that a renderer samples to maintain distinct lighting between adjacent spaces, even when the density of **light probes** is low.

| Topic | Description |
| --- | --- |
| [APV light leak prevention with rendering layer masks](rendering-layer-masks-apv-prevent-light-leaks.md) | Learn why APVs might produce light leaks between adjacent spaces and how **rendering layer masks** can solve the issue. |
| [Configure APV rendering layer masks and verify probe assignments](rendering-layer-masks-apv-configure-and-verify-probe-assignments.md) | Learn how to configure APV rendering layer masks and assign probes. |
| [Override rendering layer masks in a volume](rendering-layer-masks-apv-override-probe-in-a-volume.md) | Override or combine APV rendering layer masks in local adjustment volumes. |

## Additional resources

* [Rendering Layers in URP](rendering-layers.md)
* [Troubleshooting Adaptive Probe Volumes](../probevolumes-fixissues.md)
* [Adaptive Probe Volumes panel reference](../probevolumes-lighting-panel-reference.md)
* [Probe Adjustment Volume component reference](../probevolumes-adjustment-volume-component-reference.md)