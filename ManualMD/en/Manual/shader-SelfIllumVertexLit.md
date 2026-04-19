# Self-Illuminated Vertex-Lit

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

![Self-Illuminated Vertex-Lit shader.](../uploads/Shaders/Shader-IllumVertex.jpg)

Self-Illuminated Vertex-Lit shader.

## Self-Illuminated Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader allows you to define bright and dark parts of the object. The alpha channel of a secondary texture will define areas of the object that “emit” light by themselves, even when no light is shining on it. In the alpha channel, black is zero light, and white is full light emitted by the object. Any **scene** lights will add illumination on top of the shader’s illumination. So even if your object does not emit any light by itself, it will still be lit by lights in your scene.

## Vertex-Lit Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader is **Vertex-Lit**, which is one of the simplest shaders. All lights shining on it are rendered in a single pass and calculated at vertices only.

Because it is vertex-lit, it won’t display any pixel-based rendering effects, such as light cookies, normal mapping, or shadows. This shader is also much more sensitive to tesselation of the models. If you put a point light very close to a cube using this shader, the light will only be calculated at the corners. Pixel-lit shaders are much more effective at creating a nice round highlight, independent of tesselation. If that’s an effect you want, you may consider using a pixel-lit shader or increase tesselation of the objects instead.

## Performance

Generally, this shader is cheap to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

SelfIllumVertexLit