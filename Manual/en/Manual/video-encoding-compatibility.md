# Video encoding compatibility reference

Video encoding compresses video files using a codec to save bandwidth and storage.

## Codecs

| **Codec** | **Advantages** | **Disadvantages** |
| --- | --- | --- |
| **H.264** | Best natively supported codec for hardware acceleration. |  |
| **H.265** | Higher **compression**. | Limited to [supported devices](#H265Compatibility). |
| **VP8** | Cross-platform support and comprehensive **feature set**. | Consumes more resources compared to H.264. |

The Unity Editor supports the .ogv format, but it isn’t widely supported on other platforms. Transcode .ogv files into .mp4 (H.264) or .webm (VP8) depending on the target platform.

**Note:** Android supports VP8 using native libraries, so VP8 might be hardware-assisted on some Android devices.

## Encoding with external programs

If you use video files that the target platform supports, you can manage encoding with an external program for finer control. You can disable the **Transcode** property in the [Video Clip Importer](class-VideoClip.md) so Unity doesn’t modify the video files.

## Key encoding values

The following table lists key values in your encoding parameters:

| **Parameter** | **Description** |
| --- | --- |
| **Video Codec** | [H.264](https://en.wikipedia.org/wiki/H.264/MPEG-4_AVC), [H.265](https://en.wikipedia.org/wiki/High_Efficiency_Video_Coding), or [VP8](https://en.wikipedia.org/wiki/VP8). |
| **Resolution** | The resolution of your display. For example: 1280 × 720. |
| **Profile** | The profile is a set of capabilities and constraints, often specified by the vendor, such as Baseline or Main. Applies to H.264/H.265. See [H.264](https://en.wikipedia.org/wiki/H.264/MPEG-4_AVC#Profiles) or [H.265](https://en.wikipedia.org/wiki/High_Efficiency_Video_Coding#Profiles). |
| **Profile Level** | Applies for H.264/H.265. Within a given profile, the level specifies performance requirements, for example, Baseline 3.1. |
| **Audio Codec** | [AAC](https://en.wikipedia.org/wiki/Advanced_Audio_Coding) (for .mp4 videos using H.264/H.265) or [Vorbis](https://en.wikipedia.org/wiki/Vorbis) (for .webm videos using VP8). |
| **Audio Channels** | Dependent on the platform. Refer to the developer guide for your platform. For example, the article on [Supported media formats](https://developer.android.com/guide/topics/media/media-formats) for Android. |

## H.265 compatibility

The following table provides platform-specific requirements and information for the H.265 codec:

| **Platform** | **Requirements** | **Encoding/Decoding** | **Notes** |
| --- | --- | --- | --- |
| **macOS** | SDK 10.13+ | **Hardware encoding:** 6th Generation Intel Core processor   **Software encoding:** All Macs    **Hardware decoding:** 6th Generation Intel Core processor   **Software decoding:** All Macs |  |
| **Windows** | Windows 10 + HEVC extensions | [Encoder](https://docs.microsoft.com/en-us/windows/desktop/medfound/h-265---hevc-video-encoder)  [Decoder](https://docs.microsoft.com/en-us/windows/desktop/medfound/h-265---hevc-video-decoder) | [HEVC extension (Hardware only)](https://www.microsoft.com/en-ca/p/hevc-video-extensions-from-device-manufacturer/9n4wgh0z6vhq)  [HEVC extension (Hardware + software support)](https://www.microsoft.com/en-ca/p/hevc-video-extensions/9nmzlz57r3t7/0010) |
| **UWP** | [Windows 10+](https://docs.microsoft.com/en-us/windows/uwp/audio-video-camera/supported-codecs) |  | If a device lists support for H.265, that might not apply to all devices within the device family. |
| **Android** | [5.0+](https://developer.android.com/guide/topics/media/media-formats) |  |  |
| **iOS** | SDK 11.0+ | **Hardware decoding:** A9 Chip   **Software decoding:** All iOS Devices |  |
| **tvOS** | SDK 11.0+ |  |  |

## Additional resources

* [Understand video files](VideoSources-VideoFiles.md)
* [Video transcoding](video-transcoding.md)
* [Video file compatibility with the Unity Editor](VideoSources-FileCompatibility.md)