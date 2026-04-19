# Preview baked lighting

You can edit a **scene** and preview changes to **lightmaps** and lighting without affecting the baked lightmaps.

To preview, do the following in a scene that uses lightmapping:

1. In the [Scene view View Options toolbar](overlays-reference-view-options.md), select **Debug Draw Mode**.
2. Select a [Global Illumination](GIVis.md#global-illumination) Draw Mode.
3. In the **Lighting Visualization** overlay, set **Lighting Data** to **Preview**.
4. Make changes to the scene that affect lightmaps.

Unity creates temporary lightmaps and uses them in the **Scene view**. To check the temporary lightmaps, go to the [Lighting window](lighting-window.md) and select **Baked Lightmaps**.

To use the baked lightmaps without your changes, set **Lighting Data** to **Baked**.