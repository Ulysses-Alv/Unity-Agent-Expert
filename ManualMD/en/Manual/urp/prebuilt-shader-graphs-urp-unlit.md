# Unlit shader graph reference for URP

Explore the properties and settings you can use to customize the default Unlit **shader** graph in the Universal **Render Pipeline** (URP).

## Contexts

The default **Vertex** and **Fragment** Contexts contain the following Blocks. Which Blocks display might change depending on the [graph settings](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

### Vertex Context

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Position** | Vector3 | Sets the vertex position in object space. |
| **Normal** | Vector3 | Sets the normal of the vertex in object space. |
| **Tangent** | Vector3 | Sets the tangent of the vertex in object space. |
| **Motion Vector** | Vector3 | Sets the per‑vertex motion vector in object space. This property is available only if you set **Additional Motion Vector** to **Custom** in graph settings. |

### Fragment Context

| **Input** | **Type** | **Description** |
| --- | --- | --- |
| **Base Color** | Vector3 | Sets the base color of the material. The default is [`Color.grey`](../../ScriptReference/Color-grey.md). |
| **Alpha** | Float | Sets the alpha of the material. The range is 0 to 1. The default is 1. This property is available only if you set **Surface Type** to **Transparent** or you enable **Alpha Clipping** in graph settings. |
| **Alpha Clip Threshold** | Float | Sets the minimum alpha value for a **pixel** to be visible. The range is 0 to 1. The default is 0.5. This property is available only if you enable **Alpha Clipping** in the graph settings. |

## Graph Settings

Explore the shader graph settings you can use to customize the Vertex and Fragment Contexts.

For more details about graph settings that are common to all shader graph shaders, refer to the [Graph Settings tab](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

| **Setting** | **Description** |
| --- | --- |
| **Allow Material Override** | Enables material instances to override properties exposed in the shader graph. When disabled, materials using this shader will always use the values defined in the shader asset. |
| **Surface Type** | Determines if the material supports transparency. The options are:  * **Opaque**: The material is fully opaque and doesn’t use transparency blending. * **Transparent**: The material uses transparency blending based on the alpha channel. Unity adds the **Alpha** Block to the Fragment Context. |
| **Blending Mode** | Sets how the shader blends the color of a transparent material with the background. This property is available only when you set **Surface Type** to **Transparent**. The options are:  * **Alpha**: Calculates transparency using the alpha value of the material. * **Premultiply**: Calculates transparency using the alpha value of the material, but multiplies the RGB values by the alpha value first. * **Additive**: Adds the color of the material to the background color. * **Multiply**: Multiplies the color of the material with the background color.  For more information about blending modes, refer to [Blend modes](2d-light-blend-modes.md). |
| **Render Face** | Specifies which faces of the mesh the shader renders. The options are:  * **Front**: Renders only front‑facing triangles. This is the default. * **Back**: Renders only back‑facing triangles. * **Both**: Renders both front and back faces. |
| **Depth Write** | Determines whether the shader writes depth information to the depth buffer. The options are:  * **Auto**: Sets opaque objects to write to the depth buffer, and transparent objects to not write to the depth buffer. * **Force Enabled**: Enables all objects from writing to the depth buffer. * **Force Disabled**: Disables all objects from writing to the depth buffer. |
| **Depth Test** | Specifies the comparison function used to determine whether a pixel passes the depth test. The options are:  * **Never**: The pixel never passes the depth test. * **Less**: Passes if the incoming depth value is less than the stored depth. * **Equal**: Passes if the incoming depth value equals the stored depth. * **L Equal**: Passes if the incoming depth value is less than or equal to the stored depth. * **Greater**: Passes if the incoming depth value is greater than the stored depth. * **Not Equal**: Passes if the incoming depth value isn’t equal to the stored depth. * **G Equal**: Passes if the incoming depth value is greater than or equal to the stored depth. * **Always**: The pixel always passes the depth test. |
| **Alpha Clipping** | Discards pixels if their alpha value is lower than a threshold value. When you enable this setting, Unity adds the **Alpha Clip Threshold** Block to the Fragment context. Avoid using this setting when **Alpha** and **Alpha Clip Threshold** are constant for the entire material, as it can cause visual artifacts and unnecessary MSAA performance overhead. |
| **Cast Shadows** | Determines whether the material casts shadows onto other objects. |
| **Supports **LOD** Cross Fade** | Enables level of detail (LOD) cross‑fading when the **Mesh** Renderer has a **LOD Group** component with **Fade Mode** set to **Cross Fade**. |
| **Additional Motion Vectors** | Controls how Unity generates motion vectors for this material, for motion blur and temporal anti‑aliasing. The options are:  * **None**: Generates no motion vectors. * **Time Based**: Computes motion vectors automatically from per‑frame transformations. * **Custom**: Enables you to manually provide motion vector data through the **Motion Vector** Block in the Vertex Context.  For more information, refer to [Built-in shader support for motion vectors in URP](features/motion-vectors-shader-support.md). |
| **Alembic Motion Vectors** | Enables motion vector support for meshes streamed through Alembic files. For more information, refer to [Built-in shader support for motion vectors in URP](features/motion-vectors-shader-support.md). |
| **Keep Lighting Variants** | Adds the lighting keywords in the shader, even though it’s Unlit. Enable this property to create custom lighting and call URP lighting functions in **Custom Function** nodes without manually adding and managing shader keywords. This property is applicable only when you use the [Forward or Forward+ rendering path](rendering/forward-rendering-paths.md). |
| **Default Decal Blending** | Performs decal blending automatically. When disabled, you can implement decal blending directly in the shader graph to customize its behavior or allow decals to receive lighting from custom lighting implementations. This property is applicable only when you use the [Forward or Forward+ rendering path](rendering/forward-rendering-paths.md). |
| **Default SSAO** | Applies Screen Space **Ambient Occlusion** (SSAO) automatically by sampling the pre‑rendered SSAO texture. When disabled, you can implement SSAO manually in the shader graph and customize its functionality. This property is applicable only when you use the [Forward or Forward+ rendering path](rendering/forward-rendering-paths.md). |
| **Custom Editor GUI** | Renders a custom editor GUI in the ****Inspector**** window of the material. Enter the name of the GUI class in the field. For more information, refer to [Control material properties in the Inspector window](../writing-shader-display-types.md). |

## Additional resources

* [Get started with Shader Graph](https://learn.unity.com/tutorial/get-started-with-shader-graph) on Unity Learn
* [Introduction to motion vectors](features/motion-vectors.md)