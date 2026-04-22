# Audio in Web

This page only provides information about audio capabilities in the Web platform. To learn how to use audio in your Unity project, refer to [Audio Overview](AudioOverview.md).

Because Unity uses [FMOD](https://www.fmod.com) to manage audio for platforms, the Web platform supports limited audio functionality, which only includes the basic features. FMOD relies on threads, which the **WebGL** API doesn’t support. For this reason, Unity uses an implementation based on the internal Web Audio API, which enables the browser to handle audio playback and mixing.

**Note:** Google Chrome’s new Autoplay policy prevents autoplay of audio and video under certain conditions. For example, while your game might be set to autoplay some background music soon after the game loads, it won’t play automatically unless you click or tap on the website. For more information on how to enable or disable this policy, refer to Google Chrome’s documentation on [Autoplay policy in Chrome](https://developer.chrome.com/blog/autoplay/).

## Supported classes

Unity Web supports the following API classes:

| **Class** | **WebGL Support status** |
| --- | --- |
| **[AudioSource](../ScriptReference/AudioSource.md)** | WebGL supports some APIs. Refer to [AudioSource](../ScriptReference/AudioSource.md) for specific support details. |
| **[AudioListener](../ScriptReference/AudioListener.md)** | All APIs supported. |
| **[AudioClip](../ScriptReference/AudioClip.md)** | WebGL supports some APIs. Refer to [AudioClip](#audioclip) for specific support details. |
| **[AudioMixer](../ScriptReference/Audio.AudioMixer.md)** | WebGL supports some APIs. Refer to [Audio Mixer](#audio-mixer) for specific support details. |
| **[SystemInfo.supportsAudio](../ScriptReference/SystemInfo-supportsAudio.md)** | The browser provides audio support for WebGL. For this reason, `SystemInfo.supportsAudio` is always true. |
| **[Microphone](../ScriptReference/Microphone.md)** | Not supported. |

## AudioSource

The [AudioSource](../ScriptReference/AudioSource.md) API supports basic positional audio playback, including:

* Pausing and resuming
* Rolloff
* Pitch setting
* Doppler effect support

Unity WebGL supports the following AudioSource APIs:

| **Settings** | **Description** |
| --- | --- |
| **[Clip](../ScriptReference/AudioSource-clip.md)** | Determines the **audio clip** that plays next. |
| **[dopplerLevel](../ScriptReference/AudioSource-dopplerLevel.md)** | Sets the Doppler scale for the AudioSource. |
| **[ignoreListenerPause](../ScriptReference/AudioSource-ignoreListenerPause.md)** | Allows AudioSource to ignore `AudioListener.pause` and continue to play audio. |
| **[ignoreListenerVolume](../ScriptReference/AudioSource-ignoreListenerVolume.md)** | Ignores the end-user’s AudioSource volume. |
| **[isPlaying](../ScriptReference/AudioSource-isPlaying.md)** | Returns true if the `AudioSource.clip` is playing. |
| **[loop](../ScriptReference/AudioSource-loop.md)** | Allows the application to loop the `AudioSource.clip`. |
| **[maxDistance](../ScriptReference/AudioSource-maxDistance.md)** | Sets the maximum distance at which the `AudioSource.clip` stops attenuating or becomes inaudible. |
| **[minDistance](../ScriptReference/AudioSource-minDistance.md)** | Sets the minimum distance at which the AudioSource.clip no longer rises in volumes. The sound starts to attenuate beyond the minimum distance. |
| **[mute](../ScriptReference/AudioSource-mute.md)** | Mutes the AudioSource. |
| **[pitch](../ScriptReference/AudioSource-pitch.md)** | Sets the pitch of the `AudioSource.clip`. WebGL only supports positive pitch values. |
| **[playOnAwake](../ScriptReference/AudioSource-playOnAwake.md)** | Plays the AudioSource on Awake. |
| **[rolloffMode](../ScriptReference/AudioSource-rolloffMode.md)** | Sets the AudioSource attenuation over distance. |
| **[time](../ScriptReference/AudioSource-time.md)** | Sets the playback position in seconds. |
| **[timeSamples](../ScriptReference/AudioSource-timeSamples.md)** | Sets the playback position in Pulse-code modulation (PCM) samples. |
| **[velocityUpdateMode](../ScriptReference/AudioSource-velocityUpdateMode.md)** | Sets whether the AudioSource updates in the fixed or dynamic update loop. |
| **[volume](../ScriptReference/AudioSource-volume.md)** | Sets the volume of the AudioSource (0.0 to 1.0). |
| **[Pause](../ScriptReference/AudioSource.Pause.md)** | Pauses the `AudioSource.clip`. |
| **[Play](../ScriptReference/AudioSource.Play.md)** | Plays the `AudioSource.clip`. |
| **[PlayDelayed](../ScriptReference/AudioSource.PlayDelayed.md)** | Plays the `AudioSource.clip` with a delay you specify in seconds. |
| **[PlayOneShot](../ScriptReference/AudioSource.PlayOneShot.md)** | Plays an [AudioClip](../ScriptReference/AudioSource-clip.md) and scales the AudioSource volume by volumeScale. |
| **[PlayScheduled](../ScriptReference/AudioSource.PlayScheduled.md)** | Plays the AudioSource at a time you specify. |
| **[SetScheduledEndTime](../ScriptReference/AudioSource.SetScheduledEndTime.md)** | Sets a time that a scheduled `AudioSource.clip` ends. |
| **[SetScheduledStartTime](../ScriptReference/AudioSource.SetScheduledStartTime.md)** | Sets the time that a scheduled `AudioSource.clip` starts. |
| **[Stop](../ScriptReference/AudioSource.Stop.md)** | Stops playing the `AudioSource.clip`. |
| **[UnPause](../ScriptReference/AudioSource.UnPause.md)** | Unpauses a paused `AudioSource.clip`. |
| **[PlayClipAtPoint](../ScriptReference/AudioSource.PlayClipAtPoint.md)** | Plays an `AudioSource.clip` at a given position in the worldspace. |

**Note:** Because audio data is decoded in the browser with [`decodeAudioData`](https://webaudio.github.io/web-audio-api/#dom-baseaudiocontext-decodeaudiodata), it is possible that the runtime sampling rate of the audio file is different from the serialized audio data (44100Hz) as the browser will match the `BaseAudioContext` sampling rate, generally aligned with the device audio driver setup. Considering this, we recommend validating the [`AudioClip.frequency`](../ScriptReference/AudioClip-frequency.md) value at runtime or using [`AudioSource.time`](../ScriptReference/AudioSource-time.md) when performing seeking operations.

## AudioClip

Unity WebGL imports [AudioClip](../ScriptReference/AudioClip.md) files in the Advanced Audio Coding (AAC) Format, which is supported by most browsers. Unity WebGL supports the following AudioClip APIs:

| **Properties** | **Description** |
| --- | --- |
| **[frequency](../ScriptReference/AudioClip-frequency.md)** | The sample frequency of the clip in Hertz. |
| **[length](../ScriptReference/AudioClip-length.md)** | The length of the AudioClip in seconds. |
| **[loadState](../ScriptReference/AudioClip-loadState.md)** | Returns the current load state of the audio data associated with an AudioClip. |
| **[samples](../ScriptReference/AudioClip-samples.md)** | The length of the AudioClip in samples. |
| **[loadType](../ScriptReference/AudioClip-loadType.md)** | The load type of the clip. You can set the AudioClip load type in the **Inspector**. |

| **Method** | **Description** | **Additional information** |
| --- | --- | --- |
| **[AudioClip.Create](../ScriptReference/AudioClip.Create.md)** | Creates an AudioClip with a name and length you specify. | Unity WebGL partially supports `AudioClip.Create`. Browsers don’t support dynamic streaming, so to use `AudioClip.Create`, set the Stream to false. |
| **[AudioClip.SetData](../ScriptReference/AudioClip.SetData.md)** | Sets sample data in an AudioSource.clip. | Unity WebGL partially supports `AudioClip.SetData`. You can use this method only on compressed audio files with **Load Type** set to **Decompress on Load**. Refer to [Compressed audio](#compressed-audio). |
| **[AudioClip.GetData](../ScriptReference/AudioClip.GetData.md)** | Retrieves an array with sample data from an AudioSource.clip. | Unity WebGL partially supports `AudioClip.GetData`. You can use this method only on compressed audio files with **Load Type** set to **Decompress on Load**. Refer to [Compressed audio](#compressed-audio). |

**Note:** For audio clip support on Linux, make sure you’ve installed the [ffmpeg](https://ffmpeg.org/) package.

### Audio looping issues

When looping an audio clip on a WebGL build, audio might have audible glitches at the loop point. Because AAC is a licensed format, Unity relies on the host platform’s encoder to convert files to AAC. Encoders might alter the first 1024 samples during encoding, which affects the loop start point and causes audible glitches when the sound loops. To address this issue:

1. Use WAV as your source file format.
2. Add at least 1024 samples of silence to the beginning of your WAV file.
3. Edit the WAV’s `smpl` chunk so the loop starts after the added silence.

## Audio Mixer

Unity Web supports some functionality of [Audio Mixer](AudioMixer.md) assets.

You can do the following with Audio Mixers on Web:

* Create an Audio Mixer asset.
* Add AudioMixerGroups to the hierarchy.
* Adjust the volume of each group. To expose or change the volume with scripting, use [AudioMixer.SetFloat](../ScriptReference/Audio.AudioMixer.SetFloat.md).

**Note**: Volume is the only property you can change on Web. Other properties and sound effects aren’t supported.

## Compressed audio

To use compressed audio with WebGL in Unity, set the AudioClip [loadType](../ScriptReference/AudioClip-loadType.md) to one of the following options:

* [CompressedInMemory](../ScriptReference/AudioClipLoadType.CompressedInMemory.md)
* [DecompressOnLoad](../ScriptReference/AudioClipLoadType.DecompressOnLoad.md)

You can also change these settings in the Unity Editor. To change the setting, select your Audio Clip, then set **Load Type** in the Inspector window to your preferred option (**Decompress On Load** or **Compressed In Memory**).

| ****Compression** method** | **Description** | **Considerations** |
| --- | --- | --- |
| **CompressedInMemory** | Use this to compress the audio on disk and have it remain compressed after it loads into your application memory. | Compressed audio can cause latency and is less precise when it comes to audio playback. However, compressed audio uses less memory in your application than decompressed audio. It’s best practice to use `CompressedInMemory` for audio that’s unaffected by precision for example, background music. |
| **DecompressOnLoad** | Use this to compress the audio on disk, similar to CompressedInMemory, and decompress when it loads into your application memory. | Decompressed audio uses a significant amount of memory compared to compressed audio but has lower latency and more audio flexibility. Use `DecompressedOnLoad` for audio that’s affected by precision (for example, character dialog or sound effects). |

### Issues with Silent Mode on iOS browsers

The Silent Mode switch on iOS devices behaves differently on compressed and uncompressed audio clips in Web builds. Uncompressed sounds aren’t audible on iOS devices in Silent Mode because WebKit categorizes this audio node type differently than `MediaElementSourceNode`, which is used for `CompressedInMemory` sounds. For more information, refer to [WebKit issue #262781](https://bugs.webkit.org/show_bug.cgi?id=262781).

If you set your audio clip’s `loadType` to `DecompressOnLoad`, the clip will not be audible on iOS devices that are in Silent Mode. To ensure audio plays regardless of Silent Mode, set your audio clip `loadType` to `CompressedInMemory`. This issue occurs only in Web builds.

If you need to load in a sound when a **scene** starts, it’s recommended to load the sound completely in a temporary scene and [load the target scene additively](../ScriptReference/SceneManagement.LoadSceneMode.Additive.md).

## Audio playback and browser security

For security reasons, browsers don’t allow audio playback until an end user interacts with your application webpage via a mouse click, touch event, or key press. Use a loading screen to allow the end user to interact with your application and start audio playback before your main content begins.

## Scriptable Audio Pipeline

The scriptable audio pipeline is not supported. Attempting to use it in a Web build will trigger a warning.