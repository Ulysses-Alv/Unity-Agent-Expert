# Create a cubemap

## Creating Cubemaps from Textures

The fastest way to create **cubemaps** is to import them from specially laid out [Textures](class-TextureImporter.md).
Select the Texture in the **Project window**, to see the Import Settings in the **Inspector** window. In the Import Settings, set the **Texture Type** to **Default**, **Normal Map** or **Single Channel**, and the **Texture Shape** to **Cube**. Unity then automatically sets the Texture up as a Cubemap.

Several commonly-used cubemap layouts are supported (and in most cases, Unity detects them automatically).

Vertical and horizontal cross layouts, as well as column and row of cubemap faces are supported:

![A vertical cross layout where four faces form a column, a horizontal cross layout where four faces form a row, a column layout, and a row layout.](../uploads/Textures/CubeLayout6Faces1.png)

A vertical cross layout where four faces form a column, a horizontal cross layout where four faces form a row, a column layout, and a row layout.

Another common layout is `LatLong` (Latitude-Longitude, sometimes called cylindrical). Panorama images are
often in this layout:

![A latitude-longitude layout. Four faces form a row, and the two remaining faces stretch along the top and bottom of the texture.](../uploads/Textures/CubeLayoutLatLong.png)

A latitude-longitude layout. Four faces form a row, and the two remaining faces stretch along the top and bottom of the texture.

`SphereMap` (spherical environment map) images can also be found:

![A spherical environment map. The texture is circular, with one face in the centre and four faces surrounding it in a ring. The sixth face is the outer shell of the texture.](../uploads/Textures/CubeLayoutSphereMap.png)

A spherical environment map. The texture is circular, with one face in the centre and four faces surrounding it in a ring. The sixth face is the outer shell of the texture.

By default Unity looks at the **aspect ratio** of the imported texture to determine the most appopriate layout from
the above. When imported, a cubemap is produced which can be used for skyboxes and reflections:

Selecting `Glossy Reflection` option is useful for cubemap textures that will be used by
[Reflection Probes](class-ReflectionProbe.md). It processes cubemap mipmap levels in a special way
(specular convolution) that can be used to simulate reflections from surfaces of different smoothness:

![A cubemap used in a Reflection Probe to create reflections on surfaces of varying smoothnesses.](../uploads/Textures/CubeOptionGlossyReflections.png)

A cubemap used in a Reflection Probe to create reflections on surfaces of varying smoothnesses.

## Other Techniques

Another useful technique is to generate the cubemap from the contents of a Unity **scene** using a script.
The [Camera.RenderToCubemap](../ScriptReference/Camera.RenderToCubemap.md) function can record the six face
images from any desired position in the scene; the code example on the function’s script reference page
adds a menu command to make this task easy.

## Legacy Cubemap Assets

Unity also supports creating cubemaps out of six separate [textures](class-TextureImporter.md).
Select **Assets** > **Create** > **Rendering** > **Legacy Cubemap** from the menu,
and drag six textures into empty slots in the inspector.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Right..Back Slots** | Textures for the corresponding cubemap face. |
| **Face Size** | Width and Height of each Cubemap face in **pixels**. The textures will be scaled automatically to fit this size. |
| **Mipmap** | Should mipmaps be created? |
| **Linear** | Should the cubemap use linear color? |
| **Readable** | Should the cubemap allow **scripts** to access the pixel data? |

Note that it is preferred to create cubemaps using the Cubemap texture import type (see above) - this
way cubemap texture data can be compressed; edge fixups and glossy reflection convolution be performed;
and **HDR** cubemaps are supported.