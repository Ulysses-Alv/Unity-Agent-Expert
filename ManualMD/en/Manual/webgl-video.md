# Video playback in Web

Unity Web supports video playback using the [VideoPlayer](../ScriptReference/Video.VideoPlayer.md) API. This page provides information only about the video playback capabilities that Web supports. To learn how to use video with your Unity application, refer to the [Video Player](class-VideoPlayer.md).

## Video player

You can use the Video Player component to control how to time video playback relative to other interactive behaviors in Web. For more information, refer to the [Clock management with the Video Player component](video-clock.md).

The only exceptions are:

* Web doesn’t support frame accuracy.
* The `VideoPlayer` component doesn’t support synchronous playback with `captureFramerate`. By default, it uses the normal asynchronous playback that’s described with the Game time update mode.
* The `VideoPlayer` component corrects drift between video playback and Unity time by temporarily speeding the playback controls to up or down. However, because the video support in Safari browser has limitations that prevent this mechanism from operating with precision, the drift correction is disabled.

## Supported video playback features and formats

Unity Web supports the following video playback audio output modes:

| **Class** | **Use** |
| --- | --- |
| **[VideoAudioOutputMode.None](../ScriptReference/Video.VideoAudioOutputMode.None.md)** | Disables the embedded audio. |
| **[VideoAudioOutputMode.Direct](../ScriptReference/Video.VideoAudioOutputMode.Direct.md)** | Sends the embedded audio directly to the platform’s audio hardware. |
| **[VideoAudioOutputMode.AudioSource](../ScriptReference/Video.VideoAudioOutputMode.AudioSource.md)** | Sends the embedded audio into a specified [AudioSource](../ScriptReference/AudioSource.md). If you set the output mode to `VideoAudioOutputMode.AudioSource`, Unity ignores all AudioSource fields except mute. This is because 3D spatialization of video playback isn’t available on the web. |

### Video formats

Unity supports the following common video file formats:

| ***Format*** | ***Extensions*** |
| --- | --- |
| MPEG–4 Part 14 | .mp4 |
| MPEG–4 file used for video downloaded from the Apple iTunes store | .m4v |
| Apple’s QuickTime movie format | .mov |
| Moving Picture Experts Group (MPEG) | .mpg |
| MPEG video | .mpeg |
| WebM video | .webm |
| Ogg video file | .ogv |

The only exception to this restriction is if the video URL has no file name extension, in which case, the browser plays the video without any restrictions.

### Video clips

[`VideoClips`](class-VideoClip.md) aren’t supported on Web. Typically, when creating a **scene**, you import a [VideoClip](class-VideoClip.md) to your Unity project using `VideoClipImporter`, which is convenient if you want to reuse the same `VideoClip` across several platforms. When building a Web game that has `VideoClip` attached however, the Unity console logs the following warning for each `VideoClip` found in the game:

`Embedded video clips are not supported by the Web player: %s. \nUse the Video Player component's URL option instead`.
Where `%s` is replaced by the video clip name. At runtime, if your game has `VideoClips` assigned, then Unity logs a warning message in the developer console of your web browser.

## Additional resources:

* [Video Player](VideoPlayer.md)
* [Video Clip](class-VideoClip.md)
* [Video Player component](class-VideoPlayer.md)