# Enable Rendering Layers for Decals in URP

To enable Rendering Layers for decals in your project:

1. In the **Project** window, select a Renderer asset with a [Decal Renderer Feature](../renderer-feature-decal.md).
2. In the Decal Renderer Feature, enable **Use Rendering Layers**.

When you enable Rendering Layers for Decals, Unity shows the **Rendering Layers** property on each Decal Projector.

## Example

This section describes how to configure the following application example:

* The **scene** contains a Decal Projector.
* The Decal Projector projects a decal on the wall and the ground, but not on the paint bucket.

The following illustration shows the example:

![In image 1, the paint bucket has the Receive decals layer selected. In image 2 it does not, so the Decal Projector does not project on the bucket.](../../../uploads/urp/lighting/rendering-layers/rendering-layers-decal-example.png)  
*In image `1`, the paint bucket has the `Receive decals` layer selected. In image `2` it does not, so the Decal Projector does not project on the bucket.*

To implement the example:

1. Enable Rendering Layers for Decals in your project, as described above.
2. [Create a Decal Projector](../renderer-feature-decal-create.md) in the scene.
3. Go to **Project Settings** > **Tags and Layers**. Add a Rendering Layer called `Receive decals`.
4. Select the Decal Projector. In the Rendering Layers property, select `Receive decals`.
5. Select the paint bucket **GameObject**. In the **Rendering **Layer Mask**** field, clear the `Receive decals` layer. Now the Decal Projector does not affect this GameObject.

## Additional resources

* [`RenderingLayerMask` API](../../../ScriptReference/RenderingLayerMask.md)