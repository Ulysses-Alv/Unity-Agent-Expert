# PlayableGraph, its nodes, and output

Use a `PlayableGraph` to mix, blend, and modify multiple data sources. You can then play the result through a `PlayableOutput`. Each `Playable` in a PlayableGraph can support animation, audio, or **scripts**.

The main parts of a `PlayableGraph` are as follows:

* [The PlayableGraph itself](#playablegraph) which represents the flow of data and indicates what each node produces and consumes.
* [The nodes in the PlayableGraph](#playable). Each node is a playable.
* [The PlayableOutput](#playableoutput) which defines the outputs of the PlayableGraph.

## PlayableGraph

The PlayableGraph provides a flexible graph with time-based synchronization of multiple data sources. The Unity Engine implements data sources for animation and audio. Use scripting to create custom data sources.

In comparison, the Mecanim animation system includes a **state machine** graph that you can use to create transitions between animation states. The state machine graph only supports animation.

The PlayableGraph defines a set of playable outputs bound to a **GameObject** or a component. The PlayableGraph also defines **playables** and their relationships with each other.

![A sample PlayableGraph](../uploads/Main/PlayablesGraph0.png)

A sample PlayableGraph

The PlayableGraph manages the lifecycle of its playables and their outputs. The following lists a few of the most commonly used PlayableGraph methods:

* Create a new PlayableGraph with the `PlayableGraph.Create()` static method.
* Play a PlayableGraph with the `PlayableGraph.Play()` method.
* Stop a playing PlayableGraph with the`PlayableGraph.Stop()` method.
* Manually evaluate the state of a PlayableGraph with the `PlayableGraph.Evaluate()` method.
* Connect two playables together with the `PlayableGraph.Connect()` method.
* Destroy a PlayableGraph manually with the `PlayableGraph.Destroy()` method. This method automatically destroys all playables and playable outputs created by the PlayableGraph. If you don’t destroy a PlayableGraph, Unity issues an error message.

For a full list of PlayableGraph methods and for more details, refer to [the PlayableGraph struct](ScriptReference/Playables.PlayableGraph.md) in the Scripting Reference.

## Playable

A playable is a C# struct that implements the `IPlayable` interface. A playable is a node that connects to other playables and defines its relationship with other playables.

![Core playable types](../uploads/Main/PlayablesGraph1.png)

Core playable types

Unity implements core playable types as C# structs to avoid allocating [memory for garbage collection](performance-managed-memory.md). The `Playable` class defines a few basic methods and is the base type for all playables. The `PlayableExtensions` static class provides most of the methods for interacting with playables.

Because `Playable` is the base type for all playables, you can implicitly cast a playable to it. However, the opposite isn’t true. If you explicitly cast a `Playable` onto an incompatible type, it throws an exception.

The types that inherit from the `Playable` base type have additional, type specific methods. For example, the `AnimationClipPlayable` type includes specific methods for an **animation clip** encapsulated in a playable. To access type specific methods, you must cast your playable to the appropriate type.

Non-abstract playables have the `Create()` public static method which creates a playable of the corresponding type. The `Create()` method always takes a `PlayableGraph` as its first parameter, and that graph owns the new playable. Some playable types might require additional parameters.

## Playable output

A playable output is a C# struct that implements an `IPlayableOutput`. It defines the output of a PlayableGraph.

![Core playable output types](../uploads/Main/PlayablesGraph2.png)

Core playable output types

Unity implements playable output types as C# structs to avoid allocating [memory for garbage collection](performance-managed-memory.md).

The `PlayableOutput` class defines a few basic methods and is the base type for all playable outputs. The `PlayableOutputExtensions` static class provides most of the methods for interacting with playable outputs.

It’s recommended that you link a valid playable output to a playable. If a playable output isn’t linked to a playable, the playable output does nothing. To link a playable output to a playable, use the `PlayableOutput.SetSourcePlayable()` method. The linked playable acts as the root for that specific playable output.

Non-abstract playable outputs have a public static method, `Create()`, that creates a playable output of the corresponding type. The `Create()` method always takes a PlayableGraph as its first parameter, and that graph owns the new playable output. Some playable output types might require additional parameters.

## Additional resources

* [PlayableGraph struct](../ScriptReference/Playables.PlayableGraph.md)
* [Playable struct](../ScriptReference/Playables.Playable.md)
* [Playable Extensions static class](../ScriptReference/Playables.PlayableExtensions.md)
* [PlayableOutput Extensions static class](../ScriptReference/Playables.PlayableOutputExtensions.md)