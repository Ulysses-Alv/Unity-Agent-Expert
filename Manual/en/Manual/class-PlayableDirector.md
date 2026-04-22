# Playable Director component

Use the Playable Director component to build a `PlayableGraph` from a `PlayableAsset`. Once built, you can then use the Playable Director component to manage the playback and timing of the `PlayableGraph` through a `PlayableDirector` object.

The `PlayableDirector` object includes methods that control playback, retrieve or set binding information, and subscribe to playback state notifications. Refer to [the PlayableDirector class](../ScriptReference/Playables.PlayableDirector.md) for more information and an example of how to use the **Playables** API to interact with the Playable Director component.

The Playable Director component is primarily used by [Unity’s Timeline](com.unity.timeline.md) to do the following:

* Store the link between a Timeline instance and a Timeline asset.
* Store the list of Timeline asset tracks and their bindings.
* Control when the Timeline instance plays and how it updates its clock.
* Control what happens when the Timeline instance finishes playing.

Refer to [Playable Director component](https://docs.unity3d.com/Packages/com.unity.timeline@latest/index.html?subfolder=/manual/playable-director.html) in the Timeline package documentation for more information.

## Additional resources

* [PlayableDirector class](../ScriptReference/Playables.PlayableDirector.md).
* [PlayableAsset class](../ScriptReference/Playables.PlayableAsset.md).
* [Unity Timeline package](https://docs.unity3d.com/Packages/com.unity.timeline@latest).