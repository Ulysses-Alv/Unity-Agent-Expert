# Create a decal via a Decal Projector in URP

The [Decal Projector](renderer-feature-decal-projector-reference.md) component can project a Material as a decal if the Material uses a **Shader** Graph with the Decal Material type.

![Shader Graph with the Decal Material type](../../uploads/urp/decal/decal-shader-graph-material-type.png)  
*Shader Graph with the Decal Material type*

URP contains the pre-built Decal Shader (`Shader Graphs/Decal`).

![Decal Material properties.](../../uploads/urp/decal/decal-material-properties.png)  
*Decal Material properties and advanced options.*

## Create custom Decal shaders

The pre-built `Shader Graphs/Decal` shader serves as a simple example. You can create your own decal shaders that render decals in a way that suits your project best.

You can also create new decal materials and shaders directly from the **Material** property in the [Decal Projector](renderer-feature-decal-projector-reference.md) component with the following options:

* **URP Decal**: Creates a new material from the default **URP Decal** shader graph asset. This is a decal shader already set up for URP’s decal system.
* **ShaderGraph Decal**: Creates a new material or material variant from a new shader graph asset based on the **Decal URP Default** shader graph template. This is ideal if you want to customize the decal’s appearance or behavior using Shader Graph nodes.

To create a custom decal Shader Graph manually, in the shader target, go to **Graph Settings** and set the **Material** property to **Decal**.

![Shader Graph, Decal Material](../../uploads/urp/decal/decal-shader-graph-material-type.png)

Shader Graph, Decal Material

Enabling one of the following properties override the equivalent Lit Shader property on the surface of the Material.

To improve performance, pack data for different surface properties into a single texture. This way the shader performs fewer samples and Unity stores fewer textures.

For example, the following Shader Graph uses a **normal map** and a mask map to drive all properties in the shader. This decal is used for the damaged tarmac effect, and a hardcoded roughness value of 0 suites the use case.

![Decal Graph](../../uploads/urp/decal/decal-graph.png)

Decal Graph

The shader samples the mask and uses the color for setting the Ambient Occlusion values (Red channel), smoothness values (Green channel), Emission intensity values (Blue channel), and alpha values for the entire decal. Decals are often blended using single alpha values for all properties. The following image shows the mask map for the example tarmac cracks:  
![Decal Mask](../../uploads/urp/decal/decal-mask.png)  
*Example of mask map that packs Ambient Occlusion, Smoothness, Emission, and alpha values of a decal atlas into a single texture.*