# Develop an ambisonic audio decoder

An ambisonic decoder is an audio component that decodes the ambisonic audio format into a specific output format, such as stereo or surround sound. This format depends on your speaker configuration (menu: **Edit** > **Project Settings** > **Audio** > **Default Speaker Mode**), unless your platform overrides this.

Although Unity supports ambisonic audio, it doesn’t provide built-in decoders by default. Instead, you must either select a third-party decoder, or use your own decoder **plug-in**.

## Set up an ambisonic audio decoder plug-in

You can set up an ambisonic audio decoder in the same way as you’d set up an [Audio Spatializer](AudioSpatializerSDK.md). However, the following parameters in the `AudioPluginInterface.h` file are specific to ambisonic audio decoder plug-ins:

* The [`UnityAudioEffectDefinitionFlags_IsAmbisonicDecoder`](#unityaudioeffectdefinitionflags_isambisonicdecoder) effect definition flag
* The [`UnityAudioAmbisonicData`](#unityaudioambisonicdata-struct) data struct

### UnityAudioEffectDefinitionFlags\_IsAmbisonicDecoder

During the plug-in scanning phase, the `UnityAudioEffectDefinitionFlags_IsAmbisonicDecoder` flag notifies Unity that this is an ambisonic decoder effect.
To enable a plug-in to operate as an ambisonic decoder, set a flag in the description bit-field of the effect:

```
definition.flags |= UnityAudioEffectDefinitionFlags_IsAmbisonicDecoder;
```

Unity lists your plug-in as an option in the Project Settings window (menu: **Edit** > **Project Settings** > **Audio** > **Ambisonic Decoder Plugin**).

### UnityAudioAmbisonicData struct

The `UnityAudioAmbisonicData` struct is similar to the [`UnityAudioSpatializerData`](AudioSpatializerSDK.md#SpatializerMetaData) struct that Unity passes into spatializers, and contains an `ambisonicOutChannels` integer.

#### The ambisonicOutChannels integer

The Ambisonic decoders run early in the audio pipeline in Unity, and the `ambisonicOutChannels` variable tells the plug-in how many of the output channels Unity needs to use. `ambisonicOutChannels` is automatically set to the `DefaultSpeakerMode`’s channel count.

For example, if you play back a first order ambisonic **audio clip** that has 4 channels, and your speaker mode is stereo (which has only 2 channels):

* An ambisonic decoder’s process callback passes in 4 for the in and out channel count.
* The `ambisonicOutChannels` field is automatically set to 2.
* The plug-in outputs its spatialized data to the first 2 channels and zero out the other 2 channels.

## Steps to develop an ambisonic audio decoder plug-in

Follow these steps to develop your own ambisonic audio decoder plug-in for Unity:

1. Create a custom audio plug-in using the [Native audio plug-in SDK](AudioMixerNativeAudioPlugin.md).
2. Set a flag in the description bit-field of the effect:

   ```
   definition.flags |= UnityAudioEffectDefinitionFlags_IsAmbisonicDecoder;
   ```
3. When you are done configuring your plug-in, compile your file. Make sure it is compilable on your preferred platforms.
4. Optionally, convert your file to a .dll file.
5. Move your plug-in file into your Unity project’s Asset folder.

For more information on how to work with your ambisonic audio plug-in in Unity, see [Ambisonic Audio](AmbisonicAudio.md).

## Formats that Unity ambisonics support

The Unity ambisonic sources framework can support first order ambisonics. The plug-in interface includes information to support binaural stereo and **quad** output, but the plug-in itself determines which outputs are supported.

Initially, ambisonic decoder plug-ins support first order ambisonic sources and binaural stereo output. There is no support for second order ambisonics.

There is nothing in the framework that’s specific to any of the different ambisonic formats available. If the clip’s format matches the ambisonic decoder plug-in’s expected format, then ambisonic audio should work without issue. Unity’s preferred ambisonic format is B-format, with ACN component ordering, and SN3D normalization.

## Further information

For information on how to develop a plug-in, refer to [Native audio plug-in SDK](AudioMixerNativeAudioPlugin.md) and [Audio spatializer SDK](AudioSpatializerSDK.md). You must also download the [Audio plug-in SDK](https://github.com/Unity-Technologies/NativeAudioPlugins).