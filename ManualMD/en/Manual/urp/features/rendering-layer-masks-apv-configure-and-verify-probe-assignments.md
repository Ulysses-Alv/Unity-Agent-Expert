# Configure APV rendering layer masks and verify probe assignments

Reduce light leaks between adjacent spaces by assigning rendering layers to renderers, enabling APV rendering **layer masks**, and verifying probe assignments.

Assign rendering layers to large **GameObjects** or to GameObjects where light leaks occur. During baking, Unity assigns a rendering layer to each probe based on the rendering layers hit by its rays.

For example, in a building where most objects use the Interior rendering layer, assign nearby probes to the Interior layer. The same applies to exterior areas. This reduces light leaks because the system samples probes based on the rendering layer you assigned.

## Assign rendering layers to renderers

To assign rendering layers to your renderers:

1. Select **Edit** > **Project Settings**.
2. In the **Tags and Layers** tab, expand the **Rendering layers** section.
3. Check existing rendering layers or create new ones.
4. Select a GameObject in your **scene**.
5. In the GameObject ****Inspector****, go to **Additional Settings**.
6. Select a rendering layer from the ****Rendering Layer Mask**** dropdown.

This assigns a rendering layer to the GameObject’s **Mesh** Renderer.

## Enable rendering layer masks and rebake lighting

To enable rendering layer masks and rebake lighting:

1. Select **Window** > **Rendering** > **Lighting** > **Adaptive Probe Volumes**.
2. Set the URP asset’s ****Light Probe** System** to **Adaptive Probe Volumes**.
3. Expand the **Rendering layers** section.
4. Enable **Rendering Layer Masks**.
5. Add the rendering layer masks Unity must consider for APVs.

   You can rename the masks you select.

   You can use up to four rendering layer masks for APV. These masks define up to four zones the system uses to reduce light leaks. Each selected mask is a bitmask built with the bitwise OR operator. You can assign multiple rendering layers to each APV mask.

   In the following example:

   * APV uses two masks (zones) from the four available to reduce light leaks.
   * APV mask 1 includes the Exterior and Upper rendering layers.
   * APV mask 2 includes the Interior, Lower, and Upper rendering layers.
   * A **Mesh Renderer** assigned only to the Exterior rendering layer primarily samples APV mask 1.
   * A Mesh Renderer assigned only to the Interior or Lower rendering layer primarily samples APV mask 2.
   * A Mesh Renderer assigned to the Upper rendering layer samples both APV mask 1 and mask 2, because this layer belongs to both APV masks.
   * A Mesh Renderer with no rendering layer assigned samples both APV mask 1 and mask 2.

   | Available rendering layers | Mask 1 | Mask 2 |
   | --- | --- | --- |
   | Interior | 0 | 1 |
   | Exterior | 1 | 0 |
   | Lower | 0 | 1 |
   | Upper | 1 | 1 |
6. Select **Generate Lighting** to bake lights again.

After you rebake, Unity applies the selected APV rendering layer masks and updates probe assignments.

## Verify probe assignments in the Rendering Debugger

To view which layers contain light probes:

1. Select **Window** > **Analysis** > **Rendering Debugger** > **Probe Volumes**.
2. Select **Display probes** in the **Probe visualization** section.
3. In the **Probe Shading Mode** dropdown, select **Rendering Layer Masks**.

You can view the different areas that contain light probes at bake time.

![The Rendering Debugger displays Rendering Layers across both surfaces and probes. Green spheres fill the exterior on the left, while yellow spheres indicate probes inside a box on the right. ](../../../uploads/Main/RenderingLayersEnabled_Probes.jpg)

The Rendering Debugger displays Rendering Layers across both surfaces and probes. Green spheres fill the exterior on the left, while yellow spheres indicate probes inside a box on the right.

![A dark, unlit interior of a box on the right and a simple gradient sky and ground on the left. Gray probe spheres are scattered across the four areas defined by the horizontal ground line and the vertical box plane. At a sampling point for a pixel inside the box, Unity evaluates eight surrounding probes. The four lower probes are embedded in the ground, making them invalid, and Unity never samples them. Among the valid upper probes, rendering layers limit sampling to those that match the layer mask. The two upper-right probes assigned to the interior layer are selected.](../../../uploads/Main/RenderingLayersEnabled_Probes-1.jpg)

A dark, unlit interior of a box on the right and a simple gradient sky and ground on the left. Gray probe spheres are scattered across the four areas defined by the horizontal ground line and the vertical box plane. At a sampling point for a pixel inside the box, Unity evaluates eight surrounding probes. The four lower probes are embedded in the ground, making them invalid, and Unity never samples them. Among the valid upper probes, rendering layers limit sampling to those that match the layer mask. The two upper-right probes assigned to the interior layer are selected.

![A dark, unlit interior of a box on the right and a simple gradient sky and ground on the left. Gray probe spheres are scattered across the four areas defined by the horizontal ground line and the vertical box plane. At a sampling point for a pixel outside the box, Unity evaluates eight surrounding probes. The four lower probes are embedded in the ground, making them invalid, and Unity never samples them. Among the valid upper probes, rendering layers limit sampling to those that match the layer mask. The two upper-left probes assigned to the exterior layer are selected.](../../../uploads/Main/RenderingLayersEnabled_Probes-2.jpg)

A dark, unlit interior of a box on the right and a simple gradient sky and ground on the left. Gray probe spheres are scattered across the four areas defined by the horizontal ground line and the vertical box plane. At a sampling point for a pixel outside the box, Unity evaluates eight surrounding probes. The four lower probes are embedded in the ground, making them invalid, and Unity never samples them. Among the valid upper probes, rendering layers limit sampling to those that match the layer mask. The two upper-left probes assigned to the exterior layer are selected.

## Additional resources

* [Rendering Debugger Material properties](rendering-debugger-reference.md#material)
* [Probe Volumes Options Override reference](../probevolumes-options-override-reference.md)
* [New lighting features and workflows in Unity 6](https://www.youtube.com/watch?v=IpVuIZYFRg4&t=1078s)