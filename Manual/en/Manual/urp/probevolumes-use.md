# Use Adaptive Probe Volumes

This page provides the basic workflow you need to use Adaptive Probe Volumes in your project.

## Add and bake an Adaptive Probe Volume

### Enable Adaptive Probe Volumes

1. In the main menu, go to **Edit** > **Project Settings** > **Quality**.
2. In the **Rendering** section, double-click the active ****Render Pipeline** Asset** to open it in the ****Inspector**** window.
3. In the **Inspector** window, go to **Lighting** > **Light Probe Lighting**. Set **Light Probe System** to **Adaptive Probe Volumes**.

### Add an Adaptive Probe Volume to the Scene

1. In the main menu, go to **GameObject** > **Light** > **Adaptive Probe Volume**.
2. In the **Inspector** window of the Adaptive Probe Volume, set **Mode** to **Global**.

   The Adaptive Probe Volume now covers your entire **Scene**.

### Adjust your light and Mesh Renderer settings

1. To include a light source in an Adaptive Probe Volume’s baked lighting data, open its **Inspector** window, go to **Light** > **General**, and set **Mode** to **Mixed** or **Baked**.
2. To include a GameObject in an Adaptive Probe Volume’s baked lighting data, open its **Inspector** window, go to **Mesh Renderer** > **Lighting**, and enable **Contribute Global Illumination**.
3. To make a GameObject receive baked lighting from probes, open its **Inspector** window, go to **Mesh Renderer** > **Lighting**, and set **Receive Global Illumination** to **Light Probes**.

### Bake your lighting

1. In the main menu, go to **Window** > **Rendering** > **Lighting**.
2. In the **Scene** tab, under **Mixed Lighting**, enable **Baked **Global Illumination****.
3. In the **Adaptive Probe Volumes** tab, under **Baking**, set the baking mode to **Single Scene**.
4. To bake your lighting, in the **Adaptive Probe Volumes** tab, do one of the following:

   * To bake all the lighting data of the scene, select **Generate Lighting**.
   * To bake Adaptive Probe Volumes only, in the **Generate Lighting** dropdown, select **Bake Probe Volumes**.

If no scene in the Baking Set contains an Adaptive Probe Volume, Unity asks if you want to create an Adaptive Probe Volume automatically.

You can change baking settings in the Lighting window’s [Lightmapping Settings](https://docs.unity3d.com/Documentation/Manual/class-LightingSettings.html#LightmappingSettings).

Refer to [Bake different lighting setups with Baking Sets](probevolumes-usebakingsets.md) for more information about Baking Sets.

If there are visual artefacts in baked lighting, such as dark blotches or light leaks, refer to [Fix issues with Adaptive Probe Volumes](probevolumes-fixissues.md).

## Configure an Adaptive Probe Volume

You can use the following to configure an Adaptive Probe Volume:

* Use the [Adaptive Probe Volumes panel](probevolumes-lighting-panel-reference.md) in the **Lighting** window to change the probe spacing and behavior in all the Adaptive Probe Volumes in a Baking Set.
* Use the settings in the [Adaptive Probe Volume **Inspector** window](probevolumes-inspector-reference.md) to change the Adaptive Probe Volume size and probe density.
* Add a [Probe Adjustment Volume component](probevolumes-adjustment-volume-component-reference.md) to the Adaptive Probe Volume, to make probes invalid in a small area or fix other lighting issues.
* Add a [Volume](set-up-a-volume.md) to your scene with a [Probe Volumes Options Override](probevolumes-options-override-reference.md), to change the way URP samples Adaptive Probe Volume data when the **camera** is inside the volume. This doesn’t affect baking.

## Additional resources

* [Bake multiple scenes together with Baking Sets](probevolumes-usebakingsets.md)
* [Change lighting at runtime](probe-volumes-change-lighting-at-runtime.md)
* [Fix issues with Adaptive Probe Volumes](probevolumes-fixissues.md)
* [Work with multiple Scenes in Unity](https://docs.unity3d.com/Documentation/Manual/MultiSceneEditing.html)