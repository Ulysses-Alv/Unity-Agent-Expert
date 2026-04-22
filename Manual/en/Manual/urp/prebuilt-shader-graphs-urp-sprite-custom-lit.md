# Sprite Custom Lit shader graph reference for URP

Explore the properties and settings you can use to customize the default **Sprite** Custom Lit **shader** graph in the Universal **Render Pipeline** (URP).

This shader graph has the same Contexts and graph settings as the [Sprite Lit shader graph](prebuilt-shader-graphs-urp-sprite-lit.md). The compiled shader doesn’t use a default lighting model, to allow you to create a custom lighting model.

For more information about 2D shaders in URP, refer to [2D Lighting in URP](2d-index.md).

## Contexts

The default **Vertex** and **Fragment** Contexts contain the following Blocks. Which Blocks display might change depending on the [graph settings](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

### Vertex Context

The default values are the values from the **mesh**.

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Position** | Vector3 | Sets the vertex position in object space. |
| **Normal** | Vector3 | Sets the normal of the vertex in object space. |
| **Tangent** | Vector3 | Sets the tangent of the vertex in object space. |

### Fragment Context

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Base Color** | Vector3 | Sets the base color of the material. The default is [`Color.grey`](../../ScriptReference/Color-grey.md). |
| ****Sprite Mask**** | Vector4 | Sets a custom mask that determines whether the **pixels** of the sprite are visible. For more information, refer to [Add normal map and mask textures to a sprite](SecondaryTextures.md). |
| **Normal (Tangent Space)** | Vector3 | Sets a custom **normal map** to simulate bumps and grooves on the surface of the sprite. For more information, refer to [Add normal map and mask textures to a sprite](SecondaryTextures.md). |
| **Alpha** | Float | Sets the alpha of the material. The range is 0 to 1. The default is 1. |
| **Alpha Clip Threshold** | Float | Sets the minimum alpha value for a pixel to be visible. The range is 0 to 1. The default is 0.5. This property is only available if you enable **Alpha Clipping** in the graph settings. |

## Graph Settings

Explore the shader graph settings you can use to customize the Fragment and Vertex contexts.

For more details about graph settings that are common to all shader graph shaders, refer to [Graph Settings tab](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

| **Setting** | **Description** |
| --- | --- |
| **Blending Mode** | Sets how the shader blends the color of a transparent material with the background. The options are:  * **Alpha**: Calculates transparency using the alpha value of the material. * **Premultiply**: Calculates transparency using the alpha value of the material, but multiplies the RGB values by the alpha value first. * **Additive**: Adds the color of the material to the background color. * **Multiply**: Multiplies the color of the material with the background color.  For more information about blending modes, refer to [Blend modes](2d-light-blend-modes.md). |
| **Depth Write** | Determines whether the shader writes depth information to the depth buffer. The options are:  * **Auto**: Sets opaque objects to write to the depth buffer, and transparent objects to not write to the depth buffer. * **Force Enabled**: Enables all objects from writing to the depth buffer. * **Force Disabled**: Disables all objects from writing to the depth buffer. |
| **Alpha Clipping** | Discards pixels if their alpha value is lower than a threshold value. Enabling this setting adds the **Alpha Clip Threshold** Block to the Fragment context. |
| **Disable Color Tint** | Prevents Shader Graph tinting the sprite with the **Color** property of the sprite. If you enable this setting, use a [Vertex Color Node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Vertex-Color-Node.html) to access the **Color** property of the sprite. |
| **Sort 3D as 2D Compatible** | Uses 2D sorting rules to draw **3D objects**, to avoid **2D objects** appearing at the wrong depth. |
| **Custom Editor GUI** | Renders a custom editor GUI in the ****Inspector**** window of the material. Enter the name of the GUI class in the field. For more information, refer to [Control material properties in the Inspector window](../writing-shader-display-types.md). |

## Additional resources

* [Introduction to normal maps](../StandardShaderMaterialParameterNormalMap.md)
* [2D lighting in URP](2d-index.md)
* [20 advanced 2D shader effects - part 1](https://www.youtube.com/watch?v=wav3YUcqUCA)