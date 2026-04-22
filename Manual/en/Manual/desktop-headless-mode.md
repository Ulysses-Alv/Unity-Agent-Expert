# Desktop headless mode

The Dedicated Server build target is similar to the desktop headless mode. The only difference is that the Dedicated Server build target is optimized to increase memory and CPU performance when running as a networked application.

Desktop headless mode allows you to run applications in batchmode on any desktop platform without initializing the graphics device. You can run desktop headless mode by passing the `-batchmode` and `-nographics` command line arguments when executing the Player. You can’t select headless mode from the Unity Editor build settings, but you can add the `-batchmode` and `-nographics` command line arguments to effectively create a headless build.

Desktop headless mode doesn’t perform any optimizations for running a build as a dedicated server. Although the Dedicated Server build option performs additional optimizations, you might still want to use desktop headless mode for other purposes, such as automated testing on CI/CD platforms.