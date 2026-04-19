# Terrain Lit shader material Inspector window reference for URP

Explore the properties and settings you can use to customize the default Terrain Lit **shader** in the Universal **Render Pipeline** (URP).

URP uses the Terrain Lit shader for [Terrain](../script-Terrain.md). This shader is a simpler version of the [Lit shader](lit-shader.md).

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Enable Height-based Blend** |  | Specifies whether URP renders only the [Terrain Layer](../class-TerrainLayer.md) with the greatest height value for a particular pixel. URP takes the height values from the blue channel of the Mask Map of the Terrain Layer.  If you disable this property, URP blends the Terrain Layers based on the weights in the splatmap textures, and adds the **Opacity as Density** parameter to the [Terrain Layer Inspector window](../class-TerrainLayer.md). |
|  | **Height Transition** | Sets the size in world units of the smooth transition area between Terrain Layers. |
| **Enable Per-pixel Normal** |  | Samples the normal map that Unity generates from the terrain heightmap, for each pixel. Enabling this parameter provides high-resolution mesh normals, even if the mesh is low resolution.  **Note**: This parameter works only if you enable **Draw Instanced** in the [Terrain Settings](../terrain-OtherSettings.md). |