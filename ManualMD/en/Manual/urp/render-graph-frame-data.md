# Frame data in the render graph system in URP

Fetch the textures that the Universal **Render Pipeline** (URP) creates for the current frame or previous frames, for example a color texture or a depth texture.

| **Page** | **Description** |
| --- | --- |
| [Get data from the current frame](accessing-frame-data.md) | Fetch the textures URP creates for the current frame. |
| [Get data from previous frames](render-graph-get-previous-frames.md) | To fetch the previous frames the **camera** rendered, use the `UniversalCameraData.historyManager` API. |
| [Add textures to the camera history](render-graph-add-textures-to-camera-history.md) | To add your own texture to the camera history, create a camera history type to store the texture between frames. |
| [Get the current framebuffer from GPU memory](render-graph-framebuffer-fetch.md) | To speed up rendering, use the `SetInputAttachment` API to read the frame that Unity has rendered so far. |
| [Frame data textures reference](render-graph-frame-data-reference.md) | Explore the textures you can fetch from the current frame or previous frames. |