# Configure collisions between LowLevelPhysics2D API objects

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

In the `LowLevelPhysics2D` API, **collisions** are enabled by default.

By default the following applies:

* Objects use [GameObject layers](../Layers.md), and all objects are on the built-in layer called **Default**.
* Objects collide with objects on all other layers.
* You can use up to 32 layers.

To change this behaviour, do the following:

1. Set the `LowLevelPhysics2D` API to use its own set of 64 layers, instead of **GameObject** layers.
2. Configure which layers objects are on, and create **layer masks** that set which layers an object interacts with.

**Note:** To configure how a physics object reacts to collisions, for example how bouncy it is, refer to [Configure LowLevelPhysics2D API scenes with definitions](2d-physics-api-definitions.md).

## Set the API to use its own set of layers

Follow these steps:

1. Create a Physics Low Level Settings 2D asset. For more information, refer to [Configure global 2D physics settings](2d-physics-api-global-settings.md).
2. In the **Layers** section of the Physics Low Level Settings 2D asset, enable **Use Full Layers**.

The `PhysicsMask` API now uses the layers in the **Physics Layer Names** section of the Physics Low Level Settings 2D asset.

Use the **Physics Layer Names** section to add, edit, or remove layers. For more information, refer to [Physics Low Level Settings window reference](2d-physics-api-reference-low-level-settings).

## Change which objects collide in the Editor

Follow these steps:

1. Create the layers you need. For more information, refer to [Physics Low Level Settings window reference](2d-physics-api-reference-low-level-settings).
2. [Create a physics object](2d-physics-api-physics-object.md) with a public `PhysicsShapeDefinition` and attach it to a GameObject.
3. Select the GameObject, then in the ****Inspector**** window open the **Contact Filter** dropdown.
4. Set **Categories** to the layers you want the object to be on. For example, the first built-in layer called **Default**.
5. Set **Contacts** to the layers you want the object to collide with. For example **Nothing** for no layers.

## Change which objects collide in a script

Follow these steps:

1. Create the layers you need. For more information, refer to [Physics Low Level Settings window reference](2d-physics-api-reference-low-level-settings).
2. Create a `PhysicsMask` object with the layer you want the object to be on. For example:

   ```
   PhysicsMask objectLayer = PhysicsLayers.GetLayerMask("Car");
   ```

   `PhysicsLayers.GetLayerMask` gets the layer mask whether you’re using GameObject layers or the `LowLevelPhysics2D` API layers.
3. Create another `PhysicsMask` object with the layers you want the object to collide with.

   ```
   PhysicsMask objectLayer = PhysicsLayers.GetLayerMask("Walls");
   ```
4. Create a `ContactFilter` object with the two physics layer masks. For example:

   ```
   PhysicsShape.ContactFilter myContactFilter = new PhysicsShape.ContactFilter { categories = objectLayer, contacts = contactLayer },
   ```

   You can also use `PhysicsMask.All` to represent all layers, or `PhysicsMask.None` to represent no layers.
5. Assign the `ContactFilter` to the `PhysicsShape` when you create it. For example:

   ```
   PhysicsShapeDefinition shapeDefinition = new PhysicsShapeDefinition{
       contactFilter = myContactFilter
   };
   ```

If `PhysicsMask` is a bitmask rather than a set of layers, add the `PhysicsMask.ShowAsPhysicsMask` attribute. Unity then displays the mask as bit values.

For more information, refer to [`PhysicsMask`](../../ScriptReference/LowLevelPhysics2D.PhysicsMask.md).

## Prevent objects passing through each other

Because physics bodies move in steps, fast-moving objects can sometimes pass through other objects without activating a collision. This is called tunnelling.

The 2D physics system automatically uses continuous **collision detection** to prevent tunnelling when dynamic objects approach static objects. However if a very fast-moving object still passes through another object, enable the `fastCollisionsAllowed` property of the physics body to force **continuous collision detection**. Enabling this setting can reduce performance.

## Example

The following example sets a falling circle to be on the **MyNewLayer** layer and collide only with other **MyNewLayer** layer objects. It passes through the larger circle that’s in the default layer.

To use the example:

1. Create a Physics Low Level Settings 2D asset, enable **Use Full Layers**, and add a layer called **MyNewLayer**.
2. Attach the script to a GameObject in your **scene**, and enter Play mode.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

public class ContactFilters : MonoBehaviour
{
    void Start()
    {
        CircleGeometry smallCircleShape = new CircleGeometry{ radius = 0.5f };
        CircleGeometry largeCircleShape = new CircleGeometry{ radius = 3f };

        PhysicsWorld world = PhysicsWorld.defaultWorld;

        // Create a small falling circle
        PhysicsBody object1 = world.CreateBody(new PhysicsBodyDefinition
        { 
            position = new Vector2(0.5f, 8f),
            type = PhysicsBody.BodyType.Dynamic
        });

        // Create a shape on the MyNewLayer layer, which only collides with the MyNewLayer layer
        PhysicsMask objectLayer = PhysicsLayers.GetLayerMask("MyNewLayer");
        PhysicsMask collisionLayer = PhysicsLayers.GetLayerMask("MyNewLayer");
        PhysicsShape.ContactFilter contactFilter = new PhysicsShape.ContactFilter { 
            categories = objectLayer,
            contacts = collisionLayer
        };
        PhysicsShapeDefinition smallCircleDefinition = new PhysicsShapeDefinition{
            contactFilter = contactFilter
        };
        object1.CreateShape(smallCircleShape, smallCircleDefinition);

        // Create a larger static circle below, on the default layer
        PhysicsBody object2 = world.CreateBody(new PhysicsBodyDefinition
        { 
            position = new Vector2(0f, 0f),
            type = PhysicsBody.BodyType.Static
        });
        object2.CreateShape(largeCircleShape);

    }
}
```

## Additional resources

* [Introduction to collision in the LowLevelPhysics2D API](2d-physics-api-interactions-introduction.md)
* [Detect collisions between LowLevelPhysics2D API objects](2d-physics-api-collision-handle.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub
* [Body definition reference for the LowLevelPhysics2D API](2d-physics-api-reference-body.md)
* [Physics Low Level Settings 2D asset Inspector window reference](../class-PhysicsLowLevelSettings2D.md)