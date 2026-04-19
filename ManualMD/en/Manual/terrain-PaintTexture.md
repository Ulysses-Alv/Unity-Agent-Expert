# Paint Texture

Use the **Paint Texture** tool to add textures, such as grass, snow, or sand, to your **Terrain**. It allows you to draw areas of tiled texture directly onto the Terrain. In the Terrain **Inspector**, click the [**Paint Terrain**](terrain-Tools.md) icon, and select **Paint Texture** from the list of Terrain tools.

![Paint Texture tool in the Terrain Inspector](../uploads/Main/1.3.3-PaintTexture_grey.png)

Paint Texture tool in the Terrain Inspector

To access the **Paint Texture** tool from an overlay:

1. In the **Terrain Tools** overlay, select **Materials Mode** ![Material Mode Menu button](../uploads/Main/terrainOverlays-MaterialModeMenuButton.png). Materials Mode tools display at the end of the **Terrain Tools** overlay.
2. From the available Materials Mode tools on the **Terrain Tools** overlay, select **Paint Texture** ![Paint Texture](../uploads/Main/terrainOverlays-PaintTextureButton.png).

To configure the tool, you must first click the **Edit Terrain Layers** button to add [Terrain Layers](class-TerrainLayer.md). The first Terrain Layer you add flood-fills your Terrain with the configured texture. You can add multiple Terrain Layers. However, the number of Terrain Layers each tile supports depends on your specific **render pipeline**. See the [Rendering performance section on Terrain Layers](class-TerrainLayer.md#Performance) for more information.

Next, you must choose a [Brush](class-Brush.md) for painting. Brushes are Assets based on Textures, which define the shape of a brush. Select from the built-in Brushes or create your own, then adjust the **Brush Size** and **Opacity** (strength of the applied effect) of the brush.

Finally, in the **Scene** view, click and drag the cursor across the Terrain to create areas of tiled texture. You can paint across tile boundaries to blend adjacent regions with a natural, organic look. Be aware, however, that the Terrain system adds the selected Terrain Layer to any Terrain you paint on, which might affect performance as mentioned above.

---

* 2019–04–17 Page published
* Updated content to match new UI