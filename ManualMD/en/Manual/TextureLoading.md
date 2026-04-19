# Texture optimization

Techniques for speeding up texture loading, and reducing how much memory textures use.

| **Page** | **Description** |
| --- | --- |
| [Loading textures in the background](LoadingTextureandMeshData-introduction.md) | Resources for loading textures asynchronously to prevent your application pausing when textures load. |
| [Optimizing GPU texture memory with mipmap streaming](TextureStreaming.md) | Use mipmap streaming to limit the size of textures in GPU memory. |
| [Sparse Textures](SparseTextures.md) | Learn about creating textures that Unity can break down into rectangular tiles, so it can load only the tiles it needs. |
| [Optimize high-resolution textures with Streaming Virtual Texturing](svt-streaming-virtual-texturing.md) | Resources for using Streaming Virtual Texturing (SVT) to progressively upload tiles when they’re needed, to reduce GPU memory usage and texture loading times with high-resolution textures. |