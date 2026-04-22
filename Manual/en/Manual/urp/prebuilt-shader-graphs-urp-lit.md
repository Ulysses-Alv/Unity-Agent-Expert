# Lit shader graph reference for URP

Explore the properties and settings you can use to customize the default Lit **shader** graph in the Universal **Render Pipeline** (URP).

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
| **Normal (Tangent Space)** | Vector3 | Sets the surface normal in tangent space. This property is available only when you set **Fragment Normal Space** to **Tangent** in graph settings. |
| **Normal Object Space** | Vector3 | Sets the surface normal in object space. This property is available only when you set **Fragment Normal Space** to **Object** in graph settings. |
| **Normal World Space** | Vector3 | Sets the surface normal in world space. This property is available only when you set **Fragment Normal Space** to **World** in graph settings. |
| **Metallic** | Float | Sets the metallic value of the material, where 0 is non-metallic and 1 is metallic. The default is 0. This property is available only when **Workflow Mode** is **Metallic**. |
| ****Specular Color**** | Vector3 | Sets the brightness and tint color of specular highlights. The default is [`Color.grey`](../../ScriptReference/Color-grey.md). This property is available only when **Workflow Mode** is **Specular**. |
| **Smoothness** | Float | Sets the smoothness of the material. The default is 0.5. |
| **Emission** | Float3 | Sets the color the material emits as a light source. The default is black. For more information, refer to [Add light emission to a material](https://docs.unity3d.com/Manual/StandardShaderMaterialParameterEmission.html). |
| ****Ambient Occlusion**** | Float | Sets the ambient occlusion of the material. The range is 0 to 1, where 0 means the fragment is completely occluded and 1 means the fragment isn’t occluded at all so the ambient color doesn’t change. The default is 1. |
| **Alpha** | Float | Sets the alpha of the material. The range is 0 to 1. The default is 1. This property is available only if you set **Surface Type** to **Transparent** or you enable **Alpha Clipping** in graph settings. |
| **Alpha Clip Threshold** | Float | Sets the minimum alpha value for a **pixel** to be visible. The range is 0 to 1. The default is 0.5. This property is available only if you enable **Alpha Clipping** in the graph settings. |
| **Coat Mask** | Float | Sets the strength of the clear coat layer, where 0 disables the coat and 1 applies it fully across the surface. The default is 0. This property is available only when you enable **Clear Coat** in graph settings. |
| **Coat Smoothness** | Float | Sets the smoothness of the clear coat layer. The default is 1. This property is available only when you enable **Clear Coat** in graph settings. |

## Graph Settings

Explore the shader graph settings you can use to customize the Vertex and Fragment Contexts.

For more details about graph settings that are common to all shader graph shaders, refer to the [Graph Settings tab](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Graph-Settings-Tab.html).

| **Setting** | **Description** |
| --- | --- |
| **Allow Material Override** | Enables material instances to override properties exposed in the shader graph. When disabled, materials using this shader will always use the values defined in the shader asset. |
| **Workflow Mode** | Defines how the shader calculates reflections and specular highlights. The options are:  * **Specular**: Uses a specular color input to define the brightness and tint of specular reflections for non‑metallic surfaces. * **Metallic**: Uses a metallic input to define how much of the surface behaves like metal. Metallic surfaces reflect the environment color, while non‑metallic surfaces reflect the base color.  When you change this setting, Unity adds either the **Specular Color** Block or the **Metallic** Block to the Fragment Context. |
| **Surface Type** | Determines if the material supports transparency. The options are:  * **Opaque**: The material is fully opaque and doesn’t use transparency blending. * **Transparent**: The material uses transparency blending based on the alpha channel. Unity adds the **Alpha** Block to the Fragment Context. |
| **Blending Mode** | Sets how the shader blends the color of a transparent material with the background. This property is available only when you set **Surface Type** to **Transparent**. The options are:  * **Alpha**: Calculates transparency using the alpha value of the material. * **Premultiply**: Calculates transparency using the alpha value of the material, but multiplies the RGB values by the alpha value first. * **Additive**: Adds the color of the material to the background color. * **Multiply**: Multiplies the color of the material with the background color.  For more information about blending modes, refer to [Blend modes](2d-light-blend-modes.md). |
| **Render Face** | Specifies which faces of the mesh the shader renders. The options are:  * **Front**: Renders only front‑facing triangles. This is the default. * **Back**: Renders only back‑facing triangles. * **Both**: Renders both front and back faces. |
| **Depth Write** | Determines whether the shader writes depth information to the depth buffer. The options are:  * **Auto**: Sets opaque objects to write to the depth buffer, and transparent objects to not write to the depth buffer. * **Force Enabled**: Enables all objects from writing to the depth buffer. * **Force Disabled**: Disables all objects from writing to the depth buffer. |
| **Depth Test** | Specifies the comparison function used to determine whether a pixel passes the depth test. The options are:  * **Never**: The pixel never passes the depth test. * **Less**: Passes if the incoming depth value is less than the stored depth. * **Equal**: Passes if the incoming depth value equals the stored depth. * **L Equal**: Passes if the incoming depth value is less than or equal to the stored depth. * **Greater**: Passes if the incoming depth value is greater than the stored depth. * **Not Equal**: Passes if the incoming depth value isn’t equal to the stored depth. * **G Equal**: Passes if the incoming depth value is greater than or equal to the stored depth. * **Always**: The pixel always passes the depth test. |
| **Alpha Clipping** | Discards pixels if their alpha value is lower than a threshold value. When you enable this setting, Unity adds the **Alpha Clip Threshold** Block to the Fragment context. |
| **Cast Shadows** | Determines whether the material casts shadows onto other objects. |
| **Receive Shadows** | Determines whether the material receives and displays shadows cast from other objects. |
| **Supports **LOD** Cross Fade** | Enables support for Level of Detail (LOD) cross‑fading when used with a **Mesh** Renderer that has **LOD Group** cross‑fade enabled. |
| **Additional Motion Vectors** | Controls how Unity generates motion vectors for this material, for motion blur and temporal anti‑aliasing. The options are:  * **None**: Generates no motion vectors. * **Time Based**: Computes motion vectors automatically from per‑frame transformations. * **Custom**: Enables you to manually provide motion vector data through the **Motion Vector** Block in the Vertex Context.  For more information, refer to [Built-in shader support for motion vectors in URP](features/motion-vectors-shader-support.md). |
| **Alembic Motion Vectors** | Enables motion vector support for meshes streamed through Alembic files. For more information, refer to [Built-in shader support for motion vectors in URP](features/motion-vectors-shader-support.md). |
| **Fragment Normal Space** | Specifies the coordinate space for the normal input in the Fragment Context. The options are:  * **Tangent**: Normals are relative to the mesh surface’s UV orientation. * **Object**: Normals are relative to the mesh’s local coordinate system. * **World**: Normals are relative to the global scene coordinate system.  When you change this setting, Unity adds the corresponding **Normal** Block to the Fragment Context. |
| **Clear Coat** | Enables you to add an additional transparent, reflective layer over the base material to simulate surfaces like car paint or varnished wood. When you enable this setting, Unity adds the **Coat Mask** and **Coat Smoothness** Blocks to the Fragment Context. |
| **Custom Editor GUI** | Renders a custom editor GUI in the ****Inspector**** window of the material. Enter the name of the GUI class in the field. For more information, refer to [Control material properties in the Inspector window](../writing-shader-display-types.md). |
| **Support VFX Graph** | Enables this shader graph to support the Visual Effect Graph. |

## Additional resources

* [Get started with Shader Graph](https://learn.unity.com/tutorial/get-started-with-shader-graph) on Unity Learn
* [Introduction to motion vectors](features/motion-vectors.md)