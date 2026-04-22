# Add per-instance properties to GPU instancing shaders in the Built-In Render Pipeline

By default, Unity GPU instances **GameObjects** with different [Transforms](class-Transform.md) in each instanced draw call. To add more variation to the instances, modify the **shader** to add per-instance properties such as color. You can do this both in **surface shaders** and in vertex/fragment shaders.

Custom shaders don’t need per-instance data, but they do require an instance ID because world matrices need one to function correctly. Surface shaders automatically set up an instance ID, but custom vertex and fragment shaders don’t. To set up the ID for custom vertex and fragment shaders, use UNITY\_SETUP\_INSTANCE\_ID at the beginning of the shader. For an example of how to do this, see [Vertex and fragment shader](gpu-instancing-vertex-fragment-shader-example.md).

When you declare an instanced property, Unity gathers all the property values from the MaterialPropertyBlock objects set on GameObjects into a single draw call. For an example of how to use MaterialPropertyBlock objects to set per-instance data at runtime, see [Changing per-instance data at runtime](gpu-instancing-change-data.md).

When adding per-instance data to multi-pass shaders, keep the following in mind:

* If a multi-pass shader has more than two passes, Unity only instances the first pass. This is because Unity renders later passes together for each object, which forces material changes.
* If you use the Forward **rendering path** in the Built-in **Render Pipeline**, Unity can’t efficiently instance objects that are affected by multiple lights. Unity can only use instancing effectively for the base pass, not for additional passes. For more information about lighting passes, see documentation on [Forward Rendering and Pass Tags](RenderTech-ForwardRendering.md).

When you use multiple per-instance properties, you don’t need to fill all of them in `MaterialPropertyBlock` objects. Also, if one instance lacks a property, Unity takes the default value from the referenced material. If the material doesn’t have a default value for the property, Unity sets the value to 0. Don’t put non-instanced properties in the `MaterialPropertyBlock`, because this disables instancing. Instead, create different materials for them.