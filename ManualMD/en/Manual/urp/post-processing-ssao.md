# Introduction to screen space ambient occlusion in URP

The Screen Space **Ambient Occlusion** effect darkens creases, holes, intersections and surfaces that are close to each other in real-time. In the real world, such areas tend to block out or occlude **ambient light**, so they appear darker.

URP implements the Screen Space Ambient Occlusion (SSAO) effect as a [Renderer Feature](urp-renderer-feature.md). It works with every **shader** that the Universal **Render Pipeline** (URP) provides as well as any custom opaque Shader Graphs you create.

**Note**: The SSAO effect is a Renderer Feature and works independently from the **post-processing** effects in URP. This effect does not depend on or interact with Volumes.

The following images show a **scene** with the Ambient Occlusion effect turned off, on, and only the Ambient Occlusion texture.

![Scene with Ambient Occlusion effect turned off.](../../uploads/urp/post-proc/ssao/scene-ssao-off.png)

Scene with Ambient Occlusion effect turned off.

![Scene with Ambient Occlusion effect turned on.](../../uploads/urp/post-proc/ssao/scene-ssao-on.png)

Scene with Ambient Occlusion effect turned on.

![Scene with only the Ambient Occlusion texture.](../../uploads/urp/post-proc/ssao/scene-ssao-only-ao.png)

Scene with only the Ambient Occlusion texture.

## Implementation details

The SSAO Renderer Feature uses normal vectors for calculating how exposed each point on a surface is to ambient lighting.

URP 10.0 implements the `DepthNormals` Pass block that generates the the normal texture `_CameraNormalsTexture` for the current frame. By default, the SSAO Renderer Feature uses this texture to calculate Ambient Occlusion values.

If you implement your custom SRP and if you do not want to implement the `DepthNormals` Pass block in your shaders, you can use the SSAO Renderer Feature and set its **Source** property to **Depth**. In this case, Unity does not use the `DepthNormals` Pass to generate the normal vectors, it reconstructs the normal vectors using the depth texture instead.

Selecting the option **Depth** in the **Source** property enables the **Normal Quality** property. The options in this property (Low, Medium, and High) determine the number of samples of the depth texture that Unity takes when reconstructing the normal vector from the depth texture. The number of samples per quality level: Low: 1, Medium: 5, High: 9.