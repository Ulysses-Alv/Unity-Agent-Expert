# Select the CPU or GPU for baking

To generate baked ****lightmaps**** and ****Light Probes****, you can choose between two backends for the [Progressive Lightmapper](progressive-lightmapper.md):

* The [Progressive GPU Lightmapper](GPUProgressiveLightmapper.md) default backend uses your computer’s GPU and Video Ram (VRAM).
* The Progressive CPU **Lightmapper** backend uses your computer’s CPU and system RAM.

## Hardware and software requirements

In order to use the Progressive GPU Lightmapper, your computer must meet these minimum specifications:

* At least one GPU with OpenCL 1.2 support
* At least 2GB of VRAM dedicated to that GPU
* A CPU that supports SSE4.1 instructions

If the **scene** you are baking requires more VRAM than is available on the designated GPU, bake times can significantly increase.

The Progressive GPU Lightmapper does not support OpenCL CPU devices.

The Apple silicon version of the Unity Editor is not compatible with the CPU Progressive Lightmapper. However, it is compatible with the [Progressive GPU Lightmapper](GPUProgressiveLightmapper.md).

## Render pipeline support

Refer to [render pipeline feature comparison](render-pipelines-feature-comparison.md) for more information about support for the Progressive Lightmapper across **render pipelines**.

## Select the Progressive Lightmapper

To select the Progressive Lightmapper:

1. Go to **Window** > **Rendering** > **Lighting**
2. Navigate to **Lightmapping Settings**
3. Set **Lightmapper** to **Progressive CPU** or **Progressive GPU**

You can perform many of the functions available in this window via **scripts**, using the [LightmapEditorSettings](../ScriptReference/LightmapEditorSettings.md) and [Lightmapping](../ScriptReference/Lightmapping.md) APIs.