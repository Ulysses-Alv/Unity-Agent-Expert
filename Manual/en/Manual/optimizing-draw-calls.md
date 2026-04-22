# Introduction to optimizing draw calls

Optimizing draw calls speeds up rendering, by reducing how often the CPU sends information to the GPU.

By default, each **mesh** needs its own draw call. A draw call has the following steps:

* Update the render state. The CPU uses the graphics API to prepare and send everything the GPU needs to draw the mesh. For example **shader** code, textures, and buffers. This is the most CPU-intensive step.
* Submit a draw call. The CPU uses the graphics API to tell the GPU what to draw.

If you optimize draw calls, Unity groups data and meshes that use the same render state. As a result, Unity updates the render state less often, and submits multiple objects in single draw calls.

Optimizing draw calls provides the following benefits:

* Improve frame times, or increase the number of **GameObjects** you can render per frame.
* Reduce the amount of electricity your application needs. For battery-powered devices, this reduces the heat the device produces and the rate at which batteries run out.

## Types of draw call optimization

Draw call optimization methods in Unity work in one of the following ways:

* Drawing multiple copies of a mesh in one draw call, using the hardware capabilities of the GPU. This is GPU instancing.
* Using another technique to reduce render state updates, such as large GPU buffers. This is how the Scriptable **Render Pipeline** (SRP) Batcher works.
* Combining meshes that use the same render state and drawing them together. This is batching.

For more information, refer to [Choose a method for optimizing draw calls](optimizing-draw-calls-choose-method.md).

## Additional resources

* [Optimize shaders](SL-ShaderPerformance.md)
* [Optimize for mobile](iphone-iOS-Optimization.md)
* [Optimize for untethered XR devices](xr-untethered-device-optimization.md)