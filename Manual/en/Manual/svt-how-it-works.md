# How Streaming Virtual Texturing works

The Streaming Virtual Texturing system (SVT) divides textures into tiles.

At runtime, when SVT samples a texture, it does the following:

* It samples an indirection texture, calculates the non-virtual UVs, and samples a cache texture with these UVs.
* It translates the virtual UVs to a tile ID, and binds an additional render target that receives these tile IDs. The CPU copies this render target back to main memory asynchronously, and processes it on a separate thread, to create requests for the asynchronous read manager to load these tiles from disk into a cache of GPU memory (if they are not already present).

The cost of these runtime operations means that it’s more efficient to group textures together, and to sample them all at the same time. This process is called stacking textures. A group of textures that are sampled at the same time with the same UV coordinates is called a Texture Stack.

SVT issues requests for tiles during frame rendering, so it can take from milliseconds to seconds until the requested tiles load into the GPU cache, and in some cases they might not load into the GPU cache at all. If the requested tile does not load into the cache, SVT has an automatic fallback mechanism: it samples tiles from a higher mipmap level until the requested tile is in the cache. This results in a lower **level of detail** until the tile fully loads.

## SVT in the High Definition Render Pipeline debug view

The High Definition **Render Pipeline** (HDRP) debug view shows how many neighboring screen **pixels** request the same tile.

To open the HDRP debug view, navigate to the Unity top menu and select **Window** > **Analysis** > **Rendering Debugger** > **Rendering**. To view the debug data for SVT, set **Fullscreen Debug Mode** to **RequestedVirtualTextureTiles**.

![The debug view uses distinct colors for every tile ID.](../uploads/Main/svt-how-it-works02.png)

The debug view uses distinct colors for every tile ID.

Pixels that sample the same texture tile have the same colour. The color hue (green, red, and so on) represents the mipmap level to which a tile belongs. For example, all yellowish tiles belong to mipmap 1 of the texture.