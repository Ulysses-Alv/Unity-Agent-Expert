# Audio Clip Import Settings reference

[Switch to Scripting](../ScriptReference/AudioClip.md "Go to AudioClip page in the Scripting Reference")

Explore the properties and settings to specify the sampling rate, **compression** format, and quality of your imported audio assets.

For a list of the audio files Unity supports and information about how to import these files, refer to [Audio file compatibility](AudioFiles-compatibility.md) and [Import audio files into Unity](AudioFiles-import.md).

![Screenshot of the Audio File Import Settings Inspector window. Shows configurable options for audio files.](../uploads/Main/AudioClipImporter50.png)

Screenshot of the Audio File Import Settings Inspector window. Shows configurable options for audio files.

The audio file used in this image is an asset from the [FREE Casual Game SFX Pack](https://assetstore.unity.com/packages/audio/sound-fx/free-casual-game-sfx-pack-54116) by Dustyroom, available on the Unity **Asset Store**.

## Properties

The following properties are available in the **Inspector** for your imported audio files:

| **Property** | **Function** |
| --- | --- |
| **Force To Mono** | Enable to mix multi-channel audio down to a mono track before packing occurs. |
| **Normalize** | Enable to [normalize](https://en.wikipedia.org/wiki/Audio_normalization) audio during the **Force To Mono** mixing down process. |
| **Load In Background** | Enable to ensure Unity loads the **audio clip** asynchronously (in the background) when the game starts, which takes pressure off the main thread. This setting can enhance performance if you need to load large audio files. |
| **Ambisonic** | Enable this option if your audio file contains Ambisonic-encoded audio. Ambisonic **audio sources** store audio in a format which represents a soundfield that can be rotated based on the listener’s orientation. It’s useful for 360-degree videos and **XR** applications. |

## Platform-specific audio clip override settings

The following table shows the options available in the platform-specific overrides panel in the **Audio Clip Import Settings** Inspector window. In this panel, you can configure the audio clip’s settings for various platforms. To access the settings for a specific platform, select the tab for that platform.

![Screenshot of the default platform-specific settings section. Tabs for different platforms, such as iOS, Android, and WebGL, are visible along the top.](../uploads/Main/AudioClipImporter-platform.png)

Screenshot of the default platform-specific settings section. Tabs for different platforms, such as iOS, Android, and WebGL, are visible along the top.

| **Property** | **Function** |
| --- | --- |
| **Load Type** | Choose the method Unity uses to load audio assets at runtime:  * **Decompress On Load** - Decompress audio files as soon as they’re loaded. Use this option for smaller compressed sounds to avoid the performance overhead of decompressing during gameplay. Be aware that decompressing Vorbis-encoded sounds on load will use about ten times more memory than keeping them compressed (for ADPCM encoding it’s about 3.5 times), so don’t use this option for large files. * **Compressed In Memory** - Keep audio compressed in memory and decompress while playing. This option has a slight performance overhead, especially for Ogg/Vorbis compressed files. Use it only for files that consume excess memory on **Decompress on Load**. The decompression happens on the mixer thread, which you can monitor in the DSP CPU section in the Audio module of the Profiler window. * **Streaming** - Decode continuous audio. This method uses a minimal amount of memory to buffer compressed data that’s incrementally read from the disk and decoded spontaneously. The decompression happens on a separate streaming thread whose CPU usage you can monitor in the Streaming CPU section in the Audio module of the profiler window. **Note**: Streaming clips have an overhead of approximately 200KB, even without loaded audio data. |
| **Compression Format** | Choose the format for the sound to use at runtime. Note that the options available depend on the currently selected build target:  * **PCM** - Choose this option for short sound effects, and for higher quality audio at the expense of larger file sizes. PCM is lightweight on the CPU requirements because the sound is uncompressed and can just be read from memory. * **ADPCM** - Use this format for sounds that contain a lot of noise and play in large quantities, such as footsteps, impacts, and weapons. The compression ratio is 3.5 times smaller than PCM, but CPU usage is much lower than the MP3/Vorbis formats which makes it a better choice for these types of sounds. * **Vorbis/MP3** - Choose this compression to create smaller files but lower quality audio compared to PCM audio. Use the **Quality** slider to configure the amount of compression and discard less audible information. This format is best for medium length sound effects and music. |
| **Sample Rate Setting** | Control the sample rate of the audio, which affects audio file size and quality. PCM and ADPCM compression formats allow automatically optimized or manual sample rate reduction. The following options are available:  * **Preserve Sample Rate** - Choose this setting to keep the sample rate unmodified (default). * **Optimize Sample Rate** - Optimize the sample rate automatically according to the highest frequency content analyzed. * **Override Sample Rate** - Override the sample rate manually. If you lower the sample rate, it can reduce file size because it removes some of the sound’s frequency details. |
| **Force To Mono** | Enable to down-mix the audio clip to a single channel sound. The down-mixing process typically results in signals that are more quiet than the original. After the down-mixing, the signal is peak-normalized. The peak-normalized signal provides space for later adjustments through the volume property of [AudioSource](class-AudioSource.md). |
| **Load In Background** | Enable this setting to load the audio clip in the background, which prevents stalls on the main thread. This setting is disabled by default to ensure the standard Unity behavior, where all AudioClips are fully loaded when the **scene** starts to play. Unity defers play requests on AudioClips that are still loading in the background until the clip loads fully. You can use the `AudioClip.loadState` property to query the load state. |
| **Preload Audio Data** | Enable to preload the audio clip after the scene fully loads. This setting is enabled by default to reflect standard Unity behavior where all AudioClips complete loading as soon as the scene starts to play. If this flag isn’t set, the audio data will either load on the first [AudioSource.Play](../ScriptReference/AudioSource.Play.md), [AudioSource.PlayOneShot](../ScriptReference/AudioSource.PlayOneShot.md). It can also load through [AudioClip.LoadAudioData](../ScriptReference/AudioClip.LoadAudioData.md), and unload again through [AudioClip.UnloadAudioData](../ScriptReference/AudioClip.UnloadAudioData.md). |
| **Quality** | Determine the amount of compression to apply to a compressed clip. Doesn’t apply to PCM/ADPCM/HEVAG formats. Start with high-quality compression and gradually reduce the setting until you notice the sound quality drops. Then, increase it again slightly until the perceived loss of quality disappears. The Inspector shows statistics about the file size. **Note**: The original size relates to the original file, so if your file is an MP3 file and you set **Compression Format** to **PCM** (uncompressed), the resulting Ratio will be bigger than 100% because the file is now stored uncompressed and takes up more space than the source MP3 that it came from. |

## Preview window

The Preview window lets you play the audio in the Unity Editor outside of Play mode. This window allows you to check audio files to ensure they sound as expected before you use them in your application.

| **Property** | **Function** |
| --- | --- |
| Image of the Play button in Preview Window. | Select the **Play** button to play the selected audio clip. |
| Image of the Turn Auto Play On/Off toggle in the Preview Window. | Enable the **Turn Auto Play On/Off** toggle to play the audio clips as soon as you select any clip in your project. |
| Image of the Turn Loop On/Off toggle in the Preview Window. | Enable this toggle to play the audio clips in a continuous loop. |

**Note**: If Unity Audio is disabled in **Project Settings**, preview is unavailable.

## Additional resources

* [Introduction to audio files and Audio Clips](AudioFiles-introduction.md)
* [Import audio files into Unity](AudioFiles-import.md)
* [Audio file type compatibility](AudioFiles-compatibility.md)
* [Audio file compression in Unity](AudioFiles-compression.md)