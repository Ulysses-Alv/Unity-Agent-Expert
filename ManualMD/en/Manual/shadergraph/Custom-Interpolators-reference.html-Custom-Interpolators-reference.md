# Custom Interpolator reference | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Custom-Interpolators-reference.html)

# Custom Interpolator reference

Transfer custom data from the vertex stage to the fragment stage.

You first create a custom interpolator block node in the Vertex context. You can then add a custom interpolator node in the workspace and connect it to a block node in the Fragment context.

The following descriptions and settings apply to both the Vertex block and the Fragment node.

## Settings

| Property | Description |
| --- | --- |
| **Name** | Sets the unique name of the custom interpolator to identify and reference it in the graph. |
| **Type** | Sets the number of channels the Custom Interpolator exposes. The default value is **Vector 4**, which exposes x, y, z, and w channels. |
| **Interpolation** | Selects how Unity interpolates the value from vertex to fragment across the surface. The following options are available:  * **Linear**: Applies the default linear interpolation, which preserves correct rates of change in screen space. * **No Perspective**: Doesn't correct perspective, which can warp data, depending on the angle between the surface and the camera. * **No Interpolation**: Doesn't interpolate the data, which creates hard edges between triangles. |

## Additional resources

* [Built-in blocks](Built-In-Blocks.html)
* [Add a custom interpolator](Custom-Interpolators.html)