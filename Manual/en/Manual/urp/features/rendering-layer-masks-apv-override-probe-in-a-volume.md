# Override rendering layer masks in a volume

Override or combine rendering **layer masks** in a local volume to handle light leaks in cases such as doorways or shared walls.

A probe can belong to one or more rendering layers. When Unity bakes lighting, it uses ray tests and occlusion analysis to evaluate the visibility of geometry tagged with rendering layers. Unity then assigns probes to the appropriate layers.

During baking, probes cast rays that can hit multiple rendering layers. Unity assigns each probe a single mask from the four defined in the **Lighting** window. Unity selects the mask that best matches the rendering layers most frequently hit by its rays. You can override this automatic assignment with a Probe Adjustment Volume.

## Change or combine rendering layer masks in a bounded area

To change or combine **rendering layer masks** in a bounded area:

1. Create a [probe adjustment volume](../probevolumes-use.md) in the **scene**.

   When multiple adjustment volumes overlap, smaller volumes take precedence.
2. In the probe adjustment volume ****Inspector****, set **Mode** to **Override Rendering Layer Mask**.
3. In the **Operation** tab, select one of the following options:

   * Select **Override** to replace the probe’s existing rendering layer masks in the volume with the selected mask. Use this for strict separation between spaces.
   * Select **Add** to combine the selected mask with the probe’s existing masks. Use this for a transitional zone between spaces.
   * Select **Remove** to subtract the selected mask from the probe’s existing masks. Use this to exclude a boundary region without fully redefining the area.
4. In the **Rendering Layer Mask** dropdown, select the rendering layer mask to apply.
5. Do one of the following:

   * To bake only the probes affected by the volume, select **Preview Probe Adjustments**.
   * To bake all probes in the scene, select **Bake Probe Volumes**.

After you preview or bake, Unity updates affected probes and applies the mask changes.

**Note:** If no probes are in the same layer as a given **pixel**, all probes have the same priority, and Unity samples baked probes as if you deactivated the rendering layer feature. The **Quality** [leak reduction mode](../probevolumes-options-override-reference.md) ensures Unity samples only valid probes.

## Additional resources

* [Troubleshooting Adaptive Probe Volumes](../probevolumes-fixissues.md)
* [Adaptive Probe Volumes panel reference](../probevolumes-lighting-panel-reference.md)
* [Light Probe validity in Adaptive Probe Volumes in URP](../probevolumes-light-probe-validity.md)