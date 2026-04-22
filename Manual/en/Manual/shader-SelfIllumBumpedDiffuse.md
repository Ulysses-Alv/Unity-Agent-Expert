# Self-Illuminated Normal mapped Diffuse

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

![Self-Illuminated Normal mapped Diffuse shader.](../uploads/Shaders/Shader-IllumBump.jpg)

Self-Illuminated Normal mapped Diffuse shader.

## Self-Illuminated Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader allows you to define bright and dark parts of the object. The alpha channel of a secondary texture will define areas of the object that “emit” light by themselves, even when no light is shining on it. In the alpha channel, black is zero light, and white is full light emitted by the object. Any **scene** lights will add illumination on top of the shader’s illumination. So even if your object does not emit any light by itself, it will still be lit by lights in your scene.

## Normal Mapped Properties

Like a **Diffuse** shader, this computes a simple (Lambertian) lighting model. The lighting on the surface decreases as the angle between it and the light decreases. The lighting depends only on the angle, and does not change as the **camera** moves or rotates around.

**Normal mapping** simulates small surface details using a texture, instead of spending more polygons to actually carve out details. It does not actually change the shape of the object, but uses a special texture called a **Normal Map** to achieve this effect. In the normal map, each **pixel**’s color value represents the angle of the surface normal. Then by using this value instead of the one from geometry, lighting is computed. The normal map effectively overrides the **mesh**’s geometry when calculating lighting of the object.

### Creating Normal maps

You can import normal maps created outside of Unity, or you can import a regular grayscale image and convert it to a Normal Map from within Unity. (This page refers to a legacy shader which has been superseded by the [Standard Shader](shader-StandardShader.md), but you can learn more about how to use [Normal Maps in the Standard Shader](StandardShaderMaterialParameterNormalMap.md))

### Technical Details

The Normal Map is a tangent space type of normal map. Tangent space is the space that “follows the surface” of the model geometry. In this space, Z always points away from the surface. Tangent space Normal Maps are a bit more expensive than the other “object space” type Normal Maps, but have some advantages:

1. It’s possible to use them on deforming models - the bumps will remain on the deforming surface and will just work.
2. It’s possible to reuse parts of the normal map on different areas of a model; or use them on different models.

## Diffuse Properties

Diffuse computes a simple (Lambertian) lighting model. The lighting on the surface decreases as the angle between it and the light decreases. The lighting depends only on this angle, and does not change as the camera moves or rotates around.

## Performance

Generally, this shader is cheap to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

SelfIllumBumpedDiffuse