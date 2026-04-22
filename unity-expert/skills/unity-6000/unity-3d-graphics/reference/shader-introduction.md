# Introduction to shaders

A **shader** program, commonly referred to as a shader, is a program that runs on a GPU.

## Types of shader

In Unity, shaders are divided into three broad categories. You use each category for different things, and work with them differently.

* Shaders that are part of the [graphics pipeline](https://en.wikipedia.org/wiki/Graphics_pipeline) are the most common type of shader. They perform calculations that determine the color of **pixels** on the screen. In Unity, you usually work with this type of shader by using [Shader objects](shader-objects.md).
* [Compute shaders](class-ComputeShader.md) perform calculations on the GPU, outside of the regular graphics pipeline.
* **Ray tracing** shaders perform calculations related to ray tracing.

## Terminology

The terminology around shaders can be confusing; people commonly use the word “shader” to mean different things.

In this documentation, the terminology is as follows:

* **shader** or **shader program** - a program that runs on a GPU. Unless otherwise specified, this means shader programs that are part of the graphics pipeline.
* **Shader object** - an instance of the `Shader` class. A Shader object is a wrapper for shader programs and other information.
* ****ShaderLab**** - a Unity-specific language for writing shaders.
* **Shader Graph** - a tool for creating shaders without writing code.
* **shader asset** - a file with the `.shader` extension in your Unity project. It defines a Shader object.
* **Shader Graph asset** - a file in your Unity project. It defines a Shader object.