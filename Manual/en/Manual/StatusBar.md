# Status bar

The status bar provides notifications about various Unity processes, and quick access to related tools and settings. You cannot move or rearrange the status bar.

The status bar displays the following:

![The status bar in the main Editor window](../uploads/Main/status-bar-overview.png)

The status bar in the main Editor window

1. The most recent message logged to the [Console window](Console.md). Click the message to open the Console window.
2. A global progress bar for various asynchronous tasks (for example, shader compilation, lightmap baking, and occlusion culling). Click the progress bar to open the [Background Tasks window](BackgroundTasksWindow.md), which displays progress for individual tasks and subtasks.
3. The current [code optimization mode](managed-code-debugging.md). Click the code optimization icon to switch between debug mode and release mode.
4. The cache server status. Click the icon to get additional information about the [cache server](UnityAccelerator.md), and re-establish a lost connection.
5. An activity indicator (spinner) that shows when Unity compiles C# **scripts** or runs asynchronous tasks.