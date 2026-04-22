# Supported codecs for transparent videos (per-pixel alpha)

Choose compatible video codecs to ensure per-pixel alpha transparency in your Unity applications.

Certain video codecs support the use of transparency on individual **pixels** of a video (per-pixel alpha).

This information covers the transparent video codecs that Unity supports. For information about transparency types and how to set up per-pixel alpha videos, refer to [Introduction to video transparency support in Unity](VideoTransparency-overview.md) and [Set up transparent videos in Unity](VideoTransparency-instructions.md).

## Supported per-pixel alpha codecs

Unity supports two codecs that have per-pixel alpha:

* [Apple ProRes 4444](#apple-prores-4444)
* [WebM with VP8](#webm-with-vp8)

**Note**: If you’re not sure what codec your video uses, your operating system’s video playback software might have the functionality to identify it.

### Apple ProRes 4444

The [Apple ProRes 4444](https://support.apple.com/en-gb/HT202410) codec is a high-quality version of Apple ProRes for 4:4:4:4 image sources, including alpha channels.

Some key information about this codec:

* Provides the same level of visual fidelity as the source video.
* Most Apple platforms have support for it.
* Typically produces large files, which increases storage and bandwidth requirements.

When you import a video that uses this codec, enable both the **Transcode** and **Keep Alpha** options in the [Video Clip Importer](class-VideoClip.md).

If your video has the Apple ProRes codec but you don’t transcode the file, your [target platform](video-sources-compatibility-target-platforms.md) might not support the video. For more information about the process, refer to [How Unity transcodes videos with alpha values](VideoTransparency-overview.md#transcode-alpha).

### WebM with VP8

A WebM file that uses the VP8 video codec stores alpha information natively. Most Unity-supported platforms can read transparent videos in this format.

Unity usually only supports the VP8 video codec and Vorbis audio codec for WebM videos.

If your VP8 video has another audio codec, the audio might not be audible, and Unity doesn’t support the import of VP9 videos. However, if your WebM file uses the VP9 video or Opus audio codecs, there are some exceptions for certain platforms. For more information, refer to [WebM videos - Supported video and audio codecs](VideoSources-FileCompatibility.md#supported-video-audio).

**Note:** Most of Unity’s supported platforms use a software implementation to decode WebM files. Therefore, you don’t need to transcode the files for these platforms. However, [Android](android.md)’s built-in VP8 support doesn’t include transparency support, so you must enable [transcoding](video-transcoding.md) to ensure Unity uses its internal alpha representation.

## Additional resources

* [Introduction to video transparency support in Unity](VideoTransparency-overview.md)
* [Set up transparent videos in Unity](VideoTransparency-instructions.md)