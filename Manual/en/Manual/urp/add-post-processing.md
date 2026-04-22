# Add post-processing in URP

New scenes in URP do not use **post-processing** by default. Instead you must manually add post-processing to any new **scene** which you create. You can then configure each effect individually to create the visual effect you want.

**Note**: Some examples and scene templates in URP use post-processing by default. If you use these to create a new scene, you might not need to make any changes.

## Add post-processing to a scene

To add post-processing to a scene:

1. Select a **Camera**, then in the **Inspector** window enable **Post Processing**.
2. Add a GameObject with a [Volume](Volumes.md) component in the scene. For example, select **GameObject** > **Volume** > **Global Volume**.
3. Select the **GameObject**, then in the **Volume** component select **New** to create a new [Volume Profile](Volume-Profile.md).
4. Select **Add Override**, then select a post-processing effect [Volume Override](VolumeOverrides.md), for example **Bloom**.

Now you can use the Volume Override to enable and adjust the settings for the post-processing effect.

**Note**: Post-processing effects from a volume apply to a camera only if a value in the **Volume Mask** property of the camera contains the layer that the volume belongs to.

## Configure individual post-processing effects

Each post-processing effect in URP has individual settings you can adjust to tune the visual impact they have on your scene. For more information on the post-processing effect settings, refer to the reference pages in the [Effect List](EffectList.md).