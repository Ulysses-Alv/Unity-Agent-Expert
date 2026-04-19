# Add collision to trees

Colliders define the shape of an object and are used to calculate physical **collisions**. You can add a **collider** to a tree asset in the Tree Editor or to a SpeedTree asset in the SpeedTree Modeler.

## Set up the terrain to support tree colliders

On the **Terrain** tile’s ****Inspector**** window, go to ****Terrain Collider**** and make sure **Enable Tree Colliders** is selected.

**Note**: This option is enabled by default. If you disable it for one tile, other tiles still enable it, including any new tiles you add.

## Add a collider to Tree Editor trees

To add a **Capsule Collider** to a tree asset, add it to the tree’s **prefab**:

1. In the **Hierarchy** window, click **>** next to the tree GameObject to open its prefab.
2. In the prefab’s **Inspector** window, select **Add Component** > **Physics** > **Capsule Collider** to add the collider.
3. For information on how to configure the collider, refer to [Capsule Collider component reference](class-CapsuleCollider.md).

To return to the **scene**, click **<** next to the prefab’s name.

The collider is added to the tree prefab, so you can now access it in the **Inspector** window of the tree **GameObject**. If you’ve already added the tree to the scene using the **Paint Trees** tool, the collider is added to all instances of the tree.

## Add a collider to SpeedTree trees

If you created trees in the SpeedTrees Modeler with collision objects, the Unity Editor accounts for the collision objects when it calculates colliders on the terrain.

For more information, refer to the SpeedTree [Collision object](https://docs9.speedtree.com/modeler/doku.php?id=collision_object) documentation.

## Additional resources

* [Terrain Collider component reference](class-TerrainCollider.md)
* [Capsule Collider component reference](class-CapsuleCollider.md)
* [Physics Collision](collision-section.md)