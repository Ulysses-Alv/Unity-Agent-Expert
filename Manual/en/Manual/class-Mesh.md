# Mesh asset Inspector window reference

[Switch to Scripting](../ScriptReference/Mesh.md "Go to Mesh page in the Scripting Reference")

A **Mesh** asset represents a mesh in your Unity project.

Unity creates mesh assets by default when it imports [models](models.md). You can also create mesh assets directly in Unity, for example, by [creating a mesh with code](creating-meshes.md) and saving the results as an asset.

Select a mesh asset in the ****Project window**** to open the **Mesh **Inspector****. This Inspector displays information about the mesh, such as the total number of vertices or the names of any blend shapes the mesh has. It also displays a visual preview of the mesh.

## Mesh information

The Mesh Inspector displays information about how the mesh **asset stores** its data, but not the values in the data itself. For example, the Inspector shows that Unity stores the value for Position as a set of three Float32 values for each vertex, but it doesn’t show the Position value of a specific vertex.

![The Mesh Inspector](../uploads/Main/mesh-inspector.png)

The Mesh Inspector

| **Property** | **Description** |
| --- | --- |
| **Vertices** | The total number of vertices in the mesh, and how much on-disk storage and runtime memory Unity uses to store the vertex data for the mesh. For each vertex attribute, Unity shows how much on-disk storage and runtime memory Unity uses to store that attribute, and the data type that Unity uses for it.  For more information about this data, see [Mesh data](AnatomyofaMesh.md). For information about how Unity can compress data in a mesh, see the [Compressing mesh data](compressing-mesh-data-optimization.md). |
| **Indices** | The total number of indices in the mesh, how much on-disk storage and runtime memory Unity uses to store this data, and information about the indices for each sub-mesh.  For more information about this data, see [Mesh data](AnatomyofaMesh.md). |
| **Skin** | This section is only visible on a [skinned mesh](mesh-introduction.md#deformable-meshes).  The options are:  * **BlendWeight**: How many bones can affect a vertex. * **BlendIndices**: The indices of the bones that affect skinned vertices. * **Blend Shapes**: The total number of blend shapes in your mesh and lists their names. The number of frames next to each blend shape describes how many key frames each one has. |
| **Other** | Shows additional information.  The options are:  * **Bounds Center**: A bounding box describes where the outermost edges of a mesh lie. The “Bounds Center” stores the location of the center of the mesh’s bounding box. * **Bounds Size**: Each value represents the length of the bounding box edges in X, Y and Z directions. * **Read/Write Enabled**: The value of the [Mesh.isReadable](../ScriptReference/Mesh-isReadable.md) property, which determines whether Unity can access and modify certain mesh data at runtime. You can set this value with the **Read/Write Enabled** checkbox in the [Model Import Settings window](FBXImporter-Model.md). |

## Mesh preview

This part of the Inspector allows you to preview the appearance of a mesh and explore the mesh data in a visual way.

You can use the following properties in the UI to configure the view:

| **Property** | **Description** |
| --- | --- |
| **View mode** | Provides different ways of visualizing the mesh. For more information, refer to [View mode dropdown](#viewmode). |
| **Wireframe** | Not available in UV Layout mode.  If enabled, the preview displays a visualisation of the vertices and edges. Otherwise, it doesn’t. |
| **Channel picker** | Only available in UV Checker mode.  This drop-down menu allows you to change which UV channel the preview displays. If the mesh only uses one UV channel, then only “Channel 0” is enabled. |
| **Texture size** | Only available in UV Checker mode.  This slider allows you to change the size of the checkered texture on the mesh’s UV layout. |
| **Blend shape picker** | Only available in Blendshapes mode.  This drop-down allows you to change the blend shape that Unity visualizes in the mesh preview. When you select a different blend shape, Unity no longer applies the previous blend shape to the mesh preview. |

### View mode dropdown

For more information, refer to [Use the Mesh asset Inspector Preview](view-mesh-data-visualizations.md).

| **Value** | **Description** |
| --- | --- |
| **Shaded** | Provides a visualization of the mesh with a basic light. This is the default view. For more information, see [Shaded view](view-mesh-data-visualizations.md#shaded). |
| **UV Checker** | Applies a checkerboard texture to the mesh to visualize how the mesh’s UV layout maps textures. For more information, see [UV Checker view](view-mesh-data-visualizations.md#uv-checker) |
| **UV Layout** | Displays how the vertices of the mesh are organized in the unwrapped UV layout. This view disables the Wireframe toggle. For more information, see [UV Layout view](view-mesh-data-visualizations.md#uv-layout) |
| **Vertex Color** | Visualizes any vertex colors that the vertices in this mesh have. If no vertices have a vertex color, this menu item is disabled. For more information, see [Vertex Color view](view-mesh-data-visualizations.md#vertex-color) |
| **Normals** | Visualizes the relative directions of the normals in the mesh with color. For more information, see [Normals view](view-mesh-data-visualizations.md#normals) |
| **Tangents** | Visualizes the tangent data in the mesh with color. For more information, see [Tangents view](view-mesh-data-visualizations.md#tangents) |
| **Blendshapes** | Lets you preview blend shape deformations to the mesh. If the mesh has no blend shapes, this menu item is disabled. For more information, see [Blendshapes view](mesh-data-deformable-meshes.md#blend-shapes) |