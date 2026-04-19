# Surface Shader input structure reference for the Built-In Render Pipeline

The input structure `Input` generally has any texture coordinates needed by the **shader**. Texture coordinates must be named “`uv`” followed by texture name (or start it with “`uv2`” to use second texture coordinate set).

Additional values that can be put into Input structure:

* `float3 viewDir` - contains view direction, for computing Parallax effects, rim lighting etc.
* `float4` with `COLOR` semantic - contains interpolated per-vertex color.
* `float4 screenPos` - contains screen space position for reflection or screenspace effects. Note that this is not suitable for [GrabPass](SL-GrabPass.md); you need to compute custom UV yourself using `ComputeGrabScreenPos` function.
* `float3 worldPos` - contains world space position.
* `float3 worldRefl` - contains world reflection vector *if **surface shader** does not write to o.Normal*. See Reflect-Diffuse shader for example.
* `float3 worldNormal` - contains world normal vector *if surface shader does not write to o.Normal*.
* `float3 worldRefl; INTERNAL_DATA` - contains world reflection vector *if surface shader writes to o.Normal*. To get the reflection vector based on per-pixel **normal map**, use `WorldReflectionVector (IN, o.Normal)`. See Reflect-Bumped shader for example.
* `float3 worldNormal; INTERNAL_DATA` - contains world normal vector *if surface shader writes to o.Normal*. To get the normal vector based on per-pixel normal map, use `WorldNormalVector (IN, o.Normal)`.