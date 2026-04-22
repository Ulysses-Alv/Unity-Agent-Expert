# Create and view a Particle System

The Built-in Particle System uses a component, so placing a Particle System in a Scene is a matter of adding a pre-made GameObject (menu: **GameObject** > **Effects** > **Particle System**) or adding the component to an existing GameObject (menu: **Component** > **Effects** > **Particle System**). Because the component is quite complicated, the Inspector is divided into a number of collapsible sub-sections or **modules** that each contain a group of related properties. Additionally, you can edit one or more systems at the same time using a separate Editor window accessed via the **Open Editor** button in the Inspector. See documentation on the [Particle System component](class-ParticleSystem.md) and individual [Particle System modules](ParticleSystemModules.md) to learn more.

When a GameObject with a **Particle** System is selected, the **Scene** view contains a small **Particle Effect** panel, with some simple controls that are useful for visualising changes you make to the system’s settings.

![The Particle Effect panel in the Scene view.](../uploads/Main/PartSysEffectPanel.png)

The Particle Effect panel in the Scene view.

The **Playback Speed** allows you to speed up or slow down the particle simulation, so you can quickly see how it looks at an advanced state. The **Playback Time** indicates the time elapsed since the system was started; this may be faster or slower than real time depending on the playback speed. The **Particle Count** indicates how many particles are currently in the system. The playback time can be moved backwards and forwards by clicking on the **Playback Time** label and dragging the mouse left and right. The buttons at the top of the panel can be used to pause and resume the simulation, or to stop it and reset to the initial state.