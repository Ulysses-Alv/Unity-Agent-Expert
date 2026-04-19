# Reflective Specular

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this **shader**.

One consideration for this shader is that the Base texture’s alpha channel will double as both the Reflection Map and the Specular Map.

## Reflective Properties

**Note.** Unity 5 introduced the [Standard Shader](shader-StandardShader.md) which replaces this shader.

This shader will simulate reflective surfaces such as cars, metal objects etc. It requires an environment **Cubemap** which will define what exactly is reflected. The main texture’s alpha channel defines the strength of reflection on the object’s surface. Any **scene** lights will add illumination on top of what is reflected.

## Specular Properties

Specular computes the same simple (Lambertian) lighting as Diffuse, plus a viewer dependent specular highlight. This is called the Blinn-Phong lighting model. It has a specular highlight that is dependent on surface angle, light angle, and viewing angle. The highlight is actually just a realtime-suitable way to simulate blurred reflection of the light source. The level of blur for the highlight is controlled with the **Shininess** slider in the **Inspector**.

Additionally, the alpha channel of the main texture acts as a Specular Map (sometimes called “gloss map”), defining which areas of the object are more reflective than others. Black areas of the alpha will be zero specular reflection, while white areas will be full specular reflection. This is very useful when you want different areas of your object to reflect different levels of specularity. For example, something like rusty metal would use low specularity, while polished metal would use high specularity. Lipstick has higher specularity than skin, and skin has higher specularity than cotton clothes. A well-made Specular Map can make a huge difference in impressing the player.

## Performance

Generally, this shader is moderately expensive to render. For more details, please view the [Shader Peformance page](shader-Performance.md).

ReflectiveSpecular