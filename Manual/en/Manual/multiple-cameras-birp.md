# Set the order of multiple cameras

You can create multiple Cameras and assign each one to a different depth. Use the following properties in the **Camera** **Inspector** window:

* **Depth** if your project uses the Built-In **Render Pipeline**.
* **Priority** if your project uses the Universal Render Pipeline (URP).

Cameras are drawn from low depth to high depth. In other words, a Camera with a depth of 2 will be drawn on top of a Camera with a depth of 1. You can adjust the values of the **View Rect** property to resize and position the Camera’s view onscreen. This can create multiple mini-views like missile cams, map views, rear-view mirrors, etc.