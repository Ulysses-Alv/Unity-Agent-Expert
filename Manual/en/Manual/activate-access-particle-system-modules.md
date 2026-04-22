# Activate and access Particle System modules

The **Particle** System component has many properties, and for convenience, the **Inspector** organises them into collapsible sections called “modules”. These modules are documented in separate pages. See documentation on [Particle System Modules](ParticleSystemModules.md) to learn about each one.

To expand and collapse modules, click the bar that shows their name. Use the checkbox on the left to enable or disable the functionality of the properties in that module. For example, if you don’t want to vary the sizes of particles over their lifetime, uncheck the **Size over Lifetime** module.

The **Open Editor** button displays the options in a separate Editor window, which allows you to edit multiple systems at once.

The [Particle Effect panel](PartSysUsage.md) in the **Scene** View contains some additional options for previewing Particle Systems.

All Particle System modules are part of the [Particle System](class-ParticleSystem.md) component.

To create a new Particle System and enable a module:

1. Click **GameObject** > **Effects** > **Particle System**.
2. In the Inspector, find the Particle System component.
3. In the Particle System component, find the fold-out for the module you want to apply.
4. To the left of the fold-out header, enable the checkbox.

# Access modules via the API

Modules are part of the Particle System component, so you can access it via the [ParticleSystem](../ScriptReference/ParticleSystem.md) class.