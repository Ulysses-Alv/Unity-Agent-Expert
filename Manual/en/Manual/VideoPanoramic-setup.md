# Set up a panoramic video

Use a panoramic video to create a more immersive environment in your application.

To set up a panoramic video for use in Unity, follow these steps:

1. [Create a Video Player for your panoramic video](#create-a-video-player-for-your-panoramic-video).
2. If your video has an equirectangular layout, [render an equirectangular video to a render texture](#render-an-equirectangular-video-to-a-render-texture).
3. If your video has a **cubemap** layout, [render a cubemap video to a render texture](#render-a-cubemap-video-to-a-render-texture).
4. [Project your video content onto an object](#project-your-video-content-onto-an-object).

## Create a Video Player for your panoramic video

To create a Video Player component from your panoramic video:

1. Import your video into Unity as an [asset](ImportingAssets.md).
2. Drag the video asset from the Project view to an empty area of Unity’s Hierarchy view.

By default, this sets up the component to play the video full-screen for the default **Camera**. Press **Play** to view this.

## Render an equirectangular video to a render texture

To render an equirectangular video to a **render texture**, you must change the default behavior of the video player. To do this, use the following steps:

1. Go to **Assets** > **Create** > **Rendering** > **Render Texture**.
2. Click on the new render texture. The **Inspector** window shows.
3. Set **Size** to match your video’s dimensions.
   * To check the dimensions of your video, refer to [View the source information of your video clip](video-clips-use.md#source).
4. Click on the object that contains the **Video Player** component. The Inspector window shows.
5. Set **Render Mode** to **Render Texture**.
6. Set **Target Texture** to your new render texture.
7. If your video has the cubemap layout, consider these [Alternative steps for videos with the cubemap layout](#alternative).
8. Enter Play mode to verify that this is functioning correctly.

The video doesn’t render in the **Game** window, but you can select the render texture asset to check its content updates with each video frame.

**Note**: Sometimes the asset preview fails to refresh automatically and doesn’t display the video’s content. To fix this you might need to force a UI update, for example move your mouse or click on a button.

## Render a cubemap video to a render texture

If your video has a cubemap layout, you can render the video directly to a render texture cube instead of a 2D render texture. This is better for performance and is more suited to game engines.

To set this up:

1. Go to **Assets** > **Create** > **Rendering** > **Render Texture**. This action creates a new render texture.
2. Click on the new render texture. The Inspector window shows.
3. Set **Dimension** to **Cube**.
4. Set **Size** to be exactly the dimensions of the individual faces of the source video.
   * For example, if you have a 4 x 3 horizontal cross cubemap layout video with dimensions 4096 x 3072, set **Size** to 1024 x 1024 (4096 / 4 and 3072 / 3).
5. Click on the object that contains the **Video Player** component. The Inspector window shows.
6. Set **Render Mode** to **Render Texture**.
7. Set **Target Texture** to your new render texture.

With this setup, the video player assumes that the source video contains a cube map in either a cross or a strip layout (which it uses the video aspect ratio to determine). The video player then fills out the render texture’s faces with the correct cube faces.

## Project your video content onto an object

Once your video player contains your panoramic video as a render texture, you can display your video’s content on the surface of a **GameObject** or **skybox**.

To display the content on a GameObject, drag your render texture onto the object. This creates a new material that contains the render texture as its texture. When the Video Player plays the video, it will display on the object.

For information on how to set a panoramic video as a skybox, refer to [Set up a panoramic video as a skybox](VideoPanoramic-skybox.md).

## Additional resources

* [Panoramic Skybox Shader Material Inspector window reference](shader-skybox-panoramic.md)
* [Panoramic video](VideoPanoramic.md)
* [Introduction to panoramic videos](VideoPanoramic-introduction.md)
* [Set up your 3D panoramic video](VideoPanoramic-3D.md)
* [Set up a panoramic video as a skybox](VideoPanoramic-skybox.md)