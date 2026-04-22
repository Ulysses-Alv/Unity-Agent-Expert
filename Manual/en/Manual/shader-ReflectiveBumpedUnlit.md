# Reflective Normal Mapped Unlit

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

## Reflective Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader will simulate reflective surfaces such as cars, metal objects etc. It requires an environment **Cubemap** which will define what exactly is reflected. The main texture’s alpha channel defines the strength of reflection on the object’s surface. Any **scene** lights will add illumination on top of what is reflected.

## Normal mapped Properties

This shader does not use normal-mapping in the traditional way. The **normal map** does not affect any lights shining on the object, because the shader does not use lights at all. The normal map will only distort the reflection map.

## Special Properties

This shader is special because it does not respond to lights at all, so you don’t have to worry about performance reduction from use of multiple lights. It simply displays the reflection cube map on the model. The reflection is distorted by the normal map, so you get the benefit of detailed reflection. Because it does not respond to lights, it is quite cheap. It is somewhat of a specialized use case, but in those cases it does exactly what you want as cheaply as possible.

## Performance

Generally, this shader is quite cheap to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

ReflectiveBumpedUnlit