# Tilemap Renderer component reference

When you create a [tilemap](./create-tilemap.md), Unity automatically attaches the **Tilemap Renderer** component to the tilemap **GameObject** you created. You can determine how Unity renders tiles set on the tilemap by adjusting the following property settings of the Tilemap Renderer component.

## Properties

| Property | Description |
| --- | --- |
| **Sort Order** | Select the direction that Unity sorts the tiles on the selected tilemap from. |
| **Mode** | Select the rendering mode of the renderer. |
| **Chunk** | Select this mode to have the renderer group tiles by their location and the textures used by their sprites, and then batch them together for rendering. Groups with the same textures can be dynamically batched together even if they’re out of locational sequence of each other (based on the direction selected for the renderer’s **Sort Order**).    Select this mode for the best rendering performance with tilemaps.   **Note:** This mode isn’t compatible with the [Scriptable Render Pipeline Batcher](../../SRPBatcher.md). |
| **Individual** | Select this mode to have the renderer render each tile individually, taking into account their location and position in the **Sort Order**. This mode enables the **sprites** on the tiles to interact with other renderers in the **scene** or with a [Custom Sorting Axis](isometric-tilemaps/renderer/tilemap-renderer-isometric-modes.md). |
| **SRP Batch** | Select this mode to make the Tilemap Renderer component compatible with the [Scriptable Render Pipeline Batcher](../../SRPBatcher.md), once the compatibility requirements are also met. **Note:** This mode is only supported in the [Universal Render Pipeline](../../universal-render-pipeline.md) version 15 onwards.   This rendering mode groups tiles by their location and the textures used by their sprites, and then batch them together for rendering. Groups with the same textures will be dynamically batched together only if they’re in locational sequence of each other (based on the direction selected for the renderer’s **Sort Order**). |
| **Detect Chunk Culling Bounds** | Select the way the renderer detects the [bounds](../../../ScriptReference/Tilemaps.TilemapRenderer-chunkCullingBounds.md) used for the culling of [tilemap chunks](../../../ScriptReference/Tilemaps.TilemapRenderer-chunkSize.md). These bounds expand the boundary of tilemap chunks to ensure that oversize sprites aren’t clipped during culling. |
| **Auto** | The Renderer automatically inspects the Sprites used by the Tiles to determine the expanded culling bounds to use. |
| **Manual** | The values used to extend the bounds for culling of the tilemap chunks are manually set instead of with the Editor’s automatic detection. |
| **Chunk Culling Bounds** | This property is visible only when **Detect Chunk Culling Bounds** is set to **Manual** |
| **XYZ** | Enter the values (in Unity units) to extend the culling bounds by. |
| **Mask Interaction** | Set how the Tilemap Renderer behaves when it interacts with a [Sprite Mask](../../sprite/mask/mask-landing.md). |
| **None** | Select this option to have the Tilemap Renderer not interact with any Sprite Mask in the scene. This is the default option. |
| **Visible Inside Mask** | Select this option to have the Tilemap Renderer only render areas of the tilemap that the Sprite Mask overlays. |
| **Visible Outside Mask** | Select this option to have the Tilemap Renderer render the entire tilemap but subtract the areas that the Sprite Mask overlays. |
| **Material** | Select the material used to render the sprite textures. |

### Additional Settings

Properties in the **Additional Settings** section of the tilemap renderer component.

| Property | Description |
| --- | --- |
| **Sorting Layer** | Select an existing [Sorting Layer](../../class-TagManager.md) from the **Sorting Layer** dropdown or create a new Sorting Layer for this tilemap. |
| **Order in Layer** | Set the render priority of the tilemap within its Sorting Layer. Unity renders lower numbered layers first, and higher numbered layers overlap those below. |

## Additional resources

* [TilemapRenderer class](../../../ScriptReference/Tilemaps.TilemapRenderer.md) (Scripting API)