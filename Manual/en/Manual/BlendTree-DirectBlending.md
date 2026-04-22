# Direct blending

Use a **Direct Blend Tree** to map animator parameters to the weight of a BlendTree child. This is useful when you want exact control over blending animations rather than the indirectly control provided by 1D blending or 2D blending.

![A Direct Blend Tree with five animation clips assigned.](../uploads/Main/AnimatorDirectBlendTree.png)

A Direct Blend Tree with five animation clips assigned.

In direct blending, you use the **Inspector** window to add motions and to also assign an [Animator Parameter](AnimationParameters.md) to each blend weight. This Direct mode bypasses 2D blending algorithms, such as Freeform Directional and Freeform Cartesian, and allows you to implement a scripting solution to control the mix of blended animations.

For example, you can use Direct mode to to mix blend shape animations for facial expressions or to blend additive animations.

![Example of using Direct Blending for facial expressions](../uploads/Main/AnimatorDirectBlendTreeFacialExpressions.jpg)

Example of using Direct Blending for facial expressions