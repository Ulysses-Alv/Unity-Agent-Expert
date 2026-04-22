# Apply GPU instancing for a Particle System in the Built-In Render Pipeline

GPU instancing offers a large performance boost compared with CPU rendering. You can use it if you want your **particle** system to render **Mesh** particles (as opposed to the default [rendering mode](PartSysRendererModule.md) of rendering **billboard** particles).

To be able to use GPU instancing with your **particle systems**:

* Set your Particle System’s [renderer](PartSysRendererModule.md) mode to **Mesh**
* Use a **shader** for the [renderer](PartSysRendererModule.md) material that supports GPU Instancing
* Run your project on a platform that supports GPU instancing

To enable GPU instancing for a particle system, you must enable the **Enable GPU Instancing** checkbox in the **Renderer** module of your particle system.

![The option to enable Particle System GPU instancing in the Renderer module](../uploads/Main/PartSysInstancingEnable.png)

The option to enable Particle System GPU instancing in the Renderer module

Unity comes with a built-in particle shader that supports GPU instancing, but the default particle material does not use it, so you must change this to use GPU instancing. The particle shader that supports GPU instancing is called **Particles/Standard Surface**. To use it, you must create your own new **material**, and set the material’s shader to **Particles/Standard Surface**. You must then assign this new material to the material field in the Particle System [renderer](PartSysRendererModule.md) module.

![The built-in shader that is compatible with Particle System GPU Instancing](../uploads/Main/PartSysInstancingShader.png)

The built-in shader that is compatible with Particle System GPU Instancing

If you are using a different shader for your particles, it must use ‘#pragma target 4.5’ or higher. See [Shader Compile Targets](SL-ShaderCompileTargets.md) for more details. This requirement is higher than regular GPU Instancing in Unity because the Particle System writes all its instance data to a single large buffer, rather than breaking up the instancing into multiple draw calls.

## Custom shader examples

You can also write custom shaders that make use of GPU Instancing. See the following sections for more information:

* [Particle system GPU Instancing in a Surface Shader](example-particle-system-gpu-instancing-surface-shader.md)
* [Particle system GPU Instancing in a Custom Shader](example-particle-system-gpu-instancing-custom-shader.md)
* [Customising instance data used by the Particle System](example-particle-system-gpu-instancing-custom-vertex-streams.md) (to work alongside Custom Vertex Streams)