# Reflective Parallax Diffuse

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

## Reflective Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader will simulate reflective surfaces such as cars, metal objects etc. It requires an environment **Cubemap** which will define what exactly is reflected. The main texture’s alpha channel defines the strength of reflection on the object’s surface. Any **scene** lights will add illumination on top of what is reflected.

## Parallax Normal mapped Properties

**Parallax Normal mapped** is the same as regular **Normal mapped**, but with a better simulation of “depth”. The extra depth effect is achieved through the use of a **Height Map**. The Height Map is contained in the alpha channel of the **Normal map**. In the alpha, black is zero depth and white is full depth. This is most often used in bricks/stones to better display the cracks between them.

The Parallax mapping technique is pretty simple, so it can have artifacts and unusual effects. Specifically, very steep height transitions in the Height Map should be avoided. Adjusting the **Height** value in the **Inspector** can also cause the object to become distorted in an odd, unrealistic way. For this reason, it is recommended that you use gradual Height Map transitions or keep the **Height** slider toward the shallow end.

## Diffuse Properties

Diffuse computes a simple (Lambertian) lighting model. The lighting on the surface decreases as the angle between it and the light decreases. The lighting depends only on this angle, and does not change as the **camera** moves or rotates around.

## Performance

Generally, this shader is on the more expensive rendering side. For more details, please view the [Shader Peformance page](shader-Performance.md).

ReflectiveParallaxDiffuse