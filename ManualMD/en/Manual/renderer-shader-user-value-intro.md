# Introduction to RSUV

Use the Renderer **Shader** User Value (RSUV) when managing many objects (for example, Mesh Renderers) that need visual customization for each instance. RSUV lets you specify per-renderer data to apply shader customizations while sharing a single material across multiple renderers. In Scriptable **Render Pipeline** (SRP) projects, this is the most performant approach.

## Render pipeline compatibility

If your project uses the Built-In Render Pipeline, you can use the Material Property Blocks (MPBs). However, this method isn’t recommended if your project uses an SRP, as it has low performance at runtime and might not work as expected.

With projects that use SRPs like the Universal Render Pipeline (URP) or the High Definition Render Pipeline (HDRP), which allow the use of the SRP batcher, you have two main options:

* Material duplication: Generate a dedicated material for each customized renderer. However, this process duplicates the entire material, and in many cases, the required customization per object is limited to only a small set of parameters.
* Renderer Shader User Value (RSUV): Specify per-renderer data to apply shader customizations on a single material used by multiple renderers.

**Note**: RSUV is only available in Scriptable Render Pipelines (SRPs). It is not supported in the Built-In Render Pipeline.

## RSUV functionality

RSUV is a custom 32-bit integer value that you can assign to each renderer to provide specific data to the shader. You can set the value through the following APIs:

* **Mesh** Renderer: [`MeshRenderer.SetShaderUserValue(uint value)`](../ScriptReference/MeshRenderer.SetShaderUserValue.md)
* Skinned **Mesh Renderer**: [`SkinnedMeshRenderer.SetShaderUserValue(uint value)`](../ScriptReference/SkinnedMeshRenderer.SetShaderUserValue.md)

You can access the value within shader HLSL code through the `unity_RendererUserValue` property. This functionality doesn’t interfere with batching.

With this approach, all customized renderers continue to share a single material, which produces a notable [performance improvement](#performance) in most scenarios.

## Use cases

The type of data you can convey through RSUV is flexible and depends on your visual customization scenario requirements. For example:

| **Requirement example** | **Recommended method** |
| --- | --- |
| Different diffuse colors per object | Encode an RGBA32 color directly into the RSUV and decode it in shader code. |
| Different shirt logo textures per character | For example, with 16 available logos, allocate 4 bits of RSUV to encode the UV index across a single texture atlas containing all logos. You can use the remaining 28 bits to encode additional variations, such as color adjustments. |
| Large amounts of custom data per renderer | Use the RSUV as an index into a larger data structure stored in a global GraphicsBuffer. |

**Note**: RSUV usage and encoding are shader-specific. For instance, one shader might interpret the RSUV as a color, while another may treat it as an index into auxiliary data.

## Performance

The RSUV feature introduces no additional CPU overhead. Using RSUV to customize values per renderer is always slightly faster than duplicating a material and significantly faster than using Material Property Block (MPB).

The performance gain is maximum when using [GPU Resident Drawer (GRD)](urp/reduce-rendering-work-on-cpu.md), as RSUV doesn’t require any material duplication, which avoids breaking GRD instancing draw calls.

## Additional resources

* [Set and use the RSUV](renderer-shader-user-value-set-and-use.md)
* [RSUV sample](renderer-shader-user-value-sample.md)