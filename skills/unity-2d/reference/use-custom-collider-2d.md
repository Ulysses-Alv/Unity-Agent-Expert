# Use a Custom Collider 2D

When assigning `PhysicsShape2D` to the Custom **Collider** 2D, you can do so either during Edit mode or Play mode. When modifying the Custom Collider 2D in Edit mode, Unity saves all assigned `PhysicsShape2D` and associated vertices in the Unity **Scene**, and the `CustomCollider2D` retains its configuration when the Scene is loaded. This makes it possible to use Edit mode authoring **scripts** to create custom geometry.

When modifying the Custom Collider 2D during Play mode, the Collider will not retain any changes made to assigned `PhysicsShape2D` and their associated vertices, and all changes will be lost when exiting Play mode.