# Graph Settings tab reference | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Graph-Settings-Tab.html)

# Graph Settings tab reference

Use the **Graph Settings** tab in the [Graph Inspector](Internal-Inspector.html) window to change settings that affect the current shader graph as a whole.

## General properties

| Property | Description |
| --- | --- |
| **Precision** | Select a default [Precision Mode](Precision-Modes.html) for the entire graph. You can override the precision mode at the node level in your graph. |
| **Preview** | Select your preferred preview mode for the nodes that support preview. The options are:  * **Inherit**: The Unity Editor automatically selects the preview mode to use. * **Preview 2D**: Renders the output of the sub graph as a flat two-dimensional preview. * **Preview 3D**: Renders the output of the sub graph on a three-dimensional object such as a sphere.  This property is available only in [sub graphs](Sub-graph.html). |

## Target Settings

Add or remove graph targets to the current shader graph and set target properties according to the selected material type.

### Active Targets

A list that contains the [graph targets](Graph-Target.html) selected for the current shader graph. Select the **Add (+)** and **Remove (−)** buttons to add or remove **Active Targets**.

Shader Graph supports the following target types:

* **Custom Render Texture**: Shaders for updating [Custom Render Textures](Custom-Render-Texture.html).
* **Built-in**: Shaders for Unity’s [Built-In Render Pipeline](https://docs.unity3d.com/6000.0/Documentation/Manual/built-in-render-pipeline.html).
* **Universal**: Shaders for the [Universal Render Pipeline (URP)](https://docs.unity3d.com/Manual/universal-render-pipeline.html), available only if your project uses URP.
* **HDRP**: Shaders for the [High Definition Render Pipeline (HDRP)](https://docs.unity3d.com/6000.0/Documentation/Manual/high-definition-render-pipeline.html), available only if your project uses HDRP.

##### Important

The Built-In Render Pipeline is deprecated and will be made obsolete in a future release.  
It remains supported, including bug fixes and maintenance, through the full Unity 6.7 LTS lifecycle.  
For more information on migration, refer to [Migrating from the Built-In Render Pipeline to URP](https://docs.unity3d.com/6000.5/Documentation/Manual/urp/upgrading-from-birp.html) and [Render pipeline feature comparison](https://docs.unity3d.com/6000.5/Documentation/Manual/render-pipelines-feature-comparison.html).

### Target properties

Each graph target added in the list of **Active Targets** has its own set of properties.

| Property | Description |
| --- | --- |
| **Material** | Selects a material type for the target. The available options depend on the current target type. |
| Other properties (contextual) | A set of material and shader related properties that correspond to the current target type and the **Material** you select for the target.  * For Universal Render Pipeline (URP) target properties, refer to [Shader graph material Inspector window reference for URP](https://docs.unity3d.com/6000.6/Documentation/Manual/urp/shaders-in-universalrp-reference.html). * For High Definition Render Pipeline (HDRP) target properties, refer to HDRP's [Shader Graph materials reference](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/shader-graph-materials-reference.html). |
| **Custom Editor GUI** | Renders a custom editor GUI in the Inspector window of the material. Enter the name of the GUI class in the field. For more information, refer to [Control material properties in the Inspector window](https://docs.unity3d.com/6000.6/Documentation/Manual/writing-shader-display-types.html) and [Custom Editor block in ShaderLab reference](https://docs.unity3d.com/Manual/SL-CustomEditor.html). |
| **Support VFX Graph** | Enables this shader graph to support the [Visual Effect Graph](https://docs.unity3d.com/Packages/com.unity.visualeffectgraph@latest) to render particles. **Note**: This option is only available for certain material types. |

## Additional resources

* [Precision Modes](Precision-Modes.html)
* [Graph targets](Graph-Target.html)