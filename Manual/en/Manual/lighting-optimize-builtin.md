# Optimize lighting in the Built-In Render Pipeline

To optimize the lights in your **scene**, avoid Unity using multiple render passes to render **GameObjects**, or doing too much work to render lighting. This reduces the number of the draw calls the CPU sends, and the number of vertices and **pixels** the GPU processes.

Do the following:

* Use [lightmapping](Lightmapping-landing.md) to light static objects, instead of **realtime lights**.
* Avoid combining meshes if they’re lit by different realtime per-pixel lights, because Unity calculates every light for every combined **mesh**.
* Avoid lighting GameObjects with multiple per-pixel lights.

## Prioritize lights

To avoid lighting GameObjects with multiple per-pixel lights, prioritize which lights provide per-pixel lighting based on which GameObjects they light.

For example, in a driving game, prioritize the car headlights as a per-pixel light, but deprioritize the rear lights and distant lampposts.

To decrease the number of per-pixel lights, do any of the following:

* In the [Light component Inspector window](class-Light.md), set **Render Mode** to **Not important**. This sets the light as a per-vertex light or spherical harmonics (SH) light.
* In the [Quality settings window](class-QualitySettings.md), decrease **Pixel Light Count**. This excludes lights from the per-pixel light group.

To increase the number of per-pixel lights, do any of the following:

* In the [Light component Inspector window](class-Light.md), increase the **Intensity**. The brightest light is always a per-pixel light.
* In the [Light component Inspector window](class-Light.md), set **Render Mode** to **Important**. This sets the light as a per-pixel light.
* In the [Quality settings window](class-QualitySettings.md), increase **Pixel Light Count**. This includes more lights in the per-pixel light group.

## Disable per-vertex and spherical harmonics (SH) lights

To disable per-vertex and SH lights in a custom **shader**, add the `OnlyDirectional` tag to the Pass in your **ShaderLab** code. For more information, refer to [Pass tags in ShaderLab reference](SL-PassTags.md).

## Additional resources

* [Per-vertex and per-pixel lights](PerPixelLights.md)
* [Per-vertex and per-pixel lights in the Built-In Render Pipeline](PerPixelLights-BuiltIn.md)