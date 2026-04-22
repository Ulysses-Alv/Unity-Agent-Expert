# Video Player component reference

Configure the Video Player component’s settings to specify the video file to play and customize its playback behavior within your **scene**.

For information about the Video Player component and how to create one, refer to [Video Player component targets](VideoPlayer-intro.md) and [Create a Video Player component](VideoPlayer-instructions.md).

| **Property** | **Description** |
| --- | --- |
| **Source** | Choose the source type to retrieve your video content from. The source could be a video file that’s stored in your project, hosted externally at a specific URL, or saved in your `StreamingAssets` folder. |
| **Video Clip** | Select the [Video Clip](class-VideoClip.md) to assign to the Video Player component. This isn’t supported on the Web platform. |
| **URL** | Enter the URL (for example, http:// or file://) of the video you want to assign to the Video Player. |
| **Update Mode** | Set the clock source that the Video Player component uses to update its timing. For more information, refer to [Update Mode dropdown](#update-mode-dropdown). |
| ****Play On Awake**** | Play the video when the Scene launches. Clear it if you want to trigger the video playback at another point during runtime. Trigger it via scripting with the `Play()` command. |
| **Wait For First Frame** | Wait for the first frame of the source video to be ready for display before playback starts. Clear it to keep the video time in sync with the rest of the game, which might cause the first few frames to be discarded. |
| **Loop** | Loop the source video when it reaches its end. Clear it to stop playing the video when it reaches the end. |
| **Skip On Drop** | When you enable this option, and the Video Player component detects drift between the playback position and the game clock, the Video Player skips ahead. When you disable this option, the Video Player doesn’t correct for drift and systematically plays all frames. |
| **Playback Speed** | Set a multiplier for the playback speed, as a value between 0 and 10. By default, this value is set to 1, which is the normal speed. If you set the field to 2, the video plays at two times its normal speed. |
| **Render Mode** | Choose how the video will render. For more information, refer to [Render Mode dropdown](#render-mode-dropdown). |
| **Camera** | Define the [Camera](class-Camera.md) that receives the video. This property is available only when **Render Mode** is **Camera Far Plane** or **Camera Near Plane**. | |
| **Alpha** | Set the global transparency level to add to the source video. This property allows elements behind the plane to be visible through it. For more information about alpha channels, refer to [video transparency support](VideoTransparency.md). This property is available only when **Render Mode** is **Camera Far Plane** or **Camera Near Plane**. |
| **3D Layout** | Choose the layout of 3D content in the source video. For more information, refer to [3D Layout dropdown](#3d-layout-dropdown). |
| **Target Texture** | Define the Render Texture where the Video Player component renders its images. For recommended settings for the Render Texture, refer to [Prepare your Render Texture for video](VideoPlayer-rendertexture.md#prepare). |
| ****Aspect Ratio**** | Set the aspect ratio of the images that fill the **Camera Near Plane**, **Camera Far Plane**, or **Render Texture** when the corresponding **Render Mode** is used. For more information, refer to [Aspect Ratio dropdown](#aspect-ratio-dropdown). |
| **Renderer** | Select the [Renderer](../ScriptReference/Renderer.md) where the Video Player component renders its images. When set to **None**, the Video Player component uses the **Renderer** on the same **GameObject**. |
| **Auto-Select Property** | When you enable this option, the Video Player component selects the [Renderer](../ScriptReference/Renderer.md)’s main texture automatically. When you disable this option, you can set the **Material Property** option manually. |
| **Audio Output Mode** | Define how the source’s audio tracks are output. For more information, refer to [Audio Output Mode dropdown](#audio-output-mode-dropdown). |
| **Controlled Tracks** | The number of audio tracks in the video.  Only shown when **Source** is **URL**. When **Source** is **Video Clip**, the number of tracks is determined by examining the video file. |
| **Track Number** | Enable the associated audio track to use for playback. You need to set this property prior to playback.  The text to the left of the checkbox provides information about the audio track, specifically the track number, language, and number of channels.  When the source is a URL, this information is only available during playback.  This property only appears if your source is a Video Clip that has an audio track (or tracks), or your source is a URL (which lets you indicate how many tracks to expect from the URL during playback). |
| **Audio Source** | The [audio source](class-AudioSource.md) through which the audio track plays. The targeted audio source can also play Audio Clips.  The audio source’s playback controls (`Play On Awake` and `Play()` in scripting API) don’t apply to the video source’s audio track.  This property only appears when the **Audio Output Mode** is set to **Audio Source**. |
| **Mute** | Mute the associated audio track. In **Audio Source** mode, the audio source’s control is used.  This property is available only when **Audio Output Mode** is **Direct**. |
| **Volume** | Volume of the associated audio track. In **Audio Source** mode, the audio source’s volume is used.  This property is available only when **Audio Output Mode** is **Direct**. |

### Update Mode dropdown

The following options are available in the **Update Mode** dropdown.

| **Value** | **Description** |
| --- | --- |
| **DSP Time** | Use the same clock source that processes audio. |
| **Game Time** | Use the same clock source as the game clock. This clock source is affected by the [time scaling and capture frame rate settings](managing-time-and-frame-rate.md). |
| **Unscaled Game Time** | Use the same clock source as the game clock but this clock source isn’t affected by [time scaling or capture frame rate](managing-time-and-frame-rate.md). |

### Render Mode dropdown

The following options are available in the **Render Mode** dropdown. For further information about these targets, refer to [Video Player component targets](VideoPlayer-intro.md).

| **Value** | **Description** |
| --- | --- |
| **Camera Far Plane** | Render the video on the Camera’s [far plane](../ScriptReference/Video.VideoRenderMode.CameraFarPlane.md). If you choose this option, the video plays in the background of the Scene. |
| **Camera Near Plane** | Render the video on the Camera’s [near plane](../ScriptReference/Video.VideoRenderMode.CameraNearPlane.md). If you choose this option, the video plays in front of the objects in your Scene. |
| **Render Texture** | Render the video into a [Render Texture](class-RenderTexture.md). For instructions and recommended settings, refer to  [Set up your Render Texture to display video](VideoPlayer-rendertexture.md). |
| **Material Override** | Render the video into a selected Texture property of a GameObject through its Renderer’s [Material](../ScriptReference/Material.md). |
| **API Only** | Render the video into the [VideoPlayer.texture](../ScriptReference/Video.VideoPlayer-texture.md) Scripting API property. |

### 3D Layout dropdown

The following options are available in the **3D Layout** dropdown.

| **Value** | **Description** |
| --- | --- |
| **None** | Video doesn’t have any 3D content. |
| **Side by Side** | Video has 3D content where the left eye occupies the left half and right eye occupies the right half of video frames. |
| **Over Under** | Video has 3D content where the left eye occupies the upper half and right eye occupies the lower half of video frames. |

### Aspect Ratio dropdown

The following settings are options available in the **Aspect Ratio** dropdown.

| **Value** | **Description** |
| --- | --- |
| **No Scaling** | Use no scaling. The video is centered on the destination rectangle. |
| **Fit Vertically** | Scale the source to fit the destination rectangle vertically. This option crops the left and right sides or leaves black areas on each side if necessary. The source aspect ratio is preserved. |
| **Fit Horizontally** | Scale the source to fit the destination rectangle horizontally. This option crops the top and bottom regions or leaves black areas above and below if needed. The source aspect ratio is preserved. |
| **Fit Inside** | Scale the source to fit the destination rectangle without having to crop. Leaves black areas on the left and right or above and below as needed. The source aspect ratio is preserved. |
| **Fit Outside** | Scale the source to fit the destination rectangle without leaving black areas on the left and right or above and below, cropping as required. The source aspect ratio is preserved. |
| **Stretch** | Scale both horizontally or vertically to fit the destination rectangle. The source aspect ratio isn’t preserved. |

### Audio Output Mode dropdown

The following options are available in the **Audio Output Mode** dropdown.

| **Value** | **Description** |
| --- | --- |
| **None** | Audio isn’t played. |
| **Audio Source** | Audio samples are sent to selected [audio sources](class-AudioSource.md), enabling Unity’s audio processing to be applied. |
| **Direct** | Audio samples are sent directly to the audio output hardware, bypassing Unity’s audio processing. |
| **API Only (Experimental)** | Audio samples are sent to the associated [AudioSampleProvider](../ScriptReference/Experimental.Audio.AudioSampleProvider.md). |

## Additional resources

* [VideoPlayer API](../ScriptReference/Video.VideoPlayer.md)
* [Video Player](VideoPlayer.md)
* [Video Player component targets](VideoPlayer-intro.md)
* [Video Player component reference](class-VideoPlayer.md)
* [Set up your Render Texture to display video](VideoPlayer-rendertexture.md)
* [Create a Video Player component](VideoPlayer-instructions.md)