# Mesh Filter component reference

[Switch to Scripting](../ScriptReference/MeshFilter.md "Go to MeshFilter page in the Scripting Reference")

A **Mesh Filter** component holds a reference to a mesh. It works with a [Mesh Renderer](https://docs.unity3d.com/Manual/class-MeshRenderer.html) component on the same **GameObject**; the Mesh Renderer renders the mesh that the Mesh Filter references.

To render a deformable mesh, use a [Skinned Mesh Renderer](https://docs.unity3d.com/Manual/class-SkinnedMeshRenderer.html) instead. A Skinned Mesh Renderer component does not need a Mesh Filter component.

## Mesh Filter Inspector reference

| **Property** | **Function** |
| --- | --- |
| **Mesh** | A reference to a [mesh asset](class-Mesh.md).  To change the mesh asset that the Mesh Filter component references, use the picker (circle icon) next to the mesh’s name.  **Note:** The settings for other components on this GameObject don’t change when you change the mesh that the Mesh Filter references. For example, a Mesh Renderer doesn’t update its settings, which might cause Unity to render the mesh with unexpected properties. If this happens, adjust the settings of the other components as needed.  Corresponds to the [MeshFilter.mesh](https://docs.unity3d.com/ScriptReference/MeshFilter-mesh.html) property. |

MeshFilter