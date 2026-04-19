# Change particle color

Understand how the **Particle** System can change a particle’s color based on its speed or lifetime.

## Color by speed

The [Color by Speed](PartSysColorBySpeedModule.md) module can change a particle’s color based on its speed in distance units per second.

Burning or glowing particles (such as sparks) tend to burn more brightly when they move quickly through the air (for example, when sparks are exposed to more oxygen), but then dim slightly as they slow down. To simulate this, you might use *Color By Speed* with a gradient that has white at the upper end of the speed range, and red at the lower end (in the spark example, faster particles will appear white while slower particles are red).

## Color over lifetime

The [Color Over Lifetime](PartSysColorOverLifeModule.md) module can change a particle’s color based on how long it has existed.

Many types of natural and fantastical particles vary in color over time, and so this property has many uses. For example, white hot sparks will cool as they pass through the air and a magic spell might burst into a rainbow of colors. Equally important, though, is the variation of alpha (transparency). It is very common for particles to burn out, fade or dissipate as they reach the end of their lifetime (for example, hot sparks, fireworks and smoke particles) and a simple diminishing gradient produces this effect.

When also using the [Start Color](PartSysMainModule.md) property, this module multiples the 2 colors together, to get the final particle color.

## Fade particles near opaque GameObjects

To fade particles near opaque **GameObjects**, follow these steps:

1. Enable a [deferred rendering path](rendering-paths-introduction.md#deferred), or [output a depth texture from the camera](SL-CameraDepthTexture.md).
2. Go to **Edit** > **Project Settings** > **Quality**, then enable **Soft Particles**.

Soft particles are supported only on platforms that support [depth textures](SL-CameraDepthTexture.md).

![Soft particles disabled, with visible intersections with the scene.](../uploads/Main/SoftParticlesOff.jpg)

Soft particles disabled, with visible intersections with the scene.

![Soft particles enabled, with intersections that fade out smoothly.](../uploads/Main/SoftParticlesOn.jpg)

Soft particles enabled, with intersections that fade out smoothly.

## Additional resources

* [Quality settings tab reference](class-qualitysettings.md)