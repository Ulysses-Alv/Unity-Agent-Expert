# Use video sources

Reference your video sources in the Video Player to use video in Unity.

To use video in Unity, you must reference your files in the Video Player through the **Source** dropdown. The Video Player can play video sources from video clips or URLs.

This information covers referencing video files only. To configure the Video Player, refer to [Video Player component](class-VideoPlayer.md). To configure video clips, refer to [Video Clip Importer](class-VideoClip.md).

## Prerequisites

* Add a [Video Player component](class-VideoPlayer.md) to your **GameObject**.

## Reference your file as a video clip

To reference a file as a video clip in the Video Player:

1. [Import your file as a video clip](video-clips-use.md#import).
2. Select your **Video Player** component.
3. In the **Video Player **Inspector**** window, go to **Source**.
4. Set **Source** to **Video Clip**.
5. Click and drag the video clip asset into the **Video Clip** parameter or click the **Video Clip** picker to select your video clip.

**Note:** As video files are often large, you can also assign video clips as an [addressable video asset](https://docs.unity3d.com/Packages/com.unity.addressables@2.2/manual/index.html) or from [AssetBundles](AssetBundlesIntro.md) to help reduce your initial game install size.

## Reference your file as a URL

Reference your file as a URL to play files that aren’t bundled with your application. This can be helpful for user-generated content, if your content isn’t under Unity’s direct control, or if you want to avoid storing large video files locally.

URLs can point to files on a local file system, a web server, or your StreamingAssets folder. As the URL option bypasses asset management, you must manually ensure that Unity is able to locate the source video. For example, a local file must be in a file location that Unity can access, indicated with scripting, while a web URL needs a web server to host the source video.

To reference a file as a URL in the Video Player:

1. Select your **Video Player** component.
2. In the **Video Player Inspector** window, go to **Source**.
3. Set **Source** to **URL**.
4. Set the URL as your chosen URL.

**Note:** On the [Web](webgl.md) platform, the URL must point to a web URL because playback from the local file system and `Application.persistentDataPath` aren’t supported.

### File system

On native build platforms, you can set the URL to any file path to directly use files in your file system. You can use a `file://` prefix but it’s not necessary.

**Note:** Some web browsers allow you to manually disable browser Cross-Origin Resource Sharing (CORS) security for `file://` URL access for local development and testing purposes. For security reasons, this isn’t a recommended approach.

### Web server

You can set the URL to read videos from a web server using the `http://` and `https://` prefixes. In these cases, Unity performs the necessary pre-buffering and error management processes.

### StreamingAssets folder

You can set the URL to use files placed in Unity’s [StreamingAssets](StreamingAssets.md) folder, or by using the platform-specific path `Application.streamingAssetsPath`. Refer to [Application.streamingAssetsPath](../ScriptReference/Application-streamingAssetsPath.md) for more information.

## Additional resources

* [Video Player component](class-VideoPlayer.md)
* [Video Clip Importer](class-VideoClip.md)
* [Asset Bundles](AssetBundlesIntro.md)
* [StreamingAssets](StreamingAssets.md)