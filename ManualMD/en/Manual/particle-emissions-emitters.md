# Particle emissions and emitters

Understand particle emissions, and how Unity manages them via **Particle System** modules.

## Rate of emission

The rate of emission can be constant or can vary over the lifetime of the system according to a curve. If the [Emission module](PartSysEmissionModule.md) **Rate over Distance** mode is active, a certain number of particles are released per unit of distance moved by the parent object. This is very useful for simulating particles that are actually created by the motion of the object (for example, dust from a car’s wheels on a dirt track).

If **Rate over Time** is active, then the desired number of particles are emitted each second regardless of how the parent object moves. Additionally, you can add bursts of extra particles that appear at specific times (for example, a steam train chimney that produces puffs of smoke).

## Emitter shape

The [Shape](PartSysShapeModule.md) module defines the volume or surface from which [particles](class-ParticleSystem.md) can be emitted, and the direction of the start velocity. The **Shape** property defines the shape of the emission volume, and the rest of the module properties vary depending on the Shape you choose.

All shapes (except [Mesh](class-Mesh.md)) have properties that define their dimensions, such as the **Radius** property. To edit these, drag the handles on the wireframe emitter shape in the **Scene** view. The choice of shape affects the region from which particles can be launched, but also the initial direction of the particles. For example, a **Sphere** emits particles outward in all directions, a **Cone** emits a diverging stream of particles, and a **Mesh** emits particles in directions that are normal to the surface.

### Configuring mesh-based emitter shapes

You can choose to only emit particles from a particular material (sub-Mesh) by checking the **Single Material** property and you can offset the emission position along the Mesh’s normals by checking the **Normal Offset** property.

To ignore the color of the Mesh, check the **Use Mesh Colors** property. To read the texture colors from a mesh, assign the Texture you wish to read to the **Texture** property.

Meshes must be read/write enabled to work on the Particle System. If you assign them in the Editor, Unity handles this for you. But if you want to assign different meshes at run time, you need to check the **Read/Write Enabled** setting in the [Import Settings](FBXImporter-Model.md).

The **Mode** property allows the Particle System to emit particles in a predictable order on the surface of a Mesh. When emitting from Vertices, this property allows you to emit each new particle from the next vertex in the array of vertices in the Mesh. When emitting from Edges, the Particle System can emit particles smoothly along the edges of the Mesh’s triangles/lines.

## Particle lifetime based on emitter speed

The [Lifetime by Emitter Speed](PartSysLifetimeByEmitterSpeedModule.md) controls the initial lifetime of each particle based on the speed of the emitter when the particle spawns. It multiplies the start lifetime of particles by a value that depends on the speed of the object that spawned them. For most Particle Systems, this is the **GameObject** velocity, but for sub-emitters, the velocity comes from the parent particle that the sub-emitter particle originated from.

## Particle emissions from other particles

Many types of particles produce effects at different stages of their lifetimes that can also be implemented using Particle System [Sub Emitters](PartSysSubEmitModule.md). For example, a bullet might be accompanied by a puff of smoke powder as it leaves the gun barrel, and a fireball might explode on impact. You can use sub-emitters to create effects like these.

Sub-emitters are ordinary Particle System objects created in the Scene. This means that sub-emitters can have sub-emitters of their own (this type of arrangement can be useful for complex effects like fireworks). However, it is very easy to generate an enormous number of particles using sub-emitters, which can be resource-intensive.

To trigger a sub-emitter, you can use these are the conditions:

* **Birth**: When the particles are created.
* **Collision**: When the particles collide with an object.
* **Death**: When the particles are destroyed.
* **Trigger**: When the particles interact with a Trigger **collider**.
* **Manual**: Only triggered when requested via script. See [ParticleSystem.TriggerSubEmitter](../ScriptReference/ParticleSystem.TriggerSubEmitter.md).

Note that the **Collision**, **Trigger**, **Death** and **Manual** events can only use burst emission in the [Emission](PartSysEmissionModule.md) module.

Additionally, you can transfer properties from the parent particle to each newly created particle using the **Inherit** options. The transferrable properties are size, rotation, color and lifetime. To control how velocity is inherited, configure the [Inherit Velocity](PartSysInheritVelocity.md) module on the sub-emitter system.

It is also possible to configure a probability that a sub-emitter event will trigger, by setting the **Emit Probability** property. A value of 1 guarantees that the event will trigger, whereas lower values reduce the probability.

## Additional resources

* [Shape Module reference](PartSysShapeModule.md)
* [Emission Module reference](PartSysEmissionModule.md)
* [Sub Emitters Module reference](PartSysSubEmitModule.md)
* [Lifetime By Emitter Speed Module reference](PartSysLifetimeByEmitterSpeedModule.md)