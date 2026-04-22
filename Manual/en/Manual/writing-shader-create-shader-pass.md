# Add a shader pass in a custom shader

A Pass is the fundamental element of a **Shader** object. It contains instructions for setting the state of the GPU, and the shader programs that run on the GPU.

Simple Shader objects might contain only a single Pass, but more complex shaders can contain multiple Passes. You can use separate Passes to define parts of your **Shader object** that work differently; for example, parts that require a change to the render state, different shader programs, or a different `LightMode` Pass tag.

**Note**: In render pipelines based on the Scriptable **Render Pipeline**, you can use a RenderStateBlock to change the render state on the GPU, without requiring a separate Pass.

To define a regular Pass in **ShaderLab**, you place a `Pass` block inside a `SubShader` block.

You can also define two special types of Pass, using the `UsePass` or `GrabPass` commands. For information on those commands, see [ShaderLab commands: UsePass](SL-UsePass.md) or [ShaderLab commands: GrabPass](SL-GrabPass.md).

## Name a shader pass

A Pass can have a name. You need to reference a Pass by name in the `UsePass` command, and in some C# APIs. The name of a Pass is visible in the [Frame Debugger](FrameDebugger.md) tool.

To assign a name to a Pass in ShaderLab, you place a `Name` block inside a `Pass` block.

Internally, Unity converts the name to uppercase. When you reference the name in ShaderLab code, you must use the uppercase variant; for example, if the value is “example”, you must reference it as EXAMPLE.

If more than one Pass in the same SubShader has the same name, Unity uses the first Pass in the code.

### Using the Pass name with C# scripts

To access the name of a Pass from C# **scripts**, you can use APIs such as [Material.FindPass](../ScriptReference/Material.FindPass.md), [Material.GetPassName](../ScriptReference/Material.GetPassName.md), or [ShaderData.Pass.Name](../ScriptReference/ShaderData.Pass.Name.md).

**Note:** [Material.GetShaderPassEnabled](../ScriptReference/Material.GetShaderPassEnabled.md) and [Material.SetShaderPassEnabled](../ScriptReference/Material.SetShaderPassEnabled.md) do not reference Passes by name; instead, they reference Passes using the value of the [LightMode tag](SL-PassTags.md).

## Example

This example code demonstrates the syntax for creating a Shader object that contains a single SubShader, which in turn contains a single Pass.

```
Shader "Examples/SinglePass"
{
    SubShader
    {
        Pass
        {                
              Name "ExamplePassName"
              Tags { "ExampleTagKey" = "ExampleTagValue" }

              // ShaderLab commands go here.

              // HLSL code goes here.
        }
    }
}
```

## Additional resources

* [Pass block in ShaderLab reference](SL-Pass.md)
* [Pass tags in ShaderLab reference](SL-PassTags.md)
* [Shader code blocks in ShaderLab reference](shader-shaderlab-code-blocks.md)
* [Name directive in ShaderLab reference](SL-Name.md)