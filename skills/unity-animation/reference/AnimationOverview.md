# Mecanim Animation system

The Mecanim Animation system is based on [animation clips](AnimationClips.md). An animation clip contains information about how a character or **GameObject** changes its position, rotation, or other properties over time. Think of each animation clip as a recording. An animation clip from an external source is created by an artist or animator with third-party software, such as Autodesk® 3ds Max® or Autodesk® Maya®.

You can organize many animation clips into a structured flowchart-like system called an [Animator Controller](class-AnimatorController.md). The Animator Controller acts as a [state machine](AnimationStateMachines.md) which keeps track of the animation clip being played and the criteria for when the next animation clip should be played.

For example, a simple Animator Controller might contain a single animation clip for a collectable item that spins and bounces. An Animator Controller for a door might contain two animation clips: one for the door opening and a second for the door closing.

A more advanced Animator Controller might contain a dozen **humanoid animations** for each of a character’s actions such as idle, turning, walking, and jogging. This Animator Controller might use [blend trees](class-BlendTree.md) to blend between multiple animation clips and provide fluid motion as the player moves around the **scene**.

The Mecanim Animation system also provides special features for handling humanoid characters including the ability to [retarget](Retargeting.md) animation from one source to another and adjust [muscle definitions](MuscleDefinitions.md) to ensure that a character deforms correctly. For example, you can retarget animation from a third-party animation library to your own character then adjust the its range of motion. These special features are available through Unity’s [Avatar](AvatarCreationandSetup.md) system which maps humanoid characters to a common internal format.

The Animation Clips, the Animator Controller, and the Avatar, are connected together through a GameObject with the [Animator Component](class-Animator.md). This component references an Animator Controller and, if required, the Avatar for the model. The Animator Controller also references its animation clips.

![How the parts of the Mecanim Animation system connect together](../uploads/Main/MecanimHowItFitsTogether.jpg)

How the parts of the Mecanim Animation system connect together

The previous diagram demonstrates the following workflow:

1. Animation clips are [imported from an external source](class-AnimationClip.md) or created within Unity. In this example, they are [imported humanoid animations](ConfiguringtheAvatar.md).
2. Animation clips are placed and arranged in an Animator Controller. This demonstrates an Animator Controller in the **Animator window**. The States, which may represent animations or nested sub-state machines, display as nodes connected by lines. This Animator Controller exists as an asset in the **Project window**.
3. The rigged character model, an astronaut named `Astrella`, has a specific configuration of bones mapped to Unity’s [Avatar](class-Avatar.md) format. The mapping is stored as an Avatar asset, under the imported character model in the Project window.
4. When animating the character model, it has an Animator component attached. In the **Inspector** window, the [Animator Component](class-Animator.md) has the [Animator Controller](class-AnimatorController.md) and the [Avatar](class-Avatar.md) assigned. The animator uses these together to animate the model. The Avatar reference is only necessary when animating a humanoid character. For [other types of animation](GenericAnimations.md), only an Animator Controller is required.

The following topics provide more details on the Mecanim Animation system:

* [Animation Clips](AnimationClips.md)
* [Humanoid Avatars](AvatarCreationandSetup.md)
* [Animator component](class-Animator.md)
* [Animator Controller](class-AnimatorController.md)
* [Animator Override Controller](AnimatorOverrideController.md)
* [Playables API](Playables.md)
* [Performance and Optimization](MecanimPeformanceandOptimization.md)
* [Mecanim FAQ](MecanimFAQ.md)