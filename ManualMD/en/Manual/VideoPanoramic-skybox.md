# Set up a panoramic video as a skybox

Use your panoramic video as the backdrop of your **Scene**.

This information covers how to create a **skybox** material and assign it to your default skybox. For more information about how to work with skyboxes, refer to [Create a skybox](skyboxes-using.md).

You can use a panoramic video as a skybox to enhance immersion in various ways. You can use it to create dynamic backgrounds, such as weather systems, day-night cycles, or other visual changes that add depth and realism to the experience.

In **virtual reality** (VR), panoramic video skyboxes provide a fully dynamic, 360° environment that users can explore. As the user moves their head, their view of the video’s content shifts, which can make their surroundings feel more realistic.

To use a panoramic video as a skybox, complete the following tasks:

1. Check the [prerequisites before you continue](#prerequisites).
2. [Identify your video content type](#identify-your-video-content-type).
3. If your video has a equirectangular layout, [Create a panoramic skybox material for your video](#create-a-panoramic-skybox-material-for-your-video).
4. If your video has a **cubemap** layout, you have 2 options:
   * [Create a panoramic skybox material for your video](#create-a-panoramic-skybox-material-for-your-video).
   * [Create a cubemap skybox material for your cubemap video](#create-a-cubemap-skybox-material-for-your-video).
5. [Assign your skybox material to your skybox](#assign-your-skybox-material-to-your-skybox).

For more information about how to create a skybox, refer to [Create a skybox](skyboxes-using.md).

## Prerequisites

Before you continue to set up a panoramic video as a skybox, you need to set up a panoramic video. For more information, refer to [Set up a panoramic video](VideoPanoramic-setup.md).

## Identify your video content type

You need to correctly identify the type of content in the video (cubemap or equirectangular) for the panoramic video to display correctly. To check the type, you can either check the video’s resolution or view the video itself.

### Check the video’s resolution

You can check the video resolution to confirm what content type the video has using the following steps:

1. Click on the video.
2. In the **Inspector**, click on the video’s name.
3. Select **Source Info**.
4. Check the ****Pixels**** value.
   * Equirectangular videos have 2:1 or 1:1 ratios.
   * Cubemap videos have 1:6, 3:4, 4:3, or 6:1 ratios.

### View the video content

Click on the video to bring up a preview in the Inspector window. If the video looks like a flattened cube, it’s a cubemap. If the video looks more like a flattened sphere, it’s likely equirectangular.

## Create a panoramic skybox material for your video

If you want your panoramic video with an equirectangular layout to be the backdrop of your Scene, you need to replace the default skybox with your video content. To do this:

1. Create a new Material (menu: **Assets** > **Create** > **Material**).
2. Select the new Material. This opens the Inspector window for the Material.
3. Set **Shader** to **Skybox** > **Panoramic**.
4. In the **Texture** slot, assign your **render texture**.
5. Set **Mapping** to:
   * **Latitude Longitude Layout** for equirectangular videos.
   * **6 Frames Layout** for cubemap videos.
6. For **Image Type**:
   * Choose **360 Degrees** if the video covers a full 360° view.
   * Choose **180 Degrees** if the video only faces one way.

To check your content looks correct, view the **Preview** at the bottom of the **Material** Inspector.

For more information about the panoramic skybox **shader**, refer to [Panoramic skybox](shader-skybox-panoramic.md).

## Create a cubemap skybox material for your video

The cubemap layout (such as a cross and strip layout) is common for static skybox textures. To set up a skybox material for a cubemap video:

1. Create a new Material (menu: **Assets** > **Create** > **Material**).
2. Select the new Material. This opens the Inspector window for the Material.
3. Set **Shader** to **Skybox** > **Cubemap**.
4. In the **Texture** slot, assign your render texture.

For more information about the cubemap skybox shader, refer to [Cubemap Skybox Shader Material Inspector Window reference](shader-skybox-cubemap.md).

## Assign your skybox material to your skybox

After you create your skybox panoramic or cubemap material, you can assign it to your skybox.

1. From the main menu, select **Window** > **Rendering** > **Lighting**.
2. In the **Lighting** window, select the **Environment** tab.
3. Drag your skybox material to the **Skybox Material** property.

For more information about how to create a skybox, refer to [Create a skybox](skyboxes-using.md).

## Additional resources

* [Panoramic Skybox Shader Material Inspector window reference](shader-skybox-panoramic.md)
* [Panoramic video](VideoPanoramic.md)
* [Introduction to panoramic videos](VideoPanoramic-introduction.md)
* [Set up a panoramic video](VideoPanoramic-setup.md)
* [Set up your 3D panoramic video](VideoPanoramic-3D.md)