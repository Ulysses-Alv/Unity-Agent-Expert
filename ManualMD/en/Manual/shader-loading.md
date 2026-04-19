# Shader loading

Unity loads [compiled shaders](shader-compilation.md) from your built application in the following way:

1. When Unity loads a **scene** or a [runtime resource](LoadingResourcesatRuntime.md), it loads all the compiled **shader** variants for the scene or resource into CPU memory.
2. By default, Unity decompresses all the shader variants into another area of CPU memory. You can [control how much memory shaders use on different platforms](shader-memory.md).
3. The first time Unity needs to render geometry using a shader variant, Unity passes the shader variant and its data to the graphics API and the graphics driver.
4. The graphics driver creates a GPU-specific version of the shader variant and uploads it to the GPU.

This approach ensures that Unity and the graphics driver avoid processing and storing all the shader variants on the GPU before Unity needs them. But there might be a visible stall when the graphics driver creates a GPU-specific shader variant for the first time.

Unity caches each GPU-specific shader variant, to avoid another stall when Unity needs the shader variant again.

Unity removes the shader variant completely from CPU and GPU memory when there are no longer any objects that reference the shader variant.

## Which shaders Unity loads

Unity only loads compiled shaders that are compatible with the platform’s graphics API, hardware and [graphics tier](graphics-tiers.md).

If Unity can’t find the shader variant it needs because you or Unity [stripped the shader variant](shader-variant-stripping.md) from your built application, Unity tries to select a similar shader variant. If Unity can’t find a similar shader variant, it uses the magenta [error shader](shader-error.md).

You can [enable strict shader variant matching](shader-how-many-variants.md#highlight-missing-shaders) to stop Unity trying to select a similar shader variant.

## How Unity selects a subshader

If a shader variant contains multiple [subshaders](shader-objects.md#subshaders), Unity tries to select and use a single subshader that’s compatible with all of the following:

* The platform’s hardware.
* The current [ShaderLab level of detail (LOD)](SL-ShaderLOD.md).
* The active **render pipeline**.

Unity searches for the first compatible subshader among the following elements in the following order:

1. The subshaders in the order they appear in the shader.
2. The subshaders in any fallback **shader objects**, in the order they appear in the shader objects.

If Unity can’t find a compatible subshader, it uses the magenta error shader.

You can set which subshaders are compatible with which hardware using **ShaderLab** tags. See [ShaderLab: assigning tags to a SubShader](SL-SubShaderTags.md).

## Check when Unity loads shaders

To check when a shader loads from serialized data, search for the following **profiler markers** in the [Profiler](Profiler.md):

* `Shader.ParseThreaded`
* `Shader.ParseMainThread`

To check when a shader unloads shader variants, search for the `Shader.MainThreadCleanup` profiler marker.

## Additional resources

* [How Unity loads and uses shaders](shader-loading.md)
* [Optimizing shader runtime performance](SL-ShaderPerformance.md)