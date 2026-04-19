# Sculpt the terrain with a brush

The Raise or Lower **Terrain** tool is a good way to form the overall topography of your world. It’s the least specific of the sculpting tools, so you can use it to create large features that you later refine with the other tools, or with different brushes and settings for the Raise or Lower Terrain tool itself.

## Access the tool

To raise or lower the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Raise or Lower Terrain**.

## Tool options

The Raise or Lower Terrain tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Example workflow for a mountain

Although you can use large brushes to create a mountain quickly, you can get a more natural-looking mountain by building it up more slowly.

A workflow for creating a mountain might look like this:

1. Create the general structure of your entire terrain. For example, to create a natural, non-uniform look for the entire tile, use either:
   * The Raise or Lower Terrain tool with a large, low-opacity brush.
   * The [Add noise to your terrain](terrain-Sculpt-Noise-Landing.md).
2. Use smaller, higher-opacity brushes to create the mountain.
3. Use small brushes to add fine details to the mountain or other important areas.
4. Use additional tools, such as [Smooth Height](terrain-SmoothHeight.md), to improve the blend between the mountain and the surrounding terrain.

This method of working from the general to the specific can create a better integration between different areas.

## Additional resources

* [Stamp a shape onto the terrain](terrain-StampTerrain.md)
* [Work with heightmaps](terrain-Heightmaps.md)
* [Textures and layers](terrain-Textures-Landing.md)
* [Brushes](class-Brush.md)