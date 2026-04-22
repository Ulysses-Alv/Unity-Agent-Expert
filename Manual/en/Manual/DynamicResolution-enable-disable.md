# Enable or disable Dynamic Resolution for render targets

With **dynamic resolution**, render targets have the [DynamicallyScalable](../ScriptReference/RenderTextureCreationFlags.DynamicallyScalable.md) flag. You can set this to state whether Unity should scale this **render texture** as part of the dynamic resolution process or not. **Cameras** also have the [allowDynamicResolution](../ScriptReference/Camera-allowDynamicResolution.md) flag, which you can use to set up dynamic resolution so that there is no need to override the render target if you just want to apply dynamic resolution to a less complex **Scene**.

## MRT buffers

When you enable **Allow Dynamic Resolution** on the Camera, Unity scales all of that Camera’s targets.