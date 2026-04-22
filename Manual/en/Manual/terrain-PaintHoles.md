# Paint holes in the terrain

To simulate cave openings and overhanging formations, such as cliffs, use the Paint Holes tool to hide a part of the **terrain**. Then use other Editor tools, such as ProBuilder, to create the cave or overhang geometry.

## Access the tool

To paint a hole, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Paint Holes**.

**Note**: If your project uses the Universal **Render Pipeline** (URP), enable **Terrain Holes** in the [URP Asset](urp/universalrp-asset.md). Otherwise, terrain holes disappear when you build your application.

## Tool options

The Paint Holes tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Hide or smooth aliased edges

The Paint Holes tool is texture-based, so it can create aliased edges around holes. You can:

* Hide the aliased edges with geometry, such as rock meshes.
* Increase the **Heightmap** Resolution to add more texels. This increases performance requirements.

  **Note:** When you change resolutions, the Editor copies higher resolution texture data into lower resolution texture data. In other words, it downscales. However, it doesn’t upscale.

To try a higher resolution:

1. Erease the parts of the hole you want to redraw with the Paint Holes tool.
2. In **Terrain Settings**, increase the Heightmap Resolution.
3. Draw the hole again.

For more information, refer to [Terrain Settings reference](terrain-OtherSettings.md).

## Hole collision

A hole in the terrain is automatically excluded from the Terrain **Collider**, so that characters and objects can fall through it. You don’t need to manually adjust the **Terrain Collider**.

If you are creating holes in a script-based workflow, ensure that the resolution of the hole’s mask texture is equal to the resolution of the Terrain’s heightmap minus 1 (`-1`). For example, if the Terrain’s heightmap resolution is `1025x1025`, the hole’s mask texture needs to be `1024x1024`. For more information, refer to [TerrainData-holesResolution](../ScriptReference/TerrainData-holesResolution.md).

## Using a custom material

The holes functionality depends on the **shader**. To allow holes in a custom Terrain material, you can either:

* Use the Terrain shader the Editor provides.
* Use your own shader with the following code snippet in the fragment shader:

```
  #ifdef _ALPHATEST_ON
      ClipHoles(IN.uvMainAndLM.xy);
  #endif
```

## Additional resources

* [Digging into Terrain Paint Holes in Unity 2019.3](https://blogs.unity3d.com/2020/01/31/digging-into-terrain-paint-holes-in-unity-2019-3/)
* [Primary sculpting tools](terrain-Basic-Sculpting.md)
* [Advanced sculpting tools](terrain-Sculpt.md)
* [Change existing height with effects](terrain-Effects.md)
* [Transform terrain features](terrain-Transform.md)
* [Erode the terrain](terrain-Erosion.md)