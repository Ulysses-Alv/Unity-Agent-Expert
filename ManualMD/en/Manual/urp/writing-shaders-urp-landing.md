# Examples of writing a custom shader in URP

Techniques for writing **shaders** in the Universal **Render Pipeline** (URP).

Each example covers some extra information compared to the basic shader example. If you’re new to writing shaders using Unity’s **ShaderLab** language, consider going through the sections in the order of appearance on this page.

| **Page** | **Description** |
| --- | --- |
| [Create a sample scene in URP](writing-shaders-urp-basic-prerequisites.md) | Create a sample **scene** for writing shaders in URP. |
| [Write an unlit basic shader in URP](writing-shaders-urp-basic-unlit-structure.md) | An example of a basic URP shader that fills a **mesh** shape with a color. |
| [Write an unlit shader with color input in URP](writing-shaders-urp-unlit-color.md) | An example of a URP shader that adds the **Base Color** property to a material. |
| [Draw a texture in a shader in URP](writing-shaders-urp-unlit-texture.md) | An example of a URP shader that draws a texture on a mesh. |
| [Visualize normal vectors in a shader in URP](writing-shaders-urp-unlit-normals.md) | An example of a URP shader that visualizes the normal vector values on a mesh. |
| [Write depth only in a shader in URP](writing-shaders-urp-depth-only.md) | Add a depth prepass that writes the depth of objects before the opaque and transparent passes. |
| [Reconstruct world space positions in a shader in URP](writing-shaders-urp-reconstruct-world-position.md) | An example of a URP shader that reconstructs the world space positions for **pixels** using a depth texture and screen space UV coordinates. |