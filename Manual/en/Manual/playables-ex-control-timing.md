# Control the timing of a playable clip

This example demonstrates how to use the `SetTime()` method to manually set the start of a playable clip. This example also demonstrates how to use the `Pause()` method to pause a playable clip.

## Prerequisites

Before you use the `ControlTiming` script in your project, your project must have the following:

* A **GameObject**, such as a cube or a capsule. You don’t need to manually add an **Animator component** to this GameObject. The `RequireComponent` attribute adds this component if it’s not present.
* An animation clip that animates the properties of the GameObject. For example, a clip that animates the position and rotation of the GameObject.

## Add and run the script

To use the `ControlTiming` script in your project, follow these steps:

1. Add a script component to your GameObject. Name the script file `ControlTiming.cs` and use the following code:

   ```
   using UnityEngine;
   using UnityEngine.Playables;
   using UnityEngine.Animations;

   [RequireComponent(typeof(Animator))]
   public class ControlTiming : MonoBehaviour
   {
       public AnimationClip clip;
       public float time;
       PlayableGraph graph;
       AnimationClipPlayable clipPlayable;

       void Start()
       {
           // Create and name the graph.
           graph = PlayableGraph.Create("ControlTiming");

           var output = AnimationPlayableOutput.Create(graph, "Animation", GetComponent<Animator>());

           // Wrap the clip in a playable.
           clipPlayable = AnimationClipPlayable.Create(graph, clip);

           // Connect the Playable to an output.
           output.SetSourcePlayable(clipPlayable);

           // Play the graph.
           graph.Play();

           // Pause the clip playable. This stops time from progressing automatically.
           clipPlayable.Pause();
       }

       void Update ()
       {
           // Control the time manually.
           clipPlayable.SetTime(time);
       }

       void OnDisable()
       {
           // Destroy all Playables and Outputs created by the graph.
           graph.Destroy();
       }
   }
   ```
2. In the Script component, select the animation clip for the paused clip. You can also set the start time in seconds. If you set a value that’s greater than the length of the clip and the clip doesn’t loop, the clip plays the last frame.
3. Select **Play** to switch the Editor to Play mode.
4. In the Script component, experiment with different start times. Because the animation clip is paused, only the first frame of the animation displays.
5. If you have installed the [PlayableGraph Visualizer package](playables-visualizer.md), use it to display the PlayableGraph.

## Additional resources

* [Animation clips](AnimationClips.md)
* [The SetTime method](../ScriptReference/Playables.PlayableExtensions.SetTime.md)
* [Creating scripts](creating-scripts.md)
* [Play mode in the Game view](GameView.md)