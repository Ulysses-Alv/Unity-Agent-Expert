# Enable accurate G-buffer normals in the Deferred and Deferred+ rendering path in URP

In the Deferred and Deferred+ **rendering paths**, Unity stores normals in the G-buffer.

By default, Unity encodes each normal in the RGB channel of a normal texture, using 8 bits each for x, y and z. The values are quantized with the loss of accuracy. This increases performance, especially on mobile GPUs, but might lead to color banding artifacts on smooth surfaces.

To improve the quality of the normals, you can enable the **Accurate G-buffer normals** property in the Universal Renderer asset. Follow these steps:

1. In the **Project** window, select the Universal Renderer asset.
2. In the ****Inspector**** window, go to the **Renderer** section.
3. Set **Accurate G-buffer normals** to **On**.

When you set **Accurate G-buffer normals** to **On**, Unity uses octahedron encoding. The values of normal vectors are more accurate, but the encoding and decoding operations put extra load on the GPU. The precision of the encoded normal vectors is similar to the precision of the sampled values in the **Forward rendering** path.

The following illustration shows the visual difference between the two options when the **Camera** is very close to the **GameObject**:

![Accurate G-buffer normals, visual difference between the two options.](../../../uploads/urp/rendering-deferred/difference-accurate-g-buffer-normals-on-off.png)

Accurate G-buffer normals, visual difference between the two options.

This option does not support the following:

* Decal normal blending when used with the [Screen Space decal technique](../renderer-feature-decal.md).
* **Terrain** blending with more than four Terrain layers, because normals in different layers encoded using octahedron encoding cannot be blended together because of the bitwise nature of the encoding (2 x 12 bits).

## Additional resources

* [G-buffer layout in the Deferred rendering path](g-buffer-layout.md)
* [Normals](../../StandardShaderMaterialParameterNormalMapSurfaceNormals.md)