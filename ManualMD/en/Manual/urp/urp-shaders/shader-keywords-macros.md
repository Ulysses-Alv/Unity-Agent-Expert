# Shader keywords and macros reference in URP

Shader keywords and macros that enable or provide access to URP features in **shaders**.

| **Element** | **Description** |
| --- | --- |
| `_CLUSTER_LIGHT_LOOP` | Use this `multi_compile` keyword to make the shader compatible with the Forward+ and Deferred+ **rendering paths**. For an implementation example, refer to [Render additional lights in a shader](../use-built-in-shader-methods-additional-lights-fplus.md). |
| `_ADDITIONAL_LIGHTS` | Use this keyword to define areas in shader code that Unity should execute if per-pixel additional lights are enabled in a **scene** and URP Asset. If a renderer uses the Forward+ rendering path, Unity ignores this keyword and uses the `_CLUSTER_LIGHT_LOOP` keyword instead. For an implementation example, refer to [Render additional lights in a shader](../use-built-in-shader-methods-additional-lights-fplus.md). |
| `LIGHT_LOOP_BEGIN` | Use this macro to iterate over the additional lights. In the Forward+ rendering path, the `LIGHT_LOOP_BEGIN` macro requires the following struct to be in its scope, both the type and the variable name must match this signature: `InputData inputData`. For an implementation example, refer to [Render additional lights in a shader](../use-built-in-shader-methods-additional-lights-fplus.md). |

## Additional resources

* [Changing how shaders work via branching and keywords](../../SL-MultipleProgramVariants.md)
* [Shader methods in URP](../use-built-in-shader-methods.md)
* [Custom lighting in URP](../lighting/custom-lighting-landing.md)