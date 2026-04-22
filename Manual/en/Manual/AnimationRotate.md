# Rotation in animations

You can use Unity to animate your **GameObject**’s rotations. There are different methods of applying these rotations to suit your project best.

Consult [Rotation and orientation in Unity](QuaternionAndEulerRotationsInUnity.md) for more information on rotational representations.

## Rotation interpolation

You can use the [Animation window](AnimationEditorGuide.md) to choose how Unity applies rotation to your GameObject. Unity uses interpolation to calculate how a GameObject visually moves from one orientation to another in your animation.

Different interpolation methods look different in motion, but have the same result. Unity offers three types of interpolation for your animations:

![The Animation window, with the interpolation menu expanded to show the rotation interpolation options.](../uploads/Main/QuaternionInterpolation.png)

The Animation window, with the interpolation menu expanded to show the rotation interpolation options.

### Euler angle interpolation

**Euler Angles** interpolation applies the full range of motion to the GameObject specified by the angles you enter; if the rotation is greater than 360 degrees, the GameObject rotates fully before it stops at the correct orientation.

**Euler Angles (Quaternion)** interpolation uses the above interpolation method but bakes the information into a **quaternion** curve. This method uses more memory but results in a slightly faster runtime.

### Quaternion interpolation

**Quaternion** interpolation rotates the GameObject across the shortest distance to a particular orientation. For example, regardless of whether the rotation value is 5 degrees or 365 degrees, the GameObject rotates 5 degrees.

## External animation sources

[Animation from external sources](AnimationsImport.md) often contains rotational **keyframe** animation in Euler format. Unity resamples these animations and generates new keyframes for each frame in the animation to avoid rotations that exceed the valid range of rotational quaternions.

For example, if you have two keyframes that are six frames apart with the x value going from 0 to 270 degrees, the GameObject rotates 90 degrees in the opposite direction because it’s the shortest way to get to the same result. Instead, Unity resamples and adds a keyframe on every frame, so the rotation is only 45 degrees between keyframes and the rotation is correct.

### Resolve rotation problems with external animation sources

If the quaternion resampling of the imported animation doesn’t match the original closely enough for your needs, you can turn off animation resampling and use the original Euler animation keyframes at runtime. For more information, consult [Euler curve resampling](AnimationEulerCurveImport.md).

## Additional resources

* [Rotation and orientation in Unity](QuaternionAndEulerRotationsInUnity.md)
* [Using Animation Curves](animeditor-AnimationCurves.md)
* [Euler curve resampling](AnimationEulerCurveImport.md)