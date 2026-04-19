# Enable optional features for Embedded Linux

You can enable the following optional features to improve the performance of your applications.

## Loading screen

When you start the Embedded Linux player, a separate loading screen appears (typically, within 200 ms on our reference systems) containing the 2D image configured in the [Player Settings](embedded-linux-player-settings.md) window. The initial **scene** content is still available except that it loads behind the loading screen.

## Shader cache persistence for GLES3

Embedded Linux supports binary **shader** caching on the device where the Unity Player is installed for better startup timings. The cache is created at runtime after you load a shader. As this cache is written to the temporary folder:`[TEMP]/[COMPANY_NAME]/[PROJECT_NAME]/UnityShaderCache/`, it can be wiped when you restart the system.

To use shader caching when your system restarts, copy the cache into the Player data by following these steps:

1. Deploy the Unity Player to the target system.
2. Run the application and make sure all shaders are touched.
3. Copy all the files from `[TEMP]/[COMPANY_NAME]/[PROJECT_NAME]/UnityShaderCache/` to `[PATH_TO_PLAYER]/Data/UnityShaderCache/`.

**Note:**

* You must refresh the cache every time the player is updated.
* As the cache is device-specific, you must only share it between devices with the exact same hardware and software configuration.
* You can also specify the location of the read-write cache by using `platform-hmi-gfx-cache-path`, described in [Command line arguments](#cli-args).

## Pipeline cache persistence for Vulkan

Embedded Linux supports binary Vulkan pipeline caching on the device where the Unity Player is installed for better startup timings. The binary Vulkan pipeline cache is created at runtime when you use Vulkan pipelines. As this cache is written to the temporary `[TEMP]/[COMPANY_NAME]/[PROJECT_NAME]/vulkan_pso_cache.bin` file, you can wipe it when you restart the system.

To use pipeline caching when your system restarts, copy the cache into the Player data by following these steps:

1. Deploy the Unity Player to the target system.
2. Run the application and make sure all pipelines are being used.
3. Copy `[TEMP]/[COMPANY_NAME]/[PROJECT_NAME]/vulkan_pso_cache.bin` file to `[PATH_TO_PLAYER]/Data`.

**Note:**

* You must refresh the cache every time the player is updated.
* As the cache is device specific, you must only share it between devices with the exact same hardware and software configuration.
* You can also specify the location of the read-write cache by using `platform-hmi-gfx-cache-path`, described in [Command line arguments](#cli-args).

## Startup time logging

Startup time logging is the length of time that it takes an application to start up. It’s often used as a critical metric for system safety and regulatory requirements.

Startup time logging in Embedded Linux include the duration or total time (in milliseconds) since the application is launched. There are two types of Startup time logging:

* **Real:** This is the actual wall or clock time, similar to a stopwatch used for calculating the time.
* **User:** This is the time an application or one of its threads has spent on a CPU core. This can be higher than the Real time if multiple threads are busy when an application is starting up.

To add a startup timing log from C#, use:

`HmiPlatform.LogStartupTiming("log tag");`

The results appear in the following `Player.log` line:

`[TIMING::STARTUP] log tag: Real: xxx ms | User: yyy ms`

It contains the **log tag**, **wall time (xxx)**, and **cpu time (yyy)** in milliseconds since the player’s start time.

You can guard the code using `#if UNITY_EMBEDDED_LINUX_API ... #endif`.

**Note:** Use the same terminology as the **Time** command to refer to wall vs. CPU time. For more information, refer to the main [Linux](https://man7.org/linux/man-pages/man7/time.7.html) manual.

### Example output

```
[TIMING::STARTUP] Initial probing done: Real: 19 ms | User: 11 ms
[TIMING::STARTUP] SDL Initialized: Real: 64 ms | User: 54 ms
[TIMING::STARTUP] Scripting runtime loaded: Real: 97 ms | User: 86 ms
[TIMING::STARTUP] Plugins loaded: Real: 97 ms | User: 87 ms
[TIMING::STARTUP] Engine initialized (nogfx): Real: 104 ms | User: 94 ms
[TIMING::STARTUP] Player Prefs loaded: Real: 104 ms | User: 94 ms
[TIMING::STARTUP] Screen initialized: Real: 139 ms | User: 112 ms
[TIMING::STARTUP] Engine initialized (gfx): Real: 187 ms | User: 161 ms
[TIMING::STARTUP] Gfx initialized: Real: 190 ms | User: 163 ms
[TIMING::STARTUP] Input initialized: Real: 190 ms | User: 163 ms
[TIMING::STARTUP] SPLASH - Begin: Real: 190 ms | User: 164 ms
[TIMING::STARTUP] SPLASH - Primary scene assets loaded (async): Real: 2197 ms | User: 1670 ms
[TIMING::STARTUP] SPLASH - All engine initial states established: Real: 2197 ms | User: 1670 ms
```

Output from a custom event using the Script API

`[TIMING::STARTUP] HELLO!!: Real: 2198 ms | User: 1671 ms`

When you specify `platform-hmi-quit-after-frame` command line argument, the log contains the following information up to frame number `X`, where `X` is the number of frames after which the application quits.

`[TIMING::STARTUP] Frame 1 rendered: Real: 2209 ms | User: 1687 ms`

`[TIMING::STARTUP] Frame 2 rendered: Real: 2210 ms | User: 1692 ms`

## EVDEV input handling with Wayland

To enable the `EVDEV SDL2` input driver while running in Wayland, start the player with the `-platform-embedded-linux-wayland-enable-evdev-input` argument.

## Force to use Wayland

In a system where both X11 and Wayland windowing systems are available, you can force the Unity Player to use Wayland by setting the environment variable to `SDL_VIDEODRIVER=wayland`.

## Command line arguments

You can launch the Unity Embedded Linux Player from the command line and pass arguments to change how the Player executes.

**Note:** All command line arguments have precedence over the settings configured in the Unity Editor.

| **Command line argument** | **Description** |
| --- | --- |
| `-platform-hmi-force-srgb-blit` | Change the `Force SRGB blit` setting to control whether Gamma Color Space (sRGB) is used instead of Linear Color Space during rendering. For more information, refer to [**Player Settings**](embedded-linux-player-settings.md) > **Rendering** > **Force SRGB blit**. |
| `-platform-hmi-log-startup-times` | Log the Player’s startup timing data. This argument is required to use the startup profiling tool `-platform-hmi-quit-after-frame N`. |
| `-platform-hmi-quit-after-frame N` | Log the Player’s startup timing data for the first `N` frames and automatically exit the Player after rendering the `N`th frame. `N` represents the number of frames for which you want to log the Player’s startup time. Use this argument to profile the Player’s startup performance.   **Note**: This argument has an effect only when `-platform-hmi-log-startup-times` is enabled. |
| `-platform-hmi-single-gl-context` | Disable context sharing for OpenGL ES. By default, Unity uses two OpenGL ES contexts, one for startup and another for rendering. This argument forces Unity to use a single context instead. Use this argument to detect and troubleshoot graphics drivers issues.   **Note:** This argument disables multi-display support. |
| `-platform-hmi-cpu-configuration <configuration>` | Specify a CPU configuration for the Player. This argument expects a string containing a combination of the letters: H (high performance core), L (low performance core) and/or D (disable core). The string defines the performance mode of each CPU core.  For example, `DHLL` on a 4-core CPU disables the first core, assigns the second core as high performance, and configures the third and forth as low performance. For more information, refer to [**Player Settings**](embedded-linux-player-settings.md) > **Configuration** > **CPU configuration**. |
| `-platform-hmi-player-data-path` | Enter the directory path on the system where you want to save the `.config` and log files. For more information, refer to [**Player Settings**](embedded-linux-player-settings.md) > **Configuration** > **Player Data path**. |
| `-platform-hmi-force-vsync-count [C]` | Define the number of vertical syncs allowed to pass between each frame. Set it to `0` to disable **vsync** completely, or set it `-1` to use the value set in `QualitySettings`. |
| `-platform-hmi-enable-signal-handlers` | Configure Unity engine’s signal handler setup for embedded platforms. The default value is `1` which enables Unity’s signal handlers. If crash handling or core dump creation doesn’t work correctly, disable this argument by setting it to `0`. |
| `-platform-hmi-gfx-cache-path <path>` | Sets a custom path for the engine to store the GLES shader cache or Vulkan pipeline cache depending on the graphics back-end and platform capabilities. The custom path has precedence over the persistent cache for [GLES3](#persistent-cache-gles) and [Vulkan](#persistent-cache-vulkan).  **Note:** The path provided must exist and be writable.  The [player log file](log-files.md#player) (`Player.log`) contains information about issues and successful usage. |
| `-platform-embedded-linux-enable-gamepadinput` | Change the game controller setting. Refer to [**Player Settings**](embedded-linux-player-settings.md) > **Configuration** > **Enable Game Controllers**. |
| `-platform-embedded-linux-offscreen-video` | Configure the Player to use offscreen rendering driver from SDL2. This is helpful for simulations and setting up a render server. All rendering occurs offscreen but still remains GPU-accelerated.   **Note:** When using this feature, you can limit CPU/GPU usage by adjusting `Application.targetFrameRate`. |
| `-platform-embedded-linux-wayland-enable-evdev-input` | Enable the EVDEV SDL2 input driver while running the Player in Wayland. For more information, refer to [EVDEV input handling with Wayland](#EVDEV). |

## Additional resources

* [Embedded Linux Player Settings](embedded-linux-player-settings.md)
* [Linux manual pages](https://man7.org/linux/man-pages/man7/time.7.html)