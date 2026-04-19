# Sort sprites

Unity’s Graphics settings (menu: **Edit** > **Project Settings** > **Graphics**) provide a setting called **Transparency Sort Mode**, which you can use to control how to sort sprites depending on their position in relation to the Camera. More specifically, it uses the sprite’s position on an axis to determine which ones are transparent compared to others.

An example of when you might use this setting is to sort **sprites** along the y-axis. This is quite common in 2D games, where sprites that are higher up the y-axis are sorted behind sprites that are lower, to make them appear further away.

**Notes:**

* If you set the **Transparency Sort Mode** to **Custom Axis**, you also need to set the **Transparency Sort Axis**.
* If the **Transparency Sort Mode** is set to **Custom Axis**, renderers in the **Scene** view are sorted based on the distance of this axis from the **camera**. Use a value between –1 and 1 to define the axis. For example: X=0, Y=1, Z=0 sets the axis direction to up. X=1, Y=1, Z=0 sets the axis to a diagonal direction between X and Y.

![Example: Set the Transparency Sort Mode to Custom Axis, and set the Y value for the Transparency Sort Axis to a value higher than 0.](../../../uploads/Main/AxisDistanceSort2.png)

Example: Set the **Transparency Sort Mode** to **Custom Axis**, and set the **Y** value for the **Transparency Sort Axis** to a value higher than 0.