# Set a GameObject to use a Light Probe Proxy Volume in the Built-In Render Pipeline

Most of the Renderer components in Unity contain Light Probes. There are three options for Light Probes:

* **Off**: the Renderer doesn’t use any interpolated Light Probes.
* **Blend Probes** (default value): the Renderer uses one interpolated Light Probe.
* **Use Proxy Volume**: the Renderer uses a 3D grid of interpolated Light Probes.

When you set the **Light Probes** property in the ****Mesh** Renderer** component to **Use Proxy Volume**, the **GameObject** must have a **Light Probe Proxy Volume** (LPPV) component attached. You can add a LPPV component on the same GameObject, or you can use (borrow) a LPPV component from another GameObject using the **Proxy Volume Override** property. If Unity cannot find a LPPV component in the current GameObject or in the Proxy Volume Override GameObject, a warning message is displayed at the bottom of the Renderer.

## Example

![An example of a simple Mesh Renderer using a Light Probe Proxy Volume component](../uploads/Main/LightProbeProxyVolumeExample.png)

An example of a simple Mesh Renderer using a Light Probe Proxy Volume component

In the **Scene** above, there are two planes on the floor using Materials that emit a lot of light. Note that:

* The **ambient light** changes across the geometry when using a LPPV component. Use one interpolated Light Probe to create a constant color on each side of the geometry.
* The geometry doesn’t use static **lightmaps**, and the spheres represent interpolated Light Probes. They are part of the **Gizmo Renderer**.