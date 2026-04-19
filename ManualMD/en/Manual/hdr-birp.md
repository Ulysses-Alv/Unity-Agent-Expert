# HDR in the Built-in Render Pipeline

Support for **HDR** in the Built-In **Render Pipeline** depends on the **rendering path** in use.

## Forward Rendering

In the **Forward rendering** path, a **Camera** only supports HDR if you have a **post-processing** effect on the Camera. This is due to performance considerations. If you have no post-processing effect on the Camera, the **scene** will be rendered directly to the back buffer where HDR is not supported.

## Deferred Shading

In the **Deferred Shading** rendering path, when HDR is enabled for a Camera the lighting buffer is also allocated as a floating point buffer. This reduces banding in the lighting buffer.