# Set the render pipeline and render order for a shader

To set the **render pipeline** and render order for a subshader or a **shader** pass, add a `Tag` block.

## Add a tag

Add a `Tag` block to a `SubShader` block or a `Pass` block. For example:

```
Shader "TagsExample"
{
    SubShader
    {
        // Add a subshader tag
        Tags { "TagName" = "Value" }

        Pass
        {    
            // Add a pass tag
            Tags { "TagName" = "Value" }
            ...
        }
    }
}
```

For more information about tag names and values, refer to [`SubShader` tags in ShaderLab reference](SL-SubShaderTags.md) and [`Pass` tags in ShaderLab reference](SL-PassTags.md). Not all `Tag` blocks are compatible with both subshaders and passes.

You can also add your own custom tags and read them from a C# script.

## Set the render pipeline

To make a subshader compatible with the Universal Render Pipeline (URP) or the High Definition Render Pipeline (HDRP), add a `RenderPipeline` tag to the `SubShader` block:

* For URP, use `"RenderPipeline" = "UniversalPipeline"`
* For HDRP, use `"RenderPipeline" = "HDRenderPipeline"`

## Set the render order

Unity uses two tags to determine when to execute a shader:

* `LightMode` tag: Sets which stage of the render pipeline a shader pass runs during. For example, drawing opaque geometry, or drawing **post-processing** effects.
* `Queue` tag: Sets which group of objects a subshader belongs to. Unity draws each render queue in a specific order.

Each `LightMode` stage usually renders a specific `Queue` range. Unity warns you if you use a `Queue` tag that the `LightMode` stage doesn’t render.

**Note:** The `LightMode` tag isn’t related to the [light mode](LightModes-introduction.md) parameter in the **Inspector** window of a Light component.

### Set the `LightMode` tag

Add a `LightMode` tag to the `Pass` block. For example, to execute the shader during the **forward rendering** pass in the Universal Render Pipeline (URP) that renders the `Geometry` queue, use `Tags { "LightMode" = "UniversalForward" }`.

For a list of `LightMode` values, refer to [`LightMode` tag valid values](SL-PassTags.md#lightmode-tag) and [ShaderLab `Pass` tags in URP reference](urp/urp-shaders/urp-shaderlab-pass-tags.md).

To enable or disable a shader pass by its `LightMode` tag, use the [`Material.SetShaderPassEnabled`](../ScriptReference/Material.SetShaderPassEnabled.md) API.

### Set the Queue tag

Add a `Queue` tag to the `SubShader` block. Use one of the following approaches:

* To add the subshader to a queue, use `"Queue" = "RenderQueueName"`. For example `"Queue" = "Transparent"`.
* To add the subshader before or after a queue, use `"Queue" = "RenderQueueName-Offset"` or `"Queue" = "RenderQueueName+Offset"`. For example, `"Queue" = "Geometry+1"`.

You can also use queue values instead of names. For example, `"Queue" = "3000"` instead of `"Queue" = "Transparent"`, or `"Queue" = "2001"` instead of `"Queue" = "Geometry+1"`.

You can also set the queue in the [**Inspector** window of the material](class-Material.md) using the **Render Queue** property, or in a C# script using the [`Material.renderQueue`](../ScriptReference/Material-renderQueue.md) API.

For more information, refer to [Render queues and sorting behaviours](built-in-rendering-order.md).

## Check tags

To read tags from a C# script, use the following:

* To read a subshader tag, use the [`Material.GetTag`](../ScriptReference/Material.GetTag.md) API.
* To read a pass tag, use the [`Shader.FindPassTagValue`](../ScriptReference/Shader.FindPassTagValue.md) API.
* To read a `Queue` tag, use the [`Shader.renderQueue`](../ScriptReference/Shader-renderQueue.md) API.

## Additional resources

* [SubShader tags in ShaderLab reference](SL-SubShaderTags.md)
* [Pass tags in ShaderLab reference](SL-PassTags.md)