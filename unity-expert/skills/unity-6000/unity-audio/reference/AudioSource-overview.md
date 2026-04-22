# Introduction to the Audio Source component

Attach an **Audio Source** component to a **GameObject** to control how and where sounds play in your **scene**.

Audio sources are components that let you integrate sound effects, music, commentary, and other audio features into your application.

They interact with other audio components in Unity that allow you to edit, enhance, and output sound in your scene, including:

* Audio Clips
* Audio Random Containers
* Audio Listeners
* Audio Mixers

This page covers how the audio source interacts with these audio components. For more information about the **Audio Source** component’s properties and how to set up the component, refer to [Audio Source component reference](AudioSource-reference.md) and [Set up an Audio Source component](AudioSource-create.md).

## Audio resources

The **Audio Source** component requires an audio resource to play sound in your scene. Audio resources are containers that hold the actual audio data, so you must assign one to the Audio Source so it has audio data to edit and play. For instructions, refer to [Assign an audio resource to your audio source](AudioSource-create.md#assign-an-audio-resource-to-your-audio-source).

The following Unity file types are audio resources:

* [Audio Clip](class-AudioClip.md)
* [Audio Random Container](AudioRandomContainer.md)

Refer to those pages for more information about each type and for the audio file formats Unity supports.

## Output method of the audio source

In the **Audio Source** component, the **Output** property specifies where the audio source will send the audio signal in the audio processing pipeline.

This property accepts an **Audio Mixer Group**. The Audio Mixer is a tool that lets you post-process the audio with effects. You can then assign your Audio Mixer to the property to make sure your audio source applies your effects to the audio.

If you set the property to **None**, the sound will bypass your mixer and the audio will play without your effects. This is the default behavior.

Then, any **Audio Listener** components in the scene detects the audio from nearby audio sources, and outputs the audio to the user so they can hear it. Audio listeners are usually found on **cameras** in the scene, but you can also assign them to other objects.

For more information about these components, refer to [Audio Listener](class-AudioListener.md) and [Audio Mixer](AudioMixer.md).

## Configure your audio source

You can configure the audio source to play the clip as 2D, 3D, or as a mixture (*SpatialBlend*). The audio can be spread out between speakers (stereo to 7.1) (*Spread*) and morphed between 3D and 2D (*SpatialBlend*).

If you set **SpatialBlend** to `0.0f`, then Unity will treat the audio clip as a 2D sound. If you set it to `1.0f`, the clip is fully 3D. Anything in between is a blend of 2D and 3D.

Use falloff curves to control the spread over distance. Also, if the [listener](class-AudioListener.md) is within one or multiple [Reverb Zones](class-AudioReverbZone.md), this applies reverberation to the source. You can also apply individual filters to each audio source for an even richer audio experience. For more details, refer to [Audio Effects](class-AudioEffect.md).

For a list of Audio Source settings, refer to [Audio Source component reference](AudioSource-reference.md).

## API resources

The following is a list of useful API for AudioSource and its related properties.

* [AudioSource](../ScriptReference/AudioSource.md)
* [AudioClip](../ScriptReference/AudioClip.md)
* [AudioListener](../ScriptReference/AudioListener.md)
* [AudioMixer](../ScriptReference/Audio.AudioMixer.md)

## Additional resources

* [Audio Source](Class-AudioSource.md)
* [Introduction to the Audio Source component](AudioSource-overview.md)
* [Set up an Audio Source component](AudioSource-create.md)
* [Audio Source component reference](AudioSource-reference.md)