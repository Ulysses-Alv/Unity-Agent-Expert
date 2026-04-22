# Legacy Animation component

[Switch to Scripting](../ScriptReference/Animation.md "Go to Animation page in the Scripting Reference")

This is the Legacy Animation component, which was used on **GameObjects** for animation purposes prior to the introduction of the Mecanim Animation system.
This component is retained in Unity for backwards compatibility. For new projects, use the [Animator component](class-Animator.md).

![The Animation Inspector](../uploads/Main/AnimationInspector35.png)

The Animation Inspector

## Properties

| **Property** | **Function** |
| --- | --- |
| **Animation** | The default animation to play when **Play Automatically** is enabled. |
| **Animations** | A list of animations that you can access from **scripts**. |
| **Play Automatically** | Enable this option to play animation automatically when the game starts. |
| **Animate Physics** | Enable this option to have the animation interact with Physics. |
| **Culling Type** | Determine when not to play the animation.  * **Always Animate**: Always animate. * **Based on Renderers**: Cull based on the default animation pose. |

Consult the [Animation Window Guide](AnimationEditorGuide.md) for more information on how to create animations inside Unity.
Consult [Model import workflows](ImportingModelFiles.md) page on how to import animated characters.