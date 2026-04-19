# Transparent Cutout Vertex-Lit

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

![Transparent Cutout Vertex-Lit shader.](../uploads/Shaders/Shader-TransCutoutVertex.jpg)

Transparent Cutout Vertex-Lit shader.

## Transparent Cutout Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

Cutout shader is an alternative way of displaying transparent objects. Differences between Cutout and regular [Transparent](shader-TransparentFamily.md) shaders are:

* This shader cannot have partially transparent areas. Everything will be either fully opaque or fully transparent.
* Objects using this shader can cast and receive shadows!
* The graphical sorting problems normally associated with Transparent shaders do not occur when using this shader.

This shader uses an alpha channel contained in the **Base** Texture to determine the transparent areas. If the alpha contains a blend between transparent and opaque areas, you can manually determine the cutoff point for the which areas will be shown. You change this cutoff by adjusting the **Alpha Cutoff** slider.

## Vertex-Lit Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader is **Vertex-Lit**, which is one of the simplest shaders. All lights shining on it are rendered in a single pass and calculated at vertices only.

Because it is vertex-lit, it won’t display any pixel-based rendering effects, such as light cookies, normal mapping, or shadows. This shader is also much more sensitive to tesselation of the models. If you put a point light very close to a cube using this shader, the light will only be calculated at the corners. Pixel-lit shaders are much more effective at creating a nice round highlight, independent of tesselation. If that’s an effect you want, you may consider using a pixel-lit shader or increase tesselation of the objects instead.

## Performance

Generally, this shader is very cheap to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

TransCutVertexLit