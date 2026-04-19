# Particle triggers

The Built-in **Particle** System’s [Triggers](PartSysTriggersModule.md) module allows you to access and modify particles based on their interaction with one or more **Colliders** in the **Scene**.

When you enable this module, the **Particle System** calls the [OnParticleTrigger()](../ScriptReference/MonoBehaviour.OnParticleTrigger.md) callback on attached **scripts**, which you can use to access lists of particles depending on where they are in respect to the Colliders in the Scene.