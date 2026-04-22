# Set up your scene and lights for baking

Select **Window** > **Rendering** > **Lighting** from the Unity Editor menu to open the Lighting window. Make sure any Mesh you want to apply a lightmap to has proper UVs for lightmapping. The easiest way to do this is to open the [Mesh import settings](FBXImporter-Model.md) and enable the **Generate Lightmap UVs** setting.

Next, to control the resolution of the **lightmaps**, go to the **Lightmapping Settings** section and adjust the **Lightmap Resolution** value.

To be included in your lightmap, Renderers must meet the following criteria:

* Have a **Mesh Renderer** or **Terrain** component
* Be marked as [**Contribute GI**](StaticObjects.md)
* Use a built-in Unity Material or **shader** that supports lightmaps, or a shader in the Built-In **Render Pipeline** with a [Meta Pass](MetaPass.md).

You can adjust settings for Lights in the [Light Explorer](LightingExplorer.md). To open the Light Explorer, go to **Window** > **Rendering** > **Light Explorer**.

### Static versus Dynamic

No matter which **Global Illumination** system you use, Unity will only consider objects that are marked as [“Contribute GI”](https://docs.unity3d.com/Manual/StaticObjects.html) during the baking/precomputing of the lighting. Dynamic (i.e. non-static) objects have to rely on the Light Probes you placed throughout the **scene** to receive indirect lighting.

Because the baking/precomputing of the lighting is a relatively slow process, only large and complex assets with distinct lighting variations, such as concavity and self-shadowing, should be tagged as “Contribute GI”. Smaller and convex meshes that receive homogeneous lighting should not be marked as such, and they should, therefore, receive indirect lighting from the [Light Probes](https://docs.unity3d.com/Manual/LightProbes.html) which store a simpler approximation of the lighting. Larger dynamic objects can rely on [LPPVs](https://docs.unity3d.com/Manual/class-LightProbeProxyVolume.html), in order to receive better localized indirect lighting. Limiting the number of objects tagged as “Contribute GI” in your scene is absolutely crucial to minimize baking times while maintaining an adequate lighting quality.