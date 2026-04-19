# Create a compound collider

A compound **collider** is a collection of collider GameObjects on the same parent **Rigidbody** **GameObject**. For details on when and why you might use a compound collider, see [Introduction to compound colliders](compound-colliders-introduction.md).

A compound collider is made up of a parent GameObject which has a Rigidbody component, and child GameObjects that have colliders.

## Plan a compound collider

Before you build your compound collider, think about what you want to use the collider for, and how you want to arrange the colliders.

* If you only need the collider to provide collisions, and not any **collision** or trigger events, you can arrange colliders in any arrangement, including overlaps, as long as you have covered all of the space that should act as a collider at run time.
* If you need the collider to call collision or trigger events, you need to arrange your compound collider in such a way that ensures there are no overlaps. You can also plan to implement [Tags](Tags.md) and Layers(Layers) to ensure that only specific colliders call collision and trigger events.
* If you need to detect exactly which part of the item is involved in a collision, separate the item regions by collider. For example, on a treasure chest that opens when a player touches the front of it, you might have one collider on the front and one on the back.
* If another collider needs to slide along your compound collider (for example, a character sliding across ice), try to have just one collider along that surface, to reduce jittering during transitions from one collider contact to another.

There is no one perfect way to arrange a compound collider; the efficiency always depends on the shape, the desired behavior, and the other elements in your project. For this reason, you should always run tests against your compound collider to check that it behaves as expected, and you should use the [Physics Profiler](ProfilerPhysics.md) to test different arrangements and configurations for computational efficiency.

## Build a compound collider

1. Create or choose the parent GameObject. In most cases, this is the GameObject that contains the **Mesh** and **Mesh Renderer**.
2. Add a Rigidbody to the parent GameObject, and configure it as needed for your project (see [Introduction to rigid body physics](RigidbodiesOverview.md)).
3. Create an empty GameObject as a child of the parent GameObject.
   1. Right-click on the parent GameObject, and select **Create Empty GameObject**.
4. Add a collider to the new empty GameObject.
   1. In the **Inspector** window, select **Add Component**.
   2. Choose a collider shape. You can use any collider shape on a compound collider. In most cases, you should aim to choose the simplest shape that can most accurately represent the collider you are building.
5. Position the new collider.
   1. Use the [Transform](class-Transform.md) or the [positioning shortcuts](PositioningGameObjects.md) to position the collider.
6. Test and observe the Rigidbody’s behavior. Changes to collider configuration can change the Rigidbody’s **center of mass**, which can result in some unexpected behavior.
7. Repeat steps 4–6 for as many colliders as you need.

If you need to apply the same compound collider to multiple GameObjects, you can duplicate the parent GameObject or use [Prefabs](Prefabs.md).

## Automatically generate compound colliders

Several third-party tools are available on the [Asset Store](https://assetstore.unity.com/) that can automatically generate a compound collider based on the Mesh of the GameObject. These tools are useful and can save time, but their output still requires testing and might require some adjustment to be fully efficient. You should apply the same level of testing to them as you would to a collider that you have built manually.