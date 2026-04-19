# Lens Distortion Volume Override reference for URP

![Scene with Lens Distortion effect turned off.](../../uploads/urp/post-proc/lens-distortion-off.png)

Scene with Lens Distortion effect turned off.

![Scene with Lens Distortion effect turned on.](../../uploads/urp/post-proc/lens-distortion.png)

Scene with Lens Distortion effect turned on.

The **Lens Distortion** effect distorts the final rendered picture to simulate the shape of a real-world **camera** lens.

## Properties

| **Property** | **Description** |
| --- | --- |
| **Intensity** | Use the slider to set the overall strength of the **distortion effect**. |
| **X Multiplier** | Use the slider to set the distortion intensity on the x-axis. This value acts as a multiplier so you can set this value to 0 to disable distortion on this axis, |
| **Y Multiplier** | Use the slider to set the distortion intensity on the y-axis. This value acts as a multiplier so you can set this value to 0 to disable distortion on this axis, |
| **Center** | Set the center point of the distortion effect on the screen. |
| **Scale** | Use the slider to set the value for global screen scaling. This zooms the render to hide the borders of the screen. When you use a high distortion, **pixels** on the borders of the screen can break because they rely on information from pixels outside the screen boundaries that don’t exist. This property is useful for hiding these broken pixels around the screen border. |