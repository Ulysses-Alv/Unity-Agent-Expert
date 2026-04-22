# Debug a frame

To debug a frame using the Frame Debugger:

1. Open the Frame Debugger (menu: **Window** > **Analysis** > **Frame Debugger**).
2. Use the target selector to select the process to attach the Frame Debugger to. If you want to debug a frame in the Unity Editor, set this to **Editor**. If you want to debug a frame in a built application, see [Attach the Frame Debugger to a built project](FrameDebugger-attach.md).
3. Click **Enable**. When you do this, the Frame Debugger captures a frame. It populates the Event Hierarchy with the draw calls and other events that constitute the frame and renders the frame in the Game view.
     
   **Note**: If your application is running, the Frame Debugger pauses it.
4. Select an event from the Event Hierarchy to view the **scene** as it appears up to and including that event. This also displays information about the event in the Event Information Panel. You can use the previous event and next event button, the arrow keys, or the event scrubber to move through the frame linearly. If you don’t know which event Unity renders the geometry you want to debug in, these navigation tools are useful to move through the events linearly until you find it.

When a draw call event corresponds to the geometry of a **GameObject**, Unity highlights that GameObject in the [Hierarchy](Hierarchy.md).

If an event renders into a [RenderTexture](class-RenderTexture.md), Unity displays the contents of that RenderTexture in the Game view and Frame Debugger window. This is useful for inspecting how various off-screen render targets build up. For example:

![Viewing events that accumulate to produce the diffuse G-buffer during deferred rendering.](../uploads/Main/frame-debugger-event-accumulation.gif)

Viewing events that accumulate to produce the diffuse G-buffer during deferred rendering.

## Debug a Render Graph frame

If your project uses the [Render Graph system](urp/render-graph-introduction.md), the Frame Debugger displays the following in the [Event Hierarchy panel](frame-debugger-window-event-hierarchy.md):

* A parent rendering event called **ExecuteRenderGraph**.
* Child rendering events called **(RP <render-pass>:<subpass>)**, where **<render-pass>** is the render pass number and **<subpass>** is the subpass number.

The Frame Debugger displays only render passes that contain a draw call.