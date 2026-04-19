# Introduction to the Frame Debugger

Debug frames to help identify rendering artefacts and other issues. Unity includes a dedicated [Frame Debugger](frame-debugger-window.md) that can pause the application on a particular frame and display the list of rendering events that constitute the frame. The Frame Debugger can step through each event and display the graphical state of the **scene** at that point in the rendering process. You can use this to find where graphical issues arise or to just see how Unity constructs the scene from graphical elements.

## Render Pipeline compatibility

| **Feature** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom Scriptable Render Pipeline (SRP)** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Frame Debugger** | Yes | Yes | Yes | Yes |

**Note:** The Frame Debugger updates every frame and you can’t pause it. If the number of render passes or draw calls changes rapidly, the Frame Debugger might update its list of events too quickly to read.

If you need information about a frame that Unity’s Frame Debugger doesn’t provide, there are 3rd-party frame debugging programs that support Unity. The Unity Editor supports native launching and frame capturing for [RenderDoc](RenderDocIntegration.md). You can also build a standalone Player and attach a supported frame debugger. For information about supported frame debugging software, see [Profiling tools](performance-profiling-tools.md).