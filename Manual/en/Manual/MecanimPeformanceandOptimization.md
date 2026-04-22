# Performance and optimization

This page contains some tips to help you obtain the best performance in Unity, covering the [animation system](#anim) and [run-time optimizations](#runtime).

**Note:** For tips on modeling your character in a 3d application for best performance in Unity, see [Modeling characters for optimal performance](ModelingOptimizedCharacters.md).

## Animation system

### Controllers

The [Animator](class-Animator.md) doesn’t spend time processing when a [Controller](class-AnimatorController.md) is not set to it.

### Simple animation

Playing a single [Animation Clip](class-AnimationClip.md) with no blending can make Unity slower than the [legacy animation system](Animations.md). The old system is very direct, sampling the curve and directly writing into the transform. Unity’s current animation system has temporary buffers it uses for blending, and there is additional copying of the sampled curve and other data. The current system layout is optimized for animation blending and more complex setups.

### Scale curves

Animating scale curves is more expensive than animating translation and rotation curves. To improve performance, avoid scale animations.

**Note:** This does not apply to constant curves (curves that have the same value for the length of the [animation clip](AnimationClips.md)). Constant curves are optimized, and are less expensive than normal curves. Constant curves that have the same values as the default **scene** values do not write to the scene every frame.

### Layers

Most of the time Unity is evaluating animations, and keeps the overhead for [animation layers](AnimationLayers.md) and [Animation State Machines](AnimationStateMachines.md) to the minimum. The cost of adding another layer to the animator, synchronized or not, depends on what animations and blend trees are played by the layer. When the weight of the layer is zero, Unity skips the layer update.

### Humanoid vs. Generic animation types

These are tips to help you decide between these types:

* When importing **Humanoid animation** use an **Avatar** Mask (class-AvatarMask) to remove IK Goals or finger animation if you don’t need them.
* When you use Generic, using **root motion** is more expensive than not using it. If your animations don’t use root motion, make sure that you have not specified a root bone.

### Scene-level optimization

There are many optimizations that can be made, some useful tips include:

* Use hashes instead of strings to query the Animator.
* Implement a small AI Layer to control the Animator. You can make it provide simple callbacks for OnStateChange, OnTransitionBegin, and other events.
* Use State Tags to easily match your AI **state machine** to the Unity state machine.
* Use additional curves to simulate events.
* Use additional curves to mark up your animations; for example, in conjunction with [target matching](TargetMatching.md).

## Runtime Optimizations

### Visibility and updates

Always optimize animations by setting the animator [Culling Mode](class-Animator.md) to **Cull Completely**, and disable the [skinned mesh renderer’s](class-SkinnedMeshRenderer.md) **Update When Offscreen** property. This saves Unity from updating animations when the character is not visible.