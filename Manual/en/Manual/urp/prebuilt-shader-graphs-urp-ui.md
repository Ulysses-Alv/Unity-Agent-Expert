# UI shader graph reference

Explore the properties and settings you can use to customize the default [UI shader graph](../ui-systems/ui-shader-graph.md) in the Universal **Render Pipeline** (URP).

## Contexts

The default **Fragment** Context contains the following Blocks. Which Blocks display might change depending on the [graph settings](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

### Fragment Context

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Base Color** | Vector3 | Sets the base color of the material. The default is [`Color.grey`](../../ScriptReference/Color-grey.md). |
| **Alpha** | Float | Sets the alpha of the material. The range is 0 to 1. The default is 1. |

## Graph Settings

Explore the **shader** graph settings you can use to customize the Fragment context.

For more details about graph settings that are common to all shader graph shaders, refer to [Graph Settings tab](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

| **Setting** | **Description** |
| --- | --- |
| **Blending Mode** | Sets how the shader blends the color of a transparent material with the background. This property is available only when you set **Surface Type** to **Transparent**. The options are:  * **Alpha**: Calculates transparency using the alpha value of the material. * **Premultiply**: Calculates transparency using the alpha value of the material, but multiplies the RGB values by the alpha value first. * **Additive**: Adds the color of the material to the background color. * **Multiply**: Multiplies the color of the material with the background color.  For more information about blending modes, refer to [Blend modes](2d-light-blend-modes.md). |

## Additional resources

* [UI Toolkit](../UIElements.md)
* [UI Shaders Graph](../ui-systems/ui-shader-graph.md)