# Playables API examples

The **Playables** API examples include sample **scripts** that demonstrate how to perform basic animation tasks such as playing, mixing, and controlling **animation clips** or other playables. The Playables API examples include the following topics:

| **Topic** | **Description** |
| --- | --- |
| **[PlayableGraph Visualizer](playables-visualizer.md)** | The Playable API examples use the PlayableGraph Visualizer package to visualize the result of most scripts. This topic demonstrates how to download and install the PlayableGraph Visualizer package. **Note:** The PlayableGraph Visualizer is a discontinued experimental package that might not work with your version of Unity. |
| **[Play an animation clip](playables-ex-play-clip.md)** | Learn how to create a `PlayableGraph` that plays a single animation clip. |
| **[Blend two animation clips](playables-ex-blend-clips.md)** | Learn how to use a `AnimationMixerPlayable` to blend two animation clips. |
| **[Blend an animation clip with a controller](playables-ex-blend-clip-controller.md)** | Learn how to use a `AnimatorMixerPlayable` to blend an animation clip with an **Animator Controller**. |
| **[Control the play state](playables-ex-play-state.md)** | Learn about the `Pause()` and `Play()` methods and how to control the play state of a `PlayableGraph`. |
| **[Control the timing of a playable clip](playables-ex-control-timing.md)** | Learn about the `SetTime()` method and how to manually set the start of a paused playable. |
| **[Create a PlayableGraph with different outputs](playables-ex-different-outputs.md)** | Learn how to create a `PlayableGraph` with two different playable output types. |
| **[Create a custom playable](playables-ex-custom-playable.md)** | Learn about the `PlayableBehaviour` public class to create a custom Playable. This example also demonstrate how to override the `PrepareFrame()` virtual method. |

## Additional resources

* [The PlayableGraph, nodes, and output](Playables-Graph.md)
* [Playable struct](../ScriptReference/Playables.Playable.md)
* [Get started with packages](Packages.md)