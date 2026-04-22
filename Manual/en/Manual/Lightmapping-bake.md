# Bake lighting

To generate **lightmaps** for your Scene:

1. Open the [Lighting window](lighting-window.md) (menu: **Window** > **Rendering** > **Lighting**)
2. At the bottom of the **Scene** tab on the Lighting window, select **Generate Lighting**.
3. A progress bar appears in Unity Editor’s status bar, in the bottom-right corner.

When lightmapping is complete, Unity’s Scene and Game views update automatically and you can view the resulting lightmaps by going to the **Baked Lightmaps** tab in the Lighting Window.

When you generate lighting, Unity adds [Lighting Data Assets](LightmapSnapshot.md), [baked lightmaps](lighting-window.md) and [Reflection Probes](ReflectionProbes.md) to the [Assets](SpecialFolders.md) folder.

## Bake lightmaps automatically

To set Unity to bake lightmaps automatically when you open a scene that has no lighting data, follow these steps:

1. Go to **Window** > **Rendering** > **Lighting**.
2. Set **Bake on Scene Load** to **If Missing Lighting Data**.

If you share your project with someone else, you can use this option to reduce the size of your project by not including lighting data. When a scene is opened by someone else, Unity calculates the missing lighting data.