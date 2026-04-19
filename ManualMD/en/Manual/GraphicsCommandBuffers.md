# CommandBuffer fundamentals in the Built-In Render Pipeline

A [CommandBuffer](../ScriptReference/Rendering.CommandBuffer.md) holds a list of rendering commands (such as setting the render target, or drawing a given mesh). You can instruct Unity to schedule and execute those commands at various points in the Built-in **Render Pipeline**, which allows you to customize and extend Unity’s rendering functionality.

![Blurry refraction, using Command Buffers.](../uploads/Main/RenderingCommandBufferBlurryRefraction.jpg)

Blurry refraction, using Command Buffers.

You can execute CommandBuffers immediately using the [Graphics.ExecuteCommandBuffer](../ScriptReference/Graphics.ExecuteCommandBuffer.md) API, or you can schedule them so that they occur at a given point in the render pipeline. To schedule them, use the [Camera.AddCommandBuffer](../ScriptReference/Camera.AddCommandBuffer.md) API with the [CameraEvent enum](../ScriptReference/Rendering.CameraEvent.md), and the [Light.AddCommandBuffer](../ScriptReference/Light.AddCommandBuffer.md) API with the [LightEvent enum](../ScriptReference/Rendering.LightEvent.md). To see when Unity executes CommandBuffers that you schedule in this way, see [CameraEvent and LightEvent order of execution](built-in-rendering-order.md).

For a full list of the commands that you can execute using CommandBuffers, see the [CommandBuffer API documentation](../ScriptReference/Rendering.CommandBuffer.md). Note that some commands are supported only on certain hardware; for example, the commands relating to **ray tracing** are supported only in DX12.

* [Introduction to render pipelines](render-pipelines-overview.md)
* [Execute rendering commands in a Scriptable Render Pipeline](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/manual/index.html)
* [Custom rendering and post-processing in URP](urp/customizing-urp.md)