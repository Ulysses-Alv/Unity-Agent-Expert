# Enable Enlighten Realtime Global Illumination

To enable Enlighten Realtime Global Illumination in your Scene, open the [Lighting window](lighting-window.md) (menu: **Window** > **Rendering** > **Lighting**) and enable **Realtime Global Illumination**.

To disable the effect of **Realtime GI** on a specific Light, select the Light **GameObject** and, in the Light component, set the **Indirect Multiplier** to 0. This means that the Light doesn’t contribute any indirect light to the **Scene**.

To disable **Realtime GI** altogether, open the Lighting window (menu: **Window** > **Rendering** > **Lighting**) and uncheck **Realtime Global Illumination**.

For detailed step-by-step advice on using **Enlighten** Realtime **Global Illumination**, see the Unity tutorial on [Precomputed Realtime GI](https://learn.unity.com/project/lighting-optimization-with-precomputed-realtime-gi).

## Light Probes and Enlighten Realtime Global Illumination

Note that [Light Probes](LightProbes.md) behave differently when you enable Enlighten Realtime Global Illumination.

In order to react to runtime changes in Scene lighting, they sample lighting iteratively at runtime.

When you disable Enlighten Realtime Global Illumination in a Scene, Light Probes only use baked lighting data. This means that they don’t react to runtime changes in Scene lighting.

## Shadows and Enlighten Realtime Global Illumination

If a Light also casts shadows, Unity renders both dynamic and static GameObjects in the Scene into the Light’s [shadow map](Shadows.md). The [Material Shaders](Shaders.md) of both static and dynamic GameObjects sample this shadow map so that these GameObjects cast real-time shadows onto each other. The **Shadow Distance** setting determines the maximum distance at which shadows begin to fade out and disappear entirely, which in turn affects performance and image quality.

Enlighten Realtime Global Illumination results also include soft shadows, unless the Scene is very small. These shadows are typically more coarse-grained than what lightmapping can achieve.

To modify **Shadow Distance** settings, navigate to **Edit** > **Project Settings** > **Quality** > **Shadows**.