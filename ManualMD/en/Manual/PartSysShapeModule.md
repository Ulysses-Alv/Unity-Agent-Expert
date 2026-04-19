# Shape module reference

Explore the Shape module properties to define the volume or surface of [particle](class-ParticleSystem.md) emission, and the direction of the start velocity.

## Shapes in the Shape module

The direction of particle emission varies depending on the shape you choose.

| **Shape** | **Description** |
| --- | --- |
| **Sphere** | Emits particles uniformly in all directions. |
| **Hemisphere** | Emits particles uniformly in all directions on one side of a plane. |
| **Cone** | Emits particles from the base or body of a cone. The particles diverge in proportion to their distance from the cone’s center line. |
| **Donut** | Emits particles from a torus. The particles move outwards from the ring of the torus. |
| **Box** | Emits particles from the edge, surface, or body of a box shape. The particles move in the emitter object’s forward (Z) direction. |
| ****Mesh**** | Emits particles from any arbitrary mesh shape supplied via the **Inspector**. |
| ****Mesh Renderer**** | Emits particles from a reference to a **GameObject**’s Mesh Renderer. |
| **Skinned Mesh Renderer** | Emits particles from a reference to a GameObject’s Skinned Mesh Renderer. |
| ****Sprite**** | Emits particles from a sprite shape supplied via the Inspector. |
| ****Sprite Renderer**** | Emits particles from a reference to a GameObject’s Sprite Renderer. |
| **Circle** | Emits particles uniformly from the center or edge of a circle. The particles move only in the plane of the circle. |
| **Edge** | Emits particles from a line segment. The particles move in the emitter object’s upward (Y) direction. |
| **Rectangle** | Emits particles from a rectangle. The particles move up from the rectangle. |

## Shape properties

This section details the properties for each **Shape**.

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

### Sphere, Hemisphere, Circle

**Sphere**, **Hemisphere**, and **Circle** have the same properties.

| **Property** | **Description** |
| --- | --- |
| **Radius** | Sets the radius of the circular aspect of the shape. |
| **Radius Thickness** | Sets the proportion of the volume that emits particles. A value of 0 emits particles from the outer surface of the shape. A value of 1 emits particles from the entire volume. Values between 0 and 1 will use a proportion of the volume. |
| **Arc** | Sets the angular portion of a full circle that forms the emitter’s shape. |
| **Mode** | Defines how Unity generates particles around the arc of the shape. For information on the **Mode** dropdown options, refer to [Mode dropdown](#ModeDropdown). |
| **Spread** | Controls the discrete intervals around the arc where particles can be generated. For example, a value of 0 means that particles can spawn anywhere around the arc, and a value of 0.1 means that particles can only spawn at 10% intervals around the shape. |
| **Speed** | Controls the speed the emission position moves around the arc. Set this value to **Constant** for the value to always remain the same, or **Curve** for the value to change over time. This property is available only when **Mode** is **Loop** or **Ping-Pong**. |
| **Texture** | Selects a texture to use for tinting and discarding particles. |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a **collision**. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center. When **Shape** is **Hemisphere** or **Circle**, a value of 1 behaves identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Cone

| **Property** | **Description** |
| --- | --- |
| **Angle** | Sets the angle of the cone at its point. A value of 0 produces a cylinder while a value of 90 produces a flat disc. |
| **Radius** | Sets the radius of the circular aspect of the shape. |
| **Radius Thickness** | Sets the proportion of the volume that emits particles. A value of 0 emits particles from the outer surface of the shape. A value of 1 emits particles from the entire volume. Values between 0 and 1 will use a proportion of the volume. |
| **Arc** | Sets the angular portion of a full circle that forms the emitter’s shape. |
| **Mode** | Defines how Unity generates particles around the arc of the shape. For information on the **Mode** dropdown options, refer to [Mode dropdown](#ModeDropdown). |
| **Spread** | Controls the discrete intervals around the arc where particles can be generated. For example, a value of 0 means that particles can spawn anywhere around the arc, and a value of 0.1 means that particles can only spawn at 10% intervals around the shape. |
| **Speed** | Controls the speed the emission position moves around the arc. Set this value to **Constant** for the value to always remain the same, or **Curve** for the value to change over time. This property is available only when **Mode** is **Loop** or **Ping-Pong**. |
| **Length** | Sets the length of the cone. This property is available only when **Emit from:** is **Volume**. |
| **Emit from:** | Chooses the part of the cone to emit particles from: **Base** or **Volume**. |
| **Texture** | Selects a texture to use for tinting and discarding particles. |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a collision. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center, behaving identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Donut

| **Property** | **Description** |
| --- | --- |
| **Radius** | Sets the radius of the main donut ring. |
| **Donus Radius** | Sets the radius of the outer donut ring. |
| **Radius Thickness** | Sets the proportion of the volume that emits particles. A value of 0 emits particles from the outer surface of the shape. A value of 1 emits particles from the entire volume. Values between 0 and 1 will use a proportion of the volume. |
| **Arc** | Sets the angular portion of a full circle that forms the emitter’s shape. |
| **Mode** | Defines how Unity generates particles around the arc of the shape. For information on the **Mode** dropdown options, refer to [Mode dropdown](#ModeDropdown). |
| **Spread** | Controls the discrete intervals around the arc where particles can be generated. For example, a value of 0 means that particles can spawn anywhere around the arc, and a value of 0.1 means that particles can only spawn at 10% intervals around the shape. |
| **Speed** | Controls the speed the emission position moves around the arc. Set this value to **Constant** for the value to always remain the same, or **Curve** for the value to change over time. This property is available only when **Mode** is **Loop** or **Ping-Pong**. |
| **Texture** | Selects a texture to use for tinting and discarding particles. |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a collision. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center, behaving identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Box

| **Property** | **Description** |
| --- | --- |
| **Emit from:** | Chooses the part of the box to emit particles from: **Volume**, **Edge**, or **Shell**. The default value is **Volume**. |
| **Texture** | Selects a texture to use for tinting and discarding particles. |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a collision. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center, behaving identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Mesh, Mesh Renderer, Skinned Mesh Renderer

**Mesh**, **Mesh Renderer**, and **Skinned Mesh Renderer** have the same properties.

| **Property** | **Description** |
| --- | --- |
| **Type** | Choose where to emit particles from: **Vertex**, **Edge**, or **Triangle**. The default value is **Vertex**. |
| **Mode** | Defines how Unity generates particles around the arc of the shape. For information on the **Mode** dropdown options, refer to [Mode dropdown](#ModeDropdown). |
| **Spread** | Controls the discrete intervals around the arc where particles can be generated. For example, a value of 0 means that particles can spawn anywhere around the arc, and a value of 0.1 means that particles can only spawn at 10% intervals around the shape. |
| **Speed** | Controls the speed the emission position moves around the arc. Set this value to **Constant** for the value to always remain the same, or **Curve** for the value to change over time. This property is available only when **Mode** is **Loop** or **Ping-Pong**. |
| **Mesh** | Selects the mesh that provides the emitter’s shape. |
| **Single Material** | Selects whether to emit particles from a particular submesh, identified by the material index number. If enabled, a numeric field appears for you to specify the material index number. |
| **Use Mesh Colors** | Modulates particle color with mesh vertex colors. If mesh vertex colors don’t exist, this uses the **shader** color property `_Color` or `_TintColor` from the material. |
| **Normal Offset** | Sets the distance away from the surface of the mesh to emit particles in the direction of the surface normal. |
| **Texture** | Selects a texture to use for tinting and discarding particles. |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a collision. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center, behaving identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Sprite and Sprite Renderer

**Sprite** and **Sprite Renderer** have the same properties.

| **Property** | **Description** |
| --- | --- |
| **Type** | Choose where to emit particles from: **Vertex**, **Edge**, or **Triangle**. The default value is **Vertex**. |
| **Sprite** | Selects the sprite that defines the particle emitter’s shape. |
| **Normal Offset** | Sets the distance away from the surface of the sprite to emit particles in the direction of the surface normal. |
| **Texture** | Selects a texture to use for tinting and discarding particles. |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a collision. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center, behaving identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Edge

| **Property** | **Description** |
| --- | --- |
| **Radius** | Defines the length of the edge. |
| **Mode** | Defines how Unity generates particles around the arc of the shape. For information on the **Mode** dropdown options, refer to [Mode dropdown](#ModeDropdown). |
| **Spread** | Controls the discrete intervals around the arc where particles can be generated. For example, a value of 0 means that particles can spawn anywhere around the arc, and a value of 0.1 means that particles can only spawn at 10% intervals around the shape. |
| **Speed** | Controls the speed the emission position moves around the arc. Set this value to **Constant** for the value to always remain the same, or **Curve** for the value to change over time. This property is available only when **Mode** is **Loop** or **Ping-Pong**. |
| **Texture** | Selects a texture to use for tinting and discarding particles. |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a collision. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center, behaving identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Rectangle

| **Property** | **Description** |
| --- | --- |
| **Position** | Applies an offset to the emitter shape used for spawning particles. |
| **Rotation** | Rotates the emitter shape used for spawning particles. |
| **Scale** | Changes the size of the emitter shape used for spawning particles. |
| **Align to Direction** | Orients particles based on their initial direction of travel. This can be useful if you want to simulate, for example, chunks of car paint flying off a car’s bodywork during a collision. If the orientation isn’t satisfactory, you can override this property by applying a **Start Rotation** value in the **Main** module. |
| **Randomize Direction** | Blends particle directions towards a random direction. A value of 0 has no effect. A value of 1 completely randomizes the particle direction. |
| **Spherize Direction** | Blends particle directions towards a spherical direction, where they travel outwards from the center of their [Transform](class-Transform.md). A value of 0 has no effect. A value of 1 points the particle direction outwards from the center, behaving identically to when the **Shape** is **Sphere**. |
| **Randomize Position** | Moves particles by a random amount, up to the specified value. A value of 0 has no effect. Any other value will apply some randomness to the spawning positions of the particles. |

### Mode dropdown options

If **Shape** is **Sphere**, **Hemisphere**, **Cone**, **Donut**, **Mesh**, **Mesh Renderer**, **Skinned Mesh Renderer**, **Circle**, or **Edge**, the **Mode** property has the following options:

| **Value** | **Description** |
| --- | --- |
| **Random** (default) | Generates particles randomly along the radius. |
| **Loop** | Generates particles sequentially along the radius of the shape, and loops back to the start at the end of each cycle. |
| **Ping-Pong** | Generates particles sequentially along the radius of the shape, and loops back to the start at the end of each cycle. Each consecutive loop happens in the opposite direction to the previous loop. |
| **Burst Spread** | Generates particles evenly along the radius. Use this mode for a more evenly distributed spread of particles compared to the default randomized behavior, where particles might clump together unevenly. Burst Spread is best used with [burst emissions](PartSysEmissionModule.md). This property is available only when **Shape** is **Sphere**, **Hemisphere**, **Cone**, **Donut**, **Circle**, or **Edge**. |

## Additional resources

* [Particle emissions and emitters](particle-emissions-emitters.md)
* [Main module reference](PartSysMainModule.md)
* [Create and view a Particle System](PartSysUsage.md)