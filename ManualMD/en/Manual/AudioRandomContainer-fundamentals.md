# Audio Random Container fundamentals

An Audio Random Container is an object that lets you create audio playlists for your **scene** and apply rules to determine when and how the clips play. For example, you can add randomization, looping, and timed playback to your audio clips. This is an excellent way to create varied and dynamic audio playback for music and sound effects like footsteps, weapon hits, and background music.

The audio cycle and AudioSource API exceptions are important concepts to know before you use an Audio Random Container.

## Audio cycle

An audio cycle is the full **Audio Clips** list length. If the list has three audio clips, an audio cycle is three clips.

### AudioSource API

Use the [AudioSource API](../ScriptReference/AudioSource.md) to start, pause, and stop an Audio Random Container. This is similar to when you use the AudioSource API to play, pause, and stop an Audio Clip, but with the following exceptions:

* [AudioSource.isPlaying](../ScriptReference/AudioSource.isPlaying.md) returns true when the Audio Random Container plays an audio clip through an **audio source**.
* [AudioSource.Play](../ScriptReference/AudioSource.Play.md) behaves differently depending on whether you set the Audio Random Container to **Manual** or **Automatic**.
  + When set to **Manual**, `AudioSource.Play` plays an audio clip in the **Audio Clips** list based on the container’s **Playback Mode**. For example, if the **Playback Mode** is set to **Random**, it plays a random audio clip in the list.

## Additional resources

* [Audio playlist randomization](AudioRandomContainer.md)
* [Audio Clips](class-AudioClip.md)
* [AudioSource.Play](../ScriptReference/AudioSource.Play.md)
* [Audio Source priority](class-AudioSource.md)
* [Audio Mixer](AudioMixer.md)