# Troubleshooting light leaks in Adaptive Probe Volumes in URP

Light leaks are areas that are too light or dark, often in the corners of a wall or ceiling.

![A light leak.](../../uploads/urp/probe-volumes/probevolumes-lightleak.JPG)

Light leaks often occur when geometry receives light from a **Light Probe** that isn’t visible to the geometry, for example because the Light Probe is on the other side of a wall. Adaptive Probe Volumes use regular grids of Light Probes, so Light Probes might not follow walls or be at the boundary between different lighting areas.

To fix light leaks, you can do the following:

* [Create thicker walls](#thickerwalls).
* [Add an Adaptive Probe Volumes Options override to your scene](#volume).
* [Enable Rendering Layers](#layers).
* [Adjust Baking Set properties](#probevolumesettings).
* [Use a Probe Adjustment Volume](#probevolumeadjustment).

### Create thicker walls

Adjust walls so their width is closer to the distance between probes in the local [brick](probevolumes-concept.md#how-probe-volumes-work)

### Add an Adaptive Probe Volumes Options override to your scene

You can add a [Volume](volumes-landing-page.md), then add an **Adaptive Probe Volumes Options** override to the Volume. This adjusts the position that **GameObjects** use to sample the Light Probes.

1. Add a [Volume](set-up-a-volume.md) to your **scene** and make sure its area overlaps the **camera** position.
2. Select **Add Override**, then select **Lighting** > **Adaptive Probe Volumes Options**.
3. Enable **Normal Bias**, then adjust the value to move the position that GameObject pixels use to sample the Light Probes, along the **pixel**’s surface normal.
4. Enable **View Bias**, then adjust the value to move the position that GameObject pixels use to sample the Light Probes, towards the camera.
5. Disable and enable **Leak Reduction Mode** to check if it improves light leaks.

Volumes only affect the scene if the camera is near or inside the volume. Refer to [Understand volumes](Volumes.md) for more information.

Refer to [Probe Volumes Options Override reference](probevolumes-options-override-reference.md) for more information on **Adaptive Probe Volumes Options** settings.

### Use Rendering Layer Masks

You can configure the **Rendering **Layer Masks**** in the [Adaptive Probe Volumes panel](probevolumes-lighting-panel-reference.md) in the Lighting window. This allow APV to assign a **Rendering Layer Mask** to each Light Probe.

For performance reasons, Adaptive Probe Volumes only supports up to 4 Rendering Layers Masks. You can use the list to create a new mask and use the dropdown to assign it any Rendering Layer.
When lighting is generated, Unity will try to automatically assign a mask to each probe by looking at the Rendering Layer Masks of objects surrounding the probe. Additionally, you can use a **Probe Adjustment Volume** to override the Rendering Layer Mask assigned to Light Probes.

At runtime, renderers will only sample lighting data from probes that have a matching Rendering Layer Mask. If the object doesn’t match any of the Masks defined in the Lighting window, it will sample lighting from all the valid surrounding probes.
Note that this feature requires **Use Rendering Layers** to be enabled in the URP Asset.

For example, in order to fix light leaking issues, you can create an Interior and an Exterior Rendering Layer Mask to ensure interior objects will never sample lighting data from exterior probes and fix light leaking through the walls.
A renderer can have several Rendering Layers enabled in it’s Rendering Layer Masks. This is useful when dealing with dynamic objects that may want to sample lighting from both the exterior and interior probes.

You can visualize which layers are assigned to a probe:
- Go to the Probe Volumes tab
- Enable **Display Probes**
- Set the **Probe Shading Mode** field to **Rendering Layer Masks**
- Use the toggles in the **Scene View** Overlay to hide Probes matching a Rendering Layer Mask

### Adjust Baking Set properties

If adding a Volume doesn’t work, use the [Adaptive Probe Volumes panel](probevolumes-lighting-panel-reference.md) in the Lighting window to adjust Virtual Offset and Dilation settings.

1. In **Probe Dilation Settings**, reduce **Search Radius**. This can help in situations where invalid Light Probes are receiving lighting data from more distant Light Probes. However, a lower **Search Radius** might cause light leaks.
2. In **Virtual Offset Settings**, reduce **Search Distance Multiplier** and **Ray Origin Bias**.
3. If there are light leaks in multiple locations, adjust **Min Probe Spacing** and **Max Probe Spacing** to increase the density of Light Probes.
4. Select **Generate Lighting** to rebake the scene using the new settings.

Note: Don’t use very low values for the settings, or Dilation and Virtual Offset might not work.

### Add a Probe Adjustment Volume component

Use a Probe Adjustment Volume component to make Light Probes invalid in a small area. This triggers Dilation during baking, and improves the results of **Leak Reduction Mode** at runtime.

1. In the Adaptive Probe Volume Inspector, select **Add Component**, then select **Light** > **Probe Adjustment Volume**.
2. Set the **Size** so the **Probe Adjustment Volume** area overlaps the Light Probes causing light leaks.
3. Set **Probe Volume Overrides** > **Mode** to **Invalidate Probes**, to invalidate the Light Probes in the Volume.
4. If you have a [Volume with a Probe Volumes Options override](#volume), enable **Leak Reduction Mode**.
5. In **Probe Volume Settings**, select **Generate Lighting** to rebake the scene using the new settings.

Using a Probe Adjustment Volume component solves most light leak issues, but often not all.

If you use many Probe Adjustment Volumes in a scene, your bake will be slower, and your scene might be harder to understand and maintain.

Refer to [Probe Adjustment Volume component reference](probevolumes-adjustment-volume-component-reference.md) for more information.

## Fix seams

Seams are artefacts that appear when one lighting condition transitions immediately into another. Seams are caused when two adjacent bricks have different Light Probe densities. Refer to [bricks](probevolumes-concept.md#how-probe-volumes-work) for more information.

![Two seams.](../../uploads/urp/probe-volumes/probevolumes-seams.JPG)

To fix seams, do the following:

1. Add a [Volume](set-up-a-volume.md) to your scene and make sure its area overlaps the position of the camera.
2. Select **Add Override**, then select **Lighting** > **Probe Volumes Options**.
3. Enable **Sampling Noise**, then try adjusting the value to add noise and make the transition more diffuse. Noise can help break up noticeable edges in indirect lighting at brick boundaries.

## Additional resources

* [Configure the size and density of Adaptive Probe Volumes](probevolumes-changedensity.md)
* [Adaptive Probe Volumes panel reference](probevolumes-lighting-panel-reference.md)
* [Probe Volumes Options Override reference](probevolumes-options-override-reference.md)
* [Probe Adjustment Volume component reference](probevolumes-adjustment-volume-component-reference.md)
* [Light management with rendering layer masks](features/rendering-layer-masks-apv-landing.md)