# Set up your Render Texture to display video

Use a **Render Texture** to display your video content on the surface of multiple **GameObjects** that share the Render Texture or to apply **post-processing** effects to the video.

To display your video content on a Render Texture:

1. [Prepare your Render Texture for video](#prepare).
2. [Set a Render Texture as your Video Player’s target](#target).

## Prepare your Render Texture for video

To optimize your Render Texture for use with video content:

1. Click on your Render Texture to open the **Inspector** window.
2. Make sure the resolution of the Render Texture matches the resolution of your video content.
3. Set **Anti-aliasing** to **None**.
4. Set **Color Format** to **R8G8B8A8\_UNORM**. If you notice excessive banding with your Render Texture, particularly in dark areas where the color transitions look unnatural, set **Color Format** to **R16G16B16A16\_UNORM**.
5. Set **Depth Stencil Format** to **None**.
6. Disable **Mipmap**.
7. Set **Wrap Mode** to **Clamp**.
8. Set **Filter Mode** to **Point**.
9. Disable anisotropic filtering via script: `renderTexture.anisoLevel = 0;`

When done, your Render Texture is ready for your video content.

## Set a Render Texture as your Video Player’s target

To have your video play on your Render Texture, you need to set up your Video Player component to target your Render Texture:

1. Click on the GameObject that contains the Video Player component. The Inspector window opens.
2. Set **Render Mode** to **Render Texture**.
3. Assign your Render Texture to **Target Texture**.

Your Video Player component now targets your optimized Render Texture.

## Additional resources

* [VideoPlayer API](../ScriptReference/Video.VideoPlayer.md)
* [Video Player](VideoPlayer.md)
* [Video Player component targets](VideoPlayer-intro.md)
* [Video Player component reference](class-VideoPlayer.md)
* [Create a Video Player component](VideoPlayer-instructions.md)