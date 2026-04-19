# Troubleshooting objects appearing unlit or darker than intended

Fix issues causing some objects to appear unlit or darker than intended.

![ Three images depicting the same scene of a room filled by a 3 by 3 grid of spheres, baked in Shadowmask Lighting Mode. The first image shows that some of the spheres are unlit and appear dark compared to others. The second image shows the same scene in the Light Overlap debug view to highlight overlapping lights in red. The third image shows the same scene in the Shadowmask debug view.](../uploads/Main/ex-mixed-lights-render-as-baked.png)

Three images depicting the same scene of a room filled by a 3 by 3 grid of spheres, baked in Shadowmask Lighting Mode. The first image shows that some of the spheres are unlit and appear dark compared to others. The second image shows the same scene in the Light Overlap debug view to highlight overlapping lights in red. The third image shows the same scene in the Shadowmask debug view.

## Symptoms

Mixed lights in the **scene** are rendered as **baked lights** and some objects and areas in the scene appear unlit or darker than intended.

## Cause

This issue can occur when Unity converts mixed lights to baked lights during the lighting setup. Baked lights don’t provide real-time lighting, so they don’t illuminate dynamically in the scene. Instead, their contribution is baked into **lightmaps**, **light probes**, or **reflection probes**. If you don’t properly configure the baked lighting setup, objects may remain dark due to missing or incomplete lighting data during the bake process.

In [Shadowmask](lighting-mode.md#shadowmask) Lighting Mode, only four mixed lights can overlap. Unity converts lights that exceed this limit into baked lights.

In [Subtractive](lighting-mode.md#subtractive) Lighting Mode, no lights (except for directional lights) cast real-time shadows. This is a limitation of this lighting mode.

When you use baked lighting, real-time adjustments (such as changing the light’s color, intensity, or position during runtime) are no longer possible, which may limit dynamic or interactive lighting effects in the scene and cause the objects to appear darker or lit in unintended ways.

## Resolution - Change the light placement and radius

To verify that light overlap exceeding the limit is the main issue, switch to the [Light Overlap](GIVis.md#light-overlap-baked-global-illumination-only) **Scene** view draw mode after generating lighting. This draw mode highlights overlapping lights in red so you can identify the number of overlapping lights.

Move the affected lights so that they no longer overlap. You can also adjust the **Range** or **Spot Angle** properties in the [Light](class-Light.md) component. Generate lighting again after making the changes to check if this fixes the issue.

## Resolution - Switch to a different Lighting Mode

Each [Lighting Mode](lighting-mode.md) has different advantages and limitations. If your project will be lit using a single directional light, then consider the **Subtractive** Lighting Mode instead of **Shadowmask**. Refer to [Lighting Modes](lighting-mode.md) to view the different limits and ways Unity calculates lighting for each mode.