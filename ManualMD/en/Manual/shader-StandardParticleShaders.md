# Standard Particle Shaders Material Inspector window reference for the Built-In Render Pipeline

The Standard **Particle** **Shaders** have the same set of properties as the [Standard Shader](shader-StandardShader.md) (or a subset of those properties, depending on the Shader). This page describes the properties and options that are additional to the Standard Shader properties. For information on the Standard Shader properties, see documentation on [Material parameters](StandardShaderMaterialParameters.md).

## Blending Options

All of the Standard Particle Shaders have **Blending Options** that enable you to blend particles with the objects surrounding them in different ways.

| Property: | Function: |
| --- | --- |
| **Rendering Mode** | The Standard Particle Shaders can have the following [Rendering Mode](StandardShaderMaterialParameterRenderingMode.md) options, which control how the final color from the shader is combined with the background color:  **Additive**: Adds the background and final particle color together. This is useful for glow effects, like those you might use for fire or magic spells.  **Subtractive**: Subtracts the final particle color from the background, which darkens the particles against the background. This is useful for foggy effects, like those you might use for steam, or thick black smoke.  **Modulate**: Multiplies the final particle color with the background color. This is useful for portals and light-rays. |
| **Color Mode** | Control how the albedo texture is combined with the particle color. The **Color Mode** options are:  **Multiply**: Multiplies the particle albedo color with the particle texture.   **Additive**: Preserves a hot spot, such as a white part of the particle texture, while adding the particle albedo color to the darker pixels of the texture.   **Subtractive**: Subtracts the particle albedo color from the particle texture.   **Overlay**: Gives more contrast to the original color and adds the particle albedo color to the gray values. This is similar to Additive, but preserves the original colors.  **Color**: Uses the alpha channel from the particle texture and the albedo color from the particle itself. This is useful for overwriting particles with the same color, while keeping their original “shape”.  **Difference**: Subtracts the particle albedo color from the texture, or the texture from the color, to get a positive value. This is useful for a range of effects where you want a more dynamic color change.   See image below table for a demonstration of this effect. |

![Color Modes allow for different ways of combining the particle color with the albedo texture](../uploads/Main/SPSSColorModes.png)

Color Modes allow for different ways of combining the particle color with the albedo texture

## Main Options

| Property | Function |
| --- | --- |
| **Flip-Book Mode** | Render flip-books as individual frames or blend the frames together to give smoother animations. Set to either: **Simple** - Render frames in a flip-book as a sequence of individual frames. **Blended** - Blend the frames in a flip-book to render the flip-book as a smooth animation. |
| **Two Sided** | Render both the front and back faces of the particle. When disabled, Unity only renders the front face of the geometry, which is the face in the **camera**’s view. |
| **Enable Soft Particles** | Fade out particles when they get close to the surface of objects written into the depth buffer. This is useful for avoiding hard edges when particles intersect with opaque geometry. For example, by enabling soft particles, you can make the particle system emit particles close to an opaque surface without causing harsh intersections with the surface:  Left: Soft Particles enabled. Right: Soft Particles disabled. |
| **Enable Camera Fading** | Fade out particles when they get close to the camera. Set to: **Near fade** - The distance from the camera where particles are completely faded away. **Far fade** - The distance from the camera where particles start to fade as they get nearer. |
| **Enable Distortion** | Make particles perform fake refraction with the objects drawn before them. Distortion is ideal for creating heat haze effects for fire, for example:  Left: Fire with distortion. Right: Fire with distortion. This effect can be quite expensive because it captures the current frame to a texture. |

StandardParticleShaders