# Reflective Vertex-Lit

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

![Reflective Vertex-Lit shader.](../uploads/Shaders/Shader-ReflVertex.jpg)

Reflective Vertex-Lit shader.

## Reflective Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader will simulate reflective surfaces such as cars, metal objects etc. It requires an environment **Cubemap** which will define what exactly is reflected. The main texture’s alpha channel defines the strength of reflection on the object’s surface. Any **scene** lights will add illumination on top of what is reflected.

## Vertex-Lit Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader is **Vertex-Lit**, which is one of the simplest shaders. All lights shining on it are rendered in a single pass and calculated at vertices only.

Because it is vertex-lit, it won’t display any pixel-based rendering effects, such as light cookies, normal mapping, or shadows. This shader is also much more sensitive to tesselation of the models. If you put a point light very close to a cube using this shader, the light will only be calculated at the corners. Pixel-lit shaders are much more effective at creating a nice round highlight, independent of tesselation. If that’s an effect you want, you may consider using a pixel-lit shader or increase tesselation of the objects instead.

## Performance

Generally, this shader is not too expensive to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

ReflectiveVertexLit