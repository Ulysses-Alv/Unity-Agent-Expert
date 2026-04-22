# Configure a Light Probe Proxy Volume in the Built-In Render Pipeline

The area in which the 3D grid of interpolated **Light Probes** are generated is affected by the **Bounding Box Mode** property.

The main difference between **Automatic Local** and **Automatic World** is that in **Automatic Local**, the bounding box is more resource-intensive to compute when a large hierarchy of GameObjects uses the same LPPV component from a parent **GameObject**. However, the resulting bounding box may be smaller in size, meaning the lighting data is more compact.

The number of interpolated Light Probes from within the **bounding volume** is affected by the **Proxy Volume Resolution** property. There are two options:

* **Automatic** (default value) - The resolution on each axis is computed using the number of interpolated Light Probes per unit area that you specify, and the size of the bounding box.
* **Custom** - Allows you to specify a different resolution on each axis (see below).

**Note:** The final resolution on each axis must be a power of two, and the maximum value of the resolution is 32.

**Probe Position Mode** specifies the relative position of an interpolated Light Probe to a cell center. This option may be useful in situations when some of the interpolated Light Probes pass through walls or other geometries and cause light leaking. The example below shows the difference between **Cell Corner** and **Cell Center** in a 2D view, using a 4x4 grid resolution:

![On the left, Probe Position Mode is set to Cell Corner, and each Light Probe is at the intersection of four cells. On the right, Probe Position Mode is set to Cell Center, and each Light Probe is at the center of a cell instead.](../uploads/Main/LightProbeProxyVolumeWindow4.png)

On the left, **Probe Position Mode** is set to **Cell Corner**, and each Light Probe is at the intersection of four cells. On the right, **Probe Position Mode** is set to **Cell Center**, and each Light Probe is at the center of a cell instead.

**Data Format** specifies which format the associated 3D texture uses. There are two options:

* **Float** (default value) - The texture uses the 32-bit floating-point channel format to store the spherical harmonics coefficients.
* **Half Float** - The texture uses the half-float (16-bit float) channel format to store the spherical harmonics coefficients. The advantage of this format is that half-float linear texture sampling is supported by the majority of devices and the precision difference between this format and the 32-bit floating point channel format is not noticeable. Also, the texture sampling performance on the GPU is better with this data format.

## Additional resources

* [Light Probe Proxy Volume component reference](class-LightProbeProxyVolume-reference.md)