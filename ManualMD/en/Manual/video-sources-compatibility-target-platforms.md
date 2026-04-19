# Video file compatibility with target platforms

The [Video Player component](class-VideoPlayer.md) uses the native audio and video decoding libraries of your Unity Editor platform to play [video files in the Editor](VideoSources-FileCompatibility.md). You need to check that the files meet the system requirements for the target build platform.

## Recommendations for codec support

Follow the vendor recommendations for your platform for codec support:

* **Windows:** [Supported Media Formats](https://docs.microsoft.com/en-us/windows/desktop/medfound/supported-media-formats-in-media-foundation), [H.265](https://docs.microsoft.com/en-us/windows/desktop/medfound/h-265---hevc-video-decoder#format-constraints)
* **UWP:** [Supported Codecs](https://docs.microsoft.com/en-us/windows/uwp/audio-video-camera/supported-codecs)
* **Android:** [Supported Media Formats](https://developer.android.com/guide/topics/media/media-formats.html#recommendations)
* **iOS:** [Compare iPhone Models](http://www.apple.com/iphone/compare/)

**Note:** On older mobile platforms, codec choices are limited. You might need to inspect and convert or [re-encode](video-encoding-compatibility.md) videos that you intend to include in an application running on multiple devices.

## Additional resources

* [Understand video files](VideoSources-VideoFiles.md)
* [Video file compatibility with the Unity Editor](VideoSources-FileCompatibility.md)
* [Video encoding compatibility](video-encoding-compatibility.md)
* [Transcoding video files](video-transcoding.md)