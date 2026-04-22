# Introduction to panoramic videos

A panoramic video is a video that shows up to 360° views of a captured **scene**. Unity’s panoramic video feature enables you to:

* Include real-world 360° videos.
* Reduce scene complexity in **VR** by including a pre-rendered backdrop video instead of real geometry.

To use the panoramic video features in the Unity Editor, you must have access to panoramic video clips, or know how to author them.

## Panoramic video layouts

Unity supports 180° and 360° videos in either an [equirectangular](https://en.wikipedia.org/wiki/Equirectangular_projection) layout (longitude and latitude) or a [cubemap](https://docs.unity3d.com/2017.2/Documentation/Manual/class-Cubemap.html) layout (6 frames).

### Equirectangular 2D videos

Equirectangular 2D videos have an **aspect ratio** of exactly 2:1 for 360 content, or 1:1 for 180 content.

This layout flattens a spherical image onto a 2D plane, which makes it possible to display 360° content in a video.

The horizontal axis represents the full 360° field of view around the viewer. The vertical axis represents the 180° field of view from directly above to directly below the viewer.

![Equirectangular 2D video](../uploads/Main/Equirectangular2D.jpg)

Equirectangular 2D video

### Cubemap 2D videos

Cubemap 2D videos have an aspect ratio of 1:6, 3:4, 4:3, or 6:1, depending on face layout.

A cubemap disassembles a cube’s faces, lays them flat, and maps the 360° video content onto the faces. This projection enables the display of 360° content in a video.

When the engine rebuilds the cubemap into a cube, it recreates the scene as a seamless panoramic view.

![Cubemap 2D video](../uploads/Main/Cubemap2D.jpg)

Cubemap 2D video

## 3D panoramic videos

3D panoramic videos are videos that combine up to 360° panoramic views with stereoscopic 3D content to simulate depth. You can use 3D panoramic videos to create immersive environments in **virtual reality** (VR) with less overhead than actual geometry.

To create 3D videos, specialized **cameras** capture two slightly different perspectives of a scene (one for each eye). During **post-processing**, software formats these perspectives into layouts like Side-by-Side or Over-Under.

| **Layout** | **Description** |
| --- | --- |
| **Side-by-Side** | The left eye’s content is placed on the left side of the frame, and the right eye’s content is placed on the right side. |
| **Over-Under** | The content is arranged vertically, with the left eye’s content positioned on top and the right eye’s content on the bottom. |

When you set up your 3D panoramic videos in Unity, Unity detects which eye is currently being rendered and uses [Single Pass Stereo rendering](https://docs.unity3d.com/Manual/SinglePassStereoRendering.html) to send this information to the **Skybox** or Panoramic **shader**. The shader contains the logic to select the correct half of the video based on this information when Unity renders each eye’s content in VR.

For setup information, refer to [Set up your 3D panoramic video](VideoPanoramic-3D.md).

## Limitations and optimizations for panoramic videos

When you work with panoramic video in Unity, you need to be aware of some restrictions and optimization requirements to ensure smooth playback.

### Optimize your visuals

The panoramic video feature is resource-intensive. For the best visual results use panoramic videos in the highest possible resolution. For example, 4K or 8K resolution.

### System limitations

Large videos require more computing power and resources for decoding. Most systems have specific limits on maximum video decoding resolutions. For example, many desktop hardware video decoders have a limit of 4K resolutions. Mobile hardware video decoders are often limited to 2K or less which limits the resolution of playback in real-time on those platforms.

### 3D content

3D content requires double either the width or the height of the video. This corresponds to **Side-by-Side** or **Over-Under** layouts.

You can transcode your video to produce lower resolution versions of panoramic videos. However, be careful not to introduce bleeding at the edge between the left and right 3D content, or, between cube map faces and adjacent black areas. In general, author videos in power-of-two dimensions and transcode to other power-of-two dimensions to reduce visual artifacts.

## Additional resources

* [Panoramic video](VideoPanoramic.md)
* [Set up a panoramic video](VideoPanoramic-setup.md)
* [Set up a panoramic video as a skybox](VideoPanoramic-skybox.md)
* [Set up your 3D panoramic video](VideoPanoramic-3D.md)