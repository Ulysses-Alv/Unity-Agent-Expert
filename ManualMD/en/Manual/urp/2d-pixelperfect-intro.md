# Introduction to the Pixel Perfect Camera in URP

Understand how to use the **Pixel** Perfect **Camera** to keep your pixel art crisp and clear at different resolutions, and stable in motion.

The **2D Pixel Perfect** package contains the **Pixel Perfect Camera** component. The component calculates what Unity needs to scale the **viewport** with resolution changes to maintain the pixel perfect visual style, so that you don’t need to calculate manually. You can use the component settings to adjust the definition of the rendered pixel art within the camera viewport, and preview any changes made in the Game view.

![Pixel Perfect Camera Gizmo](../../uploads/urp/2D/2D_Pix_image_0.png)

Pixel Perfect Camera Gizmo

To begin using the component, attach the **Pixel Perfect Camera** component to the main Camera **GameObject** in the **scene**. The component is represented by two green bounding boxes centered on the **Camera** **Gizmo** in the **Scene view**. The solid green bounding box shows the visible area in Game view, while the dotted bounding box shows the [Reference Resolution](2d-pixelperfect-ref.md#reference-resolution).