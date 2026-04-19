# Optimizing GPU texture memory with mipmap streaming

Use mipmap streaming to limit the size of textures in GPU memory.

| Page | Description |
| --- | --- |
| [Introduction to mipmap streaming](TextureStreaming-introduction.md) | How mipmap streaming works, and its limitations. |
| [Use mipmap streaming](TextureStreaming-use.md) | Enable mipmap streaming, make a texture work with mipmap streaming, and change the mipmap level of a texture. |
| [Configure mipmap streaming](TextureStreaming-configure.md) | Enable or disable mipmap streaming on cameras, set which mipmap levels a **camera** uses, and set a memory budget for textures. |
| [Override the mipmap level of a texture](TextureStreaming-override-mipmap-level.md) | Request that Unity overrides the mipmap level of a texture. |
| [Preload mipmap levels](TextureStreaming-preload.md) | Use the API to preload mipmap levels to avoid textures pop-in when you enable a camera. |
| [Analyze mipmap streaming](TextureStreaming-analyze.md) | Understand and debug mipmap streaming in your **scene**. |