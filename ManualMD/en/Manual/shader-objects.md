# Shader object fundamentals

In Unity, when you work with shaders that are part of the [graphics pipeline](https://en.wikipedia.org/wiki/Graphics_pipeline), you usually work with instances of the [Shader](../ScriptReference/Shader.md) class. An instance of the `Shader` class is called a ****Shader object****.

A Shader object is a Unity-specific way of working with shader programs; it is a wrapper for shader programs and other information. It lets you define multiple shader programs in the same file, and tell Unity how to use them.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Shader objects** | Yes | Yes | Yes | Yes |

## Shader object

A Shader object contains shader programs, instructions for changing settings on the GPU (collectively called the render state), and information that tells Unity how to use them.

You use Shader objects with [materials](Materials.md) to determine the appearance of your **scene**.

## Assets

You can create Shader objects in two ways. Each has its own type of asset:

* You can [write code](writing-shader-writing-shader-programs-hlsl.md) to create a [shader asset](class-Shader.md), which is a text file with the `.shader` extension.
* You can use [Shader Graph](shader-graph.md) to create a Shader Graph asset.

Whichever way you create your Shader object, Unity represents the results in the same way internally.

## Inside a Shader object

A Shader object has a nested structure. It organizes information into structures called **SubShaders** and **Passes**. It organises shader programs into **shader variants**.

### Shader object

A Shader object contains:

* Information about itself, such as its name
* An optional fallback Shader object, which Unity uses if it can’t use this one
* One or more SubShaders

You can also define additional information such as shared shader code, or whether to use a [custom editor](editor-CustomEditors.md). For information on defining a Shader object, see [ShaderLab: defining a Shader object](SL-Shader.md).

### SubShaders

SubShaders let you separate your Shader object into parts that are compatible with different hardware, render pipelines, and runtime settings.

A SubShader contains:

* Information about which hardware, render pipelines, and runtime settings this SubShader is compatible with
* SubShader tags, which are key-value pairs that provide information about the SubShader
* One or more Passes

You can also define additional information, such as render state that is common to all of its Passes. For information on everything you can define in a SubShader, see [ShaderLab: defining a SubShader](SL-SubShader.md).

### Passes

A Pass contains:

* Pass tags, which are key-value pairs that provide information about the Pass
* Instructions for updating the render state before running its shader programs
* Shader programs, organised into one or more shader variants

You can also define additional information such as a name. For information on everything you can define in a Pass, see [ShaderLab: defining a Pass](SL-Pass.md).

### Shader variants

The shader programs that a Pass contains are organised into shader variants. Shader variants share common code, but have different functionality when a given keyword is enabled or disabled.

The number of shader variants in a Pass depends on how many keywords you define in your shader code, and the target platform. Each Pass contains at least one variant.

For more information, see [Shader variants](shader-variants.md).