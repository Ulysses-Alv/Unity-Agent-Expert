# Troubleshooting shadows in URP

Techniques for troubleshooting shadow rendering in the Universal **Render Pipeline** (URP).

## Adjust the shadow bias settings in URP

By adjusting the shadow bias values you can reduce or eliminate such shadow artifacts as shadow acne, shadow detachment (also known as peter-panning), light leaking, and self-shadowing.

In URP, you can set the shadow bias in the **URP Asset** and on individual lights.

In the **URP Asset**, the shadow bias settings are in the **Shadows** section. The values you set here are the default values for all lights in the **scene**.

To set the shadow bias values for a specific light:

1. In the **Shadows** section of the **Light** component, ensure that **Shadows Type** is set to **Soft Shadows** or **Hard Shadows**.
2. Set the **Bias** property to **Custom**. Unity shows the properties **Depth** and **Normal**. Use those properties to set the shadow bias values for the current light.

Using high shadow bias values might result in light leaking through meshes. This is where there is a visible gap between the shadow and its caster, and leads to shadow shapes that do not accurately represent their casters.

## Avoid shadow flickering caused by maximum light limit in URP

URP has the [per-camera limit](lighting/light-limits-in-urp.md) on the number of visible lights in the **camera** frustum. For example, on mobile platforms, this limit is 1 Main Light and 32 Additional Lights.

In the Forward **Rendering Path**, there is also the [per-object light limit](lighting/light-limits-in-urp.md).

If the number of visible lights exceeds the limit, Unity disables some lights. It might disable a real-time shadow-casting light, which causes shadow flickering.

To avoid this issue, ensure that the number of visible lights in the camera frustum does not exceed the limit. You can get the number of visible lights using the `cullResults.visibleLights` array of the [RenderingData.cullResults](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.2/api/UnityEngine.Rendering.Universal.RenderingData.html#UnityEngine_Rendering_Universal_RenderingData_cullResults) API.

The following example shows how to use the `cullresults.visibleLights` array:

```
public override void AddRenderPasses(ScriptableRenderer renderer, ref RenderingData renderingData)
{
    int visibleLightCount = renderingData.cullResults.visibleLights.Length;
}
```