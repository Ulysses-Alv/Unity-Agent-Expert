# Set the sort point of a sprite

This property is only available when the **Sprite** Renderer’s **Draw Mode** is set to **Simple**.

In a 2D project, the Main **Camera** is set to [Orthographic Projection mode](../../class-Camera.md) by default. In this mode, Unity renders sprites in the order of their distance to the camera, along the direction of the Camera’s view.

![Orthographic Camera: Side view (top) and Game view (bottom)](../../../uploads/Main/2DSpriteRenderer_SortPoint.png)

**Orthographic Camera:** Side view (top) and Game view (bottom)

By default, a Sprite’s **Sort Point** is set to its **Center**, and Unity measures the distance between the camera’s Transform position and the Center of the Sprite to determine their render order.

To set to a different **Sort Point** from the Center, select the **Pivot** option. Edit the Sprite’s Pivot position in the Sprite Editor. For more information, refer to [Sprite Editor tab reference for the Sprite Editor window](../sprite-editor/sprite-editor-window-reference.md).

---

SpriteRenderer