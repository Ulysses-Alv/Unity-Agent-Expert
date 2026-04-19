# Disable and strip the physics integration from your project

Disable and strip the physics integration SDK from your project to meet specific performance or feature requirements.

Unity supports a number of versatile and performant [physics engine integrations](physics-integrations.md). However, if your project doesn’t require physics simulation, you can disable and even remove the physics integration entirely from your project. This can result in a smaller build size for certain build platforms and more optimized resource usage. When you do, all physics-related functionality, such as physics APIs and components, no longer work in your project.

If you enabled **Strip Engine Code** in **Project settings**, you must set ****GameObject** SDK** in the **Physics** Project Settings **None** to strip the code of all built-in physics integrations alongside other engine code.

To strip the physics integration from your project:

1. Select **Edit > Project Settings** to open the Project Settings window.
2. Select the **Physics** tab.
3. In the **GameObject SDK** dropdown, select **None**.
4. Restart the Editor for your change to take effect.

## Additional resources

* [Physics](PhysicsSection.md)
* [Physics integrations in Unity](physics-integrations.md)
* [Physics project settings](class-PhysicsManager.md)