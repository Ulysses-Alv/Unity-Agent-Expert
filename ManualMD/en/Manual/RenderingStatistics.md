# Rendering Statistics window reference

The Game view includes a Rendering Statistics window that displays real-time rendering information about your application during Play mode. To open this window, click the **Stats** button in the top right corner of the Game view. Unity displays the Statistics window as an overlay in the top right of the Game view. The rendering statistics shown in the Graphics section window are useful for optimizing performance. The exact set of statistics available varies according to the build target.

![Statistics window showing real-time rendering statistics. ](../uploads/Main/GameViewStats.png)

Statistics window showing real-time rendering statistics.

## **Statistics**

The Graphics section of the Statistics window contains the following information:

| **Statistic** | **Description** |
| --- | --- |
| ****FPS**** | The current number of frames Unity is able to draw per second. |
| **CPU** | **Main**: The total amount of time taken to process one frame. This number includes the time Unity takes to process the frame update of your application and the time Unity takes in the Editor to update the Scene view, other Editor Windows, or process Editor-only tasks. **Render**: The amount of time taken to render one frame. This number includes the time Unity takes to process the frame update for the Game view; it doesn’t include the time Unity takes in the Editor. |
| **Batches** | The total number of [draw call batches](DrawCallBatching.md) Unity processes during a frame. This number includes static, dynamic, and [instance](GPUInstancing.md) batches. |
| **Saved by batching** | The number of draw calls Unity combined into batches. To ensure good draw call batching, share materials between different ****GameObjects**** as often as possible. Batches group draw calls with the same render state, so changing the material causes Unity to create a new batch. |
| **Tris** | The number of [triangles](../ScriptReference/Mesh-triangles.md) Unity processes during a frame. This value is important when [optimizing for low-end hardware](OptimizingGraphicsPerformance.md). |
| **Verts** | The number of [vertices](../ScriptReference/Mesh-vertices.md) Unity processes during a frame. This value is important when [optimizing for low-end hardware](OptimizingGraphicsPerformance.md). |
| **Screen** | The resolution of the screen, along with the amount of memory the screen uses. |
| **SetPass** | The number of times Unity switches which **shader** pass it uses to render GameObjects during a frame. A shader might contain several shader passes and each pass renders GameObjects differently. Each pass requires Unity to bind a new shader, which might introduce CPU overhead. |
| **Shadow casters** | The number of GameObjects in the frame that cast shadows. |
| **Visible skinned meshes** | The number of [Skinned Mesh Renderers](class-SkinnedMeshRenderer.md) in the frame. |
| **Animation components playing** | The number of [Animation](class-Animation.md) components playing during the frame. |
| ****Animator components** playing** | The number of [Animator](class-Animator.md) components playing during the frame. |

For more detailed information about your application’s rendering performance, see the [Rendering module of the Profiler window](ProfilerRendering.md).