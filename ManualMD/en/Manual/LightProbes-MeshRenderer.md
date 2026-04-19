# Set a GameObject to use light from Light Probes

To use **Light Probes** on your moving GameObjects, the **Mesh Renderer** component on the moving **GameObject** must be set correctly. The **Mesh** Renderer component has a **Light Probes** setting which is set to **Blend Probes** by default. This means that by default, all GameObjects will use light probes and will blend between the nearest probes as it changes position in your **scene**.

You can change this setting to either “off” or “use proxy volume”. Switching the light probes setting to off will disable the light probe’s effect on this GameObject.

Light Probe Proxy Volumes are a special setting which you can use for situations where a **large moving object** might be too big to be sensibly lit by the results of a single tetrahedron from the **light probe group**, and instead needs to be lit by multiple groups of light probes across the length of the model. See **Light Probe Proxy Volumes** for more information.

The other setting in the Mesh Renderer **inspector** which relates to light probes is the **Anchor Override** setting. As described previously, when a GameObject moves through your scene, Unity calculates which tetrahedron the GameObject falls within from the volume defined by the light probe groups. By default this is calculated from the centre point of the Mesh’s bounding box, however you can override the point that is used by assigning a different GameObject to the **Anchor Override** field.

If you assign a different GameObject to this field, it is up to you to move that GameObject in a way that suits the lighting you want on your mesh.

The anchor override may be useful when a GameObject contains two separate adjoining meshes; if both meshes are lit individually according to their bounding box positions then the lighting will be discontinuous at the place where they join. This can be prevented by using the same Transform (for example the parent or a child object) as the interpolation point for both Mesh Renderers or by using a Light Probe Proxy Volume.