# Understand video files

Unity stores imported video files as [video clip assets](video-clips-use.md). Video files are often referred to as containers because they can hold not just the video itself but also additional tracks like audio, subtitles, and additional video footage. Some common types of tracks that might be present in a video container include:

* Multiple points of view
* Stereo or 5.1 versions of the audio mix
* Subtitles in different languages
* Dialog in different languages

File extensions such as .mp4, .mov, .webm or .avi indicate that the data in the video file is arranged using a certain container format. For more information on common file extensions, refer to [Video file compatibility with the Unity Editor](VideoSources-FileCompatibility.md#FileFormats).

## Encoding

To save bandwidth and storage, each track’s content is [encoded](video-encoding-compatibility.md) using a [codec](video-encoding-compatibility.md#codecs), which compresses and decompresses data as required.

## Decoding

Unity provides both software and hardware decoding.

Most modern devices have hardware dedicated to decoding videos. This hardware typically requires less power than other resources like the CPU, so these resources can be used for tasks other than decoding videos. This hardware acceleration uses native custom APIs, which differ between platforms. Unity’s video architecture hides these differences by providing a common UI and [scripting API](../ScriptReference/UnityEngine.VideoModule.md) to access these capabilities.

Unity is also capable of software-based video decoding. This uses the VP8 video codec and Vorbis audio codec, and is useful for situations where a platform’s hardware decoding results in unwanted restrictions with resolution, multiple audio tracks, or [alpha transparency support](VideoTransparency.md).

## Additional resources

* [Video file compatibility with the Unity Editor](VideoSources-FileCompatibility.md)
* [Video file compatibility with target platforms](video-sources-compatibility-target-platforms.md)
* [Video encoding compatibility](video-encoding-compatibility.md)
* [Transcoding video files](video-transcoding.md)