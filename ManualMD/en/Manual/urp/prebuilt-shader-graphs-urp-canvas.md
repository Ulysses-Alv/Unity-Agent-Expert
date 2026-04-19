# Canvas shader graph reference for URP

Explore the properties and settings you can use to customize the default Canvas **shader** graph in the Universal **Render Pipeline** (URP). You can apply Canvas shader graph materials to user interface elements you create using the [uGUI (Unity UI)](https://docs.unity3d.com/Packages/com.unity.ugui@latest) (uGUI) package.

## Contexts

The default **Fragment** Context contains the following Blocks. Which Blocks display might change depending on the [graph settings](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

### Fragment Context

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Base Color** | Vector3 | Sets the base color of the material. The default is [`Color.grey`](../../ScriptReference/Color-grey.md). |
| **Alpha** | Float | Sets the alpha of the material. The range is 0 to 1. The default is 1. |
| **Emission** | Float3 | Sets the color the material emits as a light source. The default is black. For more information, refer to [Add light emission to a material](https://docs.unity3d.com/Manual/StandardShaderMaterialParameterEmission.html) |
| **Alpha Clip Threshold** | Float | Sets the minimum alpha value for a **pixel** to be visible. The range is 0 to 1. The default is 0.5. This property is only available if you enable **Alpha Clipping** in the graph settings. |

## Graph Settings

Explore the shader graph settings you can use to customize the Fragment context.

For more details about graph settings that are common to all shader graph shaders, refer to [Graph Settings tab](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

| **Setting** | **Description** |
| --- | --- |
| **Alpha Clipping** | Discards pixels if their alpha value is lower than a threshold value. Enabling this setting adds the **Alpha Clip Threshold** Block to the Fragment context. |
| **Disable Color Tint** | Prevents Shader Graph tinting the **sprite** with the **Color** property of the sprite. If you enable this setting, use a [Vertex Color Node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Vertex-Color-Node.html) to access the **Color** property of the sprite. |
| **Custom Editor GUI** | Renders a custom editor GUI in the ****Inspector**** window of the material. Enter the name of the GUI class in the field. For more information, refer to [Control material properties in the Inspector window](../writing-shader-display-types.md). |

## Additional resources

* [UI systems](https://docs.unity3d.com/Manual/UIToolkits.html)
* [UGUI Shaders](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Shader-Graph-Sample-UGUI-Shaders.html)