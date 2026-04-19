# Deploy a QNX project

Unity uses EGL handled by SDL2, which requires SDL to dynamically load `libEGL` and `libGLESv2` from the `graphics.conf file`. Unity does not parse the `conf` file but instead uses environment variables to locate these libraries.

## Setup

Use the following instructions to deploy QNX.

1. Use one of the methods to locate the `graphics.conf` file that your screen loads:

   * Start `screen` with the `-c [path/to/graphics.conf]` option.
   * Let your `screen` automatically find the `graphics.conf` file in the folder inside `GRAPHICS_ROOT`.
2. Make sure the folder that contains `graphics.conf` is part of `LD_LIBRARY_PATH`.
3. Locate `begin egl display 1` in `graphics.conf`:

   * The line starting with `egl-dlls`should contain `libEGL[-_tag].so`, which is the required `libEGL` (for example, `libEGL_viv.so`).
   * The line starting with `glesv2-dlls` should contain `libGLESv2[-_tag]`, which is the `libGLESv2` (for example, `libGLESv2_viv.so`).
   * Both libraries should be in the same folder as `graphics.conf`.
   * Both library file names must be set up as environment variables.
4. If you are using `ksh`, set the following environment variables.

   ```
   SDL_VIDEO_EGL_DRIVER=[name_of_libEGL_in_graphics_conf].so (e.g., run export SDL_VIDEO_EGL_DRIVER=libEGL_viv.so)
   SDL_VIDEO_GL_DRIVER=[name_of_libGLESv2_in_graphics_conf].so (e.g., run export SDL_VIDEO_GL_DRIVER=libGLESv2_viv.so)
   ```
5. If you are on `sh`, you need to set the environment with the unity player start. For example, `run SDL_VIDEO_EGL_DRIVER=libEGL_viv.so SDL_VIDEO_GL_DRIVER=libGLESv2_viv.so ./qnxplayer`.
6. Start the Unity Player.

## Application window configuration using environment variable

Use the environment variable `UNITY_QNX_WINDOW_CLASS_LIST` to configure the properties for your Unity application windows. This variable allows you to override Unity’s default windowing behavior by using the platform-specific window properties. You can define properties including display index, position, size, and hardware pipeline in the `graphics.conf` file for your application windows to use.

To set the `UNITY_QNX_WINDOW_CLASS_LIST` environment variable, consider the following points:

* Pass a list of `SCREEN_PROPERTY_CLASS IDs` separated by commas. For example, `main,fullscreen_overlay`. Each entry in the list corresponds to a [`SCREEN_PROPERTY_CLASS`](https://www.qnx.com/developers/docs/8.0/com.qnx.doc.screen/topic/manual/cscreen_config_intro.html#screen_winmgr_class) in your `graphics.conf` file.
* Do not include any white space.
* Specify up to 8 windows, as supported by Unity.
* Ensure that listed class IDs are defined in your `graphics.conf` file.

When you launch your application, each entry in the list applies to its corresponding window in order. For example, the first entry in the list applies to the first Unity window (main window) and the second entry applies to the second window.

**Note**: If `UNITY_QNX_WINDOW_CLASS_LIST` references class IDs that don’t exist in the `graphics.conf` file, the application will fail to start due to window creation failure.

### Example

`export UNITY_QNX_WINDOW_CLASS_LIST=main,fullscreen_overlay`

This assigns properties to windows in the following order:

* `main` window properties to the first window (the SDL window with ID 1)
* `fullscreen_overlay` window properties to the second window

## Additional resources

* [Apply your Screen configuration](https://www.qnx.com/developers/docs/8.0/com.qnx.doc.screen/topic/manual/cscreen_config_intro.html#ariaid-title10)