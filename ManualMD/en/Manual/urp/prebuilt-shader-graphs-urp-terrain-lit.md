# Terrain Lit shader graph reference for URP

Explore the properties and settings you can use to customize the default Terrain Lit **shader** graph in the Universal **Render Pipeline** (URP).

## Terrain shader passes

Some properties might not work as expected, because the terrain system adds two additional shader passes when it generates a terrain shader. The passes are the following:

* Basemap Gen, which renders the shader graph on a **quad** to create a low-resolution baked version. As a result, the **Position** and **Normal (Tangent Space)** values are 2D during this pass.
* Basemap Rendering, which uses a built-in shader and the texture Basemap Gen generates. The **Position** Block has no affect on this pass.

## Contexts

The default **Vertex** and **Fragment** Contexts contain the following Blocks. Which Blocks display might change depending on the [graph settings](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

### Vertex Context

The default values are the values from the **mesh**.

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Position** | Vector3 | Sets the vertex position in object space.  **Note**: The vertex position has no effect on the Basemap Rendering pass. |

### Fragment Context

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Base Color** | Vector3 | Sets the base color of the material. The default is [`Color.grey`](../../ScriptReference/Color-grey.md). |
| **Normal (Tangent Space)** | Vector3 | Sets a custom **normal map** to simulate bumps and grooves on the surface of the **sprite**. For more information, refer to [Add normal map and mask textures to a sprite](SecondaryTextures.md). |
| **Metallic** | float | Sets how metallic the surface of your material is. The range is between 0 and 1. A higher value means the surface reflects the environment more, and its albedo color becomes less visible. A lower value means the albedo color is clearer, and reflections from the environment are visible on top of the surface color, rather than obscuring it. A value of 1 means the the surface color is entirely driven by reflections from the environment. |
| **Emission** | Vector3 | Sets the color the material emits as a light source. The default is black. For more information, refer to [Add light emission to a material](https://docs.unity3d.com/Manual/StandardShaderMaterialParameterEmission.html) |
| **Smoothness** | float | Sets the smoothness of the material. Every light ray that hits a smooth surface bounces off at predictable and consistent angles. For a perfectly smooth surface that reflects light like a mirror, set this to a value of 1. Less smooth surfaces reflect light over a wider range of angles, so the reflections have less detail and spread across the surface in a more diffused pattern. |
| ****Ambient Occlusion**** | float | Sets how much **ambient light** reaches the surface, rather than being occluded. A value of 0 means the fragment is completely occluded and appears black. A value of 1 means the fragment is not occluded at all, and the ambient color does not change. |
| **Alpha** | Float | Sets the alpha of the material. The range is 0 to 1. The default is 1. |
| **Alpha Clip Threshold** | Float | Sets the minimum alpha value for a **pixel** to be visible. The range is 0 to 1. The default is 0.5. This property is only available if you enable **Alpha Clipping** in the graph settings. |

## Graph Settings

Explore the shader graph settings you can use to customize the Fragment and Vertex contexts.

For more details about graph settings that are common to all shader graph shaders, refer to [Graph Settings tab](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

| **Setting** | **Description** |
| --- | --- |
| **Depth Write** | Determines whether the shader writes depth information to the depth buffer. The options are:  * **Auto**: Sets opaque objects to write to the depth buffer, and transparent objects to not write to the depth buffer. * **Force Enabled**: Enables all objects from writing to the depth buffer. * **Force Disabled**: Disables all objects from writing to the depth buffer. |
| **Depth Test** | Sets the conditions under which geometry passes or fails depth testing. The GPU doesn’t draw pixels that fail a depth test. For more information, refer to [Control material properties in the Inspector window](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/surface-options.html). |
| **Alpha Clipping** | Discards pixels if their alpha value is lower than a threshold value. Enabling this setting adds the **Alpha Clip Threshold** Block to the Fragment context. To access the holes texture from the [Terrain Texture Node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Terrain-Texture-Node.html), you must enable this property. |
| **Cast Shadows** | Enables the terrain casting shadows onto itself and other **GameObjects**. |
| **Receive Shadows** | Enables the terrain receiving shadows from other GameObjects. |
| **Fragment Normal Space** | Specifies the normal map space that this material uses. The options are:  * **TangentSpace**: Defines the normals in tangent space. Use this to tile a texture on a mesh. * **ObjectSpace**: Defines the normals in object space. Use this for planar-mapping GameObjects like the terrain. * **WorldSpace**: Defines the normal maps in world space. |
| **Custom Editor GUI** | Renders a custom editor GUI in the ****Inspector**** window of the material. Enter the name of the GUI class in the field. |

## Additional resources

* [Terrain](../script-Terrain.md)
* [Shader Graph manual](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest)