# Forward rendering path in the Built-In Render Pipeline

The Forward **rendering path** renders each **GameObject** in one or more render passes, depending on the lights that affect the object.

Lights themselves are also treated differently by the **Forward rendering** path, depending on their settings and intensity. For more information, refer to [Per-pixel and per-vertex lights](PerPixelLights-BuiltIn.md).

## Render passes

For each GameObject, Unity first renders the Base Pass, which renders the following:

* One per-pixel light that affects the GameObject.
* All the per-vertex and spherical harmonics (SH) lights that affect the GameObject.
* **Lightmap** data for the GameObject
* Ambient lighting
* Emissive lighting
* Shadows from directional lights

**Note:** Lightmapped objects don’t receive lighting from SH lights.

Unity then renders one Additional Pass for each per-pixel light that affects the GameObject. Unity doesn’t render shadows for these lights.

## Additional resources

* [Introduction to rendering paths](rendering-paths-introduction.md)
* [Per-pixel and per-vertex lights](PerPixelLights-BuiltIn.md)