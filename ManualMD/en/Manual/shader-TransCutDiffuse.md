# Transparent Cutout Diffuse

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

![Transparent Cutout Diffuse shader.](../uploads/Shaders/Shader-TransCutoutDiffuse.jpg)

Transparent Cutout Diffuse shader.

## Transparent Cutout Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

Cutout shader is an alternative way of displaying transparent objects. Differences between Cutout and regular [Transparent](shader-TransparentFamily.md) shaders are:

* This shader cannot have partially transparent areas. Everything will be either fully opaque or fully transparent.
* Objects using this shader can cast and receive shadows!
* The graphical sorting problems normally associated with Transparent shaders do not occur when using this shader.

This shader uses an alpha channel contained in the **Base** Texture to determine the transparent areas. If the alpha contains a blend between transparent and opaque areas, you can manually determine the cutoff point for the which areas will be shown. You change this cutoff by adjusting the **Alpha Cutoff** slider.

## Diffuse Properties

Diffuse computes a simple (Lambertian) lighting model. The lighting on the surface decreases as the angle between it and the light decreases. The lighting depends only on this angle, and does not change as the **camera** moves or rotates around.

## Performance

Generally, this shader is cheap to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

TransCutDiffuse