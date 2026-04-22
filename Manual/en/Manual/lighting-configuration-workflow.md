# Lighting configuration workflow

To set up lighting in Unity, follow these steps:

1. [Choose a render pipeline](#choose-a-render-pipeline)
2. [Configure lighting](#configure-lighting)
3. [Fine-tune your scene lighting](#fine-tune-your-scene-lighting)

## Choose a render pipeline

Unity provides **render pipelines** that differ in customization and lighting features:

* [Built-in Render Pipeline](lighting-birp.md) (not scriptable)
* [Universal Render Pipeline (URP)](urp/lighting/lighting-in-urp.md)
* [High-Definition Render Pipeline (HDRP)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest?subfolder=/manual/Light-Component.html)
* Custom Scriptable Render Pipeline (SRP)

For more information on render pipeline selection, refer to [choose a render pipeline](choose-a-render-pipeline.md).

## Configure lighting

1. Choose baked GI, realtime GI, mixed baked and realtime GI, or opt for no GI.

   For more information, refer to [Lighting Settings Asset Inspector window reference](class-LightingSettings.md#MixedLighting)
2. Choose one of the following Lighting Modes:

   * Baked Indirect
   * Subtractive
   * Shadowmask
   * Distance Shadowmask

   For more information, refer to [Lighting Mode](lighting-mode.md).

## Fine-tune your scene lighting

To fine-tune your **scene** lighting, follow these tasks based on project requirements:

1. Add [baked, realtime, or mixed lights](LightModes-introduction.md).
2. Optionally configure emissive surfaces with [Baked GI or Realtime GI](class-LightmapParameters.md).
3. Add baked, realtime, or custom [Reflection Probes](ReflectionProbes.md).
4. If a GI mode is set, add [Light Probes](LightProbes-landing.md). You can also add [Light Probe Proxy Volumes (LPPVs)](LightProbeProxyVolume-landing.md).

## Additional resources

* [Add light emission to a material](StandardShaderMaterialParameterEmission.md)
* [Reflection Probe Inspector window reference](class-ReflectionProbe.md)
* [Light Probes](LightProbes.md)
* [SRP Core](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@latest)