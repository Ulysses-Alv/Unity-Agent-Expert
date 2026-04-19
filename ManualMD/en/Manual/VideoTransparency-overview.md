# Introduction to video transparency support in Unity

Alpha transparency allows parts of a video or image sequence to appear semi-transparent or fully transparent.

Transparent videos are useful if you want to create overlays and visual effects because you can layer videos within your Unity **scenes**.

This page provides an overview of alpha transparency, explains the different types of transparency in videos, and describes how Unity handles and transcodes alpha values for use with transparent videos.

## Alpha values

An alpha value determines the level of transparency of a video or a **pixel** within a video.

Alpha values can take on any value within the range 0 to 1 or 0 to 255. This range allows for smooth and gradual changes in transparency.

| **Value** | **Transparency level** |
| --- | --- |
| Lowest (0) | Image or pixel is fully transparent. |
| Highest (1 or 255) | Image or pixel is fully opaque. |
| Intermediate values | Image or pixel is partially transparent. You can view both the image and the background behind it simultaneously. |

Whether an alpha value controls the entire video’s image or a pixel depends on the value’s type: global or per-pixel.

### Global alpha value transparency

The Video Player component supports a global alpha value when it plays content in a **camera**’s near or far planes. When you change this value, every pixel in the video changes transparency. If you set the value to `0.0f`, the entire video becomes completely transparent.

For instructions on how to change global alpha value transparency, refer to [Change your entire video’s transparency (global alpha)](VideoTransparency-instructions.md#global-alpha).

### Per-pixel alpha value transparency

Videos can have per-pixel alpha values, which means that transparency can vary across the video image. For example, the background of a video is transparent while an object in the foreground is opaque.

For instructions on how to import a per-pixel alpha video, refer to [Change the transparency of pixels in your video (per-pixel alpha)](VideoTransparency-instructions.md#import-per-pixel).

For transparent formats Unity supports, refer to [Supported codecs for per-pixel alpha transparency](VideoTransparency-codecs.md#per-pixel).

## How Unity transcodes videos with alpha values

When you use transparent videos in Unity, in certain cases you might need to [transcode](video-transcoding.md) your video to ensure it supports the use of alpha values. You need to transcode your video if:

* Your video has the VP8 codec but the platform you want to target doesn’t support VP8’s built-in transparency.
* You want to use the H.264 codec, which doesn’t have built-in support for alpha values.

When you transcode a video with alpha values (for example, a WebM video), Unity packs the alpha values into the color stream. This process allows you to use transparent videos with both H.264 and VP8.

During this process, Unity doubles the image width of the video. For example, if you have a 1280 x 720 video with alpha values, internally this becomes a 2560 x 720 video.

After Unity doubles the video width, the resulting file contains two sections:

* The left half stores the original (without alpha) image.
* The right half stores the alpha as a monochrome image.

Both halves are then repackaged by shaders into RGBA, which produces the final texture in the Video Player. Afterwards, the video output can display transparency. For a list of these shaders (which are automatically included in your builds), go to **Project Settings** > **Graphics** > **Shader Settings** > **Video**.

**Note**: This process lowers the maximum resolution a video can have. H.264 has a typical max resolution of 4096 x 2304. If you import a video with transparency that has this resolution, it becomes 8192 x 2304, which fails because most encoders don’t support this size.

VP8 has a max resolution of 16383 x 16383 so is more likely to accept these higher resolutions. But VP8 is often implemented using a software-only path, which is slower and might result in low performance at these higher resolutions.

## Additional resources

* [Supported codecs for transparent videos (per-pixel alpha)](VideoTransparency-codecs.md)
* [Set up transparent videos in Unity](VideoTransparency-instructions.md)
* [Transcoding video files](video-transcoding.md)
* [Video Clip](class-VideoClip.md)
* [Video Player](class-VideoPlayer.md)