# Particle lights and trails

Understand how the **Particle** System can render lights and trails on particle effects.

## Lights

The [Lights](PartSysLightsModule.md) module is a fast way to add real-time lighting to your particle effects. It can be used to make systems cast light onto their surroundings, for example for a fire, fireworks or lightning. It also allows you to have the lights inherit a variety of properties from the particles they are attached to. This can make it more believable that the particle effect itself is emitting light. For example, this can be achieved by making the lights fade out with their particles and having them share the same colors.

This module makes it easy to create a lot of real-time lights very quickly, and real-time lights have a high performance cost, especially in **Forward Rendering** mode. If the lights also cast shadows, the performance cost is even higher. To help guard against an accidental tweak to the emission rate and thus causing thousands of real-time lights to be created, the Maximum Lights property should be used. Creating more lights than your target hardware is able to manage can cause slowdowns and unresponsiveness.

## Trails

The [Trails](PartSysTrailsModule.md) module adds trails to a percentage of your particles. This module shares a number of properties with the [Trail Renderer](class-TrailRenderer.md) component, but offers the ability to easily attach Trails to particles, and to inherit various properties from the particles. Trails can be useful for a variety of effects, such as bullets, smoke, and magic visuals.

![The Trails module in Particles mode](../uploads/Main/PartSysTrailsModule.png)

The Trails module in Particles mode

![The Trails module in Ribbon mode](../uploads/Main/PartSysTrailsModuleRibbon.png)

The Trails module in Ribbon mode

### Tips

* Use the [Renderer](PartSysRendererModule.md) module to specify the Trail Material.
* Unity samples colors from the Color Gradient at each vertex, and linearly interpolates colors between each vertex,. Add more vertices to your **Line Renderer** to get a closer approximation of a detailed Color Gradient.