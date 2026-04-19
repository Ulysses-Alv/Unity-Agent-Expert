# Introduction to Dedicated Server

Dedicated Server refers to a computer that’s optimized to run server applications.

The Dedicated Server [build target](../ScriptReference/BuildTarget.md) is a sub-target of the three Desktop Platforms (Linux, macOS, and Windows) that’s optimized for running as a dedicated server. Using Dedicated Server build target, you can cut back on your server builds’ CPU, memory, and disk space requirements.

Server builds often contain unnecessary assets and compiled code for headless server processes. The build data might include artifacts such as audio files, textures, meshes, and **shaders**. In a multiplayer context, rendering and asset management processes occur unnecessarily when building and executing server runtimes.

The goal of the Dedicated Server build target is to reduce the resource demand of server builds, including the disk size, the size in memory, and the CPU usage. Most of the optimizations take place through [stripping code](managed-code-stripping.md) and assets that aren’t necessary for a server build.

## Additional resources

* [Desktop headless mode](desktop-headless-mode.md)