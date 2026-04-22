# Video file compatibility with the Unity Editor

You can import different video file formats into Unity. For Unity to preview video files, the files must be compatible with the platform where you run the Unity Editor. Different platforms support different video file formats.

**Note:** Your files must also be [compatible with the target build platforms](video-sources-compatibility-target-platforms.md) for your users. Unity provides options for [transcoding video files](video-transcoding.md) to commonly supported formats, but you can also manage compatibility yourself.

## Video file formats

Each Editor platform supports different video file formats.

| Extension | Windows | macOS | Linux |
| --- | --- | --- | --- |
| .asf | X |  |  |
| .avi | X |  |  |
| .dv | X | X |  |
| .m4v | X | X |  |
| .mov | X | X |  |
| .mp4 | X | X |  |
| .mpg | X | X |  |
| .mpeg | X | X |  |
| .ogv | X | X | X |
| .vp8 | X | X | X |
| .webm | X | X | X |
| .wmv | X |  |  |

### WebM videos - Supported video and audio codecs

Unity usually only supports the import of WebM videos with VP8 (video) and Vorbis (audio) codecs. However, you can use the `StreamingAssets` folder to add WebM videos with the other codecs to the project. When you use this folder, the Editor doesn’t interpret the file itself, instead the Editor imports the raw file into the build. For more information about this folder, refer to [Include additional files in a build](StreamingAssets.md).

You can then use code to assign the raw file to the VideoPlayer component (refer to [VideoPlayer.url](../ScriptReference/Video.VideoPlayer-url.md)). This allows the target platform to read the file directly and bypass the Editor’s support limitations. However, the imported video clip isn’t visible in the Editor and you can’t drag the asset onto a Video Player in the **scene**.

When you use a video file in this way, it unlocks the following target platform capabilities:

| **Platform** | **Supported video codecs** | **Supported audio codecs** |
| --- | --- | --- |
| Android | Supports the video codecs that WebM files support on Android. Refer to the Android documentation on [Supported media formats](https://developer.android.com/media/platform/supported-formats). | Supports the audio codecs that WebM files support on Android. Refer to the Android documentation on [Supported media formats](https://developer.android.com/media/platform/supported-formats). |
| Nintendo Switch | VP8, VP9 | Vorbis |
| Web | VP8, VP9 | Vorbis, Opus |

## Codecs

Encoding must use a supported codec. Each platform supports specific codecs that can change with each version of the platform. Video files that use unsupported codecs trigger an error message in the Editor and you must convert them to a compatible codec before you can use them. For more information about codec support, refer to the official platform documentation for your Editor, such as [Windows codec compatibility](https://docs.microsoft.com/en-us/windows/desktop/medfound/supported-media-formats-in-media-foundation) and [macOS codec compatibility](https://developer.apple.com/library/content/documentation/MacOSX/Conceptual/OSX_Technology_Overview/MediaLayer/MediaLayer.html).

The optimal supported video codec for most platforms is H.264. However, the optimal encoding for Linux is a .webm container with VP8 for video and Vorbis for audio.

## Unsupported files

The easiest way to ensure video compatibility for different platforms is to enable [transcoding](video-transcoding.md) in Unity, which will re-encode the video for compatibility when you make a build. However, this can result in longer build times and lower video quality. To transcode video clips, refer to [Video Clip Importer](class-VideoClip.md#transcoding).

To use video clips that aren’t supported by the Editor platform but are supported by your target platform, refer to [Use unsupported video files in the Editor](video-files-unsupported.md).

## Additional resources

* [Video file compatibility with target platforms](video-sources-compatibility-target-platforms.md)
* [Transcoding video files](video-transcoding.md)
* [Understand video files](VideoSources-VideoFiles.md)
* [Video transparency support](VideoTransparency.md)