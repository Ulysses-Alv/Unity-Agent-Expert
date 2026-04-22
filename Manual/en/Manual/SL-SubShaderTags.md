# SubShader tags in ShaderLab reference

This page contains information on using a `Tags` block in your **ShaderLab** code to assign tags to a SubShader.

For information on defining SubShader, see [ShaderLab: defining a SubShader](SL-SubShader.md). For information on how a **Shader** object works, and the relationship between Shader objects, SubShaders and Passes, see [Shader objects](shader-objects.md).

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: SubShader Tags block** | Yes | Yes | Yes | Yes |
| **ShaderLab: RenderPipeline SubShader tag** | Yes | Yes | No | No |
| **ShaderLab: Queue SubShader tag** | Yes | Yes | Yes  **Note:** in a custom SRP, you can define your own rendering order and choose whether or not to use render queues. For more information, see DrawingSettings and SortingCriteria. | Yes |
| **ShaderLab: RenderType SubShader tag** | Yes | Yes | Yes | Yes |
| **ShaderLab: DisableBatching SubShader tag** | Yes | Yes | Yes | Yes |
| **ShaderLab: ForceNoShadowCasting SubShader tag** | Yes | Yes  This disables regular shadows, but has no effect on contact shadows. | Yes | Yes |
| **ShaderLab: CanUseSpriteAtlas SubShader tag** | Yes | Yes | Yes | Yes |
| **ShaderLab: PreviewType SubShader tag** | Yes | Yes | Yes | Yes |

## Syntax

| **Signature** | **Function** |
| --- | --- |
| Tags { “[name1]” = “[value1]” “[name2]” = “[value2]”} | Applies the given tags to the SubShader.  You can define as many tags as you like. |

### RenderPipeline tag

| **Signature** | **Function** |
| --- | --- |
| “RenderPipeline” = “[name]” | Tells Unity whether this SubShader is compatible with URP or HDRP. |

| **Parameter** | **Value** | **Function** |
| --- | --- | --- |
| [name] | `UniversalPipeline` | This SubShader is compatible with URP only. |
|  | `HDRenderPipeline` | This SubShader is compatible with HDRP only. |
|  | (any other value, or not declared) | This SubShader is not compatible with URP or HDRP. |

### Queue tag

| **Signature** | **Function** |
| --- | --- |
| “Queue” = “[queue name]” | Use the named render queue. |
| “Queue” = “[queue name] + [offset]” | Use an unnamed queue, at a given offset from the named queue.  An example of when this is useful is in the case of transparent water, which you should draw after opaque objects but before transparent objects. |

| **Signature** | **Value** | **Function** |
| --- | --- | --- |
| [queue name] | Background | Specifies the Background render queue. |
|  | Geometry | Specifies the Geometry render queue. |
|  | AlphaTest | Specifies the AlphaTest render queue. |
|  | Transparent | Specifies the Transparent render queue. |
|  | Overlay | Specifies the Overlay render queue. |
| [offset] | integer | Specifies the index at which Unity renders the unnamed queue, relative to the named queue. |

### RenderType tag

| **Signature** | **Function** |
| --- | --- |
| “RenderType” = “[renderType]” | Set the RenderType value for this SubShader. |

| **Signature** | **Value** | **Function** |
| --- | --- | --- |
| [renderType] | String | There are no set values for this parameter. To identify the RenderType value for any SubShader that you want to replace, open its shader source file.  The `RenderType` SubShader tags for Unity’s legacy built-in shaders are listed on the [shader replacement page](SL-ShaderReplacement.md).  You can also create your own values for your custom SubShaders. |

### ForceNoShadowCasting tag

| **Signature** | **Function** |
| --- | --- |
| “ForceNoShadowCasting” = “[state]” | Whether to prevent shadow casting (and sometimes receiving) for all geometry that uses this SubShader. |

| **Signature** | **Value** | **Function** |
| --- | --- | --- |
| [state] | True | Unity prevents the geometry in this SubShader from casting shadows.  In the Built in Render Pipeline, with the Forward or Legacy Vertex Lit rendering paths, Unity also prevents the geometry in this SubShader from receiving shadows.  In HDRP, this does not prevent the geometry from casting contact shadows. |
|  | False | Unity does not prevent the geometry in this SubShader from casting or receiving shadows. This is the default value. |

### DisableBatching tag

| **Signature** | **Function** |
| --- | --- |
| “DisableBatching” = “[state]” | Whether Unity prevents **Dynamic Batching** for all geometry that uses this SubShader. |

| **Signature** | **Value** | **Function** |
| --- | --- | --- |
| [state] | True | Unity prevents Dynamic Batching for geometry that uses this SubShader. |
|  | False | Unity does not prevent Dynamic Batching for geometry that uses this SubShader. This is the default value. |
|  | LODFading | Unity prevents Dynamic Batching for all geometry that is part of a [LODGroup](https://docs.unity3d.com/Manual/class-LODGroup.html) with a Fade Mode value that is not None. Otherwise, Unity does not prevent Dynamic Batching. |

## IgnoreProjector tag

| **Signature** | **Function** |
| --- | --- |
| “IgnoreProjector” = “[state]” | Whether Unity ignores Projectors when rendering this geometry. |

| **Signature** | **Value** | **Function** |
| --- | --- | --- |
| [state] | True | Unity ignores Projectors when rendering this geometry. |
|  | False | Unity does not ignore Projectors when rendering this geometry. This is the default value. |

### PreviewType tag

The `PreviewType` SubShader Tag tells the Unity Editor how to display a material that uses this SubShader in the Material **Inspector**.

| **Signature** | **Function** |
| --- | --- |
| “PreviewType” = “[shape]” | Which shape the Unity Editor uses to display a preview of a material that uses this SubShader. |

| **Signature** | **Value** | **Function** |
| --- | --- | --- |
| [shape] | Sphere | Display the material on a sphere. This is the default value. |
|  | Plane | Display the material on a plane. |
|  | **Skybox** | Display the material on a skybox. |

## Additional resources

* [Add a shader tag to a SubShader or Pass](writing-shader-tags-introduction.md)
* [Configure if or when Unity uses a shader](writing-shader-tags.md)