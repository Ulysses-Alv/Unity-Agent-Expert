# Per-pixel and per-vertex lights

If you use the default [Forward rendering path](rendering-paths-introduction.md#forward), each realtime Light component can be one of the following types:

* A per-pixel light, which lights each **pixel** of an object accurately.
* A per-vertex light, which lights each vertex of an object accurately. Unity interpolates lighting for the pixels between vertices.

Per-pixel lights give more accurate results but reduce performance.

The Built-In **Render Pipeline** also sets some lights as Spherical Harmonic (SH) per-vertex lights, which are the least accurate but render the fastest.

For more information, refer to the following:

* [Light limits in the Universal Render Pipeline (URP)](urp/lighting/light-limits-in-urp.md)
* [Per-pixel and per-vertex lights in the Built-In Render Pipeline](PerPixelLights-BuiltIn.md)

## Additional resources

* [Choose a lighting setup](choose-a-lighting-setup.md)
* [Introduction to rendering paths](rendering-paths-introduction.md)