# Introduction to Realtime Global Illumination using Enlighten

Unity uses a middleware solution called **Enlighten** for Realtime **Global Illumination**.

By default, [Realtime Lights](LightModes-introduction.md#realtime) contribute only direct lighting to a **Scene**. If you enable **Realtime Global Illumination** (Enlighten Realtime Global Illumination) in your Scene, Realtime Lights also contribute indirect lighting to a Scene.

## Render pipeline support

See [render pipeline feature comparison](render-pipelines-feature-comparison.md) for more information about support for Realtime Global Illumination using Enlighten across **render pipelines**.

## When to use Enlighten Realtime Global Illumination

Enlighten Realtime Global Illumination (Realtime GI) is useful for Lights that change slowly and have a significant visual impact on your Scene, such as the sun moving across the sky, or a slowly pulsating light in a closed corridor. This feature is not intended for special effects or Lights that change quickly, because latency and the number of CPU cycles needed make that sort of application impractical. Enlighten Realtime Global Illumination is suitable for games targeting mid-level to high-end PC systems and consoles. Some high-end mobile devices may also be powerful enough to make use of this feature, but you should keep Scenes small and the resolution for real-time **lightmaps** low to ensure acceptable performance.

## How Enlighten Realtime Global Illumination works

Enlighten Realtime Global Illumination does the following:

1. Splits the scene into small surface patches (clusters).
2. Determines the degree to which the patches are visible to each other.
3. Groups patches visible to each other into systems.

At runtime, Enlighten Realtime Global Illumination uses this precomputed visibility information to approximate how Realtime Lights bounce in the Scene, saves the results in a set of lightmaps, and then uses these lightmaps to apply indirect lighting to the Scene.
It is computationally intensive to update the lightmaps, and so the process is split across several frames. It takes Enlighten Realtime Global Illumination several frames to propagate changes to indirect lighting throughout the Scene.