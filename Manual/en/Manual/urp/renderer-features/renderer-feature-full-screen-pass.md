# Full Screen Pass Renderer Feature reference for URP

The Full Screen Pass Renderer Feature lets you inject full screen render passes at pre-defined injection points to create full screen effects.

For more information about using a Full Screen Pass Renderer Feature, refer to [Create a low-code custom post-processing effect](../post-processing/post-processing-custom-effect-low-code.md).

## Properties

The Full Screen Pass Renderer Feature contains the following properties.

| Property | Description |
| --- | --- |
| **Name** | Sets the name of the Full Screen Pass Renderer Feature. |
| **Injection Point** | Sets when Unity renders the effect. The options are the following:  * **Before Rendering Transparents**: Renders the effect after the skybox and before transparent objects. * **Before Rendering Post Processing**: Renders the effect after transparent objects and before post-processing. * **After Rendering Post Processing**: Renders the effect after post-processing, before the `AfterRendering` pass.  The default option is **After Rendering Post Processing**.  For more information, refer to [Injection points reference](../customize/custom-pass-injection-points.md). |
| **Requirements** | Adds additional render passes so the full screen render pass can use the textures they generate. The options are the following:  * **None**: Adds no additional render passes. * **Everything**: Adds depth, normal, color, and motion vector render passes. * **Depth**: Creates a depth texture copy of the depth buffer. For more information, refer to **Depth Texture** in the [URP Asset](../universalrp-asset.md). * **Normal**: Adds a normals render pass. * **Color**: Creates a texture with the opaque objects in the scene, and binds it to the `_BlitTexture` texture in the shader. For more information, refer to **Opaque Texture** in the [URP Asset](../universalrp-asset.md). * **Motion**: Creates a motion vector texture. For more information, refer to [Motion vectors render pass](../features/motion-vectors-render-pass.md). |
| **Fetch Color Buffer** | Allows the full screen render pass to access the color texture the **camera** currently targets. Unity binds the texture to the `_BlitTexture` texture in the **shader**. |
| **Bind Depth-Stencil** | Allows the full screen render pass to access the depth-stencil texture the camera currently targets. Enabling this property has an impact on performance. |
| **Pass Material** | Sets the material the Renderer Feature uses to render the effect. Use the **New** dropdown to create a new material and its associated shader from a template. The options are the following:  * **SRP Blit Shader**: Creates a new shader using Unity’s Scriptable Render Pipeline (SRP) Blit template. * **ShaderGraph Fullscreen**: Creates a new material or material variant from a new shader graph asset based on the **Fullscreen Basic URP** shader graph template, which uses the [**URP Fullscreen** material type](../prebuilt-shader-graphs-urp-fullscreen.md). **Note**: To define if Unity creates a material or a material variant from the shader graph asset, refer to the **Graph Template Workflow** option in the [Shader Graph Preferences](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Shader-Graph-Preferences.html). |
| **Pass** | Selects the specific pass inside the **Pass Material** for the Renderer Feature to use.  This property is hidden by default. To access the property, at the top of the Full Screen Pass Renderer Feature, open the **More** (⋮) menu and enable **Advanced Properties**. |

## Additional resources

* [Adding pre-built effects via Renderer Features in URP](../urp-renderer-feature-landing.md)