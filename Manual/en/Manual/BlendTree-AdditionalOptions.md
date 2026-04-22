# Common Blend Tree Options

The options in this topic are common to both 1D and 2D blending.

## Time Scale

You can alter the “natural” speed of **animation clips** with the animation speed number fields (the columns with a clock icon). For example, to double the speed of a walk clip, specify a value of `2.0`.

To rescale the speed of each clip in the motions list relative to the minimum and maximum values, Select **Adjust Time Scale** > **Homogeneous Speed**. This preserves the initial relative speed of each clip.

**Note**: The **Adjust Time Scale** is only available if all the clips are animation clips and not child Blend Trees.

## Mirroring

Enable the Mirror checkbox to mirror any **humanoid Animation** Clip in the motions list. Mirroring uses the same animation as its unmirrored original without doubling the memory and space.

However, mirroring an animation in a blend tree does not create a fully symmetrical animation. Unity automatically adds offsets to ensure that run cycles, walk cycles, and animations such as footsteps are blended correctly with other animation clips and blend trees. For example, if your blend tree has a humanoid running left, and you enable Mirror to switch the humanoid to run right, the foot cycles need to match so that the left foot touches the ground at the same time. This ensures that the mirrored Blend tree correctly blends with surrounding clips and other blend trees that have not been mirrored.