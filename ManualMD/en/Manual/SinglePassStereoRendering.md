# Stereo rendering

VR and most **MR** devices require rendering the Unity **scene** in stereo. Unity **XR** supports two stereo render modes:

* **Multi-pass**: in this mode, Unity performs a render pass for each eye. Some parts of the render loop are shared between the two passes, so multi-pass rendering is faster than rendering the scene with two unique **cameras**. Multi-pass mode provides the widest compatibility with existing **shaders** and rendering utilities, but is slower than single pass instanced mode.
* **Single-pass instanced**: in this mode, Unity renders the scene in a single pass using instanced draw calls. This mode greatly decreases CPU usage and slightly decreases GPU usage compared to the multi-pass mode.
* **Multiview**: A variation of single-pass instanced rendering supported by some OpenGL and OpenGL ES devices. This option replaces single-pass instanced when available.

**Note:** The earlier technique of rendering the scene into a double-wide texture using a single render pass is no longer available.

Single-pass instanced stereo rendering is now available on most **VR** platforms.

## Set the render mode

You can find the **Render mode** setting under **XR **Plug-in** Management** in ****Project Settings****. Each XR provider plug-in provides its own setting, if supported.

To set a render mode:

1. Open **Project Settings** (menu: **Edit > Project Settings**).
2. Expand the **XR Plugin Management** section, if necessary.
3. Select the settings page for the relevant provider plug-in.
4. Choose a mode from the list.

![Render mode options in the MockHMD provider plug-in](../uploads/Main/SinglePassStereoRendering3.png)

Render mode options in the MockHMD provider plug-in

**Note:** Some plug-ins name the setting **Stereo Rendering Mode**.

## Single-pass instanced render mode support

Single-pass instanced render mode is supported on the following platforms and devices:

* Android devices that support the Multiview extension
* HoloLens
* PlayStation VR
* PC devices (tethered):
* For DirectX on desktop, the GPU must support Direct3D 11 and the `VPAndRTArrayIndexFromAnyShaderFeedingRasterizer` extension.
* For OpenGL on desktop, the GPU must support one of the following extensions:
  + `GL_NV_viewport_array2`
  + `GL_AMD_vertex_shader_layer`
  + `GL_ARB_shader_viewport_layer_array`

If you set the **Render Mode** to **Single Pass Instanced** when that mode is not supported, then rendering falls back to multi-pass mode.

**Notes:**

* Unity doesn’t support single-pass instanced rendering in the built-in **render pipeline** when using Shader Graph.
* Unity doesn’t support single-pass instanced rendering in the built-in render pipeline when using deferred rendering.