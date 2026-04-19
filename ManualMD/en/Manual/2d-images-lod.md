# 2D images for low level of detail (LOD)

Techniques and resources for using 2D **billboards** to simplify 3D meshes that are far away.

| **Topic:** | **Description:** |
| --- | --- |
| [Applying 2D billboards for low LOD meshes](applying-2d-billboards-low-lod-meshes.md) | Understand how Unity uses Billboard Renderers and Billboard assets to create a 2D billboard for low-LOD rendering. |
| [Billboard Renderer component reference](class-BillboardRenderer.md) | Explore the properties in the Billboard Renderer to to customize how Unity renders billboards and their interaction with lighting. |
| [Billboard asset reference](class-BillboardAsset.md) | Explore the properties in the Billboard asset. |

## Render pipeline information

How you work with billboards depends on the **render pipeline** you use.

| **Feature name** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Built-in Render Pipeline** |
| --- | --- | --- | --- |
| **Billboard Renderer component** | Yes. | Yes.  Only with VFX Graph. | Yes. |