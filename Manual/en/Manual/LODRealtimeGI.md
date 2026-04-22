# LOD and Enlighten Realtime Global Illumination

Objects with **Receive **Global Illumination**** set to ****Lightmaps**** have a lighting solution with lightmaps for baked direct and [indirect lighting](LightmappingDirectional.md), and lightmaps for **Enlighten** Realtime Global Illumination. For more information about **Receive Global Illumination**, see the [Mesh Renderer settings](class-MeshRenderer.md) and the script reference for [ReceiveGI](../ScriptReference/ReceiveGI.md).

When you use Unity’s **LOD** system in a **Scene** where you have enabled Baked Global Illumination and Enlighten Realtime Global Illumination, the system lights the most detailed model out of the **LOD Group** as if it has the **Contribute Global Illumination** setting enabled and isn’t part of the LOD group."

Enlighten can only compute direct lighting for lower LODs and the LOD system must rely on [Light Probes](LightProbes.md) to sample indirect lighting.

To enable the Editor to produce lightmaps with Enlighten Realtime Global Illumination, select the **GameObject** you want to affect, view its Renderer component in the **Inspector** window, and ensure that **Contribute Global Illumination** is enabled.

For lower LODs in a LOD Group, you can only combine baked lightmaps with Enlighten Realtime Global Illumination from [Light Probes](LightProbes.md) or [Light Probe Proxy Volumes](class-LightProbeProxyVolume.md), which you must place around the LOD Group.

![An animation showing how real-time ambient color affects the Enlighten Realtime Global Illumination used by lower LODs](../uploads/Main/LODRealtimeGI.gif)

An animation showing how real-time ambient color affects the Enlighten Realtime Global Illumination used by lower LODs