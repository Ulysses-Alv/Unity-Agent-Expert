# Introduction to the deferred rendering path

The Deferred **rendering path** in the Universal **Render Pipeline** (URP) first creates a G-buffer, which is a set of textures that stores information about the **scene**, then uses the information to light all the **GameObjects** at once.

## Terrain blending accuracy with the Deferred rendering path

When blending **terrain** in the **Forward rendering** path, Unity uses multi-pass rendering to calculate lighting for four layers at a time, and alpha-blends after each set of four layers. In the Deferred rendering path, Unity combines terrain layers in the G-buffer pass using hardware blending, four layers at a time, then calculates lighting only once during the Deferred rendering pass. The approach in the Deferred rendering path limits how correct the combination of property values is. For example, **pixel** normals cannot be accurately combined using the alpha blend equation alone, because one terrain layer might contain coarse terrain detail while another layer might contain fine detail. Averaging or summing normals results in a loss of accuracy.

![Terrain layers rendered with the Forward rendering path](../../../uploads/urp/rendering-deferred/terrain-layers-forward.png)

Terrain layers rendered with the Forward rendering path

![Terrain layers rendered with the Deferred rendering path](../../../uploads/urp/rendering-deferred/terrain-layers-deferred.png)

Terrain layers rendered with the Deferred rendering path

## Default shader compatibility

Unity uses a forward render pass to render the following default **shaders**:

* Complex Lit: The lighting model is too complex to fit into the G-buffer.
* Baked Lit or Lit: There’s no realtime lighting, so Unity renders the data into the Emissive/GI/Lighting field of the G-buffer. This is quicker than using a deferred render pass.

## Additional resources

* [Introduction to rendering paths in URP](../rendering-paths-introduction-urp.md)
* [Render passes in the Deferred rendering path](render-passes-deferred.md)
* [Make a shader compatible with Deferred rendering path](make-shader-compatible-with-deferred.md)