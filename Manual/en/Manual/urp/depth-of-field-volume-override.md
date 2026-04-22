# Depth of Field Volume Override in URP

The **Depth Of Field** component applies a depth of field effect, which simulates the focus properties of a **camera** lens. In real life, a camera can only focus sharply on an object at a specific distance. Objects nearer or farther from the camera are out of focus. The blurring gives a visual cue about an object’s distance, and introduces “bokeh”, which refers to visual artifacts that appear around bright areas of the image as they fall out of focus. To read more about bokeh, refer to the [Wikipedia article on Bokeh](https://en.wikipedia.org/wiki/Bokeh).

The Universal **Render Pipeline** (URP) has two depth of field modes:

* **Gaussian**: this mode approximates camera-like effects, but doesn’t imitate them completely. It has a limited blur radius and only does far-field blurring. This mode is the fastest, and is the best mode for lower-end platforms.  
  ![Gaussian Depth Of Field](../../uploads/urp/post-proc/dof-gaussian.png)
* **Bokeh**: a slower but higher quality mode that closely imitates the effects of a real-life camera. It can do both near & far-field blurring, and generates bokeh on areas with high luminosity intensity, also known as hot spots.  
  ![Gaussian Depth Of Field](../../uploads/urp/post-proc/dof-bokeh.png)