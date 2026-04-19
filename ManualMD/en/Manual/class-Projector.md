# Projector component reference for the Built-In Render Pipeline

[Switch to Scripting](../ScriptReference/Projector.md "Go to Projector page in the Scripting Reference")

Explore the properties in the Projector component to configure a material’s projection onto a surface.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Near Clip Plane** | Objects in front of the near clip plane will not be projected upon. |
| **Far Clip Plane** | Objects beyond this distance will not be projected upon. |
| **Field Of View** | The field of view in degrees. This is only used if the Projector is not Orthographic. |
| **Aspect Ratio** | This allows you to tune the height and width of the Projector. |
| **Orthographic** | If enabled, the Projector will be Orthographic instead of perspective. |
| **Orthographic Size** | The Orthographic size of the Projection. This is only used if Orthographic is enabled. |
| **Material** | The Material that will be projected. |
| **Ignore Layers** | Objects in Layers that you specify here will not be projected upon. The default value is None. |

Projector