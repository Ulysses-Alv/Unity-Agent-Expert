# Pixel Perfect Camera component reference for URP

Explore the different properties available to customize the appearance of your 2D **pixel** art with the **Pixel Perfect **Camera****.

![Property table](../../uploads/urp/2D/2D_Pix_image_7.png)
The component’s **Inspector** window

| **Property** | **Function** |
| --- | --- |
| **Asset Pixels Per Unit** | This is the amount of pixels that make up one unit of the **scene**. Match this value to the **Pixels Per Unit** values of all **sprites** in the scene. |
| **Reference Resolution** | This is the original resolution your assets are designed for. |
| **Crop Frame** | Select what to do when there is a difference in **aspect ratio**. |
| **Grid Snapping** | Select what to do when snapping. |
| **Filter Mode** (Only available when **Stretch Fill** option is selected.) | Select the method Unity uses to upscale the final image. |
| **Current Pixel Ratio** | Shows the size ratio of the rendered sprites compared to their original size. |

### Reference Resolution

This is the original resolution your Assets are designed for. Scaling up scenes and assets from this resolution preserves your pixel art cleanly at higher resolutions.

### Grid Snapping options

#### Upscale Render Texture

By default, the scene is rendered at the pixel perfect resolution closest to the full screen resolution.

Enable this option to have the scene rendered to a temporary texture set as close as possible to the **Reference Resolution**, while maintaining the full screen aspect ratio. This temporary texture is then upscaled to fit the entire screen.

![Box examples](../../uploads/urp/2D/2D_Pix_image_8.png)

Box examples

The result is unaliased and unrotated pixels, which may be a desirable visual style for certain game projects.

#### Pixel Snapping

Enable this feature to snap **sprite Renderers** to a grid in world space at render-time. The grid size is based on the **Assets Pixels Per Unit** value.

**Pixel Snapping** prevents subpixel movement and make sprites appear to move in pixel-by-pixel increments. This does not affect any **GameObjects**’ Transform positions.

### Crop Frame

| **Option** | **Description** |
| --- | --- |
| **None** | Doesn’t crop the frame. |
| **Pillarbox** | Crops the **viewport** vertically, using black bars on the left and right to match the **Reference Resolution**. |
| **Letterbox** | Crops the viewport horizontally, using black bars at the top and bottom to match the **Reference Resolution**. |
| **Windowbox** | Crops the viewport both vertically and horizontally, using black bars on all sides to match the **Reference Resolution**. |
| **Stretch fill** | Stretches the viewport to fill the screen while maintaining the aspect ratio of the viewport. |

![An image of a 2D character, with Crop Frame set to None, Pillrabox, Letterbox, Windowbox, and Stretch Fill.](../../uploads/urp/2D/pixel-perfect-crop-options.jpg)

An image of a 2D character, with Crop Frame set to None, Pillrabox, Letterbox, Windowbox, and Stretch Fill.

### Filter Mode

**Filter Mode** is only usable when **Stretch Fill** option is selected.

Defaults to **Retro AA** upscale filtering, where the image is upscaled as close as possible to the screen resolution as a multiple of the **Reference resolution**, followed by a bilinear filtering to upscale to the target screen resolution.

**Point** filtering is also available for user preference. If you upscale the image this way, it can suffer from bad pixel placement, thus losing pixel perfectness.

| Uncropped cat | Cropped cat |
| --- | --- |
| Point | Retro AA |
| Uncropped cat | Cropped cat |
| Upscale **Render Texture** + Point | Upscale Render Texture + Retro AA |