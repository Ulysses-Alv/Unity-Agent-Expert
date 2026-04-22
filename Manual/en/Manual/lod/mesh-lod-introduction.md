# Introduction to Mesh LOD

Mesh **LOD** reduces the number of polygons Unity has to draw and provides automatic LOD creation.

Mesh LOD creates LODs automatically on model import and stores each LOD in the index buffer of the original **mesh**. For more information about LOD generation, refer to the following pages:

* [Generate LODs on import](mesh-lod-generate-lods.md)
* [How the Mesh LOD generator works](mesh-lod-generator.md)

At runtime, the Mesh LOD feature selects the appropriate LOD automatically depending on the size of the mesh on the screen, project-wide and per-object Mesh LOD settings. For more information, refer to [Mesh LOD runtime quality](mesh-lod-quality.md).

When generating LODs automatically, Unity does not create new **GameObjects** or components. This means that the Mesh LOD feature only optimizes the workload for rendering geometry, and does not provide options for configuring material or **Mesh Renderer** settings for less detailed LODs. Being focused on geometry optimization, the feature provides a [smaller memory footprint](mesh-lod-generator.md#mesh-lod-memory-footprint) compared with the LOD Group feature.

## Limitations of the Mesh LOD feature

The Mesh LOD feature has the following limitations:

* The following systems do not support Mesh LOD selection. These systems always select LOD0 with the Mesh LOD feature.

  + Entities Graphics
  + **Particle** System
  + Visual Effect Graph
  + [Static batching](../static-batching)
  + [GPU instancing](../GPUInstancing.md)
* Mesh LOD supports the [LOD cross-fade](lod-transitions-mesh-lod.md) feature with the following limitations:

  + [GPU Resident Drawer](../urp/gpu-resident-drawer.md) must be enabled in a project.
  + Mesh LOD does not support [custom transition zones](lod-transitions-lod-group.md#width).
* When Mesh LOD generates LODs for models with a Skinned Mesh Renderer component, the generator does not take skin weights or blend shape deformations into account during the mesh simplification process. As a consequence, the LOD generator may remove indices that refer to vertices that are important for the intended mesh deformation. For more information, refer to [Skinned Mesh Renderer deformation artifacts](mesh-lod-troubleshooting.md#skinned-mesh-renderer-deformation-artifacts).
* Using Mesh LODs with Skinned Mesh Renderers does not reduce the workload related to calculating mesh deformations. When performing deformations, a Skinned Mesh Renderer deforms LOD0 regardless of which LOD index Unity is currently using for rendering a mesh.
* Mesh LOD only supports meshes with triangle topology.
* Unity doesn’t provide functionality to visualize which Mesh LOD index is being rendered.
* Using Mesh LOD in combination with LOD Group might lead to unexpected outcomes and is not recommended.
* The following APIs do not support automatic Mesh LOD selection and require explicitly specifying which LOD to draw. To specify an LOD in these methods, use the `forceMeshLod` property in the [RenderParams](../../ScriptReference/RenderParams.md) struct.

  + [Graphics.RenderMeshInstanced](../../ScriptReference/Graphics.RenderMeshInstanced.md)
  + [Graphics.RenderMeshIndirect](../../ScriptReference/Graphics.RenderMeshIndirect.md)
  + [Graphics.RenderMeshPrimitives](../../ScriptReference/Graphics.RenderMeshPrimitives.md)

## Additional resources

* [Troubleshooting Mesh LOD visual artifacts](mesh-lod-troubleshooting.md)
* [Introduction to level of detail](../LevelOfDetail.md)
* [LOD Group](lod-group-landing.md)