# Sample output textures in a shader

Depth textures are available for sampling in shaders as global **shader** properties. By declaring a sampler called `_CameraDepthTexture` you will be able to sample the main depth texture for the **camera**.

`_CameraDepthTexture` always refers to the camera’s primary depth
texture. By contrast, you can use `_LastCameraDepthTexture` to refer to the last depth texture rendered by any camera. This could be useful for example if you render a half-resolution depth texture in script using a secondary camera and want to make it available to a post-process shader.

In the Built-In **Render Pipeline**, the motion vectors texture (when enabled) is available in Shaders as a global Shader property. By declaring a sampler called ‘\_CameraMotionVectorsTexture’ you can sample the Texture for the currently rendering Camera. When sampling from this texture motion from the encoded **pixel** is returned in a range of –1..1. This will be the UV offset from the last frame to the current frame.

## Additional resources

* [Sample motion vectors in a shader in URP](urp/features/motion-vectors-sample.md)