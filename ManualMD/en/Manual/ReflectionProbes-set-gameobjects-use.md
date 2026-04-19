# Set GameObjects to use Reflection Probes

To make use of the reflection **cubemap**, an object must have the **Reflection Probes** option enabled on its [Mesh Renderer](class-MeshRenderer.md) and also be using a **shader** that supports reflection probes. When the object passes within the volume set by the probe’s **Size** and **Probe Origin** properties, the probe’s cubemap will be applied to the object.

You can also manually set which reflection probe to use for a particular object using the settings on the object’s [Mesh Renderer](class-MeshRenderer.md). To do this, select one of the options for the **Mesh** Renderer’s **Reflection Probes** property (**Simple**, **Blend Probes** or **Blend Probes and Skybox**) and drag the chosen probe onto its **Anchor Override** property.

See the [Reflection Probes](ReflectionProbes.md) section in the manual for further details about principles and usage.