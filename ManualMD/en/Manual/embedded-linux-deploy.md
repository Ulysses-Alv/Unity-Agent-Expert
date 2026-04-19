# Deploy an Embedded Linux project

On the Embedded Linux player, Unity uses SDL2 to handle keyboard, mouse, and controller input to interact with the Player window. Depending on the graphics API used, it requires SDL to dynamically load `libEGL` and `libGLESv2` for OpenGL ES, or `libvulkan` for Vulkan from the user space.

## Setup for [Wayland](https://wiki.archlinux.org/title/Wayland)

Although this setup assumes that you’re using [Weston](https://gitlab.freedesktop.org/wayland/weston) (the reference Wayland server), you can use the same setup with slight modifications for another compositor.

## Prerequisites

This assumes that you have a Wayland compositor (Weston) running, which exports the Wayland socket in the directory that the environment variable `XDG_RUNTIME_DIR` is linked to.

### Setup on Desktop shell

To deploy your project on desktop shell:

1. Verify that the environment variable `XDG_RUNTIME_DIR` is set to the correct directory. If not, then run `export XDG_RUNTIME_DIR=<dir>` with the correct directory (`/run/user/1000/`is the default for a Weston installation).
2. Run Unity Player.

### Setup on IVI-shell extension

You can deploy your project using IVI-shell extension, which is an alternative shell extension for Weston.

1. Set up IVI Surface IDs that the Unity Player needs to use with the environment variable `UNITY_IVI_SURFACE_IDS`. If this isn’t set, the Unity Player uses ID `4711` and upwards for the newly created surfaces (for example, Unity Display 1 will use 4711, Unity Display 2 will use 4712 etc.).
2. **Note**: The environment variable expects a comma-separated list of IDs. For example, `export UNITY_IVI_SURFACE_IDS=100,200,300 uses ID` `100` for Unity Display 1, `200` for Unity Display 2, etc..
3. If you’re only using a single display output, a single ID is sufficient. For example, export `UNITY_IVI_SURFACE_IDS=100`.
4. Verify that the environment variable `XDG_RUNTIME_DIR`is set to the correct directory. If this isn’t set, run `export XDG_RUNTIME_DIR=<dir>` with the correct directory (`/run/user/1000/` is the default for a Weston installation).
5. Run the Unity Player.
6. Use the following steps to set up an IVI surface for Unity (example, `weston` to `fullscreen map` surface).

   1. Create IVI Layer.

      `LayerManagerControl create layer 0 <display-width> <display-height>`
   2. Add Layer on Screen.

      `LayerManagerControl set screen 0 render order 0`
   3. Add Unity Player Surface on Layer.

      `LayerManagerControl set layer 0 render order <surface-id>`
   4. Add Surface Source Region.

      `LayerManagerControl set surface <surface-id> source region 0 0 <display-width> <display-height>`
   5. Add Surface Destination Region.

      `LayerManagerControl set surface <surface-id> destination region 0 0 <display-width> <display-height>`
   6. Add Layer Visibility.

      `LayerManagerControl set layer 0 visibility 1`
   7. Add Surface Visibility.

      `LayerManagerControl set surface <surface-id> visibility 1`

Unity Player appears on screen.

### Additional information

By default, Unity creates surfaces of the same size as the physical displays. If you want to use surfaces other than physical displays, such as rendering multiple surfaces to one screen, use `UNITY_IVI_EXPORT_DISPLAYS` as the environment variable.

For example, with the setting `export` `UNITY_IVI_EXPORT_DISPLAYS=1024x768@60,1920x1080@60` Unity uses a surface size of *1024x768* for Unity Display 1, and a surface size of *1920x1080* for Unity Display 2.

You can omit `@60` and use `export UNITY_IVI_EXPORT_DISPLAYS=1024x768,1920x1080`because `@60` is automatically assumed.

## Additional resources

* [Build and deliver for Embedded Linux](embedded-linux-build-and-deliver.md)