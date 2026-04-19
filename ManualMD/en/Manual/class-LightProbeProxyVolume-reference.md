# Light Probe Proxy Volume component reference for the Built-In Render Pipeline

There are three options available:

| **Bounding Box Mode:** | **Function:** |
| --- | --- |
| **Automatic Local** (default value) | A local-space bounding box is computed. The interpolated Light Probe positions are generated inside this bounding box. If a Renderer component isn’t attached to the **GameObject**, then a default bounding box is generated. The bounding box computation encloses the current Renderer, and sets all the Renderers down the hierarchy that have the **Light Probes** property to **Use Proxy Volume**. |
| **Automatic World** | A bounding box is computed which encloses the current Renderer and all the Renderers down the hierarchy that have the **Light Probes** property set to **Use Proxy Volume**. The bounding box is world-aligned. |
| **Custom** | A custom bounding box is used. The bounding box is specified in the local-space of the GameObject. The bounding box editing tools are available. You can edit the **bounding volume** manually by modifying the **Size** and **Origin** values in the UI (see below). |