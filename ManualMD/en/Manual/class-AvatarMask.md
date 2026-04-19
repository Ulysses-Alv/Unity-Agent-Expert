# Avatar Mask window reference

[Switch to Scripting](../ScriptReference/AvatarMask.md "Go to AvatarMask page in the Scripting Reference")

There are two ways to define which parts of your animation should be masked:

* By selecting from a [Humanoid body map](#Humanoid)
* By choosing which bones to include or exclude from a [Transform hierarchy](#Transform)

## Humanoid body selection

If your animation uses a Humanoid **Avatar**, you can select or deselect portions of the simplified humanoid body diagram to indicate where to mask the animation:

![Defining an Avatar Mask using the Humanoid body](../uploads/Main/AvatarMaskInspectorHumanoid.png)

Defining an Avatar Mask using the Humanoid body

The body diagram groups body parts into these portions:

* Head
* Left Arm
* Right Arm
* Left Hand
* Right Hand
* Left Leg
* Right Leg
* Root (denoted by the “shadow” under the feet)

To include animation from one of these body portions, click the Avatar diagram for that portion until it appears as green.
To exclude animation, click the body portion until it appears red.
To include or exclude all, double-click the empty space surrounding the Avatar.

You can also toggle **Inverse** Kinematics\_\_ (IK)\_\_ for hands and feet, which determines whether or not to include IK curves in animation blending.

## Transform selection

If your animation does not use a Humanoid Avatar and you want more detailed control over which bones are masked, you can select or deselect portions of the Model’s hierarchy:

1. Assign a reference to the Avatar whose transform you would like to mask.
2. Click the **Import Skeleton** button. The hierarchy of the avatar appears in the **inspector**.
3. You can check each bone in the hierarchy to use as your mask.

![Defining an avatar mask using the Transform method](../uploads/Main/AvatarMaskInspectorTransform.png)

Defining an avatar mask using the Transform method

Mask assets can be used in [Animator Controllers](class-AnimatorController.md), when specifying [Animation Layers](AnimationLayers.md) to apply masking at runtime, or in the import settings of your animation files to apply masking while importing animation.

A benefit of using Masks is that they tend to reduce memory overheads since body parts that are not active do not need their associated **animation curves**. Also, the unused curves need not be calculated during playback which will tend to reduce the CPU overhead of the animation.

AvatarMask