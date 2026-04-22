# Diffuse

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

![Normal Diffuse shader.](../uploads/Shaders/Shader-NormalDiffuse.jpg)

Normal Diffuse shader.

## Diffuse Properties

Diffuse computes a simple (Lambertian) lighting model. The lighting on the surface decreases as the angle between it and the light decreases. The lighting depends only on this angle, and does not change as the **camera** moves or rotates around.

## Performance

Generally, this shader is cheap to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

NormalDiffuse