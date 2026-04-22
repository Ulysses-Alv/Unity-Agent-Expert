# Introduction to lighting data

Precomputed lighting data and baked lighting are not the same.

Baked lighting data is static. For example, the appearance of a **GameObject** can change as it moves around a **Light Probe** (if that GameObject has Contribute Global Illumination enabled). However, the data in the Light Probe itself does not change.
Unity can store baked lighting data in **lightmaps**, Light Probes, and Reflection Probes.

Enlighten Realtime **Global Illumination** does not rely on baked lighting. Instead, it uses precomputed visibility data to speed up the process of determining how a light moving in real-time affects the surfaces that its rays can reach.

## Generating lighting data

The Editor follows different steps to calculate **Enlighten** Realtime Global Illumination and Baked Global Illumination. The progress bar displays information about the current process. You can continue working in the Editor while the processes run.

The stages of lighting precomputation are listed below:

**Enlighten Realtime Global Illumination**

1. Create Geometry
2. Layout Systems
3. Create Systems
4. Create Atlas
5. Clustering
6. Visibility
7. Light Transport
8. Tetrahedralize Probes
9. Create ProbeSet

**Probes**

1. Ambient Probes
2. Baked/Realtime Ref. Probes

## Generating data

Unity’s precomputed lighting solutions only consider static geometry. To begin the lighting precompute process, you need at least one [Static](https://docs.unity3d.com/Manual/StaticObjects.html) **GameObject** in your ****Scene****.

When you initiate a precompute, the Unity Editor evaluates and computes all aspects of your Scene lighting. To recalculate and bake only [Reflection Probes](class-ReflectionProbe.md), select the **Generate Lighting** option **Reflection Probes** [Lighting Window](lighting-window.md).

## Before you build

Before building your game, generate the lighting data for all your Scenes, to ensure that you do not lose any lighting data.

When you generate lighting for your Scene, Unity saves your lighting data as Asset files in your project directory. This ensures that the data you need is part of your build.