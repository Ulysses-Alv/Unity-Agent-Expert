# Add a sprite to a LowLevelPhysics2D API object

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

After you [create a physics object](2d-physics-api-physics-object.md), to add a [sprite](../sprite/sprite-landing.md), add a Sprite Renderer component. Follow these steps:

1. In the **Hierarchy** window, select the **GameObject** with the `LowLevelPhysics2D` API script.
2. Select **Add Component**.
3. Select **Rendering** > **Sprite Renderer**.

You can also drag a sprite asset in the **Project** window to the GameObject in the **Hierarchy** window. This approach creates a new GameObject as a child of the physics body GameObject.

By default the sprite doesn’t move with the physics object. To make the sprite move, refer to [Move a GameObject with the physics API](2d-physics-api-move-gameobject.md).

## Additional resources

* [Sprites](../sprite/sprite-landing.md)
* [Sprite Renderer](../sprite/renderer/sprite-renderer-reference.md)