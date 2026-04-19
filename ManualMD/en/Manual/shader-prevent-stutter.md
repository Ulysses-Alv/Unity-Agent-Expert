# Prevent shader stutters

Learn how you can prevent **shader** stutters in your project.

A stutter refers to a noticeable pause in rendering, which causes motion to stop or become uneven. Stutters can occur when Unity uses a shader variant for the first time in your build application. Stutters cause a negative impact on users of your application. For example, stutters reduce the responsiveness on your game, and cam lead to severe motion sickness on **XR** platforms.

Refer to the following topics to understand the approaches you can use to prevent shaders stutters in your project:

| **Topic** | **Description** |
| --- | --- |
| [Tracing and warming up PSOs](shader-pso-trace-warming.md) | Trace and warm up graphics states with the experimental `GraphicsStateCollection` API to reduce shader compilation stutters. |
| [Other methods to warm up shaders](shader-prewarm-other.md) | Warm up shaders if your build targets the DirectX 11, OpenGL, OpenGLES, or **WebGL** graphics APIs. |