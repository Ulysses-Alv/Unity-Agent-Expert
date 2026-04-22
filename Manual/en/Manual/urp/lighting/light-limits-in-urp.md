# Light limits in URP

If you use the default Forward [rendering path](../rendering-paths-introduction-urp.md), the Universal **Render Pipeline** (URP) has a maximum of 9 lights for each **GameObject**:

* 1 Main Light, which is a per-pixel light by default. This is the main Directional Light in your **scene**, or the brightest light.
* 8 Additional Lights, which are also per-pixel lights by default. You can set them as per-vertex lights instead.

For more information about per-pixel and per-vertex lights, refer to [Per-pixel and per-vertex lights](../../PerPixelLights.md).

There are also the following limits per **camera**:

* Desktop and console platforms: 256 Additional Lights
* Mobile platforms: 32 Additional Lights.
* OpenGL ES 3.0 and earlier: 16 Additional Lights.

The Main Light is always visible per camera.

To add more lights per object or per camera, use the Forward+ or Deferred rendering paths instead. For more information, refer to [Choose a rendering path in URP](../rendering-paths-comparison.md)

## Additional resources

* [Lighting settings in the URP Asset Inspector window](../universalrp-asset.md#lighting)