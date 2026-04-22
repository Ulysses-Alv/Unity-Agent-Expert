# Enable optional features for QNX

You can enable the following optional features to improve the performance of your applications.

## Command line arguments

You can launch the Unity QNX Player from the command line and pass arguments to change how the Player executes.

**Note**: All command line arguments have precedence over the settings configured in the Unity Editor.

| **Command line argument** | **Description** |
| --- | --- |
| `-platform-qnx-graphics-conf` | Configure the path to the `graphics.conf` to override [auto detection](qnx-autodetect-plugins.md). |
| `-platform-hmi-force-srgb-blit` | Change the `Force SRGB blit` setting to control whether Gamma Color Space (sRGB) is used instead of Linear Color Space during rendering. For more information, refer to [**Player Settings**](qnx-player-settings.md) > **Rendering** > **Force SRGB blit**. |
| `-platform-hmi-log-startup-times` | Log the Player’s startup timing data. This argument is required to use the startup profiling tool `-platform-hmi-quit-after-frame N`. |
| `-platform-hmi-quit-after-frame` | Log the Player’s startup timing data for the first `N` frames and automatically exit the Player after rendering the `N`th frame. `N` represents the number of frames for which you want to log the Player’s startup time. Use this argument to profile the Player’s startup performance.   **Note**: This argument has an effect only when `-platform-hmi-log-startup-times` is enabled. |
| `-platform-hmi-single-gl-context` | Disable context sharing for OpenGL ES. By default, Unity uses two OpenGL ES contexts, one for startup and another for rendering. This argument forces Unity to use a single context instead. Use this argument to detect and troubleshoot graphics drivers issues.   **Note:** This argument disables multi-display support. |
| `-platform-hmi-cpu-configuration <configuration>` | Specify a CPU configuration for the Player. This argument expects a string containing a combination of the letters: H (high performance core), L (low performance core) and/or D (disable core). The string defines the performance mode of each CPU core.  For example, `DHLL` on a 4-core CPU disables the first core, assigns the second core as high performance, and configures the third and forth as low performance. For more information, refer to [**Player Settings**](qnx-player-settings.md) > **Configuration** > **CPU configuration**. |
| `-platform-hmi-player-data-path` | Enter the directory path on the system where you want to save the `.config` and log files. For more information, refer to [**Player Settings**](qnx-player-settings.md) > **Configuration** > **Player Data path**. |
| `-platform-hmi-force-vsync-count [C]` | Define the number of vertical syncs allowed to pass between each frame. Set it to `0` to disable **vsync** completely, or set it `-1` to use the value set in `QualitySettings`. |
| `-platform-hmi-enable-signal-handlers` | Configure Unity engine’s signal handler setup for embedded platforms. The default value is `1` which enables Unity’s signal handlers. If crash handling or core dump creation doesn’t work correctly, disable this argument by setting it to `0`. |

## Startup time logging

Startup time logging is the length of time that it takes an application to start up. It’s often used as a critical metric for system safety and regulatory requirements.

Startup time logging in QNX devices include the duration or total time from the time the application launches. There are two types of Startup time logging:

* **Real:** This is the actual wall or clock time, similar to a stopwatch used for calculating the time.
* **User:** This is the time an application or one of its threads has spent on a CPU core. This can be higher than the Real time if multiple threads are busy when an application is starting up.

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

## Webcam

**Note:** Unity’s support for Webcam in QNX is currently experimental.

#### Prerequisites

* QNX 7.1
* `libcamapi` and its dependencies installed on the system (will be loaded dynamically)
* **Camera** supporting `NV12` format

[Webcam](../ScriptReference/WebCamTexture.md) usage is optional in QNX and it’s only supported on QNX 7.1. For more information, refer to the [Webcam](../ScriptReference/WebCamTexture.md) documentation.

## Additional resources

* [Autodetect plug-ins for QNX](qnx-autodetect-plugins.md)
* [Support touch input for QNX](qnx-touch-input.md)