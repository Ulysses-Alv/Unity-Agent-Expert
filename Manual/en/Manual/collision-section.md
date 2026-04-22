# Collision

To configure **collision** between GameObjects in Unity, you need to use colliders. Colliders define the shape of a **GameObject** for the purposes of physical collisions. You can then use these colliders to manage collision events. You can configure collisions via **collider** components, or their corresponding C# class.

This documentation describes how to configure collisions and collision events, and how colliders interact with each other and their environment.

| **Topic** | **Description** |
| --- | --- |
| [Introduction to collision](CollidersOverview.md) | Overview of the fundamental concepts around physics collision in Unity. |
| [Introduction to collider types](collider-types-introduction.md) | The different collider types (static, kinematic, and dynamic), and how collider behaviour differs depending on the collider’s physics body configuration. |
| [Collider shapes](collider-shapes.md) | The different collider shapes available, and how collider shape complexity affects performance. |
| [Collider surfaces](collider-surfaces.md) | How PhysX handles friction and bounciness on a collider’s surface, and how to configure surface properties for each collider. |
| [Collider interactions and events](collider-interactions.md) | How collisions can call events and functions to trigger changes at run time. |
| [Collision detection](collision-detection.md) | How PhysX detects collisions in Unity, and how to select the right algorithm depending on your collider configuration for optimal performance. |