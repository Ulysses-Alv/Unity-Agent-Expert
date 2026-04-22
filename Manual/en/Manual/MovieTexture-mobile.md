# Play video on mobile platforms

## iOS

Movie Textures are not supported on iOS. Instead, full-screen streaming playback is provided using [Handheld.PlayFullScreenMovie](../ScriptReference/Handheld.PlayFullScreenMovie.md).

Keep your videos inside the **StreamingAssets** folder located in the **Assets** folder of your project.

Unity iOS supports any movie file types that play correctly on an iOS device, implying files with the extensions **.mov**, **.mp4**, **.mpv**, and **.3gp** and using one of the following **compression** standards:

* H.264 Baseline Profile Level 3.0 video
* MPEG–4 Part 2 video

For more information about supported compression standards, consult the iPhone SDK [MPMoviePlayerController Class Reference](http://developer.apple.com/library/ios/#documentation/MediaPlayer/Reference/MPMoviePlayerController_Class/MPMoviePlayerController/MPMoviePlayerController.html).

As soon as you call [Handheld.PlayFullScreenMovie](../ScriptReference/Handheld.PlayFullScreenMovie.md) the screen fades from your current content to the designated background color. It might take some time before the movie is ready to play. In the meantime, the player continues displaying the background color and may also display a progress indicator to let the user know the movie is loading. When playback finishes, the screen fades back to your content.

### The video player does not respect switching to mute while playing videos

Unity plays video files using Apple’s embedded player (as of SDK 3.2 and iPhone OS 3.1.2 and earlier). This contains a bug that prevents Unity from switching to mute.

### The video player does not respect the device’s orientation

The Apple video player and iPhone SDK do not provide a way to adjust the orientation of the video. To fix this, you can manually create two copies of each movie in landscape and portrait orientations. Then, the orientation of the device can be determined before playback so the right version of the movie can be chosen.

## Android

Movie Textures are not supported on Android. Instead, full-screen streaming playback is provided using [Handheld.PlayFullScreenMovie](../ScriptReference/Handheld.PlayFullScreenMovie.md).

Keep your videos inside the **StreamingAssets** folder located in the **Assets** folder of your project.

Unity Android supports any movie file type supported by Android, (ie, files with the extensions **.mp4** and **.3gp**) and using one of the following compression standards:

* H.263
* H.264 AVC
* MPEG–4 SP

However, device vendors are keen on expanding this list, so some Android devices are able to play formats other than those listed, such as HD videos.

For more information about the supported compression standards, consult the Android SDK [Core Media Formats documentation](http://developer.android.com/guide/appendix/media-formats.html).

As soon as you call [Handheld.PlayFullScreenMovie](../ScriptReference/Handheld.PlayFullScreenMovie.md) the screen fades from your current content to the designated background color. It might take some time before the movie is ready to play. In the meantime, the player continues displaying the background color and may also display a progress indicator to let the user know the movie is loading. When playback finishes, the screen fades back to your content.