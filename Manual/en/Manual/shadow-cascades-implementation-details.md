# Implementation details of shadow cascades

This page describes the technical implementation details of shadow cascades.

## Perspective aliasing

A directional light typically simulates sunlight, and a single directional light can illuminate the entire **Scene**. This means that its shadow map covers a large portion of the Scene, which can lead to a problem called perspective aliasing. Perspective aliasing means that shadow map **pixels** close to the **Camera** look enlarged and chunky compared to those farther away.

![Shadows in the distance (A) have an appropriate resolution, whereas shadows close to camera (B) show perspective aliasing.](../uploads/Main/DirShadowAliasing.jpg)

Shadows in the distance (A) have an appropriate resolution, whereas shadows close to camera (B) show perspective aliasing.

Perspective aliasing occurs because different areas of the shadow map are scaled disproportionately by the camera’s perspective. The shadow map from a light needs to cover only the part of the Scene visible to the Camera, which is defined by the Camera’s [view frustum](UnderstandingFrustum.md). If you imagine a simple case where the directional light comes directly from above, you can see the relationship between the frustum and the shadow map.

![Left: object within camera frustum. Right: the far end of the frustum has higher shadow map pixel density than the near end.](../uploads/Main/ShadMapFrustumDiagram.svg)

Left: object within camera frustum. Right: the far end of the frustum has higher shadow map pixel density than the near end.

In this simplified example, the distant end of the frustum is covered by 20 pixels of shadow map, while the near end is covered by only 4 pixels. However, both ends appear the same size on-screen. The result is that the resolution of the map is effectively much less for shadow areas that are close to the camera.

## How shadow cascades work

Perspective aliasing is less noticeable when you use the **Soft Shadows** option, and when you use a higher resolution for the shadow map. However, these solutions use more memory and bandwidth while rendering.

When using shadow cascades, Unity splits the frustum area into two zones based on distance from the camera. The zone at the near end uses a separate shadow map at a reduced size (but with the same resolution). These staged reductions in shadow map size are known as cascaded shadow maps (sometimes called parallel split shadow maps).

![Shadow cascades let Unity distribute shadow resolution more efficiently.](../uploads/Main/ShadMapCascadeDiagram.svg)

Shadow cascades let Unity distribute shadow resolution more efficiently.