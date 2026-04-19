# Add ambient occlusion in the Built-In Render Pipeline

**Ambient occlusion** (AO) is a feature that simulates the soft shadows that occur in creases, holes, and surfaces that are close to each another. These areas occlude (block out) **ambient light**, so they appear darker.

It works by approximating how much ambient light can hit a point on a surface. It then darkens creases, holes and surfaces that are close to each other.

You can use ambient occlusion to add realism to your lighting.

## Baked Ambient Occlusion

If **Baked **Global Illumination**** is enabled in your **Scene**, Unity can bake ambient occlusion into the **lightmaps**. This is known as **Baked Ambient Occlusion**.

To enable baked ambient occlusion in your Scene:

1. Open the **Lighting** window (menu: **Window** > **Rendering** > **Lighting**)
2. Navigate to the **Mixed Lighting** section
3. Enable **Baked Global Illumination**
4. Navigate to the **Lightmapping Settings** section
5. Enable **Ambient Occlusion**

## Realtime ambient occlusion

If Global Illumination is not enabled in your Scene but you still want the effect of ambient occlusion, you can use a **post-processing** effect to apply real-time ambient occlusion to your Scene.

If ****Enlighten** Realtime Global Illumination** is enabled in your scene, the resolution for indirect lighting does not capture fine details or dynamic objects. We recommend using a real-time ambient occlusion post-processing effect, which has much more detail and results in higher quality lighting.

For information on real-time ambient occlusion post-processing effects, see [post-processing effects](PostProcessingOverview.md).

## Additional resources

* [Screen space ambient occlusion in URP](urp/post-processing-ssao-landing.md)