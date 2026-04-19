# Virtual Texturing error material

You can use Streaming Virtual Texturing (SVT) with ****shaders**** you create in [Shader Graph](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest). To use SVT to stream one or more textures, you need to set up your material and Shader Graph correctly. To learn more, see [Using Streaming Virtual Texturing in Shader Graph](https://docs.unity3d.com/2020.1/Documentation/Manual/svt-use-in-shader-graph.html).

When a texture is not set up correctly, Unity can’t render the material. Instead, Unity renders a Virtual Texturing (VT) error texture with a purple and blue checkerboard pattern. This texture remains visible until you fix the validation error.

The Virtual Texturing error texture:

![Virtual Texturing Error Material](../uploads/Main/VT_Error_Material.png)

Virtual Texturing Error Material

## Validation errors

A Texture Stack is a group of textures that Unity samples at the same time, using the same UV coordinates. A validation error occurs when the Texture Stack or a texture in the Texture Stack is not in a valid state.

You can see the validation error in the **Material **Inspector****:

![Virtual Texturing Error](../uploads/Main/VT_Error.png)

Virtual Texturing Error

This validation error example is caused by a Texture Stack with three layers that only has two texture slots assigned.

Common reasons for Texture Stack validation errors include:

* A texture slot in the Texture Stack does not have an assigned texture.
* The Texture Stack is used in a material that is part of an [AssetBundle](https://docs.unity3d.com/2020.1/Documentation/Manual/AssetBundlesIntro.html).
* The Texture Stack is used in a material that is part of a [SubScene](https://docs.unity3d.com/Packages/com.unity.entities@0.13/api/Unity.Scenes.SubScene.html).
* The Texture Stack is not of a supported data type (see [Streaming Virtual Texturing requirements and compatibility](https://docs.unity3d.com/2020.1/Documentation/Manual/svt-requirements-compatibility.html)).

Errors also occur when one or more textures in the stack are not in a valid state. Causes of a texture validation error include:

* The texture is not of type Texture2D (SVT does not support arrays or cube maps).
* The texture uses a Mirror wrap mode.
* The texture uses different clamping modes per axis (for example, repeating horizontally while clamping vertically).The texture’s resolution is not a power of two.
* The **aspect ratios** of all textures are not the same.
* The texture does not meet the minimum size of 128 by 128 **pixels**.
* The texture does not have a mipmap.
* The texture has **Use Crunch **compression**** enabled in the [texture importer](https://docs.unity3d.com/Manual/class-TextureImporter.html).

## Build player error

For Unity to build a player using SVT, all materials must use valid textures with matching aspect ratios and have valid Virtual Texture properties (see [**Validation errors**](#validation-errors)). If a material in your player build is in an invalid state, it causes a build error that appears in the console log.

## Error naming convention

When Unity logs a validation error, it names the Texture Stacks according to the following convention:

* The number of layers the Texture Stack has.
* The Texture Stack’s layer names as set in ShaderGraph.

Here are two examples of this naming convention:

* An error name in a Texture Stack with two layers: Texture 2(Layer\_Name\_1, NULL)
* An error name in a Texture Stack with three layers: Texture 3 (Layer\_Name\_1, Layer\_Name\_2, NULL)

When an error appears in the console log, the error name also includes the material that uses the Texture Stack. For example:

* ‘Texture Stack 2 (3\_BaseColorMap, NULL)’ of material ‘materialname3’.