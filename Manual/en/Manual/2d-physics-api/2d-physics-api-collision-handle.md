# Detect collisions between LowLevelPhysics2D API objects

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

After you [configure collisions](2d-physics-api-collisions-enable.md), to detect when collisions occur and respond to them, use **collision** callback methods.

Follow these steps:

1. To enable collision callback methods, set the `autoContactCallbacks` property of the [physics world](2d-physics-api-world.md) to `true`.
2. Add `PhysicsCallbacks.IContactCallback` to the list of implemented interfaces in your script. For example:

   ```
   public class DetectCollisions : MonoBehaviour, PhysicsCallbacks.IContactCallback
   {
       ...
   }
   ```
3. Implement the `OnContactBegin2D` and `OnContactEnd2D` methods of the interface, and include the code you want to run when a collision starts and ends.
4. To activate the callback methods when a collision occurs, set the `contactEvents` property of the shape to `true` and the `callbackTarget` property to `this`.

For example, the following code creates a circle that logs when it collides with a larger circle:

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

// Add the PhysicsCallbacks.IContactCallback interface
public class DetectCollisions : MonoBehaviour, PhysicsCallbacks.IContactCallback
{
    void Start()
    {
        PhysicsWorld world = PhysicsWorld.defaultWorld;
        
        // Set the world to use automatic contact callbacks
        // Don't do this in production code. Configure the setting using a Physics Low Level Settings 2D asset instead.
        world.autoContactCallbacks = true;

        // Create a small falling circle
        PhysicsBody object1 = world.CreateBody(new PhysicsBodyDefinition
        { 
            position = new Vector2(0.5f, 8f),
            type = PhysicsBody.BodyType.Dynamic
        });
        PhysicsShape object1shape = object1.CreateShape(CircleGeometry.defaultGeometry);

        // Create a larger static circle below
        PhysicsBody object2 = world.CreateBody(new PhysicsBodyDefinition
        { 
            position = new Vector2(0f, 0f),
            type = PhysicsBody.BodyType.Static
        });
        object2.CreateShape(new CircleGeometry { 
            radius = 3f,
        });

        // Set object 1 to activate collisions
        object1shape.contactEvents = true;  
        object1shape.callbackTarget = this;
    }

    // Log when the small circle collides
    public void OnContactBegin2D(PhysicsEvents.ContactBeginEvent eventData)
    {
        var contact = eventData.contactId.contact;
        Debug.Log("Collision started between shapes: " + contact.shapeA + " and " + contact.shapeB);
    }

    // Log when the small circle stops colliding
    public void OnContactEnd2D(PhysicsEvents.ContactEndEvent eventData)
    {
        var contact = eventData.contactId.contact;
        Debug.Log("Collision ended between shapes: " + contact.shapeA + " and " + contact.shapeB);
    }
}
```

## Detect a trigger overlap

A trigger shape doesn’t collide with other **colliders**. Instead, other colliders pass through it. For more information, refer to [Introduction to collision in the LowLevelPhysics2D API](2d-physics-api-interactions-introduction.md).

To detect when a trigger reports an overlap with another shape, follow these steps instead:

1. To enable trigger callback methods, set the `autoTriggerCallbacks` property of the physics world to `true`.
2. Add `PhysicsCallbacks.ITriggerCallback` to the list of implemented interfaces in your script.
3. Implement the `OnTriggerBegin2D` and `OnTriggerEnd2D` methods of the interface, and include the code you want to run when contact starts and ends.
4. To activate the callback methods when contact occurs, on both shapes set the `triggerEvents` property of the shape to `true` and the `callbackTarget` property to `this`.
5. To enable the trigger, on the trigger shape, set the `isTrigger` property to `true`.

For an example, refer to the `PhysicsShapeTriggerCallback` example in the [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub.

## Get contacts without a callback

To get information about contact between objects without a callback, use the event APIs of the world instance. For example, to get collisions that have begun, use the `PhysicsWorld.contactBeginEvents` method. For more information, refer to [`PhysicsWorld`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.md).

**Important**: This approach is less safe because the methods return a `ReadOnlySpan` type that points directly to memory. For more information, refer to [Destroy LowLevelPhysics2D API objects and manage memory](2d-physics-api-destroy.md).

## Additional resources

* [Introduction to collision in the LowLevelPhysics2D API](2d-physics-api-interactions-introduction.md)
* [Configure collisions between LowLevelPhysics2D API objects](2d-physics-api-collisions-enable.md)
* [Check LowLevelPhysics2D API intersections](2d-physics-api-raycasting.md)
* [Body definition reference for the LowLevelPhysics2D API](2d-physics-api-reference-body.md)
* [Physics Low Level Settings 2D asset Inspector window reference](../class-PhysicsLowLevelSettings2D.md)