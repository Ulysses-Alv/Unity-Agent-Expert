# Reduce shader variants in custom shader code

Prevent [shader variants](shader-variants.md) from being compiled in your custom shaders, to greatly reduce build times, file size, **shader** loading times, and runtime memory usage.

You can also limit or remove shader variants using the Editor UI. For more information, refer to [Reduce shader variants](shader-variant-stripping.md).

## Limit shader variants when you declare shader keywords

The way that you declare shader keywords can limit the number of variants that they produce:

* Use `shader_feature` instead of `multi_compile` where possible - see [How Unity compiles branching shaders](shader-conditionals-choose-a-type.md).
* Ensure that you don’t define unused keywords with `multi_compile`.

For information on declaring keywords in hand-coded shaders, see [Declaring and using shader keywords in HLSL](SL-MultipleProgramVariants.md).

## Indicate which shader keywords affect which shader stage

When you [declare a keyword](SL-MultipleProgramVariants-declare.md), Unity assumes all stages of the shader contain conditional code for that keyword.

You can add the following suffixes to indicate that only certain stages contain conditional code for a keyword, so that Unity doesn’t generate unneeded shader variants.

* `_vertex`
* `_fragment`
* `_hull`
* `_domain`
* `_geometry`

For example, use `#pragma shader_feature_fragment RED GREEN BLUE` to indicate that you use the 3 keywords to create conditional code in the fragment stage only.

**Note**: You can’t add these suffixes to `#pragma dynamic_branch` because `dynamic_branch` doesn’t create variants.

These suffixes might behave differently or have no effect, depending on the [graphics API](../ScriptReference/Rendering.GraphicsDeviceType.md). For example:

* The suffixes have no effect on OpenGL, OpenGL ES or Vulkan.
* The `_geometry` and `_raytracing` suffixes have no effect on Metal. Metal treats `_vertex`, `_hull` and `_domain` as a single stage.

## Use preprocessor macros to limit variants by platform

You can create conditional shader code using a [target platform preprocessor macro](shader-branching-built-in-macros.md), so you can limit variants on platforms with limited memory.

For example:

```
#ifdef SHADER_API_DESKTOP
   #pragma multi_compile _ RED GREEN BLUE WHITE
#else
   #pragma shader_feature RED GREEN BLUE WHITE
#endif
```

The code sample does the following:

* If you build for a [`SHADER_API_DESKTOP` platform](shader-branching-built-in-macros.md), Unity builds variants for every possible keyword combination.
* If you build for another platform, Unity builds only variants for keyword combinations the Materials in your build use.

You can use target platform preprocessor macros to select between `shader_feature`, `multi_compile` and `dynamic_branch`. For more information on when to use each type of conditional, see [Shader Conditionals](shader-conditionals-choose-a-type.md).

## Create user-controlled quality settings

When you build for console and mobile platforms that have limited memory, you can limit shader variants by only allowing users to switch between a small number of quality settings.

For example, if you use the keywords `DYNAMIC_LIGHTING`, `SOFT_SHADOWS` and `HIGH_QUALITY_LIGHTMAPS`, you can create the following:

* A ‘low quality’ setting that turns on `DYNAMIC_LIGHTING`.
* A ‘high quality’ setting that turns on `DYNAMIC_LIGHTING`, `SOFT_SHADOWS` and `HIGH_QUALITY_LIGHTMAPS`.

This means Unity won’t create shader variants for `DYNAMIC_LIGHTING` when it’s off, or the many different combinations of the 3 keywords being on and off.

### Reduce user-controlled quality settings by platform

You can use [target platform preprocessor macros](shader-branching-built-in-macros.md) to conditionally create fewer quality settings and fewer variants on platforms with limited memory. For example the following code sample will allow users to switch between 8 permutations of settings on `SHADER_API_DESKTOP` platforms, but only 2 on `SHADER_API_MOBILE` platforms.

```
#if SHADER_API_DESKTOP
   #pragma multi_compile SHADOWS_LOW SHADOWS_HIGH
   #pragma multi_compile REFLECTIONS_LOW REFLECTIONS_HIGH
   #pragma multi_compile CAUSTICS_LOW CAUSTICS_HIGH
#elif SHADER_API_MOBILE
   #pragma multi_compile QUALITY_LOW QUALITY_HIGH
   #pragma shader_feature CAUSTICS // Uses shader_feature, so Unity strips variants that use CAUSTICS if there are no Materials that use the keyword at build time.
#endif
```

## Additional resources

* [Reduce shader variants](shader-variant-stripping.md)