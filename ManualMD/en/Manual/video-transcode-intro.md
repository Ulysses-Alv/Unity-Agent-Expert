# Introduction to transcoding video files

Video transcoding converts video files between different formats. The [Video Clip Importer](class-VideoClip.md) can transcode video files that you import into Unity. This is useful when your video files use a [codec that the Editor platform supports](VideoSources-FileCompatibility.md), but isn’t [compatible with your target platforms](video-sources-compatibility-target-platforms.md).

For instructions on how to transcode your video clips, refer to the [Transcode your video files](video-transcode-steps.md).

## Video codecs

You can transcode video files into one of the following video codecs:

* [H.264](https://en.wikipedia.org/wiki/H.264/MPEG-4_AVC)
* [H.265](https://en.wikipedia.org/wiki/High_Efficiency_Video_Coding)
* [VP8](https://en.wikipedia.org/wiki/VP8)

Transcoded video files use the appropriate audio codec automatically:

* [AAC](https://en.wikipedia.org/wiki/Advanced_Audio_Coding) for videos encoded with H.264 or H.265
* [Vorbis](https://en.wikipedia.org/wiki/Vorbis) for videos encoded with VP8

**Note**: The [Video Clip Importer](class-VideoClip.md) provides only basic transcoding options. Depending on the encoding of your source files, you might get optimal performance with an external transcoding program.

## Additional resources

* [Video file compatibility with the Unity Editor](VideoSources-FileCompatibility.md)
* [Video file compatibility with target platforms](video-sources-compatibility-target-platforms.md)
* [Video encoding compatibility](video-encoding-compatibility.md)
* [Understand video files](VideoSources-VideoFiles.md)