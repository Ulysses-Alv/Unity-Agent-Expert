# Ambisonic Audio

Ambisonics are a type of audio that provide a representation of sound that can completely surround a listener. They can provide an audio **skybox** for distant ambient sounds, and are particularly useful for 360-degree videos and applications.

Ambisonics are stored in a multi-channel format. Instead of mapping each channel to a specific speaker, ambisonics represent the soundfield in a more general way. You can rotate the soundfield based on the listener’s orientation (such as the user’s head rotation in VR). You can also decode the soundfield into a format that matches the speaker setup.

## Selecting an ambisonic audio decoder

To select an ambisonic audio decoder in your project, open your project’s [Audio](class-AudioManager.md) settings (menu: **Edit** > **Project Settings** > **Audio**). Select an **Ambisonic Decoder Plugin** from the list of available decoders in the project.

![Ambisonic options in the Audio settings](../uploads/Main/AmbisonicAudioSettings.png)

Ambisonic options in the Audio settings

There are no built-in decoders included with Unity, but you can do one of the following options:

* You can create your own ambisonic audio decoder **plug-in**. For more information, refer to [Ambisonic Audio Decoder](AudioDevelopAmbisonicDecoder.md).
* Use external decoders. For instance, some **VR** hardware manufacturers provide them in their audio SDKs for Unity. Check your target platform manufacturer’s documentation to learn if this is available for your project.

## Importing an ambisonic audio clip

To import an ambisonic **audio clip**, import a multi-channel B-format WAV file with ACN component ordering, and SN3D normalization. In the **Inspector** window for the audio clip, enable **Ambisonic**.

![The Ambisonic option in the audio clip inspector](../uploads/Main/AmbisonicAudioClipInspector.png)

The Ambisonic option in the audio clip inspector

## Playing the ambisonic audio clip through an Audio Source

To play an ambisonic audio clip through an **Audio Source**:

* Assign the WAV file as an **Audio Clip** on an Audio Source.
* Disable the **Spatialize** option. When you play an ambisonic audio clip, it is automatically decoded through the project’s selected ambisonic audio decoder. The decoder converts the clip from ambisonic format to the project’s selected speaker format. It also already handles spatialization as a part of this decoding operation, based on the orientation of the audio source and **audio listener**.

When Unity plays an ambisonic audio clip, it decompresses the file, if necessary, then decodes it to convert it to the project’s selected speaker mode. It then applies the Audio Source’s effects.

**Note:** Reverb zones are disabled for ambisonic audio clips.